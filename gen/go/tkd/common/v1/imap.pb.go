// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: tkd/common/v1/imap.proto

package commonv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IMAPConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Host is the IMAP host to connect to
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// TLS enables or disables TLS support.
	Tls bool `protobuf:"varint,2,opt,name=tls,proto3" json:"tls,omitempty"`
	// InsecureSkipVerify disables certificate validation when TLS is used.
	InsecureSkipVerify bool `protobuf:"varint,3,opt,name=insecure_skip_verify,json=insecureSkipVerify,proto3" json:"insecure_skip_verify,omitempty"`
	// User is the username to authenticate.
	User string `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	// Password is the password to authenticate.
	Password string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	// Folder is the IMAP folder to watch. Defaults to INBOX
	Folder string `protobuf:"bytes,6,opt,name=folder,proto3" json:"folder,omitempty"`
	// ReadOnly can be set to true to open the mailbox in read-only mode.
	ReadOnly bool `protobuf:"varint,7,opt,name=read_only,json=readOnly,proto3" json:"read_only,omitempty"`
}

func (x *IMAPConfig) Reset() {
	*x = IMAPConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tkd_common_v1_imap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IMAPConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IMAPConfig) ProtoMessage() {}

func (x *IMAPConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_common_v1_imap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IMAPConfig.ProtoReflect.Descriptor instead.
func (*IMAPConfig) Descriptor() ([]byte, []int) {
	return file_tkd_common_v1_imap_proto_rawDescGZIP(), []int{0}
}

func (x *IMAPConfig) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *IMAPConfig) GetTls() bool {
	if x != nil {
		return x.Tls
	}
	return false
}

func (x *IMAPConfig) GetInsecureSkipVerify() bool {
	if x != nil {
		return x.InsecureSkipVerify
	}
	return false
}

func (x *IMAPConfig) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *IMAPConfig) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *IMAPConfig) GetFolder() string {
	if x != nil {
		return x.Folder
	}
	return ""
}

func (x *IMAPConfig) GetReadOnly() bool {
	if x != nil {
		return x.ReadOnly
	}
	return false
}

var File_tkd_common_v1_imap_proto protoreflect.FileDescriptor

var file_tkd_common_v1_imap_proto_rawDesc = []byte{
	0x0a, 0x18, 0x74, 0x6b, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x69, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x74, 0x6b, 0x64, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x74, 0x6b, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x02, 0x0a, 0x0a, 0x49, 0x4d, 0x41, 0x50, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0xf7, 0x18, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x03, 0x74, 0x6c, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65,
	0x5f, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x12, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x53, 0x6b, 0x69, 0x70,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x1b,
	0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x3a, 0x52, 0xd2, 0x7e, 0x4f,
	0x0a, 0x3a, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x0a, 0x03, 0x74, 0x6c, 0x73, 0x0a, 0x09, 0x72,
	0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x0a, 0x14,
	0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x5f, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x0a, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x02, 0x69, 0x64,
	0x1a, 0x0d, 0x69, 0x64, 0x6d, 0x5f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x75, 0x73, 0x65, 0x72, 0x42,
	0xb9, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x49, 0x6d, 0x61, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x69, 0x65, 0x72, 0x6b, 0x6c, 0x69, 0x6e, 0x69, 0x6b, 0x2d, 0x64, 0x6f, 0x62, 0x65, 0x72, 0x73,
	0x62, 0x65, 0x72, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x74, 0x6b, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x43, 0x58, 0xaa, 0x02, 0x0d,
	0x54, 0x6b, 0x64, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d,
	0x54, 0x6b, 0x64, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19,
	0x54, 0x6b, 0x64, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x54, 0x6b, 0x64, 0x3a,
	0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_tkd_common_v1_imap_proto_rawDescOnce sync.Once
	file_tkd_common_v1_imap_proto_rawDescData = file_tkd_common_v1_imap_proto_rawDesc
)

func file_tkd_common_v1_imap_proto_rawDescGZIP() []byte {
	file_tkd_common_v1_imap_proto_rawDescOnce.Do(func() {
		file_tkd_common_v1_imap_proto_rawDescData = protoimpl.X.CompressGZIP(file_tkd_common_v1_imap_proto_rawDescData)
	})
	return file_tkd_common_v1_imap_proto_rawDescData
}

var file_tkd_common_v1_imap_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tkd_common_v1_imap_proto_goTypes = []any{
	(*IMAPConfig)(nil), // 0: tkd.common.v1.IMAPConfig
}
var file_tkd_common_v1_imap_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tkd_common_v1_imap_proto_init() }
func file_tkd_common_v1_imap_proto_init() {
	if File_tkd_common_v1_imap_proto != nil {
		return
	}
	file_tkd_common_v1_descriptor_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tkd_common_v1_imap_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*IMAPConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tkd_common_v1_imap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tkd_common_v1_imap_proto_goTypes,
		DependencyIndexes: file_tkd_common_v1_imap_proto_depIdxs,
		MessageInfos:      file_tkd_common_v1_imap_proto_msgTypes,
	}.Build()
	File_tkd_common_v1_imap_proto = out.File
	file_tkd_common_v1_imap_proto_rawDesc = nil
	file_tkd_common_v1_imap_proto_goTypes = nil
	file_tkd_common_v1_imap_proto_depIdxs = nil
}
