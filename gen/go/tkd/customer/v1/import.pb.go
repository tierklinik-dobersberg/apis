// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
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
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StartSessionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Importer      string                 `protobuf:"bytes,1,opt,name=importer,proto3" json:"importer,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         *CustomerQuery         `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Customer      *Customer              `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
	State         *ImportState           `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	state            protoimpl.MessageState `protogen:"open.v1"`
	MatchedCustomers []*ImportedCustomer    `protobuf:"bytes,1,rep,name=matched_customers,json=matchedCustomers,proto3" json:"matched_customers,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
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
	state             protoimpl.MessageState `protogen:"open.v1"`
	InternalReference string                 `protobuf:"bytes,1,opt,name=internal_reference,json=internalReference,proto3" json:"internal_reference,omitempty"`
	ExtraData         *structpb.Struct       `protobuf:"bytes,2,opt,name=extra_data,json=extraData,proto3" json:"extra_data,omitempty"`
	Customer          *Customer              `protobuf:"bytes,3,opt,name=customer,proto3" json:"customer,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
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

type UpsertPatientRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Patient       *Patient               `protobuf:"bytes,1,opt,name=patient,proto3" json:"patient,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpsertPatientRequest) Reset() {
	*x = UpsertPatientRequest{}
	mi := &file_tkd_customer_v1_import_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertPatientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertPatientRequest) ProtoMessage() {}

func (x *UpsertPatientRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UpsertPatientRequest.ProtoReflect.Descriptor instead.
func (*UpsertPatientRequest) Descriptor() ([]byte, []int) {
	return file_tkd_customer_v1_import_proto_rawDescGZIP(), []int{6}
}

func (x *UpsertPatientRequest) GetPatient() *Patient {
	if x != nil {
		return x.Patient
	}
	return nil
}

type UpsertPatientSuccess struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpsertPatientSuccess) Reset() {
	*x = UpsertPatientSuccess{}
	mi := &file_tkd_customer_v1_import_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertPatientSuccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertPatientSuccess) ProtoMessage() {}

func (x *UpsertPatientSuccess) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UpsertPatientSuccess.ProtoReflect.Descriptor instead.
func (*UpsertPatientSuccess) Descriptor() ([]byte, []int) {
	return file_tkd_customer_v1_import_proto_rawDescGZIP(), []int{7}
}

func (x *UpsertPatientSuccess) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpsertCustomerSuccess struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpsertCustomerSuccess) Reset() {
	*x = UpsertCustomerSuccess{}
	mi := &file_tkd_customer_v1_import_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertCustomerSuccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertCustomerSuccess) ProtoMessage() {}

func (x *UpsertCustomerSuccess) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UpsertCustomerSuccess.ProtoReflect.Descriptor instead.
func (*UpsertCustomerSuccess) Descriptor() ([]byte, []int) {
	return file_tkd_customer_v1_import_proto_rawDescGZIP(), []int{8}
}

func (x *UpsertCustomerSuccess) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Error struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Error         []string               `protobuf:"bytes,1,rep,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Error) Reset() {
	*x = Error{}
	mi := &file_tkd_customer_v1_import_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_tkd_customer_v1_import_proto_rawDescGZIP(), []int{9}
}

func (x *Error) GetError() []string {
	if x != nil {
		return x.Error
	}
	return nil
}

type ImportSessionComplete struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImportSessionComplete) Reset() {
	*x = ImportSessionComplete{}
	mi := &file_tkd_customer_v1_import_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImportSessionComplete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportSessionComplete) ProtoMessage() {}

func (x *ImportSessionComplete) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ImportSessionComplete.ProtoReflect.Descriptor instead.
func (*ImportSessionComplete) Descriptor() ([]byte, []int) {
	return file_tkd_customer_v1_import_proto_rawDescGZIP(), []int{10}
}

type ImportSessionRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Message:
	//
	//	*ImportSessionRequest_StartSession
	//	*ImportSessionRequest_LookupCustomer
	//	*ImportSessionRequest_UpsertCustomer
	//	*ImportSessionRequest_UpsertPatient
	//	*ImportSessionRequest_Complete
	Message       isImportSessionRequest_Message `protobuf_oneof:"message"`
	CorrelationId string                         `protobuf:"bytes,99,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImportSessionRequest) Reset() {
	*x = ImportSessionRequest{}
	mi := &file_tkd_customer_v1_import_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImportSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportSessionRequest) ProtoMessage() {}

func (x *ImportSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_customer_v1_import_proto_msgTypes[11]
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
	return file_tkd_customer_v1_import_proto_rawDescGZIP(), []int{11}
}

func (x *ImportSessionRequest) GetMessage() isImportSessionRequest_Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *ImportSessionRequest) GetStartSession() *StartSessionRequest {
	if x != nil {
		if x, ok := x.Message.(*ImportSessionRequest_StartSession); ok {
			return x.StartSession
		}
	}
	return nil
}

func (x *ImportSessionRequest) GetLookupCustomer() *LookupCustomerRequest {
	if x != nil {
		if x, ok := x.Message.(*ImportSessionRequest_LookupCustomer); ok {
			return x.LookupCustomer
		}
	}
	return nil
}

func (x *ImportSessionRequest) GetUpsertCustomer() *UpsertCustomerRequest {
	if x != nil {
		if x, ok := x.Message.(*ImportSessionRequest_UpsertCustomer); ok {
			return x.UpsertCustomer
		}
	}
	return nil
}

func (x *ImportSessionRequest) GetUpsertPatient() *UpsertPatientRequest {
	if x != nil {
		if x, ok := x.Message.(*ImportSessionRequest_UpsertPatient); ok {
			return x.UpsertPatient
		}
	}
	return nil
}

func (x *ImportSessionRequest) GetComplete() *ImportSessionComplete {
	if x != nil {
		if x, ok := x.Message.(*ImportSessionRequest_Complete); ok {
			return x.Complete
		}
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

type ImportSessionRequest_UpsertPatient struct {
	UpsertPatient *UpsertPatientRequest `protobuf:"bytes,4,opt,name=upsert_patient,json=upsertPatient,proto3,oneof"`
}

type ImportSessionRequest_Complete struct {
	Complete *ImportSessionComplete `protobuf:"bytes,10,opt,name=complete,proto3,oneof"`
}

func (*ImportSessionRequest_StartSession) isImportSessionRequest_Message() {}

func (*ImportSessionRequest_LookupCustomer) isImportSessionRequest_Message() {}

func (*ImportSessionRequest_UpsertCustomer) isImportSessionRequest_Message() {}

func (*ImportSessionRequest_UpsertPatient) isImportSessionRequest_Message() {}

func (*ImportSessionRequest_Complete) isImportSessionRequest_Message() {}

type ImportSessionResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Message:
	//
	//	*ImportSessionResponse_StartSession
	//	*ImportSessionResponse_LookupCustomer
	//	*ImportSessionResponse_UpsertCustomerSuccess
	//	*ImportSessionResponse_UpsertPatientSuccess
	//	*ImportSessionResponse_Error
	Message       isImportSessionResponse_Message `protobuf_oneof:"message"`
	CorrelationId string                          `protobuf:"bytes,99,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImportSessionResponse) Reset() {
	*x = ImportSessionResponse{}
	mi := &file_tkd_customer_v1_import_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImportSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportSessionResponse) ProtoMessage() {}

func (x *ImportSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_customer_v1_import_proto_msgTypes[12]
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
	return file_tkd_customer_v1_import_proto_rawDescGZIP(), []int{12}
}

func (x *ImportSessionResponse) GetMessage() isImportSessionResponse_Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *ImportSessionResponse) GetStartSession() *StartSessionResponse {
	if x != nil {
		if x, ok := x.Message.(*ImportSessionResponse_StartSession); ok {
			return x.StartSession
		}
	}
	return nil
}

func (x *ImportSessionResponse) GetLookupCustomer() *LookupCustomerResponse {
	if x != nil {
		if x, ok := x.Message.(*ImportSessionResponse_LookupCustomer); ok {
			return x.LookupCustomer
		}
	}
	return nil
}

func (x *ImportSessionResponse) GetUpsertCustomerSuccess() *UpsertCustomerSuccess {
	if x != nil {
		if x, ok := x.Message.(*ImportSessionResponse_UpsertCustomerSuccess); ok {
			return x.UpsertCustomerSuccess
		}
	}
	return nil
}

func (x *ImportSessionResponse) GetUpsertPatientSuccess() *UpsertPatientSuccess {
	if x != nil {
		if x, ok := x.Message.(*ImportSessionResponse_UpsertPatientSuccess); ok {
			return x.UpsertPatientSuccess
		}
	}
	return nil
}

func (x *ImportSessionResponse) GetError() *Error {
	if x != nil {
		if x, ok := x.Message.(*ImportSessionResponse_Error); ok {
			return x.Error
		}
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

type ImportSessionResponse_UpsertCustomerSuccess struct {
	UpsertCustomerSuccess *UpsertCustomerSuccess `protobuf:"bytes,3,opt,name=upsert_customer_success,json=upsertCustomerSuccess,proto3,oneof"`
}

type ImportSessionResponse_UpsertPatientSuccess struct {
	UpsertPatientSuccess *UpsertPatientSuccess `protobuf:"bytes,5,opt,name=upsert_patient_success,json=upsertPatientSuccess,proto3,oneof"`
}

type ImportSessionResponse_Error struct {
	Error *Error `protobuf:"bytes,4,opt,name=error,proto3,oneof"`
}

func (*ImportSessionResponse_StartSession) isImportSessionResponse_Message() {}

func (*ImportSessionResponse_LookupCustomer) isImportSessionResponse_Message() {}

func (*ImportSessionResponse_UpsertCustomerSuccess) isImportSessionResponse_Message() {}

func (*ImportSessionResponse_UpsertPatientSuccess) isImportSessionResponse_Message() {}

func (*ImportSessionResponse_Error) isImportSessionResponse_Message() {}

var File_tkd_customer_v1_import_proto protoreflect.FileDescriptor

const file_tkd_customer_v1_import_proto_rawDesc = "" +
	"\n" +
	"\x1ctkd/customer/v1/import.proto\x12\x0ftkd.customer.v1\x1a\x1etkd/customer/v1/customer.proto\x1a\x1dtkd/customer/v1/patient.proto\x1a\x1etkd/common/v1/descriptor.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1bbuf/validate/validate.proto\":\n" +
	"\x13StartSessionRequest\x12#\n" +
	"\bimporter\x18\x01 \x01(\tB\a\xfa\xf7\x18\x03\xc8\x01\x01R\bimporter\"\x16\n" +
	"\x14StartSessionResponse\"M\n" +
	"\x15LookupCustomerRequest\x124\n" +
	"\x05query\x18\x01 \x01(\v2\x1e.tkd.customer.v1.CustomerQueryR\x05query\"}\n" +
	"\x10ImportedCustomer\x125\n" +
	"\bcustomer\x18\x01 \x01(\v2\x19.tkd.customer.v1.CustomerR\bcustomer\x122\n" +
	"\x05state\x18\x02 \x01(\v2\x1c.tkd.customer.v1.ImportStateR\x05state\"h\n" +
	"\x16LookupCustomerResponse\x12N\n" +
	"\x11matched_customers\x18\x01 \x03(\v2!.tkd.customer.v1.ImportedCustomerR\x10matchedCustomers\"\xc7\x01\n" +
	"\x15UpsertCustomerRequest\x126\n" +
	"\x12internal_reference\x18\x01 \x01(\tB\a\xfa\xf7\x18\x03\xc8\x01\x01R\x11internalReference\x126\n" +
	"\n" +
	"extra_data\x18\x02 \x01(\v2\x17.google.protobuf.StructR\textraData\x12>\n" +
	"\bcustomer\x18\x03 \x01(\v2\x19.tkd.customer.v1.CustomerB\a\xfa\xf7\x18\x03\xc8\x01\x01R\bcustomer\"S\n" +
	"\x14UpsertPatientRequest\x12;\n" +
	"\apatient\x18\x01 \x01(\v2\x18.tkd.customer.v1.PatientB\a\xfa\xf7\x18\x03\xc8\x01\x01R\apatient\"&\n" +
	"\x14UpsertPatientSuccess\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"'\n" +
	"\x15UpsertCustomerSuccess\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"\x1d\n" +
	"\x05Error\x12\x14\n" +
	"\x05error\x18\x01 \x03(\tR\x05error\"\x17\n" +
	"\x15ImportSessionComplete\"\xd9\x03\n" +
	"\x14ImportSessionRequest\x12K\n" +
	"\rstart_session\x18\x01 \x01(\v2$.tkd.customer.v1.StartSessionRequestH\x00R\fstartSession\x12Q\n" +
	"\x0flookup_customer\x18\x02 \x01(\v2&.tkd.customer.v1.LookupCustomerRequestH\x00R\x0elookupCustomer\x12Q\n" +
	"\x0fupsert_customer\x18\x03 \x01(\v2&.tkd.customer.v1.UpsertCustomerRequestH\x00R\x0eupsertCustomer\x12N\n" +
	"\x0eupsert_patient\x18\x04 \x01(\v2%.tkd.customer.v1.UpsertPatientRequestH\x00R\rupsertPatient\x12D\n" +
	"\bcomplete\x18\n" +
	" \x01(\v2&.tkd.customer.v1.ImportSessionCompleteH\x00R\bcomplete\x12%\n" +
	"\x0ecorrelation_id\x18c \x01(\tR\rcorrelationIdB\x11\n" +
	"\amessage\x12\x06\xfa\xf7\x18\x02\b\x01\"\xdc\x03\n" +
	"\x15ImportSessionResponse\x12L\n" +
	"\rstart_session\x18\x01 \x01(\v2%.tkd.customer.v1.StartSessionResponseH\x00R\fstartSession\x12R\n" +
	"\x0flookup_customer\x18\x02 \x01(\v2'.tkd.customer.v1.LookupCustomerResponseH\x00R\x0elookupCustomer\x12`\n" +
	"\x17upsert_customer_success\x18\x03 \x01(\v2&.tkd.customer.v1.UpsertCustomerSuccessH\x00R\x15upsertCustomerSuccess\x12]\n" +
	"\x16upsert_patient_success\x18\x05 \x01(\v2%.tkd.customer.v1.UpsertPatientSuccessH\x00R\x14upsertPatientSuccess\x12.\n" +
	"\x05error\x18\x04 \x01(\v2\x16.tkd.customer.v1.ErrorH\x00R\x05error\x12%\n" +
	"\x0ecorrelation_id\x18c \x01(\tR\rcorrelationIdB\t\n" +
	"\amessage2\xa8\x01\n" +
	"\x15CustomerImportService\x12i\n" +
	"\rImportSession\x12%.tkd.customer.v1.ImportSessionRequest\x1a&.tkd.customer.v1.ImportSessionResponse\"\x05\xb2~\x02\b\x02(\x010\x01\x1a$\xba~!\n" +
	"\ridm_superuser\n" +
	"\x10customer_managerB\xc9\x01\n" +
	"\x13com.tkd.customer.v1B\vImportProtoP\x01ZGgithub.com/tierklinik-dobersberg/apis/gen/go/tkd/customer/v1;customerv1\xa2\x02\x03TCX\xaa\x02\x0fTkd.Customer.V1\xca\x02\x0fTkd\\Customer\\V1\xe2\x02\x1bTkd\\Customer\\V1\\GPBMetadata\xea\x02\x11Tkd::Customer::V1b\x06proto3"

var (
	file_tkd_customer_v1_import_proto_rawDescOnce sync.Once
	file_tkd_customer_v1_import_proto_rawDescData []byte
)

func file_tkd_customer_v1_import_proto_rawDescGZIP() []byte {
	file_tkd_customer_v1_import_proto_rawDescOnce.Do(func() {
		file_tkd_customer_v1_import_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_tkd_customer_v1_import_proto_rawDesc), len(file_tkd_customer_v1_import_proto_rawDesc)))
	})
	return file_tkd_customer_v1_import_proto_rawDescData
}

var file_tkd_customer_v1_import_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_tkd_customer_v1_import_proto_goTypes = []any{
	(*StartSessionRequest)(nil),    // 0: tkd.customer.v1.StartSessionRequest
	(*StartSessionResponse)(nil),   // 1: tkd.customer.v1.StartSessionResponse
	(*LookupCustomerRequest)(nil),  // 2: tkd.customer.v1.LookupCustomerRequest
	(*ImportedCustomer)(nil),       // 3: tkd.customer.v1.ImportedCustomer
	(*LookupCustomerResponse)(nil), // 4: tkd.customer.v1.LookupCustomerResponse
	(*UpsertCustomerRequest)(nil),  // 5: tkd.customer.v1.UpsertCustomerRequest
	(*UpsertPatientRequest)(nil),   // 6: tkd.customer.v1.UpsertPatientRequest
	(*UpsertPatientSuccess)(nil),   // 7: tkd.customer.v1.UpsertPatientSuccess
	(*UpsertCustomerSuccess)(nil),  // 8: tkd.customer.v1.UpsertCustomerSuccess
	(*Error)(nil),                  // 9: tkd.customer.v1.Error
	(*ImportSessionComplete)(nil),  // 10: tkd.customer.v1.ImportSessionComplete
	(*ImportSessionRequest)(nil),   // 11: tkd.customer.v1.ImportSessionRequest
	(*ImportSessionResponse)(nil),  // 12: tkd.customer.v1.ImportSessionResponse
	(*CustomerQuery)(nil),          // 13: tkd.customer.v1.CustomerQuery
	(*Customer)(nil),               // 14: tkd.customer.v1.Customer
	(*ImportState)(nil),            // 15: tkd.customer.v1.ImportState
	(*structpb.Struct)(nil),        // 16: google.protobuf.Struct
	(*Patient)(nil),                // 17: tkd.customer.v1.Patient
}
var file_tkd_customer_v1_import_proto_depIdxs = []int32{
	13, // 0: tkd.customer.v1.LookupCustomerRequest.query:type_name -> tkd.customer.v1.CustomerQuery
	14, // 1: tkd.customer.v1.ImportedCustomer.customer:type_name -> tkd.customer.v1.Customer
	15, // 2: tkd.customer.v1.ImportedCustomer.state:type_name -> tkd.customer.v1.ImportState
	3,  // 3: tkd.customer.v1.LookupCustomerResponse.matched_customers:type_name -> tkd.customer.v1.ImportedCustomer
	16, // 4: tkd.customer.v1.UpsertCustomerRequest.extra_data:type_name -> google.protobuf.Struct
	14, // 5: tkd.customer.v1.UpsertCustomerRequest.customer:type_name -> tkd.customer.v1.Customer
	17, // 6: tkd.customer.v1.UpsertPatientRequest.patient:type_name -> tkd.customer.v1.Patient
	0,  // 7: tkd.customer.v1.ImportSessionRequest.start_session:type_name -> tkd.customer.v1.StartSessionRequest
	2,  // 8: tkd.customer.v1.ImportSessionRequest.lookup_customer:type_name -> tkd.customer.v1.LookupCustomerRequest
	5,  // 9: tkd.customer.v1.ImportSessionRequest.upsert_customer:type_name -> tkd.customer.v1.UpsertCustomerRequest
	6,  // 10: tkd.customer.v1.ImportSessionRequest.upsert_patient:type_name -> tkd.customer.v1.UpsertPatientRequest
	10, // 11: tkd.customer.v1.ImportSessionRequest.complete:type_name -> tkd.customer.v1.ImportSessionComplete
	1,  // 12: tkd.customer.v1.ImportSessionResponse.start_session:type_name -> tkd.customer.v1.StartSessionResponse
	4,  // 13: tkd.customer.v1.ImportSessionResponse.lookup_customer:type_name -> tkd.customer.v1.LookupCustomerResponse
	8,  // 14: tkd.customer.v1.ImportSessionResponse.upsert_customer_success:type_name -> tkd.customer.v1.UpsertCustomerSuccess
	7,  // 15: tkd.customer.v1.ImportSessionResponse.upsert_patient_success:type_name -> tkd.customer.v1.UpsertPatientSuccess
	9,  // 16: tkd.customer.v1.ImportSessionResponse.error:type_name -> tkd.customer.v1.Error
	11, // 17: tkd.customer.v1.CustomerImportService.ImportSession:input_type -> tkd.customer.v1.ImportSessionRequest
	12, // 18: tkd.customer.v1.CustomerImportService.ImportSession:output_type -> tkd.customer.v1.ImportSessionResponse
	18, // [18:19] is the sub-list for method output_type
	17, // [17:18] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_tkd_customer_v1_import_proto_init() }
func file_tkd_customer_v1_import_proto_init() {
	if File_tkd_customer_v1_import_proto != nil {
		return
	}
	file_tkd_customer_v1_customer_proto_init()
	file_tkd_customer_v1_patient_proto_init()
	file_tkd_customer_v1_import_proto_msgTypes[11].OneofWrappers = []any{
		(*ImportSessionRequest_StartSession)(nil),
		(*ImportSessionRequest_LookupCustomer)(nil),
		(*ImportSessionRequest_UpsertCustomer)(nil),
		(*ImportSessionRequest_UpsertPatient)(nil),
		(*ImportSessionRequest_Complete)(nil),
	}
	file_tkd_customer_v1_import_proto_msgTypes[12].OneofWrappers = []any{
		(*ImportSessionResponse_StartSession)(nil),
		(*ImportSessionResponse_LookupCustomer)(nil),
		(*ImportSessionResponse_UpsertCustomerSuccess)(nil),
		(*ImportSessionResponse_UpsertPatientSuccess)(nil),
		(*ImportSessionResponse_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_tkd_customer_v1_import_proto_rawDesc), len(file_tkd_customer_v1_import_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tkd_customer_v1_import_proto_goTypes,
		DependencyIndexes: file_tkd_customer_v1_import_proto_depIdxs,
		MessageInfos:      file_tkd_customer_v1_import_proto_msgTypes,
	}.Build()
	File_tkd_customer_v1_import_proto = out.File
	file_tkd_customer_v1_import_proto_goTypes = nil
	file_tkd_customer_v1_import_proto_depIdxs = nil
}
