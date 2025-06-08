package log

import (
	"context"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/tierklinik-dobersberg/apis/pkg/internal/timing"
)

var loggerContextKey = struct{ s string }{s: "logger"}

func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerContextKey, logger)
}

func L(ctx context.Context) *slog.Logger {
	l, _ := ctx.Value(loggerContextKey).(*slog.Logger)
	if l == nil {
		return slog.New(slog.NewTextHandler(os.Stderr, nil))
	}

	return l
}

func NewLoggingInterceptor() connect.UnaryInterceptorFunc {
	return func(uf connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, ar connect.AnyRequest) (connect.AnyResponse, error) {
			start := time.Now()

			l := L(ctx).With("method", ar.Spec().Procedure)

			ctx = WithLogger(ctx, l)

			var resp connect.AnyResponse
			err := timing.Track(ctx, strings.ReplaceAll(ar.Spec().Procedure, "/", "."), func() error {
				var err error

				resp, err = uf(ctx, ar)

				return err
			})

			l = l.With("duration", time.Since(start))

			if err != nil {
				l = l.With("error", err)
			}

			l.Info("connect request handled")

			return resp, err
		}
	}
}
