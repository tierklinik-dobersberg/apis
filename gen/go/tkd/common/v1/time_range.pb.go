// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: tkd/common/v1/time_range.proto

package commonv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// TimeRange describes a time range with a start and end-time.
// If both, start and end time is unset, the time range is not valid
// and no times will match.
type TimeRange struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// From holds the (inclusive) start time of the timerange.
	// If from is unspecified, the time-range has an open-start.
	From *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	// To holds the (inclusive) end time of the timerange.
	// If unspecified, the time range has an open end.
	To            *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TimeRange) Reset() {
	*x = TimeRange{}
	mi := &file_tkd_common_v1_time_range_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimeRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeRange) ProtoMessage() {}

func (x *TimeRange) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_common_v1_time_range_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeRange.ProtoReflect.Descriptor instead.
func (*TimeRange) Descriptor() ([]byte, []int) {
	return file_tkd_common_v1_time_range_proto_rawDescGZIP(), []int{0}
}

func (x *TimeRange) GetFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *TimeRange) GetTo() *timestamppb.Timestamp {
	if x != nil {
		return x.To
	}
	return nil
}

var File_tkd_common_v1_time_range_proto protoreflect.FileDescriptor

const file_tkd_common_v1_time_range_proto_rawDesc = "" +
	"\n" +
	"\x1etkd/common/v1/time_range.proto\x12\rtkd.common.v1\x1a\x1fgoogle/protobuf/timestamp.proto\"g\n" +
	"\tTimeRange\x12.\n" +
	"\x04from\x18\x01 \x01(\v2\x1a.google.protobuf.TimestampR\x04from\x12*\n" +
	"\x02to\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\x02toB\xbe\x01\n" +
	"\x11com.tkd.common.v1B\x0eTimeRangeProtoP\x01ZCgithub.com/tierklinik-dobersberg/apis/gen/go/tkd/common/v1;commonv1\xa2\x02\x03TCX\xaa\x02\rTkd.Common.V1\xca\x02\rTkd\\Common\\V1\xe2\x02\x19Tkd\\Common\\V1\\GPBMetadata\xea\x02\x0fTkd::Common::V1b\x06proto3"

var (
	file_tkd_common_v1_time_range_proto_rawDescOnce sync.Once
	file_tkd_common_v1_time_range_proto_rawDescData []byte
)

func file_tkd_common_v1_time_range_proto_rawDescGZIP() []byte {
	file_tkd_common_v1_time_range_proto_rawDescOnce.Do(func() {
		file_tkd_common_v1_time_range_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_tkd_common_v1_time_range_proto_rawDesc), len(file_tkd_common_v1_time_range_proto_rawDesc)))
	})
	return file_tkd_common_v1_time_range_proto_rawDescData
}

var file_tkd_common_v1_time_range_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tkd_common_v1_time_range_proto_goTypes = []any{
	(*TimeRange)(nil),             // 0: tkd.common.v1.TimeRange
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_tkd_common_v1_time_range_proto_depIdxs = []int32{
	1, // 0: tkd.common.v1.TimeRange.from:type_name -> google.protobuf.Timestamp
	1, // 1: tkd.common.v1.TimeRange.to:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tkd_common_v1_time_range_proto_init() }
func file_tkd_common_v1_time_range_proto_init() {
	if File_tkd_common_v1_time_range_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_tkd_common_v1_time_range_proto_rawDesc), len(file_tkd_common_v1_time_range_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tkd_common_v1_time_range_proto_goTypes,
		DependencyIndexes: file_tkd_common_v1_time_range_proto_depIdxs,
		MessageInfos:      file_tkd_common_v1_time_range_proto_msgTypes,
	}.Build()
	File_tkd_common_v1_time_range_proto = out.File
	file_tkd_common_v1_time_range_proto_goTypes = nil
	file_tkd_common_v1_time_range_proto_depIdxs = nil
}
