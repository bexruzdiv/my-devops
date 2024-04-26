// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: rkeeper.proto

package rkeeper_service

import (
	proto "github.com/golang/protobuf/proto"
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

type RkeeperBranch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Guid      string `protobuf:"bytes,2,opt,name=guid,proto3" json:"guid,omitempty"`
	BranchId  string `protobuf:"bytes,3,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	ShipperId string `protobuf:"bytes,4,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
	Url       string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	TableId   string `protobuf:"bytes,6,opt,name=table_id,json=tableId,proto3" json:"table_id,omitempty"`
}

func (x *RkeeperBranch) Reset() {
	*x = RkeeperBranch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rkeeper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RkeeperBranch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RkeeperBranch) ProtoMessage() {}

func (x *RkeeperBranch) ProtoReflect() protoreflect.Message {
	mi := &file_rkeeper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RkeeperBranch.ProtoReflect.Descriptor instead.
func (*RkeeperBranch) Descriptor() ([]byte, []int) {
	return file_rkeeper_proto_rawDescGZIP(), []int{0}
}

func (x *RkeeperBranch) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RkeeperBranch) GetGuid() string {
	if x != nil {
		return x.Guid
	}
	return ""
}

func (x *RkeeperBranch) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *RkeeperBranch) GetShipperId() string {
	if x != nil {
		return x.ShipperId
	}
	return ""
}

func (x *RkeeperBranch) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *RkeeperBranch) GetTableId() string {
	if x != nil {
		return x.TableId
	}
	return ""
}

type Menu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	CategoryPath string `protobuf:"bytes,5,opt,name=category_path,json=categoryPath,proto3" json:"category_path,omitempty"`
}

func (x *Menu) Reset() {
	*x = Menu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rkeeper_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Menu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Menu) ProtoMessage() {}

func (x *Menu) ProtoReflect() protoreflect.Message {
	mi := &file_rkeeper_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Menu.ProtoReflect.Descriptor instead.
func (*Menu) Descriptor() ([]byte, []int) {
	return file_rkeeper_proto_rawDescGZIP(), []int{1}
}

func (x *Menu) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Menu) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Menu) GetCategoryPath() string {
	if x != nil {
		return x.CategoryPath
	}
	return ""
}

type RkeeperOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId        string `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	RkeeperOrderId string `protobuf:"bytes,3,opt,name=rkeeper_order_id,json=rkeeperOrderId,proto3" json:"rkeeper_order_id,omitempty"`
	BranchId       string `protobuf:"bytes,4,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	ShipperId      string `protobuf:"bytes,5,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
	CreatedAt      string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	SentAt         string `protobuf:"bytes,7,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
	SendCount      string `protobuf:"bytes,8,opt,name=send_count,json=sendCount,proto3" json:"send_count,omitempty"`
}

func (x *RkeeperOrder) Reset() {
	*x = RkeeperOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rkeeper_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RkeeperOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RkeeperOrder) ProtoMessage() {}

func (x *RkeeperOrder) ProtoReflect() protoreflect.Message {
	mi := &file_rkeeper_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RkeeperOrder.ProtoReflect.Descriptor instead.
func (*RkeeperOrder) Descriptor() ([]byte, []int) {
	return file_rkeeper_proto_rawDescGZIP(), []int{2}
}

func (x *RkeeperOrder) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RkeeperOrder) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *RkeeperOrder) GetRkeeperOrderId() string {
	if x != nil {
		return x.RkeeperOrderId
	}
	return ""
}

func (x *RkeeperOrder) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *RkeeperOrder) GetShipperId() string {
	if x != nil {
		return x.ShipperId
	}
	return ""
}

func (x *RkeeperOrder) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *RkeeperOrder) GetSentAt() string {
	if x != nil {
		return x.SentAt
	}
	return ""
}

func (x *RkeeperOrder) GetSendCount() string {
	if x != nil {
		return x.SendCount
	}
	return ""
}

var File_rkeeper_proto protoreflect.FileDescriptor

var file_rkeeper_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a, 0x0d, 0x52, 0x6b,
	0x65, 0x65, 0x70, 0x65, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x67,
	0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x67, 0x75, 0x69, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x19, 0x0a,
	0x08, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x04, 0x4d, 0x65, 0x6e, 0x75,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x50, 0x61, 0x74, 0x68, 0x22, 0xf6, 0x01, 0x0a, 0x0c, 0x52, 0x6b,
	0x65, 0x65, 0x70, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x72, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72,
	0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x72, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x65,
	0x6e, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e,
	0x74, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72,
	0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rkeeper_proto_rawDescOnce sync.Once
	file_rkeeper_proto_rawDescData = file_rkeeper_proto_rawDesc
)

func file_rkeeper_proto_rawDescGZIP() []byte {
	file_rkeeper_proto_rawDescOnce.Do(func() {
		file_rkeeper_proto_rawDescData = protoimpl.X.CompressGZIP(file_rkeeper_proto_rawDescData)
	})
	return file_rkeeper_proto_rawDescData
}

var file_rkeeper_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rkeeper_proto_goTypes = []interface{}{
	(*RkeeperBranch)(nil), // 0: genproto.RkeeperBranch
	(*Menu)(nil),          // 1: genproto.Menu
	(*RkeeperOrder)(nil),  // 2: genproto.RkeeperOrder
}
var file_rkeeper_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rkeeper_proto_init() }
func file_rkeeper_proto_init() {
	if File_rkeeper_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rkeeper_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RkeeperBranch); i {
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
		file_rkeeper_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Menu); i {
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
		file_rkeeper_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RkeeperOrder); i {
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
			RawDescriptor: file_rkeeper_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rkeeper_proto_goTypes,
		DependencyIndexes: file_rkeeper_proto_depIdxs,
		MessageInfos:      file_rkeeper_proto_msgTypes,
	}.Build()
	File_rkeeper_proto = out.File
	file_rkeeper_proto_rawDesc = nil
	file_rkeeper_proto_goTypes = nil
	file_rkeeper_proto_depIdxs = nil
}