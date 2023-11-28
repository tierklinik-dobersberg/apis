package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/sebest/xff"
	"github.com/tierklinik-dobersberg/apis/pkg/data"
	"github.com/tierklinik-dobersberg/apis/pkg/log"
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

func getLocalAddresses() (map[string][]string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	result := make(map[string][]string, len(ifaces))

	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			log.L(context.TODO()).Errorf("failed to get IP addresses for interface %s", iface.Name)

			continue
		}

		for _, a := range addrs {
			switch v := a.(type) {
			case *net.IPAddr:
				if v.IP.To4() == nil {
					log.L(context.TODO()).Infof("skipping IPv6 address %s", v.IP)

					continue
				}

				if xff.IsPublicIP(v.IP) {
					log.L(context.TODO()).Infof("skipping public IP address %s", v.IP)

					continue
				}

				ones, _ := v.IP.DefaultMask().Size()
				result[iface.Name] = append(result[iface.Name], fmt.Sprintf("%s/%d", v.IP, ones))

			case *net.IPNet:
				if v.IP.To4() == nil {
					log.L(context.TODO()).Infof("skipping IPv6 address %s", v.IP)

					continue
				}

				if xff.IsPublicIP(v.IP) {
					log.L(context.TODO()).Infof("skipping public IP address %s", v.IP)

					continue
				}

				ones, _ := v.Mask.Size()
				result[iface.Name] = append(result[iface.Name], fmt.Sprintf("%s/%d", v.IP, ones))

			default:
				log.L(context.TODO()).Warnf("unsupported interface address type %T", a)
			}
		}
	}

	return result, nil
}

// WithTrustedProxies adds a HTTP middleware that extracts the
// real client IP address from X-Forwareded-For headers.
func WithTrustedProxies(networks []string) CreateOption {
	return func(srv *http.Server) error {
		allIfaces, err := getLocalAddresses()
		if err != nil {
			return fmt.Errorf("failed to enumerate interace addresses: %w", err)
		}

		nets := make([]string, 0, len(networks))

		for _, n := range networks {
			ifaceAddr, ok := allIfaces[n]
			if ok {
				// this is an interface name
				nets = append(nets, ifaceAddr...)

				continue
			}

			if n == "local" {
				// all iface addresses
				for iface, addrs := range allIfaces {
					log.L(context.TODO()).Infof("adding addresses from interface %s as trusted networks: %v", iface, addrs)

					nets = append(nets, addrs...)
				}

				continue
			}

			nets = append(nets, n)
		}

		// remove all duplicates
		nets = data.MapToSlice(
			data.IndexSlice(nets, func(n string) string { return n }),
		)

		log.L(context.TODO()).Infof("trusted networks for X-Forwarded-For headers: %v", nets)

		xffHandler, err := xff.New(xff.Options{
			AllowedSubnets: nets,
			Debug:          true,
		})
		if err != nil {
			return err
		}

		prevHandler := srv.Handler

		xffHandlerFunc := xffHandler.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		if os.Getenv("DEBUG") != "" {
			srv.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				log.L(r.Context()).Infof("RemoteAddr: %s | X-Forwarded-For: %s", r.RemoteAddr, r.Header.Get("X-Forwarded-For"))

				xffHandlerFunc.ServeHTTP(w, r)
			})
		} else {
			srv.Handler = xffHandlerFunc
		}

		return nil
	}
}
