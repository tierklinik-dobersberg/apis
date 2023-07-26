package cors

import (
	"net/http"

	"github.com/rs/cors"
)

type Config struct {
	AllowedOrigins   []string
	AllowCredentials bool
	Debug            bool
}

func Wrap(cfg Config, next http.Handler) http.Handler {
	// Setup CORS middleware
	corsOpts := cors.Options{
		AllowedOrigins:   cfg.AllowedOrigins,
		AllowCredentials: cfg.AllowCredentials,
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		AllowedHeaders: []string{
			"Accept-Encoding",
			"Content-Encoding",
			"Content-Type",
			"Connect-Protocol-Version",
			"Connect-Timeout-Ms",
			"Connect-Accept-Encoding",  // Unused in web browsers, but added for future-proofing
			"Connect-Content-Encoding", // Unused in web browsers, but added for future-proofing
			"Grpc-Timeout",             // Used for gRPC-web
			"X-Grpc-Web",               // Used for gRPC-web
			"X-User-Agent",             // Used for gRPC-web
		},
		ExposedHeaders: []string{
			"Content-Encoding",         // Unused in web browsers, but added for future-proofing
			"Connect-Content-Encoding", // Unused in web browsers, but added for future-proofing
			"Grpc-Status",              // Required for gRPC-web
			"Grpc-Message",             // Required for gRPC-web
		},
		Debug: cfg.Debug,
	}

	c := cors.New(corsOpts)

	return c.Handler(next)
}
