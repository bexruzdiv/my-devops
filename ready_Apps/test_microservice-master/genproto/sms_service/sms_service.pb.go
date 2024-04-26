// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: sms_service.proto

package sms_service

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

var File_sms_service_proto protoreflect.FileDescriptor

var file_sms_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x6d, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x73,
	0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x3d, 0x0a, 0x0a, 0x53, 0x6d, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x0d, 0x2e, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6d, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x6d, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_sms_service_proto_goTypes = []interface{}{
	(*Sms)(nil),         // 0: genproto.Sms
	(*empty.Empty)(nil), // 1: google.protobuf.Empty
}
var file_sms_service_proto_depIdxs = []int32{
	0, // 0: genproto.SmsService.Send:input_type -> genproto.Sms
	1, // 1: genproto.SmsService.Send:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sms_service_proto_init() }
func file_sms_service_proto_init() {
	if File_sms_service_proto != nil {
		return
	}
	file_sms_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sms_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sms_service_proto_goTypes,
		DependencyIndexes: file_sms_service_proto_depIdxs,
	}.Build()
	File_sms_service_proto = out.File
	file_sms_service_proto_rawDesc = nil
	file_sms_service_proto_goTypes = nil
	file_sms_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SmsServiceClient is the client API for SmsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SmsServiceClient interface {
	Send(ctx context.Context, in *Sms, opts ...grpc.CallOption) (*empty.Empty, error)
}

type smsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSmsServiceClient(cc grpc.ClientConnInterface) SmsServiceClient {
	return &smsServiceClient{cc}
}

func (c *smsServiceClient) Send(ctx context.Context, in *Sms, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/genproto.SmsService/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmsServiceServer is the server API for SmsService service.
type SmsServiceServer interface {
	Send(context.Context, *Sms) (*empty.Empty, error)
}

// UnimplementedSmsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSmsServiceServer struct {
}

func (*UnimplementedSmsServiceServer) Send(context.Context, *Sms) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}

func RegisterSmsServiceServer(s *grpc.Server, srv SmsServiceServer) {
	s.RegisterService(&_SmsService_serviceDesc, srv)
}

func _SmsService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sms)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.SmsService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServiceServer).Send(ctx, req.(*Sms))
	}
	return interceptor(ctx, in, info, handler)
}

var _SmsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.SmsService",
	HandlerType: (*SmsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _SmsService_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sms_service.proto",
}
