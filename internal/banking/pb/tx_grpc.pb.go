// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.0
// source: banking/tx.proto

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
	TxService_Transfer_FullMethodName      = "/banking.TxService/Transfer"
	TxService_CheckTxStatus_FullMethodName = "/banking.TxService/CheckTxStatus"
)

// TxServiceClient is the client API for TxService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TxServiceClient interface {
	Transfer(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransferResponse, error)
	CheckTxStatus(ctx context.Context, in *CheckTxStatusRequest, opts ...grpc.CallOption) (*CheckTxStatusResponse, error)
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

func (c *txServiceClient) CheckTxStatus(ctx context.Context, in *CheckTxStatusRequest, opts ...grpc.CallOption) (*CheckTxStatusResponse, error) {
	out := new(CheckTxStatusResponse)
	err := c.cc.Invoke(ctx, TxService_CheckTxStatus_FullMethodName, in, out, opts...)
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
	CheckTxStatus(context.Context, *CheckTxStatusRequest) (*CheckTxStatusResponse, error)
	mustEmbedUnimplementedTxServiceServer()
}

// UnimplementedTxServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTxServiceServer struct {
}

func (UnimplementedTxServiceServer) Transfer(context.Context, *TransferRequest) (*TransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}
func (UnimplementedTxServiceServer) CheckTxStatus(context.Context, *CheckTxStatusRequest) (*CheckTxStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckTxStatus not implemented")
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

func _TxService_CheckTxStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckTxStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxServiceServer).CheckTxStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TxService_CheckTxStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxServiceServer).CheckTxStatus(ctx, req.(*CheckTxStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TxService_ServiceDesc is the grpc.ServiceDesc for TxService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TxService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "banking.TxService",
	HandlerType: (*TxServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Transfer",
			Handler:    _TxService_Transfer_Handler,
		},
		{
			MethodName: "CheckTxStatus",
			Handler:    _TxService_CheckTxStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "banking/tx.proto",
}
