// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: tkd/common/v1/date.proto

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

type Month int32

const (
	Month_MONTH_UNSPECIFIED Month = 0
	Month_January           Month = 1
	Month_Feburary          Month = 2
	Month_March             Month = 3
	Month_April             Month = 4
	Month_May               Month = 5
	Month_June              Month = 6
	Month_July              Month = 7
	Month_August            Month = 8
	Month_September         Month = 9
	Month_October           Month = 10
	Month_November          Month = 11
	Month_December          Month = 12
)

// Enum value maps for Month.
var (
	Month_name = map[int32]string{
		0:  "MONTH_UNSPECIFIED",
		1:  "January",
		2:  "Feburary",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}
	Month_value = map[string]int32{
		"MONTH_UNSPECIFIED": 0,
		"January":           1,
		"Feburary":          2,
		"March":             3,
		"April":             4,
		"May":               5,
		"June":              6,
		"July":              7,
		"August":            8,
		"September":         9,
		"October":           10,
		"November":          11,
		"December":          12,
	}
)

func (x Month) Enum() *Month {
	p := new(Month)
	*p = x
	return p
}

func (x Month) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Month) Descriptor() protoreflect.EnumDescriptor {
	return file_tkd_common_v1_date_proto_enumTypes[0].Descriptor()
}

func (Month) Type() protoreflect.EnumType {
	return &file_tkd_common_v1_date_proto_enumTypes[0]
}

func (x Month) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Month.Descriptor instead.
func (Month) EnumDescriptor() ([]byte, []int) {
	return file_tkd_common_v1_date_proto_rawDescGZIP(), []int{0}
}

type Date struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Year holds the year of the date.
	Year int64 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	// Month holds the month of the date.
	Month Month `protobuf:"varint,2,opt,name=month,proto3,enum=tkd.common.v1.Month" json:"month,omitempty"`
	// Day holds the day-of-month.
	Day int32 `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
	// Timezone optionally holds the timezone.
	// If unset, the server will default to the local
	// timezone.
	Timezone      string `protobuf:"bytes,4,opt,name=timezone,proto3" json:"timezone,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Date) Reset() {
	*x = Date{}
	mi := &file_tkd_common_v1_date_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Date) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Date) ProtoMessage() {}

func (x *Date) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_common_v1_date_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Date.ProtoReflect.Descriptor instead.
func (*Date) Descriptor() ([]byte, []int) {
	return file_tkd_common_v1_date_proto_rawDescGZIP(), []int{0}
}

func (x *Date) GetYear() int64 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Date) GetMonth() Month {
	if x != nil {
		return x.Month
	}
	return Month_MONTH_UNSPECIFIED
}

func (x *Date) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *Date) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

type DateRange struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	From          *Date                  `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To            *Date                  `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DateRange) Reset() {
	*x = DateRange{}
	mi := &file_tkd_common_v1_date_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DateRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateRange) ProtoMessage() {}

func (x *DateRange) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_common_v1_date_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DateRange.ProtoReflect.Descriptor instead.
func (*DateRange) Descriptor() ([]byte, []int) {
	return file_tkd_common_v1_date_proto_rawDescGZIP(), []int{1}
}

func (x *DateRange) GetFrom() *Date {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *DateRange) GetTo() *Date {
	if x != nil {
		return x.To
	}
	return nil
}

var File_tkd_common_v1_date_proto protoreflect.FileDescriptor

var file_tkd_common_v1_date_proto_rawDesc = []byte{
	0x0a, 0x18, 0x74, 0x6b, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x74, 0x6b, 0x64, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0x74, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x2a, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74,
	0x68, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x64, 0x61, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x22,
	0x59, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x6b, 0x64,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x23, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x02, 0x74, 0x6f, 0x2a, 0xb0, 0x01, 0x0a, 0x05, 0x4d,
	0x6f, 0x6e, 0x74, 0x68, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4a,
	0x61, 0x6e, 0x75, 0x61, 0x72, 0x79, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x65, 0x62, 0x75,
	0x72, 0x61, 0x72, 0x79, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x61, 0x72, 0x63, 0x68, 0x10,
	0x03, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x70, 0x72, 0x69, 0x6c, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03,
	0x4d, 0x61, 0x79, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x75, 0x6e, 0x65, 0x10, 0x06, 0x12,
	0x08, 0x0a, 0x04, 0x4a, 0x75, 0x6c, 0x79, 0x10, 0x07, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x75, 0x67,
	0x75, 0x73, 0x74, 0x10, 0x08, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x65, 0x70, 0x74, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x10, 0x09, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x63, 0x74, 0x6f, 0x62, 0x65, 0x72, 0x10,
	0x0a, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x6f, 0x76, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x10, 0x0b, 0x12,
	0x0c, 0x0a, 0x08, 0x44, 0x65, 0x63, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x10, 0x0c, 0x42, 0xb9, 0x01,
	0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x42, 0x09, 0x44, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x65,
	0x72, 0x6b, 0x6c, 0x69, 0x6e, 0x69, 0x6b, 0x2d, 0x64, 0x6f, 0x62, 0x65, 0x72, 0x73, 0x62, 0x65,
	0x72, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74,
	0x6b, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x43, 0x58, 0xaa, 0x02, 0x0d, 0x54, 0x6b,
	0x64, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x54, 0x6b,
	0x64, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x54, 0x6b,
	0x64, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x54, 0x6b, 0x64, 0x3a, 0x3a, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_tkd_common_v1_date_proto_rawDescOnce sync.Once
	file_tkd_common_v1_date_proto_rawDescData = file_tkd_common_v1_date_proto_rawDesc
)

func file_tkd_common_v1_date_proto_rawDescGZIP() []byte {
	file_tkd_common_v1_date_proto_rawDescOnce.Do(func() {
		file_tkd_common_v1_date_proto_rawDescData = protoimpl.X.CompressGZIP(file_tkd_common_v1_date_proto_rawDescData)
	})
	return file_tkd_common_v1_date_proto_rawDescData
}

var file_tkd_common_v1_date_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tkd_common_v1_date_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tkd_common_v1_date_proto_goTypes = []any{
	(Month)(0),        // 0: tkd.common.v1.Month
	(*Date)(nil),      // 1: tkd.common.v1.Date
	(*DateRange)(nil), // 2: tkd.common.v1.DateRange
}
var file_tkd_common_v1_date_proto_depIdxs = []int32{
	0, // 0: tkd.common.v1.Date.month:type_name -> tkd.common.v1.Month
	1, // 1: tkd.common.v1.DateRange.from:type_name -> tkd.common.v1.Date
	1, // 2: tkd.common.v1.DateRange.to:type_name -> tkd.common.v1.Date
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tkd_common_v1_date_proto_init() }
func file_tkd_common_v1_date_proto_init() {
	if File_tkd_common_v1_date_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tkd_common_v1_date_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tkd_common_v1_date_proto_goTypes,
		DependencyIndexes: file_tkd_common_v1_date_proto_depIdxs,
		EnumInfos:         file_tkd_common_v1_date_proto_enumTypes,
		MessageInfos:      file_tkd_common_v1_date_proto_msgTypes,
	}.Build()
	File_tkd_common_v1_date_proto = out.File
	file_tkd_common_v1_date_proto_rawDesc = nil
	file_tkd_common_v1_date_proto_goTypes = nil
	file_tkd_common_v1_date_proto_depIdxs = nil
}
