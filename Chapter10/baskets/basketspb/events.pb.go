// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: basketspb/events.proto

package basketspb

import (
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

type BasketStarted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *BasketStarted) Reset() {
	*x = BasketStarted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basketspb_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasketStarted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasketStarted) ProtoMessage() {}

func (x *BasketStarted) ProtoReflect() protoreflect.Message {
	mi := &file_basketspb_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasketStarted.ProtoReflect.Descriptor instead.
func (*BasketStarted) Descriptor() ([]byte, []int) {
	return file_basketspb_events_proto_rawDescGZIP(), []int{0}
}

func (x *BasketStarted) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BasketStarted) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

type BasketCanceled struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BasketCanceled) Reset() {
	*x = BasketCanceled{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basketspb_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasketCanceled) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasketCanceled) ProtoMessage() {}

func (x *BasketCanceled) ProtoReflect() protoreflect.Message {
	mi := &file_basketspb_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasketCanceled.ProtoReflect.Descriptor instead.
func (*BasketCanceled) Descriptor() ([]byte, []int) {
	return file_basketspb_events_proto_rawDescGZIP(), []int{1}
}

func (x *BasketCanceled) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type BasketCheckedOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId string                   `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	PaymentId  string                   `protobuf:"bytes,3,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	Items      []*BasketCheckedOut_Item `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *BasketCheckedOut) Reset() {
	*x = BasketCheckedOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basketspb_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasketCheckedOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasketCheckedOut) ProtoMessage() {}

func (x *BasketCheckedOut) ProtoReflect() protoreflect.Message {
	mi := &file_basketspb_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasketCheckedOut.ProtoReflect.Descriptor instead.
func (*BasketCheckedOut) Descriptor() ([]byte, []int) {
	return file_basketspb_events_proto_rawDescGZIP(), []int{2}
}

func (x *BasketCheckedOut) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BasketCheckedOut) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *BasketCheckedOut) GetPaymentId() string {
	if x != nil {
		return x.PaymentId
	}
	return ""
}

func (x *BasketCheckedOut) GetItems() []*BasketCheckedOut_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type BasketCheckedOut_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoreId     string  `protobuf:"bytes,1,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	ProductId   string  `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	StoreName   string  `protobuf:"bytes,3,opt,name=store_name,json=storeName,proto3" json:"store_name,omitempty"`
	ProductName string  `protobuf:"bytes,4,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	Price       float64 `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`
	Quantity    int32   `protobuf:"varint,6,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *BasketCheckedOut_Item) Reset() {
	*x = BasketCheckedOut_Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basketspb_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasketCheckedOut_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasketCheckedOut_Item) ProtoMessage() {}

func (x *BasketCheckedOut_Item) ProtoReflect() protoreflect.Message {
	mi := &file_basketspb_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasketCheckedOut_Item.ProtoReflect.Descriptor instead.
func (*BasketCheckedOut_Item) Descriptor() ([]byte, []int) {
	return file_basketspb_events_proto_rawDescGZIP(), []int{2, 0}
}

func (x *BasketCheckedOut_Item) GetStoreId() string {
	if x != nil {
		return x.StoreId
	}
	return ""
}

func (x *BasketCheckedOut_Item) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *BasketCheckedOut_Item) GetStoreName() string {
	if x != nil {
		return x.StoreName
	}
	return ""
}

func (x *BasketCheckedOut_Item) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *BasketCheckedOut_Item) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *BasketCheckedOut_Item) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

var File_basketspb_events_proto protoreflect.FileDescriptor

var file_basketspb_events_proto_rawDesc = []byte{
	0x0a, 0x16, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74,
	0x73, 0x70, 0x62, 0x22, 0x40, 0x0a, 0x0d, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xd1, 0x02, 0x0a, 0x10, 0x42, 0x61, 0x73, 0x6b,
	0x65, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x61,
	0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x1a, 0xb4, 0x01, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0xbe, 0x01, 0x0a, 0x0d,
	0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0x42, 0x0b, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x62, 0x69, 0x73, 0x63, 0x75, 0x6d,
	0x2f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2d, 0x44, 0x72, 0x69, 0x76, 0x65, 0x6e, 0x2d, 0x41, 0x72,
	0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2d, 0x69, 0x6e, 0x2d, 0x47, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x31, 0x30, 0x2f, 0x62,
	0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62,
	0x2f, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58,
	0xaa, 0x02, 0x09, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0xca, 0x02, 0x09, 0x42,
	0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0xe2, 0x02, 0x15, 0x42, 0x61, 0x73, 0x6b, 0x65,
	0x74, 0x73, 0x70, 0x62, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x09, 0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_basketspb_events_proto_rawDescOnce sync.Once
	file_basketspb_events_proto_rawDescData = file_basketspb_events_proto_rawDesc
)

func file_basketspb_events_proto_rawDescGZIP() []byte {
	file_basketspb_events_proto_rawDescOnce.Do(func() {
		file_basketspb_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_basketspb_events_proto_rawDescData)
	})
	return file_basketspb_events_proto_rawDescData
}

var file_basketspb_events_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_basketspb_events_proto_goTypes = []interface{}{
	(*BasketStarted)(nil),         // 0: basketspb.BasketStarted
	(*BasketCanceled)(nil),        // 1: basketspb.BasketCanceled
	(*BasketCheckedOut)(nil),      // 2: basketspb.BasketCheckedOut
	(*BasketCheckedOut_Item)(nil), // 3: basketspb.BasketCheckedOut.Item
}
var file_basketspb_events_proto_depIdxs = []int32{
	3, // 0: basketspb.BasketCheckedOut.items:type_name -> basketspb.BasketCheckedOut.Item
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_basketspb_events_proto_init() }
func file_basketspb_events_proto_init() {
	if File_basketspb_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_basketspb_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasketStarted); i {
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
		file_basketspb_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasketCanceled); i {
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
		file_basketspb_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasketCheckedOut); i {
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
		file_basketspb_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasketCheckedOut_Item); i {
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
			RawDescriptor: file_basketspb_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_basketspb_events_proto_goTypes,
		DependencyIndexes: file_basketspb_events_proto_depIdxs,
		MessageInfos:      file_basketspb_events_proto_msgTypes,
	}.Build()
	File_basketspb_events_proto = out.File
	file_basketspb_events_proto_rawDesc = nil
	file_basketspb_events_proto_goTypes = nil
	file_basketspb_events_proto_depIdxs = nil
}
