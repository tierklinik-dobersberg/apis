// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: tkd/survey/survey.proto

package surveyv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	survey "github.com/tierklinik-dobersberg/apis/gen/go/tkd/survey"
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
	// SurveyServiceName is the fully-qualified name of the SurveyService service.
	SurveyServiceName = "tkd.survey.v1.SurveyService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SurveyServiceCreateSurveyProcedure is the fully-qualified name of the SurveyService's
	// CreateSurvey RPC.
	SurveyServiceCreateSurveyProcedure = "/tkd.survey.v1.SurveyService/CreateSurvey"
	// SurveyServiceDeleteSurveyProcedure is the fully-qualified name of the SurveyService's
	// DeleteSurvey RPC.
	SurveyServiceDeleteSurveyProcedure = "/tkd.survey.v1.SurveyService/DeleteSurvey"
	// SurveyServiceListSurveyDefinitionsProcedure is the fully-qualified name of the SurveyService's
	// ListSurveyDefinitions RPC.
	SurveyServiceListSurveyDefinitionsProcedure = "/tkd.survey.v1.SurveyService/ListSurveyDefinitions"
	// SurveyServiceSaveSurveyVoteProcedure is the fully-qualified name of the SurveyService's
	// SaveSurveyVote RPC.
	SurveyServiceSaveSurveyVoteProcedure = "/tkd.survey.v1.SurveyService/SaveSurveyVote"
)

// SurveyServiceClient is a client for the tkd.survey.v1.SurveyService service.
type SurveyServiceClient interface {
	CreateSurvey(context.Context, *connect_go.Request[survey.CreateSurveyRequest]) (*connect_go.Response[survey.Survey], error)
	DeleteSurvey(context.Context, *connect_go.Request[survey.DeleteSurveyRequest]) (*connect_go.Response[emptypb.Empty], error)
	ListSurveyDefinitions(context.Context, *connect_go.Request[survey.ListSurveyDefinitionsRequest]) (*connect_go.Response[survey.ListSurveyDefinitionsResponse], error)
	SaveSurveyVote(context.Context, *connect_go.Request[survey.SaveSurveyVoteRequest]) (*connect_go.Response[emptypb.Empty], error)
}

// NewSurveyServiceClient constructs a client for the tkd.survey.v1.SurveyService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSurveyServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) SurveyServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &surveyServiceClient{
		createSurvey: connect_go.NewClient[survey.CreateSurveyRequest, survey.Survey](
			httpClient,
			baseURL+SurveyServiceCreateSurveyProcedure,
			opts...,
		),
		deleteSurvey: connect_go.NewClient[survey.DeleteSurveyRequest, emptypb.Empty](
			httpClient,
			baseURL+SurveyServiceDeleteSurveyProcedure,
			opts...,
		),
		listSurveyDefinitions: connect_go.NewClient[survey.ListSurveyDefinitionsRequest, survey.ListSurveyDefinitionsResponse](
			httpClient,
			baseURL+SurveyServiceListSurveyDefinitionsProcedure,
			opts...,
		),
		saveSurveyVote: connect_go.NewClient[survey.SaveSurveyVoteRequest, emptypb.Empty](
			httpClient,
			baseURL+SurveyServiceSaveSurveyVoteProcedure,
			opts...,
		),
	}
}

// surveyServiceClient implements SurveyServiceClient.
type surveyServiceClient struct {
	createSurvey          *connect_go.Client[survey.CreateSurveyRequest, survey.Survey]
	deleteSurvey          *connect_go.Client[survey.DeleteSurveyRequest, emptypb.Empty]
	listSurveyDefinitions *connect_go.Client[survey.ListSurveyDefinitionsRequest, survey.ListSurveyDefinitionsResponse]
	saveSurveyVote        *connect_go.Client[survey.SaveSurveyVoteRequest, emptypb.Empty]
}

// CreateSurvey calls tkd.survey.v1.SurveyService.CreateSurvey.
func (c *surveyServiceClient) CreateSurvey(ctx context.Context, req *connect_go.Request[survey.CreateSurveyRequest]) (*connect_go.Response[survey.Survey], error) {
	return c.createSurvey.CallUnary(ctx, req)
}

// DeleteSurvey calls tkd.survey.v1.SurveyService.DeleteSurvey.
func (c *surveyServiceClient) DeleteSurvey(ctx context.Context, req *connect_go.Request[survey.DeleteSurveyRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.deleteSurvey.CallUnary(ctx, req)
}

// ListSurveyDefinitions calls tkd.survey.v1.SurveyService.ListSurveyDefinitions.
func (c *surveyServiceClient) ListSurveyDefinitions(ctx context.Context, req *connect_go.Request[survey.ListSurveyDefinitionsRequest]) (*connect_go.Response[survey.ListSurveyDefinitionsResponse], error) {
	return c.listSurveyDefinitions.CallUnary(ctx, req)
}

// SaveSurveyVote calls tkd.survey.v1.SurveyService.SaveSurveyVote.
func (c *surveyServiceClient) SaveSurveyVote(ctx context.Context, req *connect_go.Request[survey.SaveSurveyVoteRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.saveSurveyVote.CallUnary(ctx, req)
}

// SurveyServiceHandler is an implementation of the tkd.survey.v1.SurveyService service.
type SurveyServiceHandler interface {
	CreateSurvey(context.Context, *connect_go.Request[survey.CreateSurveyRequest]) (*connect_go.Response[survey.Survey], error)
	DeleteSurvey(context.Context, *connect_go.Request[survey.DeleteSurveyRequest]) (*connect_go.Response[emptypb.Empty], error)
	ListSurveyDefinitions(context.Context, *connect_go.Request[survey.ListSurveyDefinitionsRequest]) (*connect_go.Response[survey.ListSurveyDefinitionsResponse], error)
	SaveSurveyVote(context.Context, *connect_go.Request[survey.SaveSurveyVoteRequest]) (*connect_go.Response[emptypb.Empty], error)
}

// NewSurveyServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSurveyServiceHandler(svc SurveyServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	surveyServiceCreateSurveyHandler := connect_go.NewUnaryHandler(
		SurveyServiceCreateSurveyProcedure,
		svc.CreateSurvey,
		opts...,
	)
	surveyServiceDeleteSurveyHandler := connect_go.NewUnaryHandler(
		SurveyServiceDeleteSurveyProcedure,
		svc.DeleteSurvey,
		opts...,
	)
	surveyServiceListSurveyDefinitionsHandler := connect_go.NewUnaryHandler(
		SurveyServiceListSurveyDefinitionsProcedure,
		svc.ListSurveyDefinitions,
		opts...,
	)
	surveyServiceSaveSurveyVoteHandler := connect_go.NewUnaryHandler(
		SurveyServiceSaveSurveyVoteProcedure,
		svc.SaveSurveyVote,
		opts...,
	)
	return "/tkd.survey.v1.SurveyService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SurveyServiceCreateSurveyProcedure:
			surveyServiceCreateSurveyHandler.ServeHTTP(w, r)
		case SurveyServiceDeleteSurveyProcedure:
			surveyServiceDeleteSurveyHandler.ServeHTTP(w, r)
		case SurveyServiceListSurveyDefinitionsProcedure:
			surveyServiceListSurveyDefinitionsHandler.ServeHTTP(w, r)
		case SurveyServiceSaveSurveyVoteProcedure:
			surveyServiceSaveSurveyVoteHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSurveyServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSurveyServiceHandler struct{}

func (UnimplementedSurveyServiceHandler) CreateSurvey(context.Context, *connect_go.Request[survey.CreateSurveyRequest]) (*connect_go.Response[survey.Survey], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.survey.v1.SurveyService.CreateSurvey is not implemented"))
}

func (UnimplementedSurveyServiceHandler) DeleteSurvey(context.Context, *connect_go.Request[survey.DeleteSurveyRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.survey.v1.SurveyService.DeleteSurvey is not implemented"))
}

func (UnimplementedSurveyServiceHandler) ListSurveyDefinitions(context.Context, *connect_go.Request[survey.ListSurveyDefinitionsRequest]) (*connect_go.Response[survey.ListSurveyDefinitionsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.survey.v1.SurveyService.ListSurveyDefinitions is not implemented"))
}

func (UnimplementedSurveyServiceHandler) SaveSurveyVote(context.Context, *connect_go.Request[survey.SaveSurveyVoteRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.survey.v1.SurveyService.SaveSurveyVote is not implemented"))
}
