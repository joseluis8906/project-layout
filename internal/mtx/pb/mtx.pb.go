// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: mtx/mtx.proto

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

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNumber string                 `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Owner       *RegisterRequest_Owner `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mtx_mtx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mtx_mtx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_mtx_mtx_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *RegisterRequest) GetOwner() *RegisterRequest_Owner {
	if x != nil {
		return x.Owner
	}
	return nil
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mtx_mtx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mtx_mtx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_mtx_mtx_proto_rawDescGZIP(), []int{1}
}

type PutMoneyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PutMoneyRequest) Reset() {
	*x = PutMoneyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mtx_mtx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutMoneyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutMoneyRequest) ProtoMessage() {}

func (x *PutMoneyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mtx_mtx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutMoneyRequest.ProtoReflect.Descriptor instead.
func (*PutMoneyRequest) Descriptor() ([]byte, []int) {
	return file_mtx_mtx_proto_rawDescGZIP(), []int{2}
}

type PutMoneyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PutMoneyResponse) Reset() {
	*x = PutMoneyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mtx_mtx_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutMoneyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutMoneyResponse) ProtoMessage() {}

func (x *PutMoneyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mtx_mtx_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutMoneyResponse.ProtoReflect.Descriptor instead.
func (*PutMoneyResponse) Descriptor() ([]byte, []int) {
	return file_mtx_mtx_proto_rawDescGZIP(), []int{3}
}

type SendMoneyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendMoneyRequest) Reset() {
	*x = SendMoneyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mtx_mtx_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMoneyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMoneyRequest) ProtoMessage() {}

func (x *SendMoneyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mtx_mtx_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMoneyRequest.ProtoReflect.Descriptor instead.
func (*SendMoneyRequest) Descriptor() ([]byte, []int) {
	return file_mtx_mtx_proto_rawDescGZIP(), []int{4}
}

type SendMoneyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendMoneyResponse) Reset() {
	*x = SendMoneyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mtx_mtx_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMoneyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMoneyResponse) ProtoMessage() {}

func (x *SendMoneyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mtx_mtx_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMoneyResponse.ProtoReflect.Descriptor instead.
func (*SendMoneyResponse) Descriptor() ([]byte, []int) {
	return file_mtx_mtx_proto_rawDescGZIP(), []int{5}
}

type WithdrawRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WithdrawRequest) Reset() {
	*x = WithdrawRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mtx_mtx_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawRequest) ProtoMessage() {}

func (x *WithdrawRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mtx_mtx_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawRequest.ProtoReflect.Descriptor instead.
func (*WithdrawRequest) Descriptor() ([]byte, []int) {
	return file_mtx_mtx_proto_rawDescGZIP(), []int{6}
}

type WithdrawResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WithdrawResponse) Reset() {
	*x = WithdrawResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mtx_mtx_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WithdrawResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithdrawResponse) ProtoMessage() {}

func (x *WithdrawResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mtx_mtx_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithdrawResponse.ProtoReflect.Descriptor instead.
func (*WithdrawResponse) Descriptor() ([]byte, []int) {
	return file_mtx_mtx_proto_rawDescGZIP(), []int{7}
}

type GetBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetBalanceRequest) Reset() {
	*x = GetBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mtx_mtx_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceRequest) ProtoMessage() {}

func (x *GetBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mtx_mtx_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceRequest.ProtoReflect.Descriptor instead.
func (*GetBalanceRequest) Descriptor() ([]byte, []int) {
	return file_mtx_mtx_proto_rawDescGZIP(), []int{8}
}

type GetBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetBalanceResponse) Reset() {
	*x = GetBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mtx_mtx_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceResponse) ProtoMessage() {}

func (x *GetBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mtx_mtx_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceResponse.ProtoReflect.Descriptor instead.
func (*GetBalanceResponse) Descriptor() ([]byte, []int) {
	return file_mtx_mtx_proto_rawDescGZIP(), []int{9}
}

type RegisterRequest_Owner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	FullName string `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
}

func (x *RegisterRequest_Owner) Reset() {
	*x = RegisterRequest_Owner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mtx_mtx_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest_Owner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest_Owner) ProtoMessage() {}

func (x *RegisterRequest_Owner) ProtoReflect() protoreflect.Message {
	mi := &file_mtx_mtx_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest_Owner.ProtoReflect.Descriptor instead.
func (*RegisterRequest_Owner) Descriptor() ([]byte, []int) {
	return file_mtx_mtx_proto_rawDescGZIP(), []int{0, 0}
}

func (x *RegisterRequest_Owner) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RegisterRequest_Owner) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest_Owner) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

var File_mtx_mtx_proto protoreflect.FileDescriptor

var file_mtx_mtx_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x74, 0x78, 0x2f, 0x6d, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x6d, 0x74, 0x78, 0x22, 0xb2, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x74, 0x78,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x1a, 0x4a, 0x0a,
	0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09,
	0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x0a,
	0x0f, 0x50, 0x75, 0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x12, 0x0a, 0x10, 0x50, 0x75, 0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x6f, 0x6e, 0x65,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64,
	0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x0a,
	0x0f, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x12, 0x0a, 0x10, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xb6, 0x02, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x14,
	0x2e, 0x6d, 0x74, 0x78, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x74, 0x78, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x50,
	0x75, 0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x14, 0x2e, 0x6d, 0x74, 0x78, 0x2e, 0x50, 0x75,
	0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x6d, 0x74, 0x78, 0x2e, 0x50, 0x75, 0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x6f, 0x6e, 0x65,
	0x79, 0x12, 0x15, 0x2e, 0x6d, 0x74, 0x78, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x6f, 0x6e, 0x65,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x74, 0x78, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x37, 0x0a, 0x08, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x12, 0x14, 0x2e, 0x6d,
	0x74, 0x78, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x74, 0x78, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x2e, 0x6d, 0x74, 0x78, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x6d, 0x74, 0x78, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x6f, 0x73, 0x65, 0x6c, 0x75, 0x69, 0x73, 0x38,
	0x39, 0x30, 0x36, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x6c, 0x61, 0x79, 0x6f,
	0x75, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x74, 0x78, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mtx_mtx_proto_rawDescOnce sync.Once
	file_mtx_mtx_proto_rawDescData = file_mtx_mtx_proto_rawDesc
)

func file_mtx_mtx_proto_rawDescGZIP() []byte {
	file_mtx_mtx_proto_rawDescOnce.Do(func() {
		file_mtx_mtx_proto_rawDescData = protoimpl.X.CompressGZIP(file_mtx_mtx_proto_rawDescData)
	})
	return file_mtx_mtx_proto_rawDescData
}

var file_mtx_mtx_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_mtx_mtx_proto_goTypes = []interface{}{
	(*RegisterRequest)(nil),       // 0: mtx.RegisterRequest
	(*RegisterResponse)(nil),      // 1: mtx.RegisterResponse
	(*PutMoneyRequest)(nil),       // 2: mtx.PutMoneyRequest
	(*PutMoneyResponse)(nil),      // 3: mtx.PutMoneyResponse
	(*SendMoneyRequest)(nil),      // 4: mtx.SendMoneyRequest
	(*SendMoneyResponse)(nil),     // 5: mtx.SendMoneyResponse
	(*WithdrawRequest)(nil),       // 6: mtx.WithdrawRequest
	(*WithdrawResponse)(nil),      // 7: mtx.WithdrawResponse
	(*GetBalanceRequest)(nil),     // 8: mtx.GetBalanceRequest
	(*GetBalanceResponse)(nil),    // 9: mtx.GetBalanceResponse
	(*RegisterRequest_Owner)(nil), // 10: mtx.RegisterRequest.Owner
}
var file_mtx_mtx_proto_depIdxs = []int32{
	10, // 0: mtx.RegisterRequest.owner:type_name -> mtx.RegisterRequest.Owner
	0,  // 1: mtx.AccountService.Register:input_type -> mtx.RegisterRequest
	2,  // 2: mtx.AccountService.PutMoney:input_type -> mtx.PutMoneyRequest
	4,  // 3: mtx.AccountService.SendMoney:input_type -> mtx.SendMoneyRequest
	6,  // 4: mtx.AccountService.Withdraw:input_type -> mtx.WithdrawRequest
	8,  // 5: mtx.AccountService.GetBalance:input_type -> mtx.GetBalanceRequest
	1,  // 6: mtx.AccountService.Register:output_type -> mtx.RegisterResponse
	3,  // 7: mtx.AccountService.PutMoney:output_type -> mtx.PutMoneyResponse
	5,  // 8: mtx.AccountService.SendMoney:output_type -> mtx.SendMoneyResponse
	7,  // 9: mtx.AccountService.Withdraw:output_type -> mtx.WithdrawResponse
	9,  // 10: mtx.AccountService.GetBalance:output_type -> mtx.GetBalanceResponse
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_mtx_mtx_proto_init() }
func file_mtx_mtx_proto_init() {
	if File_mtx_mtx_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mtx_mtx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_mtx_mtx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_mtx_mtx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutMoneyRequest); i {
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
		file_mtx_mtx_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutMoneyResponse); i {
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
		file_mtx_mtx_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMoneyRequest); i {
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
		file_mtx_mtx_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMoneyResponse); i {
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
		file_mtx_mtx_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithdrawRequest); i {
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
		file_mtx_mtx_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WithdrawResponse); i {
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
		file_mtx_mtx_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceRequest); i {
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
		file_mtx_mtx_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceResponse); i {
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
		file_mtx_mtx_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest_Owner); i {
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
			RawDescriptor: file_mtx_mtx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mtx_mtx_proto_goTypes,
		DependencyIndexes: file_mtx_mtx_proto_depIdxs,
		MessageInfos:      file_mtx_mtx_proto_msgTypes,
	}.Build()
	File_mtx_mtx_proto = out.File
	file_mtx_mtx_proto_rawDesc = nil
	file_mtx_mtx_proto_goTypes = nil
	file_mtx_mtx_proto_depIdxs = nil
}
