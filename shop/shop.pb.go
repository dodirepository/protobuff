// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: shop/shop.proto

package shop

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

type ShopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopId   int32  `protobuf:"varint,1,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	ShopName string `protobuf:"bytes,2,opt,name=shop_name,json=shopName,proto3" json:"shop_name,omitempty"`
}

func (x *ShopResponse) Reset() {
	*x = ShopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_shop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopResponse) ProtoMessage() {}

func (x *ShopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shop_shop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopResponse.ProtoReflect.Descriptor instead.
func (*ShopResponse) Descriptor() ([]byte, []int) {
	return file_shop_shop_proto_rawDescGZIP(), []int{0}
}

func (x *ShopResponse) GetShopId() int32 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *ShopResponse) GetShopName() string {
	if x != nil {
		return x.ShopName
	}
	return ""
}

// Define the request message
type ShopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopId int32 `protobuf:"varint,1,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
}

func (x *ShopRequest) Reset() {
	*x = ShopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_shop_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShopRequest) ProtoMessage() {}

func (x *ShopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shop_shop_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShopRequest.ProtoReflect.Descriptor instead.
func (*ShopRequest) Descriptor() ([]byte, []int) {
	return file_shop_shop_proto_rawDescGZIP(), []int{1}
}

func (x *ShopRequest) GetShopId() int32 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

var File_shop_shop_proto protoreflect.FileDescriptor

var file_shop_shop_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x22, 0x44, 0x0a, 0x0c, 0x53, 0x68, 0x6f, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a,
	0x0b, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x68, 0x6f, 0x70, 0x49, 0x64, 0x32, 0x46, 0x0a, 0x0b, 0x53, 0x68, 0x6f, 0x70, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x11, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x53, 0x68,
	0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x73, 0x68, 0x6f, 0x70,
	0x2e, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2a, 0x5a,
	0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x64, 0x69,
	0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x66, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_shop_shop_proto_rawDescOnce sync.Once
	file_shop_shop_proto_rawDescData = file_shop_shop_proto_rawDesc
)

func file_shop_shop_proto_rawDescGZIP() []byte {
	file_shop_shop_proto_rawDescOnce.Do(func() {
		file_shop_shop_proto_rawDescData = protoimpl.X.CompressGZIP(file_shop_shop_proto_rawDescData)
	})
	return file_shop_shop_proto_rawDescData
}

var file_shop_shop_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_shop_shop_proto_goTypes = []any{
	(*ShopResponse)(nil), // 0: shop.ShopResponse
	(*ShopRequest)(nil),  // 1: shop.ShopRequest
}
var file_shop_shop_proto_depIdxs = []int32{
	1, // 0: shop.ShopService.GetShopDetails:input_type -> shop.ShopRequest
	0, // 1: shop.ShopService.GetShopDetails:output_type -> shop.ShopResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shop_shop_proto_init() }
func file_shop_shop_proto_init() {
	if File_shop_shop_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shop_shop_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ShopResponse); i {
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
		file_shop_shop_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ShopRequest); i {
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
			RawDescriptor: file_shop_shop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shop_shop_proto_goTypes,
		DependencyIndexes: file_shop_shop_proto_depIdxs,
		MessageInfos:      file_shop_shop_proto_msgTypes,
	}.Build()
	File_shop_shop_proto = out.File
	file_shop_shop_proto_rawDesc = nil
	file_shop_shop_proto_goTypes = nil
	file_shop_shop_proto_depIdxs = nil
}
