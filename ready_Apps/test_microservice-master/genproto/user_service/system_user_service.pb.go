// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: system_user_service.proto

package user_service

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SystemUserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SystemUserId) Reset() {
	*x = SystemUserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_user_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemUserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemUserId) ProtoMessage() {}

func (x *SystemUserId) ProtoReflect() protoreflect.Message {
	mi := &file_system_user_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemUserId.ProtoReflect.Descriptor instead.
func (*SystemUserId) Descriptor() ([]byte, []int) {
	return file_system_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *SystemUserId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetSystemUserByUsernameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GetSystemUserByUsernameRequest) Reset() {
	*x = GetSystemUserByUsernameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_user_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSystemUserByUsernameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSystemUserByUsernameRequest) ProtoMessage() {}

func (x *GetSystemUserByUsernameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_user_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSystemUserByUsernameRequest.ProtoReflect.Descriptor instead.
func (*GetSystemUserByUsernameRequest) Descriptor() ([]byte, []int) {
	return file_system_user_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetSystemUserByUsernameRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type GetSystemUserByCredentialsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *GetSystemUserByCredentialsRequest) Reset() {
	*x = GetSystemUserByCredentialsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_user_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSystemUserByCredentialsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSystemUserByCredentialsRequest) ProtoMessage() {}

func (x *GetSystemUserByCredentialsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_user_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSystemUserByCredentialsRequest.ProtoReflect.Descriptor instead.
func (*GetSystemUserByCredentialsRequest) Descriptor() ([]byte, []int) {
	return file_system_user_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetSystemUserByCredentialsRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetSystemUserByCredentialsRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ChangeSystemUserPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *ChangeSystemUserPasswordRequest) Reset() {
	*x = ChangeSystemUserPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_user_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeSystemUserPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeSystemUserPasswordRequest) ProtoMessage() {}

func (x *ChangeSystemUserPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_user_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeSystemUserPasswordRequest.ProtoReflect.Descriptor instead.
func (*ChangeSystemUserPasswordRequest) Descriptor() ([]byte, []int) {
	return file_system_user_service_proto_rawDescGZIP(), []int{3}
}

func (x *ChangeSystemUserPasswordRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChangeSystemUserPasswordRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GetAllSystemUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   uint64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit  uint64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetAllSystemUsersRequest) Reset() {
	*x = GetAllSystemUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_user_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllSystemUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllSystemUsersRequest) ProtoMessage() {}

func (x *GetAllSystemUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_user_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllSystemUsersRequest.ProtoReflect.Descriptor instead.
func (*GetAllSystemUsersRequest) Descriptor() ([]byte, []int) {
	return file_system_user_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllSystemUsersRequest) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllSystemUsersRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllSystemUsersRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetAllSystemUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SystemUsers []*SystemUser `protobuf:"bytes,1,rep,name=system_users,json=systemUsers,proto3" json:"system_users,omitempty"`
	Count       uint64        `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllSystemUsersResponse) Reset() {
	*x = GetAllSystemUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_user_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllSystemUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllSystemUsersResponse) ProtoMessage() {}

func (x *GetAllSystemUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_user_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllSystemUsersResponse.ProtoReflect.Descriptor instead.
func (*GetAllSystemUsersResponse) Descriptor() ([]byte, []int) {
	return file_system_user_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllSystemUsersResponse) GetSystemUsers() []*SystemUser {
	if x != nil {
		return x.SystemUsers
	}
	return nil
}

func (x *GetAllSystemUsersResponse) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type DeleteSystemUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ShipperId string `protobuf:"bytes,2,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
}

func (x *DeleteSystemUserRequest) Reset() {
	*x = DeleteSystemUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_user_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSystemUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSystemUserRequest) ProtoMessage() {}

func (x *DeleteSystemUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_user_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSystemUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteSystemUserRequest) Descriptor() ([]byte, []int) {
	return file_system_user_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteSystemUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteSystemUserRequest) GetShipperId() string {
	if x != nil {
		return x.ShipperId
	}
	return ""
}

var File_system_user_service_proto protoreflect.FileDescriptor

var file_system_user_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x0c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x5b, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x4d, 0x0a, 0x1f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x5c, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x6a, 0x0a,
	0x19, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x52, 0x0b, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x48, 0x0a, 0x17, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65,
	0x72, 0x49, 0x64, 0x32, 0xdd, 0x04, 0x0a, 0x11, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x1a, 0x14, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x06, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x12, 0x22, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x38, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x51, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x28, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65,
	0x72, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x79, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x2b, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x79, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0e,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x29,
	0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_system_user_service_proto_rawDescOnce sync.Once
	file_system_user_service_proto_rawDescData = file_system_user_service_proto_rawDesc
)

func file_system_user_service_proto_rawDescGZIP() []byte {
	file_system_user_service_proto_rawDescOnce.Do(func() {
		file_system_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_system_user_service_proto_rawDescData)
	})
	return file_system_user_service_proto_rawDescData
}

var file_system_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_system_user_service_proto_goTypes = []interface{}{
	(*SystemUserId)(nil),                      // 0: genproto.SystemUserId
	(*GetSystemUserByUsernameRequest)(nil),    // 1: genproto.GetSystemUserByUsernameRequest
	(*GetSystemUserByCredentialsRequest)(nil), // 2: genproto.GetSystemUserByCredentialsRequest
	(*ChangeSystemUserPasswordRequest)(nil),   // 3: genproto.ChangeSystemUserPasswordRequest
	(*GetAllSystemUsersRequest)(nil),          // 4: genproto.GetAllSystemUsersRequest
	(*GetAllSystemUsersResponse)(nil),         // 5: genproto.GetAllSystemUsersResponse
	(*DeleteSystemUserRequest)(nil),           // 6: genproto.DeleteSystemUserRequest
	(*SystemUser)(nil),                        // 7: genproto.SystemUser
	(*empty.Empty)(nil),                       // 8: google.protobuf.Empty
}
var file_system_user_service_proto_depIdxs = []int32{
	7, // 0: genproto.GetAllSystemUsersResponse.system_users:type_name -> genproto.SystemUser
	7, // 1: genproto.SystemUserService.Create:input_type -> genproto.SystemUser
	0, // 2: genproto.SystemUserService.Get:input_type -> genproto.SystemUserId
	4, // 3: genproto.SystemUserService.GetAll:input_type -> genproto.GetAllSystemUsersRequest
	7, // 4: genproto.SystemUserService.Update:input_type -> genproto.SystemUser
	6, // 5: genproto.SystemUserService.Delete:input_type -> genproto.DeleteSystemUserRequest
	1, // 6: genproto.SystemUserService.GetByUsername:input_type -> genproto.GetSystemUserByUsernameRequest
	2, // 7: genproto.SystemUserService.GetByCredentials:input_type -> genproto.GetSystemUserByCredentialsRequest
	3, // 8: genproto.SystemUserService.ChangePassword:input_type -> genproto.ChangeSystemUserPasswordRequest
	0, // 9: genproto.SystemUserService.Create:output_type -> genproto.SystemUserId
	7, // 10: genproto.SystemUserService.Get:output_type -> genproto.SystemUser
	5, // 11: genproto.SystemUserService.GetAll:output_type -> genproto.GetAllSystemUsersResponse
	8, // 12: genproto.SystemUserService.Update:output_type -> google.protobuf.Empty
	8, // 13: genproto.SystemUserService.Delete:output_type -> google.protobuf.Empty
	7, // 14: genproto.SystemUserService.GetByUsername:output_type -> genproto.SystemUser
	7, // 15: genproto.SystemUserService.GetByCredentials:output_type -> genproto.SystemUser
	8, // 16: genproto.SystemUserService.ChangePassword:output_type -> google.protobuf.Empty
	9, // [9:17] is the sub-list for method output_type
	1, // [1:9] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_system_user_service_proto_init() }
func file_system_user_service_proto_init() {
	if File_system_user_service_proto != nil {
		return
	}
	file_system_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_system_user_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemUserId); i {
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
		file_system_user_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSystemUserByUsernameRequest); i {
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
		file_system_user_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSystemUserByCredentialsRequest); i {
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
		file_system_user_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeSystemUserPasswordRequest); i {
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
		file_system_user_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllSystemUsersRequest); i {
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
		file_system_user_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllSystemUsersResponse); i {
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
		file_system_user_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSystemUserRequest); i {
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
			RawDescriptor: file_system_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_system_user_service_proto_goTypes,
		DependencyIndexes: file_system_user_service_proto_depIdxs,
		MessageInfos:      file_system_user_service_proto_msgTypes,
	}.Build()
	File_system_user_service_proto = out.File
	file_system_user_service_proto_rawDesc = nil
	file_system_user_service_proto_goTypes = nil
	file_system_user_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SystemUserServiceClient is the client API for SystemUserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SystemUserServiceClient interface {
	Create(ctx context.Context, in *SystemUser, opts ...grpc.CallOption) (*SystemUserId, error)
	Get(ctx context.Context, in *SystemUserId, opts ...grpc.CallOption) (*SystemUser, error)
	GetAll(ctx context.Context, in *GetAllSystemUsersRequest, opts ...grpc.CallOption) (*GetAllSystemUsersResponse, error)
	Update(ctx context.Context, in *SystemUser, opts ...grpc.CallOption) (*empty.Empty, error)
	Delete(ctx context.Context, in *DeleteSystemUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetByUsername(ctx context.Context, in *GetSystemUserByUsernameRequest, opts ...grpc.CallOption) (*SystemUser, error)
	GetByCredentials(ctx context.Context, in *GetSystemUserByCredentialsRequest, opts ...grpc.CallOption) (*SystemUser, error)
	ChangePassword(ctx context.Context, in *ChangeSystemUserPasswordRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type systemUserServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemUserServiceClient(cc grpc.ClientConnInterface) SystemUserServiceClient {
	return &systemUserServiceClient{cc}
}

func (c *systemUserServiceClient) Create(ctx context.Context, in *SystemUser, opts ...grpc.CallOption) (*SystemUserId, error) {
	out := new(SystemUserId)
	err := c.cc.Invoke(ctx, "/genproto.SystemUserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserServiceClient) Get(ctx context.Context, in *SystemUserId, opts ...grpc.CallOption) (*SystemUser, error) {
	out := new(SystemUser)
	err := c.cc.Invoke(ctx, "/genproto.SystemUserService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserServiceClient) GetAll(ctx context.Context, in *GetAllSystemUsersRequest, opts ...grpc.CallOption) (*GetAllSystemUsersResponse, error) {
	out := new(GetAllSystemUsersResponse)
	err := c.cc.Invoke(ctx, "/genproto.SystemUserService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserServiceClient) Update(ctx context.Context, in *SystemUser, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/genproto.SystemUserService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserServiceClient) Delete(ctx context.Context, in *DeleteSystemUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/genproto.SystemUserService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserServiceClient) GetByUsername(ctx context.Context, in *GetSystemUserByUsernameRequest, opts ...grpc.CallOption) (*SystemUser, error) {
	out := new(SystemUser)
	err := c.cc.Invoke(ctx, "/genproto.SystemUserService/GetByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserServiceClient) GetByCredentials(ctx context.Context, in *GetSystemUserByCredentialsRequest, opts ...grpc.CallOption) (*SystemUser, error) {
	out := new(SystemUser)
	err := c.cc.Invoke(ctx, "/genproto.SystemUserService/GetByCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemUserServiceClient) ChangePassword(ctx context.Context, in *ChangeSystemUserPasswordRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/genproto.SystemUserService/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemUserServiceServer is the server API for SystemUserService service.
type SystemUserServiceServer interface {
	Create(context.Context, *SystemUser) (*SystemUserId, error)
	Get(context.Context, *SystemUserId) (*SystemUser, error)
	GetAll(context.Context, *GetAllSystemUsersRequest) (*GetAllSystemUsersResponse, error)
	Update(context.Context, *SystemUser) (*empty.Empty, error)
	Delete(context.Context, *DeleteSystemUserRequest) (*empty.Empty, error)
	GetByUsername(context.Context, *GetSystemUserByUsernameRequest) (*SystemUser, error)
	GetByCredentials(context.Context, *GetSystemUserByCredentialsRequest) (*SystemUser, error)
	ChangePassword(context.Context, *ChangeSystemUserPasswordRequest) (*empty.Empty, error)
}

// UnimplementedSystemUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSystemUserServiceServer struct {
}

func (*UnimplementedSystemUserServiceServer) Create(context.Context, *SystemUser) (*SystemUserId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedSystemUserServiceServer) Get(context.Context, *SystemUserId) (*SystemUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedSystemUserServiceServer) GetAll(context.Context, *GetAllSystemUsersRequest) (*GetAllSystemUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedSystemUserServiceServer) Update(context.Context, *SystemUser) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedSystemUserServiceServer) Delete(context.Context, *DeleteSystemUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedSystemUserServiceServer) GetByUsername(context.Context, *GetSystemUserByUsernameRequest) (*SystemUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByUsername not implemented")
}
func (*UnimplementedSystemUserServiceServer) GetByCredentials(context.Context, *GetSystemUserByCredentialsRequest) (*SystemUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByCredentials not implemented")
}
func (*UnimplementedSystemUserServiceServer) ChangePassword(context.Context, *ChangeSystemUserPasswordRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}

func RegisterSystemUserServiceServer(s *grpc.Server, srv SystemUserServiceServer) {
	s.RegisterService(&_SystemUserService_serviceDesc, srv)
}

func _SystemUserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.SystemUserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).Create(ctx, req.(*SystemUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.SystemUserService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).Get(ctx, req.(*SystemUserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllSystemUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.SystemUserService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).GetAll(ctx, req.(*GetAllSystemUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.SystemUserService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).Update(ctx, req.(*SystemUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSystemUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.SystemUserService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).Delete(ctx, req.(*DeleteSystemUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserService_GetByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSystemUserByUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).GetByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.SystemUserService/GetByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).GetByUsername(ctx, req.(*GetSystemUserByUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserService_GetByCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSystemUserByCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).GetByCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.SystemUserService/GetByCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).GetByCredentials(ctx, req.(*GetSystemUserByCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemUserService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeSystemUserPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemUserServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.SystemUserService/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemUserServiceServer).ChangePassword(ctx, req.(*ChangeSystemUserPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SystemUserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.SystemUserService",
	HandlerType: (*SystemUserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SystemUserService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _SystemUserService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _SystemUserService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SystemUserService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SystemUserService_Delete_Handler,
		},
		{
			MethodName: "GetByUsername",
			Handler:    _SystemUserService_GetByUsername_Handler,
		},
		{
			MethodName: "GetByCredentials",
			Handler:    _SystemUserService_GetByCredentials_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _SystemUserService_ChangePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "system_user_service.proto",
}
