// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: sms_payments_service.proto

package sms_service

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

type SmsPaymentsId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SmsPaymentsId) Reset() {
	*x = SmsPaymentsId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_payments_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SmsPaymentsId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmsPaymentsId) ProtoMessage() {}

func (x *SmsPaymentsId) ProtoReflect() protoreflect.Message {
	mi := &file_sms_payments_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmsPaymentsId.ProtoReflect.Descriptor instead.
func (*SmsPaymentsId) Descriptor() ([]byte, []int) {
	return file_sms_payments_service_proto_rawDescGZIP(), []int{0}
}

func (x *SmsPaymentsId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAllSmsPaymentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShipperId string `protobuf:"bytes,1,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
	Page      int64  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit     int64  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetAllSmsPaymentsRequest) Reset() {
	*x = GetAllSmsPaymentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_payments_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllSmsPaymentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllSmsPaymentsRequest) ProtoMessage() {}

func (x *GetAllSmsPaymentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sms_payments_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllSmsPaymentsRequest.ProtoReflect.Descriptor instead.
func (*GetAllSmsPaymentsRequest) Descriptor() ([]byte, []int) {
	return file_sms_payments_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllSmsPaymentsRequest) GetShipperId() string {
	if x != nil {
		return x.ShipperId
	}
	return ""
}

func (x *GetAllSmsPaymentsRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllSmsPaymentsRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAllSmsPaymentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Smspayments []*SmsPayments `protobuf:"bytes,1,rep,name=smspayments,proto3" json:"smspayments,omitempty"`
	Count       int64          `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllSmsPaymentsResponse) Reset() {
	*x = GetAllSmsPaymentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_payments_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllSmsPaymentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllSmsPaymentsResponse) ProtoMessage() {}

func (x *GetAllSmsPaymentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sms_payments_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllSmsPaymentsResponse.ProtoReflect.Descriptor instead.
func (*GetAllSmsPaymentsResponse) Descriptor() ([]byte, []int) {
	return file_sms_payments_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllSmsPaymentsResponse) GetSmspayments() []*SmsPayments {
	if x != nil {
		return x.Smspayments
	}
	return nil
}

func (x *GetAllSmsPaymentsResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_sms_payments_service_proto protoreflect.FileDescriptor

var file_sms_payments_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x6d, 0x73, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x73, 0x6d, 0x73, 0x5f, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x0d, 0x53, 0x6d,
	0x73, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x63, 0x0a, 0x18, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x6d, 0x73, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x70,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x69,
	0x70, 0x70, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x22, 0x6a, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x6d, 0x73, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a,
	0x0b, 0x73, 0x6d, 0x73, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6d,
	0x73, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x0b, 0x73, 0x6d, 0x73, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xd8, 0x01, 0x0a,
	0x12, 0x53, 0x6d, 0x73, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6d, 0x73, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x1a, 0x17, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x6d, 0x73, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x64, 0x12, 0x35, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x6d, 0x73, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x64, 0x1a, 0x15, 0x2e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6d, 0x73, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x51, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x22,
	0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x53, 0x6d, 0x73, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x53, 0x6d, 0x73, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x16, 0x5a, 0x14, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6d, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sms_payments_service_proto_rawDescOnce sync.Once
	file_sms_payments_service_proto_rawDescData = file_sms_payments_service_proto_rawDesc
)

func file_sms_payments_service_proto_rawDescGZIP() []byte {
	file_sms_payments_service_proto_rawDescOnce.Do(func() {
		file_sms_payments_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_sms_payments_service_proto_rawDescData)
	})
	return file_sms_payments_service_proto_rawDescData
}

var file_sms_payments_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sms_payments_service_proto_goTypes = []interface{}{
	(*SmsPaymentsId)(nil),             // 0: genproto.SmsPaymentsId
	(*GetAllSmsPaymentsRequest)(nil),  // 1: genproto.GetAllSmsPaymentsRequest
	(*GetAllSmsPaymentsResponse)(nil), // 2: genproto.GetAllSmsPaymentsResponse
	(*SmsPayments)(nil),               // 3: genproto.SmsPayments
}
var file_sms_payments_service_proto_depIdxs = []int32{
	3, // 0: genproto.GetAllSmsPaymentsResponse.smspayments:type_name -> genproto.SmsPayments
	3, // 1: genproto.SmsPaymentsService.Create:input_type -> genproto.SmsPayments
	0, // 2: genproto.SmsPaymentsService.Get:input_type -> genproto.SmsPaymentsId
	1, // 3: genproto.SmsPaymentsService.GetAll:input_type -> genproto.GetAllSmsPaymentsRequest
	0, // 4: genproto.SmsPaymentsService.Create:output_type -> genproto.SmsPaymentsId
	3, // 5: genproto.SmsPaymentsService.Get:output_type -> genproto.SmsPayments
	2, // 6: genproto.SmsPaymentsService.GetAll:output_type -> genproto.GetAllSmsPaymentsResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sms_payments_service_proto_init() }
func file_sms_payments_service_proto_init() {
	if File_sms_payments_service_proto != nil {
		return
	}
	file_sms_payments_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sms_payments_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SmsPaymentsId); i {
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
		file_sms_payments_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllSmsPaymentsRequest); i {
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
		file_sms_payments_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllSmsPaymentsResponse); i {
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
			RawDescriptor: file_sms_payments_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sms_payments_service_proto_goTypes,
		DependencyIndexes: file_sms_payments_service_proto_depIdxs,
		MessageInfos:      file_sms_payments_service_proto_msgTypes,
	}.Build()
	File_sms_payments_service_proto = out.File
	file_sms_payments_service_proto_rawDesc = nil
	file_sms_payments_service_proto_goTypes = nil
	file_sms_payments_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SmsPaymentsServiceClient is the client API for SmsPaymentsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SmsPaymentsServiceClient interface {
	Create(ctx context.Context, in *SmsPayments, opts ...grpc.CallOption) (*SmsPaymentsId, error)
	Get(ctx context.Context, in *SmsPaymentsId, opts ...grpc.CallOption) (*SmsPayments, error)
	GetAll(ctx context.Context, in *GetAllSmsPaymentsRequest, opts ...grpc.CallOption) (*GetAllSmsPaymentsResponse, error)
}

type smsPaymentsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSmsPaymentsServiceClient(cc grpc.ClientConnInterface) SmsPaymentsServiceClient {
	return &smsPaymentsServiceClient{cc}
}

func (c *smsPaymentsServiceClient) Create(ctx context.Context, in *SmsPayments, opts ...grpc.CallOption) (*SmsPaymentsId, error) {
	out := new(SmsPaymentsId)
	err := c.cc.Invoke(ctx, "/genproto.SmsPaymentsService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsPaymentsServiceClient) Get(ctx context.Context, in *SmsPaymentsId, opts ...grpc.CallOption) (*SmsPayments, error) {
	out := new(SmsPayments)
	err := c.cc.Invoke(ctx, "/genproto.SmsPaymentsService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsPaymentsServiceClient) GetAll(ctx context.Context, in *GetAllSmsPaymentsRequest, opts ...grpc.CallOption) (*GetAllSmsPaymentsResponse, error) {
	out := new(GetAllSmsPaymentsResponse)
	err := c.cc.Invoke(ctx, "/genproto.SmsPaymentsService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmsPaymentsServiceServer is the server API for SmsPaymentsService service.
type SmsPaymentsServiceServer interface {
	Create(context.Context, *SmsPayments) (*SmsPaymentsId, error)
	Get(context.Context, *SmsPaymentsId) (*SmsPayments, error)
	GetAll(context.Context, *GetAllSmsPaymentsRequest) (*GetAllSmsPaymentsResponse, error)
}

// UnimplementedSmsPaymentsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSmsPaymentsServiceServer struct {
}

func (*UnimplementedSmsPaymentsServiceServer) Create(context.Context, *SmsPayments) (*SmsPaymentsId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedSmsPaymentsServiceServer) Get(context.Context, *SmsPaymentsId) (*SmsPayments, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedSmsPaymentsServiceServer) GetAll(context.Context, *GetAllSmsPaymentsRequest) (*GetAllSmsPaymentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}

func RegisterSmsPaymentsServiceServer(s *grpc.Server, srv SmsPaymentsServiceServer) {
	s.RegisterService(&_SmsPaymentsService_serviceDesc, srv)
}

func _SmsPaymentsService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmsPayments)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsPaymentsServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.SmsPaymentsService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsPaymentsServiceServer).Create(ctx, req.(*SmsPayments))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmsPaymentsService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SmsPaymentsId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsPaymentsServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.SmsPaymentsService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsPaymentsServiceServer).Get(ctx, req.(*SmsPaymentsId))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmsPaymentsService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllSmsPaymentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsPaymentsServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.SmsPaymentsService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsPaymentsServiceServer).GetAll(ctx, req.(*GetAllSmsPaymentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SmsPaymentsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.SmsPaymentsService",
	HandlerType: (*SmsPaymentsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SmsPaymentsService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _SmsPaymentsService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _SmsPaymentsService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sms_payments_service.proto",
}
