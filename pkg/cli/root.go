package cli

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/calendar/v1/calendarv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/idm/v1/idmv1connect"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/roster/v1/rosterv1connect"
)

type BaseURLS struct {
	Idm      string
	Calendar string
	Roster   string
}

type Root struct {
	*cobra.Command

	BaseURLS

	HttpClient         connect.HTTPClient
	TokenPath          string
	insecureSkipVerify bool
	outputYAML         bool
	Print              OutputFunc
	Transport          http.RoundTripper
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

func New(name string) *Root {
	root := &Root{}

	root.Print = root.defaultPrintFunc

	root.Command = &cobra.Command{
		Use: name,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			content, err := os.ReadFile(root.TokenPath)
			if err != nil {
				if !os.IsNotExist(err) {
					return fmt.Errorf("failed to read token at %s: %w", root.TokenPath, err)
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
					InsecureSkipVerify: root.insecureSkipVerify,
				},
			}

			if root.insecureSkipVerify {
				logrus.Warn("TLS certificate validation is disabled")
			}

			var rt = root.Transport
			if len(content) > 0 {
				rt = &addHeaderRT{RoundTripper: root.Transport, token: string(content)}
			}

			root.HttpClient = &http.Client{
				Transport: rt,
			}

			return nil
		},
	}

	flags := root.PersistentFlags()
	{
		defaultTokenPath := os.Getenv("CIS_IDM_TOKEN_FILE")
		if defaultTokenPath == "" {
			defaultTokenPath = filepath.Join(os.Getenv("HOME"), ".idm-token")
		}

		flags.StringVar(&root.BaseURLS.Idm, "idm-url", os.Getenv("IDM_URL"), "The Base URL for the  server")
		flags.StringVar(&root.BaseURLS.Calendar, "cal-url", os.Getenv("CAL_URL"), "The Base URL for the  server")
		flags.StringVar(&root.BaseURLS.Roster, "roster-url", os.Getenv("ROSTER_URL"), "The Base URL for the  server")
		flags.StringVar(&root.TokenPath, "token-file", defaultTokenPath, "The path to the cached access token")
		flags.BoolVar(&root.insecureSkipVerify, "insecure", false, "Do not validate TLS certificates")
		flags.BoolVar(&root.outputYAML, "yaml", false, "Display YAML output instead of JSON")
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
