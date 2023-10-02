package cli

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/ghodss/yaml"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	idmv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/idm/v1"
)

type BaseURLS struct {
	Idm         string `json:"idm"`
	Calendar    string `json:"calendar"`
	Roster      string `json:"roster"`
	CallService string `json:"callService"`
}

type Root struct {
	*cobra.Command `json:"-"`

	ctx    context.Context
	tokens TokenFile

	BaseURLS `json:"urls"`

	TokenPath          string `json:"tokenPath"`
	InsecureSkipVerify bool   `json:"insecure"`
	OutputYAML         bool   `json:"outputYaml"`

	HttpClient connect.HTTPClient `json:"-"`
	Print      OutputFunc         `json:"-"`
	Transport  http.RoundTripper  `json:"-"`
}

func (root *Root) RoundTrip(req *http.Request) (*http.Response, error) {
	if root.tokens.AccessToken != "" {
		req.Header.Add("Authentication", "Bearer "+root.tokens.AccessToken)
	}

	return root.Transport.RoundTrip(req)
}

func (root *Root) Context() context.Context {
	return root.ctx
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
			if err == nil {
				if err := root.readTokenFile(); err != nil {
					logrus.Warnf("failed to read token file: %s", err)
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

			root.HttpClient = &http.Client{
				Transport: root,
			}

			if root.tokens.AccessToken != "" {

				_, err := root.Auth().Introspect(root.Context(), connect.NewRequest(&idmv1.IntrospectRequest{}))
				if err != nil {
					var cerr *connect.Error
					if errors.As(err, &cerr) && cerr.Code() == connect.CodeUnauthenticated {
						if err := root.refreshToken(); err != nil {
							logrus.Warnf("failed to refresh access token: %s", err)
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
			defaultTokenPath = filepath.Join(configDir, "token.json")
		}

		flags.StringVar(&root.BaseURLS.Idm, "idm-url", root.BaseURLS.Idm, "The Base URL for the IDM server")
		flags.StringVar(&root.BaseURLS.Calendar, "cal-url", root.BaseURLS.Calendar, "The Base URL for the Calendar server")
		flags.StringVar(&root.BaseURLS.Roster, "roster-url", root.BaseURLS.Roster, "The Base URL for the Roster server")
		flags.StringVar(&root.BaseURLS.CallService, "call-service-url", root.BaseURLS.CallService, "The Base URL for the CallService server")
		flags.StringVar(&root.TokenPath, "token-file", defaultTokenPath, "The path to the cached access token")
		flags.BoolVar(&root.InsecureSkipVerify, "insecure", false, "Do not validate TLS certificates")
		flags.BoolVar(&root.OutputYAML, "yaml", false, "Display YAML output instead of JSON")
	}

	return root
}
