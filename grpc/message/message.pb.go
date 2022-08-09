// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: grpc/proto/message.proto

package message

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value1 int32 `protobuf:"varint,1,opt,name=Value1,proto3" json:"Value1,omitempty"`
	Value2 int32 `protobuf:"varint,2,opt,name=Value2,proto3" json:"Value2,omitempty"`
	Value3 int32 `protobuf:"varint,3,opt,name=Value3,proto3" json:"Value3,omitempty"`
	Value4 int32 `protobuf:"varint,4,opt,name=Value4,proto3" json:"Value4,omitempty"`
	Value5 int32 `protobuf:"varint,5,opt,name=Value5,proto3" json:"Value5,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_grpc_proto_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetValue1() int32 {
	if x != nil {
		return x.Value1
	}
	return 0
}

func (x *Message) GetValue2() int32 {
	if x != nil {
		return x.Value2
	}
	return 0
}

func (x *Message) GetValue3() int32 {
	if x != nil {
		return x.Value3
	}
	return 0
}

func (x *Message) GetValue4() int32 {
	if x != nil {
		return x.Value4
	}
	return 0
}

func (x *Message) GetValue5() int32 {
	if x != nil {
		return x.Value5
	}
	return 0
}

var File_grpc_proto_message_proto protoreflect.FileDescriptor

var file_grpc_proto_message_proto_rawDesc = []byte{
	0x0a, 0x18, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x31,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x31, 0x12, 0x16,
	0x0a, 0x06, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x33,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x33, 0x12, 0x16,
	0x0a, 0x06, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x34, 0x12, 0x16, 0x0a, 0x06, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x35,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x35, 0x32, 0x35,
	0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x23, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x08, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x08, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_message_proto_rawDescOnce sync.Once
	file_grpc_proto_message_proto_rawDescData = file_grpc_proto_message_proto_rawDesc
)

func file_grpc_proto_message_proto_rawDescGZIP() []byte {
	file_grpc_proto_message_proto_rawDescOnce.Do(func() {
		file_grpc_proto_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_message_proto_rawDescData)
	})
	return file_grpc_proto_message_proto_rawDescData
}

var file_grpc_proto_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_grpc_proto_message_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: Message
}
var file_grpc_proto_message_proto_depIdxs = []int32{
	0, // 0: MessageService.SendMessage:input_type -> Message
	0, // 1: MessageService.SendMessage:output_type -> Message
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_proto_message_proto_init() }
func file_grpc_proto_message_proto_init() {
	if File_grpc_proto_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_grpc_proto_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_message_proto_goTypes,
		DependencyIndexes: file_grpc_proto_message_proto_depIdxs,
		MessageInfos:      file_grpc_proto_message_proto_msgTypes,
	}.Build()
	File_grpc_proto_message_proto = out.File
	file_grpc_proto_message_proto_rawDesc = nil
	file_grpc_proto_message_proto_goTypes = nil
	file_grpc_proto_message_proto_depIdxs = nil
}