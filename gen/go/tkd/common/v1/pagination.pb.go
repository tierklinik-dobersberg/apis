// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: tkd/common/v1/pagination.proto

package commonv1

import (
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

type SortDirection int32

const (
	SortDirection_SORT_DIRECTION_UNSPECIFIED SortDirection = 0
	SortDirection_SORT_DIRECTION_ASC         SortDirection = 1
	SortDirection_SORT_DIRECTION_DESC        SortDirection = 2
)

// Enum value maps for SortDirection.
var (
	SortDirection_name = map[int32]string{
		0: "SORT_DIRECTION_UNSPECIFIED",
		1: "SORT_DIRECTION_ASC",
		2: "SORT_DIRECTION_DESC",
	}
	SortDirection_value = map[string]int32{
		"SORT_DIRECTION_UNSPECIFIED": 0,
		"SORT_DIRECTION_ASC":         1,
		"SORT_DIRECTION_DESC":        2,
	}
)

func (x SortDirection) Enum() *SortDirection {
	p := new(SortDirection)
	*p = x
	return p
}

func (x SortDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_tkd_common_v1_pagination_proto_enumTypes[0].Descriptor()
}

func (SortDirection) Type() protoreflect.EnumType {
	return &file_tkd_common_v1_pagination_proto_enumTypes[0]
}

func (x SortDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortDirection.Descriptor instead.
func (SortDirection) EnumDescriptor() ([]byte, []int) {
	return file_tkd_common_v1_pagination_proto_rawDescGZIP(), []int{0}
}

type Sort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldName string        `protobuf:"bytes,1,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	Direction SortDirection `protobuf:"varint,2,opt,name=direction,proto3,enum=tkd.common.v1.SortDirection" json:"direction,omitempty"`
}

func (x *Sort) Reset() {
	*x = Sort{}
	mi := &file_tkd_common_v1_pagination_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Sort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sort) ProtoMessage() {}

func (x *Sort) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_common_v1_pagination_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sort.ProtoReflect.Descriptor instead.
func (*Sort) Descriptor() ([]byte, []int) {
	return file_tkd_common_v1_pagination_proto_rawDescGZIP(), []int{0}
}

func (x *Sort) GetFieldName() string {
	if x != nil {
		return x.FieldName
	}
	return ""
}

func (x *Sort) GetDirection() SortDirection {
	if x != nil {
		return x.Direction
	}
	return SortDirection_SORT_DIRECTION_UNSPECIFIED
}

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Kind:
	//
	//	*Pagination_Page
	//	*Pagination_NextPageToken
	Kind isPagination_Kind `protobuf_oneof:"kind"`
	// PageSize is the size of a page. If it is 0, the complete list
	// will be returned. Not that server implementation might enforce a
	// maximum page size!
	PageSize int32   `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	SortBy   []*Sort `protobuf:"bytes,4,rep,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	mi := &file_tkd_common_v1_pagination_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_common_v1_pagination_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_tkd_common_v1_pagination_proto_rawDescGZIP(), []int{1}
}

func (m *Pagination) GetKind() isPagination_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *Pagination) GetPage() int32 {
	if x, ok := x.GetKind().(*Pagination_Page); ok {
		return x.Page
	}
	return 0
}

func (x *Pagination) GetNextPageToken() string {
	if x, ok := x.GetKind().(*Pagination_NextPageToken); ok {
		return x.NextPageToken
	}
	return ""
}

func (x *Pagination) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *Pagination) GetSortBy() []*Sort {
	if x != nil {
		return x.SortBy
	}
	return nil
}

type isPagination_Kind interface {
	isPagination_Kind()
}

type Pagination_Page struct {
	// Page is the page number to retrieve.
	// Page numbers start at 0!
	Page int32 `protobuf:"varint,1,opt,name=page,proto3,oneof"`
}

type Pagination_NextPageToken struct {
	// Not yet supported!
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3,oneof"`
}

func (*Pagination_Page) isPagination_Kind() {}

func (*Pagination_NextPageToken) isPagination_Kind() {}

var File_tkd_common_v1_pagination_proto protoreflect.FileDescriptor

var file_tkd_common_v1_pagination_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x74, 0x6b, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x74, 0x6b, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22,
	0x61, 0x0a, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x74, 0x6b, 0x64, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x9f, 0x01, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x00, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2c,
	0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x6f, 0x72, 0x74, 0x52, 0x06, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x79, 0x42, 0x06, 0x0a, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x2a, 0x60, 0x0a, 0x0d, 0x53, 0x6f, 0x72, 0x74, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x44, 0x49,
	0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x44, 0x49,
	0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x53, 0x43, 0x10, 0x01, 0x12, 0x17, 0x0a,
	0x13, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x44, 0x45, 0x53, 0x43, 0x10, 0x02, 0x42, 0xbf, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x74,
	0x6b, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x65, 0x72,
	0x6b, 0x6c, 0x69, 0x6e, 0x69, 0x6b, 0x2d, 0x64, 0x6f, 0x62, 0x65, 0x72, 0x73, 0x62, 0x65, 0x72,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x6b,
	0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x43, 0x58, 0xaa, 0x02, 0x0d, 0x54, 0x6b, 0x64,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x54, 0x6b, 0x64,
	0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x54, 0x6b, 0x64,
	0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x54, 0x6b, 0x64, 0x3a, 0x3a, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tkd_common_v1_pagination_proto_rawDescOnce sync.Once
	file_tkd_common_v1_pagination_proto_rawDescData = file_tkd_common_v1_pagination_proto_rawDesc
)

func file_tkd_common_v1_pagination_proto_rawDescGZIP() []byte {
	file_tkd_common_v1_pagination_proto_rawDescOnce.Do(func() {
		file_tkd_common_v1_pagination_proto_rawDescData = protoimpl.X.CompressGZIP(file_tkd_common_v1_pagination_proto_rawDescData)
	})
	return file_tkd_common_v1_pagination_proto_rawDescData
}

var file_tkd_common_v1_pagination_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tkd_common_v1_pagination_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tkd_common_v1_pagination_proto_goTypes = []any{
	(SortDirection)(0), // 0: tkd.common.v1.SortDirection
	(*Sort)(nil),       // 1: tkd.common.v1.Sort
	(*Pagination)(nil), // 2: tkd.common.v1.Pagination
}
var file_tkd_common_v1_pagination_proto_depIdxs = []int32{
	0, // 0: tkd.common.v1.Sort.direction:type_name -> tkd.common.v1.SortDirection
	1, // 1: tkd.common.v1.Pagination.sort_by:type_name -> tkd.common.v1.Sort
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tkd_common_v1_pagination_proto_init() }
func file_tkd_common_v1_pagination_proto_init() {
	if File_tkd_common_v1_pagination_proto != nil {
		return
	}
	file_tkd_common_v1_pagination_proto_msgTypes[1].OneofWrappers = []any{
		(*Pagination_Page)(nil),
		(*Pagination_NextPageToken)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tkd_common_v1_pagination_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tkd_common_v1_pagination_proto_goTypes,
		DependencyIndexes: file_tkd_common_v1_pagination_proto_depIdxs,
		EnumInfos:         file_tkd_common_v1_pagination_proto_enumTypes,
		MessageInfos:      file_tkd_common_v1_pagination_proto_msgTypes,
	}.Build()
	File_tkd_common_v1_pagination_proto = out.File
	file_tkd_common_v1_pagination_proto_rawDesc = nil
	file_tkd_common_v1_pagination_proto_goTypes = nil
	file_tkd_common_v1_pagination_proto_depIdxs = nil
}
