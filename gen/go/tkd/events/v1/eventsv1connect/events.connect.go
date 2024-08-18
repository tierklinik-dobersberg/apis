// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: tkd/events/v1/events.proto

package eventsv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1"
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
	// EventServiceName is the fully-qualified name of the EventService service.
	EventServiceName = "tkd.events.v1.EventService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// EventServiceSubscribeProcedure is the fully-qualified name of the EventService's Subscribe RPC.
	EventServiceSubscribeProcedure = "/tkd.events.v1.EventService/Subscribe"
	// EventServicePublishProcedure is the fully-qualified name of the EventService's Publish RPC.
	EventServicePublishProcedure = "/tkd.events.v1.EventService/Publish"
	// EventServicePublishStreamProcedure is the fully-qualified name of the EventService's
	// PublishStream RPC.
	EventServicePublishStreamProcedure = "/tkd.events.v1.EventService/PublishStream"
)

// EventServiceClient is a client for the tkd.events.v1.EventService service.
type EventServiceClient interface {
	Subscribe(context.Context) *connect_go.BidiStreamForClient[v1.SubscribeRequest, v1.Event]
	Publish(context.Context, *connect_go.Request[v1.Event]) (*connect_go.Response[emptypb.Empty], error)
	PublishStream(context.Context) *connect_go.ClientStreamForClient[v1.Event, emptypb.Empty]
}

// NewEventServiceClient constructs a client for the tkd.events.v1.EventService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewEventServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) EventServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &eventServiceClient{
		subscribe: connect_go.NewClient[v1.SubscribeRequest, v1.Event](
			httpClient,
			baseURL+EventServiceSubscribeProcedure,
			opts...,
		),
		publish: connect_go.NewClient[v1.Event, emptypb.Empty](
			httpClient,
			baseURL+EventServicePublishProcedure,
			opts...,
		),
		publishStream: connect_go.NewClient[v1.Event, emptypb.Empty](
			httpClient,
			baseURL+EventServicePublishStreamProcedure,
			opts...,
		),
	}
}

// eventServiceClient implements EventServiceClient.
type eventServiceClient struct {
	subscribe     *connect_go.Client[v1.SubscribeRequest, v1.Event]
	publish       *connect_go.Client[v1.Event, emptypb.Empty]
	publishStream *connect_go.Client[v1.Event, emptypb.Empty]
}

// Subscribe calls tkd.events.v1.EventService.Subscribe.
func (c *eventServiceClient) Subscribe(ctx context.Context) *connect_go.BidiStreamForClient[v1.SubscribeRequest, v1.Event] {
	return c.subscribe.CallBidiStream(ctx)
}

// Publish calls tkd.events.v1.EventService.Publish.
func (c *eventServiceClient) Publish(ctx context.Context, req *connect_go.Request[v1.Event]) (*connect_go.Response[emptypb.Empty], error) {
	return c.publish.CallUnary(ctx, req)
}

// PublishStream calls tkd.events.v1.EventService.PublishStream.
func (c *eventServiceClient) PublishStream(ctx context.Context) *connect_go.ClientStreamForClient[v1.Event, emptypb.Empty] {
	return c.publishStream.CallClientStream(ctx)
}

// EventServiceHandler is an implementation of the tkd.events.v1.EventService service.
type EventServiceHandler interface {
	Subscribe(context.Context, *connect_go.BidiStream[v1.SubscribeRequest, v1.Event]) error
	Publish(context.Context, *connect_go.Request[v1.Event]) (*connect_go.Response[emptypb.Empty], error)
	PublishStream(context.Context, *connect_go.ClientStream[v1.Event]) (*connect_go.Response[emptypb.Empty], error)
}

// NewEventServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewEventServiceHandler(svc EventServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	eventServiceSubscribeHandler := connect_go.NewBidiStreamHandler(
		EventServiceSubscribeProcedure,
		svc.Subscribe,
		opts...,
	)
	eventServicePublishHandler := connect_go.NewUnaryHandler(
		EventServicePublishProcedure,
		svc.Publish,
		opts...,
	)
	eventServicePublishStreamHandler := connect_go.NewClientStreamHandler(
		EventServicePublishStreamProcedure,
		svc.PublishStream,
		opts...,
	)
	return "/tkd.events.v1.EventService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case EventServiceSubscribeProcedure:
			eventServiceSubscribeHandler.ServeHTTP(w, r)
		case EventServicePublishProcedure:
			eventServicePublishHandler.ServeHTTP(w, r)
		case EventServicePublishStreamProcedure:
			eventServicePublishStreamHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedEventServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedEventServiceHandler struct{}

func (UnimplementedEventServiceHandler) Subscribe(context.Context, *connect_go.BidiStream[v1.SubscribeRequest, v1.Event]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.events.v1.EventService.Subscribe is not implemented"))
}

func (UnimplementedEventServiceHandler) Publish(context.Context, *connect_go.Request[v1.Event]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.events.v1.EventService.Publish is not implemented"))
}

func (UnimplementedEventServiceHandler) PublishStream(context.Context, *connect_go.ClientStream[v1.Event]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.events.v1.EventService.PublishStream is not implemented"))
}
