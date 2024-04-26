// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: branch_user.proto

package user_service

import (
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type BranchUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone      string                `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	IsActive   bool                  `protobuf:"varint,4,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	IsBlocked  bool                  `protobuf:"varint,5,opt,name=is_blocked,json=isBlocked,proto3" json:"is_blocked,omitempty"`
	CreatedAt  string                `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  string                `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	ShipperId  string                `protobuf:"bytes,8,opt,name=shipper_id,json=shipperId,proto3" json:"shipper_id,omitempty"`
	BranchId   string                `protobuf:"bytes,9,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	UserRoleId string                `protobuf:"bytes,10,opt,name=user_role_id,json=userRoleId,proto3" json:"user_role_id,omitempty"`
	FcmToken   *wrappers.StringValue `protobuf:"bytes,11,opt,name=fcm_token,json=fcmToken,proto3" json:"fcm_token,omitempty"`
	PlatformId *wrappers.StringValue `protobuf:"bytes,12,opt,name=platform_id,json=platformId,proto3" json:"platform_id,omitempty"`
}

func (x *BranchUser) Reset() {
	*x = BranchUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BranchUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BranchUser) ProtoMessage() {}

func (x *BranchUser) ProtoReflect() protoreflect.Message {
	mi := &file_branch_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BranchUser.ProtoReflect.Descriptor instead.
func (*BranchUser) Descriptor() ([]byte, []int) {
	return file_branch_user_proto_rawDescGZIP(), []int{0}
}

func (x *BranchUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BranchUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BranchUser) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *BranchUser) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *BranchUser) GetIsBlocked() bool {
	if x != nil {
		return x.IsBlocked
	}
	return false
}

func (x *BranchUser) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *BranchUser) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *BranchUser) GetShipperId() string {
	if x != nil {
		return x.ShipperId
	}
	return ""
}

func (x *BranchUser) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *BranchUser) GetUserRoleId() string {
	if x != nil {
		return x.UserRoleId
	}
	return ""
}

func (x *BranchUser) GetFcmToken() *wrappers.StringValue {
	if x != nil {
		return x.FcmToken
	}
	return nil
}

func (x *BranchUser) GetPlatformId() *wrappers.StringValue {
	if x != nil {
		return x.PlatformId
	}
	return nil
}

var File_branch_user_proto protoreflect.FileDescriptor

var file_branch_user_proto_rawDesc = []byte{
	0x0a, 0x11, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x03,
	0x0a, 0x0a, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x39,
	0x0a, 0x09, 0x66, 0x63, 0x6d, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x08, 0x66, 0x63, 0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x3d, 0x0a, 0x0b, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x49, 0x64, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_branch_user_proto_rawDescOnce sync.Once
	file_branch_user_proto_rawDescData = file_branch_user_proto_rawDesc
)

func file_branch_user_proto_rawDescGZIP() []byte {
	file_branch_user_proto_rawDescOnce.Do(func() {
		file_branch_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_branch_user_proto_rawDescData)
	})
	return file_branch_user_proto_rawDescData
}

var file_branch_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_branch_user_proto_goTypes = []interface{}{
	(*BranchUser)(nil),           // 0: genproto.BranchUser
	(*wrappers.StringValue)(nil), // 1: google.protobuf.StringValue
}
var file_branch_user_proto_depIdxs = []int32{
	1, // 0: genproto.BranchUser.fcm_token:type_name -> google.protobuf.StringValue
	1, // 1: genproto.BranchUser.platform_id:type_name -> google.protobuf.StringValue
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_branch_user_proto_init() }
func file_branch_user_proto_init() {
	if File_branch_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_branch_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BranchUser); i {
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
			RawDescriptor: file_branch_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_branch_user_proto_goTypes,
		DependencyIndexes: file_branch_user_proto_depIdxs,
		MessageInfos:      file_branch_user_proto_msgTypes,
	}.Build()
	File_branch_user_proto = out.File
	file_branch_user_proto_rawDesc = nil
	file_branch_user_proto_goTypes = nil
	file_branch_user_proto_depIdxs = nil
}