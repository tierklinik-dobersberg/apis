// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: tkd/booking/v1/species.proto

package bookingv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Species is a well-known species for which treatments are available.
type Species struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name is the unique name of the species.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// DisplayName is a human readable name for the species to be used in the
	// self-booking user interface.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// RequestCastrationStatus can be set to true if the self-booking user interface
	// should request the castration status of the patient.
	RequestCastrationStatus bool `protobuf:"varint,3,opt,name=request_castration_status,json=requestCastrationStatus,proto3" json:"request_castration_status,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *Species) Reset() {
	*x = Species{}
	mi := &file_tkd_booking_v1_species_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Species) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Species) ProtoMessage() {}

func (x *Species) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_booking_v1_species_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Species.ProtoReflect.Descriptor instead.
func (*Species) Descriptor() ([]byte, []int) {
	return file_tkd_booking_v1_species_proto_rawDescGZIP(), []int{0}
}

func (x *Species) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Species) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Species) GetRequestCastrationStatus() bool {
	if x != nil {
		return x.RequestCastrationStatus
	}
	return false
}

type ListSpeciesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Search        string                 `protobuf:"bytes,1,opt,name=search,proto3" json:"search,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListSpeciesRequest) Reset() {
	*x = ListSpeciesRequest{}
	mi := &file_tkd_booking_v1_species_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSpeciesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSpeciesRequest) ProtoMessage() {}

func (x *ListSpeciesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_booking_v1_species_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSpeciesRequest.ProtoReflect.Descriptor instead.
func (*ListSpeciesRequest) Descriptor() ([]byte, []int) {
	return file_tkd_booking_v1_species_proto_rawDescGZIP(), []int{1}
}

func (x *ListSpeciesRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type ListSpeciesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Species       []*Species             `protobuf:"bytes,1,rep,name=species,proto3" json:"species,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListSpeciesResponse) Reset() {
	*x = ListSpeciesResponse{}
	mi := &file_tkd_booking_v1_species_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSpeciesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSpeciesResponse) ProtoMessage() {}

func (x *ListSpeciesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_booking_v1_species_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSpeciesResponse.ProtoReflect.Descriptor instead.
func (*ListSpeciesResponse) Descriptor() ([]byte, []int) {
	return file_tkd_booking_v1_species_proto_rawDescGZIP(), []int{2}
}

func (x *ListSpeciesResponse) GetSpecies() []*Species {
	if x != nil {
		return x.Species
	}
	return nil
}

type UpdateSpeciesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Species       *Species               `protobuf:"bytes,2,opt,name=species,proto3" json:"species,omitempty"`
	UpdateMask    *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateSpeciesRequest) Reset() {
	*x = UpdateSpeciesRequest{}
	mi := &file_tkd_booking_v1_species_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateSpeciesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSpeciesRequest) ProtoMessage() {}

func (x *UpdateSpeciesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_booking_v1_species_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSpeciesRequest.ProtoReflect.Descriptor instead.
func (*UpdateSpeciesRequest) Descriptor() ([]byte, []int) {
	return file_tkd_booking_v1_species_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateSpeciesRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateSpeciesRequest) GetSpecies() *Species {
	if x != nil {
		return x.Species
	}
	return nil
}

func (x *UpdateSpeciesRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type DeleteSpeciesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteSpeciesRequest) Reset() {
	*x = DeleteSpeciesRequest{}
	mi := &file_tkd_booking_v1_species_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteSpeciesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSpeciesRequest) ProtoMessage() {}

func (x *DeleteSpeciesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_booking_v1_species_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSpeciesRequest.ProtoReflect.Descriptor instead.
func (*DeleteSpeciesRequest) Descriptor() ([]byte, []int) {
	return file_tkd_booking_v1_species_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteSpeciesRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_tkd_booking_v1_species_proto protoreflect.FileDescriptor

var file_tkd_booking_v1_species_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x74, 0x6b, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x74, 0x6b, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62,
	0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x07, 0x53, 0x70,
	0x65, 0x63, 0x69, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x19,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x61, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x17, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2c, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x48, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x70,
	0x65, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a,
	0x07, 0x73, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x52, 0x07, 0x73, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73,
	0x22, 0xac, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0xf7, 0x18, 0x03, 0xc8, 0x01, 0x01,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x63, 0x69, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73,
	0x42, 0x07, 0xfa, 0xf7, 0x18, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x07, 0x73, 0x70, 0x65, 0x63, 0x69,
	0x65, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22,
	0x33, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0xf7, 0x18, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x32, 0xd2, 0x02, 0x0a, 0x0e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x12, 0x17, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x65,
	0x73, 0x1a, 0x17, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0b,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x12, 0x22, 0x2e, 0x74, 0x6b,
	0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x12, 0x24, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x74, 0x6b, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x12, 0x24, 0x2e, 0x74, 0x6b, 0x64, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0xc3, 0x01, 0x0a, 0x12, 0x63, 0x6f,
	0x6d, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x42, 0x0c, 0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x65,
	0x72, 0x6b, 0x6c, 0x69, 0x6e, 0x69, 0x6b, 0x2d, 0x64, 0x6f, 0x62, 0x65, 0x72, 0x73, 0x62, 0x65,
	0x72, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74,
	0x6b, 0x64, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x42, 0x58, 0xaa, 0x02, 0x0e,
	0x54, 0x6b, 0x64, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x0e, 0x54, 0x6b, 0x64, 0x5c, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x1a, 0x54, 0x6b, 0x64, 0x5c, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x54,
	0x6b, 0x64, 0x3a, 0x3a, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tkd_booking_v1_species_proto_rawDescOnce sync.Once
	file_tkd_booking_v1_species_proto_rawDescData = file_tkd_booking_v1_species_proto_rawDesc
)

func file_tkd_booking_v1_species_proto_rawDescGZIP() []byte {
	file_tkd_booking_v1_species_proto_rawDescOnce.Do(func() {
		file_tkd_booking_v1_species_proto_rawDescData = protoimpl.X.CompressGZIP(file_tkd_booking_v1_species_proto_rawDescData)
	})
	return file_tkd_booking_v1_species_proto_rawDescData
}

var file_tkd_booking_v1_species_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_tkd_booking_v1_species_proto_goTypes = []any{
	(*Species)(nil),               // 0: tkd.booking.v1.Species
	(*ListSpeciesRequest)(nil),    // 1: tkd.booking.v1.ListSpeciesRequest
	(*ListSpeciesResponse)(nil),   // 2: tkd.booking.v1.ListSpeciesResponse
	(*UpdateSpeciesRequest)(nil),  // 3: tkd.booking.v1.UpdateSpeciesRequest
	(*DeleteSpeciesRequest)(nil),  // 4: tkd.booking.v1.DeleteSpeciesRequest
	(*fieldmaskpb.FieldMask)(nil), // 5: google.protobuf.FieldMask
	(*emptypb.Empty)(nil),         // 6: google.protobuf.Empty
}
var file_tkd_booking_v1_species_proto_depIdxs = []int32{
	0, // 0: tkd.booking.v1.ListSpeciesResponse.species:type_name -> tkd.booking.v1.Species
	0, // 1: tkd.booking.v1.UpdateSpeciesRequest.species:type_name -> tkd.booking.v1.Species
	5, // 2: tkd.booking.v1.UpdateSpeciesRequest.update_mask:type_name -> google.protobuf.FieldMask
	0, // 3: tkd.booking.v1.SpeciesService.CreateSpecies:input_type -> tkd.booking.v1.Species
	1, // 4: tkd.booking.v1.SpeciesService.ListSpecies:input_type -> tkd.booking.v1.ListSpeciesRequest
	3, // 5: tkd.booking.v1.SpeciesService.UpdateSpecies:input_type -> tkd.booking.v1.UpdateSpeciesRequest
	4, // 6: tkd.booking.v1.SpeciesService.DeleteSpecies:input_type -> tkd.booking.v1.DeleteSpeciesRequest
	0, // 7: tkd.booking.v1.SpeciesService.CreateSpecies:output_type -> tkd.booking.v1.Species
	2, // 8: tkd.booking.v1.SpeciesService.ListSpecies:output_type -> tkd.booking.v1.ListSpeciesResponse
	0, // 9: tkd.booking.v1.SpeciesService.UpdateSpecies:output_type -> tkd.booking.v1.Species
	6, // 10: tkd.booking.v1.SpeciesService.DeleteSpecies:output_type -> google.protobuf.Empty
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tkd_booking_v1_species_proto_init() }
func file_tkd_booking_v1_species_proto_init() {
	if File_tkd_booking_v1_species_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tkd_booking_v1_species_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tkd_booking_v1_species_proto_goTypes,
		DependencyIndexes: file_tkd_booking_v1_species_proto_depIdxs,
		MessageInfos:      file_tkd_booking_v1_species_proto_msgTypes,
	}.Build()
	File_tkd_booking_v1_species_proto = out.File
	file_tkd_booking_v1_species_proto_rawDesc = nil
	file_tkd_booking_v1_species_proto_goTypes = nil
	file_tkd_booking_v1_species_proto_depIdxs = nil
}
