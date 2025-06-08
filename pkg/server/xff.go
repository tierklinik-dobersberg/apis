package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

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
			slog.Error("failed to get IP addresses for interface", "error", iface.Name)

			continue
		}

		for _, a := range addrs {
			switch v := a.(type) {
			case *net.IPAddr:
				if v.IP.To4() == nil {
					continue
				}

				if xff.IsPublicIP(v.IP) {
					continue
				}

				ones, _ := v.IP.DefaultMask().Size()
				result[iface.Name] = append(result[iface.Name], fmt.Sprintf("%s/%d", v.IP, ones))

			case *net.IPNet:
				if v.IP.To4() == nil {
					continue
				}

				if xff.IsPublicIP(v.IP) {
					continue
				}

				ones, _ := v.Mask.Size()
				result[iface.Name] = append(result[iface.Name], fmt.Sprintf("%s/%d", v.IP, ones))

			default:
				slog.Warn("unsupported interface address type", "type", fmt.Sprintf("%T", a))
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

		// create a slice of net.IPNet
		var ipNets []*net.IPNet
		var lock sync.RWMutex

		parseNetworks := func() error {
			lock.Lock()
			defer lock.Unlock()

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
					for _, addrs := range allIfaces {
						nets = append(nets, addrs...)
					}

					continue
				}

				// try to parse it as net.IPNet
				if _, _, err := net.ParseCIDR(n); err == nil {
					nets = append(nets, n)

					continue
				}

				if ips, err := net.LookupIP(n); err == nil {
					for _, ip := range ips {
						netStr := fmt.Sprintf("%s/32", ip)

						nets = append(nets, netStr)
					}

					continue
				}

				return fmt.Errorf("failed to determine IP address or network for %q", n)
			}

			// remove all duplicates
			nets = data.MapToSlice(
				data.IndexSlice(nets, func(n string) string { return n }),
			)

			ipNets = make([]*net.IPNet, len(nets))
			for idx, n := range nets {
				_, parsed, err := net.ParseCIDR(n)
				if err != nil {
					return fmt.Errorf("failed to parse %s: %w", n, err)
				}

				ipNets[idx] = parsed
			}

			return nil
		}

		if err := parseNetworks(); err != nil {
			return err
		}

		go func() {
			ticker := time.NewTicker(time.Minute)

			for range ticker.C {
				if err := parseNetworks(); err != nil {
					slog.Error("failed to refresh trusted networks", "error", err)
				}
			}
		}()

		isAllowed := func(ip net.IP) bool {
			lock.RLock()
			defer lock.RUnlock()

			for _, n := range ipNets {
				if n.Contains(ip) {
					return true
				}
			}

			return false
		}

		prevHandler := srv.Handler

		srv.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			var realIP net.IP

			if r.RemoteAddr != "" {
				sip, _, err := net.SplitHostPort(r.RemoteAddr)
				if err == nil {
					ip := net.ParseIP(sip)
					switch {
					case ip == nil:
						// there's nothing we can do here
					case !isAllowed(ip):

						// this seems to be the real client IP
						realIP = ip

					case isAllowed(ip):
						if xffh := r.Header.Get("X-Forwarded-For"); xffh != "" {
							for _, ip := range strings.Split(xffh, ",") {
								ip = strings.TrimSpace(ip)

								parsedIP := net.ParseIP(ip)
								if parsedIP == nil {
									continue
								}

								if isAllowed(parsedIP) {
									continue
								}

								realIP = parsedIP

								break
							}
						} else {
							realIP = ip
						}
					}
				}
			}

			if realIP != nil {
				l := log.L(ctx)

				ctx = WithRealIP(ctx, realIP)
				ctx = log.WithLogger(ctx, l.With("remoteIP", realIP.String()))

				r = r.WithContext(ctx)
			}

			prevHandler.ServeHTTP(w, r)
		})

		return nil
	}
}
