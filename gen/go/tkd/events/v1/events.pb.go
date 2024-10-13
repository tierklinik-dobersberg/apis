// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: tkd/events/v1/events.proto

package eventsv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event    *anypb.Any `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Retained bool       `protobuf:"varint,2,opt,name=retained,proto3" json:"retained,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	mi := &file_tkd_events_v1_events_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_events_v1_events_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_tkd_events_v1_events_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetEvent() *anypb.Any {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *Event) GetRetained() bool {
	if x != nil {
		return x.Retained
	}
	return false
}

type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Kind:
	//
	//	*SubscribeRequest_Subscribe
	//	*SubscribeRequest_Unsubscribe
	Kind isSubscribeRequest_Kind `protobuf_oneof:"kind"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	mi := &file_tkd_events_v1_events_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_events_v1_events_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_tkd_events_v1_events_proto_rawDescGZIP(), []int{1}
}

func (m *SubscribeRequest) GetKind() isSubscribeRequest_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *SubscribeRequest) GetSubscribe() string {
	if x, ok := x.GetKind().(*SubscribeRequest_Subscribe); ok {
		return x.Subscribe
	}
	return ""
}

func (x *SubscribeRequest) GetUnsubscribe() string {
	if x, ok := x.GetKind().(*SubscribeRequest_Unsubscribe); ok {
		return x.Unsubscribe
	}
	return ""
}

type isSubscribeRequest_Kind interface {
	isSubscribeRequest_Kind()
}

type SubscribeRequest_Subscribe struct {
	Subscribe string `protobuf:"bytes,1,opt,name=subscribe,proto3,oneof"`
}

type SubscribeRequest_Unsubscribe struct {
	Unsubscribe string `protobuf:"bytes,2,opt,name=unsubscribe,proto3,oneof"`
}

func (*SubscribeRequest_Subscribe) isSubscribeRequest_Kind() {}

func (*SubscribeRequest_Unsubscribe) isSubscribeRequest_Kind() {}

type SubscribeOnceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeUrls []string `protobuf:"bytes,1,rep,name=type_urls,json=typeUrls,proto3" json:"type_urls,omitempty"`
}

func (x *SubscribeOnceRequest) Reset() {
	*x = SubscribeOnceRequest{}
	mi := &file_tkd_events_v1_events_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubscribeOnceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeOnceRequest) ProtoMessage() {}

func (x *SubscribeOnceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_events_v1_events_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeOnceRequest.ProtoReflect.Descriptor instead.
func (*SubscribeOnceRequest) Descriptor() ([]byte, []int) {
	return file_tkd_events_v1_events_proto_rawDescGZIP(), []int{2}
}

func (x *SubscribeOnceRequest) GetTypeUrls() []string {
	if x != nil {
		return x.TypeUrls
	}
	return nil
}

var File_tkd_events_v1_events_proto protoreflect.FileDescriptor

var file_tkd_events_v1_events_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x6b, 0x64, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x74, 0x6b,
	0x64, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x64, 0x22, 0x5e, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x22, 0x0a, 0x0b, 0x75, 0x6e, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0b, 0x75, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x06, 0x0a, 0x04,
	0x6b, 0x69, 0x6e, 0x64, 0x22, 0x33, 0x0a, 0x14, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x4f, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x79, 0x70, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x32, 0x9e, 0x02, 0x0a, 0x0c, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x1f, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x28, 0x01,
	0x30, 0x01, 0x12, 0x4c, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4f,
	0x6e, 0x63, 0x65, 0x12, 0x23, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4f, 0x6e, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x30, 0x01,
	0x12, 0x37, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x14, 0x2e, 0x74, 0x6b,
	0x64, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3f, 0x0a, 0x0d, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x14, 0x2e, 0x74, 0x6b, 0x64,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x28, 0x01, 0x42, 0xbb, 0x01, 0x0a, 0x11, 0x63,
	0x6f, 0x6d, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x42, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x65, 0x72,
	0x6b, 0x6c, 0x69, 0x6e, 0x69, 0x6b, 0x2d, 0x64, 0x6f, 0x62, 0x65, 0x72, 0x73, 0x62, 0x65, 0x72,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x6b,
	0x64, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x45, 0x58, 0xaa, 0x02, 0x0d, 0x54, 0x6b, 0x64,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x54, 0x6b, 0x64,
	0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x54, 0x6b, 0x64,
	0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x54, 0x6b, 0x64, 0x3a, 0x3a, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tkd_events_v1_events_proto_rawDescOnce sync.Once
	file_tkd_events_v1_events_proto_rawDescData = file_tkd_events_v1_events_proto_rawDesc
)

func file_tkd_events_v1_events_proto_rawDescGZIP() []byte {
	file_tkd_events_v1_events_proto_rawDescOnce.Do(func() {
		file_tkd_events_v1_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_tkd_events_v1_events_proto_rawDescData)
	})
	return file_tkd_events_v1_events_proto_rawDescData
}

var file_tkd_events_v1_events_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tkd_events_v1_events_proto_goTypes = []any{
	(*Event)(nil),                // 0: tkd.events.v1.Event
	(*SubscribeRequest)(nil),     // 1: tkd.events.v1.SubscribeRequest
	(*SubscribeOnceRequest)(nil), // 2: tkd.events.v1.SubscribeOnceRequest
	(*anypb.Any)(nil),            // 3: google.protobuf.Any
	(*emptypb.Empty)(nil),        // 4: google.protobuf.Empty
}
var file_tkd_events_v1_events_proto_depIdxs = []int32{
	3, // 0: tkd.events.v1.Event.event:type_name -> google.protobuf.Any
	1, // 1: tkd.events.v1.EventService.Subscribe:input_type -> tkd.events.v1.SubscribeRequest
	2, // 2: tkd.events.v1.EventService.SubscribeOnce:input_type -> tkd.events.v1.SubscribeOnceRequest
	0, // 3: tkd.events.v1.EventService.Publish:input_type -> tkd.events.v1.Event
	0, // 4: tkd.events.v1.EventService.PublishStream:input_type -> tkd.events.v1.Event
	0, // 5: tkd.events.v1.EventService.Subscribe:output_type -> tkd.events.v1.Event
	0, // 6: tkd.events.v1.EventService.SubscribeOnce:output_type -> tkd.events.v1.Event
	4, // 7: tkd.events.v1.EventService.Publish:output_type -> google.protobuf.Empty
	4, // 8: tkd.events.v1.EventService.PublishStream:output_type -> google.protobuf.Empty
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tkd_events_v1_events_proto_init() }
func file_tkd_events_v1_events_proto_init() {
	if File_tkd_events_v1_events_proto != nil {
		return
	}
	file_tkd_events_v1_events_proto_msgTypes[1].OneofWrappers = []any{
		(*SubscribeRequest_Subscribe)(nil),
		(*SubscribeRequest_Unsubscribe)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tkd_events_v1_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tkd_events_v1_events_proto_goTypes,
		DependencyIndexes: file_tkd_events_v1_events_proto_depIdxs,
		MessageInfos:      file_tkd_events_v1_events_proto_msgTypes,
	}.Build()
	File_tkd_events_v1_events_proto = out.File
	file_tkd_events_v1_events_proto_rawDesc = nil
	file_tkd_events_v1_events_proto_goTypes = nil
	file_tkd_events_v1_events_proto_depIdxs = nil
}
