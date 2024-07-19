// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: itau/hello.proto

package pb

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

type HelloWorldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *HelloWorldRequest) Reset() {
	*x = HelloWorldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itau_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloWorldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloWorldRequest) ProtoMessage() {}

func (x *HelloWorldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_itau_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloWorldRequest.ProtoReflect.Descriptor instead.
func (*HelloWorldRequest) Descriptor() ([]byte, []int) {
	return file_itau_hello_proto_rawDescGZIP(), []int{0}
}

func (x *HelloWorldRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type HelloWorldResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *HelloWorldResponse) Reset() {
	*x = HelloWorldResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itau_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloWorldResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloWorldResponse) ProtoMessage() {}

func (x *HelloWorldResponse) ProtoReflect() protoreflect.Message {
	mi := &file_itau_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloWorldResponse.ProtoReflect.Descriptor instead.
func (*HelloWorldResponse) Descriptor() ([]byte, []int) {
	return file_itau_hello_proto_rawDescGZIP(), []int{1}
}

func (x *HelloWorldResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_itau_hello_proto protoreflect.FileDescriptor

var file_itau_hello_proto_rawDesc = []byte{
	0x0a, 0x10, 0x69, 0x74, 0x61, 0x75, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x69, 0x74, 0x61, 0x75, 0x22, 0x25, 0x0a, 0x11, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22,
	0x26, 0x0a, 0x12, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x4a, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x57, 0x6f, 0x72, 0x6c, 0x64,
	0x12, 0x17, 0x2e, 0x69, 0x74, 0x61, 0x75, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72,
	0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x69, 0x74, 0x61, 0x75,
	0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6a, 0x6f, 0x73, 0x65, 0x6c, 0x75, 0x69, 0x73, 0x38, 0x39, 0x30, 0x36, 0x2f, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x74, 0x61, 0x75, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_itau_hello_proto_rawDescOnce sync.Once
	file_itau_hello_proto_rawDescData = file_itau_hello_proto_rawDesc
)

func file_itau_hello_proto_rawDescGZIP() []byte {
	file_itau_hello_proto_rawDescOnce.Do(func() {
		file_itau_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_itau_hello_proto_rawDescData)
	})
	return file_itau_hello_proto_rawDescData
}

var file_itau_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_itau_hello_proto_goTypes = []interface{}{
	(*HelloWorldRequest)(nil),  // 0: itau.HelloWorldRequest
	(*HelloWorldResponse)(nil), // 1: itau.HelloWorldResponse
}
var file_itau_hello_proto_depIdxs = []int32{
	0, // 0: itau.HelloService.World:input_type -> itau.HelloWorldRequest
	1, // 1: itau.HelloService.World:output_type -> itau.HelloWorldResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_itau_hello_proto_init() }
func file_itau_hello_proto_init() {
	if File_itau_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_itau_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloWorldRequest); i {
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
		file_itau_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloWorldResponse); i {
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
			RawDescriptor: file_itau_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_itau_hello_proto_goTypes,
		DependencyIndexes: file_itau_hello_proto_depIdxs,
		MessageInfos:      file_itau_hello_proto_msgTypes,
	}.Build()
	File_itau_hello_proto = out.File
	file_itau_hello_proto_rawDesc = nil
	file_itau_hello_proto_goTypes = nil
	file_itau_hello_proto_depIdxs = nil
}
