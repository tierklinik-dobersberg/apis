// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
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
	unsafe "unsafe"
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

const file_tkd_booking_v1_species_proto_rawDesc = "" +
	"\n" +
	"\x1ctkd/booking/v1/species.proto\x12\x0etkd.booking.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a google/protobuf/field_mask.proto\x1a\x1bbuf/validate/validate.proto\"|\n" +
	"\aSpecies\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12!\n" +
	"\fdisplay_name\x18\x02 \x01(\tR\vdisplayName\x12:\n" +
	"\x19request_castration_status\x18\x03 \x01(\bR\x17requestCastrationStatus\",\n" +
	"\x12ListSpeciesRequest\x12\x16\n" +
	"\x06search\x18\x01 \x01(\tR\x06search\"H\n" +
	"\x13ListSpeciesResponse\x121\n" +
	"\aspecies\x18\x01 \x03(\v2\x17.tkd.booking.v1.SpeciesR\aspecies\"\xac\x01\n" +
	"\x14UpdateSpeciesRequest\x12\x1b\n" +
	"\x04name\x18\x01 \x01(\tB\a\xfa\xf7\x18\x03\xc8\x01\x01R\x04name\x12:\n" +
	"\aspecies\x18\x02 \x01(\v2\x17.tkd.booking.v1.SpeciesB\a\xfa\xf7\x18\x03\xc8\x01\x01R\aspecies\x12;\n" +
	"\vupdate_mask\x18\x03 \x01(\v2\x1a.google.protobuf.FieldMaskR\n" +
	"updateMask\"3\n" +
	"\x14DeleteSpeciesRequest\x12\x1b\n" +
	"\x04name\x18\x01 \x01(\tB\a\xfa\xf7\x18\x03\xc8\x01\x01R\x04name2\xd2\x02\n" +
	"\x0eSpeciesService\x12C\n" +
	"\rCreateSpecies\x12\x17.tkd.booking.v1.Species\x1a\x17.tkd.booking.v1.Species\"\x00\x12X\n" +
	"\vListSpecies\x12\".tkd.booking.v1.ListSpeciesRequest\x1a#.tkd.booking.v1.ListSpeciesResponse\"\x00\x12P\n" +
	"\rUpdateSpecies\x12$.tkd.booking.v1.UpdateSpeciesRequest\x1a\x17.tkd.booking.v1.Species\"\x00\x12O\n" +
	"\rDeleteSpecies\x12$.tkd.booking.v1.DeleteSpeciesRequest\x1a\x16.google.protobuf.Empty\"\x00B\xc3\x01\n" +
	"\x12com.tkd.booking.v1B\fSpeciesProtoP\x01ZEgithub.com/tierklinik-dobersberg/apis/gen/go/tkd/booking/v1;bookingv1\xa2\x02\x03TBX\xaa\x02\x0eTkd.Booking.V1\xca\x02\x0eTkd\\Booking\\V1\xe2\x02\x1aTkd\\Booking\\V1\\GPBMetadata\xea\x02\x10Tkd::Booking::V1b\x06proto3"

var (
	file_tkd_booking_v1_species_proto_rawDescOnce sync.Once
	file_tkd_booking_v1_species_proto_rawDescData []byte
)

func file_tkd_booking_v1_species_proto_rawDescGZIP() []byte {
	file_tkd_booking_v1_species_proto_rawDescOnce.Do(func() {
		file_tkd_booking_v1_species_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_tkd_booking_v1_species_proto_rawDesc), len(file_tkd_booking_v1_species_proto_rawDesc)))
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_tkd_booking_v1_species_proto_rawDesc), len(file_tkd_booking_v1_species_proto_rawDesc)),
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
	file_tkd_booking_v1_species_proto_goTypes = nil
	file_tkd_booking_v1_species_proto_depIdxs = nil
}
