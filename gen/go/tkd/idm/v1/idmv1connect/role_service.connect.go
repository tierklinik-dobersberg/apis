// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: tkd/idm/v1/role_service.proto

package idmv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/idm/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// RoleServiceName is the fully-qualified name of the RoleService service.
	RoleServiceName = "tkd.idm.v1.RoleService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// RoleServiceCreateRoleProcedure is the fully-qualified name of the RoleService's CreateRole RPC.
	RoleServiceCreateRoleProcedure = "/tkd.idm.v1.RoleService/CreateRole"
	// RoleServiceUpdateRoleProcedure is the fully-qualified name of the RoleService's UpdateRole RPC.
	RoleServiceUpdateRoleProcedure = "/tkd.idm.v1.RoleService/UpdateRole"
	// RoleServiceDeleteRoleProcedure is the fully-qualified name of the RoleService's DeleteRole RPC.
	RoleServiceDeleteRoleProcedure = "/tkd.idm.v1.RoleService/DeleteRole"
	// RoleServiceListRolesProcedure is the fully-qualified name of the RoleService's ListRoles RPC.
	RoleServiceListRolesProcedure = "/tkd.idm.v1.RoleService/ListRoles"
	// RoleServiceGetRoleProcedure is the fully-qualified name of the RoleService's GetRole RPC.
	RoleServiceGetRoleProcedure = "/tkd.idm.v1.RoleService/GetRole"
	// RoleServiceAssignRoleToUserProcedure is the fully-qualified name of the RoleService's
	// AssignRoleToUser RPC.
	RoleServiceAssignRoleToUserProcedure = "/tkd.idm.v1.RoleService/AssignRoleToUser"
	// RoleServiceUnassignRoleToUserProcedure is the fully-qualified name of the RoleService's
	// UnassignRoleToUser RPC.
	RoleServiceUnassignRoleToUserProcedure = "/tkd.idm.v1.RoleService/UnassignRoleToUser"
)

// RoleServiceClient is a client for the tkd.idm.v1.RoleService service.
type RoleServiceClient interface {
	CreateRole(context.Context, *connect_go.Request[v1.CreateRoleRequest]) (*connect_go.Response[v1.CreateRoleResponse], error)
	UpdateRole(context.Context, *connect_go.Request[v1.UpdateRoleRequest]) (*connect_go.Response[v1.UpdateRoleResponse], error)
	DeleteRole(context.Context, *connect_go.Request[v1.DeleteRoleRequest]) (*connect_go.Response[v1.DeleteRoleResponse], error)
	ListRoles(context.Context, *connect_go.Request[v1.ListRolesRequest]) (*connect_go.Response[v1.ListRolesResponse], error)
	GetRole(context.Context, *connect_go.Request[v1.GetRoleRequest]) (*connect_go.Response[v1.GetRoleResponse], error)
	AssignRoleToUser(context.Context, *connect_go.Request[v1.AssignRoleToUserRequest]) (*connect_go.Response[v1.AssignRoleToUserResponse], error)
	UnassignRoleToUser(context.Context, *connect_go.Request[v1.UnassignRoleToUserRequest]) (*connect_go.Response[v1.UnassignRoleToUserResponse], error)
}

// NewRoleServiceClient constructs a client for the tkd.idm.v1.RoleService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewRoleServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) RoleServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &roleServiceClient{
		createRole: connect_go.NewClient[v1.CreateRoleRequest, v1.CreateRoleResponse](
			httpClient,
			baseURL+RoleServiceCreateRoleProcedure,
			opts...,
		),
		updateRole: connect_go.NewClient[v1.UpdateRoleRequest, v1.UpdateRoleResponse](
			httpClient,
			baseURL+RoleServiceUpdateRoleProcedure,
			opts...,
		),
		deleteRole: connect_go.NewClient[v1.DeleteRoleRequest, v1.DeleteRoleResponse](
			httpClient,
			baseURL+RoleServiceDeleteRoleProcedure,
			opts...,
		),
		listRoles: connect_go.NewClient[v1.ListRolesRequest, v1.ListRolesResponse](
			httpClient,
			baseURL+RoleServiceListRolesProcedure,
			opts...,
		),
		getRole: connect_go.NewClient[v1.GetRoleRequest, v1.GetRoleResponse](
			httpClient,
			baseURL+RoleServiceGetRoleProcedure,
			opts...,
		),
		assignRoleToUser: connect_go.NewClient[v1.AssignRoleToUserRequest, v1.AssignRoleToUserResponse](
			httpClient,
			baseURL+RoleServiceAssignRoleToUserProcedure,
			opts...,
		),
		unassignRoleToUser: connect_go.NewClient[v1.UnassignRoleToUserRequest, v1.UnassignRoleToUserResponse](
			httpClient,
			baseURL+RoleServiceUnassignRoleToUserProcedure,
			opts...,
		),
	}
}

// roleServiceClient implements RoleServiceClient.
type roleServiceClient struct {
	createRole         *connect_go.Client[v1.CreateRoleRequest, v1.CreateRoleResponse]
	updateRole         *connect_go.Client[v1.UpdateRoleRequest, v1.UpdateRoleResponse]
	deleteRole         *connect_go.Client[v1.DeleteRoleRequest, v1.DeleteRoleResponse]
	listRoles          *connect_go.Client[v1.ListRolesRequest, v1.ListRolesResponse]
	getRole            *connect_go.Client[v1.GetRoleRequest, v1.GetRoleResponse]
	assignRoleToUser   *connect_go.Client[v1.AssignRoleToUserRequest, v1.AssignRoleToUserResponse]
	unassignRoleToUser *connect_go.Client[v1.UnassignRoleToUserRequest, v1.UnassignRoleToUserResponse]
}

// CreateRole calls tkd.idm.v1.RoleService.CreateRole.
func (c *roleServiceClient) CreateRole(ctx context.Context, req *connect_go.Request[v1.CreateRoleRequest]) (*connect_go.Response[v1.CreateRoleResponse], error) {
	return c.createRole.CallUnary(ctx, req)
}

// UpdateRole calls tkd.idm.v1.RoleService.UpdateRole.
func (c *roleServiceClient) UpdateRole(ctx context.Context, req *connect_go.Request[v1.UpdateRoleRequest]) (*connect_go.Response[v1.UpdateRoleResponse], error) {
	return c.updateRole.CallUnary(ctx, req)
}

// DeleteRole calls tkd.idm.v1.RoleService.DeleteRole.
func (c *roleServiceClient) DeleteRole(ctx context.Context, req *connect_go.Request[v1.DeleteRoleRequest]) (*connect_go.Response[v1.DeleteRoleResponse], error) {
	return c.deleteRole.CallUnary(ctx, req)
}

// ListRoles calls tkd.idm.v1.RoleService.ListRoles.
func (c *roleServiceClient) ListRoles(ctx context.Context, req *connect_go.Request[v1.ListRolesRequest]) (*connect_go.Response[v1.ListRolesResponse], error) {
	return c.listRoles.CallUnary(ctx, req)
}

// GetRole calls tkd.idm.v1.RoleService.GetRole.
func (c *roleServiceClient) GetRole(ctx context.Context, req *connect_go.Request[v1.GetRoleRequest]) (*connect_go.Response[v1.GetRoleResponse], error) {
	return c.getRole.CallUnary(ctx, req)
}

// AssignRoleToUser calls tkd.idm.v1.RoleService.AssignRoleToUser.
func (c *roleServiceClient) AssignRoleToUser(ctx context.Context, req *connect_go.Request[v1.AssignRoleToUserRequest]) (*connect_go.Response[v1.AssignRoleToUserResponse], error) {
	return c.assignRoleToUser.CallUnary(ctx, req)
}

// UnassignRoleToUser calls tkd.idm.v1.RoleService.UnassignRoleToUser.
func (c *roleServiceClient) UnassignRoleToUser(ctx context.Context, req *connect_go.Request[v1.UnassignRoleToUserRequest]) (*connect_go.Response[v1.UnassignRoleToUserResponse], error) {
	return c.unassignRoleToUser.CallUnary(ctx, req)
}

// RoleServiceHandler is an implementation of the tkd.idm.v1.RoleService service.
type RoleServiceHandler interface {
	CreateRole(context.Context, *connect_go.Request[v1.CreateRoleRequest]) (*connect_go.Response[v1.CreateRoleResponse], error)
	UpdateRole(context.Context, *connect_go.Request[v1.UpdateRoleRequest]) (*connect_go.Response[v1.UpdateRoleResponse], error)
	DeleteRole(context.Context, *connect_go.Request[v1.DeleteRoleRequest]) (*connect_go.Response[v1.DeleteRoleResponse], error)
	ListRoles(context.Context, *connect_go.Request[v1.ListRolesRequest]) (*connect_go.Response[v1.ListRolesResponse], error)
	GetRole(context.Context, *connect_go.Request[v1.GetRoleRequest]) (*connect_go.Response[v1.GetRoleResponse], error)
	AssignRoleToUser(context.Context, *connect_go.Request[v1.AssignRoleToUserRequest]) (*connect_go.Response[v1.AssignRoleToUserResponse], error)
	UnassignRoleToUser(context.Context, *connect_go.Request[v1.UnassignRoleToUserRequest]) (*connect_go.Response[v1.UnassignRoleToUserResponse], error)
}

// NewRoleServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewRoleServiceHandler(svc RoleServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	roleServiceCreateRoleHandler := connect_go.NewUnaryHandler(
		RoleServiceCreateRoleProcedure,
		svc.CreateRole,
		opts...,
	)
	roleServiceUpdateRoleHandler := connect_go.NewUnaryHandler(
		RoleServiceUpdateRoleProcedure,
		svc.UpdateRole,
		opts...,
	)
	roleServiceDeleteRoleHandler := connect_go.NewUnaryHandler(
		RoleServiceDeleteRoleProcedure,
		svc.DeleteRole,
		opts...,
	)
	roleServiceListRolesHandler := connect_go.NewUnaryHandler(
		RoleServiceListRolesProcedure,
		svc.ListRoles,
		opts...,
	)
	roleServiceGetRoleHandler := connect_go.NewUnaryHandler(
		RoleServiceGetRoleProcedure,
		svc.GetRole,
		opts...,
	)
	roleServiceAssignRoleToUserHandler := connect_go.NewUnaryHandler(
		RoleServiceAssignRoleToUserProcedure,
		svc.AssignRoleToUser,
		opts...,
	)
	roleServiceUnassignRoleToUserHandler := connect_go.NewUnaryHandler(
		RoleServiceUnassignRoleToUserProcedure,
		svc.UnassignRoleToUser,
		opts...,
	)
	return "/tkd.idm.v1.RoleService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case RoleServiceCreateRoleProcedure:
			roleServiceCreateRoleHandler.ServeHTTP(w, r)
		case RoleServiceUpdateRoleProcedure:
			roleServiceUpdateRoleHandler.ServeHTTP(w, r)
		case RoleServiceDeleteRoleProcedure:
			roleServiceDeleteRoleHandler.ServeHTTP(w, r)
		case RoleServiceListRolesProcedure:
			roleServiceListRolesHandler.ServeHTTP(w, r)
		case RoleServiceGetRoleProcedure:
			roleServiceGetRoleHandler.ServeHTTP(w, r)
		case RoleServiceAssignRoleToUserProcedure:
			roleServiceAssignRoleToUserHandler.ServeHTTP(w, r)
		case RoleServiceUnassignRoleToUserProcedure:
			roleServiceUnassignRoleToUserHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedRoleServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedRoleServiceHandler struct{}

func (UnimplementedRoleServiceHandler) CreateRole(context.Context, *connect_go.Request[v1.CreateRoleRequest]) (*connect_go.Response[v1.CreateRoleResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.idm.v1.RoleService.CreateRole is not implemented"))
}

func (UnimplementedRoleServiceHandler) UpdateRole(context.Context, *connect_go.Request[v1.UpdateRoleRequest]) (*connect_go.Response[v1.UpdateRoleResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.idm.v1.RoleService.UpdateRole is not implemented"))
}

func (UnimplementedRoleServiceHandler) DeleteRole(context.Context, *connect_go.Request[v1.DeleteRoleRequest]) (*connect_go.Response[v1.DeleteRoleResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.idm.v1.RoleService.DeleteRole is not implemented"))
}

func (UnimplementedRoleServiceHandler) ListRoles(context.Context, *connect_go.Request[v1.ListRolesRequest]) (*connect_go.Response[v1.ListRolesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.idm.v1.RoleService.ListRoles is not implemented"))
}

func (UnimplementedRoleServiceHandler) GetRole(context.Context, *connect_go.Request[v1.GetRoleRequest]) (*connect_go.Response[v1.GetRoleResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.idm.v1.RoleService.GetRole is not implemented"))
}

func (UnimplementedRoleServiceHandler) AssignRoleToUser(context.Context, *connect_go.Request[v1.AssignRoleToUserRequest]) (*connect_go.Response[v1.AssignRoleToUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.idm.v1.RoleService.AssignRoleToUser is not implemented"))
}

func (UnimplementedRoleServiceHandler) UnassignRoleToUser(context.Context, *connect_go.Request[v1.UnassignRoleToUserRequest]) (*connect_go.Response[v1.UnassignRoleToUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.idm.v1.RoleService.UnassignRoleToUser is not implemented"))
}
