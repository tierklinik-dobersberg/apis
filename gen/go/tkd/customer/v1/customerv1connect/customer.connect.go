// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: tkd/customer/v1/customer.proto

package customerv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/customer/v1"
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
	// CustomerServiceName is the fully-qualified name of the CustomerService service.
	CustomerServiceName = "tkd.customer.v1.CustomerService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// CustomerServiceSearchCustomerProcedure is the fully-qualified name of the CustomerService's
	// SearchCustomer RPC.
	CustomerServiceSearchCustomerProcedure = "/tkd.customer.v1.CustomerService/SearchCustomer"
	// CustomerServiceSearchCustomerStreamProcedure is the fully-qualified name of the CustomerService's
	// SearchCustomerStream RPC.
	CustomerServiceSearchCustomerStreamProcedure = "/tkd.customer.v1.CustomerService/SearchCustomerStream"
	// CustomerServiceUpdateCustomerProcedure is the fully-qualified name of the CustomerService's
	// UpdateCustomer RPC.
	CustomerServiceUpdateCustomerProcedure = "/tkd.customer.v1.CustomerService/UpdateCustomer"
)

// CustomerServiceClient is a client for the tkd.customer.v1.CustomerService service.
type CustomerServiceClient interface {
	SearchCustomer(context.Context, *connect_go.Request[v1.SearchCustomerRequest]) (*connect_go.Response[v1.SearchCustomerResponse], error)
	SearchCustomerStream(context.Context) *connect_go.BidiStreamForClient[v1.SearchCustomerRequest, v1.SearchCustomerResponse]
	UpdateCustomer(context.Context, *connect_go.Request[v1.UpdateCustomerRequest]) (*connect_go.Response[v1.UpdateCustomerResponse], error)
}

// NewCustomerServiceClient constructs a client for the tkd.customer.v1.CustomerService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCustomerServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) CustomerServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &customerServiceClient{
		searchCustomer: connect_go.NewClient[v1.SearchCustomerRequest, v1.SearchCustomerResponse](
			httpClient,
			baseURL+CustomerServiceSearchCustomerProcedure,
			opts...,
		),
		searchCustomerStream: connect_go.NewClient[v1.SearchCustomerRequest, v1.SearchCustomerResponse](
			httpClient,
			baseURL+CustomerServiceSearchCustomerStreamProcedure,
			opts...,
		),
		updateCustomer: connect_go.NewClient[v1.UpdateCustomerRequest, v1.UpdateCustomerResponse](
			httpClient,
			baseURL+CustomerServiceUpdateCustomerProcedure,
			opts...,
		),
	}
}

// customerServiceClient implements CustomerServiceClient.
type customerServiceClient struct {
	searchCustomer       *connect_go.Client[v1.SearchCustomerRequest, v1.SearchCustomerResponse]
	searchCustomerStream *connect_go.Client[v1.SearchCustomerRequest, v1.SearchCustomerResponse]
	updateCustomer       *connect_go.Client[v1.UpdateCustomerRequest, v1.UpdateCustomerResponse]
}

// SearchCustomer calls tkd.customer.v1.CustomerService.SearchCustomer.
func (c *customerServiceClient) SearchCustomer(ctx context.Context, req *connect_go.Request[v1.SearchCustomerRequest]) (*connect_go.Response[v1.SearchCustomerResponse], error) {
	return c.searchCustomer.CallUnary(ctx, req)
}

// SearchCustomerStream calls tkd.customer.v1.CustomerService.SearchCustomerStream.
func (c *customerServiceClient) SearchCustomerStream(ctx context.Context) *connect_go.BidiStreamForClient[v1.SearchCustomerRequest, v1.SearchCustomerResponse] {
	return c.searchCustomerStream.CallBidiStream(ctx)
}

// UpdateCustomer calls tkd.customer.v1.CustomerService.UpdateCustomer.
func (c *customerServiceClient) UpdateCustomer(ctx context.Context, req *connect_go.Request[v1.UpdateCustomerRequest]) (*connect_go.Response[v1.UpdateCustomerResponse], error) {
	return c.updateCustomer.CallUnary(ctx, req)
}

// CustomerServiceHandler is an implementation of the tkd.customer.v1.CustomerService service.
type CustomerServiceHandler interface {
	SearchCustomer(context.Context, *connect_go.Request[v1.SearchCustomerRequest]) (*connect_go.Response[v1.SearchCustomerResponse], error)
	SearchCustomerStream(context.Context, *connect_go.BidiStream[v1.SearchCustomerRequest, v1.SearchCustomerResponse]) error
	UpdateCustomer(context.Context, *connect_go.Request[v1.UpdateCustomerRequest]) (*connect_go.Response[v1.UpdateCustomerResponse], error)
}

// NewCustomerServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCustomerServiceHandler(svc CustomerServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	customerServiceSearchCustomerHandler := connect_go.NewUnaryHandler(
		CustomerServiceSearchCustomerProcedure,
		svc.SearchCustomer,
		opts...,
	)
	customerServiceSearchCustomerStreamHandler := connect_go.NewBidiStreamHandler(
		CustomerServiceSearchCustomerStreamProcedure,
		svc.SearchCustomerStream,
		opts...,
	)
	customerServiceUpdateCustomerHandler := connect_go.NewUnaryHandler(
		CustomerServiceUpdateCustomerProcedure,
		svc.UpdateCustomer,
		opts...,
	)
	return "/tkd.customer.v1.CustomerService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case CustomerServiceSearchCustomerProcedure:
			customerServiceSearchCustomerHandler.ServeHTTP(w, r)
		case CustomerServiceSearchCustomerStreamProcedure:
			customerServiceSearchCustomerStreamHandler.ServeHTTP(w, r)
		case CustomerServiceUpdateCustomerProcedure:
			customerServiceUpdateCustomerHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedCustomerServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedCustomerServiceHandler struct{}

func (UnimplementedCustomerServiceHandler) SearchCustomer(context.Context, *connect_go.Request[v1.SearchCustomerRequest]) (*connect_go.Response[v1.SearchCustomerResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.customer.v1.CustomerService.SearchCustomer is not implemented"))
}

func (UnimplementedCustomerServiceHandler) SearchCustomerStream(context.Context, *connect_go.BidiStream[v1.SearchCustomerRequest, v1.SearchCustomerResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.customer.v1.CustomerService.SearchCustomerStream is not implemented"))
}

func (UnimplementedCustomerServiceHandler) UpdateCustomer(context.Context, *connect_go.Request[v1.UpdateCustomerRequest]) (*connect_go.Response[v1.UpdateCustomerResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.customer.v1.CustomerService.UpdateCustomer is not implemented"))
}
