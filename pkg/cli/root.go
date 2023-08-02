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

		flags.StringVarP(&root.BaseURLS.Idm, "idm-url", "U", os.Getenv("IDM_URL"), "The Base URL for the  server")
		flags.StringVarP(&root.BaseURLS.Calendar, "cal-url", "U", os.Getenv("CAL_URL"), "The Base URL for the  server")
		flags.StringVarP(&root.BaseURLS.Roster, "roster-url", "U", os.Getenv("ROSTER_URL"), "The Base URL for the  server")
		flags.StringVarP(&root.TokenPath, "token-file", "t", defaultTokenPath, "The path to the cached access token")
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
