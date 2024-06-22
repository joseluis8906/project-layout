// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: banking/account.proto

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

type CreateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     string                      `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Currency string                      `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
	Owner    *CreateAccountRequest_Owner `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banking_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_banking_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_banking_account_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAccountRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateAccountRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *CreateAccountRequest) GetOwner() *CreateAccountRequest_Owner {
	if x != nil {
		return x.Owner
	}
	return nil
}

type CreateAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number string `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *CreateAccountResponse) Reset() {
	*x = CreateAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banking_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountResponse) ProtoMessage() {}

func (x *CreateAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_banking_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountResponse.ProtoReflect.Descriptor instead.
func (*CreateAccountResponse) Descriptor() ([]byte, []int) {
	return file_banking_account_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAccountResponse) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

type CreditAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   string    `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Number string    `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	Amount *pb.Money `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *CreditAccountRequest) Reset() {
	*x = CreditAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banking_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreditAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreditAccountRequest) ProtoMessage() {}

func (x *CreditAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_banking_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreditAccountRequest.ProtoReflect.Descriptor instead.
func (*CreditAccountRequest) Descriptor() ([]byte, []int) {
	return file_banking_account_proto_rawDescGZIP(), []int{2}
}

func (x *CreditAccountRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreditAccountRequest) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *CreditAccountRequest) GetAmount() *pb.Money {
	if x != nil {
		return x.Amount
	}
	return nil
}

type CreditAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreditAccountResponse) Reset() {
	*x = CreditAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banking_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreditAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreditAccountResponse) ProtoMessage() {}

func (x *CreditAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_banking_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreditAccountResponse.ProtoReflect.Descriptor instead.
func (*CreditAccountResponse) Descriptor() ([]byte, []int) {
	return file_banking_account_proto_rawDescGZIP(), []int{3}
}

type CreateAccountRequest_Owner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Country  string `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	FullName string `protobuf:"bytes,4,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
}

func (x *CreateAccountRequest_Owner) Reset() {
	*x = CreateAccountRequest_Owner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_banking_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountRequest_Owner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest_Owner) ProtoMessage() {}

func (x *CreateAccountRequest_Owner) ProtoReflect() protoreflect.Message {
	mi := &file_banking_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest_Owner.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest_Owner) Descriptor() ([]byte, []int) {
	return file_banking_account_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CreateAccountRequest_Owner) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateAccountRequest_Owner) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *CreateAccountRequest_Owner) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateAccountRequest_Owner) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

var File_banking_account_proto protoreflect.FileDescriptor

var file_banking_account_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67,
	0x1a, 0x0d, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe7, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x39, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x1a, 0x64, 0x0a, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09,
	0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x6a, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xb0, 0x01, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1d, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1d, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6a, 0x6f, 0x73, 0x65, 0x6c, 0x75, 0x69, 0x73, 0x38, 0x39, 0x30, 0x36, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_banking_account_proto_rawDescOnce sync.Once
	file_banking_account_proto_rawDescData = file_banking_account_proto_rawDesc
)

func file_banking_account_proto_rawDescGZIP() []byte {
	file_banking_account_proto_rawDescOnce.Do(func() {
		file_banking_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_banking_account_proto_rawDescData)
	})
	return file_banking_account_proto_rawDescData
}

var file_banking_account_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_banking_account_proto_goTypes = []interface{}{
	(*CreateAccountRequest)(nil),       // 0: banking.CreateAccountRequest
	(*CreateAccountResponse)(nil),      // 1: banking.CreateAccountResponse
	(*CreditAccountRequest)(nil),       // 2: banking.CreditAccountRequest
	(*CreditAccountResponse)(nil),      // 3: banking.CreditAccountResponse
	(*CreateAccountRequest_Owner)(nil), // 4: banking.CreateAccountRequest.Owner
	(*pb.Money)(nil),                   // 5: company.Money
}
var file_banking_account_proto_depIdxs = []int32{
	4, // 0: banking.CreateAccountRequest.owner:type_name -> banking.CreateAccountRequest.Owner
	5, // 1: banking.CreditAccountRequest.amount:type_name -> company.Money
	0, // 2: banking.AccountService.CreateAccount:input_type -> banking.CreateAccountRequest
	2, // 3: banking.AccountService.CreditAccount:input_type -> banking.CreditAccountRequest
	1, // 4: banking.AccountService.CreateAccount:output_type -> banking.CreateAccountResponse
	3, // 5: banking.AccountService.CreditAccount:output_type -> banking.CreditAccountResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_banking_account_proto_init() }
func file_banking_account_proto_init() {
	if File_banking_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_banking_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountRequest); i {
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
		file_banking_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountResponse); i {
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
		file_banking_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreditAccountRequest); i {
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
		file_banking_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreditAccountResponse); i {
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
		file_banking_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountRequest_Owner); i {
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
			RawDescriptor: file_banking_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_banking_account_proto_goTypes,
		DependencyIndexes: file_banking_account_proto_depIdxs,
		MessageInfos:      file_banking_account_proto_msgTypes,
	}.Build()
	File_banking_account_proto = out.File
	file_banking_account_proto_rawDesc = nil
	file_banking_account_proto_goTypes = nil
	file_banking_account_proto_depIdxs = nil
}
