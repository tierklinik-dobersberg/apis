// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: tkd/idm/v1/user_service.proto

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
const _ = connect_go.IsAtLeastVersion1_7_0

const (
	// UserServiceName is the fully-qualified name of the UserService service.
	UserServiceName = "tkd.idm.v1.UserService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// UserServiceGetUserProcedure is the fully-qualified name of the UserService's GetUser RPC.
	UserServiceGetUserProcedure = "/tkd.idm.v1.UserService/GetUser"
	// UserServiceInviteUserProcedure is the fully-qualified name of the UserService's InviteUser RPC.
	UserServiceInviteUserProcedure = "/tkd.idm.v1.UserService/InviteUser"
	// UserServiceListUsersProcedure is the fully-qualified name of the UserService's ListUsers RPC.
	UserServiceListUsersProcedure = "/tkd.idm.v1.UserService/ListUsers"
	// UserServiceCreateUserProcedure is the fully-qualified name of the UserService's CreateUser RPC.
	UserServiceCreateUserProcedure = "/tkd.idm.v1.UserService/CreateUser"
	// UserServiceUpdateUserProcedure is the fully-qualified name of the UserService's UpdateUser RPC.
	UserServiceUpdateUserProcedure = "/tkd.idm.v1.UserService/UpdateUser"
	// UserServiceDeleteUserProcedure is the fully-qualified name of the UserService's DeleteUser RPC.
	UserServiceDeleteUserProcedure = "/tkd.idm.v1.UserService/DeleteUser"
)

// UserServiceClient is a client for the tkd.idm.v1.UserService service.
type UserServiceClient interface {
	GetUser(context.Context, *connect_go.Request[v1.GetUserRequest]) (*connect_go.Response[v1.GetUserResponse], error)
	InviteUser(context.Context, *connect_go.Request[v1.InviteUserRequest]) (*connect_go.Response[v1.InviteUserResponse], error)
	ListUsers(context.Context, *connect_go.Request[v1.ListUsersRequest]) (*connect_go.Response[v1.ListUsersResponse], error)
	CreateUser(context.Context, *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CreateUserResponse], error)
	UpdateUser(context.Context, *connect_go.Request[v1.UpdateUserRequest]) (*connect_go.Response[v1.UpdateUserResponse], error)
	DeleteUser(context.Context, *connect_go.Request[v1.DeleteUserRequest]) (*connect_go.Response[v1.DeleteUserResponse], error)
}

// NewUserServiceClient constructs a client for the tkd.idm.v1.UserService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) UserServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userServiceClient{
		getUser: connect_go.NewClient[v1.GetUserRequest, v1.GetUserResponse](
			httpClient,
			baseURL+UserServiceGetUserProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		inviteUser: connect_go.NewClient[v1.InviteUserRequest, v1.InviteUserResponse](
			httpClient,
			baseURL+UserServiceInviteUserProcedure,
			opts...,
		),
		listUsers: connect_go.NewClient[v1.ListUsersRequest, v1.ListUsersResponse](
			httpClient,
			baseURL+UserServiceListUsersProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		createUser: connect_go.NewClient[v1.CreateUserRequest, v1.CreateUserResponse](
			httpClient,
			baseURL+UserServiceCreateUserProcedure,
			opts...,
		),
		updateUser: connect_go.NewClient[v1.UpdateUserRequest, v1.UpdateUserResponse](
			httpClient,
			baseURL+UserServiceUpdateUserProcedure,
			opts...,
		),
		deleteUser: connect_go.NewClient[v1.DeleteUserRequest, v1.DeleteUserResponse](
			httpClient,
			baseURL+UserServiceDeleteUserProcedure,
			opts...,
		),
	}
}

// userServiceClient implements UserServiceClient.
type userServiceClient struct {
	getUser    *connect_go.Client[v1.GetUserRequest, v1.GetUserResponse]
	inviteUser *connect_go.Client[v1.InviteUserRequest, v1.InviteUserResponse]
	listUsers  *connect_go.Client[v1.ListUsersRequest, v1.ListUsersResponse]
	createUser *connect_go.Client[v1.CreateUserRequest, v1.CreateUserResponse]
	updateUser *connect_go.Client[v1.UpdateUserRequest, v1.UpdateUserResponse]
	deleteUser *connect_go.Client[v1.DeleteUserRequest, v1.DeleteUserResponse]
}

// GetUser calls tkd.idm.v1.UserService.GetUser.
func (c *userServiceClient) GetUser(ctx context.Context, req *connect_go.Request[v1.GetUserRequest]) (*connect_go.Response[v1.GetUserResponse], error) {
	return c.getUser.CallUnary(ctx, req)
}

// InviteUser calls tkd.idm.v1.UserService.InviteUser.
func (c *userServiceClient) InviteUser(ctx context.Context, req *connect_go.Request[v1.InviteUserRequest]) (*connect_go.Response[v1.InviteUserResponse], error) {
	return c.inviteUser.CallUnary(ctx, req)
}

// ListUsers calls tkd.idm.v1.UserService.ListUsers.
func (c *userServiceClient) ListUsers(ctx context.Context, req *connect_go.Request[v1.ListUsersRequest]) (*connect_go.Response[v1.ListUsersResponse], error) {
	return c.listUsers.CallUnary(ctx, req)
}

// CreateUser calls tkd.idm.v1.UserService.CreateUser.
func (c *userServiceClient) CreateUser(ctx context.Context, req *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CreateUserResponse], error) {
	return c.createUser.CallUnary(ctx, req)
}

// UpdateUser calls tkd.idm.v1.UserService.UpdateUser.
func (c *userServiceClient) UpdateUser(ctx context.Context, req *connect_go.Request[v1.UpdateUserRequest]) (*connect_go.Response[v1.UpdateUserResponse], error) {
	return c.updateUser.CallUnary(ctx, req)
}

// DeleteUser calls tkd.idm.v1.UserService.DeleteUser.
func (c *userServiceClient) DeleteUser(ctx context.Context, req *connect_go.Request[v1.DeleteUserRequest]) (*connect_go.Response[v1.DeleteUserResponse], error) {
	return c.deleteUser.CallUnary(ctx, req)
}

// UserServiceHandler is an implementation of the tkd.idm.v1.UserService service.
type UserServiceHandler interface {
	GetUser(context.Context, *connect_go.Request[v1.GetUserRequest]) (*connect_go.Response[v1.GetUserResponse], error)
	InviteUser(context.Context, *connect_go.Request[v1.InviteUserRequest]) (*connect_go.Response[v1.InviteUserResponse], error)
	ListUsers(context.Context, *connect_go.Request[v1.ListUsersRequest]) (*connect_go.Response[v1.ListUsersResponse], error)
	CreateUser(context.Context, *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CreateUserResponse], error)
	UpdateUser(context.Context, *connect_go.Request[v1.UpdateUserRequest]) (*connect_go.Response[v1.UpdateUserResponse], error)
	DeleteUser(context.Context, *connect_go.Request[v1.DeleteUserRequest]) (*connect_go.Response[v1.DeleteUserResponse], error)
}

// NewUserServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserServiceHandler(svc UserServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	userServiceGetUserHandler := connect_go.NewUnaryHandler(
		UserServiceGetUserProcedure,
		svc.GetUser,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	)
	userServiceInviteUserHandler := connect_go.NewUnaryHandler(
		UserServiceInviteUserProcedure,
		svc.InviteUser,
		opts...,
	)
	userServiceListUsersHandler := connect_go.NewUnaryHandler(
		UserServiceListUsersProcedure,
		svc.ListUsers,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	)
	userServiceCreateUserHandler := connect_go.NewUnaryHandler(
		UserServiceCreateUserProcedure,
		svc.CreateUser,
		opts...,
	)
	userServiceUpdateUserHandler := connect_go.NewUnaryHandler(
		UserServiceUpdateUserProcedure,
		svc.UpdateUser,
		opts...,
	)
	userServiceDeleteUserHandler := connect_go.NewUnaryHandler(
		UserServiceDeleteUserProcedure,
		svc.DeleteUser,
		opts...,
	)
	return "/tkd.idm.v1.UserService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case UserServiceGetUserProcedure:
			userServiceGetUserHandler.ServeHTTP(w, r)
		case UserServiceInviteUserProcedure:
			userServiceInviteUserHandler.ServeHTTP(w, r)
		case UserServiceListUsersProcedure:
			userServiceListUsersHandler.ServeHTTP(w, r)
		case UserServiceCreateUserProcedure:
			userServiceCreateUserHandler.ServeHTTP(w, r)
		case UserServiceUpdateUserProcedure:
			userServiceUpdateUserHandler.ServeHTTP(w, r)
		case UserServiceDeleteUserProcedure:
			userServiceDeleteUserHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedUserServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserServiceHandler struct{}

func (UnimplementedUserServiceHandler) GetUser(context.Context, *connect_go.Request[v1.GetUserRequest]) (*connect_go.Response[v1.GetUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.idm.v1.UserService.GetUser is not implemented"))
}

func (UnimplementedUserServiceHandler) InviteUser(context.Context, *connect_go.Request[v1.InviteUserRequest]) (*connect_go.Response[v1.InviteUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.idm.v1.UserService.InviteUser is not implemented"))
}

func (UnimplementedUserServiceHandler) ListUsers(context.Context, *connect_go.Request[v1.ListUsersRequest]) (*connect_go.Response[v1.ListUsersResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.idm.v1.UserService.ListUsers is not implemented"))
}

func (UnimplementedUserServiceHandler) CreateUser(context.Context, *connect_go.Request[v1.CreateUserRequest]) (*connect_go.Response[v1.CreateUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.idm.v1.UserService.CreateUser is not implemented"))
}

func (UnimplementedUserServiceHandler) UpdateUser(context.Context, *connect_go.Request[v1.UpdateUserRequest]) (*connect_go.Response[v1.UpdateUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.idm.v1.UserService.UpdateUser is not implemented"))
}

func (UnimplementedUserServiceHandler) DeleteUser(context.Context, *connect_go.Request[v1.DeleteUserRequest]) (*connect_go.Response[v1.DeleteUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.idm.v1.UserService.DeleteUser is not implemented"))
}
