// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: tkd/userd/v1/descriptor.proto

package userdv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_tkd_userd_v1_descriptor_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         2021,
		Name:          "tkd.userd.v1.allow_self_service",
		Tag:           "varint,2021,opt,name=allow_self_service",
		Filename:      "tkd/userd/v1/descriptor.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional bool allow_self_service = 2021;
	E_AllowSelfService = &file_tkd_userd_v1_descriptor_proto_extTypes[0]
)

var File_tkd_userd_v1_descriptor_proto protoreflect.FileDescriptor

var file_tkd_userd_v1_descriptor_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x74, 0x6b, 0x64, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x74, 0x6b, 0x64, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a,
	0x4c, 0x0a, 0x12, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x6c, 0x66, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe5, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x53, 0x65, 0x6c, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0xac, 0x01,
	0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x64, 0x2e,
	0x76, 0x31, 0x42, 0x0f, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x62, 0x75, 0x66, 0x2d, 0x74,
	0x6f, 0x75, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x6b, 0x64, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x64, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x64, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54,
	0x55, 0x58, 0xaa, 0x02, 0x0c, 0x54, 0x6b, 0x64, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x64, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x0c, 0x54, 0x6b, 0x64, 0x5c, 0x55, 0x73, 0x65, 0x72, 0x64, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x18, 0x54, 0x6b, 0x64, 0x5c, 0x55, 0x73, 0x65, 0x72, 0x64, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x54, 0x6b,
	0x64, 0x3a, 0x3a, 0x55, 0x73, 0x65, 0x72, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_tkd_userd_v1_descriptor_proto_goTypes = []interface{}{
	(*descriptorpb.FieldOptions)(nil), // 0: google.protobuf.FieldOptions
}
var file_tkd_userd_v1_descriptor_proto_depIdxs = []int32{
	0, // 0: tkd.userd.v1.allow_self_service:extendee -> google.protobuf.FieldOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tkd_userd_v1_descriptor_proto_init() }
func file_tkd_userd_v1_descriptor_proto_init() {
	if File_tkd_userd_v1_descriptor_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tkd_userd_v1_descriptor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_tkd_userd_v1_descriptor_proto_goTypes,
		DependencyIndexes: file_tkd_userd_v1_descriptor_proto_depIdxs,
		ExtensionInfos:    file_tkd_userd_v1_descriptor_proto_extTypes,
	}.Build()
	File_tkd_userd_v1_descriptor_proto = out.File
	file_tkd_userd_v1_descriptor_proto_rawDesc = nil
	file_tkd_userd_v1_descriptor_proto_goTypes = nil
	file_tkd_userd_v1_descriptor_proto_depIdxs = nil
}
