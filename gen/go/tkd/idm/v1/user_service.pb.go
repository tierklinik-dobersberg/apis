// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tkd/idm/v1/user_service.proto

package idmv1

import (
	_ "github.com/tierklinik-dobersberg/apis/gen/go/tkd/common/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tkd_idm_v1_user_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_idm_v1_user_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_tkd_idm_v1_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tkd_idm_v1_user_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_idm_v1_user_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_tkd_idm_v1_user_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type ListUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldMask *fieldmaskpb.FieldMask `protobuf:"bytes,1,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
}

func (x *ListUsersRequest) Reset() {
	*x = ListUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tkd_idm_v1_user_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersRequest) ProtoMessage() {}

func (x *ListUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_idm_v1_user_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersRequest.ProtoReflect.Descriptor instead.
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return file_tkd_idm_v1_user_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListUsersRequest) GetFieldMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.FieldMask
	}
	return nil
}

type ListUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*Profile `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *ListUsersResponse) Reset() {
	*x = ListUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tkd_idm_v1_user_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersResponse) ProtoMessage() {}

func (x *ListUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_idm_v1_user_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersResponse.ProtoReflect.Descriptor instead.
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return file_tkd_idm_v1_user_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListUsersResponse) GetUsers() []*Profile {
	if x != nil {
		return x.Users
	}
	return nil
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile          *Profile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	Password         string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	PasswordIsBcrypt bool     `protobuf:"varint,3,opt,name=password_is_bcrypt,json=passwordIsBcrypt,proto3" json:"password_is_bcrypt,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tkd_idm_v1_user_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_idm_v1_user_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_tkd_idm_v1_user_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateUserRequest) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateUserRequest) GetPasswordIsBcrypt() bool {
	if x != nil {
		return x.PasswordIsBcrypt
	}
	return false
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profile           *Profile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	GeneratedPassword string   `protobuf:"bytes,2,opt,name=generated_password,json=generatedPassword,proto3" json:"generated_password,omitempty"`
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tkd_idm_v1_user_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_idm_v1_user_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_tkd_idm_v1_user_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateUserResponse) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *CreateUserResponse) GetGeneratedPassword() string {
	if x != nil {
		return x.GeneratedPassword
	}
	return ""
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tkd_idm_v1_user_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_idm_v1_user_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_tkd_idm_v1_user_service_proto_rawDescGZIP(), []int{6}
}

type UpdateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateUserResponse) Reset() {
	*x = UpdateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tkd_idm_v1_user_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserResponse) ProtoMessage() {}

func (x *UpdateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_idm_v1_user_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserResponse.ProtoReflect.Descriptor instead.
func (*UpdateUserResponse) Descriptor() ([]byte, []int) {
	return file_tkd_idm_v1_user_service_proto_rawDescGZIP(), []int{7}
}

type DeleteUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tkd_idm_v1_user_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_idm_v1_user_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_tkd_idm_v1_user_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteUserResponse) Reset() {
	*x = DeleteUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tkd_idm_v1_user_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserResponse) ProtoMessage() {}

func (x *DeleteUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_idm_v1_user_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserResponse.ProtoReflect.Descriptor instead.
func (*DeleteUserResponse) Descriptor() ([]byte, []int) {
	return file_tkd_idm_v1_user_service_proto_rawDescGZIP(), []int{9}
}

type UserInvite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UserInvite) Reset() {
	*x = UserInvite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tkd_idm_v1_user_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInvite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInvite) ProtoMessage() {}

func (x *UserInvite) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_idm_v1_user_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInvite.ProtoReflect.Descriptor instead.
func (*UserInvite) Descriptor() ([]byte, []int) {
	return file_tkd_idm_v1_user_service_proto_rawDescGZIP(), []int{10}
}

func (x *UserInvite) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInvite) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type InviteUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Invite       []*UserInvite `protobuf:"bytes,1,rep,name=invite,proto3" json:"invite,omitempty"`
	InitialRoles []string      `protobuf:"bytes,2,rep,name=initial_roles,json=initialRoles,proto3" json:"initial_roles,omitempty"`
}

func (x *InviteUserRequest) Reset() {
	*x = InviteUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tkd_idm_v1_user_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteUserRequest) ProtoMessage() {}

func (x *InviteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_idm_v1_user_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteUserRequest.ProtoReflect.Descriptor instead.
func (*InviteUserRequest) Descriptor() ([]byte, []int) {
	return file_tkd_idm_v1_user_service_proto_rawDescGZIP(), []int{11}
}

func (x *InviteUserRequest) GetInvite() []*UserInvite {
	if x != nil {
		return x.Invite
	}
	return nil
}

func (x *InviteUserRequest) GetInitialRoles() []string {
	if x != nil {
		return x.InitialRoles
	}
	return nil
}

type InviteUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InviteUserResponse) Reset() {
	*x = InviteUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tkd_idm_v1_user_service_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteUserResponse) ProtoMessage() {}

func (x *InviteUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_idm_v1_user_service_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteUserResponse.ProtoReflect.Descriptor instead.
func (*InviteUserResponse) Descriptor() ([]byte, []int) {
	return file_tkd_idm_v1_user_service_proto_rawDescGZIP(), []int{12}
}

var File_tkd_idm_v1_user_service_proto protoreflect.FileDescriptor

var file_tkd_idm_v1_user_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x74, 0x6b, 0x64, 0x2f, 0x69, 0x64, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x74, 0x6b, 0x64, 0x2e, 0x69, 0x64, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x15, 0x74, 0x6b, 0x64, 0x2f, 0x69, 0x64, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x74, 0x6b, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x6b, 0x64, 0x2e,
	0x69, 0x64, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x22, 0x4d, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d,
	0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x22, 0x3e, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x69, 0x64, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x22, 0x8c, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x69, 0x64,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x73,
	0x5f, 0x62, 0x63, 0x72, 0x79, 0x70, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x49, 0x73, 0x42, 0x63, 0x72, 0x79, 0x70, 0x74, 0x22,
	0x72, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x69, 0x64, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23,
	0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0x0a, 0x0a, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x68, 0x0a, 0x11, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x69, 0x64, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x06,
	0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x6c, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x14, 0x0a, 0x12, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0xac, 0x04, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x74,
	0x6b, 0x64, 0x2e, 0x69, 0x64, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x69,
	0x64, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x08, 0xb2, 0x7e, 0x02, 0x08, 0x01, 0x90, 0x02, 0x01, 0x12,
	0x52, 0x0a, 0x0a, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x2e,
	0x74, 0x6b, 0x64, 0x2e, 0x69, 0x64, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74,
	0x6b, 0x64, 0x2e, 0x69, 0x64, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x05, 0xb2, 0x7e,
	0x02, 0x08, 0x01, 0x12, 0x52, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x12, 0x1c, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x69, 0x64, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x69, 0x64, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x08, 0xb2,
	0x7e, 0x02, 0x08, 0x01, 0x90, 0x02, 0x01, 0x12, 0x61, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x69, 0x64, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x69, 0x64, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0xb2, 0x7e, 0x11, 0x08, 0x01, 0x12, 0x0d, 0x69, 0x64, 0x6d,
	0x5f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x75, 0x73, 0x65, 0x72, 0x12, 0x61, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x69,
	0x64, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x69, 0x64,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0xb2, 0x7e, 0x11, 0x08, 0x01, 0x12, 0x0d,
	0x69, 0x64, 0x6d, 0x5f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x75, 0x73, 0x65, 0x72, 0x12, 0x61, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x74, 0x6b,
	0x64, 0x2e, 0x69, 0x64, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x6b, 0x64,
	0x2e, 0x69, 0x64, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0xb2, 0x7e, 0x11, 0x08,
	0x01, 0x12, 0x0d, 0x69, 0x64, 0x6d, 0x5f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x75, 0x73, 0x65, 0x72,
	0x42, 0xab, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x69, 0x64, 0x6d,
	0x2e, 0x76, 0x31, 0x42, 0x10, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x65, 0x72, 0x6b, 0x6c, 0x69, 0x6e, 0x69, 0x6b, 0x2d, 0x64,
	0x6f, 0x62, 0x65, 0x72, 0x73, 0x62, 0x65, 0x72, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x6b, 0x64, 0x2f, 0x69, 0x64, 0x6d, 0x2f, 0x76, 0x31,
	0x3b, 0x69, 0x64, 0x6d, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x49, 0x58, 0xaa, 0x02, 0x0a, 0x54,
	0x6b, 0x64, 0x2e, 0x49, 0x64, 0x6d, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0a, 0x54, 0x6b, 0x64, 0x5c,
	0x49, 0x64, 0x6d, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x16, 0x54, 0x6b, 0x64, 0x5c, 0x49, 0x64, 0x6d,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0c, 0x54, 0x6b, 0x64, 0x3a, 0x3a, 0x49, 0x64, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tkd_idm_v1_user_service_proto_rawDescOnce sync.Once
	file_tkd_idm_v1_user_service_proto_rawDescData = file_tkd_idm_v1_user_service_proto_rawDesc
)

func file_tkd_idm_v1_user_service_proto_rawDescGZIP() []byte {
	file_tkd_idm_v1_user_service_proto_rawDescOnce.Do(func() {
		file_tkd_idm_v1_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_tkd_idm_v1_user_service_proto_rawDescData)
	})
	return file_tkd_idm_v1_user_service_proto_rawDescData
}

var file_tkd_idm_v1_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_tkd_idm_v1_user_service_proto_goTypes = []interface{}{
	(*GetUserRequest)(nil),        // 0: tkd.idm.v1.GetUserRequest
	(*GetUserResponse)(nil),       // 1: tkd.idm.v1.GetUserResponse
	(*ListUsersRequest)(nil),      // 2: tkd.idm.v1.ListUsersRequest
	(*ListUsersResponse)(nil),     // 3: tkd.idm.v1.ListUsersResponse
	(*CreateUserRequest)(nil),     // 4: tkd.idm.v1.CreateUserRequest
	(*CreateUserResponse)(nil),    // 5: tkd.idm.v1.CreateUserResponse
	(*UpdateUserRequest)(nil),     // 6: tkd.idm.v1.UpdateUserRequest
	(*UpdateUserResponse)(nil),    // 7: tkd.idm.v1.UpdateUserResponse
	(*DeleteUserRequest)(nil),     // 8: tkd.idm.v1.DeleteUserRequest
	(*DeleteUserResponse)(nil),    // 9: tkd.idm.v1.DeleteUserResponse
	(*UserInvite)(nil),            // 10: tkd.idm.v1.UserInvite
	(*InviteUserRequest)(nil),     // 11: tkd.idm.v1.InviteUserRequest
	(*InviteUserResponse)(nil),    // 12: tkd.idm.v1.InviteUserResponse
	(*User)(nil),                  // 13: tkd.idm.v1.User
	(*fieldmaskpb.FieldMask)(nil), // 14: google.protobuf.FieldMask
	(*Profile)(nil),               // 15: tkd.idm.v1.Profile
}
var file_tkd_idm_v1_user_service_proto_depIdxs = []int32{
	13, // 0: tkd.idm.v1.GetUserResponse.user:type_name -> tkd.idm.v1.User
	14, // 1: tkd.idm.v1.ListUsersRequest.field_mask:type_name -> google.protobuf.FieldMask
	15, // 2: tkd.idm.v1.ListUsersResponse.users:type_name -> tkd.idm.v1.Profile
	15, // 3: tkd.idm.v1.CreateUserRequest.profile:type_name -> tkd.idm.v1.Profile
	15, // 4: tkd.idm.v1.CreateUserResponse.profile:type_name -> tkd.idm.v1.Profile
	10, // 5: tkd.idm.v1.InviteUserRequest.invite:type_name -> tkd.idm.v1.UserInvite
	0,  // 6: tkd.idm.v1.UserService.GetUser:input_type -> tkd.idm.v1.GetUserRequest
	11, // 7: tkd.idm.v1.UserService.InviteUser:input_type -> tkd.idm.v1.InviteUserRequest
	2,  // 8: tkd.idm.v1.UserService.ListUsers:input_type -> tkd.idm.v1.ListUsersRequest
	4,  // 9: tkd.idm.v1.UserService.CreateUser:input_type -> tkd.idm.v1.CreateUserRequest
	6,  // 10: tkd.idm.v1.UserService.UpdateUser:input_type -> tkd.idm.v1.UpdateUserRequest
	8,  // 11: tkd.idm.v1.UserService.DeleteUser:input_type -> tkd.idm.v1.DeleteUserRequest
	1,  // 12: tkd.idm.v1.UserService.GetUser:output_type -> tkd.idm.v1.GetUserResponse
	12, // 13: tkd.idm.v1.UserService.InviteUser:output_type -> tkd.idm.v1.InviteUserResponse
	3,  // 14: tkd.idm.v1.UserService.ListUsers:output_type -> tkd.idm.v1.ListUsersResponse
	5,  // 15: tkd.idm.v1.UserService.CreateUser:output_type -> tkd.idm.v1.CreateUserResponse
	7,  // 16: tkd.idm.v1.UserService.UpdateUser:output_type -> tkd.idm.v1.UpdateUserResponse
	9,  // 17: tkd.idm.v1.UserService.DeleteUser:output_type -> tkd.idm.v1.DeleteUserResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_tkd_idm_v1_user_service_proto_init() }
func file_tkd_idm_v1_user_service_proto_init() {
	if File_tkd_idm_v1_user_service_proto != nil {
		return
	}
	file_tkd_idm_v1_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tkd_idm_v1_user_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tkd_idm_v1_user_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tkd_idm_v1_user_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUsersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tkd_idm_v1_user_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUsersResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tkd_idm_v1_user_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tkd_idm_v1_user_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tkd_idm_v1_user_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tkd_idm_v1_user_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tkd_idm_v1_user_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tkd_idm_v1_user_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tkd_idm_v1_user_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInvite); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tkd_idm_v1_user_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tkd_idm_v1_user_service_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteUserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tkd_idm_v1_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tkd_idm_v1_user_service_proto_goTypes,
		DependencyIndexes: file_tkd_idm_v1_user_service_proto_depIdxs,
		MessageInfos:      file_tkd_idm_v1_user_service_proto_msgTypes,
	}.Build()
	File_tkd_idm_v1_user_service_proto = out.File
	file_tkd_idm_v1_user_service_proto_rawDesc = nil
	file_tkd_idm_v1_user_service_proto_goTypes = nil
	file_tkd_idm_v1_user_service_proto_depIdxs = nil
}
