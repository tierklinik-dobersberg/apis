package server

import (
	"context"
	"net"
	"net/http"
)

func WithBaseContext(ctx context.Context) CreateOption {
	return func(srv *http.Server) error {
		srv.BaseContext = func(l net.Listener) context.Context {
			return ctx
		}

		return nil
	}
}

func WithBaseContextFunc(fn func(net.Listener) context.Context) CreateOption {
	return func(srv *http.Server) error {
		srv.BaseContext = fn

		return nil
	}
}
