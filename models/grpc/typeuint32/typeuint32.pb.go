// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: models/grpc/proto/typeuint32.proto

package typeuint32

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uint32 uint32 `protobuf:"varint,1,opt,name=uint32,proto3" json:"uint32,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_grpc_proto_typeuint32_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_models_grpc_proto_typeuint32_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_models_grpc_proto_typeuint32_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetUint32() uint32 {
	if x != nil {
		return x.Uint32
	}
	return 0
}

type EmptyResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResult) Reset() {
	*x = EmptyResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_grpc_proto_typeuint32_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResult) ProtoMessage() {}

func (x *EmptyResult) ProtoReflect() protoreflect.Message {
	mi := &file_models_grpc_proto_typeuint32_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResult.ProtoReflect.Descriptor instead.
func (*EmptyResult) Descriptor() ([]byte, []int) {
	return file_models_grpc_proto_typeuint32_proto_rawDescGZIP(), []int{1}
}

var File_models_grpc_proto_typeuint32_proto protoreflect.FileDescriptor

var file_models_grpc_proto_typeuint32_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x22, 0x0d, 0x0a, 0x0b, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x2e, 0x0a, 0x07, 0x54, 0x79, 0x70, 0x65, 0x41, 0x6c,
	0x6c, 0x12, 0x23, 0x0a, 0x07, 0x54, 0x79, 0x70, 0x65, 0x41, 0x6c, 0x6c, 0x12, 0x08, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x6d, 0x65, 0x69, 0x64, 0x61, 0x2d, 0x72, 0x61, 0x70,
	0x68, 0x61, 0x65, 0x6c, 0x2f, 0x61, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_grpc_proto_typeuint32_proto_rawDescOnce sync.Once
	file_models_grpc_proto_typeuint32_proto_rawDescData = file_models_grpc_proto_typeuint32_proto_rawDesc
)

func file_models_grpc_proto_typeuint32_proto_rawDescGZIP() []byte {
	file_models_grpc_proto_typeuint32_proto_rawDescOnce.Do(func() {
		file_models_grpc_proto_typeuint32_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_grpc_proto_typeuint32_proto_rawDescData)
	})
	return file_models_grpc_proto_typeuint32_proto_rawDescData
}

var file_models_grpc_proto_typeuint32_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_models_grpc_proto_typeuint32_proto_goTypes = []interface{}{
	(*Request)(nil),     // 0: Request
	(*EmptyResult)(nil), // 1: EmptyResult
}
var file_models_grpc_proto_typeuint32_proto_depIdxs = []int32{
	0, // 0: TypeAll.TypeAll:input_type -> Request
	1, // 1: TypeAll.TypeAll:output_type -> EmptyResult
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_grpc_proto_typeuint32_proto_init() }
func file_models_grpc_proto_typeuint32_proto_init() {
	if File_models_grpc_proto_typeuint32_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_grpc_proto_typeuint32_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_models_grpc_proto_typeuint32_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResult); i {
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
			RawDescriptor: file_models_grpc_proto_typeuint32_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_models_grpc_proto_typeuint32_proto_goTypes,
		DependencyIndexes: file_models_grpc_proto_typeuint32_proto_depIdxs,
		MessageInfos:      file_models_grpc_proto_typeuint32_proto_msgTypes,
	}.Build()
	File_models_grpc_proto_typeuint32_proto = out.File
	file_models_grpc_proto_typeuint32_proto_rawDesc = nil
	file_models_grpc_proto_typeuint32_proto_goTypes = nil
	file_models_grpc_proto_typeuint32_proto_depIdxs = nil
}
