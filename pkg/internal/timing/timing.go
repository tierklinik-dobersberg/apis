package timing

import (
	"context"

	servertiming "github.com/mitchellh/go-server-timing"
)

func Track(ctx context.Context, name string, fn func() error) error {
	t := servertiming.FromContext(ctx)

	if t != nil {
		defer t.NewMetric(name).Start().Stop()
	}

	return fn()
}
