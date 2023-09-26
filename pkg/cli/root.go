package cli

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/ghodss/yaml"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/calendar/v1/calendarv1connect"
	idmv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/idm/v1"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/idm/v1/idmv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/pbx3cx/v1/pbx3cxv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/roster/v1/rosterv1connect"
)

type BaseURLS struct {
	Idm         string `json:"idm"`
	Calendar    string `json:"calendar"`
	Roster      string `json:"roster"`
	CallService string `json:"callService"`
}

type Root struct {
	*cobra.Command `json:"-"`

	ctx context.Context

	BaseURLS `json:"urls"`

	TokenPath          string `json:"tokenPath"`
	InsecureSkipVerify bool   `json:"insecure"`
	OutputYAML         bool   `json:"outputYaml"`

	HttpClient connect.HTTPClient `json:"-"`
	Print      OutputFunc         `json:"-"`
	Transport  http.RoundTripper  `json:"-"`
}

type TokenFile struct {
	AccessToken   string `json:"accessToken"`
	RefreshCookie string `json:"refreshCookie"`
}

func (root *Root) Context() context.Context {
	return root.ctx
}

func (root *Root) OffTime() rosterv1connect.OffTimeServiceClient {
	return rosterv1connect.NewOffTimeServiceClient(root.HttpClient, root.BaseURLS.Roster)
}

func (root *Root) WorkTime() rosterv1connect.WorkTimeServiceClient {
	return rosterv1connect.NewWorkTimeServiceClient(root.HttpClient, root.BaseURLS.Roster)
}

func (root *Root) WorkShift() rosterv1connect.WorkShiftServiceClient {
	return rosterv1connect.NewWorkShiftServiceClient(root.HttpClient, root.BaseURLS.Roster)
}

func (root *Root) Roster() rosterv1connect.RosterServiceClient {
	return rosterv1connect.NewRosterServiceClient(root.HttpClient, root.BaseURLS.Roster)
}

func (root *Root) Constraints() rosterv1connect.ConstraintServiceClient {
	return rosterv1connect.NewConstraintServiceClient(root.HttpClient, root.BaseURLS.Roster)
}

func (root *Root) Users() idmv1connect.UserServiceClient {
	return idmv1connect.NewUserServiceClient(root.HttpClient, root.BaseURLS.Idm)
}

func (root *Root) Roles() idmv1connect.RoleServiceClient {
	return idmv1connect.NewRoleServiceClient(root.HttpClient, root.BaseURLS.Idm)
}

func (root *Root) SelfService() idmv1connect.SelfServiceServiceClient {
	return idmv1connect.NewSelfServiceServiceClient(root.HttpClient, root.BaseURLS.Idm)
}

func (root *Root) Auth() idmv1connect.AuthServiceClient {
	return idmv1connect.NewAuthServiceClient(root.HttpClient, root.BaseURLS.Idm)
}

func (root *Root) Calendar() calendarv1connect.CalendarServiceClient {
	return calendarv1connect.NewCalendarServiceClient(root.HttpClient, root.BaseURLS.Calendar)
}

func (root *Root) CallService() pbx3cxv1connect.CallServiceClient {
	return pbx3cxv1connect.NewCallServiceClient(root.HttpClient, root.BaseURLS.CallService)
}

func New(name string) *Root {
	defaultConfigDirectory := filepath.Join(
		os.Getenv("HOME"),
		".config",
		"cis",
	)

	configDir := os.Getenv("CIS_CONFIG_DIR")
	if configDir == "" {
		configDir = defaultConfigDirectory
	}

	root := &Root{
		ctx: context.Background(),
	}

	// try to read the configuration file
	configFileContent, err := os.ReadFile(filepath.Join(configDir, "config.yml"))
	if err == nil {
		configFileContent, err = yaml.YAMLToJSON(configFileContent)
		if err != nil {
			logrus.Fatalf("failed to read YAML configuration: %s", err)
		}

		if err := json.Unmarshal(configFileContent, root); err != nil {
			logrus.Fatalf("failed to parse configuration: %s", err)
		}
	} else if !os.IsNotExist(err) {
		logrus.Fatalf("failed to open configuration file: %s", err)
	}

	// apply default from environment
	if root.TokenPath == "" {
		root.TokenPath = os.Getenv("CIS_CONFIG_TOKEN_PATH")
	}

	if root.TokenPath == "" {
		root.TokenPath = filepath.Join(
			configDir,
			"token.json",
		)
	}

	root.Print = root.defaultPrintFunc

	root.Command = &cobra.Command{
		Use: name,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			dirStat, err := os.Stat(configDir)
			if err != nil {
				if os.IsNotExist(err) {
					os.MkdirAll(configDir, 0700)
				} else {
					logrus.Fatal(err)
				}
			} else if !dirStat.IsDir() {
				logrus.Fatalf("expected %q to be a directory but found a file", configDir)
			}

			// Parse the token path
			var tokens TokenFile
			if err == nil {
				content, err := os.ReadFile(root.TokenPath)
				if err != nil {
					if !os.IsNotExist(err) {
						return fmt.Errorf("failed to read token at %s: %w", root.TokenPath, err)
					} else {
						logrus.Warnf("failed to get token: %s", err)
					}
				}

				if err := json.Unmarshal(content, &tokens); err != nil {
					logrus.Warnf("failed to parse token file: s", err)
				}
			}

			// if we have an access token we wrap the default transport
			// in a custom round-tripper that adds the authentication header
			root.Transport = &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				ForceAttemptHTTP2:     true,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: root.InsecureSkipVerify,
				},
			}

			if root.InsecureSkipVerify {
				logrus.Warn("TLS certificate validation is disabled")
			}

			var rt = root.Transport
			if tokens.AccessToken != "" {
				// try to get the current profile just to see if the access token is valid
				// TODO(ppacher): better parse the access token instead of trying to get the profile
				rt = &addHeaderRT{RoundTripper: root.Transport, token: tokens.AccessToken}

				root.HttpClient = &http.Client{
					Transport: rt,
				}

				_, err := root.Auth().Introspect(root.Context(), connect.NewRequest(&idmv1.IntrospectRequest{}))
				if err != nil {
					var cerr *connect.Error
					if errors.As(err, &cerr) && cerr.Code() == connect.CodeUnauthenticated {
						// try to refresh the token
						if tokens.RefreshCookie != "" {
							req := connect.NewRequest(&idmv1.RefreshTokenRequest{})
							req.Header().Add("Cookie", tokens.RefreshCookie)

							res, err := root.Auth().RefreshToken(root.Context(), req)
							if err == nil {
								tokens.AccessToken = res.Msg.GetAccessToken().GetToken()

								blob, err := json.Marshal(tokens)
								if err != nil {
									logrus.Fatalf("failed to marshal token file")
								}

								if err := os.WriteFile(root.TokenPath, blob, 0600); err != nil {
									logrus.Warnf("failed to write token file to %q: %s", root.TokenPath, err)
								}

								rt.(*addHeaderRT).token = tokens.AccessToken
							} else {
								logrus.Warnf("failed to refresh access token: %s", err)
							}

						} else {
							logrus.Warnf("access token expired but refreshing failed because no refresh token is available")
						}
					}
				}
			}

			return nil
		},
	}

	flags := root.PersistentFlags()
	{

		defaultTokenPath := os.Getenv("CIS_TOKEN_FILE")
		if defaultTokenPath == "" {
			defaultTokenPath = filepath.Join(configDir, "access_token")
		}

		flags.StringVar(&root.BaseURLS.Idm, "idm-url", os.Getenv("IDM_URL"), "The Base URL for the IDM server")
		flags.StringVar(&root.BaseURLS.Calendar, "cal-url", os.Getenv("CAL_URL"), "The Base URL for the Calendar server")
		flags.StringVar(&root.BaseURLS.Roster, "roster-url", os.Getenv("ROSTER_URL"), "The Base URL for the Roster server")
		flags.StringVar(&root.BaseURLS.CallService, "call-service-url", os.Getenv("CALL_SERVICE_URL"), "The Base URL for the CallService server")
		flags.StringVar(&root.TokenPath, "token-file", defaultTokenPath, "The path to the cached access token")
		flags.BoolVar(&root.InsecureSkipVerify, "insecure", false, "Do not validate TLS certificates")
		flags.BoolVar(&root.OutputYAML, "yaml", false, "Display YAML output instead of JSON")
	}

	return root
}

type addHeaderRT struct {
	http.RoundTripper

	token string
}

func (ahr *addHeaderRT) RoundTrip(req *http.Request) (*http.Response, error) {
	if ahr.token != "" {
		req.Header.Add("Authentication", "Bearer "+ahr.token)
	}
	return ahr.RoundTripper.RoundTrip(req)
}
