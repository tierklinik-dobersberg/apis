// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: tkd/pbx3cx/v1/events.proto

package pbx3cxv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type OverwriteCreatedEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Overwrite     *Overwrite             `protobuf:"bytes,1,opt,name=overwrite,proto3" json:"overwrite,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OverwriteCreatedEvent) Reset() {
	*x = OverwriteCreatedEvent{}
	mi := &file_tkd_pbx3cx_v1_events_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OverwriteCreatedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverwriteCreatedEvent) ProtoMessage() {}

func (x *OverwriteCreatedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_pbx3cx_v1_events_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverwriteCreatedEvent.ProtoReflect.Descriptor instead.
func (*OverwriteCreatedEvent) Descriptor() ([]byte, []int) {
	return file_tkd_pbx3cx_v1_events_proto_rawDescGZIP(), []int{0}
}

func (x *OverwriteCreatedEvent) GetOverwrite() *Overwrite {
	if x != nil {
		return x.Overwrite
	}
	return nil
}

type OverwriteDeletedEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Overwrite     *Overwrite             `protobuf:"bytes,1,opt,name=overwrite,proto3" json:"overwrite,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OverwriteDeletedEvent) Reset() {
	*x = OverwriteDeletedEvent{}
	mi := &file_tkd_pbx3cx_v1_events_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OverwriteDeletedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverwriteDeletedEvent) ProtoMessage() {}

func (x *OverwriteDeletedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_pbx3cx_v1_events_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverwriteDeletedEvent.ProtoReflect.Descriptor instead.
func (*OverwriteDeletedEvent) Descriptor() ([]byte, []int) {
	return file_tkd_pbx3cx_v1_events_proto_rawDescGZIP(), []int{1}
}

func (x *OverwriteDeletedEvent) GetOverwrite() *Overwrite {
	if x != nil {
		return x.Overwrite
	}
	return nil
}

type CallRecordReceived struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CallEntry     *CallEntry             `protobuf:"bytes,1,opt,name=call_entry,json=callEntry,proto3" json:"call_entry,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CallRecordReceived) Reset() {
	*x = CallRecordReceived{}
	mi := &file_tkd_pbx3cx_v1_events_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CallRecordReceived) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallRecordReceived) ProtoMessage() {}

func (x *CallRecordReceived) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_pbx3cx_v1_events_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallRecordReceived.ProtoReflect.Descriptor instead.
func (*CallRecordReceived) Descriptor() ([]byte, []int) {
	return file_tkd_pbx3cx_v1_events_proto_rawDescGZIP(), []int{2}
}

func (x *CallRecordReceived) GetCallEntry() *CallEntry {
	if x != nil {
		return x.CallEntry
	}
	return nil
}

type VoiceMailReceivedEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Voicemail     *VoiceMail             `protobuf:"bytes,1,opt,name=voicemail,proto3" json:"voicemail,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VoiceMailReceivedEvent) Reset() {
	*x = VoiceMailReceivedEvent{}
	mi := &file_tkd_pbx3cx_v1_events_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VoiceMailReceivedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoiceMailReceivedEvent) ProtoMessage() {}

func (x *VoiceMailReceivedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_pbx3cx_v1_events_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoiceMailReceivedEvent.ProtoReflect.Descriptor instead.
func (*VoiceMailReceivedEvent) Descriptor() ([]byte, []int) {
	return file_tkd_pbx3cx_v1_events_proto_rawDescGZIP(), []int{3}
}

func (x *VoiceMailReceivedEvent) GetVoicemail() *VoiceMail {
	if x != nil {
		return x.Voicemail
	}
	return nil
}

type OnCallChangeEvent struct {
	state                 protoimpl.MessageState `protogen:"open.v1"`
	OnCall                []*OnCall              `protobuf:"bytes,1,rep,name=on_call,json=onCall,proto3" json:"on_call,omitempty"`
	RosterDate            string                 `protobuf:"bytes,2,opt,name=roster_date,json=rosterDate,proto3" json:"roster_date,omitempty"` // YYYY-MM-DD
	IsOverwrite           bool                   `protobuf:"varint,3,opt,name=is_overwrite,json=isOverwrite,proto3" json:"is_overwrite,omitempty"`
	PrimaryTransferTarget string                 `protobuf:"bytes,4,opt,name=primary_transfer_target,json=primaryTransferTarget,proto3" json:"primary_transfer_target,omitempty"`
	InboundNumber         string                 `protobuf:"bytes,5,opt,name=inbound_number,json=inboundNumber,proto3" json:"inbound_number,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *OnCallChangeEvent) Reset() {
	*x = OnCallChangeEvent{}
	mi := &file_tkd_pbx3cx_v1_events_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OnCallChangeEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnCallChangeEvent) ProtoMessage() {}

func (x *OnCallChangeEvent) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_pbx3cx_v1_events_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnCallChangeEvent.ProtoReflect.Descriptor instead.
func (*OnCallChangeEvent) Descriptor() ([]byte, []int) {
	return file_tkd_pbx3cx_v1_events_proto_rawDescGZIP(), []int{4}
}

func (x *OnCallChangeEvent) GetOnCall() []*OnCall {
	if x != nil {
		return x.OnCall
	}
	return nil
}

func (x *OnCallChangeEvent) GetRosterDate() string {
	if x != nil {
		return x.RosterDate
	}
	return ""
}

func (x *OnCallChangeEvent) GetIsOverwrite() bool {
	if x != nil {
		return x.IsOverwrite
	}
	return false
}

func (x *OnCallChangeEvent) GetPrimaryTransferTarget() string {
	if x != nil {
		return x.PrimaryTransferTarget
	}
	return ""
}

func (x *OnCallChangeEvent) GetInboundNumber() string {
	if x != nil {
		return x.InboundNumber
	}
	return ""
}

var File_tkd_pbx3cx_v1_events_proto protoreflect.FileDescriptor

const file_tkd_pbx3cx_v1_events_proto_rawDesc = "" +
	"\n" +
	"\x1atkd/pbx3cx/v1/events.proto\x12\rtkd.pbx3cx.v1\x1a\x1btkd/pbx3cx/v1/calllog.proto\x1a\x1dtkd/pbx3cx/v1/voicemail.proto\"O\n" +
	"\x15OverwriteCreatedEvent\x126\n" +
	"\toverwrite\x18\x01 \x01(\v2\x18.tkd.pbx3cx.v1.OverwriteR\toverwrite\"O\n" +
	"\x15OverwriteDeletedEvent\x126\n" +
	"\toverwrite\x18\x01 \x01(\v2\x18.tkd.pbx3cx.v1.OverwriteR\toverwrite\"M\n" +
	"\x12CallRecordReceived\x127\n" +
	"\n" +
	"call_entry\x18\x01 \x01(\v2\x18.tkd.pbx3cx.v1.CallEntryR\tcallEntry\"P\n" +
	"\x16VoiceMailReceivedEvent\x126\n" +
	"\tvoicemail\x18\x01 \x01(\v2\x18.tkd.pbx3cx.v1.VoiceMailR\tvoicemail\"\xe6\x01\n" +
	"\x11OnCallChangeEvent\x12.\n" +
	"\aon_call\x18\x01 \x03(\v2\x15.tkd.pbx3cx.v1.OnCallR\x06onCall\x12\x1f\n" +
	"\vroster_date\x18\x02 \x01(\tR\n" +
	"rosterDate\x12!\n" +
	"\fis_overwrite\x18\x03 \x01(\bR\visOverwrite\x126\n" +
	"\x17primary_transfer_target\x18\x04 \x01(\tR\x15primaryTransferTarget\x12%\n" +
	"\x0einbound_number\x18\x05 \x01(\tR\rinboundNumberB\xbb\x01\n" +
	"\x11com.tkd.pbx3cx.v1B\vEventsProtoP\x01ZCgithub.com/tierklinik-dobersberg/apis/gen/go/tkd/pbx3cx/v1;pbx3cxv1\xa2\x02\x03TPX\xaa\x02\rTkd.Pbx3cx.V1\xca\x02\rTkd\\Pbx3cx\\V1\xe2\x02\x19Tkd\\Pbx3cx\\V1\\GPBMetadata\xea\x02\x0fTkd::Pbx3cx::V1b\x06proto3"

var (
	file_tkd_pbx3cx_v1_events_proto_rawDescOnce sync.Once
	file_tkd_pbx3cx_v1_events_proto_rawDescData []byte
)

func file_tkd_pbx3cx_v1_events_proto_rawDescGZIP() []byte {
	file_tkd_pbx3cx_v1_events_proto_rawDescOnce.Do(func() {
		file_tkd_pbx3cx_v1_events_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_tkd_pbx3cx_v1_events_proto_rawDesc), len(file_tkd_pbx3cx_v1_events_proto_rawDesc)))
	})
	return file_tkd_pbx3cx_v1_events_proto_rawDescData
}

var file_tkd_pbx3cx_v1_events_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_tkd_pbx3cx_v1_events_proto_goTypes = []any{
	(*OverwriteCreatedEvent)(nil),  // 0: tkd.pbx3cx.v1.OverwriteCreatedEvent
	(*OverwriteDeletedEvent)(nil),  // 1: tkd.pbx3cx.v1.OverwriteDeletedEvent
	(*CallRecordReceived)(nil),     // 2: tkd.pbx3cx.v1.CallRecordReceived
	(*VoiceMailReceivedEvent)(nil), // 3: tkd.pbx3cx.v1.VoiceMailReceivedEvent
	(*OnCallChangeEvent)(nil),      // 4: tkd.pbx3cx.v1.OnCallChangeEvent
	(*Overwrite)(nil),              // 5: tkd.pbx3cx.v1.Overwrite
	(*CallEntry)(nil),              // 6: tkd.pbx3cx.v1.CallEntry
	(*VoiceMail)(nil),              // 7: tkd.pbx3cx.v1.VoiceMail
	(*OnCall)(nil),                 // 8: tkd.pbx3cx.v1.OnCall
}
var file_tkd_pbx3cx_v1_events_proto_depIdxs = []int32{
	5, // 0: tkd.pbx3cx.v1.OverwriteCreatedEvent.overwrite:type_name -> tkd.pbx3cx.v1.Overwrite
	5, // 1: tkd.pbx3cx.v1.OverwriteDeletedEvent.overwrite:type_name -> tkd.pbx3cx.v1.Overwrite
	6, // 2: tkd.pbx3cx.v1.CallRecordReceived.call_entry:type_name -> tkd.pbx3cx.v1.CallEntry
	7, // 3: tkd.pbx3cx.v1.VoiceMailReceivedEvent.voicemail:type_name -> tkd.pbx3cx.v1.VoiceMail
	8, // 4: tkd.pbx3cx.v1.OnCallChangeEvent.on_call:type_name -> tkd.pbx3cx.v1.OnCall
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_tkd_pbx3cx_v1_events_proto_init() }
func file_tkd_pbx3cx_v1_events_proto_init() {
	if File_tkd_pbx3cx_v1_events_proto != nil {
		return
	}
	file_tkd_pbx3cx_v1_calllog_proto_init()
	file_tkd_pbx3cx_v1_voicemail_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_tkd_pbx3cx_v1_events_proto_rawDesc), len(file_tkd_pbx3cx_v1_events_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tkd_pbx3cx_v1_events_proto_goTypes,
		DependencyIndexes: file_tkd_pbx3cx_v1_events_proto_depIdxs,
		MessageInfos:      file_tkd_pbx3cx_v1_events_proto_msgTypes,
	}.Build()
	File_tkd_pbx3cx_v1_events_proto = out.File
	file_tkd_pbx3cx_v1_events_proto_goTypes = nil
	file_tkd_pbx3cx_v1_events_proto_depIdxs = nil
}
