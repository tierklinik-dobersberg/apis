// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: tkd/roster/v1/constraint.proto

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
	// ConstraintServiceName is the fully-qualified name of the ConstraintService service.
	ConstraintServiceName = "tkd.roster.v1.ConstraintService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ConstraintServiceCreateConstraintProcedure is the fully-qualified name of the ConstraintService's
	// CreateConstraint RPC.
	ConstraintServiceCreateConstraintProcedure = "/tkd.roster.v1.ConstraintService/CreateConstraint"
	// ConstraintServiceUpdateConstraintProcedure is the fully-qualified name of the ConstraintService's
	// UpdateConstraint RPC.
	ConstraintServiceUpdateConstraintProcedure = "/tkd.roster.v1.ConstraintService/UpdateConstraint"
	// ConstraintServiceDeleteConstraintProcedure is the fully-qualified name of the ConstraintService's
	// DeleteConstraint RPC.
	ConstraintServiceDeleteConstraintProcedure = "/tkd.roster.v1.ConstraintService/DeleteConstraint"
	// ConstraintServiceFindConstraintsProcedure is the fully-qualified name of the ConstraintService's
	// FindConstraints RPC.
	ConstraintServiceFindConstraintsProcedure = "/tkd.roster.v1.ConstraintService/FindConstraints"
)

// ConstraintServiceClient is a client for the tkd.roster.v1.ConstraintService service.
type ConstraintServiceClient interface {
	CreateConstraint(context.Context, *connect_go.Request[v1.CreateConstraintRequest]) (*connect_go.Response[v1.CreateConstraintResponse], error)
	UpdateConstraint(context.Context, *connect_go.Request[v1.UpdateConstraintRequest]) (*connect_go.Response[v1.UpdateConstraintResponse], error)
	DeleteConstraint(context.Context, *connect_go.Request[v1.DeleteConstraintRequest]) (*connect_go.Response[v1.DeleteConstraintResponse], error)
	FindConstraints(context.Context, *connect_go.Request[v1.FindConstraintsRequest]) (*connect_go.Response[v1.FindConstraintsResponse], error)
}

// NewConstraintServiceClient constructs a client for the tkd.roster.v1.ConstraintService service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewConstraintServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ConstraintServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &constraintServiceClient{
		createConstraint: connect_go.NewClient[v1.CreateConstraintRequest, v1.CreateConstraintResponse](
			httpClient,
			baseURL+ConstraintServiceCreateConstraintProcedure,
			opts...,
		),
		updateConstraint: connect_go.NewClient[v1.UpdateConstraintRequest, v1.UpdateConstraintResponse](
			httpClient,
			baseURL+ConstraintServiceUpdateConstraintProcedure,
			opts...,
		),
		deleteConstraint: connect_go.NewClient[v1.DeleteConstraintRequest, v1.DeleteConstraintResponse](
			httpClient,
			baseURL+ConstraintServiceDeleteConstraintProcedure,
			opts...,
		),
		findConstraints: connect_go.NewClient[v1.FindConstraintsRequest, v1.FindConstraintsResponse](
			httpClient,
			baseURL+ConstraintServiceFindConstraintsProcedure,
			opts...,
		),
	}
}

// constraintServiceClient implements ConstraintServiceClient.
type constraintServiceClient struct {
	createConstraint *connect_go.Client[v1.CreateConstraintRequest, v1.CreateConstraintResponse]
	updateConstraint *connect_go.Client[v1.UpdateConstraintRequest, v1.UpdateConstraintResponse]
	deleteConstraint *connect_go.Client[v1.DeleteConstraintRequest, v1.DeleteConstraintResponse]
	findConstraints  *connect_go.Client[v1.FindConstraintsRequest, v1.FindConstraintsResponse]
}

// CreateConstraint calls tkd.roster.v1.ConstraintService.CreateConstraint.
func (c *constraintServiceClient) CreateConstraint(ctx context.Context, req *connect_go.Request[v1.CreateConstraintRequest]) (*connect_go.Response[v1.CreateConstraintResponse], error) {
	return c.createConstraint.CallUnary(ctx, req)
}

// UpdateConstraint calls tkd.roster.v1.ConstraintService.UpdateConstraint.
func (c *constraintServiceClient) UpdateConstraint(ctx context.Context, req *connect_go.Request[v1.UpdateConstraintRequest]) (*connect_go.Response[v1.UpdateConstraintResponse], error) {
	return c.updateConstraint.CallUnary(ctx, req)
}

// DeleteConstraint calls tkd.roster.v1.ConstraintService.DeleteConstraint.
func (c *constraintServiceClient) DeleteConstraint(ctx context.Context, req *connect_go.Request[v1.DeleteConstraintRequest]) (*connect_go.Response[v1.DeleteConstraintResponse], error) {
	return c.deleteConstraint.CallUnary(ctx, req)
}

// FindConstraints calls tkd.roster.v1.ConstraintService.FindConstraints.
func (c *constraintServiceClient) FindConstraints(ctx context.Context, req *connect_go.Request[v1.FindConstraintsRequest]) (*connect_go.Response[v1.FindConstraintsResponse], error) {
	return c.findConstraints.CallUnary(ctx, req)
}

// ConstraintServiceHandler is an implementation of the tkd.roster.v1.ConstraintService service.
type ConstraintServiceHandler interface {
	CreateConstraint(context.Context, *connect_go.Request[v1.CreateConstraintRequest]) (*connect_go.Response[v1.CreateConstraintResponse], error)
	UpdateConstraint(context.Context, *connect_go.Request[v1.UpdateConstraintRequest]) (*connect_go.Response[v1.UpdateConstraintResponse], error)
	DeleteConstraint(context.Context, *connect_go.Request[v1.DeleteConstraintRequest]) (*connect_go.Response[v1.DeleteConstraintResponse], error)
	FindConstraints(context.Context, *connect_go.Request[v1.FindConstraintsRequest]) (*connect_go.Response[v1.FindConstraintsResponse], error)
}

// NewConstraintServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewConstraintServiceHandler(svc ConstraintServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	constraintServiceCreateConstraintHandler := connect_go.NewUnaryHandler(
		ConstraintServiceCreateConstraintProcedure,
		svc.CreateConstraint,
		opts...,
	)
	constraintServiceUpdateConstraintHandler := connect_go.NewUnaryHandler(
		ConstraintServiceUpdateConstraintProcedure,
		svc.UpdateConstraint,
		opts...,
	)
	constraintServiceDeleteConstraintHandler := connect_go.NewUnaryHandler(
		ConstraintServiceDeleteConstraintProcedure,
		svc.DeleteConstraint,
		opts...,
	)
	constraintServiceFindConstraintsHandler := connect_go.NewUnaryHandler(
		ConstraintServiceFindConstraintsProcedure,
		svc.FindConstraints,
		opts...,
	)
	return "/tkd.roster.v1.ConstraintService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ConstraintServiceCreateConstraintProcedure:
			constraintServiceCreateConstraintHandler.ServeHTTP(w, r)
		case ConstraintServiceUpdateConstraintProcedure:
			constraintServiceUpdateConstraintHandler.ServeHTTP(w, r)
		case ConstraintServiceDeleteConstraintProcedure:
			constraintServiceDeleteConstraintHandler.ServeHTTP(w, r)
		case ConstraintServiceFindConstraintsProcedure:
			constraintServiceFindConstraintsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedConstraintServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedConstraintServiceHandler struct{}

func (UnimplementedConstraintServiceHandler) CreateConstraint(context.Context, *connect_go.Request[v1.CreateConstraintRequest]) (*connect_go.Response[v1.CreateConstraintResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.roster.v1.ConstraintService.CreateConstraint is not implemented"))
}

func (UnimplementedConstraintServiceHandler) UpdateConstraint(context.Context, *connect_go.Request[v1.UpdateConstraintRequest]) (*connect_go.Response[v1.UpdateConstraintResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.roster.v1.ConstraintService.UpdateConstraint is not implemented"))
}

func (UnimplementedConstraintServiceHandler) DeleteConstraint(context.Context, *connect_go.Request[v1.DeleteConstraintRequest]) (*connect_go.Response[v1.DeleteConstraintResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.roster.v1.ConstraintService.DeleteConstraint is not implemented"))
}

func (UnimplementedConstraintServiceHandler) FindConstraints(context.Context, *connect_go.Request[v1.FindConstraintsRequest]) (*connect_go.Response[v1.FindConstraintsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.roster.v1.ConstraintService.FindConstraints is not implemented"))
}