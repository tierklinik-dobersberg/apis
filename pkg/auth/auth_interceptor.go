package auth

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/bufbuild/connect-go"
	"github.com/tierklinik-dobersberg/apis/pkg/data"
	"github.com/tierklinik-dobersberg/apis/pkg/internal/timing"
	"github.com/tierklinik-dobersberg/apis/pkg/log"
	"golang.org/x/sync/singleflight"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"

	commonv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/common/v1"
	idmv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/idm/v1"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/idm/v1/idmv1connect"
)

type (
	// RemoteUser defines the user that is currently accessing a RPC method.
	RemoteUser struct {
		ID                  string
		Username            string
		DisplayName         string
		RoleIDs             []string
		PrimaryMail         string
		PrimaryMailVerified *bool
		Admin               bool
		ResolvedRoles       []*idmv1.Role
	}

	// ExtractorFunc extracts the user ID and a list of assigned user roles from req.
	ExtractorFunc func(ctx context.Context, req connect.AnyRequest) (RemoteUser, error)

	// RoleResolverFunc should return the role definition for the specified ID.
	RoleResolverFunc func(ctx context.Context, roleId string) (*idmv1.Role, error)
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
		Username:    headers.Get("X-Remote-User"),
		DisplayName: headers.Get("X-Remote-User-Display-Name"),
		RoleIDs:     headers.Values("X-Remote-Role"),
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

func NewIDMRoleResolver(cli idmv1connect.RoleServiceClient) RoleResolverFunc {
	return func(ctx context.Context, roleId string) (*idmv1.Role, error) {
		res, err := cli.GetRole(ctx, connect.NewRequest(&idmv1.GetRoleRequest{
			Search: &idmv1.GetRoleRequest_Id{
				Id: roleId,
			},
		}))

		if err != nil {
			return nil, err
		}

		return res.Msg.Role, nil
	}
}

func NewAuthAnnotationInterceptor(registry *protoregistry.Files, roleResolver RoleResolverFunc, extractor ExtractorFunc) connect.UnaryInterceptorFunc {
	if extractor == nil {
		extractor = RemoteHeaderExtractor
	}

	// role caching
	var (
		roleLock      sync.RWMutex
		roles         = make(map[string]*idmv1.Role)
		inFlightGroup = new(singleflight.Group)
	)

	// a small utility method for fetching and caching role definitions.
	getRole := func(ctx context.Context, roleId string) (*idmv1.Role, error) {
		l := log.L(ctx)

		roleLock.RLock()
		if r, ok := roles[roleId]; ok {
			roleLock.RUnlock()

			l.Debugf("resolved role id %q from cache: name=%q", roleId, r.Name)

			return r, nil
		}
		roleLock.RUnlock()

		res, err, _ := inFlightGroup.Do(roleId, func() (any, error) {
			var role *idmv1.Role
			err := timing.Track(ctx, "resolve-role-id", func() error {
				var err error
				role, err = roleResolver(context.Background(), roleId)

				return err
			})

			if err != nil {
				l.Errorf("failed to resolve role id %q: %s", roleId, err)

				return nil, fmt.Errorf("failed to resolve role: %w", err)
			}

			l.Debugf("resolved role id %q: name=%q", roleId, role.Name)

			roleLock.Lock()
			defer roleLock.Unlock()

			roles[roleId] = role

			return role, nil
		})

		if err != nil {
			return nil, err
		}

		return res.(*idmv1.Role), nil
	}

	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			parts := strings.Split(req.Spec().Procedure, "/")

			serviceDesc, methodDesc := getMethodDesc(registry, parts[1], parts[2])
			if serviceDesc == nil || methodDesc == nil {
				return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to find method descriptor for %s", req.Spec().Procedure))
			}

			l := log.L(ctx)

			// get the remote user
			usr, err := extractor(ctx, req)
			if err != nil {
				return nil, err
			}

			// make sure we resolve all user roles
			for _, roleId := range usr.RoleIDs {
				role, err := getRole(ctx, roleId)
				if err != nil {
					return nil, err
				}

				usr.ResolvedRoles = append(usr.ResolvedRoles, role)
			}

			// get the service options and check if the user is administrator.
			serviceOptions, _ := proto.GetExtension(serviceDesc.Options(), commonv1.E_ServiceAuth).(*commonv1.ServiceAuthDecorator)
			if serviceOptions != nil {
				// first check if the user has the specified role ID assigned.
				usr.Admin = data.ElemInBothSlices(serviceOptions.AdminRoles, usr.RoleIDs)

				if !usr.Admin {
					// the tkd.common.v1.service_auth option may also be used to specify role names instead of IDs
					// so we need to check for the names as well.
					usr.Admin = data.ElemInBothSlicesFunc(serviceOptions.AdminRoles, usr.ResolvedRoles, func(role *idmv1.Role) string {
						return role.Name
					})
				}
			}

			methodOptions, _ := proto.GetExtension(methodDesc.Options(), commonv1.E_Auth).(*commonv1.AuthDecorator)

			// if we have a user ID add some more information to the logger and
			// append the user information to the request context
			if usr.ID != "" {
				l = l.WithField("user.id", usr.ID).WithField("user.displayName", usr.DisplayName)
				ctx = context.WithValue(ctx, remoteUserContextKey, &usr)
			}

			// check method authentication requirements.
			if methodOptions != nil {
				switch methodOptions.Require {
				case commonv1.AuthRequirement_AUTH_REQ_ADMIN:
					if usr.ID == "" {
						return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("not access token provided"))
					}

					if !usr.Admin {
						return nil, connect.NewError(connect.CodePermissionDenied, fmt.Errorf("you're not allowed to perfom this operation"))
					}

				case commonv1.AuthRequirement_AUTH_REQ_REQUIRED:
					if usr.ID == "" {
						return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("not access token provided"))
					}

					// make sure the user has at least one of the required roles assigned
					if len(methodOptions.AllowedRoles) > 0 {
						isAllowed := data.ElemInBothSlices(methodOptions.AllowedRoles, usr.RoleIDs)

						if !isAllowed {
							isAllowed = data.ElemInBothSlicesFunc(methodOptions.AllowedRoles, usr.ResolvedRoles, func(role *idmv1.Role) string {
								return role.Name
							})
						}

						if !isAllowed {
							return nil, connect.NewError(connect.CodePermissionDenied, fmt.Errorf("access token does not include one of the required roles"))
						}
					}

				case commonv1.AuthRequirement_AUTH_REQ_UNSPECIFIED:
					// nothing to do
				default:
					l.WithField("requirement", methodOptions.String()).Infof("unhandeled authentication requirement")
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

func getMethodDesc(reg *protoregistry.Files, fqServiceName string, methodName string) (protoreflect.ServiceDescriptor, protoreflect.MethodDescriptor) {
	serviceNameParts := strings.Split(fqServiceName, ".")
	serviceName := serviceNameParts[len(serviceNameParts)-1]

	var (
		methodDesc  protoreflect.MethodDescriptor
		serviceDesc protoreflect.ServiceDescriptor
	)

	reg.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		if strings.HasPrefix(fqServiceName, string(fd.FullName())) {
			serviceDesc = fd.Services().ByName(protoreflect.Name(serviceName))
			if serviceDesc != nil {

				methodDesc = serviceDesc.Methods().ByName(protoreflect.Name(methodName))
				if methodDesc != nil {
					return false
				}
			}
		}

		return true
	})

	return serviceDesc, methodDesc
}
