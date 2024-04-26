// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: promo.proto

package content_service

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

type Promo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       *Language `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description *Language `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	IsActive    bool      `protobuf:"varint,4,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	Image       string    `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	ShipperId   string    `protobuf:"bytes,6,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
	StartTime   string    `protobuf:"bytes,7,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime     string    `protobuf:"bytes,8,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	CreatedAt   string    `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string    `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt   string    `protobuf:"bytes,11,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Promo) Reset() {
	*x = Promo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Promo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Promo) ProtoMessage() {}

func (x *Promo) ProtoReflect() protoreflect.Message {
	mi := &file_promo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Promo.ProtoReflect.Descriptor instead.
func (*Promo) Descriptor() ([]byte, []int) {
	return file_promo_proto_rawDescGZIP(), []int{0}
}

func (x *Promo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Promo) GetTitle() *Language {
	if x != nil {
		return x.Title
	}
	return nil
}

func (x *Promo) GetDescription() *Language {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *Promo) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Promo) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Promo) GetShipperId() string {
	if x != nil {
		return x.ShipperId
	}
	return ""
}

func (x *Promo) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *Promo) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *Promo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Promo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Promo) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type GetAllPromosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page      int64  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	IsActive  bool   `protobuf:"varint,2,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	Limit     int64  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	ShipperId string `protobuf:"bytes,4,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
	Search    string `protobuf:"bytes,5,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetAllPromosRequest) Reset() {
	*x = GetAllPromosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPromosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPromosRequest) ProtoMessage() {}

func (x *GetAllPromosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_promo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPromosRequest.ProtoReflect.Descriptor instead.
func (*GetAllPromosRequest) Descriptor() ([]byte, []int) {
	return file_promo_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllPromosRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllPromosRequest) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *GetAllPromosRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllPromosRequest) GetShipperId() string {
	if x != nil {
		return x.ShipperId
	}
	return ""
}

func (x *GetAllPromosRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type PromoId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PromoId) Reset() {
	*x = PromoId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromoId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromoId) ProtoMessage() {}

func (x *PromoId) ProtoReflect() protoreflect.Message {
	mi := &file_promo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromoId.ProtoReflect.Descriptor instead.
func (*PromoId) Descriptor() ([]byte, []int) {
	return file_promo_proto_rawDescGZIP(), []int{2}
}

func (x *PromoId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAllPromosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Promos []*Promo `protobuf:"bytes,1,rep,name=promos,proto3" json:"promos,omitempty"`
	Count  int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllPromosResponse) Reset() {
	*x = GetAllPromosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_promo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPromosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPromosResponse) ProtoMessage() {}

func (x *GetAllPromosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_promo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPromosResponse.ProtoReflect.Descriptor instead.
func (*GetAllPromosResponse) Descriptor() ([]byte, []int) {
	return file_promo_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllPromosResponse) GetPromos() []*Promo {
	if x != nil {
		return x.Promos
	}
	return nil
}

func (x *GetAllPromosResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_promo_proto protoreflect.FileDescriptor

var file_promo_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x02, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x28, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x93, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x70,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x69,
	0x70, 0x70, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x19,
	0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x55, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x27, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f,
	0x6d, 0x6f, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x32, 0x94, 0x02, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x1a, 0x11, 0x2e, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x49, 0x64, 0x22,
	0x00, 0x12, 0x2c, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x1a, 0x0f, 0x2e, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x49, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x1d, 0x2e, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x6d, 0x6f,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x11, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f,
	0x6d, 0x6f, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x11, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x72, 0x6f,
	0x6d, 0x6f, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_promo_proto_rawDescOnce sync.Once
	file_promo_proto_rawDescData = file_promo_proto_rawDesc
)

func file_promo_proto_rawDescGZIP() []byte {
	file_promo_proto_rawDescOnce.Do(func() {
		file_promo_proto_rawDescData = protoimpl.X.CompressGZIP(file_promo_proto_rawDescData)
	})
	return file_promo_proto_rawDescData
}

var file_promo_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_promo_proto_goTypes = []interface{}{
	(*Promo)(nil),                // 0: genproto.Promo
	(*GetAllPromosRequest)(nil),  // 1: genproto.GetAllPromosRequest
	(*PromoId)(nil),              // 2: genproto.PromoId
	(*GetAllPromosResponse)(nil), // 3: genproto.GetAllPromosResponse
	(*Language)(nil),             // 4: genproto.Language
	(*Empty)(nil),                // 5: genproto.Empty
}
var file_promo_proto_depIdxs = []int32{
	4, // 0: genproto.Promo.title:type_name -> genproto.Language
	4, // 1: genproto.Promo.description:type_name -> genproto.Language
	0, // 2: genproto.GetAllPromosResponse.promos:type_name -> genproto.Promo
	0, // 3: genproto.PromoService.Create:input_type -> genproto.Promo
	0, // 4: genproto.PromoService.Update:input_type -> genproto.Promo
	1, // 5: genproto.PromoService.GetAll:input_type -> genproto.GetAllPromosRequest
	2, // 6: genproto.PromoService.Get:input_type -> genproto.PromoId
	2, // 7: genproto.PromoService.Delete:input_type -> genproto.PromoId
	2, // 8: genproto.PromoService.Create:output_type -> genproto.PromoId
	5, // 9: genproto.PromoService.Update:output_type -> genproto.Empty
	3, // 10: genproto.PromoService.GetAll:output_type -> genproto.GetAllPromosResponse
	0, // 11: genproto.PromoService.Get:output_type -> genproto.Promo
	5, // 12: genproto.PromoService.Delete:output_type -> genproto.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_promo_proto_init() }
func file_promo_proto_init() {
	if File_promo_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_promo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Promo); i {
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
		file_promo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPromosRequest); i {
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
		file_promo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromoId); i {
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
		file_promo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPromosResponse); i {
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
			RawDescriptor: file_promo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_promo_proto_goTypes,
		DependencyIndexes: file_promo_proto_depIdxs,
		MessageInfos:      file_promo_proto_msgTypes,
	}.Build()
	File_promo_proto = out.File
	file_promo_proto_rawDesc = nil
	file_promo_proto_goTypes = nil
	file_promo_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PromoServiceClient is the client API for PromoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PromoServiceClient interface {
	Create(ctx context.Context, in *Promo, opts ...grpc.CallOption) (*PromoId, error)
	Update(ctx context.Context, in *Promo, opts ...grpc.CallOption) (*Empty, error)
	GetAll(ctx context.Context, in *GetAllPromosRequest, opts ...grpc.CallOption) (*GetAllPromosResponse, error)
	Get(ctx context.Context, in *PromoId, opts ...grpc.CallOption) (*Promo, error)
	Delete(ctx context.Context, in *PromoId, opts ...grpc.CallOption) (*Empty, error)
}

type promoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPromoServiceClient(cc grpc.ClientConnInterface) PromoServiceClient {
	return &promoServiceClient{cc}
}

func (c *promoServiceClient) Create(ctx context.Context, in *Promo, opts ...grpc.CallOption) (*PromoId, error) {
	out := new(PromoId)
	err := c.cc.Invoke(ctx, "/genproto.PromoService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promoServiceClient) Update(ctx context.Context, in *Promo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/genproto.PromoService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promoServiceClient) GetAll(ctx context.Context, in *GetAllPromosRequest, opts ...grpc.CallOption) (*GetAllPromosResponse, error) {
	out := new(GetAllPromosResponse)
	err := c.cc.Invoke(ctx, "/genproto.PromoService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promoServiceClient) Get(ctx context.Context, in *PromoId, opts ...grpc.CallOption) (*Promo, error) {
	out := new(Promo)
	err := c.cc.Invoke(ctx, "/genproto.PromoService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promoServiceClient) Delete(ctx context.Context, in *PromoId, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/genproto.PromoService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PromoServiceServer is the server API for PromoService service.
type PromoServiceServer interface {
	Create(context.Context, *Promo) (*PromoId, error)
	Update(context.Context, *Promo) (*Empty, error)
	GetAll(context.Context, *GetAllPromosRequest) (*GetAllPromosResponse, error)
	Get(context.Context, *PromoId) (*Promo, error)
	Delete(context.Context, *PromoId) (*Empty, error)
}

// UnimplementedPromoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPromoServiceServer struct {
}

func (*UnimplementedPromoServiceServer) Create(context.Context, *Promo) (*PromoId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedPromoServiceServer) Update(context.Context, *Promo) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedPromoServiceServer) GetAll(context.Context, *GetAllPromosRequest) (*GetAllPromosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedPromoServiceServer) Get(context.Context, *PromoId) (*Promo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedPromoServiceServer) Delete(context.Context, *PromoId) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterPromoServiceServer(s *grpc.Server, srv PromoServiceServer) {
	s.RegisterService(&_PromoService_serviceDesc, srv)
}

func _PromoService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Promo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromoServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.PromoService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromoServiceServer).Create(ctx, req.(*Promo))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromoService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Promo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromoServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.PromoService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromoServiceServer).Update(ctx, req.(*Promo))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromoService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPromosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromoServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.PromoService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromoServiceServer).GetAll(ctx, req.(*GetAllPromosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromoService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromoId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromoServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.PromoService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromoServiceServer).Get(ctx, req.(*PromoId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromoService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromoId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromoServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.PromoService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromoServiceServer).Delete(ctx, req.(*PromoId))
	}
	return interceptor(ctx, in, info, handler)
}

var _PromoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.PromoService",
	HandlerType: (*PromoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PromoService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PromoService_Update_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _PromoService_GetAll_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _PromoService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PromoService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "promo.proto",
}