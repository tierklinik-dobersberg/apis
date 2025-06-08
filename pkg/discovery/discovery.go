package discovery

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"
)

const (
	MetaKeyTransport = "transport"
	MetaKeyHTTPPath  = "http:path"
)

const (
	TransportHTTP              = "http" // default if unset
	TransportH2C               = "h2c"
	TransportHTTPS             = "https"
	TransportHTTPSWithInsecure = "https+insecure"
)

type ServiceInstance struct {
	// Name is the name of the service
	Name string

	// Instance is a unique instance id used to differentiate between
	// multiple instances of the same service
	Instance string

	// Meta may hold additional service metadata.
	Meta map[string]string

	// Address its the address at which the service can be reached.
	Address string
}

type Discoverer interface {
	// Register registers a new service instance at the service registry.
	Register(ctx context.Context, instance ServiceInstance) error

	// Deregister removes a service instance registration from the service
	// registry.
	Deregister(ctx context.Context, instance ServiceInstance) error

	// MarkHealth marks a service instance as healthy.
	// Discoverer implementations should ensure a service is marked as unhealthy if
	// MarkHealthy is not called within 10 seconds.
	MarkHealthy(ctx context.Context, instance ServiceInstance) error

	// Discover
	Discover(ctx context.Context, name string) ([]ServiceInstance, error)
}

// Register one or more service using discoverer and update the service instance health
// every 5 seconds.
func Register(ctx context.Context, discoverer Discoverer, instances ...*ServiceInstance) error {
	hostname, err := os.Hostname()
	if err != nil {
		return fmt.Errorf("failed to get hostname: %w", err)
	}

	for _, instance := range instances {
		parts := strings.Split(instance.Address, ":")
		if len(parts) != 2 {
			return fmt.Errorf("%s: expected ServiceInstance.Address to be <host>:<port>", instance.Instance)
		}

		if parts[0] == "" {
			var err error

			parts[0], err = os.Hostname()
			if err != nil {
				return fmt.Errorf("%s: failed to get hostname: %w", instance.Instance, err)
			}
		}

		// Ensure there's a valid instance-ID
		if instance.Instance == "" {
			instance.Instance = hostname
		}

		instance.Address = strings.Join(parts, ":")

		if err := discoverer.Register(ctx, *instance); err != nil {
			return fmt.Errorf("%s: failed to register service instance: %w", instance.Instance, err)
		}

		go func() {
			ticker := time.NewTicker(5 * time.Second)
			defer ticker.Stop()

			defer func() {
				if err := discoverer.Deregister(ctx, *instance); err != nil {
					slog.Error("failed to deregister service from catalog", "error", err, "instance", instance.Instance)
				}
			}()

			for {
				if err := discoverer.MarkHealthy(ctx, *instance); err != nil {
					slog.Error("failed to mark service instance as healthy", "error", err, "instance", instance.Instance)
				}

				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
				}
			}
		}()
	}

	return nil
}
