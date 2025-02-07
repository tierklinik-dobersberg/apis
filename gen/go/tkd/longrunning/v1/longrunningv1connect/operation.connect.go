// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: tkd/longrunning/v1/operation.proto

package longrunningv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/longrunning/v1"
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
	// LongRunningServiceName is the fully-qualified name of the LongRunningService service.
	LongRunningServiceName = "tkd.longrunning.v1.LongRunningService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// LongRunningServiceRegisterOperationProcedure is the fully-qualified name of the
	// LongRunningService's RegisterOperation RPC.
	LongRunningServiceRegisterOperationProcedure = "/tkd.longrunning.v1.LongRunningService/RegisterOperation"
	// LongRunningServiceUpdateOperationProcedure is the fully-qualified name of the
	// LongRunningService's UpdateOperation RPC.
	LongRunningServiceUpdateOperationProcedure = "/tkd.longrunning.v1.LongRunningService/UpdateOperation"
	// LongRunningServiceCompleteOperationProcedure is the fully-qualified name of the
	// LongRunningService's CompleteOperation RPC.
	LongRunningServiceCompleteOperationProcedure = "/tkd.longrunning.v1.LongRunningService/CompleteOperation"
	// LongRunningServiceQueryOperationsProcedure is the fully-qualified name of the
	// LongRunningService's QueryOperations RPC.
	LongRunningServiceQueryOperationsProcedure = "/tkd.longrunning.v1.LongRunningService/QueryOperations"
	// LongRunningServiceGetOperationProcedure is the fully-qualified name of the LongRunningService's
	// GetOperation RPC.
	LongRunningServiceGetOperationProcedure = "/tkd.longrunning.v1.LongRunningService/GetOperation"
)

// LongRunningServiceClient is a client for the tkd.longrunning.v1.LongRunningService service.
type LongRunningServiceClient interface {
	RegisterOperation(context.Context, *connect_go.Request[v1.RegisterOperationRequest]) (*connect_go.Response[v1.Operation], error)
	UpdateOperation(context.Context, *connect_go.Request[v1.UpdateOperationRequest]) (*connect_go.Response[v1.Operation], error)
	CompleteOperation(context.Context, *connect_go.Request[v1.CompleteOperationRequest]) (*connect_go.Response[v1.Operation], error)
	QueryOperations(context.Context, *connect_go.Request[v1.QueryOperationsRequest]) (*connect_go.Response[v1.QueryOperationsResponse], error)
	GetOperation(context.Context, *connect_go.Request[v1.GetOperationRequest]) (*connect_go.Response[v1.Operation], error)
}

// NewLongRunningServiceClient constructs a client for the tkd.longrunning.v1.LongRunningService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewLongRunningServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) LongRunningServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &longRunningServiceClient{
		registerOperation: connect_go.NewClient[v1.RegisterOperationRequest, v1.Operation](
			httpClient,
			baseURL+LongRunningServiceRegisterOperationProcedure,
			opts...,
		),
		updateOperation: connect_go.NewClient[v1.UpdateOperationRequest, v1.Operation](
			httpClient,
			baseURL+LongRunningServiceUpdateOperationProcedure,
			opts...,
		),
		completeOperation: connect_go.NewClient[v1.CompleteOperationRequest, v1.Operation](
			httpClient,
			baseURL+LongRunningServiceCompleteOperationProcedure,
			opts...,
		),
		queryOperations: connect_go.NewClient[v1.QueryOperationsRequest, v1.QueryOperationsResponse](
			httpClient,
			baseURL+LongRunningServiceQueryOperationsProcedure,
			opts...,
		),
		getOperation: connect_go.NewClient[v1.GetOperationRequest, v1.Operation](
			httpClient,
			baseURL+LongRunningServiceGetOperationProcedure,
			opts...,
		),
	}
}

// longRunningServiceClient implements LongRunningServiceClient.
type longRunningServiceClient struct {
	registerOperation *connect_go.Client[v1.RegisterOperationRequest, v1.Operation]
	updateOperation   *connect_go.Client[v1.UpdateOperationRequest, v1.Operation]
	completeOperation *connect_go.Client[v1.CompleteOperationRequest, v1.Operation]
	queryOperations   *connect_go.Client[v1.QueryOperationsRequest, v1.QueryOperationsResponse]
	getOperation      *connect_go.Client[v1.GetOperationRequest, v1.Operation]
}

// RegisterOperation calls tkd.longrunning.v1.LongRunningService.RegisterOperation.
func (c *longRunningServiceClient) RegisterOperation(ctx context.Context, req *connect_go.Request[v1.RegisterOperationRequest]) (*connect_go.Response[v1.Operation], error) {
	return c.registerOperation.CallUnary(ctx, req)
}

// UpdateOperation calls tkd.longrunning.v1.LongRunningService.UpdateOperation.
func (c *longRunningServiceClient) UpdateOperation(ctx context.Context, req *connect_go.Request[v1.UpdateOperationRequest]) (*connect_go.Response[v1.Operation], error) {
	return c.updateOperation.CallUnary(ctx, req)
}

// CompleteOperation calls tkd.longrunning.v1.LongRunningService.CompleteOperation.
func (c *longRunningServiceClient) CompleteOperation(ctx context.Context, req *connect_go.Request[v1.CompleteOperationRequest]) (*connect_go.Response[v1.Operation], error) {
	return c.completeOperation.CallUnary(ctx, req)
}

// QueryOperations calls tkd.longrunning.v1.LongRunningService.QueryOperations.
func (c *longRunningServiceClient) QueryOperations(ctx context.Context, req *connect_go.Request[v1.QueryOperationsRequest]) (*connect_go.Response[v1.QueryOperationsResponse], error) {
	return c.queryOperations.CallUnary(ctx, req)
}

// GetOperation calls tkd.longrunning.v1.LongRunningService.GetOperation.
func (c *longRunningServiceClient) GetOperation(ctx context.Context, req *connect_go.Request[v1.GetOperationRequest]) (*connect_go.Response[v1.Operation], error) {
	return c.getOperation.CallUnary(ctx, req)
}

// LongRunningServiceHandler is an implementation of the tkd.longrunning.v1.LongRunningService
// service.
type LongRunningServiceHandler interface {
	RegisterOperation(context.Context, *connect_go.Request[v1.RegisterOperationRequest]) (*connect_go.Response[v1.Operation], error)
	UpdateOperation(context.Context, *connect_go.Request[v1.UpdateOperationRequest]) (*connect_go.Response[v1.Operation], error)
	CompleteOperation(context.Context, *connect_go.Request[v1.CompleteOperationRequest]) (*connect_go.Response[v1.Operation], error)
	QueryOperations(context.Context, *connect_go.Request[v1.QueryOperationsRequest]) (*connect_go.Response[v1.QueryOperationsResponse], error)
	GetOperation(context.Context, *connect_go.Request[v1.GetOperationRequest]) (*connect_go.Response[v1.Operation], error)
}

// NewLongRunningServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewLongRunningServiceHandler(svc LongRunningServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	longRunningServiceRegisterOperationHandler := connect_go.NewUnaryHandler(
		LongRunningServiceRegisterOperationProcedure,
		svc.RegisterOperation,
		opts...,
	)
	longRunningServiceUpdateOperationHandler := connect_go.NewUnaryHandler(
		LongRunningServiceUpdateOperationProcedure,
		svc.UpdateOperation,
		opts...,
	)
	longRunningServiceCompleteOperationHandler := connect_go.NewUnaryHandler(
		LongRunningServiceCompleteOperationProcedure,
		svc.CompleteOperation,
		opts...,
	)
	longRunningServiceQueryOperationsHandler := connect_go.NewUnaryHandler(
		LongRunningServiceQueryOperationsProcedure,
		svc.QueryOperations,
		opts...,
	)
	longRunningServiceGetOperationHandler := connect_go.NewUnaryHandler(
		LongRunningServiceGetOperationProcedure,
		svc.GetOperation,
		opts...,
	)
	return "/tkd.longrunning.v1.LongRunningService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case LongRunningServiceRegisterOperationProcedure:
			longRunningServiceRegisterOperationHandler.ServeHTTP(w, r)
		case LongRunningServiceUpdateOperationProcedure:
			longRunningServiceUpdateOperationHandler.ServeHTTP(w, r)
		case LongRunningServiceCompleteOperationProcedure:
			longRunningServiceCompleteOperationHandler.ServeHTTP(w, r)
		case LongRunningServiceQueryOperationsProcedure:
			longRunningServiceQueryOperationsHandler.ServeHTTP(w, r)
		case LongRunningServiceGetOperationProcedure:
			longRunningServiceGetOperationHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedLongRunningServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedLongRunningServiceHandler struct{}

func (UnimplementedLongRunningServiceHandler) RegisterOperation(context.Context, *connect_go.Request[v1.RegisterOperationRequest]) (*connect_go.Response[v1.Operation], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.longrunning.v1.LongRunningService.RegisterOperation is not implemented"))
}

func (UnimplementedLongRunningServiceHandler) UpdateOperation(context.Context, *connect_go.Request[v1.UpdateOperationRequest]) (*connect_go.Response[v1.Operation], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.longrunning.v1.LongRunningService.UpdateOperation is not implemented"))
}

func (UnimplementedLongRunningServiceHandler) CompleteOperation(context.Context, *connect_go.Request[v1.CompleteOperationRequest]) (*connect_go.Response[v1.Operation], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.longrunning.v1.LongRunningService.CompleteOperation is not implemented"))
}

func (UnimplementedLongRunningServiceHandler) QueryOperations(context.Context, *connect_go.Request[v1.QueryOperationsRequest]) (*connect_go.Response[v1.QueryOperationsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.longrunning.v1.LongRunningService.QueryOperations is not implemented"))
}

func (UnimplementedLongRunningServiceHandler) GetOperation(context.Context, *connect_go.Request[v1.GetOperationRequest]) (*connect_go.Response[v1.Operation], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.longrunning.v1.LongRunningService.GetOperation is not implemented"))
}
