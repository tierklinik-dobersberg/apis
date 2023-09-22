package server

import (
	"context"
	"errors"
	"net/http"
	"time"

	servertiming "github.com/mitchellh/go-server-timing"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/sync/errgroup"
)

func Create(addr string, handler http.Handler) *http.Server {
	handler = servertiming.Middleware(handler, nil)

	srv := &http.Server{
		Addr:    addr,
		Handler: h2c.NewHandler(handler, &http2.Server{}),
	}

	return srv
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
