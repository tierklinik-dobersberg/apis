// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: tkd/orthanc_bridge/v1/orthanc-bridge.proto

package orthanc_bridgev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type SearchLevel int32

const (
	SearchLevel_SEARCH_LEVEL_UNSPECIFIED SearchLevel = 0
	SearchLevel_SEARCH_LEVEL_PATIENT     SearchLevel = 1
	SearchLevel_SEARCH_LEVEL_STUDY       SearchLevel = 2
	SearchLevel_SEARCH_LEVEL_SERIES      SearchLevel = 3
	SearchLevel_SEARCH_LEVEL_INSTANCE    SearchLevel = 4
)

// Enum value maps for SearchLevel.
var (
	SearchLevel_name = map[int32]string{
		0: "SEARCH_LEVEL_UNSPECIFIED",
		1: "SEARCH_LEVEL_PATIENT",
		2: "SEARCH_LEVEL_STUDY",
		3: "SEARCH_LEVEL_SERIES",
		4: "SEARCH_LEVEL_INSTANCE",
	}
	SearchLevel_value = map[string]int32{
		"SEARCH_LEVEL_UNSPECIFIED": 0,
		"SEARCH_LEVEL_PATIENT":     1,
		"SEARCH_LEVEL_STUDY":       2,
		"SEARCH_LEVEL_SERIES":      3,
		"SEARCH_LEVEL_INSTANCE":    4,
	}
)

func (x SearchLevel) Enum() *SearchLevel {
	p := new(SearchLevel)
	*p = x
	return p
}

func (x SearchLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SearchLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_enumTypes[0].Descriptor()
}

func (SearchLevel) Type() protoreflect.EnumType {
	return &file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_enumTypes[0]
}

func (x SearchLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SearchLevel.Descriptor instead.
func (SearchLevel) EnumDescriptor() ([]byte, []int) {
	return file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_rawDescGZIP(), []int{0}
}

type SearchPatientsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CaseSensitive     bool     `protobuf:"varint,1,opt,name=case_sensitive,json=caseSensitive,proto3" json:"case_sensitive,omitempty"`
	Limit             int64    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Since             int64    `protobuf:"varint,3,opt,name=since,proto3" json:"since,omitempty"`
	RequestedTags     []string `protobuf:"bytes,4,rep,name=requested_tags,json=requestedTags,proto3" json:"requested_tags,omitempty"`
	PatientName       string   `protobuf:"bytes,5,opt,name=patient_name,json=patientName,proto3" json:"patient_name,omitempty"`
	ResponsiblePerson string   `protobuf:"bytes,6,opt,name=responsible_person,json=responsiblePerson,proto3" json:"responsible_person,omitempty"`
	// The date to search for. Might either be in one of the following formats:
	//   - YYYY : search withing a specified year
	//   - YYYY-MM: search within a specified month/year
	//   - YYYY-MM-DD: search at a specifc date.
	Date string `protobuf:"bytes,7,opt,name=date,proto3" json:"date,omitempty"`
	// The orthanc instance that should be searched. If unspecified
	// orthanc-bridge will try to determine the best instance to search
	// given date is set.
	Instance string `protobuf:"bytes,8,opt,name=instance,proto3" json:"instance,omitempty"`
}

func (x *SearchPatientsRequest) Reset() {
	*x = SearchPatientsRequest{}
	mi := &file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchPatientsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchPatientsRequest) ProtoMessage() {}

func (x *SearchPatientsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchPatientsRequest.ProtoReflect.Descriptor instead.
func (*SearchPatientsRequest) Descriptor() ([]byte, []int) {
	return file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_rawDescGZIP(), []int{0}
}

func (x *SearchPatientsRequest) GetCaseSensitive() bool {
	if x != nil {
		return x.CaseSensitive
	}
	return false
}

func (x *SearchPatientsRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *SearchPatientsRequest) GetSince() int64 {
	if x != nil {
		return x.Since
	}
	return 0
}

func (x *SearchPatientsRequest) GetRequestedTags() []string {
	if x != nil {
		return x.RequestedTags
	}
	return nil
}

func (x *SearchPatientsRequest) GetPatientName() string {
	if x != nil {
		return x.PatientName
	}
	return ""
}

func (x *SearchPatientsRequest) GetResponsiblePerson() string {
	if x != nil {
		return x.ResponsiblePerson
	}
	return ""
}

func (x *SearchPatientsRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *SearchPatientsRequest) GetInstance() string {
	if x != nil {
		return x.Instance
	}
	return ""
}

type SearchPatientsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MainTags *PatientMainTags `protobuf:"bytes,2,opt,name=main_tags,json=mainTags,proto3" json:"main_tags,omitempty"`
}

func (x *SearchPatientsResponse) Reset() {
	*x = SearchPatientsResponse{}
	mi := &file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchPatientsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchPatientsResponse) ProtoMessage() {}

func (x *SearchPatientsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchPatientsResponse.ProtoReflect.Descriptor instead.
func (*SearchPatientsResponse) Descriptor() ([]byte, []int) {
	return file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_rawDescGZIP(), []int{1}
}

func (x *SearchPatientsResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SearchPatientsResponse) GetMainTags() *PatientMainTags {
	if x != nil {
		return x.MainTags
	}
	return nil
}

type PatientMainTags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientName      string   `protobuf:"bytes,1,opt,name=patient_name,json=patientName,proto3" json:"patient_name,omitempty"`
	PatientId        string   `protobuf:"bytes,2,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`
	PatientBirthDate string   `protobuf:"bytes,3,opt,name=patient_birth_date,json=patientBirthDate,proto3" json:"patient_birth_date,omitempty"` // YYYY-MM-DD
	PatientSex       string   `protobuf:"bytes,4,opt,name=patient_sex,json=patientSex,proto3" json:"patient_sex,omitempty"`
	OtherPatientIds  []string `protobuf:"bytes,5,rep,name=other_patient_ids,json=otherPatientIds,proto3" json:"other_patient_ids,omitempty"`
}

func (x *PatientMainTags) Reset() {
	*x = PatientMainTags{}
	mi := &file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PatientMainTags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatientMainTags) ProtoMessage() {}

func (x *PatientMainTags) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatientMainTags.ProtoReflect.Descriptor instead.
func (*PatientMainTags) Descriptor() ([]byte, []int) {
	return file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_rawDescGZIP(), []int{2}
}

func (x *PatientMainTags) GetPatientName() string {
	if x != nil {
		return x.PatientName
	}
	return ""
}

func (x *PatientMainTags) GetPatientId() string {
	if x != nil {
		return x.PatientId
	}
	return ""
}

func (x *PatientMainTags) GetPatientBirthDate() string {
	if x != nil {
		return x.PatientBirthDate
	}
	return ""
}

func (x *PatientMainTags) GetPatientSex() string {
	if x != nil {
		return x.PatientSex
	}
	return ""
}

func (x *PatientMainTags) GetOtherPatientIds() []string {
	if x != nil {
		return x.OtherPatientIds
	}
	return nil
}

type StudyMainTags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time                          *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"` // from StudyDate and StudyTime
	StudyId                       string                 `protobuf:"bytes,2,opt,name=study_id,json=studyId,proto3" json:"study_id,omitempty"`
	StudyDescription              string                 `protobuf:"bytes,3,opt,name=study_description,json=studyDescription,proto3" json:"study_description,omitempty"`
	AccessionNumber               string                 `protobuf:"bytes,4,opt,name=accession_number,json=accessionNumber,proto3" json:"accession_number,omitempty"`
	StudyInstanceUid              string                 `protobuf:"bytes,5,opt,name=study_instance_uid,json=studyInstanceUid,proto3" json:"study_instance_uid,omitempty"`
	RequestedPrecedureDescription string                 `protobuf:"bytes,6,opt,name=requested_precedure_description,json=requestedPrecedureDescription,proto3" json:"requested_precedure_description,omitempty"`
	InstitutionName               string                 `protobuf:"bytes,7,opt,name=institution_name,json=institutionName,proto3" json:"institution_name,omitempty"`
	RequestingPhysician           string                 `protobuf:"bytes,8,opt,name=requesting_physician,json=requestingPhysician,proto3" json:"requesting_physician,omitempty"`
	RefferingPysicianName         string                 `protobuf:"bytes,9,opt,name=reffering_pysician_name,json=refferingPysicianName,proto3" json:"reffering_pysician_name,omitempty"`
}

func (x *StudyMainTags) Reset() {
	*x = StudyMainTags{}
	mi := &file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StudyMainTags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudyMainTags) ProtoMessage() {}

func (x *StudyMainTags) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudyMainTags.ProtoReflect.Descriptor instead.
func (*StudyMainTags) Descriptor() ([]byte, []int) {
	return file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_rawDescGZIP(), []int{3}
}

func (x *StudyMainTags) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *StudyMainTags) GetStudyId() string {
	if x != nil {
		return x.StudyId
	}
	return ""
}

func (x *StudyMainTags) GetStudyDescription() string {
	if x != nil {
		return x.StudyDescription
	}
	return ""
}

func (x *StudyMainTags) GetAccessionNumber() string {
	if x != nil {
		return x.AccessionNumber
	}
	return ""
}

func (x *StudyMainTags) GetStudyInstanceUid() string {
	if x != nil {
		return x.StudyInstanceUid
	}
	return ""
}

func (x *StudyMainTags) GetRequestedPrecedureDescription() string {
	if x != nil {
		return x.RequestedPrecedureDescription
	}
	return ""
}

func (x *StudyMainTags) GetInstitutionName() string {
	if x != nil {
		return x.InstitutionName
	}
	return ""
}

func (x *StudyMainTags) GetRequestingPhysician() string {
	if x != nil {
		return x.RequestingPhysician
	}
	return ""
}

func (x *StudyMainTags) GetRefferingPysicianName() string {
	if x != nil {
		return x.RefferingPysicianName
	}
	return ""
}

type SeriesMainTags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time                              *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"` // from SeriesDate and SeriesTime
	Modality                          string                 `protobuf:"bytes,2,opt,name=modality,proto3" json:"modality,omitempty"`
	Manufacturer                      string                 `protobuf:"bytes,3,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	StationName                       string                 `protobuf:"bytes,4,opt,name=station_name,json=stationName,proto3" json:"station_name,omitempty"`
	SeriesDescription                 string                 `protobuf:"bytes,5,opt,name=series_description,json=seriesDescription,proto3" json:"series_description,omitempty"`
	BodyPartExamined                  string                 `protobuf:"bytes,6,opt,name=body_part_examined,json=bodyPartExamined,proto3" json:"body_part_examined,omitempty"`
	SequenceName                      string                 `protobuf:"bytes,7,opt,name=sequence_name,json=sequenceName,proto3" json:"sequence_name,omitempty"`
	ProtocolName                      string                 `protobuf:"bytes,8,opt,name=protocol_name,json=protocolName,proto3" json:"protocol_name,omitempty"`
	SeriesNumber                      string                 `protobuf:"bytes,9,opt,name=series_number,json=seriesNumber,proto3" json:"series_number,omitempty"`
	SeriesInstanceUid                 string                 `protobuf:"bytes,10,opt,name=series_instance_uid,json=seriesInstanceUid,proto3" json:"series_instance_uid,omitempty"`
	SeriesType                        string                 `protobuf:"bytes,11,opt,name=series_type,json=seriesType,proto3" json:"series_type,omitempty"`
	PerformedProcedureStepDescription string                 `protobuf:"bytes,12,opt,name=performed_procedure_step_description,json=performedProcedureStepDescription,proto3" json:"performed_procedure_step_description,omitempty"`
}

func (x *SeriesMainTags) Reset() {
	*x = SeriesMainTags{}
	mi := &file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SeriesMainTags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeriesMainTags) ProtoMessage() {}

func (x *SeriesMainTags) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeriesMainTags.ProtoReflect.Descriptor instead.
func (*SeriesMainTags) Descriptor() ([]byte, []int) {
	return file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_rawDescGZIP(), []int{4}
}

func (x *SeriesMainTags) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *SeriesMainTags) GetModality() string {
	if x != nil {
		return x.Modality
	}
	return ""
}

func (x *SeriesMainTags) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *SeriesMainTags) GetStationName() string {
	if x != nil {
		return x.StationName
	}
	return ""
}

func (x *SeriesMainTags) GetSeriesDescription() string {
	if x != nil {
		return x.SeriesDescription
	}
	return ""
}

func (x *SeriesMainTags) GetBodyPartExamined() string {
	if x != nil {
		return x.BodyPartExamined
	}
	return ""
}

func (x *SeriesMainTags) GetSequenceName() string {
	if x != nil {
		return x.SequenceName
	}
	return ""
}

func (x *SeriesMainTags) GetProtocolName() string {
	if x != nil {
		return x.ProtocolName
	}
	return ""
}

func (x *SeriesMainTags) GetSeriesNumber() string {
	if x != nil {
		return x.SeriesNumber
	}
	return ""
}

func (x *SeriesMainTags) GetSeriesInstanceUid() string {
	if x != nil {
		return x.SeriesInstanceUid
	}
	return ""
}

func (x *SeriesMainTags) GetSeriesType() string {
	if x != nil {
		return x.SeriesType
	}
	return ""
}

func (x *SeriesMainTags) GetPerformedProcedureStepDescription() string {
	if x != nil {
		return x.PerformedProcedureStepDescription
	}
	return ""
}

type InstanceMainTags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time              *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"` // from InstanceCreationDate and InstanceCreationTime
	AcquisitionNumber int64                  `protobuf:"varint,2,opt,name=acquisition_number,json=acquisitionNumber,proto3" json:"acquisition_number,omitempty"`
	ImageComments     string                 `protobuf:"bytes,3,opt,name=image_comments,json=imageComments,proto3" json:"image_comments,omitempty"`
}

func (x *InstanceMainTags) Reset() {
	*x = InstanceMainTags{}
	mi := &file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InstanceMainTags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceMainTags) ProtoMessage() {}

func (x *InstanceMainTags) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceMainTags.ProtoReflect.Descriptor instead.
func (*InstanceMainTags) Descriptor() ([]byte, []int) {
	return file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_rawDescGZIP(), []int{5}
}

func (x *InstanceMainTags) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *InstanceMainTags) GetAcquisitionNumber() int64 {
	if x != nil {
		return x.AcquisitionNumber
	}
	return 0
}

func (x *InstanceMainTags) GetImageComments() string {
	if x != nil {
		return x.ImageComments
	}
	return ""
}

var File_tkd_orthanc_bridge_v1_orthanc_bridge_proto protoreflect.FileDescriptor

var file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x74, 0x6b, 0x64, 0x2f, 0x6f, 0x72, 0x74, 0x68, 0x61, 0x6e, 0x63, 0x5f, 0x62, 0x72,
	0x69, 0x64, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x74, 0x68, 0x61, 0x6e, 0x63, 0x2d,
	0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74, 0x6b,
	0x64, 0x2e, 0x6f, 0x72, 0x74, 0x68, 0x61, 0x6e, 0x63, 0x5f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x02, 0x0a, 0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x63, 0x61, 0x73, 0x65, 0x53, 0x65, 0x6e, 0x73,
	0x69, 0x74, 0x69, 0x76, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x69, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x69, 0x6e, 0x63,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x54, 0x61, 0x67, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x69, 0x62, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x6d, 0x0a, 0x16, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x43, 0x0a, 0x09, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x61, 0x67,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x6f, 0x72,
	0x74, 0x68, 0x61, 0x6e, 0x63, 0x5f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x69, 0x6e, 0x54, 0x61, 0x67, 0x73, 0x52,
	0x08, 0x6d, 0x61, 0x69, 0x6e, 0x54, 0x61, 0x67, 0x73, 0x22, 0xce, 0x01, 0x0a, 0x0f, 0x50, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x69, 0x6e, 0x54, 0x61, 0x67, 0x73, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x2c, 0x0a, 0x12, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x61, 0x74,
	0x69, 0x65, 0x6e, 0x74, 0x42, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x78, 0x12, 0x2a,
	0x0a, 0x11, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x74, 0x68, 0x65, 0x72,
	0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x73, 0x22, 0xbe, 0x03, 0x0a, 0x0d, 0x53,
	0x74, 0x75, 0x64, 0x79, 0x4d, 0x61, 0x69, 0x6e, 0x54, 0x61, 0x67, 0x73, 0x12, 0x2e, 0x0a, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x74, 0x75, 0x64, 0x79, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x74, 0x75, 0x64, 0x79,
	0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x73, 0x74, 0x75, 0x64, 0x79, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x2c, 0x0a, 0x12, 0x73, 0x74, 0x75, 0x64, 0x79, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x74, 0x75,
	0x64, 0x79, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x69, 0x64, 0x12, 0x46, 0x0a,
	0x1f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x65, 0x63, 0x65,
	0x64, 0x75, 0x72, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x50, 0x72, 0x65, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x31, 0x0a, 0x14, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70,
	0x68, 0x79, 0x73, 0x69, 0x63, 0x69, 0x61, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x68, 0x79, 0x73, 0x69, 0x63,
	0x69, 0x61, 0x6e, 0x12, 0x36, 0x0a, 0x17, 0x72, 0x65, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x70, 0x79, 0x73, 0x69, 0x63, 0x69, 0x61, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x72, 0x65, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x50,
	0x79, 0x73, 0x69, 0x63, 0x69, 0x61, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x91, 0x04, 0x0a, 0x0e,
	0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x4d, 0x61, 0x69, 0x6e, 0x54, 0x61, 0x67, 0x73, 0x12, 0x2e,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6d, 0x6f, 0x64, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61,
	0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x21,
	0x0a, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x2d, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2c, 0x0a, 0x12, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x5f, 0x65, 0x78,
	0x61, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x62, 0x6f,
	0x64, 0x79, 0x50, 0x61, 0x72, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2e, 0x0a,
	0x13, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x5f, 0x75, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x69, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4f,
	0x0a, 0x24, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x64, 0x75, 0x72, 0x65, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x21, 0x70, 0x65,
	0x72, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65,
	0x53, 0x74, 0x65, 0x70, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x98, 0x01, 0x0a, 0x10, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4d, 0x61, 0x69, 0x6e,
	0x54, 0x61, 0x67, 0x73, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x11, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2a, 0x91, 0x01, 0x0a, 0x0b, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x45,
	0x41, 0x52, 0x43, 0x48, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x45, 0x41, 0x52,
	0x43, 0x48, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x50, 0x41, 0x54, 0x49, 0x45, 0x4e, 0x54,
	0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x4c, 0x45, 0x56,
	0x45, 0x4c, 0x5f, 0x53, 0x54, 0x55, 0x44, 0x59, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x45,
	0x41, 0x52, 0x43, 0x48, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x53, 0x45, 0x52, 0x49, 0x45,
	0x53, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x4c, 0x45,
	0x56, 0x45, 0x4c, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x04, 0x32, 0x80,
	0x01, 0x0a, 0x0d, 0x4f, 0x72, 0x74, 0x68, 0x61, 0x6e, 0x63, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65,
	0x12, 0x6f, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x2c, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x6f, 0x72, 0x74, 0x68, 0x61, 0x6e, 0x63,
	0x5f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2d, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x6f, 0x72, 0x74, 0x68, 0x61, 0x6e, 0x63, 0x5f, 0x62,
	0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0xf6, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x6f, 0x72,
	0x74, 0x68, 0x61, 0x6e, 0x63, 0x5f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x42,
	0x12, 0x4f, 0x72, 0x74, 0x68, 0x61, 0x6e, 0x63, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x74, 0x69, 0x65, 0x72, 0x6b, 0x6c, 0x69, 0x6e, 0x69, 0x6b, 0x2d, 0x64, 0x6f, 0x62,
	0x65, 0x72, 0x73, 0x62, 0x65, 0x72, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x6b, 0x64, 0x2f, 0x6f, 0x72, 0x74, 0x68, 0x61, 0x6e, 0x63, 0x5f,
	0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6f, 0x72, 0x74, 0x68, 0x61, 0x6e,
	0x63, 0x5f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x4f, 0x58,
	0xaa, 0x02, 0x14, 0x54, 0x6b, 0x64, 0x2e, 0x4f, 0x72, 0x74, 0x68, 0x61, 0x6e, 0x63, 0x42, 0x72,
	0x69, 0x64, 0x67, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x14, 0x54, 0x6b, 0x64, 0x5c, 0x4f, 0x72,
	0x74, 0x68, 0x61, 0x6e, 0x63, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x20, 0x54, 0x6b, 0x64, 0x5c, 0x4f, 0x72, 0x74, 0x68, 0x61, 0x6e, 0x63, 0x42, 0x72, 0x69, 0x64,
	0x67, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x16, 0x54, 0x6b, 0x64, 0x3a, 0x3a, 0x4f, 0x72, 0x74, 0x68, 0x61, 0x6e, 0x63,
	0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_rawDescOnce sync.Once
	file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_rawDescData = file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_rawDesc
)

func file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_rawDescGZIP() []byte {
	file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_rawDescOnce.Do(func() {
		file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_rawDescData = protoimpl.X.CompressGZIP(file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_rawDescData)
	})
	return file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_rawDescData
}

var file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_goTypes = []any{
	(SearchLevel)(0),               // 0: tkd.orthanc_bridge.v1.SearchLevel
	(*SearchPatientsRequest)(nil),  // 1: tkd.orthanc_bridge.v1.SearchPatientsRequest
	(*SearchPatientsResponse)(nil), // 2: tkd.orthanc_bridge.v1.SearchPatientsResponse
	(*PatientMainTags)(nil),        // 3: tkd.orthanc_bridge.v1.PatientMainTags
	(*StudyMainTags)(nil),          // 4: tkd.orthanc_bridge.v1.StudyMainTags
	(*SeriesMainTags)(nil),         // 5: tkd.orthanc_bridge.v1.SeriesMainTags
	(*InstanceMainTags)(nil),       // 6: tkd.orthanc_bridge.v1.InstanceMainTags
	(*timestamppb.Timestamp)(nil),  // 7: google.protobuf.Timestamp
}
var file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_depIdxs = []int32{
	3, // 0: tkd.orthanc_bridge.v1.SearchPatientsResponse.main_tags:type_name -> tkd.orthanc_bridge.v1.PatientMainTags
	7, // 1: tkd.orthanc_bridge.v1.StudyMainTags.time:type_name -> google.protobuf.Timestamp
	7, // 2: tkd.orthanc_bridge.v1.SeriesMainTags.time:type_name -> google.protobuf.Timestamp
	7, // 3: tkd.orthanc_bridge.v1.InstanceMainTags.time:type_name -> google.protobuf.Timestamp
	1, // 4: tkd.orthanc_bridge.v1.OrthancBridge.SearchPatients:input_type -> tkd.orthanc_bridge.v1.SearchPatientsRequest
	2, // 5: tkd.orthanc_bridge.v1.OrthancBridge.SearchPatients:output_type -> tkd.orthanc_bridge.v1.SearchPatientsResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_init() }
func file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_init() {
	if File_tkd_orthanc_bridge_v1_orthanc_bridge_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_goTypes,
		DependencyIndexes: file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_depIdxs,
		EnumInfos:         file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_enumTypes,
		MessageInfos:      file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_msgTypes,
	}.Build()
	File_tkd_orthanc_bridge_v1_orthanc_bridge_proto = out.File
	file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_rawDesc = nil
	file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_goTypes = nil
	file_tkd_orthanc_bridge_v1_orthanc_bridge_proto_depIdxs = nil
}
