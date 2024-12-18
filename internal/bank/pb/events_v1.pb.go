// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: bank/events_v1.proto

package pb

import (
	pb "github.com/joseluis8906/project-layout/pkg/pb"
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

type Events_V1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Events_V1) Reset() {
	*x = Events_V1{}
	mi := &file_bank_events_v1_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Events_V1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Events_V1) ProtoMessage() {}

func (x *Events_V1) ProtoReflect() protoreflect.Message {
	mi := &file_bank_events_v1_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Events_V1.ProtoReflect.Descriptor instead.
func (*Events_V1) Descriptor() ([]byte, []int) {
	return file_bank_events_v1_proto_rawDescGZIP(), []int{0}
}

type Events_V1_AccountCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OccurredOn int64                                `protobuf:"varint,2,opt,name=occurred_on,json=occurredOn,proto3" json:"occurred_on,omitempty"`
	Attributes *Events_V1_AccountCreated_Attributes `protobuf:"bytes,3,opt,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *Events_V1_AccountCreated) Reset() {
	*x = Events_V1_AccountCreated{}
	mi := &file_bank_events_v1_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Events_V1_AccountCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Events_V1_AccountCreated) ProtoMessage() {}

func (x *Events_V1_AccountCreated) ProtoReflect() protoreflect.Message {
	mi := &file_bank_events_v1_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Events_V1_AccountCreated.ProtoReflect.Descriptor instead.
func (*Events_V1_AccountCreated) Descriptor() ([]byte, []int) {
	return file_bank_events_v1_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Events_V1_AccountCreated) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Events_V1_AccountCreated) GetOccurredOn() int64 {
	if x != nil {
		return x.OccurredOn
	}
	return 0
}

func (x *Events_V1_AccountCreated) GetAttributes() *Events_V1_AccountCreated_Attributes {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type Events_V1_AccountDebited struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OccurredOn int64                                `protobuf:"varint,2,opt,name=occurred_on,json=occurredOn,proto3" json:"occurred_on,omitempty"`
	Attributes *Events_V1_AccountDebited_Attributes `protobuf:"bytes,3,opt,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *Events_V1_AccountDebited) Reset() {
	*x = Events_V1_AccountDebited{}
	mi := &file_bank_events_v1_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Events_V1_AccountDebited) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Events_V1_AccountDebited) ProtoMessage() {}

func (x *Events_V1_AccountDebited) ProtoReflect() protoreflect.Message {
	mi := &file_bank_events_v1_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Events_V1_AccountDebited.ProtoReflect.Descriptor instead.
func (*Events_V1_AccountDebited) Descriptor() ([]byte, []int) {
	return file_bank_events_v1_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Events_V1_AccountDebited) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Events_V1_AccountDebited) GetOccurredOn() int64 {
	if x != nil {
		return x.OccurredOn
	}
	return 0
}

func (x *Events_V1_AccountDebited) GetAttributes() *Events_V1_AccountDebited_Attributes {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type Events_V1_AccountCredited struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OccurredOn int64                                 `protobuf:"varint,2,opt,name=occurred_on,json=occurredOn,proto3" json:"occurred_on,omitempty"`
	Attributes *Events_V1_AccountCredited_Attributes `protobuf:"bytes,3,opt,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *Events_V1_AccountCredited) Reset() {
	*x = Events_V1_AccountCredited{}
	mi := &file_bank_events_v1_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Events_V1_AccountCredited) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Events_V1_AccountCredited) ProtoMessage() {}

func (x *Events_V1_AccountCredited) ProtoReflect() protoreflect.Message {
	mi := &file_bank_events_v1_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Events_V1_AccountCredited.ProtoReflect.Descriptor instead.
func (*Events_V1_AccountCredited) Descriptor() ([]byte, []int) {
	return file_bank_events_v1_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Events_V1_AccountCredited) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Events_V1_AccountCredited) GetOccurredOn() int64 {
	if x != nil {
		return x.OccurredOn
	}
	return 0
}

func (x *Events_V1_AccountCredited) GetAttributes() *Events_V1_AccountCredited_Attributes {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type Events_V1_AccountCreated_Attributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Number string `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *Events_V1_AccountCreated_Attributes) Reset() {
	*x = Events_V1_AccountCreated_Attributes{}
	mi := &file_bank_events_v1_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Events_V1_AccountCreated_Attributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Events_V1_AccountCreated_Attributes) ProtoMessage() {}

func (x *Events_V1_AccountCreated_Attributes) ProtoReflect() protoreflect.Message {
	mi := &file_bank_events_v1_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Events_V1_AccountCreated_Attributes.ProtoReflect.Descriptor instead.
func (*Events_V1_AccountCreated_Attributes) Descriptor() ([]byte, []int) {
	return file_bank_events_v1_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Events_V1_AccountCreated_Attributes) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Events_V1_AccountCreated_Attributes) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

type Events_V1_AccountDebited_Attributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   string    `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Number string    `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	Amount *pb.Money `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Events_V1_AccountDebited_Attributes) Reset() {
	*x = Events_V1_AccountDebited_Attributes{}
	mi := &file_bank_events_v1_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Events_V1_AccountDebited_Attributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Events_V1_AccountDebited_Attributes) ProtoMessage() {}

func (x *Events_V1_AccountDebited_Attributes) ProtoReflect() protoreflect.Message {
	mi := &file_bank_events_v1_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Events_V1_AccountDebited_Attributes.ProtoReflect.Descriptor instead.
func (*Events_V1_AccountDebited_Attributes) Descriptor() ([]byte, []int) {
	return file_bank_events_v1_proto_rawDescGZIP(), []int{0, 1, 0}
}

func (x *Events_V1_AccountDebited_Attributes) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Events_V1_AccountDebited_Attributes) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Events_V1_AccountDebited_Attributes) GetAmount() *pb.Money {
	if x != nil {
		return x.Amount
	}
	return nil
}

type Events_V1_AccountCredited_Attributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   string    `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Number string    `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	Amount *pb.Money `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Events_V1_AccountCredited_Attributes) Reset() {
	*x = Events_V1_AccountCredited_Attributes{}
	mi := &file_bank_events_v1_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Events_V1_AccountCredited_Attributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Events_V1_AccountCredited_Attributes) ProtoMessage() {}

func (x *Events_V1_AccountCredited_Attributes) ProtoReflect() protoreflect.Message {
	mi := &file_bank_events_v1_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Events_V1_AccountCredited_Attributes.ProtoReflect.Descriptor instead.
func (*Events_V1_AccountCredited_Attributes) Descriptor() ([]byte, []int) {
	return file_bank_events_v1_proto_rawDescGZIP(), []int{0, 2, 0}
}

func (x *Events_V1_AccountCredited_Attributes) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Events_V1_AccountCredited_Attributes) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Events_V1_AccountCredited_Attributes) GetAmount() *pb.Money {
	if x != nil {
		return x.Amount
	}
	return nil
}

var File_bank_events_v1_proto protoreflect.FileDescriptor

var file_bank_events_v1_proto_rawDesc = []byte{
	0x0a, 0x14, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x76, 0x31,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x1a, 0x0d, 0x66, 0x69,
	0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x05, 0x0a, 0x09,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x56, 0x31, 0x1a, 0xc6, 0x01, 0x0a, 0x0e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x49, 0x0a,
	0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5f,
	0x56, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x0a, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x1a, 0x38, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x1a, 0xee, 0x01, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65,
	0x62, 0x69, 0x74, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6f, 0x63, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x49, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x62, 0x61, 0x6e,
	0x6b, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x56, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x44, 0x65, 0x62, 0x69, 0x74, 0x65, 0x64, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x1a, 0x60, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x1a, 0xf0, 0x01, 0x0a, 0x0f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x63, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6f, 0x63,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x4a, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x62,
	0x61, 0x6e, 0x6b, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x56, 0x31, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x65, 0x64, 0x2e, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x1a, 0x60, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x6f, 0x73, 0x65, 0x6c, 0x75, 0x69, 0x73, 0x38, 0x39, 0x30,
	0x36, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bank_events_v1_proto_rawDescOnce sync.Once
	file_bank_events_v1_proto_rawDescData = file_bank_events_v1_proto_rawDesc
)

func file_bank_events_v1_proto_rawDescGZIP() []byte {
	file_bank_events_v1_proto_rawDescOnce.Do(func() {
		file_bank_events_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_bank_events_v1_proto_rawDescData)
	})
	return file_bank_events_v1_proto_rawDescData
}

var file_bank_events_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_bank_events_v1_proto_goTypes = []any{
	(*Events_V1)(nil),                            // 0: bank.Events_V1
	(*Events_V1_AccountCreated)(nil),             // 1: bank.Events_V1.AccountCreated
	(*Events_V1_AccountDebited)(nil),             // 2: bank.Events_V1.AccountDebited
	(*Events_V1_AccountCredited)(nil),            // 3: bank.Events_V1.AccountCredited
	(*Events_V1_AccountCreated_Attributes)(nil),  // 4: bank.Events_V1.AccountCreated.Attributes
	(*Events_V1_AccountDebited_Attributes)(nil),  // 5: bank.Events_V1.AccountDebited.Attributes
	(*Events_V1_AccountCredited_Attributes)(nil), // 6: bank.Events_V1.AccountCredited.Attributes
	(*pb.Money)(nil),                             // 7: company.Money
}
var file_bank_events_v1_proto_depIdxs = []int32{
	4, // 0: bank.Events_V1.AccountCreated.attributes:type_name -> bank.Events_V1.AccountCreated.Attributes
	5, // 1: bank.Events_V1.AccountDebited.attributes:type_name -> bank.Events_V1.AccountDebited.Attributes
	6, // 2: bank.Events_V1.AccountCredited.attributes:type_name -> bank.Events_V1.AccountCredited.Attributes
	7, // 3: bank.Events_V1.AccountDebited.Attributes.amount:type_name -> company.Money
	7, // 4: bank.Events_V1.AccountCredited.Attributes.amount:type_name -> company.Money
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_bank_events_v1_proto_init() }
func file_bank_events_v1_proto_init() {
	if File_bank_events_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bank_events_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bank_events_v1_proto_goTypes,
		DependencyIndexes: file_bank_events_v1_proto_depIdxs,
		MessageInfos:      file_bank_events_v1_proto_msgTypes,
	}.Build()
	File_bank_events_v1_proto = out.File
	file_bank_events_v1_proto_rawDesc = nil
	file_bank_events_v1_proto_goTypes = nil
	file_bank_events_v1_proto_depIdxs = nil
}
