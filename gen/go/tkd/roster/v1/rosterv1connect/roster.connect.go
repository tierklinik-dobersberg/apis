// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: tkd/roster/v1/roster.proto

package rosterv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/roster/v1"
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
	// RosterServiceName is the fully-qualified name of the RosterService service.
	RosterServiceName = "tkd.roster.v1.RosterService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// RosterServiceSessionProcedure is the fully-qualified name of the RosterService's Session RPC.
	RosterServiceSessionProcedure = "/tkd.roster.v1.RosterService/Session"
	// RosterServiceApproveRosterProcedure is the fully-qualified name of the RosterService's
	// ApproveRoster RPC.
	RosterServiceApproveRosterProcedure = "/tkd.roster.v1.RosterService/ApproveRoster"
	// RosterServiceDeleteRosterProcedure is the fully-qualified name of the RosterService's
	// DeleteRoster RPC.
	RosterServiceDeleteRosterProcedure = "/tkd.roster.v1.RosterService/DeleteRoster"
	// RosterServiceGetRosterProcedure is the fully-qualified name of the RosterService's GetRoster RPC.
	RosterServiceGetRosterProcedure = "/tkd.roster.v1.RosterService/GetRoster"
	// RosterServiceGetWorkingStaffProcedure is the fully-qualified name of the RosterService's
	// GetWorkingStaff RPC.
	RosterServiceGetWorkingStaffProcedure = "/tkd.roster.v1.RosterService/GetWorkingStaff"
)

// RosterServiceClient is a client for the tkd.roster.v1.RosterService service.
type RosterServiceClient interface {
	Session(context.Context) *connect_go.BidiStreamForClient[v1.SessionRequest, v1.SessionResponse]
	ApproveRoster(context.Context, *connect_go.Request[v1.ApproveRosterRequest]) (*connect_go.Response[v1.ApproveRosterResponse], error)
	DeleteRoster(context.Context, *connect_go.Request[v1.DeleteRosterRequest]) (*connect_go.Response[v1.DeleteRosterResponse], error)
	GetRoster(context.Context, *connect_go.Request[v1.GetRosterRequest]) (*connect_go.Response[v1.GetRosterResponse], error)
	GetWorkingStaff(context.Context, *connect_go.Request[v1.GetWorkingStaffRequest]) (*connect_go.Response[v1.GetWorkingStaffResponse], error)
}

// NewRosterServiceClient constructs a client for the tkd.roster.v1.RosterService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewRosterServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) RosterServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &rosterServiceClient{
		session: connect_go.NewClient[v1.SessionRequest, v1.SessionResponse](
			httpClient,
			baseURL+RosterServiceSessionProcedure,
			opts...,
		),
		approveRoster: connect_go.NewClient[v1.ApproveRosterRequest, v1.ApproveRosterResponse](
			httpClient,
			baseURL+RosterServiceApproveRosterProcedure,
			opts...,
		),
		deleteRoster: connect_go.NewClient[v1.DeleteRosterRequest, v1.DeleteRosterResponse](
			httpClient,
			baseURL+RosterServiceDeleteRosterProcedure,
			opts...,
		),
		getRoster: connect_go.NewClient[v1.GetRosterRequest, v1.GetRosterResponse](
			httpClient,
			baseURL+RosterServiceGetRosterProcedure,
			opts...,
		),
		getWorkingStaff: connect_go.NewClient[v1.GetWorkingStaffRequest, v1.GetWorkingStaffResponse](
			httpClient,
			baseURL+RosterServiceGetWorkingStaffProcedure,
			opts...,
		),
	}
}

// rosterServiceClient implements RosterServiceClient.
type rosterServiceClient struct {
	session         *connect_go.Client[v1.SessionRequest, v1.SessionResponse]
	approveRoster   *connect_go.Client[v1.ApproveRosterRequest, v1.ApproveRosterResponse]
	deleteRoster    *connect_go.Client[v1.DeleteRosterRequest, v1.DeleteRosterResponse]
	getRoster       *connect_go.Client[v1.GetRosterRequest, v1.GetRosterResponse]
	getWorkingStaff *connect_go.Client[v1.GetWorkingStaffRequest, v1.GetWorkingStaffResponse]
}

// Session calls tkd.roster.v1.RosterService.Session.
func (c *rosterServiceClient) Session(ctx context.Context) *connect_go.BidiStreamForClient[v1.SessionRequest, v1.SessionResponse] {
	return c.session.CallBidiStream(ctx)
}

// ApproveRoster calls tkd.roster.v1.RosterService.ApproveRoster.
func (c *rosterServiceClient) ApproveRoster(ctx context.Context, req *connect_go.Request[v1.ApproveRosterRequest]) (*connect_go.Response[v1.ApproveRosterResponse], error) {
	return c.approveRoster.CallUnary(ctx, req)
}

// DeleteRoster calls tkd.roster.v1.RosterService.DeleteRoster.
func (c *rosterServiceClient) DeleteRoster(ctx context.Context, req *connect_go.Request[v1.DeleteRosterRequest]) (*connect_go.Response[v1.DeleteRosterResponse], error) {
	return c.deleteRoster.CallUnary(ctx, req)
}

// GetRoster calls tkd.roster.v1.RosterService.GetRoster.
func (c *rosterServiceClient) GetRoster(ctx context.Context, req *connect_go.Request[v1.GetRosterRequest]) (*connect_go.Response[v1.GetRosterResponse], error) {
	return c.getRoster.CallUnary(ctx, req)
}

// GetWorkingStaff calls tkd.roster.v1.RosterService.GetWorkingStaff.
func (c *rosterServiceClient) GetWorkingStaff(ctx context.Context, req *connect_go.Request[v1.GetWorkingStaffRequest]) (*connect_go.Response[v1.GetWorkingStaffResponse], error) {
	return c.getWorkingStaff.CallUnary(ctx, req)
}

// RosterServiceHandler is an implementation of the tkd.roster.v1.RosterService service.
type RosterServiceHandler interface {
	Session(context.Context, *connect_go.BidiStream[v1.SessionRequest, v1.SessionResponse]) error
	ApproveRoster(context.Context, *connect_go.Request[v1.ApproveRosterRequest]) (*connect_go.Response[v1.ApproveRosterResponse], error)
	DeleteRoster(context.Context, *connect_go.Request[v1.DeleteRosterRequest]) (*connect_go.Response[v1.DeleteRosterResponse], error)
	GetRoster(context.Context, *connect_go.Request[v1.GetRosterRequest]) (*connect_go.Response[v1.GetRosterResponse], error)
	GetWorkingStaff(context.Context, *connect_go.Request[v1.GetWorkingStaffRequest]) (*connect_go.Response[v1.GetWorkingStaffResponse], error)
}

// NewRosterServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewRosterServiceHandler(svc RosterServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	rosterServiceSessionHandler := connect_go.NewBidiStreamHandler(
		RosterServiceSessionProcedure,
		svc.Session,
		opts...,
	)
	rosterServiceApproveRosterHandler := connect_go.NewUnaryHandler(
		RosterServiceApproveRosterProcedure,
		svc.ApproveRoster,
		opts...,
	)
	rosterServiceDeleteRosterHandler := connect_go.NewUnaryHandler(
		RosterServiceDeleteRosterProcedure,
		svc.DeleteRoster,
		opts...,
	)
	rosterServiceGetRosterHandler := connect_go.NewUnaryHandler(
		RosterServiceGetRosterProcedure,
		svc.GetRoster,
		opts...,
	)
	rosterServiceGetWorkingStaffHandler := connect_go.NewUnaryHandler(
		RosterServiceGetWorkingStaffProcedure,
		svc.GetWorkingStaff,
		opts...,
	)
	return "/tkd.roster.v1.RosterService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case RosterServiceSessionProcedure:
			rosterServiceSessionHandler.ServeHTTP(w, r)
		case RosterServiceApproveRosterProcedure:
			rosterServiceApproveRosterHandler.ServeHTTP(w, r)
		case RosterServiceDeleteRosterProcedure:
			rosterServiceDeleteRosterHandler.ServeHTTP(w, r)
		case RosterServiceGetRosterProcedure:
			rosterServiceGetRosterHandler.ServeHTTP(w, r)
		case RosterServiceGetWorkingStaffProcedure:
			rosterServiceGetWorkingStaffHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedRosterServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedRosterServiceHandler struct{}

func (UnimplementedRosterServiceHandler) Session(context.Context, *connect_go.BidiStream[v1.SessionRequest, v1.SessionResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.roster.v1.RosterService.Session is not implemented"))
}

func (UnimplementedRosterServiceHandler) ApproveRoster(context.Context, *connect_go.Request[v1.ApproveRosterRequest]) (*connect_go.Response[v1.ApproveRosterResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.roster.v1.RosterService.ApproveRoster is not implemented"))
}

func (UnimplementedRosterServiceHandler) DeleteRoster(context.Context, *connect_go.Request[v1.DeleteRosterRequest]) (*connect_go.Response[v1.DeleteRosterResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.roster.v1.RosterService.DeleteRoster is not implemented"))
}

func (UnimplementedRosterServiceHandler) GetRoster(context.Context, *connect_go.Request[v1.GetRosterRequest]) (*connect_go.Response[v1.GetRosterResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.roster.v1.RosterService.GetRoster is not implemented"))
}

func (UnimplementedRosterServiceHandler) GetWorkingStaff(context.Context, *connect_go.Request[v1.GetWorkingStaffRequest]) (*connect_go.Response[v1.GetWorkingStaffResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.roster.v1.RosterService.GetWorkingStaff is not implemented"))
}
