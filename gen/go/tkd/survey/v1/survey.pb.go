// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: tkd/survey/v1/survey.proto

package surveyv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/tierklinik-dobersberg/apis/gen/go/tkd/common/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

type Survey struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name is the name of the survey and must be unique.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// DisplayName holds a human readable display name for the survey.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Definition holds the surveyjs.io JSON defintion.
	Definition *structpb.Struct `protobuf:"bytes,3,opt,name=definition,proto3" json:"definition,omitempty"`
	// Schema holds the JSON schema used to validate
	// answers.
	Schema *structpb.Struct `protobuf:"bytes,10,opt,name=schema,proto3" json:"schema,omitempty"`
	// CreateTime is the time when the survey has been created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Creator is the ID of the user that created the survey
	Creator string `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	// ValidUntil holds the time until the survey is valid and
	// can be filled out.
	ValidUntil *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=valid_until,json=validUntil,proto3" json:"valid_until,omitempty"`
	// Published indicates whether or not the survey is ready
	// and can be filled out.
	Published bool `protobuf:"varint,7,opt,name=published,proto3" json:"published,omitempty"`
	// EligibleUserIds holds a list of user ids that are permitted
	// for this survey.
	//
	// If eligible_user_ids and eligible_role_ids are both empty,
	// every user is allowed to perform the survey.
	EligibleUserIds []string `protobuf:"bytes,8,rep,name=eligible_user_ids,json=eligibleUserIds,proto3" json:"eligible_user_ids,omitempty"`
	// EligibleRoleIds holds a list of role ids that are permitted
	// for this survey.
	//
	// If eligible_user_ids and eligible_role_ids are both empty,
	// every user is allowed to perform the survey.
	EligibleRoleIds []string `protobuf:"bytes,9,rep,name=eligible_role_ids,json=eligibleRoleIds,proto3" json:"eligible_role_ids,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Survey) Reset() {
	*x = Survey{}
	mi := &file_tkd_survey_v1_survey_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Survey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Survey) ProtoMessage() {}

func (x *Survey) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_survey_v1_survey_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Survey.ProtoReflect.Descriptor instead.
func (*Survey) Descriptor() ([]byte, []int) {
	return file_tkd_survey_v1_survey_proto_rawDescGZIP(), []int{0}
}

func (x *Survey) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Survey) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Survey) GetDefinition() *structpb.Struct {
	if x != nil {
		return x.Definition
	}
	return nil
}

func (x *Survey) GetSchema() *structpb.Struct {
	if x != nil {
		return x.Schema
	}
	return nil
}

func (x *Survey) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Survey) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *Survey) GetValidUntil() *timestamppb.Timestamp {
	if x != nil {
		return x.ValidUntil
	}
	return nil
}

func (x *Survey) GetPublished() bool {
	if x != nil {
		return x.Published
	}
	return false
}

func (x *Survey) GetEligibleUserIds() []string {
	if x != nil {
		return x.EligibleUserIds
	}
	return nil
}

func (x *Survey) GetEligibleRoleIds() []string {
	if x != nil {
		return x.EligibleRoleIds
	}
	return nil
}

type CreateSurveyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Survey        *Survey                `protobuf:"bytes,1,opt,name=survey,proto3" json:"survey,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateSurveyRequest) Reset() {
	*x = CreateSurveyRequest{}
	mi := &file_tkd_survey_v1_survey_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSurveyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSurveyRequest) ProtoMessage() {}

func (x *CreateSurveyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_survey_v1_survey_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSurveyRequest.ProtoReflect.Descriptor instead.
func (*CreateSurveyRequest) Descriptor() ([]byte, []int) {
	return file_tkd_survey_v1_survey_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSurveyRequest) GetSurvey() *Survey {
	if x != nil {
		return x.Survey
	}
	return nil
}

type DeleteSurveyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SurveyName    string                 `protobuf:"bytes,1,opt,name=survey_name,json=surveyName,proto3" json:"survey_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteSurveyRequest) Reset() {
	*x = DeleteSurveyRequest{}
	mi := &file_tkd_survey_v1_survey_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteSurveyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSurveyRequest) ProtoMessage() {}

func (x *DeleteSurveyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_survey_v1_survey_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSurveyRequest.ProtoReflect.Descriptor instead.
func (*DeleteSurveyRequest) Descriptor() ([]byte, []int) {
	return file_tkd_survey_v1_survey_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteSurveyRequest) GetSurveyName() string {
	if x != nil {
		return x.SurveyName
	}
	return ""
}

type ListSurveyDefinitionsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListSurveyDefinitionsRequest) Reset() {
	*x = ListSurveyDefinitionsRequest{}
	mi := &file_tkd_survey_v1_survey_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSurveyDefinitionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSurveyDefinitionsRequest) ProtoMessage() {}

func (x *ListSurveyDefinitionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_survey_v1_survey_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSurveyDefinitionsRequest.ProtoReflect.Descriptor instead.
func (*ListSurveyDefinitionsRequest) Descriptor() ([]byte, []int) {
	return file_tkd_survey_v1_survey_proto_rawDescGZIP(), []int{3}
}

type ListSurveyDefinitionsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Surveys       []*Survey              `protobuf:"bytes,1,rep,name=surveys,proto3" json:"surveys,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListSurveyDefinitionsResponse) Reset() {
	*x = ListSurveyDefinitionsResponse{}
	mi := &file_tkd_survey_v1_survey_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSurveyDefinitionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSurveyDefinitionsResponse) ProtoMessage() {}

func (x *ListSurveyDefinitionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_survey_v1_survey_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSurveyDefinitionsResponse.ProtoReflect.Descriptor instead.
func (*ListSurveyDefinitionsResponse) Descriptor() ([]byte, []int) {
	return file_tkd_survey_v1_survey_proto_rawDescGZIP(), []int{4}
}

func (x *ListSurveyDefinitionsResponse) GetSurveys() []*Survey {
	if x != nil {
		return x.Surveys
	}
	return nil
}

type SaveSurveyVoteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SurveyName    string                 `protobuf:"bytes,1,opt,name=survey_name,json=surveyName,proto3" json:"survey_name,omitempty"`
	Values        *structpb.Struct       `protobuf:"bytes,2,opt,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SaveSurveyVoteRequest) Reset() {
	*x = SaveSurveyVoteRequest{}
	mi := &file_tkd_survey_v1_survey_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SaveSurveyVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveSurveyVoteRequest) ProtoMessage() {}

func (x *SaveSurveyVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_survey_v1_survey_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveSurveyVoteRequest.ProtoReflect.Descriptor instead.
func (*SaveSurveyVoteRequest) Descriptor() ([]byte, []int) {
	return file_tkd_survey_v1_survey_proto_rawDescGZIP(), []int{5}
}

func (x *SaveSurveyVoteRequest) GetSurveyName() string {
	if x != nil {
		return x.SurveyName
	}
	return ""
}

func (x *SaveSurveyVoteRequest) GetValues() *structpb.Struct {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_tkd_survey_v1_survey_proto protoreflect.FileDescriptor

var file_tkd_survey_v1_survey_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x6b, 0x64, 0x2f, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x74, 0x6b,
	0x64, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x74, 0x6b, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x03, 0x0a, 0x06, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x12,
	0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0xf7, 0x18, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x0c,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0xf7, 0x18, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x07, 0xfa, 0xf7, 0x18, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0a,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x42, 0x07, 0xfa, 0xf7, 0x18, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x06, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x3b, 0x0a, 0x0b, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x5f, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x6c, 0x69, 0x67, 0x69, 0x62,
	0x6c, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0f, 0x65, 0x6c, 0x69, 0x67, 0x69, 0x62, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x6c, 0x69, 0x67, 0x69, 0x62, 0x6c, 0x65, 0x5f, 0x72,
	0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x65,
	0x6c, 0x69, 0x67, 0x69, 0x62, 0x6c, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x22, 0x4d,
	0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x73, 0x75, 0x72, 0x76,
	0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x42, 0x07, 0xfa, 0xf7,
	0x18, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x06, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x22, 0x3f, 0x0a,
	0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0b, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0xf7, 0x18, 0x03, 0xc8,
	0x01, 0x01, 0x52, 0x0a, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x1e,
	0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x50,
	0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2f, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x07, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x73,
	0x22, 0x72, 0x0a, 0x15, 0x53, 0x61, 0x76, 0x65, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x56, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0b, 0x73, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0xf7, 0x18, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0a, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x32, 0x86, 0x03, 0x0a, 0x0d, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x12, 0x22, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x73, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x6b, 0x64,
	0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x72, 0x76, 0x65,
	0x79, 0x22, 0x05, 0xb2, 0x7e, 0x02, 0x08, 0x01, 0x12, 0x51, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x12, 0x22, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x73,
	0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x75, 0x72, 0x76, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x05, 0xb2, 0x7e, 0x02, 0x08, 0x01, 0x12, 0x79, 0x0a, 0x15, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2b, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x44,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2c, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x05, 0xb2, 0x7e, 0x02, 0x08, 0x01, 0x12, 0x55, 0x0a, 0x0e, 0x53, 0x61, 0x76, 0x65, 0x53, 0x75,
	0x72, 0x76, 0x65, 0x79, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x24, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x73,
	0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x53, 0x75, 0x72,
	0x76, 0x65, 0x79, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x05, 0xb2, 0x7e, 0x02, 0x08, 0x01, 0x42, 0xbb, 0x01,
	0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79,
	0x2e, 0x76, 0x31, 0x42, 0x0b, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x69, 0x65, 0x72, 0x6b, 0x6c, 0x69, 0x6e, 0x69, 0x6b, 0x2d, 0x64, 0x6f, 0x62, 0x65, 0x72, 0x73,
	0x62, 0x65, 0x72, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x74, 0x6b, 0x64, 0x2f, 0x73, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x73,
	0x75, 0x72, 0x76, 0x65, 0x79, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x53, 0x58, 0xaa, 0x02, 0x0d,
	0x54, 0x6b, 0x64, 0x2e, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d,
	0x54, 0x6b, 0x64, 0x5c, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19,
	0x54, 0x6b, 0x64, 0x5c, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x54, 0x6b, 0x64, 0x3a,
	0x3a, 0x53, 0x75, 0x72, 0x76, 0x65, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_tkd_survey_v1_survey_proto_rawDescOnce sync.Once
	file_tkd_survey_v1_survey_proto_rawDescData = file_tkd_survey_v1_survey_proto_rawDesc
)

func file_tkd_survey_v1_survey_proto_rawDescGZIP() []byte {
	file_tkd_survey_v1_survey_proto_rawDescOnce.Do(func() {
		file_tkd_survey_v1_survey_proto_rawDescData = protoimpl.X.CompressGZIP(file_tkd_survey_v1_survey_proto_rawDescData)
	})
	return file_tkd_survey_v1_survey_proto_rawDescData
}

var file_tkd_survey_v1_survey_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_tkd_survey_v1_survey_proto_goTypes = []any{
	(*Survey)(nil),                        // 0: tkd.survey.v1.Survey
	(*CreateSurveyRequest)(nil),           // 1: tkd.survey.v1.CreateSurveyRequest
	(*DeleteSurveyRequest)(nil),           // 2: tkd.survey.v1.DeleteSurveyRequest
	(*ListSurveyDefinitionsRequest)(nil),  // 3: tkd.survey.v1.ListSurveyDefinitionsRequest
	(*ListSurveyDefinitionsResponse)(nil), // 4: tkd.survey.v1.ListSurveyDefinitionsResponse
	(*SaveSurveyVoteRequest)(nil),         // 5: tkd.survey.v1.SaveSurveyVoteRequest
	(*structpb.Struct)(nil),               // 6: google.protobuf.Struct
	(*timestamppb.Timestamp)(nil),         // 7: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                 // 8: google.protobuf.Empty
}
var file_tkd_survey_v1_survey_proto_depIdxs = []int32{
	6,  // 0: tkd.survey.v1.Survey.definition:type_name -> google.protobuf.Struct
	6,  // 1: tkd.survey.v1.Survey.schema:type_name -> google.protobuf.Struct
	7,  // 2: tkd.survey.v1.Survey.create_time:type_name -> google.protobuf.Timestamp
	7,  // 3: tkd.survey.v1.Survey.valid_until:type_name -> google.protobuf.Timestamp
	0,  // 4: tkd.survey.v1.CreateSurveyRequest.survey:type_name -> tkd.survey.v1.Survey
	0,  // 5: tkd.survey.v1.ListSurveyDefinitionsResponse.surveys:type_name -> tkd.survey.v1.Survey
	6,  // 6: tkd.survey.v1.SaveSurveyVoteRequest.values:type_name -> google.protobuf.Struct
	1,  // 7: tkd.survey.v1.SurveyService.CreateSurvey:input_type -> tkd.survey.v1.CreateSurveyRequest
	2,  // 8: tkd.survey.v1.SurveyService.DeleteSurvey:input_type -> tkd.survey.v1.DeleteSurveyRequest
	3,  // 9: tkd.survey.v1.SurveyService.ListSurveyDefinitions:input_type -> tkd.survey.v1.ListSurveyDefinitionsRequest
	5,  // 10: tkd.survey.v1.SurveyService.SaveSurveyVote:input_type -> tkd.survey.v1.SaveSurveyVoteRequest
	0,  // 11: tkd.survey.v1.SurveyService.CreateSurvey:output_type -> tkd.survey.v1.Survey
	8,  // 12: tkd.survey.v1.SurveyService.DeleteSurvey:output_type -> google.protobuf.Empty
	4,  // 13: tkd.survey.v1.SurveyService.ListSurveyDefinitions:output_type -> tkd.survey.v1.ListSurveyDefinitionsResponse
	8,  // 14: tkd.survey.v1.SurveyService.SaveSurveyVote:output_type -> google.protobuf.Empty
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_tkd_survey_v1_survey_proto_init() }
func file_tkd_survey_v1_survey_proto_init() {
	if File_tkd_survey_v1_survey_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tkd_survey_v1_survey_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tkd_survey_v1_survey_proto_goTypes,
		DependencyIndexes: file_tkd_survey_v1_survey_proto_depIdxs,
		MessageInfos:      file_tkd_survey_v1_survey_proto_msgTypes,
	}.Build()
	File_tkd_survey_v1_survey_proto = out.File
	file_tkd_survey_v1_survey_proto_rawDesc = nil
	file_tkd_survey_v1_survey_proto_goTypes = nil
	file_tkd_survey_v1_survey_proto_depIdxs = nil
}
