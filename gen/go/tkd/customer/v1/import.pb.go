// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: tkd/customer/v1/import.proto

package customerv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/tierklinik-dobersberg/apis/gen/go/tkd/common/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StartSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Importer string `protobuf:"bytes,1,opt,name=importer,proto3" json:"importer,omitempty"`
}

func (x *StartSessionRequest) Reset() {
	*x = StartSessionRequest{}
	mi := &file_tkd_customer_v1_import_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartSessionRequest) ProtoMessage() {}

func (x *StartSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_customer_v1_import_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartSessionRequest.ProtoReflect.Descriptor instead.
func (*StartSessionRequest) Descriptor() ([]byte, []int) {
	return file_tkd_customer_v1_import_proto_rawDescGZIP(), []int{0}
}

func (x *StartSessionRequest) GetImporter() string {
	if x != nil {
		return x.Importer
	}
	return ""
}

type StartSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartSessionResponse) Reset() {
	*x = StartSessionResponse{}
	mi := &file_tkd_customer_v1_import_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartSessionResponse) ProtoMessage() {}

func (x *StartSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_customer_v1_import_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartSessionResponse.ProtoReflect.Descriptor instead.
func (*StartSessionResponse) Descriptor() ([]byte, []int) {
	return file_tkd_customer_v1_import_proto_rawDescGZIP(), []int{1}
}

type LookupCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query *CustomerQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *LookupCustomerRequest) Reset() {
	*x = LookupCustomerRequest{}
	mi := &file_tkd_customer_v1_import_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LookupCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupCustomerRequest) ProtoMessage() {}

func (x *LookupCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_customer_v1_import_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupCustomerRequest.ProtoReflect.Descriptor instead.
func (*LookupCustomerRequest) Descriptor() ([]byte, []int) {
	return file_tkd_customer_v1_import_proto_rawDescGZIP(), []int{2}
}

func (x *LookupCustomerRequest) GetQuery() *CustomerQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

type ImportedCustomer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customer *Customer    `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
	State    *ImportState `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *ImportedCustomer) Reset() {
	*x = ImportedCustomer{}
	mi := &file_tkd_customer_v1_import_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImportedCustomer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportedCustomer) ProtoMessage() {}

func (x *ImportedCustomer) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_customer_v1_import_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportedCustomer.ProtoReflect.Descriptor instead.
func (*ImportedCustomer) Descriptor() ([]byte, []int) {
	return file_tkd_customer_v1_import_proto_rawDescGZIP(), []int{3}
}

func (x *ImportedCustomer) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

func (x *ImportedCustomer) GetState() *ImportState {
	if x != nil {
		return x.State
	}
	return nil
}

type LookupCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchedCustomers []*ImportedCustomer `protobuf:"bytes,1,rep,name=matched_customers,json=matchedCustomers,proto3" json:"matched_customers,omitempty"`
}

func (x *LookupCustomerResponse) Reset() {
	*x = LookupCustomerResponse{}
	mi := &file_tkd_customer_v1_import_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LookupCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupCustomerResponse) ProtoMessage() {}

func (x *LookupCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_customer_v1_import_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupCustomerResponse.ProtoReflect.Descriptor instead.
func (*LookupCustomerResponse) Descriptor() ([]byte, []int) {
	return file_tkd_customer_v1_import_proto_rawDescGZIP(), []int{4}
}

func (x *LookupCustomerResponse) GetMatchedCustomers() []*ImportedCustomer {
	if x != nil {
		return x.MatchedCustomers
	}
	return nil
}

type UpsertCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InternalReference string           `protobuf:"bytes,1,opt,name=internal_reference,json=internalReference,proto3" json:"internal_reference,omitempty"`
	ExtraData         *structpb.Struct `protobuf:"bytes,2,opt,name=extra_data,json=extraData,proto3" json:"extra_data,omitempty"`
	Customer          *Customer        `protobuf:"bytes,3,opt,name=customer,proto3" json:"customer,omitempty"`
}

func (x *UpsertCustomerRequest) Reset() {
	*x = UpsertCustomerRequest{}
	mi := &file_tkd_customer_v1_import_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertCustomerRequest) ProtoMessage() {}

func (x *UpsertCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_customer_v1_import_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertCustomerRequest.ProtoReflect.Descriptor instead.
func (*UpsertCustomerRequest) Descriptor() ([]byte, []int) {
	return file_tkd_customer_v1_import_proto_rawDescGZIP(), []int{5}
}

func (x *UpsertCustomerRequest) GetInternalReference() string {
	if x != nil {
		return x.InternalReference
	}
	return ""
}

func (x *UpsertCustomerRequest) GetExtraData() *structpb.Struct {
	if x != nil {
		return x.ExtraData
	}
	return nil
}

func (x *UpsertCustomerRequest) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

type UpsertCustomerSuccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpsertCustomerSuccess) Reset() {
	*x = UpsertCustomerSuccess{}
	mi := &file_tkd_customer_v1_import_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertCustomerSuccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertCustomerSuccess) ProtoMessage() {}

func (x *UpsertCustomerSuccess) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_customer_v1_import_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertCustomerSuccess.ProtoReflect.Descriptor instead.
func (*UpsertCustomerSuccess) Descriptor() ([]byte, []int) {
	return file_tkd_customer_v1_import_proto_rawDescGZIP(), []int{6}
}

func (x *UpsertCustomerSuccess) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error []string `protobuf:"bytes,1,rep,name=error,proto3" json:"error,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	mi := &file_tkd_customer_v1_import_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_customer_v1_import_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_tkd_customer_v1_import_proto_rawDescGZIP(), []int{7}
}

func (x *Error) GetError() []string {
	if x != nil {
		return x.Error
	}
	return nil
}

type ImportSessionComplete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ImportSessionComplete) Reset() {
	*x = ImportSessionComplete{}
	mi := &file_tkd_customer_v1_import_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImportSessionComplete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportSessionComplete) ProtoMessage() {}

func (x *ImportSessionComplete) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_customer_v1_import_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportSessionComplete.ProtoReflect.Descriptor instead.
func (*ImportSessionComplete) Descriptor() ([]byte, []int) {
	return file_tkd_customer_v1_import_proto_rawDescGZIP(), []int{8}
}

type ImportSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//
	//	*ImportSessionRequest_StartSession
	//	*ImportSessionRequest_LookupCustomer
	//	*ImportSessionRequest_UpsertCustomer
	//	*ImportSessionRequest_Complete
	Message       isImportSessionRequest_Message `protobuf_oneof:"message"`
	CorrelationId string                         `protobuf:"bytes,99,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
}

func (x *ImportSessionRequest) Reset() {
	*x = ImportSessionRequest{}
	mi := &file_tkd_customer_v1_import_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImportSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportSessionRequest) ProtoMessage() {}

func (x *ImportSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_customer_v1_import_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportSessionRequest.ProtoReflect.Descriptor instead.
func (*ImportSessionRequest) Descriptor() ([]byte, []int) {
	return file_tkd_customer_v1_import_proto_rawDescGZIP(), []int{9}
}

func (m *ImportSessionRequest) GetMessage() isImportSessionRequest_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *ImportSessionRequest) GetStartSession() *StartSessionRequest {
	if x, ok := x.GetMessage().(*ImportSessionRequest_StartSession); ok {
		return x.StartSession
	}
	return nil
}

func (x *ImportSessionRequest) GetLookupCustomer() *LookupCustomerRequest {
	if x, ok := x.GetMessage().(*ImportSessionRequest_LookupCustomer); ok {
		return x.LookupCustomer
	}
	return nil
}

func (x *ImportSessionRequest) GetUpsertCustomer() *UpsertCustomerRequest {
	if x, ok := x.GetMessage().(*ImportSessionRequest_UpsertCustomer); ok {
		return x.UpsertCustomer
	}
	return nil
}

func (x *ImportSessionRequest) GetComplete() *ImportSessionComplete {
	if x, ok := x.GetMessage().(*ImportSessionRequest_Complete); ok {
		return x.Complete
	}
	return nil
}

func (x *ImportSessionRequest) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

type isImportSessionRequest_Message interface {
	isImportSessionRequest_Message()
}

type ImportSessionRequest_StartSession struct {
	// StartSessionRequest must be the very first message sent
	// for a ImportSession RPC. It must not be sent more than once.
	StartSession *StartSessionRequest `protobuf:"bytes,1,opt,name=start_session,json=startSession,proto3,oneof"`
}

type ImportSessionRequest_LookupCustomer struct {
	LookupCustomer *LookupCustomerRequest `protobuf:"bytes,2,opt,name=lookup_customer,json=lookupCustomer,proto3,oneof"`
}

type ImportSessionRequest_UpsertCustomer struct {
	UpsertCustomer *UpsertCustomerRequest `protobuf:"bytes,3,opt,name=upsert_customer,json=upsertCustomer,proto3,oneof"`
}

type ImportSessionRequest_Complete struct {
	Complete *ImportSessionComplete `protobuf:"bytes,10,opt,name=complete,proto3,oneof"`
}

func (*ImportSessionRequest_StartSession) isImportSessionRequest_Message() {}

func (*ImportSessionRequest_LookupCustomer) isImportSessionRequest_Message() {}

func (*ImportSessionRequest_UpsertCustomer) isImportSessionRequest_Message() {}

func (*ImportSessionRequest_Complete) isImportSessionRequest_Message() {}

type ImportSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//
	//	*ImportSessionResponse_StartSession
	//	*ImportSessionResponse_LookupCustomer
	//	*ImportSessionResponse_UpsertSuccess
	//	*ImportSessionResponse_Error
	Message       isImportSessionResponse_Message `protobuf_oneof:"message"`
	CorrelationId string                          `protobuf:"bytes,99,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
}

func (x *ImportSessionResponse) Reset() {
	*x = ImportSessionResponse{}
	mi := &file_tkd_customer_v1_import_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImportSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportSessionResponse) ProtoMessage() {}

func (x *ImportSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_customer_v1_import_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportSessionResponse.ProtoReflect.Descriptor instead.
func (*ImportSessionResponse) Descriptor() ([]byte, []int) {
	return file_tkd_customer_v1_import_proto_rawDescGZIP(), []int{10}
}

func (m *ImportSessionResponse) GetMessage() isImportSessionResponse_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *ImportSessionResponse) GetStartSession() *StartSessionResponse {
	if x, ok := x.GetMessage().(*ImportSessionResponse_StartSession); ok {
		return x.StartSession
	}
	return nil
}

func (x *ImportSessionResponse) GetLookupCustomer() *LookupCustomerResponse {
	if x, ok := x.GetMessage().(*ImportSessionResponse_LookupCustomer); ok {
		return x.LookupCustomer
	}
	return nil
}

func (x *ImportSessionResponse) GetUpsertSuccess() *UpsertCustomerSuccess {
	if x, ok := x.GetMessage().(*ImportSessionResponse_UpsertSuccess); ok {
		return x.UpsertSuccess
	}
	return nil
}

func (x *ImportSessionResponse) GetError() *Error {
	if x, ok := x.GetMessage().(*ImportSessionResponse_Error); ok {
		return x.Error
	}
	return nil
}

func (x *ImportSessionResponse) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

type isImportSessionResponse_Message interface {
	isImportSessionResponse_Message()
}

type ImportSessionResponse_StartSession struct {
	StartSession *StartSessionResponse `protobuf:"bytes,1,opt,name=start_session,json=startSession,proto3,oneof"`
}

type ImportSessionResponse_LookupCustomer struct {
	LookupCustomer *LookupCustomerResponse `protobuf:"bytes,2,opt,name=lookup_customer,json=lookupCustomer,proto3,oneof"`
}

type ImportSessionResponse_UpsertSuccess struct {
	UpsertSuccess *UpsertCustomerSuccess `protobuf:"bytes,3,opt,name=upsert_success,json=upsertSuccess,proto3,oneof"`
}

type ImportSessionResponse_Error struct {
	Error *Error `protobuf:"bytes,4,opt,name=error,proto3,oneof"`
}

func (*ImportSessionResponse_StartSession) isImportSessionResponse_Message() {}

func (*ImportSessionResponse_LookupCustomer) isImportSessionResponse_Message() {}

func (*ImportSessionResponse_UpsertSuccess) isImportSessionResponse_Message() {}

func (*ImportSessionResponse_Error) isImportSessionResponse_Message() {}

var File_tkd_customer_v1_import_proto protoreflect.FileDescriptor

var file_tkd_customer_v1_import_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x74, 0x6b, 0x64, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x74, 0x6b, 0x64, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a,
	0x1e, 0x74, 0x6b, 0x64, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x74, 0x6b, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62,
	0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x13, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x23, 0x0a, 0x08, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0xf7, 0x18, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x69, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x22, 0x16, 0x0a, 0x14, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4d,
	0x0a, 0x15, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x7d, 0x0a,
	0x10, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x12, 0x35, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x08,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x68, 0x0a, 0x16,
	0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x11, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x64, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x52, 0x10, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x64, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x22, 0xc7, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x73, 0x65, 0x72,
	0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x36, 0x0a, 0x12, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0xf7,
	0x18, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x09, 0x65, 0x78, 0x74, 0x72, 0x61, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x3e, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x07, 0xfa,
	0xf7, 0x18, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x22, 0x27, 0x0a, 0x15, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1d, 0x0a, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x17, 0x0a, 0x15, 0x49, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x22, 0x89, 0x03, 0x0a, 0x14, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x0d, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x51, 0x0a, 0x0f, 0x6c, 0x6f, 0x6f, 0x6b, 0x75,
	0x70, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x6c, 0x6f, 0x6f, 0x6b,
	0x75, 0x70, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x51, 0x0a, 0x0f, 0x75, 0x70,
	0x73, 0x65, 0x72, 0x74, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0e, 0x75,
	0x70, 0x73, 0x65, 0x72, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x44, 0x0a,
	0x08, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x00, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x63, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x42, 0x11, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x06, 0xfa, 0xf7, 0x18, 0x02, 0x08, 0x01, 0x22, 0xec, 0x02,
	0x0a, 0x15, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x52, 0x0a, 0x0f, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x5f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x6c, 0x6f, 0x6f, 0x6b, 0x75,
	0x70, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x4f, 0x0a, 0x0e, 0x75, 0x70, 0x73,
	0x65, 0x72, 0x74, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x48, 0x00, 0x52, 0x0d, 0x75, 0x70, 0x73,
	0x65, 0x72, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2e, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x6b, 0x64, 0x2e,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f,
	0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x63, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xa8, 0x01, 0x0a,
	0x15, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x0d, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x05, 0xb2, 0x7e, 0x02, 0x08, 0x02, 0x28, 0x01, 0x30,
	0x01, 0x1a, 0x24, 0xba, 0x7e, 0x21, 0x0a, 0x0d, 0x69, 0x64, 0x6d, 0x5f, 0x73, 0x75, 0x70, 0x65,
	0x72, 0x75, 0x73, 0x65, 0x72, 0x0a, 0x10, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x42, 0xc9, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e,
	0x74, 0x6b, 0x64, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42,
	0x0b, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x47,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x65, 0x72, 0x6b,
	0x6c, 0x69, 0x6e, 0x69, 0x6b, 0x2d, 0x64, 0x6f, 0x62, 0x65, 0x72, 0x73, 0x62, 0x65, 0x72, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x6b, 0x64,
	0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x43, 0x58, 0xaa, 0x02, 0x0f,
	0x54, 0x6b, 0x64, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x0f, 0x54, 0x6b, 0x64, 0x5c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x1b, 0x54, 0x6b, 0x64, 0x5c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x11, 0x54, 0x6b, 0x64, 0x3a, 0x3a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tkd_customer_v1_import_proto_rawDescOnce sync.Once
	file_tkd_customer_v1_import_proto_rawDescData = file_tkd_customer_v1_import_proto_rawDesc
)

func file_tkd_customer_v1_import_proto_rawDescGZIP() []byte {
	file_tkd_customer_v1_import_proto_rawDescOnce.Do(func() {
		file_tkd_customer_v1_import_proto_rawDescData = protoimpl.X.CompressGZIP(file_tkd_customer_v1_import_proto_rawDescData)
	})
	return file_tkd_customer_v1_import_proto_rawDescData
}

var file_tkd_customer_v1_import_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_tkd_customer_v1_import_proto_goTypes = []any{
	(*StartSessionRequest)(nil),    // 0: tkd.customer.v1.StartSessionRequest
	(*StartSessionResponse)(nil),   // 1: tkd.customer.v1.StartSessionResponse
	(*LookupCustomerRequest)(nil),  // 2: tkd.customer.v1.LookupCustomerRequest
	(*ImportedCustomer)(nil),       // 3: tkd.customer.v1.ImportedCustomer
	(*LookupCustomerResponse)(nil), // 4: tkd.customer.v1.LookupCustomerResponse
	(*UpsertCustomerRequest)(nil),  // 5: tkd.customer.v1.UpsertCustomerRequest
	(*UpsertCustomerSuccess)(nil),  // 6: tkd.customer.v1.UpsertCustomerSuccess
	(*Error)(nil),                  // 7: tkd.customer.v1.Error
	(*ImportSessionComplete)(nil),  // 8: tkd.customer.v1.ImportSessionComplete
	(*ImportSessionRequest)(nil),   // 9: tkd.customer.v1.ImportSessionRequest
	(*ImportSessionResponse)(nil),  // 10: tkd.customer.v1.ImportSessionResponse
	(*CustomerQuery)(nil),          // 11: tkd.customer.v1.CustomerQuery
	(*Customer)(nil),               // 12: tkd.customer.v1.Customer
	(*ImportState)(nil),            // 13: tkd.customer.v1.ImportState
	(*structpb.Struct)(nil),        // 14: google.protobuf.Struct
}
var file_tkd_customer_v1_import_proto_depIdxs = []int32{
	11, // 0: tkd.customer.v1.LookupCustomerRequest.query:type_name -> tkd.customer.v1.CustomerQuery
	12, // 1: tkd.customer.v1.ImportedCustomer.customer:type_name -> tkd.customer.v1.Customer
	13, // 2: tkd.customer.v1.ImportedCustomer.state:type_name -> tkd.customer.v1.ImportState
	3,  // 3: tkd.customer.v1.LookupCustomerResponse.matched_customers:type_name -> tkd.customer.v1.ImportedCustomer
	14, // 4: tkd.customer.v1.UpsertCustomerRequest.extra_data:type_name -> google.protobuf.Struct
	12, // 5: tkd.customer.v1.UpsertCustomerRequest.customer:type_name -> tkd.customer.v1.Customer
	0,  // 6: tkd.customer.v1.ImportSessionRequest.start_session:type_name -> tkd.customer.v1.StartSessionRequest
	2,  // 7: tkd.customer.v1.ImportSessionRequest.lookup_customer:type_name -> tkd.customer.v1.LookupCustomerRequest
	5,  // 8: tkd.customer.v1.ImportSessionRequest.upsert_customer:type_name -> tkd.customer.v1.UpsertCustomerRequest
	8,  // 9: tkd.customer.v1.ImportSessionRequest.complete:type_name -> tkd.customer.v1.ImportSessionComplete
	1,  // 10: tkd.customer.v1.ImportSessionResponse.start_session:type_name -> tkd.customer.v1.StartSessionResponse
	4,  // 11: tkd.customer.v1.ImportSessionResponse.lookup_customer:type_name -> tkd.customer.v1.LookupCustomerResponse
	6,  // 12: tkd.customer.v1.ImportSessionResponse.upsert_success:type_name -> tkd.customer.v1.UpsertCustomerSuccess
	7,  // 13: tkd.customer.v1.ImportSessionResponse.error:type_name -> tkd.customer.v1.Error
	9,  // 14: tkd.customer.v1.CustomerImportService.ImportSession:input_type -> tkd.customer.v1.ImportSessionRequest
	10, // 15: tkd.customer.v1.CustomerImportService.ImportSession:output_type -> tkd.customer.v1.ImportSessionResponse
	15, // [15:16] is the sub-list for method output_type
	14, // [14:15] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_tkd_customer_v1_import_proto_init() }
func file_tkd_customer_v1_import_proto_init() {
	if File_tkd_customer_v1_import_proto != nil {
		return
	}
	file_tkd_customer_v1_customer_proto_init()
	file_tkd_customer_v1_import_proto_msgTypes[9].OneofWrappers = []any{
		(*ImportSessionRequest_StartSession)(nil),
		(*ImportSessionRequest_LookupCustomer)(nil),
		(*ImportSessionRequest_UpsertCustomer)(nil),
		(*ImportSessionRequest_Complete)(nil),
	}
	file_tkd_customer_v1_import_proto_msgTypes[10].OneofWrappers = []any{
		(*ImportSessionResponse_StartSession)(nil),
		(*ImportSessionResponse_LookupCustomer)(nil),
		(*ImportSessionResponse_UpsertSuccess)(nil),
		(*ImportSessionResponse_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tkd_customer_v1_import_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tkd_customer_v1_import_proto_goTypes,
		DependencyIndexes: file_tkd_customer_v1_import_proto_depIdxs,
		MessageInfos:      file_tkd_customer_v1_import_proto_msgTypes,
	}.Build()
	File_tkd_customer_v1_import_proto = out.File
	file_tkd_customer_v1_import_proto_rawDesc = nil
	file_tkd_customer_v1_import_proto_goTypes = nil
	file_tkd_customer_v1_import_proto_depIdxs = nil
}
