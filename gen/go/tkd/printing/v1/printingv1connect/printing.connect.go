// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: tkd/printing/v1/printing.proto

package printingv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v11 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/longrunning/v1"
	v1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/printing/v1"
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
	// PrintServiceName is the fully-qualified name of the PrintService service.
	PrintServiceName = "tkd.printing.v1.PrintService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PrintServiceListPrintersProcedure is the fully-qualified name of the PrintService's ListPrinters
	// RPC.
	PrintServiceListPrintersProcedure = "/tkd.printing.v1.PrintService/ListPrinters"
	// PrintServicePrintDocumentProcedure is the fully-qualified name of the PrintService's
	// PrintDocument RPC.
	PrintServicePrintDocumentProcedure = "/tkd.printing.v1.PrintService/PrintDocument"
	// PrintServicePrintDocumentStreamProcedure is the fully-qualified name of the PrintService's
	// PrintDocumentStream RPC.
	PrintServicePrintDocumentStreamProcedure = "/tkd.printing.v1.PrintService/PrintDocumentStream"
)

// PrintServiceClient is a client for the tkd.printing.v1.PrintService service.
type PrintServiceClient interface {
	// ListPrinters lists all available printers.
	ListPrinters(context.Context, *connect_go.Request[v1.ListPrintersRequest]) (*connect_go.Response[v1.ListPrintersResponse], error)
	// PrintDocument prints a document and returns a tkd.longrunning.v1.Operation to track
	// printing progress.
	PrintDocument(context.Context, *connect_go.Request[v1.Document]) (*connect_go.Response[v11.Operation], error)
	PrintDocumentStream(context.Context) *connect_go.ClientStreamForClient[v1.PrintDocumentRequest, v11.Operation]
}

// NewPrintServiceClient constructs a client for the tkd.printing.v1.PrintService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPrintServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) PrintServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &printServiceClient{
		listPrinters: connect_go.NewClient[v1.ListPrintersRequest, v1.ListPrintersResponse](
			httpClient,
			baseURL+PrintServiceListPrintersProcedure,
			opts...,
		),
		printDocument: connect_go.NewClient[v1.Document, v11.Operation](
			httpClient,
			baseURL+PrintServicePrintDocumentProcedure,
			opts...,
		),
		printDocumentStream: connect_go.NewClient[v1.PrintDocumentRequest, v11.Operation](
			httpClient,
			baseURL+PrintServicePrintDocumentStreamProcedure,
			opts...,
		),
	}
}

// printServiceClient implements PrintServiceClient.
type printServiceClient struct {
	listPrinters        *connect_go.Client[v1.ListPrintersRequest, v1.ListPrintersResponse]
	printDocument       *connect_go.Client[v1.Document, v11.Operation]
	printDocumentStream *connect_go.Client[v1.PrintDocumentRequest, v11.Operation]
}

// ListPrinters calls tkd.printing.v1.PrintService.ListPrinters.
func (c *printServiceClient) ListPrinters(ctx context.Context, req *connect_go.Request[v1.ListPrintersRequest]) (*connect_go.Response[v1.ListPrintersResponse], error) {
	return c.listPrinters.CallUnary(ctx, req)
}

// PrintDocument calls tkd.printing.v1.PrintService.PrintDocument.
func (c *printServiceClient) PrintDocument(ctx context.Context, req *connect_go.Request[v1.Document]) (*connect_go.Response[v11.Operation], error) {
	return c.printDocument.CallUnary(ctx, req)
}

// PrintDocumentStream calls tkd.printing.v1.PrintService.PrintDocumentStream.
func (c *printServiceClient) PrintDocumentStream(ctx context.Context) *connect_go.ClientStreamForClient[v1.PrintDocumentRequest, v11.Operation] {
	return c.printDocumentStream.CallClientStream(ctx)
}

// PrintServiceHandler is an implementation of the tkd.printing.v1.PrintService service.
type PrintServiceHandler interface {
	// ListPrinters lists all available printers.
	ListPrinters(context.Context, *connect_go.Request[v1.ListPrintersRequest]) (*connect_go.Response[v1.ListPrintersResponse], error)
	// PrintDocument prints a document and returns a tkd.longrunning.v1.Operation to track
	// printing progress.
	PrintDocument(context.Context, *connect_go.Request[v1.Document]) (*connect_go.Response[v11.Operation], error)
	PrintDocumentStream(context.Context, *connect_go.ClientStream[v1.PrintDocumentRequest]) (*connect_go.Response[v11.Operation], error)
}

// NewPrintServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPrintServiceHandler(svc PrintServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	printServiceListPrintersHandler := connect_go.NewUnaryHandler(
		PrintServiceListPrintersProcedure,
		svc.ListPrinters,
		opts...,
	)
	printServicePrintDocumentHandler := connect_go.NewUnaryHandler(
		PrintServicePrintDocumentProcedure,
		svc.PrintDocument,
		opts...,
	)
	printServicePrintDocumentStreamHandler := connect_go.NewClientStreamHandler(
		PrintServicePrintDocumentStreamProcedure,
		svc.PrintDocumentStream,
		opts...,
	)
	return "/tkd.printing.v1.PrintService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PrintServiceListPrintersProcedure:
			printServiceListPrintersHandler.ServeHTTP(w, r)
		case PrintServicePrintDocumentProcedure:
			printServicePrintDocumentHandler.ServeHTTP(w, r)
		case PrintServicePrintDocumentStreamProcedure:
			printServicePrintDocumentStreamHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPrintServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPrintServiceHandler struct{}

func (UnimplementedPrintServiceHandler) ListPrinters(context.Context, *connect_go.Request[v1.ListPrintersRequest]) (*connect_go.Response[v1.ListPrintersResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.printing.v1.PrintService.ListPrinters is not implemented"))
}

func (UnimplementedPrintServiceHandler) PrintDocument(context.Context, *connect_go.Request[v1.Document]) (*connect_go.Response[v11.Operation], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.printing.v1.PrintService.PrintDocument is not implemented"))
}

func (UnimplementedPrintServiceHandler) PrintDocumentStream(context.Context, *connect_go.ClientStream[v1.PrintDocumentRequest]) (*connect_go.Response[v11.Operation], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.printing.v1.PrintService.PrintDocumentStream is not implemented"))
}
