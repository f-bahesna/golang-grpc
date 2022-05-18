// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: cart/cart.proto

package cart

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

type CartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku string `protobuf:"bytes,1,opt,name=sku,proto3" json:"sku,omitempty"`
}

func (x *CartRequest) Reset() {
	*x = CartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartRequest) ProtoMessage() {}

func (x *CartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cart_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartRequest.ProtoReflect.Descriptor instead.
func (*CartRequest) Descriptor() ([]byte, []int) {
	return file_cart_cart_proto_rawDescGZIP(), []int{0}
}

func (x *CartRequest) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

type CartData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku         string `protobuf:"bytes,1,opt,name=sku,proto3" json:"sku,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CartData) Reset() {
	*x = CartData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cart_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartData) ProtoMessage() {}

func (x *CartData) ProtoReflect() protoreflect.Message {
	mi := &file_cart_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartData.ProtoReflect.Descriptor instead.
func (*CartData) Descriptor() ([]byte, []int) {
	return file_cart_cart_proto_rawDescGZIP(), []int{1}
}

func (x *CartData) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

func (x *CartData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CartData) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_cart_cart_proto protoreflect.FileDescriptor

var file_cart_cart_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x1f, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73,
	0x6b, 0x75, 0x22, 0x52, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x6b, 0x75,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x32, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12, 0x2a,
	0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x61, 0x72, 0x74, 0x42, 0x79, 0x53, 0x6b, 0x75, 0x12,
	0x0c, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e,
	0x43, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x2d, 0x62, 0x61, 0x68, 0x65, 0x73,
	0x6e, 0x61, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63,
	0x61, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cart_cart_proto_rawDescOnce sync.Once
	file_cart_cart_proto_rawDescData = file_cart_cart_proto_rawDesc
)

func file_cart_cart_proto_rawDescGZIP() []byte {
	file_cart_cart_proto_rawDescOnce.Do(func() {
		file_cart_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_cart_cart_proto_rawDescData)
	})
	return file_cart_cart_proto_rawDescData
}

var file_cart_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cart_cart_proto_goTypes = []interface{}{
	(*CartRequest)(nil), // 0: CartRequest
	(*CartData)(nil),    // 1: CartData
}
var file_cart_cart_proto_depIdxs = []int32{
	0, // 0: Cart.FindCartBySku:input_type -> CartRequest
	1, // 1: Cart.FindCartBySku:output_type -> CartData
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cart_cart_proto_init() }
func file_cart_cart_proto_init() {
	if File_cart_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cart_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartRequest); i {
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
		file_cart_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartData); i {
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
			RawDescriptor: file_cart_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cart_cart_proto_goTypes,
		DependencyIndexes: file_cart_cart_proto_depIdxs,
		MessageInfos:      file_cart_cart_proto_msgTypes,
	}.Build()
	File_cart_cart_proto = out.File
	file_cart_cart_proto_rawDesc = nil
	file_cart_cart_proto_goTypes = nil
	file_cart_cart_proto_depIdxs = nil
}
