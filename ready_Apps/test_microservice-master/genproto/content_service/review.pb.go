// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: review.proto

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

type Review struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message        *Language `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Type           string    `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	ShipperId      string    `protobuf:"bytes,4,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
	RelatedSubject string    `protobuf:"bytes,5,opt,name=related_subject,json=relatedSubject,proto3" json:"related_subject,omitempty"`
	Active         bool      `protobuf:"varint,6,opt,name=active,proto3" json:"active,omitempty"`
	CreatedAt      string    `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt      string    `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Review) Reset() {
	*x = Review{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Review) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Review) ProtoMessage() {}

func (x *Review) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Review.ProtoReflect.Descriptor instead.
func (*Review) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{0}
}

func (x *Review) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Review) GetMessage() *Language {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *Review) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Review) GetShipperId() string {
	if x != nil {
		return x.ShipperId
	}
	return ""
}

func (x *Review) GetRelatedSubject() string {
	if x != nil {
		return x.RelatedSubject
	}
	return ""
}

func (x *Review) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *Review) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Review) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type GetAllReviewsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reviews []*Review `protobuf:"bytes,1,rep,name=reviews,proto3" json:"reviews,omitempty"`
	Count   int64     `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Search  string    `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetAllReviewsResponse) Reset() {
	*x = GetAllReviewsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllReviewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllReviewsResponse) ProtoMessage() {}

func (x *GetAllReviewsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllReviewsResponse.ProtoReflect.Descriptor instead.
func (*GetAllReviewsResponse) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllReviewsResponse) GetReviews() []*Review {
	if x != nil {
		return x.Reviews
	}
	return nil
}

func (x *GetAllReviewsResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetAllReviewsResponse) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetAllReviewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page      int64  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Search    string `protobuf:"bytes,2,opt,name=search,proto3" json:"search,omitempty"`
	Limit     int64  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	ShipperId string `protobuf:"bytes,4,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
}

func (x *GetAllReviewsRequest) Reset() {
	*x = GetAllReviewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllReviewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllReviewsRequest) ProtoMessage() {}

func (x *GetAllReviewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllReviewsRequest.ProtoReflect.Descriptor instead.
func (*GetAllReviewsRequest) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllReviewsRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllReviewsRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *GetAllReviewsRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllReviewsRequest) GetShipperId() string {
	if x != nil {
		return x.ShipperId
	}
	return ""
}

type ReviewId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReviewId) Reset() {
	*x = ReviewId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReviewId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewId) ProtoMessage() {}

func (x *ReviewId) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewId.ProtoReflect.Descriptor instead.
func (*ReviewId) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{3}
}

func (x *ReviewId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_review_proto protoreflect.FileDescriptor

var file_review_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x2c, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x65, 0x64, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x71, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x07, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x22, 0x77, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x49, 0x64, 0x22, 0x1a, 0x0a,
	0x08, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x9d, 0x02, 0x0a, 0x0d, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x1a, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x64, 0x22, 0x00, 0x12, 0x2d, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x1a, 0x0f, 0x2e, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x06,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x1e, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_review_proto_rawDescOnce sync.Once
	file_review_proto_rawDescData = file_review_proto_rawDesc
)

func file_review_proto_rawDescGZIP() []byte {
	file_review_proto_rawDescOnce.Do(func() {
		file_review_proto_rawDescData = protoimpl.X.CompressGZIP(file_review_proto_rawDescData)
	})
	return file_review_proto_rawDescData
}

var file_review_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_review_proto_goTypes = []interface{}{
	(*Review)(nil),                // 0: genproto.Review
	(*GetAllReviewsResponse)(nil), // 1: genproto.GetAllReviewsResponse
	(*GetAllReviewsRequest)(nil),  // 2: genproto.GetAllReviewsRequest
	(*ReviewId)(nil),              // 3: genproto.ReviewId
	(*Language)(nil),              // 4: genproto.Language
	(*Empty)(nil),                 // 5: genproto.Empty
}
var file_review_proto_depIdxs = []int32{
	4, // 0: genproto.Review.message:type_name -> genproto.Language
	0, // 1: genproto.GetAllReviewsResponse.reviews:type_name -> genproto.Review
	0, // 2: genproto.ReviewService.Create:input_type -> genproto.Review
	0, // 3: genproto.ReviewService.Update:input_type -> genproto.Review
	2, // 4: genproto.ReviewService.GetAll:input_type -> genproto.GetAllReviewsRequest
	3, // 5: genproto.ReviewService.Get:input_type -> genproto.ReviewId
	3, // 6: genproto.ReviewService.Delete:input_type -> genproto.ReviewId
	3, // 7: genproto.ReviewService.Create:output_type -> genproto.ReviewId
	5, // 8: genproto.ReviewService.Update:output_type -> genproto.Empty
	1, // 9: genproto.ReviewService.GetAll:output_type -> genproto.GetAllReviewsResponse
	0, // 10: genproto.ReviewService.Get:output_type -> genproto.Review
	5, // 11: genproto.ReviewService.Delete:output_type -> genproto.Empty
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_review_proto_init() }
func file_review_proto_init() {
	if File_review_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_review_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Review); i {
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
		file_review_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllReviewsResponse); i {
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
		file_review_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllReviewsRequest); i {
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
		file_review_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReviewId); i {
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
			RawDescriptor: file_review_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_review_proto_goTypes,
		DependencyIndexes: file_review_proto_depIdxs,
		MessageInfos:      file_review_proto_msgTypes,
	}.Build()
	File_review_proto = out.File
	file_review_proto_rawDesc = nil
	file_review_proto_goTypes = nil
	file_review_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ReviewServiceClient is the client API for ReviewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReviewServiceClient interface {
	Create(ctx context.Context, in *Review, opts ...grpc.CallOption) (*ReviewId, error)
	Update(ctx context.Context, in *Review, opts ...grpc.CallOption) (*Empty, error)
	GetAll(ctx context.Context, in *GetAllReviewsRequest, opts ...grpc.CallOption) (*GetAllReviewsResponse, error)
	Get(ctx context.Context, in *ReviewId, opts ...grpc.CallOption) (*Review, error)
	Delete(ctx context.Context, in *ReviewId, opts ...grpc.CallOption) (*Empty, error)
}

type reviewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReviewServiceClient(cc grpc.ClientConnInterface) ReviewServiceClient {
	return &reviewServiceClient{cc}
}

func (c *reviewServiceClient) Create(ctx context.Context, in *Review, opts ...grpc.CallOption) (*ReviewId, error) {
	out := new(ReviewId)
	err := c.cc.Invoke(ctx, "/genproto.ReviewService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) Update(ctx context.Context, in *Review, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/genproto.ReviewService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) GetAll(ctx context.Context, in *GetAllReviewsRequest, opts ...grpc.CallOption) (*GetAllReviewsResponse, error) {
	out := new(GetAllReviewsResponse)
	err := c.cc.Invoke(ctx, "/genproto.ReviewService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) Get(ctx context.Context, in *ReviewId, opts ...grpc.CallOption) (*Review, error) {
	out := new(Review)
	err := c.cc.Invoke(ctx, "/genproto.ReviewService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reviewServiceClient) Delete(ctx context.Context, in *ReviewId, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/genproto.ReviewService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReviewServiceServer is the server API for ReviewService service.
type ReviewServiceServer interface {
	Create(context.Context, *Review) (*ReviewId, error)
	Update(context.Context, *Review) (*Empty, error)
	GetAll(context.Context, *GetAllReviewsRequest) (*GetAllReviewsResponse, error)
	Get(context.Context, *ReviewId) (*Review, error)
	Delete(context.Context, *ReviewId) (*Empty, error)
}

// UnimplementedReviewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedReviewServiceServer struct {
}

func (*UnimplementedReviewServiceServer) Create(context.Context, *Review) (*ReviewId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedReviewServiceServer) Update(context.Context, *Review) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedReviewServiceServer) GetAll(context.Context, *GetAllReviewsRequest) (*GetAllReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedReviewServiceServer) Get(context.Context, *ReviewId) (*Review, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedReviewServiceServer) Delete(context.Context, *ReviewId) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterReviewServiceServer(s *grpc.Server, srv ReviewServiceServer) {
	s.RegisterService(&_ReviewService_serviceDesc, srv)
}

func _ReviewService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Review)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ReviewService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).Create(ctx, req.(*Review))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Review)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ReviewService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).Update(ctx, req.(*Review))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ReviewService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).GetAll(ctx, req.(*GetAllReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ReviewService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).Get(ctx, req.(*ReviewId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.ReviewService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).Delete(ctx, req.(*ReviewId))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReviewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.ReviewService",
	HandlerType: (*ReviewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ReviewService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ReviewService_Update_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ReviewService_GetAll_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ReviewService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ReviewService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "review.proto",
}
