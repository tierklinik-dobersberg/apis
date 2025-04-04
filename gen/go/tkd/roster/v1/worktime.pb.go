// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: tkd/roster/v1/worktime.proto

package rosterv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/tierklinik-dobersberg/apis/gen/go/tkd/common/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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

// WorkTime describes the regular work time for a given user.
type WorkTime struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Id is a unique identifier for this work-time entry.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// UserId holds the ID of the user this work-time entry belongs to.
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// TimePerWeek defines the regular working time per week.
	TimePerWeek *durationpb.Duration `protobuf:"bytes,3,opt,name=time_per_week,json=timePerWeek,proto3" json:"time_per_week,omitempty"`
	// ApplicableAfter defines the date at which this work-time entry
	// is considered active.
	//
	// Format: YYYY-MM-DD
	ApplicableAfter string `protobuf:"bytes,4,opt,name=applicable_after,json=applicableAfter,proto3" json:"applicable_after,omitempty"`
	// VacationWeeksPerYear defines how many weeks of vacation should
	// be granted to UserId in a full-year of work time.
	VacationWeeksPerYear float32 `protobuf:"fixed32,5,opt,name=vacation_weeks_per_year,json=vacationWeeksPerYear,proto3" json:"vacation_weeks_per_year,omitempty"`
	// Overtime Allowance per month.
	OvertimeAllowancePerMonth *durationpb.Duration `protobuf:"bytes,6,opt,name=overtime_allowance_per_month,json=overtimeAllowancePerMonth,proto3" json:"overtime_allowance_per_month,omitempty"`
	// Whether or not the user should be excluded from time-tracking.
	ExcludeFromTimeTracking bool `protobuf:"varint,7,opt,name=exclude_from_time_tracking,json=excludeFromTimeTracking,proto3" json:"exclude_from_time_tracking,omitempty"`
	// An optional timestamp _AFTER_ which this work-time entry ends.
	//
	// Format: YYYY-MM-DD
	EndsWith      string `protobuf:"bytes,8,opt,name=ends_with,json=endsWith,proto3" json:"ends_with,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WorkTime) Reset() {
	*x = WorkTime{}
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkTime) ProtoMessage() {}

func (x *WorkTime) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkTime.ProtoReflect.Descriptor instead.
func (*WorkTime) Descriptor() ([]byte, []int) {
	return file_tkd_roster_v1_worktime_proto_rawDescGZIP(), []int{0}
}

func (x *WorkTime) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WorkTime) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *WorkTime) GetTimePerWeek() *durationpb.Duration {
	if x != nil {
		return x.TimePerWeek
	}
	return nil
}

func (x *WorkTime) GetApplicableAfter() string {
	if x != nil {
		return x.ApplicableAfter
	}
	return ""
}

func (x *WorkTime) GetVacationWeeksPerYear() float32 {
	if x != nil {
		return x.VacationWeeksPerYear
	}
	return 0
}

func (x *WorkTime) GetOvertimeAllowancePerMonth() *durationpb.Duration {
	if x != nil {
		return x.OvertimeAllowancePerMonth
	}
	return nil
}

func (x *WorkTime) GetExcludeFromTimeTracking() bool {
	if x != nil {
		return x.ExcludeFromTimeTracking
	}
	return false
}

func (x *WorkTime) GetEndsWith() string {
	if x != nil {
		return x.EndsWith
	}
	return ""
}

type SetWorkTimeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WorkTimes     []*WorkTime            `protobuf:"bytes,1,rep,name=work_times,json=workTimes,proto3" json:"work_times,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetWorkTimeRequest) Reset() {
	*x = SetWorkTimeRequest{}
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetWorkTimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetWorkTimeRequest) ProtoMessage() {}

func (x *SetWorkTimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetWorkTimeRequest.ProtoReflect.Descriptor instead.
func (*SetWorkTimeRequest) Descriptor() ([]byte, []int) {
	return file_tkd_roster_v1_worktime_proto_rawDescGZIP(), []int{1}
}

func (x *SetWorkTimeRequest) GetWorkTimes() []*WorkTime {
	if x != nil {
		return x.WorkTimes
	}
	return nil
}

type SetWorkTimeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	WorkTimes     []*WorkTime            `protobuf:"bytes,1,rep,name=work_times,json=workTimes,proto3" json:"work_times,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetWorkTimeResponse) Reset() {
	*x = SetWorkTimeResponse{}
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetWorkTimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetWorkTimeResponse) ProtoMessage() {}

func (x *SetWorkTimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetWorkTimeResponse.ProtoReflect.Descriptor instead.
func (*SetWorkTimeResponse) Descriptor() ([]byte, []int) {
	return file_tkd_roster_v1_worktime_proto_rawDescGZIP(), []int{2}
}

func (x *SetWorkTimeResponse) GetWorkTimes() []*WorkTime {
	if x != nil {
		return x.WorkTimes
	}
	return nil
}

type GetWorkTimeRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// UserIds is a list of user IDs for which the work-time should
	// be returned.
	UserIds []string `protobuf:"bytes,1,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	// ReadMask defines which fields of the response should be populated.
	// The field-mask is applied to the `result` field of the GetWorkTimeResponse.
	// For example, to only retrieve the user_id and the current work time specify
	// a read_mask like this:
	//
	//	{
	//	   read_mask: {
	//	     paths: [
	//	         "user_id",
	//	         "current",
	//	     ]
	//	   }
	//	}
	ReadMask      *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=read_mask,json=readMask,proto3" json:"read_mask,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetWorkTimeRequest) Reset() {
	*x = GetWorkTimeRequest{}
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWorkTimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkTimeRequest) ProtoMessage() {}

func (x *GetWorkTimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkTimeRequest.ProtoReflect.Descriptor instead.
func (*GetWorkTimeRequest) Descriptor() ([]byte, []int) {
	return file_tkd_roster_v1_worktime_proto_rawDescGZIP(), []int{3}
}

func (x *GetWorkTimeRequest) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *GetWorkTimeRequest) GetReadMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.ReadMask
	}
	return nil
}

type UserWorkTime struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Current       *WorkTime              `protobuf:"bytes,2,opt,name=current,proto3" json:"current,omitempty"`
	History       []*WorkTime            `protobuf:"bytes,3,rep,name=history,proto3" json:"history,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserWorkTime) Reset() {
	*x = UserWorkTime{}
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserWorkTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserWorkTime) ProtoMessage() {}

func (x *UserWorkTime) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserWorkTime.ProtoReflect.Descriptor instead.
func (*UserWorkTime) Descriptor() ([]byte, []int) {
	return file_tkd_roster_v1_worktime_proto_rawDescGZIP(), []int{4}
}

func (x *UserWorkTime) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserWorkTime) GetCurrent() *WorkTime {
	if x != nil {
		return x.Current
	}
	return nil
}

func (x *UserWorkTime) GetHistory() []*WorkTime {
	if x != nil {
		return x.History
	}
	return nil
}

type GetWorkTimeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Results       []*UserWorkTime        `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetWorkTimeResponse) Reset() {
	*x = GetWorkTimeResponse{}
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWorkTimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkTimeResponse) ProtoMessage() {}

func (x *GetWorkTimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkTimeResponse.ProtoReflect.Descriptor instead.
func (*GetWorkTimeResponse) Descriptor() ([]byte, []int) {
	return file_tkd_roster_v1_worktime_proto_rawDescGZIP(), []int{5}
}

func (x *GetWorkTimeResponse) GetResults() []*UserWorkTime {
	if x != nil {
		return x.Results
	}
	return nil
}

type SumForUsers struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserIds       []string               `protobuf:"bytes,1,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SumForUsers) Reset() {
	*x = SumForUsers{}
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SumForUsers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumForUsers) ProtoMessage() {}

func (x *SumForUsers) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumForUsers.ProtoReflect.Descriptor instead.
func (*SumForUsers) Descriptor() ([]byte, []int) {
	return file_tkd_roster_v1_worktime_proto_rawDescGZIP(), []int{6}
}

func (x *SumForUsers) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

type GetVacationCreditsLeftRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ForUsers      *SumForUsers           `protobuf:"bytes,1,opt,name=for_users,json=forUsers,proto3" json:"for_users,omitempty"`
	Until         *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=until,proto3" json:"until,omitempty"`
	Analyze       bool                   `protobuf:"varint,3,opt,name=analyze,proto3" json:"analyze,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetVacationCreditsLeftRequest) Reset() {
	*x = GetVacationCreditsLeftRequest{}
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetVacationCreditsLeftRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVacationCreditsLeftRequest) ProtoMessage() {}

func (x *GetVacationCreditsLeftRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVacationCreditsLeftRequest.ProtoReflect.Descriptor instead.
func (*GetVacationCreditsLeftRequest) Descriptor() ([]byte, []int) {
	return file_tkd_roster_v1_worktime_proto_rawDescGZIP(), []int{7}
}

func (x *GetVacationCreditsLeftRequest) GetForUsers() *SumForUsers {
	if x != nil {
		return x.ForUsers
	}
	return nil
}

func (x *GetVacationCreditsLeftRequest) GetUntil() *timestamppb.Timestamp {
	if x != nil {
		return x.Until
	}
	return nil
}

func (x *GetVacationCreditsLeftRequest) GetAnalyze() bool {
	if x != nil {
		return x.Analyze
	}
	return false
}

type AnalyzeVacationSum struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	WorkTime            *WorkTime              `protobuf:"bytes,1,opt,name=work_time,json=workTime,proto3" json:"work_time,omitempty"`
	EndsAt              *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=ends_at,json=endsAt,proto3" json:"ends_at,omitempty"`
	NumberOfDays        float64                `protobuf:"fixed64,3,opt,name=number_of_days,json=numberOfDays,proto3" json:"number_of_days,omitempty"`
	VacationWeeksPerDay float64                `protobuf:"fixed64,4,opt,name=vacation_weeks_per_day,json=vacationWeeksPerDay,proto3" json:"vacation_weeks_per_day,omitempty"`
	VacationPerWorkTime *durationpb.Duration   `protobuf:"bytes,5,opt,name=vacation_per_work_time,json=vacationPerWorkTime,proto3" json:"vacation_per_work_time,omitempty"`
	Costs               []*OffTimeCosts        `protobuf:"bytes,6,rep,name=costs,proto3" json:"costs,omitempty"`
	CostsSum            *durationpb.Duration   `protobuf:"bytes,7,opt,name=costs_sum,json=costsSum,proto3" json:"costs_sum,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *AnalyzeVacationSum) Reset() {
	*x = AnalyzeVacationSum{}
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnalyzeVacationSum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzeVacationSum) ProtoMessage() {}

func (x *AnalyzeVacationSum) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzeVacationSum.ProtoReflect.Descriptor instead.
func (*AnalyzeVacationSum) Descriptor() ([]byte, []int) {
	return file_tkd_roster_v1_worktime_proto_rawDescGZIP(), []int{8}
}

func (x *AnalyzeVacationSum) GetWorkTime() *WorkTime {
	if x != nil {
		return x.WorkTime
	}
	return nil
}

func (x *AnalyzeVacationSum) GetEndsAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndsAt
	}
	return nil
}

func (x *AnalyzeVacationSum) GetNumberOfDays() float64 {
	if x != nil {
		return x.NumberOfDays
	}
	return 0
}

func (x *AnalyzeVacationSum) GetVacationWeeksPerDay() float64 {
	if x != nil {
		return x.VacationWeeksPerDay
	}
	return 0
}

func (x *AnalyzeVacationSum) GetVacationPerWorkTime() *durationpb.Duration {
	if x != nil {
		return x.VacationPerWorkTime
	}
	return nil
}

func (x *AnalyzeVacationSum) GetCosts() []*OffTimeCosts {
	if x != nil {
		return x.Costs
	}
	return nil
}

func (x *AnalyzeVacationSum) GetCostsSum() *durationpb.Duration {
	if x != nil {
		return x.CostsSum
	}
	return nil
}

type AnalyzeVacation struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Slices        []*AnalyzeVacationSum  `protobuf:"bytes,1,rep,name=slices,proto3" json:"slices,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AnalyzeVacation) Reset() {
	*x = AnalyzeVacation{}
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnalyzeVacation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalyzeVacation) ProtoMessage() {}

func (x *AnalyzeVacation) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalyzeVacation.ProtoReflect.Descriptor instead.
func (*AnalyzeVacation) Descriptor() ([]byte, []int) {
	return file_tkd_roster_v1_worktime_proto_rawDescGZIP(), []int{9}
}

func (x *AnalyzeVacation) GetSlices() []*AnalyzeVacationSum {
	if x != nil {
		return x.Slices
	}
	return nil
}

type UserVacationSum struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	UserId              string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	VacationCreditsLeft *durationpb.Duration   `protobuf:"bytes,2,opt,name=vacation_credits_left,json=vacationCreditsLeft,proto3" json:"vacation_credits_left,omitempty"`
	TimeOffCredits      *durationpb.Duration   `protobuf:"bytes,3,opt,name=time_off_credits,json=timeOffCredits,proto3" json:"time_off_credits,omitempty"`
	Analysis            *AnalyzeVacation       `protobuf:"bytes,4,opt,name=analysis,proto3" json:"analysis,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *UserVacationSum) Reset() {
	*x = UserVacationSum{}
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserVacationSum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserVacationSum) ProtoMessage() {}

func (x *UserVacationSum) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserVacationSum.ProtoReflect.Descriptor instead.
func (*UserVacationSum) Descriptor() ([]byte, []int) {
	return file_tkd_roster_v1_worktime_proto_rawDescGZIP(), []int{10}
}

func (x *UserVacationSum) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserVacationSum) GetVacationCreditsLeft() *durationpb.Duration {
	if x != nil {
		return x.VacationCreditsLeft
	}
	return nil
}

func (x *UserVacationSum) GetTimeOffCredits() *durationpb.Duration {
	if x != nil {
		return x.TimeOffCredits
	}
	return nil
}

func (x *UserVacationSum) GetAnalysis() *AnalyzeVacation {
	if x != nil {
		return x.Analysis
	}
	return nil
}

type GetVacationCreditsLeftResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Results       []*UserVacationSum     `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetVacationCreditsLeftResponse) Reset() {
	*x = GetVacationCreditsLeftResponse{}
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetVacationCreditsLeftResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVacationCreditsLeftResponse) ProtoMessage() {}

func (x *GetVacationCreditsLeftResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVacationCreditsLeftResponse.ProtoReflect.Descriptor instead.
func (*GetVacationCreditsLeftResponse) Descriptor() ([]byte, []int) {
	return file_tkd_roster_v1_worktime_proto_rawDescGZIP(), []int{11}
}

func (x *GetVacationCreditsLeftResponse) GetResults() []*UserVacationSum {
	if x != nil {
		return x.Results
	}
	return nil
}

type DeleteWorkTimeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ids           []string               `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteWorkTimeRequest) Reset() {
	*x = DeleteWorkTimeRequest{}
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteWorkTimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWorkTimeRequest) ProtoMessage() {}

func (x *DeleteWorkTimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWorkTimeRequest.ProtoReflect.Descriptor instead.
func (*DeleteWorkTimeRequest) Descriptor() ([]byte, []int) {
	return file_tkd_roster_v1_worktime_proto_rawDescGZIP(), []int{12}
}

func (x *DeleteWorkTimeRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type DeleteWorkTimeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteWorkTimeResponse) Reset() {
	*x = DeleteWorkTimeResponse{}
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteWorkTimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWorkTimeResponse) ProtoMessage() {}

func (x *DeleteWorkTimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWorkTimeResponse.ProtoReflect.Descriptor instead.
func (*DeleteWorkTimeResponse) Descriptor() ([]byte, []int) {
	return file_tkd_roster_v1_worktime_proto_rawDescGZIP(), []int{13}
}

type UpdateWorkTimeRequest struct {
	state                   protoimpl.MessageState `protogen:"open.v1"`
	Id                      string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ExcludeFromTimeTracking bool                   `protobuf:"varint,2,opt,name=exclude_from_time_tracking,json=excludeFromTimeTracking,proto3" json:"exclude_from_time_tracking,omitempty"`
	// Format: YYYY-MM-DD
	EndsWith      string                 `protobuf:"bytes,3,opt,name=ends_with,json=endsWith,proto3" json:"ends_with,omitempty"`
	FieldMask     *fieldmaskpb.FieldMask `protobuf:"bytes,4,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateWorkTimeRequest) Reset() {
	*x = UpdateWorkTimeRequest{}
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateWorkTimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWorkTimeRequest) ProtoMessage() {}

func (x *UpdateWorkTimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWorkTimeRequest.ProtoReflect.Descriptor instead.
func (*UpdateWorkTimeRequest) Descriptor() ([]byte, []int) {
	return file_tkd_roster_v1_worktime_proto_rawDescGZIP(), []int{14}
}

func (x *UpdateWorkTimeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateWorkTimeRequest) GetExcludeFromTimeTracking() bool {
	if x != nil {
		return x.ExcludeFromTimeTracking
	}
	return false
}

func (x *UpdateWorkTimeRequest) GetEndsWith() string {
	if x != nil {
		return x.EndsWith
	}
	return ""
}

func (x *UpdateWorkTimeRequest) GetFieldMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.FieldMask
	}
	return nil
}

type UpdateWorkTimeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Worktime      *WorkTime              `protobuf:"bytes,1,opt,name=worktime,proto3" json:"worktime,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateWorkTimeResponse) Reset() {
	*x = UpdateWorkTimeResponse{}
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateWorkTimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWorkTimeResponse) ProtoMessage() {}

func (x *UpdateWorkTimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_roster_v1_worktime_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWorkTimeResponse.ProtoReflect.Descriptor instead.
func (*UpdateWorkTimeResponse) Descriptor() ([]byte, []int) {
	return file_tkd_roster_v1_worktime_proto_rawDescGZIP(), []int{15}
}

func (x *UpdateWorkTimeResponse) GetWorktime() *WorkTime {
	if x != nil {
		return x.Worktime
	}
	return nil
}

var File_tkd_roster_v1_worktime_proto protoreflect.FileDescriptor

const file_tkd_roster_v1_worktime_proto_rawDesc = "" +
	"\n" +
	"\x1ctkd/roster/v1/worktime.proto\x12\rtkd.roster.v1\x1a\x1egoogle/protobuf/duration.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a google/protobuf/field_mask.proto\x1a\x1etkd/common/v1/descriptor.proto\x1a\x1btkd/roster/v1/offtime.proto\x1a\x1bbuf/validate/validate.proto\"\xab\x03\n" +
	"\bWorkTime\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12 \n" +
	"\auser_id\x18\x02 \x01(\tB\a\xfa\xf7\x18\x03\xc8\x01\x01R\x06userId\x12F\n" +
	"\rtime_per_week\x18\x03 \x01(\v2\x19.google.protobuf.DurationB\a\xfa\xf7\x18\x03\xc8\x01\x01R\vtimePerWeek\x12)\n" +
	"\x10applicable_after\x18\x04 \x01(\tR\x0fapplicableAfter\x12D\n" +
	"\x17vacation_weeks_per_year\x18\x05 \x01(\x02B\r\xfa\xf7\x18\t\xc8\x01\x01\x1a\x04\x104 \x00R\x14vacationWeeksPerYear\x12Z\n" +
	"\x1covertime_allowance_per_month\x18\x06 \x01(\v2\x19.google.protobuf.DurationR\x19overtimeAllowancePerMonth\x12;\n" +
	"\x1aexclude_from_time_tracking\x18\a \x01(\bR\x17excludeFromTimeTracking\x12\x1b\n" +
	"\tends_with\x18\b \x01(\tR\bendsWith\"U\n" +
	"\x12SetWorkTimeRequest\x12?\n" +
	"\n" +
	"work_times\x18\x01 \x03(\v2\x17.tkd.roster.v1.WorkTimeB\a\xfa\xf7\x18\x03\xc8\x01\x01R\tworkTimes\"M\n" +
	"\x13SetWorkTimeResponse\x126\n" +
	"\n" +
	"work_times\x18\x01 \x03(\v2\x17.tkd.roster.v1.WorkTimeR\tworkTimes\"h\n" +
	"\x12GetWorkTimeRequest\x12\x19\n" +
	"\buser_ids\x18\x01 \x03(\tR\auserIds\x127\n" +
	"\tread_mask\x18\x02 \x01(\v2\x1a.google.protobuf.FieldMaskR\breadMask\"\x8d\x01\n" +
	"\fUserWorkTime\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x121\n" +
	"\acurrent\x18\x02 \x01(\v2\x17.tkd.roster.v1.WorkTimeR\acurrent\x121\n" +
	"\ahistory\x18\x03 \x03(\v2\x17.tkd.roster.v1.WorkTimeR\ahistory\"L\n" +
	"\x13GetWorkTimeResponse\x125\n" +
	"\aresults\x18\x01 \x03(\v2\x1b.tkd.roster.v1.UserWorkTimeR\aresults\"(\n" +
	"\vSumForUsers\x12\x19\n" +
	"\buser_ids\x18\x01 \x03(\tR\auserIds\"\xa4\x01\n" +
	"\x1dGetVacationCreditsLeftRequest\x127\n" +
	"\tfor_users\x18\x01 \x01(\v2\x1a.tkd.roster.v1.SumForUsersR\bforUsers\x120\n" +
	"\x05until\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\x05until\x12\x18\n" +
	"\aanalyze\x18\x03 \x01(\bR\aanalyze\"\x95\x03\n" +
	"\x12AnalyzeVacationSum\x124\n" +
	"\twork_time\x18\x01 \x01(\v2\x17.tkd.roster.v1.WorkTimeR\bworkTime\x123\n" +
	"\aends_at\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\x06endsAt\x12$\n" +
	"\x0enumber_of_days\x18\x03 \x01(\x01R\fnumberOfDays\x123\n" +
	"\x16vacation_weeks_per_day\x18\x04 \x01(\x01R\x13vacationWeeksPerDay\x12N\n" +
	"\x16vacation_per_work_time\x18\x05 \x01(\v2\x19.google.protobuf.DurationR\x13vacationPerWorkTime\x121\n" +
	"\x05costs\x18\x06 \x03(\v2\x1b.tkd.roster.v1.OffTimeCostsR\x05costs\x126\n" +
	"\tcosts_sum\x18\a \x01(\v2\x19.google.protobuf.DurationR\bcostsSum\"L\n" +
	"\x0fAnalyzeVacation\x129\n" +
	"\x06slices\x18\x01 \x03(\v2!.tkd.roster.v1.AnalyzeVacationSumR\x06slices\"\xfa\x01\n" +
	"\x0fUserVacationSum\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12M\n" +
	"\x15vacation_credits_left\x18\x02 \x01(\v2\x19.google.protobuf.DurationR\x13vacationCreditsLeft\x12C\n" +
	"\x10time_off_credits\x18\x03 \x01(\v2\x19.google.protobuf.DurationR\x0etimeOffCredits\x12:\n" +
	"\banalysis\x18\x04 \x01(\v2\x1e.tkd.roster.v1.AnalyzeVacationR\banalysis\"Z\n" +
	"\x1eGetVacationCreditsLeftResponse\x128\n" +
	"\aresults\x18\x01 \x03(\v2\x1e.tkd.roster.v1.UserVacationSumR\aresults\")\n" +
	"\x15DeleteWorkTimeRequest\x12\x10\n" +
	"\x03ids\x18\x01 \x03(\tR\x03ids\"\x18\n" +
	"\x16DeleteWorkTimeResponse\"\xc5\x01\n" +
	"\x15UpdateWorkTimeRequest\x12\x17\n" +
	"\x02id\x18\x01 \x01(\tB\a\xfa\xf7\x18\x03\xc8\x01\x01R\x02id\x12;\n" +
	"\x1aexclude_from_time_tracking\x18\x02 \x01(\bR\x17excludeFromTimeTracking\x12\x1b\n" +
	"\tends_with\x18\x03 \x01(\tR\bendsWith\x129\n" +
	"\n" +
	"field_mask\x18\x04 \x01(\v2\x1a.google.protobuf.FieldMaskR\tfieldMask\"M\n" +
	"\x16UpdateWorkTimeResponse\x123\n" +
	"\bworktime\x18\x01 \x01(\v2\x17.tkd.roster.v1.WorkTimeR\bworktime2\xaa\x04\n" +
	"\x0fWorkTimeService\x12[\n" +
	"\vSetWorkTime\x12!.tkd.roster.v1.SetWorkTimeRequest\x1a\".tkd.roster.v1.SetWorkTimeResponse\"\x05\xb2~\x02\b\x02\x12[\n" +
	"\vGetWorkTime\x12!.tkd.roster.v1.GetWorkTimeRequest\x1a\".tkd.roster.v1.GetWorkTimeResponse\"\x05\xb2~\x02\b\x01\x12d\n" +
	"\x0eUpdateWorkTime\x12$.tkd.roster.v1.UpdateWorkTimeRequest\x1a%.tkd.roster.v1.UpdateWorkTimeResponse\"\x05\xb2~\x02\b\x02\x12|\n" +
	"\x16GetVacationCreditsLeft\x12,.tkd.roster.v1.GetVacationCreditsLeftRequest\x1a-.tkd.roster.v1.GetVacationCreditsLeftResponse\"\x05\xb2~\x02\b\x01\x12d\n" +
	"\x0eDeleteWorkTime\x12$.tkd.roster.v1.DeleteWorkTimeRequest\x1a%.tkd.roster.v1.DeleteWorkTimeResponse\"\x05\xb2~\x02\b\x02\x1a\x13\xba~\x10\n" +
	"\x0eroster_managerB\xbd\x01\n" +
	"\x11com.tkd.roster.v1B\rWorktimeProtoP\x01ZCgithub.com/tierklinik-dobersberg/apis/gen/go/tkd/roster/v1;rosterv1\xa2\x02\x03TRX\xaa\x02\rTkd.Roster.V1\xca\x02\rTkd\\Roster\\V1\xe2\x02\x19Tkd\\Roster\\V1\\GPBMetadata\xea\x02\x0fTkd::Roster::V1b\x06proto3"

var (
	file_tkd_roster_v1_worktime_proto_rawDescOnce sync.Once
	file_tkd_roster_v1_worktime_proto_rawDescData []byte
)

func file_tkd_roster_v1_worktime_proto_rawDescGZIP() []byte {
	file_tkd_roster_v1_worktime_proto_rawDescOnce.Do(func() {
		file_tkd_roster_v1_worktime_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_tkd_roster_v1_worktime_proto_rawDesc), len(file_tkd_roster_v1_worktime_proto_rawDesc)))
	})
	return file_tkd_roster_v1_worktime_proto_rawDescData
}

var file_tkd_roster_v1_worktime_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_tkd_roster_v1_worktime_proto_goTypes = []any{
	(*WorkTime)(nil),                       // 0: tkd.roster.v1.WorkTime
	(*SetWorkTimeRequest)(nil),             // 1: tkd.roster.v1.SetWorkTimeRequest
	(*SetWorkTimeResponse)(nil),            // 2: tkd.roster.v1.SetWorkTimeResponse
	(*GetWorkTimeRequest)(nil),             // 3: tkd.roster.v1.GetWorkTimeRequest
	(*UserWorkTime)(nil),                   // 4: tkd.roster.v1.UserWorkTime
	(*GetWorkTimeResponse)(nil),            // 5: tkd.roster.v1.GetWorkTimeResponse
	(*SumForUsers)(nil),                    // 6: tkd.roster.v1.SumForUsers
	(*GetVacationCreditsLeftRequest)(nil),  // 7: tkd.roster.v1.GetVacationCreditsLeftRequest
	(*AnalyzeVacationSum)(nil),             // 8: tkd.roster.v1.AnalyzeVacationSum
	(*AnalyzeVacation)(nil),                // 9: tkd.roster.v1.AnalyzeVacation
	(*UserVacationSum)(nil),                // 10: tkd.roster.v1.UserVacationSum
	(*GetVacationCreditsLeftResponse)(nil), // 11: tkd.roster.v1.GetVacationCreditsLeftResponse
	(*DeleteWorkTimeRequest)(nil),          // 12: tkd.roster.v1.DeleteWorkTimeRequest
	(*DeleteWorkTimeResponse)(nil),         // 13: tkd.roster.v1.DeleteWorkTimeResponse
	(*UpdateWorkTimeRequest)(nil),          // 14: tkd.roster.v1.UpdateWorkTimeRequest
	(*UpdateWorkTimeResponse)(nil),         // 15: tkd.roster.v1.UpdateWorkTimeResponse
	(*durationpb.Duration)(nil),            // 16: google.protobuf.Duration
	(*fieldmaskpb.FieldMask)(nil),          // 17: google.protobuf.FieldMask
	(*timestamppb.Timestamp)(nil),          // 18: google.protobuf.Timestamp
	(*OffTimeCosts)(nil),                   // 19: tkd.roster.v1.OffTimeCosts
}
var file_tkd_roster_v1_worktime_proto_depIdxs = []int32{
	16, // 0: tkd.roster.v1.WorkTime.time_per_week:type_name -> google.protobuf.Duration
	16, // 1: tkd.roster.v1.WorkTime.overtime_allowance_per_month:type_name -> google.protobuf.Duration
	0,  // 2: tkd.roster.v1.SetWorkTimeRequest.work_times:type_name -> tkd.roster.v1.WorkTime
	0,  // 3: tkd.roster.v1.SetWorkTimeResponse.work_times:type_name -> tkd.roster.v1.WorkTime
	17, // 4: tkd.roster.v1.GetWorkTimeRequest.read_mask:type_name -> google.protobuf.FieldMask
	0,  // 5: tkd.roster.v1.UserWorkTime.current:type_name -> tkd.roster.v1.WorkTime
	0,  // 6: tkd.roster.v1.UserWorkTime.history:type_name -> tkd.roster.v1.WorkTime
	4,  // 7: tkd.roster.v1.GetWorkTimeResponse.results:type_name -> tkd.roster.v1.UserWorkTime
	6,  // 8: tkd.roster.v1.GetVacationCreditsLeftRequest.for_users:type_name -> tkd.roster.v1.SumForUsers
	18, // 9: tkd.roster.v1.GetVacationCreditsLeftRequest.until:type_name -> google.protobuf.Timestamp
	0,  // 10: tkd.roster.v1.AnalyzeVacationSum.work_time:type_name -> tkd.roster.v1.WorkTime
	18, // 11: tkd.roster.v1.AnalyzeVacationSum.ends_at:type_name -> google.protobuf.Timestamp
	16, // 12: tkd.roster.v1.AnalyzeVacationSum.vacation_per_work_time:type_name -> google.protobuf.Duration
	19, // 13: tkd.roster.v1.AnalyzeVacationSum.costs:type_name -> tkd.roster.v1.OffTimeCosts
	16, // 14: tkd.roster.v1.AnalyzeVacationSum.costs_sum:type_name -> google.protobuf.Duration
	8,  // 15: tkd.roster.v1.AnalyzeVacation.slices:type_name -> tkd.roster.v1.AnalyzeVacationSum
	16, // 16: tkd.roster.v1.UserVacationSum.vacation_credits_left:type_name -> google.protobuf.Duration
	16, // 17: tkd.roster.v1.UserVacationSum.time_off_credits:type_name -> google.protobuf.Duration
	9,  // 18: tkd.roster.v1.UserVacationSum.analysis:type_name -> tkd.roster.v1.AnalyzeVacation
	10, // 19: tkd.roster.v1.GetVacationCreditsLeftResponse.results:type_name -> tkd.roster.v1.UserVacationSum
	17, // 20: tkd.roster.v1.UpdateWorkTimeRequest.field_mask:type_name -> google.protobuf.FieldMask
	0,  // 21: tkd.roster.v1.UpdateWorkTimeResponse.worktime:type_name -> tkd.roster.v1.WorkTime
	1,  // 22: tkd.roster.v1.WorkTimeService.SetWorkTime:input_type -> tkd.roster.v1.SetWorkTimeRequest
	3,  // 23: tkd.roster.v1.WorkTimeService.GetWorkTime:input_type -> tkd.roster.v1.GetWorkTimeRequest
	14, // 24: tkd.roster.v1.WorkTimeService.UpdateWorkTime:input_type -> tkd.roster.v1.UpdateWorkTimeRequest
	7,  // 25: tkd.roster.v1.WorkTimeService.GetVacationCreditsLeft:input_type -> tkd.roster.v1.GetVacationCreditsLeftRequest
	12, // 26: tkd.roster.v1.WorkTimeService.DeleteWorkTime:input_type -> tkd.roster.v1.DeleteWorkTimeRequest
	2,  // 27: tkd.roster.v1.WorkTimeService.SetWorkTime:output_type -> tkd.roster.v1.SetWorkTimeResponse
	5,  // 28: tkd.roster.v1.WorkTimeService.GetWorkTime:output_type -> tkd.roster.v1.GetWorkTimeResponse
	15, // 29: tkd.roster.v1.WorkTimeService.UpdateWorkTime:output_type -> tkd.roster.v1.UpdateWorkTimeResponse
	11, // 30: tkd.roster.v1.WorkTimeService.GetVacationCreditsLeft:output_type -> tkd.roster.v1.GetVacationCreditsLeftResponse
	13, // 31: tkd.roster.v1.WorkTimeService.DeleteWorkTime:output_type -> tkd.roster.v1.DeleteWorkTimeResponse
	27, // [27:32] is the sub-list for method output_type
	22, // [22:27] is the sub-list for method input_type
	22, // [22:22] is the sub-list for extension type_name
	22, // [22:22] is the sub-list for extension extendee
	0,  // [0:22] is the sub-list for field type_name
}

func init() { file_tkd_roster_v1_worktime_proto_init() }
func file_tkd_roster_v1_worktime_proto_init() {
	if File_tkd_roster_v1_worktime_proto != nil {
		return
	}
	file_tkd_roster_v1_offtime_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_tkd_roster_v1_worktime_proto_rawDesc), len(file_tkd_roster_v1_worktime_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tkd_roster_v1_worktime_proto_goTypes,
		DependencyIndexes: file_tkd_roster_v1_worktime_proto_depIdxs,
		MessageInfos:      file_tkd_roster_v1_worktime_proto_msgTypes,
	}.Build()
	File_tkd_roster_v1_worktime_proto = out.File
	file_tkd_roster_v1_worktime_proto_goTypes = nil
	file_tkd_roster_v1_worktime_proto_depIdxs = nil
}
