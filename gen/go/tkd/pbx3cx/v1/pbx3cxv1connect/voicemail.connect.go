// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: tkd/pbx3cx/v1/voicemail.proto

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
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// VoiceMailServiceName is the fully-qualified name of the VoiceMailService service.
	VoiceMailServiceName = "tkd.pbx3cx.v1.VoiceMailService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// VoiceMailServiceCreateMailboxProcedure is the fully-qualified name of the VoiceMailService's
	// CreateMailbox RPC.
	VoiceMailServiceCreateMailboxProcedure = "/tkd.pbx3cx.v1.VoiceMailService/CreateMailbox"
	// VoiceMailServiceListMailboxesProcedure is the fully-qualified name of the VoiceMailService's
	// ListMailboxes RPC.
	VoiceMailServiceListMailboxesProcedure = "/tkd.pbx3cx.v1.VoiceMailService/ListMailboxes"
	// VoiceMailServiceDeleteMailboxProcedure is the fully-qualified name of the VoiceMailService's
	// DeleteMailbox RPC.
	VoiceMailServiceDeleteMailboxProcedure = "/tkd.pbx3cx.v1.VoiceMailService/DeleteMailbox"
	// VoiceMailServiceUpdateMailboxProcedure is the fully-qualified name of the VoiceMailService's
	// UpdateMailbox RPC.
	VoiceMailServiceUpdateMailboxProcedure = "/tkd.pbx3cx.v1.VoiceMailService/UpdateMailbox"
	// VoiceMailServiceListVoiceMailsProcedure is the fully-qualified name of the VoiceMailService's
	// ListVoiceMails RPC.
	VoiceMailServiceListVoiceMailsProcedure = "/tkd.pbx3cx.v1.VoiceMailService/ListVoiceMails"
	// VoiceMailServiceGetVoiceMailProcedure is the fully-qualified name of the VoiceMailService's
	// GetVoiceMail RPC.
	VoiceMailServiceGetVoiceMailProcedure = "/tkd.pbx3cx.v1.VoiceMailService/GetVoiceMail"
	// VoiceMailServiceMarkVoiceMailsProcedure is the fully-qualified name of the VoiceMailService's
	// MarkVoiceMails RPC.
	VoiceMailServiceMarkVoiceMailsProcedure = "/tkd.pbx3cx.v1.VoiceMailService/MarkVoiceMails"
	// VoiceMailServiceSearchVoiceMailsProcedure is the fully-qualified name of the VoiceMailService's
	// SearchVoiceMails RPC.
	VoiceMailServiceSearchVoiceMailsProcedure = "/tkd.pbx3cx.v1.VoiceMailService/SearchVoiceMails"
)

// VoiceMailServiceClient is a client for the tkd.pbx3cx.v1.VoiceMailService service.
type VoiceMailServiceClient interface {
	CreateMailbox(context.Context, *connect_go.Request[v1.CreateMailboxRequest]) (*connect_go.Response[v1.CreateMailboxResponse], error)
	ListMailboxes(context.Context, *connect_go.Request[v1.ListMailboxesRequest]) (*connect_go.Response[v1.ListMailboxesResponse], error)
	DeleteMailbox(context.Context, *connect_go.Request[v1.DeleteMailboxRequest]) (*connect_go.Response[emptypb.Empty], error)
	UpdateMailbox(context.Context, *connect_go.Request[v1.UpdateMailboxRequest]) (*connect_go.Response[v1.UpdateMailboxResponse], error)
	ListVoiceMails(context.Context, *connect_go.Request[v1.ListVoiceMailsRequest]) (*connect_go.Response[v1.ListVoiceMailsResponse], error)
	GetVoiceMail(context.Context, *connect_go.Request[v1.GetVoiceMailRequest]) (*connect_go.Response[v1.GetVoiceMailResponse], error)
	MarkVoiceMails(context.Context, *connect_go.Request[v1.MarkVoiceMailsRequest]) (*connect_go.Response[v1.MarkVoiceMailsResponse], error)
	SearchVoiceMails(context.Context, *connect_go.Request[v1.SearchVoiceMailsRequest]) (*connect_go.Response[v1.SearchVoiceMailsResponse], error)
}

// NewVoiceMailServiceClient constructs a client for the tkd.pbx3cx.v1.VoiceMailService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewVoiceMailServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) VoiceMailServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &voiceMailServiceClient{
		createMailbox: connect_go.NewClient[v1.CreateMailboxRequest, v1.CreateMailboxResponse](
			httpClient,
			baseURL+VoiceMailServiceCreateMailboxProcedure,
			opts...,
		),
		listMailboxes: connect_go.NewClient[v1.ListMailboxesRequest, v1.ListMailboxesResponse](
			httpClient,
			baseURL+VoiceMailServiceListMailboxesProcedure,
			opts...,
		),
		deleteMailbox: connect_go.NewClient[v1.DeleteMailboxRequest, emptypb.Empty](
			httpClient,
			baseURL+VoiceMailServiceDeleteMailboxProcedure,
			opts...,
		),
		updateMailbox: connect_go.NewClient[v1.UpdateMailboxRequest, v1.UpdateMailboxResponse](
			httpClient,
			baseURL+VoiceMailServiceUpdateMailboxProcedure,
			opts...,
		),
		listVoiceMails: connect_go.NewClient[v1.ListVoiceMailsRequest, v1.ListVoiceMailsResponse](
			httpClient,
			baseURL+VoiceMailServiceListVoiceMailsProcedure,
			opts...,
		),
		getVoiceMail: connect_go.NewClient[v1.GetVoiceMailRequest, v1.GetVoiceMailResponse](
			httpClient,
			baseURL+VoiceMailServiceGetVoiceMailProcedure,
			opts...,
		),
		markVoiceMails: connect_go.NewClient[v1.MarkVoiceMailsRequest, v1.MarkVoiceMailsResponse](
			httpClient,
			baseURL+VoiceMailServiceMarkVoiceMailsProcedure,
			opts...,
		),
		searchVoiceMails: connect_go.NewClient[v1.SearchVoiceMailsRequest, v1.SearchVoiceMailsResponse](
			httpClient,
			baseURL+VoiceMailServiceSearchVoiceMailsProcedure,
			opts...,
		),
	}
}

// voiceMailServiceClient implements VoiceMailServiceClient.
type voiceMailServiceClient struct {
	createMailbox    *connect_go.Client[v1.CreateMailboxRequest, v1.CreateMailboxResponse]
	listMailboxes    *connect_go.Client[v1.ListMailboxesRequest, v1.ListMailboxesResponse]
	deleteMailbox    *connect_go.Client[v1.DeleteMailboxRequest, emptypb.Empty]
	updateMailbox    *connect_go.Client[v1.UpdateMailboxRequest, v1.UpdateMailboxResponse]
	listVoiceMails   *connect_go.Client[v1.ListVoiceMailsRequest, v1.ListVoiceMailsResponse]
	getVoiceMail     *connect_go.Client[v1.GetVoiceMailRequest, v1.GetVoiceMailResponse]
	markVoiceMails   *connect_go.Client[v1.MarkVoiceMailsRequest, v1.MarkVoiceMailsResponse]
	searchVoiceMails *connect_go.Client[v1.SearchVoiceMailsRequest, v1.SearchVoiceMailsResponse]
}

// CreateMailbox calls tkd.pbx3cx.v1.VoiceMailService.CreateMailbox.
func (c *voiceMailServiceClient) CreateMailbox(ctx context.Context, req *connect_go.Request[v1.CreateMailboxRequest]) (*connect_go.Response[v1.CreateMailboxResponse], error) {
	return c.createMailbox.CallUnary(ctx, req)
}

// ListMailboxes calls tkd.pbx3cx.v1.VoiceMailService.ListMailboxes.
func (c *voiceMailServiceClient) ListMailboxes(ctx context.Context, req *connect_go.Request[v1.ListMailboxesRequest]) (*connect_go.Response[v1.ListMailboxesResponse], error) {
	return c.listMailboxes.CallUnary(ctx, req)
}

// DeleteMailbox calls tkd.pbx3cx.v1.VoiceMailService.DeleteMailbox.
func (c *voiceMailServiceClient) DeleteMailbox(ctx context.Context, req *connect_go.Request[v1.DeleteMailboxRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.deleteMailbox.CallUnary(ctx, req)
}

// UpdateMailbox calls tkd.pbx3cx.v1.VoiceMailService.UpdateMailbox.
func (c *voiceMailServiceClient) UpdateMailbox(ctx context.Context, req *connect_go.Request[v1.UpdateMailboxRequest]) (*connect_go.Response[v1.UpdateMailboxResponse], error) {
	return c.updateMailbox.CallUnary(ctx, req)
}

// ListVoiceMails calls tkd.pbx3cx.v1.VoiceMailService.ListVoiceMails.
func (c *voiceMailServiceClient) ListVoiceMails(ctx context.Context, req *connect_go.Request[v1.ListVoiceMailsRequest]) (*connect_go.Response[v1.ListVoiceMailsResponse], error) {
	return c.listVoiceMails.CallUnary(ctx, req)
}

// GetVoiceMail calls tkd.pbx3cx.v1.VoiceMailService.GetVoiceMail.
func (c *voiceMailServiceClient) GetVoiceMail(ctx context.Context, req *connect_go.Request[v1.GetVoiceMailRequest]) (*connect_go.Response[v1.GetVoiceMailResponse], error) {
	return c.getVoiceMail.CallUnary(ctx, req)
}

// MarkVoiceMails calls tkd.pbx3cx.v1.VoiceMailService.MarkVoiceMails.
func (c *voiceMailServiceClient) MarkVoiceMails(ctx context.Context, req *connect_go.Request[v1.MarkVoiceMailsRequest]) (*connect_go.Response[v1.MarkVoiceMailsResponse], error) {
	return c.markVoiceMails.CallUnary(ctx, req)
}

// SearchVoiceMails calls tkd.pbx3cx.v1.VoiceMailService.SearchVoiceMails.
func (c *voiceMailServiceClient) SearchVoiceMails(ctx context.Context, req *connect_go.Request[v1.SearchVoiceMailsRequest]) (*connect_go.Response[v1.SearchVoiceMailsResponse], error) {
	return c.searchVoiceMails.CallUnary(ctx, req)
}

// VoiceMailServiceHandler is an implementation of the tkd.pbx3cx.v1.VoiceMailService service.
type VoiceMailServiceHandler interface {
	CreateMailbox(context.Context, *connect_go.Request[v1.CreateMailboxRequest]) (*connect_go.Response[v1.CreateMailboxResponse], error)
	ListMailboxes(context.Context, *connect_go.Request[v1.ListMailboxesRequest]) (*connect_go.Response[v1.ListMailboxesResponse], error)
	DeleteMailbox(context.Context, *connect_go.Request[v1.DeleteMailboxRequest]) (*connect_go.Response[emptypb.Empty], error)
	UpdateMailbox(context.Context, *connect_go.Request[v1.UpdateMailboxRequest]) (*connect_go.Response[v1.UpdateMailboxResponse], error)
	ListVoiceMails(context.Context, *connect_go.Request[v1.ListVoiceMailsRequest]) (*connect_go.Response[v1.ListVoiceMailsResponse], error)
	GetVoiceMail(context.Context, *connect_go.Request[v1.GetVoiceMailRequest]) (*connect_go.Response[v1.GetVoiceMailResponse], error)
	MarkVoiceMails(context.Context, *connect_go.Request[v1.MarkVoiceMailsRequest]) (*connect_go.Response[v1.MarkVoiceMailsResponse], error)
	SearchVoiceMails(context.Context, *connect_go.Request[v1.SearchVoiceMailsRequest]) (*connect_go.Response[v1.SearchVoiceMailsResponse], error)
}

// NewVoiceMailServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewVoiceMailServiceHandler(svc VoiceMailServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	voiceMailServiceCreateMailboxHandler := connect_go.NewUnaryHandler(
		VoiceMailServiceCreateMailboxProcedure,
		svc.CreateMailbox,
		opts...,
	)
	voiceMailServiceListMailboxesHandler := connect_go.NewUnaryHandler(
		VoiceMailServiceListMailboxesProcedure,
		svc.ListMailboxes,
		opts...,
	)
	voiceMailServiceDeleteMailboxHandler := connect_go.NewUnaryHandler(
		VoiceMailServiceDeleteMailboxProcedure,
		svc.DeleteMailbox,
		opts...,
	)
	voiceMailServiceUpdateMailboxHandler := connect_go.NewUnaryHandler(
		VoiceMailServiceUpdateMailboxProcedure,
		svc.UpdateMailbox,
		opts...,
	)
	voiceMailServiceListVoiceMailsHandler := connect_go.NewUnaryHandler(
		VoiceMailServiceListVoiceMailsProcedure,
		svc.ListVoiceMails,
		opts...,
	)
	voiceMailServiceGetVoiceMailHandler := connect_go.NewUnaryHandler(
		VoiceMailServiceGetVoiceMailProcedure,
		svc.GetVoiceMail,
		opts...,
	)
	voiceMailServiceMarkVoiceMailsHandler := connect_go.NewUnaryHandler(
		VoiceMailServiceMarkVoiceMailsProcedure,
		svc.MarkVoiceMails,
		opts...,
	)
	voiceMailServiceSearchVoiceMailsHandler := connect_go.NewUnaryHandler(
		VoiceMailServiceSearchVoiceMailsProcedure,
		svc.SearchVoiceMails,
		opts...,
	)
	return "/tkd.pbx3cx.v1.VoiceMailService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case VoiceMailServiceCreateMailboxProcedure:
			voiceMailServiceCreateMailboxHandler.ServeHTTP(w, r)
		case VoiceMailServiceListMailboxesProcedure:
			voiceMailServiceListMailboxesHandler.ServeHTTP(w, r)
		case VoiceMailServiceDeleteMailboxProcedure:
			voiceMailServiceDeleteMailboxHandler.ServeHTTP(w, r)
		case VoiceMailServiceUpdateMailboxProcedure:
			voiceMailServiceUpdateMailboxHandler.ServeHTTP(w, r)
		case VoiceMailServiceListVoiceMailsProcedure:
			voiceMailServiceListVoiceMailsHandler.ServeHTTP(w, r)
		case VoiceMailServiceGetVoiceMailProcedure:
			voiceMailServiceGetVoiceMailHandler.ServeHTTP(w, r)
		case VoiceMailServiceMarkVoiceMailsProcedure:
			voiceMailServiceMarkVoiceMailsHandler.ServeHTTP(w, r)
		case VoiceMailServiceSearchVoiceMailsProcedure:
			voiceMailServiceSearchVoiceMailsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedVoiceMailServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedVoiceMailServiceHandler struct{}

func (UnimplementedVoiceMailServiceHandler) CreateMailbox(context.Context, *connect_go.Request[v1.CreateMailboxRequest]) (*connect_go.Response[v1.CreateMailboxResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.pbx3cx.v1.VoiceMailService.CreateMailbox is not implemented"))
}

func (UnimplementedVoiceMailServiceHandler) ListMailboxes(context.Context, *connect_go.Request[v1.ListMailboxesRequest]) (*connect_go.Response[v1.ListMailboxesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.pbx3cx.v1.VoiceMailService.ListMailboxes is not implemented"))
}

func (UnimplementedVoiceMailServiceHandler) DeleteMailbox(context.Context, *connect_go.Request[v1.DeleteMailboxRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.pbx3cx.v1.VoiceMailService.DeleteMailbox is not implemented"))
}

func (UnimplementedVoiceMailServiceHandler) UpdateMailbox(context.Context, *connect_go.Request[v1.UpdateMailboxRequest]) (*connect_go.Response[v1.UpdateMailboxResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.pbx3cx.v1.VoiceMailService.UpdateMailbox is not implemented"))
}

func (UnimplementedVoiceMailServiceHandler) ListVoiceMails(context.Context, *connect_go.Request[v1.ListVoiceMailsRequest]) (*connect_go.Response[v1.ListVoiceMailsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.pbx3cx.v1.VoiceMailService.ListVoiceMails is not implemented"))
}

func (UnimplementedVoiceMailServiceHandler) GetVoiceMail(context.Context, *connect_go.Request[v1.GetVoiceMailRequest]) (*connect_go.Response[v1.GetVoiceMailResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.pbx3cx.v1.VoiceMailService.GetVoiceMail is not implemented"))
}

func (UnimplementedVoiceMailServiceHandler) MarkVoiceMails(context.Context, *connect_go.Request[v1.MarkVoiceMailsRequest]) (*connect_go.Response[v1.MarkVoiceMailsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.pbx3cx.v1.VoiceMailService.MarkVoiceMails is not implemented"))
}

func (UnimplementedVoiceMailServiceHandler) SearchVoiceMails(context.Context, *connect_go.Request[v1.SearchVoiceMailsRequest]) (*connect_go.Response[v1.SearchVoiceMailsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.pbx3cx.v1.VoiceMailService.SearchVoiceMails is not implemented"))
}
