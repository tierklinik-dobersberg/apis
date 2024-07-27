// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: tkd/pbx3cx/v1/calllog.proto

package pbx3cxv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/pbx3cx/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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
	// CallServiceName is the fully-qualified name of the CallService service.
	CallServiceName = "tkd.pbx3cx.v1.CallService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// CallServiceRecordCallProcedure is the fully-qualified name of the CallService's RecordCall RPC.
	CallServiceRecordCallProcedure = "/tkd.pbx3cx.v1.CallService/RecordCall"
	// CallServiceGetOnCallProcedure is the fully-qualified name of the CallService's GetOnCall RPC.
	CallServiceGetOnCallProcedure = "/tkd.pbx3cx.v1.CallService/GetOnCall"
	// CallServiceCreateInboundNumberProcedure is the fully-qualified name of the CallService's
	// CreateInboundNumber RPC.
	CallServiceCreateInboundNumberProcedure = "/tkd.pbx3cx.v1.CallService/CreateInboundNumber"
	// CallServiceUpdateInboundNumberProcedure is the fully-qualified name of the CallService's
	// UpdateInboundNumber RPC.
	CallServiceUpdateInboundNumberProcedure = "/tkd.pbx3cx.v1.CallService/UpdateInboundNumber"
	// CallServiceDeleteInboundNumberProcedure is the fully-qualified name of the CallService's
	// DeleteInboundNumber RPC.
	CallServiceDeleteInboundNumberProcedure = "/tkd.pbx3cx.v1.CallService/DeleteInboundNumber"
	// CallServiceListInboundNumberProcedure is the fully-qualified name of the CallService's
	// ListInboundNumber RPC.
	CallServiceListInboundNumberProcedure = "/tkd.pbx3cx.v1.CallService/ListInboundNumber"
	// CallServiceCreateOverwriteProcedure is the fully-qualified name of the CallService's
	// CreateOverwrite RPC.
	CallServiceCreateOverwriteProcedure = "/tkd.pbx3cx.v1.CallService/CreateOverwrite"
	// CallServiceDeleteOverwriteProcedure is the fully-qualified name of the CallService's
	// DeleteOverwrite RPC.
	CallServiceDeleteOverwriteProcedure = "/tkd.pbx3cx.v1.CallService/DeleteOverwrite"
	// CallServiceGetOverwriteProcedure is the fully-qualified name of the CallService's GetOverwrite
	// RPC.
	CallServiceGetOverwriteProcedure = "/tkd.pbx3cx.v1.CallService/GetOverwrite"
	// CallServiceGetLogsForDateProcedure is the fully-qualified name of the CallService's
	// GetLogsForDate RPC.
	CallServiceGetLogsForDateProcedure = "/tkd.pbx3cx.v1.CallService/GetLogsForDate"
	// CallServiceSearchCallLogsProcedure is the fully-qualified name of the CallService's
	// SearchCallLogs RPC.
	CallServiceSearchCallLogsProcedure = "/tkd.pbx3cx.v1.CallService/SearchCallLogs"
	// CallServiceGetLogsForCustomerProcedure is the fully-qualified name of the CallService's
	// GetLogsForCustomer RPC.
	CallServiceGetLogsForCustomerProcedure = "/tkd.pbx3cx.v1.CallService/GetLogsForCustomer"
)

// CallServiceClient is a client for the tkd.pbx3cx.v1.CallService service.
type CallServiceClient interface {
	// Recording APIS
	RecordCall(context.Context, *connect_go.Request[v1.RecordCallRequest]) (*connect_go.Response[emptypb.Empty], error)
	// On-Duty/Call APIs
	GetOnCall(context.Context, *connect_go.Request[v1.GetOnCallRequest]) (*connect_go.Response[v1.GetOnCallResponse], error)
	CreateInboundNumber(context.Context, *connect_go.Request[v1.CreateInboundNumberRequest]) (*connect_go.Response[v1.CreateInboundNumberResponse], error)
	UpdateInboundNumber(context.Context, *connect_go.Request[v1.UpdateInboundNumberRequest]) (*connect_go.Response[v1.UpdateInboundNumberResponse], error)
	DeleteInboundNumber(context.Context, *connect_go.Request[v1.DeleteInboundNumberRequest]) (*connect_go.Response[v1.ListInboundNumberResponse], error)
	ListInboundNumber(context.Context, *connect_go.Request[v1.ListInboundNumberRequest]) (*connect_go.Response[v1.ListInboundNumberResponse], error)
	// Overwrite APIS
	CreateOverwrite(context.Context, *connect_go.Request[v1.CreateOverwriteRequest]) (*connect_go.Response[v1.CreateOverwriteResponse], error)
	DeleteOverwrite(context.Context, *connect_go.Request[v1.DeleteOverwriteRequest]) (*connect_go.Response[v1.DeleteOverwriteResponse], error)
	GetOverwrite(context.Context, *connect_go.Request[v1.GetOverwriteRequest]) (*connect_go.Response[v1.GetOverwriteResponse], error)
	// Calllog APIs
	GetLogsForDate(context.Context, *connect_go.Request[v1.GetLogsForDateRequest]) (*connect_go.Response[v1.GetLogsForDateResponse], error)
	SearchCallLogs(context.Context, *connect_go.Request[v1.SearchCallLogsRequest]) (*connect_go.Response[v1.SearchCallLogsResponse], error)
	GetLogsForCustomer(context.Context, *connect_go.Request[v1.GetLogsForCustomerRequest]) (*connect_go.Response[v1.GetLogsForCustomerResponse], error)
}

// NewCallServiceClient constructs a client for the tkd.pbx3cx.v1.CallService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCallServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) CallServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &callServiceClient{
		recordCall: connect_go.NewClient[v1.RecordCallRequest, emptypb.Empty](
			httpClient,
			baseURL+CallServiceRecordCallProcedure,
			opts...,
		),
		getOnCall: connect_go.NewClient[v1.GetOnCallRequest, v1.GetOnCallResponse](
			httpClient,
			baseURL+CallServiceGetOnCallProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		createInboundNumber: connect_go.NewClient[v1.CreateInboundNumberRequest, v1.CreateInboundNumberResponse](
			httpClient,
			baseURL+CallServiceCreateInboundNumberProcedure,
			opts...,
		),
		updateInboundNumber: connect_go.NewClient[v1.UpdateInboundNumberRequest, v1.UpdateInboundNumberResponse](
			httpClient,
			baseURL+CallServiceUpdateInboundNumberProcedure,
			opts...,
		),
		deleteInboundNumber: connect_go.NewClient[v1.DeleteInboundNumberRequest, v1.ListInboundNumberResponse](
			httpClient,
			baseURL+CallServiceDeleteInboundNumberProcedure,
			opts...,
		),
		listInboundNumber: connect_go.NewClient[v1.ListInboundNumberRequest, v1.ListInboundNumberResponse](
			httpClient,
			baseURL+CallServiceListInboundNumberProcedure,
			opts...,
		),
		createOverwrite: connect_go.NewClient[v1.CreateOverwriteRequest, v1.CreateOverwriteResponse](
			httpClient,
			baseURL+CallServiceCreateOverwriteProcedure,
			opts...,
		),
		deleteOverwrite: connect_go.NewClient[v1.DeleteOverwriteRequest, v1.DeleteOverwriteResponse](
			httpClient,
			baseURL+CallServiceDeleteOverwriteProcedure,
			opts...,
		),
		getOverwrite: connect_go.NewClient[v1.GetOverwriteRequest, v1.GetOverwriteResponse](
			httpClient,
			baseURL+CallServiceGetOverwriteProcedure,
			opts...,
		),
		getLogsForDate: connect_go.NewClient[v1.GetLogsForDateRequest, v1.GetLogsForDateResponse](
			httpClient,
			baseURL+CallServiceGetLogsForDateProcedure,
			opts...,
		),
		searchCallLogs: connect_go.NewClient[v1.SearchCallLogsRequest, v1.SearchCallLogsResponse](
			httpClient,
			baseURL+CallServiceSearchCallLogsProcedure,
			opts...,
		),
		getLogsForCustomer: connect_go.NewClient[v1.GetLogsForCustomerRequest, v1.GetLogsForCustomerResponse](
			httpClient,
			baseURL+CallServiceGetLogsForCustomerProcedure,
			opts...,
		),
	}
}

// callServiceClient implements CallServiceClient.
type callServiceClient struct {
	recordCall          *connect_go.Client[v1.RecordCallRequest, emptypb.Empty]
	getOnCall           *connect_go.Client[v1.GetOnCallRequest, v1.GetOnCallResponse]
	createInboundNumber *connect_go.Client[v1.CreateInboundNumberRequest, v1.CreateInboundNumberResponse]
	updateInboundNumber *connect_go.Client[v1.UpdateInboundNumberRequest, v1.UpdateInboundNumberResponse]
	deleteInboundNumber *connect_go.Client[v1.DeleteInboundNumberRequest, v1.ListInboundNumberResponse]
	listInboundNumber   *connect_go.Client[v1.ListInboundNumberRequest, v1.ListInboundNumberResponse]
	createOverwrite     *connect_go.Client[v1.CreateOverwriteRequest, v1.CreateOverwriteResponse]
	deleteOverwrite     *connect_go.Client[v1.DeleteOverwriteRequest, v1.DeleteOverwriteResponse]
	getOverwrite        *connect_go.Client[v1.GetOverwriteRequest, v1.GetOverwriteResponse]
	getLogsForDate      *connect_go.Client[v1.GetLogsForDateRequest, v1.GetLogsForDateResponse]
	searchCallLogs      *connect_go.Client[v1.SearchCallLogsRequest, v1.SearchCallLogsResponse]
	getLogsForCustomer  *connect_go.Client[v1.GetLogsForCustomerRequest, v1.GetLogsForCustomerResponse]
}

// RecordCall calls tkd.pbx3cx.v1.CallService.RecordCall.
func (c *callServiceClient) RecordCall(ctx context.Context, req *connect_go.Request[v1.RecordCallRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.recordCall.CallUnary(ctx, req)
}

// GetOnCall calls tkd.pbx3cx.v1.CallService.GetOnCall.
func (c *callServiceClient) GetOnCall(ctx context.Context, req *connect_go.Request[v1.GetOnCallRequest]) (*connect_go.Response[v1.GetOnCallResponse], error) {
	return c.getOnCall.CallUnary(ctx, req)
}

// CreateInboundNumber calls tkd.pbx3cx.v1.CallService.CreateInboundNumber.
func (c *callServiceClient) CreateInboundNumber(ctx context.Context, req *connect_go.Request[v1.CreateInboundNumberRequest]) (*connect_go.Response[v1.CreateInboundNumberResponse], error) {
	return c.createInboundNumber.CallUnary(ctx, req)
}

// UpdateInboundNumber calls tkd.pbx3cx.v1.CallService.UpdateInboundNumber.
func (c *callServiceClient) UpdateInboundNumber(ctx context.Context, req *connect_go.Request[v1.UpdateInboundNumberRequest]) (*connect_go.Response[v1.UpdateInboundNumberResponse], error) {
	return c.updateInboundNumber.CallUnary(ctx, req)
}

// DeleteInboundNumber calls tkd.pbx3cx.v1.CallService.DeleteInboundNumber.
func (c *callServiceClient) DeleteInboundNumber(ctx context.Context, req *connect_go.Request[v1.DeleteInboundNumberRequest]) (*connect_go.Response[v1.ListInboundNumberResponse], error) {
	return c.deleteInboundNumber.CallUnary(ctx, req)
}

// ListInboundNumber calls tkd.pbx3cx.v1.CallService.ListInboundNumber.
func (c *callServiceClient) ListInboundNumber(ctx context.Context, req *connect_go.Request[v1.ListInboundNumberRequest]) (*connect_go.Response[v1.ListInboundNumberResponse], error) {
	return c.listInboundNumber.CallUnary(ctx, req)
}

// CreateOverwrite calls tkd.pbx3cx.v1.CallService.CreateOverwrite.
func (c *callServiceClient) CreateOverwrite(ctx context.Context, req *connect_go.Request[v1.CreateOverwriteRequest]) (*connect_go.Response[v1.CreateOverwriteResponse], error) {
	return c.createOverwrite.CallUnary(ctx, req)
}

// DeleteOverwrite calls tkd.pbx3cx.v1.CallService.DeleteOverwrite.
func (c *callServiceClient) DeleteOverwrite(ctx context.Context, req *connect_go.Request[v1.DeleteOverwriteRequest]) (*connect_go.Response[v1.DeleteOverwriteResponse], error) {
	return c.deleteOverwrite.CallUnary(ctx, req)
}

// GetOverwrite calls tkd.pbx3cx.v1.CallService.GetOverwrite.
func (c *callServiceClient) GetOverwrite(ctx context.Context, req *connect_go.Request[v1.GetOverwriteRequest]) (*connect_go.Response[v1.GetOverwriteResponse], error) {
	return c.getOverwrite.CallUnary(ctx, req)
}

// GetLogsForDate calls tkd.pbx3cx.v1.CallService.GetLogsForDate.
func (c *callServiceClient) GetLogsForDate(ctx context.Context, req *connect_go.Request[v1.GetLogsForDateRequest]) (*connect_go.Response[v1.GetLogsForDateResponse], error) {
	return c.getLogsForDate.CallUnary(ctx, req)
}

// SearchCallLogs calls tkd.pbx3cx.v1.CallService.SearchCallLogs.
func (c *callServiceClient) SearchCallLogs(ctx context.Context, req *connect_go.Request[v1.SearchCallLogsRequest]) (*connect_go.Response[v1.SearchCallLogsResponse], error) {
	return c.searchCallLogs.CallUnary(ctx, req)
}

// GetLogsForCustomer calls tkd.pbx3cx.v1.CallService.GetLogsForCustomer.
func (c *callServiceClient) GetLogsForCustomer(ctx context.Context, req *connect_go.Request[v1.GetLogsForCustomerRequest]) (*connect_go.Response[v1.GetLogsForCustomerResponse], error) {
	return c.getLogsForCustomer.CallUnary(ctx, req)
}

// CallServiceHandler is an implementation of the tkd.pbx3cx.v1.CallService service.
type CallServiceHandler interface {
	// Recording APIS
	RecordCall(context.Context, *connect_go.Request[v1.RecordCallRequest]) (*connect_go.Response[emptypb.Empty], error)
	// On-Duty/Call APIs
	GetOnCall(context.Context, *connect_go.Request[v1.GetOnCallRequest]) (*connect_go.Response[v1.GetOnCallResponse], error)
	CreateInboundNumber(context.Context, *connect_go.Request[v1.CreateInboundNumberRequest]) (*connect_go.Response[v1.CreateInboundNumberResponse], error)
	UpdateInboundNumber(context.Context, *connect_go.Request[v1.UpdateInboundNumberRequest]) (*connect_go.Response[v1.UpdateInboundNumberResponse], error)
	DeleteInboundNumber(context.Context, *connect_go.Request[v1.DeleteInboundNumberRequest]) (*connect_go.Response[v1.ListInboundNumberResponse], error)
	ListInboundNumber(context.Context, *connect_go.Request[v1.ListInboundNumberRequest]) (*connect_go.Response[v1.ListInboundNumberResponse], error)
	// Overwrite APIS
	CreateOverwrite(context.Context, *connect_go.Request[v1.CreateOverwriteRequest]) (*connect_go.Response[v1.CreateOverwriteResponse], error)
	DeleteOverwrite(context.Context, *connect_go.Request[v1.DeleteOverwriteRequest]) (*connect_go.Response[v1.DeleteOverwriteResponse], error)
	GetOverwrite(context.Context, *connect_go.Request[v1.GetOverwriteRequest]) (*connect_go.Response[v1.GetOverwriteResponse], error)
	// Calllog APIs
	GetLogsForDate(context.Context, *connect_go.Request[v1.GetLogsForDateRequest]) (*connect_go.Response[v1.GetLogsForDateResponse], error)
	SearchCallLogs(context.Context, *connect_go.Request[v1.SearchCallLogsRequest]) (*connect_go.Response[v1.SearchCallLogsResponse], error)
	GetLogsForCustomer(context.Context, *connect_go.Request[v1.GetLogsForCustomerRequest]) (*connect_go.Response[v1.GetLogsForCustomerResponse], error)
}

// NewCallServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCallServiceHandler(svc CallServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	callServiceRecordCallHandler := connect_go.NewUnaryHandler(
		CallServiceRecordCallProcedure,
		svc.RecordCall,
		opts...,
	)
	callServiceGetOnCallHandler := connect_go.NewUnaryHandler(
		CallServiceGetOnCallProcedure,
		svc.GetOnCall,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	)
	callServiceCreateInboundNumberHandler := connect_go.NewUnaryHandler(
		CallServiceCreateInboundNumberProcedure,
		svc.CreateInboundNumber,
		opts...,
	)
	callServiceUpdateInboundNumberHandler := connect_go.NewUnaryHandler(
		CallServiceUpdateInboundNumberProcedure,
		svc.UpdateInboundNumber,
		opts...,
	)
	callServiceDeleteInboundNumberHandler := connect_go.NewUnaryHandler(
		CallServiceDeleteInboundNumberProcedure,
		svc.DeleteInboundNumber,
		opts...,
	)
	callServiceListInboundNumberHandler := connect_go.NewUnaryHandler(
		CallServiceListInboundNumberProcedure,
		svc.ListInboundNumber,
		opts...,
	)
	callServiceCreateOverwriteHandler := connect_go.NewUnaryHandler(
		CallServiceCreateOverwriteProcedure,
		svc.CreateOverwrite,
		opts...,
	)
	callServiceDeleteOverwriteHandler := connect_go.NewUnaryHandler(
		CallServiceDeleteOverwriteProcedure,
		svc.DeleteOverwrite,
		opts...,
	)
	callServiceGetOverwriteHandler := connect_go.NewUnaryHandler(
		CallServiceGetOverwriteProcedure,
		svc.GetOverwrite,
		opts...,
	)
	callServiceGetLogsForDateHandler := connect_go.NewUnaryHandler(
		CallServiceGetLogsForDateProcedure,
		svc.GetLogsForDate,
		opts...,
	)
	callServiceSearchCallLogsHandler := connect_go.NewUnaryHandler(
		CallServiceSearchCallLogsProcedure,
		svc.SearchCallLogs,
		opts...,
	)
	callServiceGetLogsForCustomerHandler := connect_go.NewUnaryHandler(
		CallServiceGetLogsForCustomerProcedure,
		svc.GetLogsForCustomer,
		opts...,
	)
	return "/tkd.pbx3cx.v1.CallService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case CallServiceRecordCallProcedure:
			callServiceRecordCallHandler.ServeHTTP(w, r)
		case CallServiceGetOnCallProcedure:
			callServiceGetOnCallHandler.ServeHTTP(w, r)
		case CallServiceCreateInboundNumberProcedure:
			callServiceCreateInboundNumberHandler.ServeHTTP(w, r)
		case CallServiceUpdateInboundNumberProcedure:
			callServiceUpdateInboundNumberHandler.ServeHTTP(w, r)
		case CallServiceDeleteInboundNumberProcedure:
			callServiceDeleteInboundNumberHandler.ServeHTTP(w, r)
		case CallServiceListInboundNumberProcedure:
			callServiceListInboundNumberHandler.ServeHTTP(w, r)
		case CallServiceCreateOverwriteProcedure:
			callServiceCreateOverwriteHandler.ServeHTTP(w, r)
		case CallServiceDeleteOverwriteProcedure:
			callServiceDeleteOverwriteHandler.ServeHTTP(w, r)
		case CallServiceGetOverwriteProcedure:
			callServiceGetOverwriteHandler.ServeHTTP(w, r)
		case CallServiceGetLogsForDateProcedure:
			callServiceGetLogsForDateHandler.ServeHTTP(w, r)
		case CallServiceSearchCallLogsProcedure:
			callServiceSearchCallLogsHandler.ServeHTTP(w, r)
		case CallServiceGetLogsForCustomerProcedure:
			callServiceGetLogsForCustomerHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedCallServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedCallServiceHandler struct{}

func (UnimplementedCallServiceHandler) RecordCall(context.Context, *connect_go.Request[v1.RecordCallRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.pbx3cx.v1.CallService.RecordCall is not implemented"))
}

func (UnimplementedCallServiceHandler) GetOnCall(context.Context, *connect_go.Request[v1.GetOnCallRequest]) (*connect_go.Response[v1.GetOnCallResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.pbx3cx.v1.CallService.GetOnCall is not implemented"))
}

func (UnimplementedCallServiceHandler) CreateInboundNumber(context.Context, *connect_go.Request[v1.CreateInboundNumberRequest]) (*connect_go.Response[v1.CreateInboundNumberResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.pbx3cx.v1.CallService.CreateInboundNumber is not implemented"))
}

func (UnimplementedCallServiceHandler) UpdateInboundNumber(context.Context, *connect_go.Request[v1.UpdateInboundNumberRequest]) (*connect_go.Response[v1.UpdateInboundNumberResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.pbx3cx.v1.CallService.UpdateInboundNumber is not implemented"))
}

func (UnimplementedCallServiceHandler) DeleteInboundNumber(context.Context, *connect_go.Request[v1.DeleteInboundNumberRequest]) (*connect_go.Response[v1.ListInboundNumberResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.pbx3cx.v1.CallService.DeleteInboundNumber is not implemented"))
}

func (UnimplementedCallServiceHandler) ListInboundNumber(context.Context, *connect_go.Request[v1.ListInboundNumberRequest]) (*connect_go.Response[v1.ListInboundNumberResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.pbx3cx.v1.CallService.ListInboundNumber is not implemented"))
}

func (UnimplementedCallServiceHandler) CreateOverwrite(context.Context, *connect_go.Request[v1.CreateOverwriteRequest]) (*connect_go.Response[v1.CreateOverwriteResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.pbx3cx.v1.CallService.CreateOverwrite is not implemented"))
}

func (UnimplementedCallServiceHandler) DeleteOverwrite(context.Context, *connect_go.Request[v1.DeleteOverwriteRequest]) (*connect_go.Response[v1.DeleteOverwriteResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.pbx3cx.v1.CallService.DeleteOverwrite is not implemented"))
}

func (UnimplementedCallServiceHandler) GetOverwrite(context.Context, *connect_go.Request[v1.GetOverwriteRequest]) (*connect_go.Response[v1.GetOverwriteResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.pbx3cx.v1.CallService.GetOverwrite is not implemented"))
}

func (UnimplementedCallServiceHandler) GetLogsForDate(context.Context, *connect_go.Request[v1.GetLogsForDateRequest]) (*connect_go.Response[v1.GetLogsForDateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.pbx3cx.v1.CallService.GetLogsForDate is not implemented"))
}

func (UnimplementedCallServiceHandler) SearchCallLogs(context.Context, *connect_go.Request[v1.SearchCallLogsRequest]) (*connect_go.Response[v1.SearchCallLogsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.pbx3cx.v1.CallService.SearchCallLogs is not implemented"))
}

func (UnimplementedCallServiceHandler) GetLogsForCustomer(context.Context, *connect_go.Request[v1.GetLogsForCustomerRequest]) (*connect_go.Response[v1.GetLogsForCustomerResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.pbx3cx.v1.CallService.GetLogsForCustomer is not implemented"))
}
