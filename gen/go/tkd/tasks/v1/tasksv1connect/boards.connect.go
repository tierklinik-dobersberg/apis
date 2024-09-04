// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: tkd/tasks/v1/boards.proto

package tasksv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/tasks/v1"
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
	// BoardServiceName is the fully-qualified name of the BoardService service.
	BoardServiceName = "tkd.tasks.v1.BoardService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// BoardServiceCreateBoardProcedure is the fully-qualified name of the BoardService's CreateBoard
	// RPC.
	BoardServiceCreateBoardProcedure = "/tkd.tasks.v1.BoardService/CreateBoard"
	// BoardServiceListBoardsProcedure is the fully-qualified name of the BoardService's ListBoards RPC.
	BoardServiceListBoardsProcedure = "/tkd.tasks.v1.BoardService/ListBoards"
	// BoardServiceDeleteBoardProcedure is the fully-qualified name of the BoardService's DeleteBoard
	// RPC.
	BoardServiceDeleteBoardProcedure = "/tkd.tasks.v1.BoardService/DeleteBoard"
	// BoardServiceGetBoardProcedure is the fully-qualified name of the BoardService's GetBoard RPC.
	BoardServiceGetBoardProcedure = "/tkd.tasks.v1.BoardService/GetBoard"
	// BoardServiceSaveNotificationProcedure is the fully-qualified name of the BoardService's
	// SaveNotification RPC.
	BoardServiceSaveNotificationProcedure = "/tkd.tasks.v1.BoardService/SaveNotification"
	// BoardServiceDeleteNotificationProcedure is the fully-qualified name of the BoardService's
	// DeleteNotification RPC.
	BoardServiceDeleteNotificationProcedure = "/tkd.tasks.v1.BoardService/DeleteNotification"
)

// BoardServiceClient is a client for the tkd.tasks.v1.BoardService service.
type BoardServiceClient interface {
	CreateBoard(context.Context, *connect_go.Request[v1.CreateBoardRequest]) (*connect_go.Response[v1.CreateBoardResponse], error)
	ListBoards(context.Context, *connect_go.Request[v1.ListBoardsRequest]) (*connect_go.Response[v1.ListBoardsResponse], error)
	DeleteBoard(context.Context, *connect_go.Request[v1.DeleteBoardRequest]) (*connect_go.Response[emptypb.Empty], error)
	GetBoard(context.Context, *connect_go.Request[v1.GetBoardRequest]) (*connect_go.Response[v1.GetBoardResponse], error)
	SaveNotification(context.Context, *connect_go.Request[v1.SaveNotificationRequest]) (*connect_go.Response[v1.SaveNotificationResponse], error)
	DeleteNotification(context.Context, *connect_go.Request[v1.DeleteNotificationRequest]) (*connect_go.Response[v1.DeleteNotificationResponse], error)
}

// NewBoardServiceClient constructs a client for the tkd.tasks.v1.BoardService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewBoardServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) BoardServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &boardServiceClient{
		createBoard: connect_go.NewClient[v1.CreateBoardRequest, v1.CreateBoardResponse](
			httpClient,
			baseURL+BoardServiceCreateBoardProcedure,
			opts...,
		),
		listBoards: connect_go.NewClient[v1.ListBoardsRequest, v1.ListBoardsResponse](
			httpClient,
			baseURL+BoardServiceListBoardsProcedure,
			opts...,
		),
		deleteBoard: connect_go.NewClient[v1.DeleteBoardRequest, emptypb.Empty](
			httpClient,
			baseURL+BoardServiceDeleteBoardProcedure,
			opts...,
		),
		getBoard: connect_go.NewClient[v1.GetBoardRequest, v1.GetBoardResponse](
			httpClient,
			baseURL+BoardServiceGetBoardProcedure,
			opts...,
		),
		saveNotification: connect_go.NewClient[v1.SaveNotificationRequest, v1.SaveNotificationResponse](
			httpClient,
			baseURL+BoardServiceSaveNotificationProcedure,
			opts...,
		),
		deleteNotification: connect_go.NewClient[v1.DeleteNotificationRequest, v1.DeleteNotificationResponse](
			httpClient,
			baseURL+BoardServiceDeleteNotificationProcedure,
			opts...,
		),
	}
}

// boardServiceClient implements BoardServiceClient.
type boardServiceClient struct {
	createBoard        *connect_go.Client[v1.CreateBoardRequest, v1.CreateBoardResponse]
	listBoards         *connect_go.Client[v1.ListBoardsRequest, v1.ListBoardsResponse]
	deleteBoard        *connect_go.Client[v1.DeleteBoardRequest, emptypb.Empty]
	getBoard           *connect_go.Client[v1.GetBoardRequest, v1.GetBoardResponse]
	saveNotification   *connect_go.Client[v1.SaveNotificationRequest, v1.SaveNotificationResponse]
	deleteNotification *connect_go.Client[v1.DeleteNotificationRequest, v1.DeleteNotificationResponse]
}

// CreateBoard calls tkd.tasks.v1.BoardService.CreateBoard.
func (c *boardServiceClient) CreateBoard(ctx context.Context, req *connect_go.Request[v1.CreateBoardRequest]) (*connect_go.Response[v1.CreateBoardResponse], error) {
	return c.createBoard.CallUnary(ctx, req)
}

// ListBoards calls tkd.tasks.v1.BoardService.ListBoards.
func (c *boardServiceClient) ListBoards(ctx context.Context, req *connect_go.Request[v1.ListBoardsRequest]) (*connect_go.Response[v1.ListBoardsResponse], error) {
	return c.listBoards.CallUnary(ctx, req)
}

// DeleteBoard calls tkd.tasks.v1.BoardService.DeleteBoard.
func (c *boardServiceClient) DeleteBoard(ctx context.Context, req *connect_go.Request[v1.DeleteBoardRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.deleteBoard.CallUnary(ctx, req)
}

// GetBoard calls tkd.tasks.v1.BoardService.GetBoard.
func (c *boardServiceClient) GetBoard(ctx context.Context, req *connect_go.Request[v1.GetBoardRequest]) (*connect_go.Response[v1.GetBoardResponse], error) {
	return c.getBoard.CallUnary(ctx, req)
}

// SaveNotification calls tkd.tasks.v1.BoardService.SaveNotification.
func (c *boardServiceClient) SaveNotification(ctx context.Context, req *connect_go.Request[v1.SaveNotificationRequest]) (*connect_go.Response[v1.SaveNotificationResponse], error) {
	return c.saveNotification.CallUnary(ctx, req)
}

// DeleteNotification calls tkd.tasks.v1.BoardService.DeleteNotification.
func (c *boardServiceClient) DeleteNotification(ctx context.Context, req *connect_go.Request[v1.DeleteNotificationRequest]) (*connect_go.Response[v1.DeleteNotificationResponse], error) {
	return c.deleteNotification.CallUnary(ctx, req)
}

// BoardServiceHandler is an implementation of the tkd.tasks.v1.BoardService service.
type BoardServiceHandler interface {
	CreateBoard(context.Context, *connect_go.Request[v1.CreateBoardRequest]) (*connect_go.Response[v1.CreateBoardResponse], error)
	ListBoards(context.Context, *connect_go.Request[v1.ListBoardsRequest]) (*connect_go.Response[v1.ListBoardsResponse], error)
	DeleteBoard(context.Context, *connect_go.Request[v1.DeleteBoardRequest]) (*connect_go.Response[emptypb.Empty], error)
	GetBoard(context.Context, *connect_go.Request[v1.GetBoardRequest]) (*connect_go.Response[v1.GetBoardResponse], error)
	SaveNotification(context.Context, *connect_go.Request[v1.SaveNotificationRequest]) (*connect_go.Response[v1.SaveNotificationResponse], error)
	DeleteNotification(context.Context, *connect_go.Request[v1.DeleteNotificationRequest]) (*connect_go.Response[v1.DeleteNotificationResponse], error)
}

// NewBoardServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewBoardServiceHandler(svc BoardServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	boardServiceCreateBoardHandler := connect_go.NewUnaryHandler(
		BoardServiceCreateBoardProcedure,
		svc.CreateBoard,
		opts...,
	)
	boardServiceListBoardsHandler := connect_go.NewUnaryHandler(
		BoardServiceListBoardsProcedure,
		svc.ListBoards,
		opts...,
	)
	boardServiceDeleteBoardHandler := connect_go.NewUnaryHandler(
		BoardServiceDeleteBoardProcedure,
		svc.DeleteBoard,
		opts...,
	)
	boardServiceGetBoardHandler := connect_go.NewUnaryHandler(
		BoardServiceGetBoardProcedure,
		svc.GetBoard,
		opts...,
	)
	boardServiceSaveNotificationHandler := connect_go.NewUnaryHandler(
		BoardServiceSaveNotificationProcedure,
		svc.SaveNotification,
		opts...,
	)
	boardServiceDeleteNotificationHandler := connect_go.NewUnaryHandler(
		BoardServiceDeleteNotificationProcedure,
		svc.DeleteNotification,
		opts...,
	)
	return "/tkd.tasks.v1.BoardService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case BoardServiceCreateBoardProcedure:
			boardServiceCreateBoardHandler.ServeHTTP(w, r)
		case BoardServiceListBoardsProcedure:
			boardServiceListBoardsHandler.ServeHTTP(w, r)
		case BoardServiceDeleteBoardProcedure:
			boardServiceDeleteBoardHandler.ServeHTTP(w, r)
		case BoardServiceGetBoardProcedure:
			boardServiceGetBoardHandler.ServeHTTP(w, r)
		case BoardServiceSaveNotificationProcedure:
			boardServiceSaveNotificationHandler.ServeHTTP(w, r)
		case BoardServiceDeleteNotificationProcedure:
			boardServiceDeleteNotificationHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedBoardServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedBoardServiceHandler struct{}

func (UnimplementedBoardServiceHandler) CreateBoard(context.Context, *connect_go.Request[v1.CreateBoardRequest]) (*connect_go.Response[v1.CreateBoardResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.tasks.v1.BoardService.CreateBoard is not implemented"))
}

func (UnimplementedBoardServiceHandler) ListBoards(context.Context, *connect_go.Request[v1.ListBoardsRequest]) (*connect_go.Response[v1.ListBoardsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.tasks.v1.BoardService.ListBoards is not implemented"))
}

func (UnimplementedBoardServiceHandler) DeleteBoard(context.Context, *connect_go.Request[v1.DeleteBoardRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.tasks.v1.BoardService.DeleteBoard is not implemented"))
}

func (UnimplementedBoardServiceHandler) GetBoard(context.Context, *connect_go.Request[v1.GetBoardRequest]) (*connect_go.Response[v1.GetBoardResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.tasks.v1.BoardService.GetBoard is not implemented"))
}

func (UnimplementedBoardServiceHandler) SaveNotification(context.Context, *connect_go.Request[v1.SaveNotificationRequest]) (*connect_go.Response[v1.SaveNotificationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.tasks.v1.BoardService.SaveNotification is not implemented"))
}

func (UnimplementedBoardServiceHandler) DeleteNotification(context.Context, *connect_go.Request[v1.DeleteNotificationRequest]) (*connect_go.Response[v1.DeleteNotificationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.tasks.v1.BoardService.DeleteNotification is not implemented"))
}
