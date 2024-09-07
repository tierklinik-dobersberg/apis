// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: tkd/tasks/v1/tasks.proto

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
	// TaskServiceName is the fully-qualified name of the TaskService service.
	TaskServiceName = "tkd.tasks.v1.TaskService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TaskServiceCreateTaskProcedure is the fully-qualified name of the TaskService's CreateTask RPC.
	TaskServiceCreateTaskProcedure = "/tkd.tasks.v1.TaskService/CreateTask"
	// TaskServiceUpdateTaskProcedure is the fully-qualified name of the TaskService's UpdateTask RPC.
	TaskServiceUpdateTaskProcedure = "/tkd.tasks.v1.TaskService/UpdateTask"
	// TaskServiceAssignTaskProcedure is the fully-qualified name of the TaskService's AssignTask RPC.
	TaskServiceAssignTaskProcedure = "/tkd.tasks.v1.TaskService/AssignTask"
	// TaskServiceCompleteTaskProcedure is the fully-qualified name of the TaskService's CompleteTask
	// RPC.
	TaskServiceCompleteTaskProcedure = "/tkd.tasks.v1.TaskService/CompleteTask"
	// TaskServiceDeleteTaskProcedure is the fully-qualified name of the TaskService's DeleteTask RPC.
	TaskServiceDeleteTaskProcedure = "/tkd.tasks.v1.TaskService/DeleteTask"
	// TaskServiceListTasksProcedure is the fully-qualified name of the TaskService's ListTasks RPC.
	TaskServiceListTasksProcedure = "/tkd.tasks.v1.TaskService/ListTasks"
	// TaskServiceGetTaskProcedure is the fully-qualified name of the TaskService's GetTask RPC.
	TaskServiceGetTaskProcedure = "/tkd.tasks.v1.TaskService/GetTask"
	// TaskServiceAddTaskAttachmentProcedure is the fully-qualified name of the TaskService's
	// AddTaskAttachment RPC.
	TaskServiceAddTaskAttachmentProcedure = "/tkd.tasks.v1.TaskService/AddTaskAttachment"
	// TaskServiceDeleteTaskAttachmentProcedure is the fully-qualified name of the TaskService's
	// DeleteTaskAttachment RPC.
	TaskServiceDeleteTaskAttachmentProcedure = "/tkd.tasks.v1.TaskService/DeleteTaskAttachment"
)

// TaskServiceClient is a client for the tkd.tasks.v1.TaskService service.
type TaskServiceClient interface {
	CreateTask(context.Context, *connect_go.Request[v1.CreateTaskRequest]) (*connect_go.Response[v1.CreateTaskResponse], error)
	UpdateTask(context.Context, *connect_go.Request[v1.UpdateTaskRequest]) (*connect_go.Response[v1.UpdateTaskResponse], error)
	AssignTask(context.Context, *connect_go.Request[v1.AssignTaskRequest]) (*connect_go.Response[v1.AssignTaskResponse], error)
	CompleteTask(context.Context, *connect_go.Request[v1.CompleteTaskRequest]) (*connect_go.Response[v1.CompleteTaskResponse], error)
	DeleteTask(context.Context, *connect_go.Request[v1.DeleteTaskRequest]) (*connect_go.Response[emptypb.Empty], error)
	ListTasks(context.Context, *connect_go.Request[v1.ListTasksRequest]) (*connect_go.Response[v1.ListTasksResponse], error)
	GetTask(context.Context, *connect_go.Request[v1.GetTaskRequest]) (*connect_go.Response[v1.GetTaskResponse], error)
	AddTaskAttachment(context.Context, *connect_go.Request[v1.AddTaskAttachmentRequest]) (*connect_go.Response[v1.AddTaskAttachmentResponse], error)
	DeleteTaskAttachment(context.Context, *connect_go.Request[v1.DeleteTaskAttachmentRequest]) (*connect_go.Response[v1.DeleteTaskAttachmentResponse], error)
}

// NewTaskServiceClient constructs a client for the tkd.tasks.v1.TaskService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTaskServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) TaskServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &taskServiceClient{
		createTask: connect_go.NewClient[v1.CreateTaskRequest, v1.CreateTaskResponse](
			httpClient,
			baseURL+TaskServiceCreateTaskProcedure,
			opts...,
		),
		updateTask: connect_go.NewClient[v1.UpdateTaskRequest, v1.UpdateTaskResponse](
			httpClient,
			baseURL+TaskServiceUpdateTaskProcedure,
			opts...,
		),
		assignTask: connect_go.NewClient[v1.AssignTaskRequest, v1.AssignTaskResponse](
			httpClient,
			baseURL+TaskServiceAssignTaskProcedure,
			opts...,
		),
		completeTask: connect_go.NewClient[v1.CompleteTaskRequest, v1.CompleteTaskResponse](
			httpClient,
			baseURL+TaskServiceCompleteTaskProcedure,
			opts...,
		),
		deleteTask: connect_go.NewClient[v1.DeleteTaskRequest, emptypb.Empty](
			httpClient,
			baseURL+TaskServiceDeleteTaskProcedure,
			opts...,
		),
		listTasks: connect_go.NewClient[v1.ListTasksRequest, v1.ListTasksResponse](
			httpClient,
			baseURL+TaskServiceListTasksProcedure,
			opts...,
		),
		getTask: connect_go.NewClient[v1.GetTaskRequest, v1.GetTaskResponse](
			httpClient,
			baseURL+TaskServiceGetTaskProcedure,
			opts...,
		),
		addTaskAttachment: connect_go.NewClient[v1.AddTaskAttachmentRequest, v1.AddTaskAttachmentResponse](
			httpClient,
			baseURL+TaskServiceAddTaskAttachmentProcedure,
			opts...,
		),
		deleteTaskAttachment: connect_go.NewClient[v1.DeleteTaskAttachmentRequest, v1.DeleteTaskAttachmentResponse](
			httpClient,
			baseURL+TaskServiceDeleteTaskAttachmentProcedure,
			opts...,
		),
	}
}

// taskServiceClient implements TaskServiceClient.
type taskServiceClient struct {
	createTask           *connect_go.Client[v1.CreateTaskRequest, v1.CreateTaskResponse]
	updateTask           *connect_go.Client[v1.UpdateTaskRequest, v1.UpdateTaskResponse]
	assignTask           *connect_go.Client[v1.AssignTaskRequest, v1.AssignTaskResponse]
	completeTask         *connect_go.Client[v1.CompleteTaskRequest, v1.CompleteTaskResponse]
	deleteTask           *connect_go.Client[v1.DeleteTaskRequest, emptypb.Empty]
	listTasks            *connect_go.Client[v1.ListTasksRequest, v1.ListTasksResponse]
	getTask              *connect_go.Client[v1.GetTaskRequest, v1.GetTaskResponse]
	addTaskAttachment    *connect_go.Client[v1.AddTaskAttachmentRequest, v1.AddTaskAttachmentResponse]
	deleteTaskAttachment *connect_go.Client[v1.DeleteTaskAttachmentRequest, v1.DeleteTaskAttachmentResponse]
}

// CreateTask calls tkd.tasks.v1.TaskService.CreateTask.
func (c *taskServiceClient) CreateTask(ctx context.Context, req *connect_go.Request[v1.CreateTaskRequest]) (*connect_go.Response[v1.CreateTaskResponse], error) {
	return c.createTask.CallUnary(ctx, req)
}

// UpdateTask calls tkd.tasks.v1.TaskService.UpdateTask.
func (c *taskServiceClient) UpdateTask(ctx context.Context, req *connect_go.Request[v1.UpdateTaskRequest]) (*connect_go.Response[v1.UpdateTaskResponse], error) {
	return c.updateTask.CallUnary(ctx, req)
}

// AssignTask calls tkd.tasks.v1.TaskService.AssignTask.
func (c *taskServiceClient) AssignTask(ctx context.Context, req *connect_go.Request[v1.AssignTaskRequest]) (*connect_go.Response[v1.AssignTaskResponse], error) {
	return c.assignTask.CallUnary(ctx, req)
}

// CompleteTask calls tkd.tasks.v1.TaskService.CompleteTask.
func (c *taskServiceClient) CompleteTask(ctx context.Context, req *connect_go.Request[v1.CompleteTaskRequest]) (*connect_go.Response[v1.CompleteTaskResponse], error) {
	return c.completeTask.CallUnary(ctx, req)
}

// DeleteTask calls tkd.tasks.v1.TaskService.DeleteTask.
func (c *taskServiceClient) DeleteTask(ctx context.Context, req *connect_go.Request[v1.DeleteTaskRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.deleteTask.CallUnary(ctx, req)
}

// ListTasks calls tkd.tasks.v1.TaskService.ListTasks.
func (c *taskServiceClient) ListTasks(ctx context.Context, req *connect_go.Request[v1.ListTasksRequest]) (*connect_go.Response[v1.ListTasksResponse], error) {
	return c.listTasks.CallUnary(ctx, req)
}

// GetTask calls tkd.tasks.v1.TaskService.GetTask.
func (c *taskServiceClient) GetTask(ctx context.Context, req *connect_go.Request[v1.GetTaskRequest]) (*connect_go.Response[v1.GetTaskResponse], error) {
	return c.getTask.CallUnary(ctx, req)
}

// AddTaskAttachment calls tkd.tasks.v1.TaskService.AddTaskAttachment.
func (c *taskServiceClient) AddTaskAttachment(ctx context.Context, req *connect_go.Request[v1.AddTaskAttachmentRequest]) (*connect_go.Response[v1.AddTaskAttachmentResponse], error) {
	return c.addTaskAttachment.CallUnary(ctx, req)
}

// DeleteTaskAttachment calls tkd.tasks.v1.TaskService.DeleteTaskAttachment.
func (c *taskServiceClient) DeleteTaskAttachment(ctx context.Context, req *connect_go.Request[v1.DeleteTaskAttachmentRequest]) (*connect_go.Response[v1.DeleteTaskAttachmentResponse], error) {
	return c.deleteTaskAttachment.CallUnary(ctx, req)
}

// TaskServiceHandler is an implementation of the tkd.tasks.v1.TaskService service.
type TaskServiceHandler interface {
	CreateTask(context.Context, *connect_go.Request[v1.CreateTaskRequest]) (*connect_go.Response[v1.CreateTaskResponse], error)
	UpdateTask(context.Context, *connect_go.Request[v1.UpdateTaskRequest]) (*connect_go.Response[v1.UpdateTaskResponse], error)
	AssignTask(context.Context, *connect_go.Request[v1.AssignTaskRequest]) (*connect_go.Response[v1.AssignTaskResponse], error)
	CompleteTask(context.Context, *connect_go.Request[v1.CompleteTaskRequest]) (*connect_go.Response[v1.CompleteTaskResponse], error)
	DeleteTask(context.Context, *connect_go.Request[v1.DeleteTaskRequest]) (*connect_go.Response[emptypb.Empty], error)
	ListTasks(context.Context, *connect_go.Request[v1.ListTasksRequest]) (*connect_go.Response[v1.ListTasksResponse], error)
	GetTask(context.Context, *connect_go.Request[v1.GetTaskRequest]) (*connect_go.Response[v1.GetTaskResponse], error)
	AddTaskAttachment(context.Context, *connect_go.Request[v1.AddTaskAttachmentRequest]) (*connect_go.Response[v1.AddTaskAttachmentResponse], error)
	DeleteTaskAttachment(context.Context, *connect_go.Request[v1.DeleteTaskAttachmentRequest]) (*connect_go.Response[v1.DeleteTaskAttachmentResponse], error)
}

// NewTaskServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTaskServiceHandler(svc TaskServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	taskServiceCreateTaskHandler := connect_go.NewUnaryHandler(
		TaskServiceCreateTaskProcedure,
		svc.CreateTask,
		opts...,
	)
	taskServiceUpdateTaskHandler := connect_go.NewUnaryHandler(
		TaskServiceUpdateTaskProcedure,
		svc.UpdateTask,
		opts...,
	)
	taskServiceAssignTaskHandler := connect_go.NewUnaryHandler(
		TaskServiceAssignTaskProcedure,
		svc.AssignTask,
		opts...,
	)
	taskServiceCompleteTaskHandler := connect_go.NewUnaryHandler(
		TaskServiceCompleteTaskProcedure,
		svc.CompleteTask,
		opts...,
	)
	taskServiceDeleteTaskHandler := connect_go.NewUnaryHandler(
		TaskServiceDeleteTaskProcedure,
		svc.DeleteTask,
		opts...,
	)
	taskServiceListTasksHandler := connect_go.NewUnaryHandler(
		TaskServiceListTasksProcedure,
		svc.ListTasks,
		opts...,
	)
	taskServiceGetTaskHandler := connect_go.NewUnaryHandler(
		TaskServiceGetTaskProcedure,
		svc.GetTask,
		opts...,
	)
	taskServiceAddTaskAttachmentHandler := connect_go.NewUnaryHandler(
		TaskServiceAddTaskAttachmentProcedure,
		svc.AddTaskAttachment,
		opts...,
	)
	taskServiceDeleteTaskAttachmentHandler := connect_go.NewUnaryHandler(
		TaskServiceDeleteTaskAttachmentProcedure,
		svc.DeleteTaskAttachment,
		opts...,
	)
	return "/tkd.tasks.v1.TaskService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TaskServiceCreateTaskProcedure:
			taskServiceCreateTaskHandler.ServeHTTP(w, r)
		case TaskServiceUpdateTaskProcedure:
			taskServiceUpdateTaskHandler.ServeHTTP(w, r)
		case TaskServiceAssignTaskProcedure:
			taskServiceAssignTaskHandler.ServeHTTP(w, r)
		case TaskServiceCompleteTaskProcedure:
			taskServiceCompleteTaskHandler.ServeHTTP(w, r)
		case TaskServiceDeleteTaskProcedure:
			taskServiceDeleteTaskHandler.ServeHTTP(w, r)
		case TaskServiceListTasksProcedure:
			taskServiceListTasksHandler.ServeHTTP(w, r)
		case TaskServiceGetTaskProcedure:
			taskServiceGetTaskHandler.ServeHTTP(w, r)
		case TaskServiceAddTaskAttachmentProcedure:
			taskServiceAddTaskAttachmentHandler.ServeHTTP(w, r)
		case TaskServiceDeleteTaskAttachmentProcedure:
			taskServiceDeleteTaskAttachmentHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTaskServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTaskServiceHandler struct{}

func (UnimplementedTaskServiceHandler) CreateTask(context.Context, *connect_go.Request[v1.CreateTaskRequest]) (*connect_go.Response[v1.CreateTaskResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.tasks.v1.TaskService.CreateTask is not implemented"))
}

func (UnimplementedTaskServiceHandler) UpdateTask(context.Context, *connect_go.Request[v1.UpdateTaskRequest]) (*connect_go.Response[v1.UpdateTaskResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.tasks.v1.TaskService.UpdateTask is not implemented"))
}

func (UnimplementedTaskServiceHandler) AssignTask(context.Context, *connect_go.Request[v1.AssignTaskRequest]) (*connect_go.Response[v1.AssignTaskResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.tasks.v1.TaskService.AssignTask is not implemented"))
}

func (UnimplementedTaskServiceHandler) CompleteTask(context.Context, *connect_go.Request[v1.CompleteTaskRequest]) (*connect_go.Response[v1.CompleteTaskResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.tasks.v1.TaskService.CompleteTask is not implemented"))
}

func (UnimplementedTaskServiceHandler) DeleteTask(context.Context, *connect_go.Request[v1.DeleteTaskRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.tasks.v1.TaskService.DeleteTask is not implemented"))
}

func (UnimplementedTaskServiceHandler) ListTasks(context.Context, *connect_go.Request[v1.ListTasksRequest]) (*connect_go.Response[v1.ListTasksResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.tasks.v1.TaskService.ListTasks is not implemented"))
}

func (UnimplementedTaskServiceHandler) GetTask(context.Context, *connect_go.Request[v1.GetTaskRequest]) (*connect_go.Response[v1.GetTaskResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.tasks.v1.TaskService.GetTask is not implemented"))
}

func (UnimplementedTaskServiceHandler) AddTaskAttachment(context.Context, *connect_go.Request[v1.AddTaskAttachmentRequest]) (*connect_go.Response[v1.AddTaskAttachmentResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.tasks.v1.TaskService.AddTaskAttachment is not implemented"))
}

func (UnimplementedTaskServiceHandler) DeleteTaskAttachment(context.Context, *connect_go.Request[v1.DeleteTaskAttachmentRequest]) (*connect_go.Response[v1.DeleteTaskAttachmentResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tkd.tasks.v1.TaskService.DeleteTaskAttachment is not implemented"))
}