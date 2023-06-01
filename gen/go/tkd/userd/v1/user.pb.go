// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: tkd/userd/v1/user.proto

package userdv1

import (
	_ "github.com/bufbuild/buf-tour/gen/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EMail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address  string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Verified bool   `protobuf:"varint,2,opt,name=verified,proto3" json:"verified,omitempty"`
}

func (x *EMail) Reset() {
	*x = EMail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tkd_userd_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EMail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EMail) ProtoMessage() {}

func (x *EMail) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_userd_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EMail.ProtoReflect.Descriptor instead.
func (*EMail) Descriptor() ([]byte, []int) {
	return file_tkd_userd_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *EMail) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *EMail) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CityCode string `protobuf:"bytes,1,opt,name=city_code,json=cityCode,proto3" json:"city_code,omitempty"`
	CityName string `protobuf:"bytes,2,opt,name=city_name,json=cityName,proto3" json:"city_name,omitempty"`
	Street   string `protobuf:"bytes,3,opt,name=street,proto3" json:"street,omitempty"`
	Extra    string `protobuf:"bytes,4,opt,name=extra,proto3" json:"extra,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tkd_userd_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_userd_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_tkd_userd_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *Address) GetCityCode() string {
	if x != nil {
		return x.CityCode
	}
	return ""
}

func (x *Address) GetCityName() string {
	if x != nil {
		return x.CityName
	}
	return ""
}

func (x *Address) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Address) GetExtra() string {
	if x != nil {
		return x.Extra
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username       string                     `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	DisplayName    string                     `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	FirstName      string                     `protobuf:"bytes,4,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName       string                     `protobuf:"bytes,5,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Addresses      []*Address                 `protobuf:"bytes,6,rep,name=addresses,proto3" json:"addresses,omitempty"`
	PhoneNumbers   []string                   `protobuf:"bytes,7,rep,name=phone_numbers,json=phoneNumbers,proto3" json:"phone_numbers,omitempty"`
	EmailAddresses []*EMail                   `protobuf:"bytes,8,rep,name=email_addresses,json=emailAddresses,proto3" json:"email_addresses,omitempty"`
	Extra          map[string]*structpb.Value `protobuf:"bytes,9,rep,name=extra,proto3" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Avatar         string                     `protobuf:"bytes,10,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tkd_userd_v1_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_tkd_userd_v1_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_tkd_userd_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetAddresses() []*Address {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *User) GetPhoneNumbers() []string {
	if x != nil {
		return x.PhoneNumbers
	}
	return nil
}

func (x *User) GetEmailAddresses() []*EMail {
	if x != nil {
		return x.EmailAddresses
	}
	return nil
}

func (x *User) GetExtra() map[string]*structpb.Value {
	if x != nil {
		return x.Extra
	}
	return nil
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

var File_tkd_userd_v1_user_proto protoreflect.FileDescriptor

var file_tkd_userd_v1_user_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x6b, 0x64, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x74, 0x6b, 0x64, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x74, 0x6b, 0x64, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a,
	0x05, 0x45, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x60, 0x01,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x22, 0x94, 0x01, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x2c, 0x0a, 0x09, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xfa, 0x42, 0x0c, 0x72, 0x0a, 0x28, 0x10, 0x32, 0x06, 0x5b,
	0x30, 0x2d, 0x39, 0x5d, 0x2b, 0x52, 0x08, 0x63, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x24, 0x0a, 0x09, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x63, 0x69, 0x74,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06,
	0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x22, 0xfb, 0x03, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x03,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x0c, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0a, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x03, 0xa8, 0x7e, 0x01, 0x52, 0x0b, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xa8,
	0x7e, 0x01, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xa8, 0x7e, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x38, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x03, 0xa8, 0x7e, 0x01, 0x52, 0x09,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x0d, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09,
	0x42, 0x03, 0xa8, 0x7e, 0x01, 0x52, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x12, 0x41, 0x0a, 0x0f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74,
	0x6b, 0x64, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x4d, 0x61, 0x69,
	0x6c, 0x42, 0x03, 0xa8, 0x7e, 0x01, 0x52, 0x0e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x12, 0x1b, 0x0a, 0x06, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xa8, 0x7e, 0x01,
	0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x1a, 0x50, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0xa6, 0x01, 0x0a, 0x10, 0x63,
	0x6f, 0x6d, 0x2e, 0x74, 0x6b, 0x64, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x42,
	0x09, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x2f, 0x62, 0x75, 0x66, 0x2d, 0x74, 0x6f, 0x75, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x74,
	0x6b, 0x64, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72,
	0x64, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x55, 0x58, 0xaa, 0x02, 0x0c, 0x54, 0x6b, 0x64, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x64, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x54, 0x6b, 0x64, 0x5c, 0x55,
	0x73, 0x65, 0x72, 0x64, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x54, 0x6b, 0x64, 0x5c, 0x55, 0x73,
	0x65, 0x72, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0e, 0x54, 0x6b, 0x64, 0x3a, 0x3a, 0x55, 0x73, 0x65, 0x72, 0x64, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tkd_userd_v1_user_proto_rawDescOnce sync.Once
	file_tkd_userd_v1_user_proto_rawDescData = file_tkd_userd_v1_user_proto_rawDesc
)

func file_tkd_userd_v1_user_proto_rawDescGZIP() []byte {
	file_tkd_userd_v1_user_proto_rawDescOnce.Do(func() {
		file_tkd_userd_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_tkd_userd_v1_user_proto_rawDescData)
	})
	return file_tkd_userd_v1_user_proto_rawDescData
}

var file_tkd_userd_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tkd_userd_v1_user_proto_goTypes = []interface{}{
	(*EMail)(nil),          // 0: tkd.userd.v1.EMail
	(*Address)(nil),        // 1: tkd.userd.v1.Address
	(*User)(nil),           // 2: tkd.userd.v1.User
	nil,                    // 3: tkd.userd.v1.User.ExtraEntry
	(*structpb.Value)(nil), // 4: google.protobuf.Value
}
var file_tkd_userd_v1_user_proto_depIdxs = []int32{
	1, // 0: tkd.userd.v1.User.addresses:type_name -> tkd.userd.v1.Address
	0, // 1: tkd.userd.v1.User.email_addresses:type_name -> tkd.userd.v1.EMail
	3, // 2: tkd.userd.v1.User.extra:type_name -> tkd.userd.v1.User.ExtraEntry
	4, // 3: tkd.userd.v1.User.ExtraEntry.value:type_name -> google.protobuf.Value
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tkd_userd_v1_user_proto_init() }
func file_tkd_userd_v1_user_proto_init() {
	if File_tkd_userd_v1_user_proto != nil {
		return
	}
	file_tkd_userd_v1_descriptor_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tkd_userd_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EMail); i {
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
		file_tkd_userd_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_tkd_userd_v1_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_tkd_userd_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tkd_userd_v1_user_proto_goTypes,
		DependencyIndexes: file_tkd_userd_v1_user_proto_depIdxs,
		MessageInfos:      file_tkd_userd_v1_user_proto_msgTypes,
	}.Build()
	File_tkd_userd_v1_user_proto = out.File
	file_tkd_userd_v1_user_proto_rawDesc = nil
	file_tkd_userd_v1_user_proto_goTypes = nil
	file_tkd_userd_v1_user_proto_depIdxs = nil
}
