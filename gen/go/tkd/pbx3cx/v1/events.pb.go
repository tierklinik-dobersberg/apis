// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: tkd/pbx3cx/v1/events.proto

package pbx3cxv1

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

type OverwriteCreatedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Overwrite *Overwrite `protobuf:"bytes,1,opt,name=overwrite,proto3" json:"overwrite,omitempty"`
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Overwrite *Overwrite `protobuf:"bytes,1,opt,name=overwrite,proto3" json:"overwrite,omitempty"`
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallEntry *CallEntry `protobuf:"bytes,1,opt,name=call_entry,json=callEntry,proto3" json:"call_entry,omitempty"`
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Voicemail *VoiceMail `protobuf:"bytes,1,opt,name=voicemail,proto3" json:"voicemail,omitempty"`
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OnCall                []*OnCall `protobuf:"bytes,1,rep,name=on_call,json=onCall,proto3" json:"on_call,omitempty"`
	RosterDate            string    `protobuf:"bytes,2,opt,name=roster_date,json=rosterDate,proto3" json:"roster_date,omitempty"` // YYYY-MM-DD
	IsOverwrite           bool      `protobuf:"varint,3,opt,name=is_overwrite,json=isOverwrite,proto3" json:"is_overwrite,omitempty"`
	PrimaryTransferTarget string    `protobuf:"bytes,4,opt,name=primary_transfer_target,json=primaryTransferTarget,proto3" json:"primary_transfer_target,omitempty"`
	InboundNumber         string    `protobuf:"bytes,5,opt,name=inbound_number,json=inboundNumber,proto3" json:"inbound_number,omitempty"`
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

var file_tkd_pbx3cx_v1_events_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x6b, 0x64, 0x2f, 0x70, 0x62, 0x78, 0x33, 0x63, 0x78, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x74, 0x6b,
	0x64, 0x2e, 0x70, 0x62, 0x78, 0x33, 0x63, 0x78, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x74, 0x6b, 0x64,
	0x2f, 0x70, 0x62, 0x78, 0x33, 0x63, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x6c,
	0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x74, 0x6b, 0x64, 0x2f, 0x70, 0x62,
	0x78, 0x33, 0x63, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x15, 0x4f, 0x76, 0x65, 0x72, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x36, 0x0a, 0x09, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x70, 0x62, 0x78, 0x33, 0x63, 0x78,
	0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x52, 0x09, 0x6f,
	0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x22, 0x4f, 0x0a, 0x15, 0x4f, 0x76, 0x65, 0x72,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x36, 0x0a, 0x09, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x70, 0x62, 0x78, 0x33, 0x63,
	0x78, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x52, 0x09,
	0x6f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x22, 0x4d, 0x0a, 0x12, 0x43, 0x61, 0x6c,
	0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12,
	0x37, 0x0a, 0x0a, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x70, 0x62, 0x78, 0x33, 0x63, 0x78,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x63,
	0x61, 0x6c, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x50, 0x0a, 0x16, 0x56, 0x6f, 0x69, 0x63,
	0x65, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x36, 0x0a, 0x09, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x70, 0x62, 0x78, 0x33,
	0x63, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x52,
	0x09, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xe6, 0x01, 0x0a, 0x11, 0x4f,
	0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x2e, 0x0a, 0x07, 0x6f, 0x6e, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x70, 0x62, 0x78, 0x33, 0x63, 0x78, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x06, 0x6f, 0x6e, 0x43, 0x61, 0x6c, 0x6c,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x4f, 0x76, 0x65, 0x72, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x12, 0x36, 0x0a, 0x17, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x25, 0x0a, 0x0e,
	0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x42, 0xbb, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x6b, 0x64, 0x2e,
	0x70, 0x62, 0x78, 0x33, 0x63, 0x78, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x65, 0x72, 0x6b, 0x6c, 0x69, 0x6e, 0x69, 0x6b, 0x2d,
	0x64, 0x6f, 0x62, 0x65, 0x72, 0x73, 0x62, 0x65, 0x72, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x6b, 0x64, 0x2f, 0x70, 0x62, 0x78, 0x33, 0x63,
	0x78, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x78, 0x33, 0x63, 0x78, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x54, 0x50, 0x58, 0xaa, 0x02, 0x0d, 0x54, 0x6b, 0x64, 0x2e, 0x50, 0x62, 0x78, 0x33, 0x63, 0x78,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x54, 0x6b, 0x64, 0x5c, 0x50, 0x62, 0x78, 0x33, 0x63, 0x78,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x54, 0x6b, 0x64, 0x5c, 0x50, 0x62, 0x78, 0x33, 0x63, 0x78,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0f, 0x54, 0x6b, 0x64, 0x3a, 0x3a, 0x50, 0x62, 0x78, 0x33, 0x63, 0x78, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tkd_pbx3cx_v1_events_proto_rawDescOnce sync.Once
	file_tkd_pbx3cx_v1_events_proto_rawDescData = file_tkd_pbx3cx_v1_events_proto_rawDesc
)

func file_tkd_pbx3cx_v1_events_proto_rawDescGZIP() []byte {
	file_tkd_pbx3cx_v1_events_proto_rawDescOnce.Do(func() {
		file_tkd_pbx3cx_v1_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_tkd_pbx3cx_v1_events_proto_rawDescData)
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
			RawDescriptor: file_tkd_pbx3cx_v1_events_proto_rawDesc,
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
	file_tkd_pbx3cx_v1_events_proto_rawDesc = nil
	file_tkd_pbx3cx_v1_events_proto_goTypes = nil
	file_tkd_pbx3cx_v1_events_proto_depIdxs = nil
}
