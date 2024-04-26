// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: tracking.proto

package tracking_service

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

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Long float64 `protobuf:"fixed64,1,opt,name=long,proto3" json:"long,omitempty"`
	Lat  float64 `protobuf:"fixed64,2,opt,name=lat,proto3" json:"lat,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_tracking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_tracking_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetLong() float64 {
	if x != nil {
		return x.Long
	}
	return 0
}

func (x *Location) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

type Courier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Phone     string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	IsActive  bool   `protobuf:"varint,4,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	IsOnline  bool   `protobuf:"varint,5,opt,name=is_online,json=isOnline,proto3" json:"is_online,omitempty"`
}

func (x *Courier) Reset() {
	*x = Courier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Courier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Courier) ProtoMessage() {}

func (x *Courier) ProtoReflect() protoreflect.Message {
	mi := &file_tracking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Courier.ProtoReflect.Descriptor instead.
func (*Courier) Descriptor() ([]byte, []int) {
	return file_tracking_proto_rawDescGZIP(), []int{1}
}

func (x *Courier) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Courier) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Courier) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Courier) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *Courier) GetIsOnline() bool {
	if x != nil {
		return x.IsOnline
	}
	return false
}

type Tracking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CourierId string    `protobuf:"bytes,2,opt,name=courier_id,json=courierId,proto3" json:"courier_id,omitempty"`
	CreatedAt string    `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Location  *Location `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	UpdatedAt string    `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Tracking) Reset() {
	*x = Tracking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tracking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tracking) ProtoMessage() {}

func (x *Tracking) ProtoReflect() protoreflect.Message {
	mi := &file_tracking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tracking.ProtoReflect.Descriptor instead.
func (*Tracking) Descriptor() ([]byte, []int) {
	return file_tracking_proto_rawDescGZIP(), []int{2}
}

func (x *Tracking) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tracking) GetCourierId() string {
	if x != nil {
		return x.CourierId
	}
	return ""
}

func (x *Tracking) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Tracking) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Tracking) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderNum uint64 `protobuf:"varint,2,opt,name=order_num,json=orderNum,proto3" json:"order_num,omitempty"`
	StatusId string `protobuf:"bytes,3,opt,name=status_id,json=statusId,proto3" json:"status_id,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_tracking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_tracking_proto_rawDescGZIP(), []int{3}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetOrderNum() uint64 {
	if x != nil {
		return x.OrderNum
	}
	return 0
}

func (x *Order) GetStatusId() string {
	if x != nil {
		return x.StatusId
	}
	return ""
}

type CourierTracking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Courier     *Courier  `protobuf:"bytes,1,opt,name=courier,proto3" json:"courier,omitempty"`
	Tracking    *Tracking `protobuf:"bytes,2,opt,name=tracking,proto3" json:"tracking,omitempty"`
	Orders      []*Order  `protobuf:"bytes,3,rep,name=orders,proto3" json:"orders,omitempty"`
	OrdersCount uint64    `protobuf:"varint,4,opt,name=orders_count,json=ordersCount,proto3" json:"orders_count,omitempty"`
}

func (x *CourierTracking) Reset() {
	*x = CourierTracking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracking_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourierTracking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourierTracking) ProtoMessage() {}

func (x *CourierTracking) ProtoReflect() protoreflect.Message {
	mi := &file_tracking_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourierTracking.ProtoReflect.Descriptor instead.
func (*CourierTracking) Descriptor() ([]byte, []int) {
	return file_tracking_proto_rawDescGZIP(), []int{4}
}

func (x *CourierTracking) GetCourier() *Courier {
	if x != nil {
		return x.Courier
	}
	return nil
}

func (x *CourierTracking) GetTracking() *Tracking {
	if x != nil {
		return x.Tracking
	}
	return nil
}

func (x *CourierTracking) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

func (x *CourierTracking) GetOrdersCount() uint64 {
	if x != nil {
		return x.OrdersCount
	}
	return 0
}

var File_tracking_proto protoreflect.FileDescriptor

var file_tracking_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x08, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x22, 0x95, 0x01, 0x0a,
	0x07, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x6f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x4f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x22, 0xa7, 0x01, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e,
	0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x2e, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x51,
	0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x4e, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49,
	0x64, 0x22, 0xba, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x2b, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x72, 0x69,
	0x65, 0x72, 0x12, 0x2e, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69,
	0x6e, 0x67, 0x12, 0x27, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x1b,
	0x5a, 0x19, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_tracking_proto_rawDescOnce sync.Once
	file_tracking_proto_rawDescData = file_tracking_proto_rawDesc
)

func file_tracking_proto_rawDescGZIP() []byte {
	file_tracking_proto_rawDescOnce.Do(func() {
		file_tracking_proto_rawDescData = protoimpl.X.CompressGZIP(file_tracking_proto_rawDescData)
	})
	return file_tracking_proto_rawDescData
}

var file_tracking_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_tracking_proto_goTypes = []interface{}{
	(*Location)(nil),        // 0: genproto.Location
	(*Courier)(nil),         // 1: genproto.Courier
	(*Tracking)(nil),        // 2: genproto.Tracking
	(*Order)(nil),           // 3: genproto.Order
	(*CourierTracking)(nil), // 4: genproto.CourierTracking
}
var file_tracking_proto_depIdxs = []int32{
	0, // 0: genproto.Tracking.location:type_name -> genproto.Location
	1, // 1: genproto.CourierTracking.courier:type_name -> genproto.Courier
	2, // 2: genproto.CourierTracking.tracking:type_name -> genproto.Tracking
	3, // 3: genproto.CourierTracking.orders:type_name -> genproto.Order
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tracking_proto_init() }
func file_tracking_proto_init() {
	if File_tracking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tracking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_tracking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Courier); i {
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
		file_tracking_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tracking); i {
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
		file_tracking_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_tracking_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourierTracking); i {
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
			RawDescriptor: file_tracking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tracking_proto_goTypes,
		DependencyIndexes: file_tracking_proto_depIdxs,
		MessageInfos:      file_tracking_proto_msgTypes,
	}.Build()
	File_tracking_proto = out.File
	file_tracking_proto_rawDesc = nil
	file_tracking_proto_goTypes = nil
	file_tracking_proto_depIdxs = nil
}
