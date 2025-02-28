// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: tkd/calendar/v1/holiday_service.proto

package calendarv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	v1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/common/v1"
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

// HolidayType specifies the type of a public holiday.
type HolidayType int32

const (
	HolidayType_HOLIDAY_TYPE_UNSPECIFIED HolidayType = 0
	HolidayType_PUBLIC                   HolidayType = 1
	HolidayType_BANK                     HolidayType = 2
	HolidayType_SCHOOL                   HolidayType = 3
	HolidayType_AUTHORITIES              HolidayType = 4
	HolidayType_OPTIONAL                 HolidayType = 5
	HolidayType_OBSERVANCE               HolidayType = 6
)

// Enum value maps for HolidayType.
var (
	HolidayType_name = map[int32]string{
		0: "HOLIDAY_TYPE_UNSPECIFIED",
		1: "PUBLIC",
		2: "BANK",
		3: "SCHOOL",
		4: "AUTHORITIES",
		5: "OPTIONAL",
		6: "OBSERVANCE",
	}
	HolidayType_value = map[string]int32{
		"HOLIDAY_TYPE_UNSPECIFIED": 0,
		"PUBLIC":                   1,
		"BANK":                     2,
		"SCHOOL":                   3,
		"AUTHORITIES":              4,
		"OPTIONAL":                 5,
		"OBSERVANCE":               6,
	}
)

func (x HolidayType) Enum() *HolidayType {
	p := new(HolidayType)
	*p = x
	return p
}

func (x HolidayType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HolidayType) Descriptor() protoreflect.EnumDescriptor {
	return file_tkd_calendar_v1_holiday_service_proto_enumTypes[0].Descriptor()
}

func (HolidayType) Type() protoreflect.EnumType {
	return &file_tkd_calendar_v1_holiday_service_proto_enumTypes[0]
}

func (x HolidayType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HolidayType.Descriptor instead.
func (HolidayType) EnumDescriptor() ([]byte, []int) {
	return file_tkd_calendar_v1_holiday_service_proto_rawDescGZIP(), []int{0}
}

// PublicHoliday describes a public holiday at a specified country.
type PublicHoliday struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Date is the date of the public holidy in the format of YYYY-MM-DD.
	Date string `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"` // YYYY-MM-DD
	// LocalName is the localized name of the public holiday in the offical
	// language of the specified country.
	LocalName string `protobuf:"bytes,2,opt,name=local_name,json=localName,proto3" json:"local_name,omitempty"`
	// Name is the international, english name of the public holiday.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// CountryCode holds the ISO 2-Letter country code.
	CountryCode string `protobuf:"bytes,4,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	// Fixed is set to true if this public holiday is always at the same date
	// like the National Public Holiday.
	Fixed bool `protobuf:"varint,5,opt,name=fixed,proto3" json:"fixed,omitempty"`
	// Global is set to true if the public holiday is globally accepted.
	Global bool `protobuf:"varint,6,opt,name=global,proto3" json:"global,omitempty"`
	// Type holds the type of the public holiday.
	Type          HolidayType `protobuf:"varint,7,opt,name=type,proto3,enum=tkd.calendar.v1.HolidayType" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PublicHoliday) Reset() {
	*x = PublicHoliday{}
	mi := &file_tkd_calendar_v1_holiday_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublicHoliday) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicHoliday) ProtoMessage() {}

func (x *PublicHoliday) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_calendar_v1_holiday_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicHoliday.ProtoReflect.Descriptor instead.
func (*PublicHoliday) Descriptor() ([]byte, []int) {
	return file_tkd_calendar_v1_holiday_service_proto_rawDescGZIP(), []int{0}
}

func (x *PublicHoliday) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *PublicHoliday) GetLocalName() string {
	if x != nil {
		return x.LocalName
	}
	return ""
}

func (x *PublicHoliday) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PublicHoliday) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *PublicHoliday) GetFixed() bool {
	if x != nil {
		return x.Fixed
	}
	return false
}

func (x *PublicHoliday) GetGlobal() bool {
	if x != nil {
		return x.Global
	}
	return false
}

func (x *PublicHoliday) GetType() HolidayType {
	if x != nil {
		return x.Type
	}
	return HolidayType_HOLIDAY_TYPE_UNSPECIFIED
}

// GetHolidayRequest is the request message for the GetHoliday RPC.
type GetHolidayRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Year holds the year for which holidays should be queried.
	Year uint64 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	// Month may be set to a month (1 to 12). If set, only holidays for that
	// month will be returned.
	Month uint64 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	// CountryCode might be set to the country code for which holidays should be
	// queried. If left empty, the default country code from cis-cal
	// configuration is used.
	CountryCode   string `protobuf:"bytes,3,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetHolidayRequest) Reset() {
	*x = GetHolidayRequest{}
	mi := &file_tkd_calendar_v1_holiday_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetHolidayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHolidayRequest) ProtoMessage() {}

func (x *GetHolidayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_calendar_v1_holiday_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHolidayRequest.ProtoReflect.Descriptor instead.
func (*GetHolidayRequest) Descriptor() ([]byte, []int) {
	return file_tkd_calendar_v1_holiday_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetHolidayRequest) GetYear() uint64 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *GetHolidayRequest) GetMonth() uint64 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *GetHolidayRequest) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

// GetHolidayResponse is the response message of the GetHoliday RPC and contains
// a list of public holidays.
type GetHolidayResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Holidays is the list of public holidays that matched the search query.
	Holidays      []*PublicHoliday `protobuf:"bytes,1,rep,name=holidays,proto3" json:"holidays,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetHolidayResponse) Reset() {
	*x = GetHolidayResponse{}
	mi := &file_tkd_calendar_v1_holiday_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetHolidayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHolidayResponse) ProtoMessage() {}

func (x *GetHolidayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_calendar_v1_holiday_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHolidayResponse.ProtoReflect.Descriptor instead.
func (*GetHolidayResponse) Descriptor() ([]byte, []int) {
	return file_tkd_calendar_v1_holiday_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetHolidayResponse) GetHolidays() []*PublicHoliday {
	if x != nil {
		return x.Holidays
	}
	return nil
}

// NumberOfWorkDaysRequest is the request message for the NumberOfWorkDays RPC.
type NumberOfWorkDaysRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Country specified the country for which the number of working days should
	// be calculated. If left empty, the default country from the cis-cal
	// configuration will be used.
	Country string `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
	// From defines the start time (inclusive).
	From *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	// To defines the end time (inclusive).
	To            *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NumberOfWorkDaysRequest) Reset() {
	*x = NumberOfWorkDaysRequest{}
	mi := &file_tkd_calendar_v1_holiday_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NumberOfWorkDaysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumberOfWorkDaysRequest) ProtoMessage() {}

func (x *NumberOfWorkDaysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_calendar_v1_holiday_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumberOfWorkDaysRequest.ProtoReflect.Descriptor instead.
func (*NumberOfWorkDaysRequest) Descriptor() ([]byte, []int) {
	return file_tkd_calendar_v1_holiday_service_proto_rawDescGZIP(), []int{3}
}

func (x *NumberOfWorkDaysRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *NumberOfWorkDaysRequest) GetFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *NumberOfWorkDaysRequest) GetTo() *timestamppb.Timestamp {
	if x != nil {
		return x.To
	}
	return nil
}

// NumberOfWorkDaysResponse is the response message of the NumberOfWorkDays RPC.
type NumberOfWorkDaysResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// NumberOfWorkDays is the number of working days in the specified time
	// range.
	NumberOfWorkDays uint32 `protobuf:"varint,1,opt,name=number_of_work_days,json=numberOfWorkDays,proto3" json:"number_of_work_days,omitempty"`
	// NumberOfWeekendDays is the number of weekend days in the specified time
	// range. Note that holidays on weekends are counted for both,
	// number_of_weekend_days and number_of_holidays.
	NumberOfWeekendDays uint32 `protobuf:"varint,2,opt,name=number_of_weekend_days,json=numberOfWeekendDays,proto3" json:"number_of_weekend_days,omitempty"`
	// NumberOfHolidays is the number of holidays in the specified time-range.
	// Note that holidays on weekends are counted for both,
	// number_of_weekend_days and number_of_holidays.
	NumberOfHolidays uint32 `protobuf:"varint,3,opt,name=number_of_holidays,json=numberOfHolidays,proto3" json:"number_of_holidays,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *NumberOfWorkDaysResponse) Reset() {
	*x = NumberOfWorkDaysResponse{}
	mi := &file_tkd_calendar_v1_holiday_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NumberOfWorkDaysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumberOfWorkDaysResponse) ProtoMessage() {}

func (x *NumberOfWorkDaysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_calendar_v1_holiday_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumberOfWorkDaysResponse.ProtoReflect.Descriptor instead.
func (*NumberOfWorkDaysResponse) Descriptor() ([]byte, []int) {
	return file_tkd_calendar_v1_holiday_service_proto_rawDescGZIP(), []int{4}
}

func (x *NumberOfWorkDaysResponse) GetNumberOfWorkDays() uint32 {
	if x != nil {
		return x.NumberOfWorkDays
	}
	return 0
}

func (x *NumberOfWorkDaysResponse) GetNumberOfWeekendDays() uint32 {
	if x != nil {
		return x.NumberOfWeekendDays
	}
	return 0
}

func (x *NumberOfWorkDaysResponse) GetNumberOfHolidays() uint32 {
	if x != nil {
		return x.NumberOfHolidays
	}
	return 0
}

type IsHolidayRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Date is the date to check.
	// If unset, the current day in the server's timezone is used.
	Date          *v1.Date `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IsHolidayRequest) Reset() {
	*x = IsHolidayRequest{}
	mi := &file_tkd_calendar_v1_holiday_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsHolidayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsHolidayRequest) ProtoMessage() {}

func (x *IsHolidayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_calendar_v1_holiday_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsHolidayRequest.ProtoReflect.Descriptor instead.
func (*IsHolidayRequest) Descriptor() ([]byte, []int) {
	return file_tkd_calendar_v1_holiday_service_proto_rawDescGZIP(), []int{5}
}

func (x *IsHolidayRequest) GetDate() *v1.Date {
	if x != nil {
		return x.Date
	}
	return nil
}

type IsHolidayResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Whether or not the queried date is a holiday.
	IsHoliday bool `protobuf:"varint,1,opt,name=is_holiday,json=isHoliday,proto3" json:"is_holiday,omitempty"`
	// The holiday, if any.
	Holiday *PublicHoliday `protobuf:"bytes,2,opt,name=holiday,proto3" json:"holiday,omitempty"`
	// QueriedDate holds the queried date.
	QueriedDate   *v1.Date `protobuf:"bytes,3,opt,name=queried_date,json=queriedDate,proto3" json:"queried_date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IsHolidayResponse) Reset() {
	*x = IsHolidayResponse{}
	mi := &file_tkd_calendar_v1_holiday_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsHolidayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsHolidayResponse) ProtoMessage() {}

func (x *IsHolidayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_calendar_v1_holiday_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsHolidayResponse.ProtoReflect.Descriptor instead.
func (*IsHolidayResponse) Descriptor() ([]byte, []int) {
	return file_tkd_calendar_v1_holiday_service_proto_rawDescGZIP(), []int{6}
}

func (x *IsHolidayResponse) GetIsHoliday() bool {
	if x != nil {
		return x.IsHoliday
	}
	return false
}

func (x *IsHolidayResponse) GetHoliday() *PublicHoliday {
	if x != nil {
		return x.Holiday
	}
	return nil
}

func (x *IsHolidayResponse) GetQueriedDate() *v1.Date {
	if x != nil {
		return x.QueriedDate
	}
	return nil
}

var File_tkd_calendar_v1_holiday_service_proto protoreflect.FileDescriptor

var file_tkd_calendar_v1_holiday_service_proto_rawDesc = string([]byte{
	0x0a, 0x25, 0x74, 0x6b, 0x64, 0x2f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x74, 0x6b, 0x64, 0x2e, 0x63, 0x61, 0x6c,
	0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x74, 0x6b, 0x64, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd9, 0x01, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x48, 0x6f, 0x6c, 0x69, 0x64,
	0x61, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x63,
	0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x6c, 0x69, 0x64,
	0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x78, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x07, 0xfa, 0xf7, 0x18, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x23,
	0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0d, 0xfa,
	0xf7, 0x18, 0x09, 0xd0, 0x01, 0x01, 0x32, 0x04, 0x18, 0x0c, 0x28, 0x01, 0x52, 0x05, 0x6d, 0x6f,
	0x6e, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x50, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6c,
	0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x08,
	0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x08,
	0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x73, 0x22, 0xa1, 0x01, 0x0a, 0x17, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x4f, 0x66, 0x57, 0x6f, 0x72, 0x6b, 0x44, 0x61, 0x79, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x37,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x07, 0xfa, 0xf7, 0x18, 0x03, 0xc8, 0x01,
	0x01, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x33, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x07, 0xfa, 0xf7, 0x18, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x02, 0x74, 0x6f, 0x22, 0xac, 0x01, 0x0a,
	0x18, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x57, 0x6f, 0x72, 0x6b, 0x44, 0x61, 0x79,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x13, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x64, 0x61, 0x79, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66,
	0x57, 0x6f, 0x72, 0x6b, 0x44, 0x61, 0x79, 0x73, 0x12, 0x33, 0x0a, 0x16, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61,
	0x79, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x79, 0x73, 0x12, 0x2c, 0x0a,
	0x12, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x68, 0x6f, 0x6c, 0x69, 0x64,
	0x61, 0x79, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x4f, 0x66, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x73, 0x22, 0x3b, 0x0a, 0x10, 0x49,
	0x73, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x74, 0x6b, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61,
	0x74, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0xa4, 0x01, 0x0a, 0x11, 0x49, 0x73, 0x48,
	0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x12, 0x38, 0x0a,
	0x07, 0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x07,
	0x68, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x12, 0x36, 0x0a, 0x0c, 0x71, 0x75, 0x65, 0x72, 0x69,
	0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x74, 0x6b, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61,
	0x74, 0x65, 0x52, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x2a,
	0x7c, 0x0a, 0x0b, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c,
	0x0a, 0x18, 0x48, 0x4f, 0x4c, 0x49, 0x44, 0x41, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x41, 0x4e, 0x4b,
	0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x43, 0x48, 0x4f, 0x4f, 0x4c, 0x10, 0x03, 0x12, 0x0f,
	0x0a, 0x0b, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x54, 0x49, 0x45, 0x53, 0x10, 0x04, 0x12,
	0x0c, 0x0a, 0x08, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x4c, 0x10, 0x05, 0x12, 0x0e, 0x0a,
	0x0a, 0x4f, 0x42, 0x53, 0x45, 0x52, 0x56, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x06, 0x32, 0xaa, 0x02,
	0x0a, 0x0e, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x57, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x12, 0x22,
	0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x09, 0x49, 0x73, 0x48,
	0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x12, 0x21, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x63, 0x61, 0x6c,
	0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x48, 0x6f, 0x6c, 0x69, 0x64,
	0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x74, 0x6b, 0x64, 0x2e,
	0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x48, 0x6f,
	0x6c, 0x69, 0x64, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x69, 0x0a, 0x10, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x57, 0x6f, 0x72, 0x6b, 0x44,
	0x61, 0x79, 0x73, 0x12, 0x28, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64,
	0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x57, 0x6f,
	0x72, 0x6b, 0x44, 0x61, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x74, 0x6b, 0x64, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x57, 0x6f, 0x72, 0x6b, 0x44, 0x61, 0x79, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xd1, 0x01, 0x0a, 0x13, 0x63,
	0x6f, 0x6d, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e,
	0x76, 0x31, 0x42, 0x13, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x65, 0x72, 0x6b, 0x6c, 0x69, 0x6e, 0x69, 0x6b,
	0x2d, 0x64, 0x6f, 0x62, 0x65, 0x72, 0x73, 0x62, 0x65, 0x72, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x6b, 0x64, 0x2f, 0x63, 0x61, 0x6c, 0x65,
	0x6e, 0x64, 0x61, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x43, 0x58, 0xaa, 0x02, 0x0f, 0x54, 0x6b, 0x64, 0x2e, 0x43,
	0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x54, 0x6b, 0x64,
	0x5c, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x54,
	0x6b, 0x64, 0x5c, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x54, 0x6b, 0x64,
	0x3a, 0x3a, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_tkd_calendar_v1_holiday_service_proto_rawDescOnce sync.Once
	file_tkd_calendar_v1_holiday_service_proto_rawDescData []byte
)

func file_tkd_calendar_v1_holiday_service_proto_rawDescGZIP() []byte {
	file_tkd_calendar_v1_holiday_service_proto_rawDescOnce.Do(func() {
		file_tkd_calendar_v1_holiday_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_tkd_calendar_v1_holiday_service_proto_rawDesc), len(file_tkd_calendar_v1_holiday_service_proto_rawDesc)))
	})
	return file_tkd_calendar_v1_holiday_service_proto_rawDescData
}

var file_tkd_calendar_v1_holiday_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tkd_calendar_v1_holiday_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_tkd_calendar_v1_holiday_service_proto_goTypes = []any{
	(HolidayType)(0),                 // 0: tkd.calendar.v1.HolidayType
	(*PublicHoliday)(nil),            // 1: tkd.calendar.v1.PublicHoliday
	(*GetHolidayRequest)(nil),        // 2: tkd.calendar.v1.GetHolidayRequest
	(*GetHolidayResponse)(nil),       // 3: tkd.calendar.v1.GetHolidayResponse
	(*NumberOfWorkDaysRequest)(nil),  // 4: tkd.calendar.v1.NumberOfWorkDaysRequest
	(*NumberOfWorkDaysResponse)(nil), // 5: tkd.calendar.v1.NumberOfWorkDaysResponse
	(*IsHolidayRequest)(nil),         // 6: tkd.calendar.v1.IsHolidayRequest
	(*IsHolidayResponse)(nil),        // 7: tkd.calendar.v1.IsHolidayResponse
	(*timestamppb.Timestamp)(nil),    // 8: google.protobuf.Timestamp
	(*v1.Date)(nil),                  // 9: tkd.common.v1.Date
}
var file_tkd_calendar_v1_holiday_service_proto_depIdxs = []int32{
	0,  // 0: tkd.calendar.v1.PublicHoliday.type:type_name -> tkd.calendar.v1.HolidayType
	1,  // 1: tkd.calendar.v1.GetHolidayResponse.holidays:type_name -> tkd.calendar.v1.PublicHoliday
	8,  // 2: tkd.calendar.v1.NumberOfWorkDaysRequest.from:type_name -> google.protobuf.Timestamp
	8,  // 3: tkd.calendar.v1.NumberOfWorkDaysRequest.to:type_name -> google.protobuf.Timestamp
	9,  // 4: tkd.calendar.v1.IsHolidayRequest.date:type_name -> tkd.common.v1.Date
	1,  // 5: tkd.calendar.v1.IsHolidayResponse.holiday:type_name -> tkd.calendar.v1.PublicHoliday
	9,  // 6: tkd.calendar.v1.IsHolidayResponse.queried_date:type_name -> tkd.common.v1.Date
	2,  // 7: tkd.calendar.v1.HolidayService.GetHoliday:input_type -> tkd.calendar.v1.GetHolidayRequest
	6,  // 8: tkd.calendar.v1.HolidayService.IsHoliday:input_type -> tkd.calendar.v1.IsHolidayRequest
	4,  // 9: tkd.calendar.v1.HolidayService.NumberOfWorkDays:input_type -> tkd.calendar.v1.NumberOfWorkDaysRequest
	3,  // 10: tkd.calendar.v1.HolidayService.GetHoliday:output_type -> tkd.calendar.v1.GetHolidayResponse
	7,  // 11: tkd.calendar.v1.HolidayService.IsHoliday:output_type -> tkd.calendar.v1.IsHolidayResponse
	5,  // 12: tkd.calendar.v1.HolidayService.NumberOfWorkDays:output_type -> tkd.calendar.v1.NumberOfWorkDaysResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_tkd_calendar_v1_holiday_service_proto_init() }
func file_tkd_calendar_v1_holiday_service_proto_init() {
	if File_tkd_calendar_v1_holiday_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_tkd_calendar_v1_holiday_service_proto_rawDesc), len(file_tkd_calendar_v1_holiday_service_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tkd_calendar_v1_holiday_service_proto_goTypes,
		DependencyIndexes: file_tkd_calendar_v1_holiday_service_proto_depIdxs,
		EnumInfos:         file_tkd_calendar_v1_holiday_service_proto_enumTypes,
		MessageInfos:      file_tkd_calendar_v1_holiday_service_proto_msgTypes,
	}.Build()
	File_tkd_calendar_v1_holiday_service_proto = out.File
	file_tkd_calendar_v1_holiday_service_proto_goTypes = nil
	file_tkd_calendar_v1_holiday_service_proto_depIdxs = nil
}
