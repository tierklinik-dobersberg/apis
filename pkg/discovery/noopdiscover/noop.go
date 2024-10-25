package noopdiscover

import (
	"context"

	"github.com/tierklinik-dobersberg/apis/pkg/discovery"
)

type NoOpDiscoverer struct{}

func (*NoOpDiscoverer) Register(context.Context, discovery.ServiceInstance) error {
	return nil
}

func (*NoOpDiscoverer) Deregister(context.Context, discovery.ServiceInstance) error {
	return nil
}

func (*NoOpDiscoverer) MarkHealthy(context.Context, discovery.ServiceInstance) error {
	return nil
}

func (*NoOpDiscoverer) Discover(context.Context, string) ([]discovery.ServiceInstance, error) {
	return nil, nil
}

var _ discovery.Discoverer = (*NoOpDiscoverer)(nil)
