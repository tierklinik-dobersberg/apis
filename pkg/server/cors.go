package server

import (
	"net/http"

	"github.com/tierklinik-dobersberg/apis/pkg/cors"
)

func WithCORS(cfg cors.Config) CreateOption {
	return func(srv *http.Server) error {
		srv.Handler = cors.Wrap(cfg, srv.Handler)

		return nil
	}
}
