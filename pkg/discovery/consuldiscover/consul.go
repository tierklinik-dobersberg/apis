package consuldiscover

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	consul "github.com/hashicorp/consul/api"
	"github.com/tierklinik-dobersberg/apis/pkg/discovery"
	"github.com/tierklinik-dobersberg/apis/pkg/discovery/noopdiscover"
)

// Registry defines a consul based service registry
type Registry struct {
	client *consul.Client
}

// NewRegistery creates a new consul registry instance.
// Addr is the address where the consul agent is running
func NewRegistery(addr string) (*Registry, error) {
	config := consul.DefaultConfig()
	config.Address = addr
	client, err := consul.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Registry{client: client}, nil
}

// NewFromEnv tries to configure a consul discoverer from the environment.
func NewFromEnv() (discovery.Discoverer, error) {
	addr := os.Getenv("CONSUL")
	if addr == "" {
		return &noopdiscover.NoOpDiscoverer{}, nil
	}

	return NewRegistery(addr)
}

// Register creates a service record in the registry
func (r *Registry) Register(ctx context.Context, instance discovery.ServiceInstance) error {
	parts := strings.Split(instance.Address, ":")
	if len(parts) != 2 {
		return errors.New("invalid host:port format. Eg: localhost:8081")
	}
	port, err := strconv.Atoi(parts[1])
	if err != nil {
		return err
	}
	host := parts[0]

	err = r.client.Agent().ServiceRegister(&consul.AgentServiceRegistration{
		Address: host,
		Port:    port,
		ID:      instance.Instance,
		Name:    instance.Name,
		Check: &consul.AgentServiceCheck{
			CheckID: instance.Instance,
			TTL:     "10s",
		},
		Meta: instance.Meta,
	})
	return err
}

// Deregister removes a service record from the registry
func (r *Registry) Deregister(ctx context.Context, instance discovery.ServiceInstance) error {
	err := r.client.Agent().ServiceDeregister(instance.Instance)

	return err
}

// HealthCheck is a push mechanism to update the health status of a service instance
func (r *Registry) MarkHealthy(_ context.Context, instance discovery.ServiceInstance) error {

	err := r.client.Agent().UpdateTTL(instance.Instance, "", "pass")
	return err
}

// Discover returns a list of addresses of active instances of the given service
func (r *Registry) Discover(ctx context.Context, serviceName string) ([]discovery.ServiceInstance, error) {
	entries, _, err := r.client.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return nil, err
	}
	var instances []discovery.ServiceInstance
	for _, entry := range entries {
		instances = append(instances, discovery.ServiceInstance{
			Name:     entry.Service.Service,
			Instance: entry.Service.ID,
			Meta:     entry.Service.Meta,
			Address:  fmt.Sprintf("%s:%d", entry.Service.Address, entry.Service.Port),
		})
	}

	return instances, nil
}

var _ discovery.Discoverer = (*Registry)(nil)
