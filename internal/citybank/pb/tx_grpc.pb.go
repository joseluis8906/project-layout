// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.0
// source: citybank/tx.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TxService_Transfer_FullMethodName      = "/citybank.TxService/Transfer"
	TxService_CheckStatus_FullMethodName   = "/citybank.TxService/CheckStatus"
	TxService_Withdraw_FullMethodName      = "/citybank.TxService/Withdraw"
	TxService_DirectDeposit_FullMethodName = "/citybank.TxService/DirectDeposit"
)

// TxServiceClient is the client API for TxService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TxServiceClient interface {
	Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error)
	CheckStatus(ctx context.Context, in *CheckStatusRequest, opts ...grpc.CallOption) (*CheckStatusResponse, error)
	Withdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*WithdrawResponse, error)
	DirectDeposit(ctx context.Context, in *DirectDepositRequest, opts ...grpc.CallOption) (*DirectDepositResponse, error)
}

type txServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTxServiceClient(cc grpc.ClientConnInterface) TxServiceClient {
	return &txServiceClient{cc}
}

func (c *txServiceClient) Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error) {
	out := new(TransferResponse)
	err := c.cc.Invoke(ctx, TxService_Transfer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txServiceClient) CheckStatus(ctx context.Context, in *CheckStatusRequest, opts ...grpc.CallOption) (*CheckStatusResponse, error) {
	out := new(CheckStatusResponse)
	err := c.cc.Invoke(ctx, TxService_CheckStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txServiceClient) Withdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*WithdrawResponse, error) {
	out := new(WithdrawResponse)
	err := c.cc.Invoke(ctx, TxService_Withdraw_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txServiceClient) DirectDeposit(ctx context.Context, in *DirectDepositRequest, opts ...grpc.CallOption) (*DirectDepositResponse, error) {
	out := new(DirectDepositResponse)
	err := c.cc.Invoke(ctx, TxService_DirectDeposit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TxServiceServer is the server API for TxService service.
// All implementations must embed UnimplementedTxServiceServer
// for forward compatibility
type TxServiceServer interface {
	Transfer(context.Context, *TransferRequest) (*TransferResponse, error)
	CheckStatus(context.Context, *CheckStatusRequest) (*CheckStatusResponse, error)
	Withdraw(context.Context, *WithdrawRequest) (*WithdrawResponse, error)
	DirectDeposit(context.Context, *DirectDepositRequest) (*DirectDepositResponse, error)
	mustEmbedUnimplementedTxServiceServer()
}

// UnimplementedTxServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTxServiceServer struct {
}

func (UnimplementedTxServiceServer) Transfer(context.Context, *TransferRequest) (*TransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}
func (UnimplementedTxServiceServer) CheckStatus(context.Context, *CheckStatusRequest) (*CheckStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckStatus not implemented")
}
func (UnimplementedTxServiceServer) Withdraw(context.Context, *WithdrawRequest) (*WithdrawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdraw not implemented")
}
func (UnimplementedTxServiceServer) DirectDeposit(context.Context, *DirectDepositRequest) (*DirectDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DirectDeposit not implemented")
}
func (UnimplementedTxServiceServer) mustEmbedUnimplementedTxServiceServer() {}

// UnsafeTxServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TxServiceServer will
// result in compilation errors.
type UnsafeTxServiceServer interface {
	mustEmbedUnimplementedTxServiceServer()
}

func RegisterTxServiceServer(s grpc.ServiceRegistrar, srv TxServiceServer) {
	s.RegisterService(&TxService_ServiceDesc, srv)
}

func _TxService_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxServiceServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TxService_Transfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxServiceServer).Transfer(ctx, req.(*TransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TxService_CheckStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxServiceServer).CheckStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TxService_CheckStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxServiceServer).CheckStatus(ctx, req.(*CheckStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TxService_Withdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxServiceServer).Withdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TxService_Withdraw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxServiceServer).Withdraw(ctx, req.(*WithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TxService_DirectDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DirectDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxServiceServer).DirectDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TxService_DirectDeposit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxServiceServer).DirectDeposit(ctx, req.(*DirectDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TxService_ServiceDesc is the grpc.ServiceDesc for TxService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TxService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "citybank.TxService",
	HandlerType: (*TxServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Transfer",
			Handler:    _TxService_Transfer_Handler,
		},
		{
			MethodName: "CheckStatus",
			Handler:    _TxService_CheckStatus_Handler,
		},
		{
			MethodName: "Withdraw",
			Handler:    _TxService_Withdraw_Handler,
		},
		{
			MethodName: "DirectDeposit",
			Handler:    _TxService_DirectDeposit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "citybank/tx.proto",
}
