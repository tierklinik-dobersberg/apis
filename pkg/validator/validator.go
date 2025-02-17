package validator

import (
	"context"
	"fmt"

	"github.com/bufbuild/connect-go"
	"github.com/bufbuild/protovalidate-go"
	"github.com/tierklinik-dobersberg/apis/pkg/internal/timing"
	"google.golang.org/protobuf/proto"
)

func NewInterceptor(validator protovalidate.Validator) connect.UnaryInterceptorFunc {
	return func(uf connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, ar connect.AnyRequest) (connect.AnyResponse, error) {
			if err := timing.Track(ctx, "validator", func() error {
				return validator.Validate(ar.Any().(proto.Message))
			}); err != nil {
				return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid request: %w", err))
			}

			return uf(ctx, ar)
		}
	}
}
