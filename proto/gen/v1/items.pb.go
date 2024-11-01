// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: v1/items.proto

package genv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Item_Kind int32

const (
	Item_KIND_UNSPECIFIED Item_Kind = 0
	Item_KIND_GAME        Item_Kind = 1
)

// Enum value maps for Item_Kind.
var (
	Item_Kind_name = map[int32]string{
		0: "KIND_UNSPECIFIED",
		1: "KIND_GAME",
	}
	Item_Kind_value = map[string]int32{
		"KIND_UNSPECIFIED": 0,
		"KIND_GAME":        1,
	}
)

func (x Item_Kind) Enum() *Item_Kind {
	p := new(Item_Kind)
	*p = x
	return p
}

func (x Item_Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Item_Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_items_proto_enumTypes[0].Descriptor()
}

func (Item_Kind) Type() protoreflect.EnumType {
	return &file_v1_items_proto_enumTypes[0]
}

func (x Item_Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Item_Kind.Descriptor instead.
func (Item_Kind) EnumDescriptor() ([]byte, []int) {
	return file_v1_items_proto_rawDescGZIP(), []int{0, 0}
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Kind  Item_Kind `protobuf:"varint,2,opt,name=kind,proto3,enum=v1.Item_Kind" json:"kind,omitempty"`
	Title string    `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_items_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_v1_items_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_v1_items_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Item) GetKind() Item_Kind {
	if x != nil {
		return x.Kind
	}
	return Item_KIND_UNSPECIFIED
}

func (x *Item) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type GetItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetItemRequest) Reset() {
	*x = GetItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_items_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemRequest) ProtoMessage() {}

func (x *GetItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_items_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemRequest.ProtoReflect.Descriptor instead.
func (*GetItemRequest) Descriptor() ([]byte, []int) {
	return file_v1_items_proto_rawDescGZIP(), []int{1}
}

func (x *GetItemRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ImportGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *ImportGameRequest) Reset() {
	*x = ImportGameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_items_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportGameRequest) ProtoMessage() {}

func (x *ImportGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_items_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportGameRequest.ProtoReflect.Descriptor instead.
func (*ImportGameRequest) Descriptor() ([]byte, []int) {
	return file_v1_items_proto_rawDescGZIP(), []int{2}
}

func (x *ImportGameRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type ImportGameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ImportGameResponse) Reset() {
	*x = ImportGameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_items_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportGameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportGameResponse) ProtoMessage() {}

func (x *ImportGameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_items_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportGameResponse.ProtoReflect.Descriptor instead.
func (*ImportGameResponse) Descriptor() ([]byte, []int) {
	return file_v1_items_proto_rawDescGZIP(), []int{3}
}

func (x *ImportGameResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_v1_items_proto protoreflect.FileDescriptor

var file_v1_items_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x76, 0x31, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x2e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x22, 0x2b, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x10, 0x4b,
	0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x10, 0x01,
	0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x29, 0x0a, 0x11, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x24, 0x0a,
	0x12, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x32, 0xac, 0x01, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x45, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01,
	0x2a, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x47, 0x65, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x5c, 0x0a, 0x0a, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x47, 0x61,
	0x6d, 0x65, 0x12, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x47, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x76,
	0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x47, 0x61,
	0x6d, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x68, 0x65, 0x6c, 0x6f, 0x6a, 0x61, 0x72, 0x61, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x65, 0x6e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_items_proto_rawDescOnce sync.Once
	file_v1_items_proto_rawDescData = file_v1_items_proto_rawDesc
)

func file_v1_items_proto_rawDescGZIP() []byte {
	file_v1_items_proto_rawDescOnce.Do(func() {
		file_v1_items_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_items_proto_rawDescData)
	})
	return file_v1_items_proto_rawDescData
}

var file_v1_items_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_items_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_items_proto_goTypes = []any{
	(Item_Kind)(0),             // 0: v1.Item.Kind
	(*Item)(nil),               // 1: v1.Item
	(*GetItemRequest)(nil),     // 2: v1.GetItemRequest
	(*ImportGameRequest)(nil),  // 3: v1.ImportGameRequest
	(*ImportGameResponse)(nil), // 4: v1.ImportGameResponse
}
var file_v1_items_proto_depIdxs = []int32{
	0, // 0: v1.Item.kind:type_name -> v1.Item.Kind
	2, // 1: v1.Items.GetItem:input_type -> v1.GetItemRequest
	3, // 2: v1.Items.ImportGame:input_type -> v1.ImportGameRequest
	1, // 3: v1.Items.GetItem:output_type -> v1.Item
	4, // 4: v1.Items.ImportGame:output_type -> v1.ImportGameResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_items_proto_init() }
func file_v1_items_proto_init() {
	if File_v1_items_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_items_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Item); i {
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
		file_v1_items_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetItemRequest); i {
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
		file_v1_items_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ImportGameRequest); i {
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
		file_v1_items_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ImportGameResponse); i {
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
			RawDescriptor: file_v1_items_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_items_proto_goTypes,
		DependencyIndexes: file_v1_items_proto_depIdxs,
		EnumInfos:         file_v1_items_proto_enumTypes,
		MessageInfos:      file_v1_items_proto_msgTypes,
	}.Build()
	File_v1_items_proto = out.File
	file_v1_items_proto_rawDesc = nil
	file_v1_items_proto_goTypes = nil
	file_v1_items_proto_depIdxs = nil
}
