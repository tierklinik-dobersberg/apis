package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	servertiming "github.com/mitchellh/go-server-timing"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/sync/errgroup"
)

// CreateOption can be used to customize the server returned by Create.
type CreateOption func(srv *http.Server) error

// Create creates a new http.Server that uses h2c to support cleartext HTTP/2.
//
// Deprecated: use CreateWithOptions instead.
func Create(addr string, handler http.Handler, opts ...CreateOption) *http.Server {
	srv, err := CreateWithOptions(addr, handler)
	if err != nil {
		// This shouldn't happen as only CreateOption may return an error and there
		// aren't any.
		panic(fmt.Sprintf("failed to create server: %s", err))
	}

	return srv
}

// Create creates a new http.Server that uses h2c to support cleartext HTTP/2 and supports server
// customization using CreateOption's.
func CreateWithOptions(addr string, handler http.Handler, opts ...CreateOption) (*http.Server, error) {
	handler = servertiming.Middleware(handler, nil)

	srv := &http.Server{
		Addr:    addr,
		Handler: handler, // h2c.NewHandler(handler, &http2.Server{}),
	}

	for _, opt := range opts {
		if err := opt(srv); err != nil {
			return nil, fmt.Errorf("create option: %w", err)
		}
	}

	// finally, apply h2c middleware
	srv.Handler = h2c.NewHandler(srv.Handler, &http2.Server{})

	return srv, nil
}

type ServeAndShutdown interface {
	ListenAndServe() error
	Shutdown(context.Context) error
}

func Serve(ctx context.Context, servers ...ServeAndShutdown) error {
	errgrp, ctx := errgroup.WithContext(ctx)

	go func() {
		<-ctx.Done()

		timeOutCtx, cancel := context.WithTimeout(
			context.Background(),
			time.Minute,
		)
		defer cancel()

		for _, s := range servers {
			go s.Shutdown(timeOutCtx)
		}
	}()

	for idx := range servers {
		s := servers[idx]

		errgrp.Go(func() error {
			err := s.ListenAndServe()
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				return err
			}

			return nil
		})
	}

	if err := errgrp.Wait(); err != nil {
		return err
	}

	return nil
}
