package cli

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/ghodss/yaml"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	idmv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/idm/v1"
	"golang.org/x/net/http2"
)

type BaseURLS struct {
	Idm               string `json:"idm"`
	Calendar          string `json:"calendar"`
	Roster            string `json:"roster"`
	CallService       string `json:"callService"`
	CommentService    string `json:"commentService"`
	CustomerService   string `json:"customerService"`
	TaskService       string `json:"taskService"`
	OfficeHourService string `json:"officeHourService"`
	EventsService     string `json:"eventsService"`
	OrthancBridge     string `json:"orthancBridge"`
}

type Config struct {
	BaseURLS `json:"urls"`
	Debug    bool `json:"debug"`
	Verbose  bool `json:"verbose"`

	TokenPath          string `json:"tokenPath"`
	InsecureSkipVerify bool   `json:"insecure"`
	OutputYAML         bool   `json:"outputYaml"`
	ProtoFieldNames    bool   `json:"useProtoFieldNames"`
	Format             string `json:"-"`
}

type Root struct {
	*cobra.Command `json:"-"`

	ctx    context.Context
	tokens TokenFile

	ConfigurationDirectory string             `json:"-"`
	Configurations         map[string]*Config `json:"configs"`

	activeConfig string

	HttpClient connect.HTTPClient `json:"-"`
	Print      OutputFunc         `json:"-"`
	Transport  http.RoundTripper  `json:"-"`
}

func (root *Root) Config() *Config {
	cfg := root.activeConfig
	if cfg == "" {
		cfg = "default"
	}

	c, ok := root.Configurations[cfg]
	if !ok {
		return &Config{}
	}

	return c
}

func (root *Root) Debug() bool              { return root.Config().Debug }
func (root *Root) Verbose() bool            { return root.Config().Verbose }
func (root *Root) TokenPath() string        { return root.Config().TokenPath }
func (root *Root) OutputYAML() bool         { return root.Config().OutputYAML }
func (root *Root) UseProtoFieldNames() bool { return root.Config().ProtoFieldNames }
func (root *Root) Insecure() bool           { return root.Config().InsecureSkipVerify }
func (root *Root) ActiveConfig() string     { return root.activeConfig }
func (root *Root) Tokens() TokenFile        { return root.tokens }

func (root *Root) RoundTrip(req *http.Request) (*http.Response, error) {
	token := root.tokens.AccessToken

	if envToken := os.Getenv("CIS_ACCESS_TOKEN"); envToken != "" {
		logrus.Infof("using access token from environment variable CIS_ACCESS_TOKEN")
		token = envToken
	}

	if token != "" {
		req.Header.Add("Authorization", "Bearer "+token)
	}

	if root.Debug() {
		if root.Verbose() {
			blob, _ := httputil.DumpRequestOut(req, true)
			_, _ = os.Stderr.Write(blob)
		} else {
			logrus.Debugf("[http] %s %q", req.Method, req.URL.String())
		}
	}

	res, err := root.Transport.RoundTrip(req)

	if root.Debug() {
		if err != nil {
			logrus.Debugf("[http] %s %q -> error: %s", req.Method, req.URL.String(), err)
		} else {
			if root.Verbose() {
				blob, _ := httputil.DumpResponse(res, true)
				_, _ = os.Stderr.Write(blob)
			} else {
				logrus.Debugf("[http] %s %q -> status: %d", req.Method, req.URL.String(), res.StatusCode)
			}
		}
	}

	return res, err
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
		ctx:                    context.Background(),
		ConfigurationDirectory: configDir,
	}

	// try to read the configuration file
	configFile := filepath.Join(configDir, "config.yml")
	configFileContent, err := os.ReadFile(configFile)
	if err == nil {
		configFileContent, err = yaml.YAMLToJSON(configFileContent)
		if err != nil {
			logrus.Fatalf("failed to read YAML configuration: %s", err)
		}

		if err := json.Unmarshal(configFileContent, &root.Configurations); err != nil {
			logrus.Fatalf("failed to parse configuration: %s", err)
		}
	} else if !os.IsNotExist(err) {
		logrus.Fatalf("failed to open configuration file: %s", err)
	}

	root.Print = root.defaultPrintFunc

	var flagConfig Config

	root.Command = &cobra.Command{
		Use: name,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cfg := root.Config()
			if cfg.TokenPath == "" {
				cfg.TokenPath = filepath.Join(
					configDir,
					root.activeConfig+"-tokens.json",
				)
			}

			// apply overwrites from command line flags
			if cmd.Flag("debug").Changed {
				cfg.Debug = flagConfig.Debug
			}
			if cmd.Flag("verbose").Changed {
				cfg.Verbose = flagConfig.Verbose
			}
			if cmd.Flag("insecure").Changed {
				cfg.InsecureSkipVerify = flagConfig.InsecureSkipVerify
			}
			if cmd.Flag("yaml").Changed {
				cfg.OutputYAML = flagConfig.OutputYAML
			}
			if cmd.Flag("proto-field-names").Changed {
				cfg.ProtoFieldNames = flagConfig.ProtoFieldNames
			}

			if cmd.Flag("format").Changed {
				cfg.Format = flagConfig.Format
			}

			if cfg.Debug {
				logrus.SetLevel(logrus.DebugLevel)
			}

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
					InsecureSkipVerify: cfg.InsecureSkipVerify,
				},
			}

			if cfg.InsecureSkipVerify {
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
		flags.BoolVar(&flagConfig.InsecureSkipVerify, "insecure", false, "Do not validate TLS certificates")
		flags.BoolVar(&flagConfig.OutputYAML, "yaml", false, "Display YAML output instead of JSON")
		flags.BoolVar(&flagConfig.Debug, "debug", false, "Enable debug mode")
		flags.BoolVar(&flagConfig.Verbose, "verbose", false, "Enable verbose output mode")
		flags.StringVarP(&root.activeConfig, "configuration", "c", "default", "Which configuration to use.")
		flags.StringVar(&flagConfig.Format, "format", "", "Use go text/template for output formatting")
		flags.BoolVar(&flagConfig.ProtoFieldNames, "proto-field-names", false, "Use protobuf field names instead of camelCase field names")
	}

	return root
}

func NewInsecureHttp2Client() *http.Client {
	return &http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, _ *tls.Config) (net.Conn, error) {
				// If you're also using this client for non-h2c traffic, you may want
				// to delegate to tls.Dial if the network isn't TCP or the addr isn't
				// in an allowlist.
				return net.Dial(network, addr)
			},
			// Don't forget timeouts!
		},
	}
}
