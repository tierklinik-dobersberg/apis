// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: tkd/calendar/v1/event.proto

package calendarv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Calendar describes a event calendar.
type Calendar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is a unique ID for the calendar. The format of the ID may depend on
	// the application and the actual calendar backend.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name is a human friendly name of the calendar.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Timezone holds the timezone information for the calendar. In the format
	// of Europe/Vienna.
	Timezone string `protobuf:"bytes,3,opt,name=timezone,proto3" json:"timezone,omitempty"`
	// Color holds the calendar color.
	Color string `protobuf:"bytes,4,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *Calendar) Reset() {
	*x = Calendar{}
	mi := &file_tkd_calendar_v1_event_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Calendar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Calendar) ProtoMessage() {}

func (x *Calendar) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_calendar_v1_event_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Calendar.ProtoReflect.Descriptor instead.
func (*Calendar) Descriptor() ([]byte, []int) {
	return file_tkd_calendar_v1_event_proto_rawDescGZIP(), []int{0}
}

func (x *Calendar) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Calendar) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Calendar) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *Calendar) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

type CalendarEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is a unique ID for the calendar event.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// CalendarId is the unique ID of the calendar this event belongs to.
	CalendarId string `protobuf:"bytes,2,opt,name=calendar_id,json=calendarId,proto3" json:"calendar_id,omitempty"`
	// StartTime holds the time the event begins.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// EndTime optionally holds the time the event ends.
	EndTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// FullDay may be set to true for events the last the whole day. In this
	// case, the time part (HH:MM:SS) of StartTime is not important!
	FullDay bool `protobuf:"varint,5,opt,name=full_day,json=fullDay,proto3" json:"full_day,omitempty"`
	// Summary of the calendar event.
	Summary string `protobuf:"bytes,6,opt,name=summary,proto3" json:"summary,omitempty"`
	// Description of the calendar event.
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	// ExtraData may holds additional information about the calendar entry. In
	// most cases, this should be CustomerAnnotation
	ExtraData *anypb.Any `protobuf:"bytes,8,opt,name=extra_data,json=extraData,proto3" json:"extra_data,omitempty"`
}

func (x *CalendarEvent) Reset() {
	*x = CalendarEvent{}
	mi := &file_tkd_calendar_v1_event_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CalendarEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalendarEvent) ProtoMessage() {}

func (x *CalendarEvent) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_calendar_v1_event_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalendarEvent.ProtoReflect.Descriptor instead.
func (*CalendarEvent) Descriptor() ([]byte, []int) {
	return file_tkd_calendar_v1_event_proto_rawDescGZIP(), []int{1}
}

func (x *CalendarEvent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CalendarEvent) GetCalendarId() string {
	if x != nil {
		return x.CalendarId
	}
	return ""
}

func (x *CalendarEvent) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *CalendarEvent) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *CalendarEvent) GetFullDay() bool {
	if x != nil {
		return x.FullDay
	}
	return false
}

func (x *CalendarEvent) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *CalendarEvent) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CalendarEvent) GetExtraData() *anypb.Any {
	if x != nil {
		return x.ExtraData
	}
	return nil
}

type CustomerAnnotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// CustomerSource may be set to the source of the customer record.
	CustomerSource string `protobuf:"bytes,1,opt,name=customer_source,json=customerSource,proto3" json:"customer_source,omitempty"`
	// CustomerId is the ID of the customer within the specified source.
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// AnimalIds is a list of animals that are expected to show up during the
	// event and need treatment.
	AnimalIds []string `protobuf:"bytes,3,rep,name=animal_ids,json=animalIds,proto3" json:"animal_ids,omitempty"`
	// CreatedByUserId holds the ID of the user that created the event.
	CreatedByUserId string `protobuf:"bytes,4,opt,name=created_by_user_id,json=createdByUserId,proto3" json:"created_by_user_id,omitempty"`
}

func (x *CustomerAnnotation) Reset() {
	*x = CustomerAnnotation{}
	mi := &file_tkd_calendar_v1_event_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomerAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerAnnotation) ProtoMessage() {}

func (x *CustomerAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_calendar_v1_event_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerAnnotation.ProtoReflect.Descriptor instead.
func (*CustomerAnnotation) Descriptor() ([]byte, []int) {
	return file_tkd_calendar_v1_event_proto_rawDescGZIP(), []int{2}
}

func (x *CustomerAnnotation) GetCustomerSource() string {
	if x != nil {
		return x.CustomerSource
	}
	return ""
}

func (x *CustomerAnnotation) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *CustomerAnnotation) GetAnimalIds() []string {
	if x != nil {
		return x.AnimalIds
	}
	return nil
}

func (x *CustomerAnnotation) GetCreatedByUserId() string {
	if x != nil {
		return x.CreatedByUserId
	}
	return ""
}

var File_tkd_calendar_v1_event_proto protoreflect.FileDescriptor

var file_tkd_calendar_v1_event_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x74, 0x6b, 0x64, 0x2f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x74,
	0x6b, 0x64, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x08, 0x43, 0x61,
	0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69,
	0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69,
	0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x22, 0xbe, 0x02, 0x0a,
	0x0d, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x49, 0x64, 0x12,
	0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x64, 0x61, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x66, 0x75, 0x6c, 0x6c, 0x44, 0x61, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x09, 0x65, 0x78, 0x74, 0x72, 0x61, 0x44, 0x61, 0x74, 0x61, 0x22, 0xaa, 0x01,
	0x0a, 0x12, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x49, 0x64, 0x73, 0x12, 0x2b, 0x0a,
	0x12, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x42, 0xc8, 0x01, 0x0a, 0x13, 0x63,
	0x6f, 0x6d, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e,
	0x76, 0x31, 0x42, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x65,
	0x72, 0x6b, 0x6c, 0x69, 0x6e, 0x69, 0x6b, 0x2d, 0x64, 0x6f, 0x62, 0x65, 0x72, 0x73, 0x62, 0x65,
	0x72, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74,
	0x6b, 0x64, 0x2f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x63,
	0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x43, 0x58, 0xaa,
	0x02, 0x0f, 0x54, 0x6b, 0x64, 0x2e, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x0f, 0x54, 0x6b, 0x64, 0x5c, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x54, 0x6b, 0x64, 0x5c, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64,
	0x61, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x11, 0x54, 0x6b, 0x64, 0x3a, 0x3a, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61,
	0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tkd_calendar_v1_event_proto_rawDescOnce sync.Once
	file_tkd_calendar_v1_event_proto_rawDescData = file_tkd_calendar_v1_event_proto_rawDesc
)

func file_tkd_calendar_v1_event_proto_rawDescGZIP() []byte {
	file_tkd_calendar_v1_event_proto_rawDescOnce.Do(func() {
		file_tkd_calendar_v1_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_tkd_calendar_v1_event_proto_rawDescData)
	})
	return file_tkd_calendar_v1_event_proto_rawDescData
}

var file_tkd_calendar_v1_event_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tkd_calendar_v1_event_proto_goTypes = []any{
	(*Calendar)(nil),              // 0: tkd.calendar.v1.Calendar
	(*CalendarEvent)(nil),         // 1: tkd.calendar.v1.CalendarEvent
	(*CustomerAnnotation)(nil),    // 2: tkd.calendar.v1.CustomerAnnotation
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*anypb.Any)(nil),             // 4: google.protobuf.Any
}
var file_tkd_calendar_v1_event_proto_depIdxs = []int32{
	3, // 0: tkd.calendar.v1.CalendarEvent.start_time:type_name -> google.protobuf.Timestamp
	3, // 1: tkd.calendar.v1.CalendarEvent.end_time:type_name -> google.protobuf.Timestamp
	4, // 2: tkd.calendar.v1.CalendarEvent.extra_data:type_name -> google.protobuf.Any
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tkd_calendar_v1_event_proto_init() }
func file_tkd_calendar_v1_event_proto_init() {
	if File_tkd_calendar_v1_event_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tkd_calendar_v1_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tkd_calendar_v1_event_proto_goTypes,
		DependencyIndexes: file_tkd_calendar_v1_event_proto_depIdxs,
		MessageInfos:      file_tkd_calendar_v1_event_proto_msgTypes,
	}.Build()
	File_tkd_calendar_v1_event_proto = out.File
	file_tkd_calendar_v1_event_proto_rawDesc = nil
	file_tkd_calendar_v1_event_proto_goTypes = nil
	file_tkd_calendar_v1_event_proto_depIdxs = nil
}
