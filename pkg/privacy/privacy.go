package privacy

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/tierklinik-dobersberg/apis/pkg/internal/timing"
	"google.golang.org/protobuf/proto"
)

type SubjectResolver interface {
	GetRequestSubject(ctx context.Context, ar connect.AnyRequest) (userID string, roleIDs []string, err error)
}

type SubjectResolverFunc func(ctx context.Context, ar connect.AnyRequest) (string, []string, error)

func (fn SubjectResolverFunc) GetRequestSubject(ctx context.Context, ar connect.AnyRequest) (string, []string, error) {
	return fn(ctx, ar)
}

func NewFilterInterceptor(resolver SubjectResolver) connect.UnaryInterceptorFunc {
	return func(uf connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, ar connect.AnyRequest) (connect.AnyResponse, error) {
			req, err := uf(ctx, ar)
			if err != nil {
				return nil, err
			}

			if err := timing.Track(ctx, "privacy", func() error {
				userID, roles, err := resolver.GetRequestSubject(ctx, ar)
				if err != nil {
					return err
				}

				if err := FilterAllowedFields(req.Any().(proto.Message), userID, roles); err != nil {
					return connect.NewError(connect.CodeInternal, err)
				}

				return nil
			}); err != nil {
				return nil, err
			}

			return req, nil
		}
	}
}
