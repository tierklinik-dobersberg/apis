package server

import (
	"context"
	"net"
	"net/http"

	"github.com/sebest/xff"
)

var clientIPContextKey = struct{ s string }{s: "clientIPContextKey"}

// WithRealIP associates the IP address ip with a context.
func WithRealIP(ctx context.Context, ip net.IP) context.Context {
	return context.WithValue(ctx, clientIPContextKey, ip)
}

// RealIPFromContext returns the IP address of the client associated
// with ctx. This requires WithRealIP to construct an appropriate context
// value before. Use the WithTrustedProxies CreateOption to automatically
// extract XFF headers and add the client IP to the request context.
func RealIPFromContext(ctx context.Context) net.IP {
	v, _ := ctx.Value(clientIPContextKey).(net.IP)

	return v
}

// WithTrustedProxies adds a HTTP middleware that extracts the
// real client IP address from X-Forwareded-For headers.
func WithTrustedProxies(networks []string) CreateOption {
	return func(srv *http.Server) error {
		xffHandler, err := xff.New(xff.Options{
			AllowedSubnets: networks,
		})
		if err != nil {
			return err
		}

		prevHandler := srv.Handler

		srv.Handler = xffHandler.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			if r.RemoteAddr != "" {
				sip, _, err := net.SplitHostPort(r.RemoteAddr)
				if err == nil {
					ip := net.ParseIP(sip)
					if ip != nil {
						ctx = WithRealIP(ctx, ip)
					}
				} else {
					// TODO(ppacher): log the output?
				}
			}

			r = r.WithContext(ctx)

			prevHandler.ServeHTTP(w, r)
		}))

		return nil
	}
}
