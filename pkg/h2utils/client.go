package h2utils

import (
	"crypto/tls"
	"log/slog"
	"net"
	"net/http"

	"github.com/tierklinik-dobersberg/apis/pkg/discovery"
	"github.com/tierklinik-dobersberg/apis/pkg/discovery/consuldiscover"
	"golang.org/x/net/http2"
)

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

// WithDiscovery updates an http.Client to use service discovery as implemented in
// the discovery package.
//
// Discovery is implemented by intercepting the RoundTrip function of http.Client.Transport
// by doing service discovery before passing the request to the original RoundTripper set.
// If Client.Transport is nil, it defaults to http.DefaultTransport.
//
// Any error that occurs during service detection is logged using slog but otherwise ignored
// and the original host name is used.
func WithDiscovery(disc discovery.Discoverer, cli *http.Client) *http.Client {
	cli.Transport = WithTransportDiscovery(disc, cli.Transport)

	return cli
}

func WithTransportDiscovery(disc discovery.Discoverer, t http.RoundTripper) http.RoundTripper {
	if disc == nil {
		var err error
		disc, err = consuldiscover.NewFromEnv()

		if err != nil {
			slog.Error("failed to get consul address", "error", err)
		}
	}

	if t == nil {
		t = http.DefaultTransport
	}

	return &discoveryTransport{
		disc: disc,
		t:    t,
	}
}

type discoveryTransport struct {
	disc discovery.Discoverer

	t http.RoundTripper
}

func (dt *discoveryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	host := req.URL.Hostname()

	if dt.disc != nil {
		// try to discover the service by name
		res, err := dt.disc.Discover(req.Context(), host)
		if err == nil && len(res) > 0 {
			r := req.Clone(req.Context())
			r.URL.Host = res[0].Address

			slog.Info("switching request host to discovered address", "original", req.URL.Host, "discovered", r.URL.Host)

			req = r
		} else {
			slog.Error("failed to discover service addresses", "host", host)
		}
	}

	return dt.t.RoundTrip(req)
}
