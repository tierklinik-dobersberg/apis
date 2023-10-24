package auth

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/bufbuild/connect-go"
	"github.com/tierklinik-dobersberg/apis/pkg/log"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"

	commonv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/common/v1"
)

type (
	// RemoteUser defines the user that is currently accessing a RPC method.
	RemoteUser struct {
		ID                  string
		DisplayName         string
		Roles               []string
		PrimaryMail         string
		PrimaryMailVerified *bool
	}

	// ExtractorFunc extracts the user ID and a list of assigned user roles from req.
	ExtractorFunc func(ctx context.Context, req connect.AnyRequest) (RemoteUser, error)
)

var remoteUserContextKey = struct{ S string }{S: "remoteUserContextKey"}

// From returns the remote user associated with ctx. Note that from may only
// be used if NewAuthAnnotationInterceptor has been added to the RPC middlewares.
func From(ctx context.Context) *RemoteUser {
	v, _ := ctx.Value(remoteUserContextKey).(*RemoteUser)
	return v
}

// RemoteHeaderExtractor extracts remote user information from X-Remote-* headers
// in the request. This method is aligned with the remote headers added by the forwared authentication
// endpoint of github.com/tierklinik-dobersberg/cis-idm.
func RemoteHeaderExtractor(ctx context.Context, req connect.AnyRequest) (RemoteUser, error) {
	headers := req.Header()

	remoteUser := RemoteUser{
		ID:          headers.Get("X-Remote-User-ID"),
		DisplayName: headers.Get("X-Remote-User"),
		Roles:       headers.Values("X-Remote-Role"),
		PrimaryMail: headers.Get("X-Remote-Mail"),
	}

	if val := headers.Get("X-Remote-Mail-Verified"); val != "" {
		b, err := strconv.ParseBool(val)
		if err == nil {
			remoteUser.PrimaryMailVerified = &b
		}
	}

	return remoteUser, nil
}

func NewAuthAnnotationInterceptor(registry *protoregistry.Files, extractor ExtractorFunc) connect.UnaryInterceptorFunc {
	if extractor == nil {
		extractor = RemoteHeaderExtractor
	}

	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			parts := strings.Split(req.Spec().Procedure, "/")

			methodDesc := getMethodDesc(registry, parts[1], parts[2])
			if methodDesc == nil {
				return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to find method descriptor for %s", req.Spec().Procedure))
			}

			l := log.L(ctx).WithField("method", methodDesc.FullName())

			// get the remote user
			usr, err := extractor(ctx, req)
			if err != nil {
				return nil, err
			}

			// if we have a user ID add some more information to the logger and
			// append the user information to the request context
			if usr.ID != "" {
				l = l.WithField("user.id", usr.ID).WithField("user.displayName", usr.DisplayName)
				ctx = context.WithValue(ctx, remoteUserContextKey, &usr)
			}

			opts, ok := proto.GetExtension(methodDesc.Options(), commonv1.E_Auth).(*commonv1.AuthDecorator)

			if ok && opts != nil {
				switch opts.Require {
				case commonv1.AuthRequirement_AUTH_REQ_REQUIRED:
					if usr.ID == "" {
						return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("not access token provided"))
					}

					// make sure the user has at least one of the required roles assigned
					if len(opts.AllowedRoles) > 0 {
						isAllowed := false

						for _, allowedRole := range opts.AllowedRoles {
							if slices.Contains(usr.Roles, allowedRole) {
								isAllowed = true
								break
							}
						}

						if !isAllowed {
							return nil, connect.NewError(connect.CodePermissionDenied, fmt.Errorf("access token does not include one of the required roles"))
						}
					}

				case commonv1.AuthRequirement_AUTH_REQ_UNSPECIFIED:
					// nothing to do
				default:
					l.WithField("requirement", opts.String()).Infof("unhandeled authentication requirement")
				}
			} else {
				l.Infof("not authentication requirement specified for service method")
			}

			ctx = log.WithLogger(ctx, l)

			return next(ctx, req)
		})
	}

	return interceptor
}

func getMethodDesc(reg *protoregistry.Files, fqServiceName string, methodName string) protoreflect.MethodDescriptor {
	serviceNameParts := strings.Split(fqServiceName, ".")
	serviceName := serviceNameParts[len(serviceNameParts)-1]

	var methodDesc protoreflect.MethodDescriptor

	reg.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		if strings.HasPrefix(fqServiceName, string(fd.FullName())) {
			serviceDesc := fd.Services().ByName(protoreflect.Name(serviceName))
			if serviceDesc != nil {

				methodDesc = serviceDesc.Methods().ByName(protoreflect.Name(methodName))
				if methodDesc != nil {
					return false
				}
			}
		}

		return true
	})

	return methodDesc
}
