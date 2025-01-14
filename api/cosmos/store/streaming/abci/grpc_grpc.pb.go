// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cosmos/store/streaming/abci/grpc.proto

package abci

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
	ABCIListenerService_ListenBeginBlock_FullMethodName = "/cosmos.store.streaming.abci.ABCIListenerService/ListenBeginBlock"
	ABCIListenerService_ListenEndBlock_FullMethodName   = "/cosmos.store.streaming.abci.ABCIListenerService/ListenEndBlock"
	ABCIListenerService_ListenDeliverTx_FullMethodName  = "/cosmos.store.streaming.abci.ABCIListenerService/ListenDeliverTx"
	ABCIListenerService_ListenCommit_FullMethodName     = "/cosmos.store.streaming.abci.ABCIListenerService/ListenCommit"
)

// ABCIListenerServiceClient is the client API for ABCIListenerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ABCIListenerServiceClient interface {
	// ListenBeginBlock is the corresponding endpoint for ABCIListener.ListenBeginBlock
	ListenBeginBlock(ctx context.Context, in *ListenBeginBlockRequest, opts ...grpc.CallOption) (*ListenBeginBlockResponse, error)
	// ListenEndBlock is the corresponding endpoint for ABCIListener.ListenEndBlock
	ListenEndBlock(ctx context.Context, in *ListenEndBlockRequest, opts ...grpc.CallOption) (*ListenEndBlockResponse, error)
	// ListenDeliverTx is the corresponding endpoint for ABCIListener.ListenDeliverTx
	ListenDeliverTx(ctx context.Context, in *ListenDeliverTxRequest, opts ...grpc.CallOption) (*ListenDeliverTxResponse, error)
	// ListenCommit is the corresponding endpoint for ABCIListener.ListenCommit
	ListenCommit(ctx context.Context, in *ListenCommitRequest, opts ...grpc.CallOption) (*ListenCommitResponse, error)
}

type aBCIListenerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewABCIListenerServiceClient(cc grpc.ClientConnInterface) ABCIListenerServiceClient {
	return &aBCIListenerServiceClient{cc}
}

func (c *aBCIListenerServiceClient) ListenBeginBlock(ctx context.Context, in *ListenBeginBlockRequest, opts ...grpc.CallOption) (*ListenBeginBlockResponse, error) {
	out := new(ListenBeginBlockResponse)
	err := c.cc.Invoke(ctx, ABCIListenerService_ListenBeginBlock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIListenerServiceClient) ListenEndBlock(ctx context.Context, in *ListenEndBlockRequest, opts ...grpc.CallOption) (*ListenEndBlockResponse, error) {
	out := new(ListenEndBlockResponse)
	err := c.cc.Invoke(ctx, ABCIListenerService_ListenEndBlock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIListenerServiceClient) ListenDeliverTx(ctx context.Context, in *ListenDeliverTxRequest, opts ...grpc.CallOption) (*ListenDeliverTxResponse, error) {
	out := new(ListenDeliverTxResponse)
	err := c.cc.Invoke(ctx, ABCIListenerService_ListenDeliverTx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBCIListenerServiceClient) ListenCommit(ctx context.Context, in *ListenCommitRequest, opts ...grpc.CallOption) (*ListenCommitResponse, error) {
	out := new(ListenCommitResponse)
	err := c.cc.Invoke(ctx, ABCIListenerService_ListenCommit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ABCIListenerServiceServer is the server API for ABCIListenerService service.
// All implementations must embed UnimplementedABCIListenerServiceServer
// for forward compatibility
type ABCIListenerServiceServer interface {
	// ListenBeginBlock is the corresponding endpoint for ABCIListener.ListenBeginBlock
	ListenBeginBlock(context.Context, *ListenBeginBlockRequest) (*ListenBeginBlockResponse, error)
	// ListenEndBlock is the corresponding endpoint for ABCIListener.ListenEndBlock
	ListenEndBlock(context.Context, *ListenEndBlockRequest) (*ListenEndBlockResponse, error)
	// ListenDeliverTx is the corresponding endpoint for ABCIListener.ListenDeliverTx
	ListenDeliverTx(context.Context, *ListenDeliverTxRequest) (*ListenDeliverTxResponse, error)
	// ListenCommit is the corresponding endpoint for ABCIListener.ListenCommit
	ListenCommit(context.Context, *ListenCommitRequest) (*ListenCommitResponse, error)
	mustEmbedUnimplementedABCIListenerServiceServer()
}

// UnimplementedABCIListenerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedABCIListenerServiceServer struct {
}

func (UnimplementedABCIListenerServiceServer) ListenBeginBlock(context.Context, *ListenBeginBlockRequest) (*ListenBeginBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListenBeginBlock not implemented")
}
func (UnimplementedABCIListenerServiceServer) ListenEndBlock(context.Context, *ListenEndBlockRequest) (*ListenEndBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListenEndBlock not implemented")
}
func (UnimplementedABCIListenerServiceServer) ListenDeliverTx(context.Context, *ListenDeliverTxRequest) (*ListenDeliverTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListenDeliverTx not implemented")
}
func (UnimplementedABCIListenerServiceServer) ListenCommit(context.Context, *ListenCommitRequest) (*ListenCommitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListenCommit not implemented")
}
func (UnimplementedABCIListenerServiceServer) mustEmbedUnimplementedABCIListenerServiceServer() {}

// UnsafeABCIListenerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ABCIListenerServiceServer will
// result in compilation errors.
type UnsafeABCIListenerServiceServer interface {
	mustEmbedUnimplementedABCIListenerServiceServer()
}

func RegisterABCIListenerServiceServer(s grpc.ServiceRegistrar, srv ABCIListenerServiceServer) {
	s.RegisterService(&ABCIListenerService_ServiceDesc, srv)
}

func _ABCIListenerService_ListenBeginBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListenBeginBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIListenerServiceServer).ListenBeginBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ABCIListenerService_ListenBeginBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIListenerServiceServer).ListenBeginBlock(ctx, req.(*ListenBeginBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIListenerService_ListenEndBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListenEndBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIListenerServiceServer).ListenEndBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ABCIListenerService_ListenEndBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIListenerServiceServer).ListenEndBlock(ctx, req.(*ListenEndBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIListenerService_ListenDeliverTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListenDeliverTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIListenerServiceServer).ListenDeliverTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ABCIListenerService_ListenDeliverTx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIListenerServiceServer).ListenDeliverTx(ctx, req.(*ListenDeliverTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABCIListenerService_ListenCommit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListenCommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABCIListenerServiceServer).ListenCommit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ABCIListenerService_ListenCommit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABCIListenerServiceServer).ListenCommit(ctx, req.(*ListenCommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ABCIListenerService_ServiceDesc is the grpc.ServiceDesc for ABCIListenerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ABCIListenerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.store.streaming.abci.ABCIListenerService",
	HandlerType: (*ABCIListenerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListenBeginBlock",
			Handler:    _ABCIListenerService_ListenBeginBlock_Handler,
		},
		{
			MethodName: "ListenEndBlock",
			Handler:    _ABCIListenerService_ListenEndBlock_Handler,
		},
		{
			MethodName: "ListenDeliverTx",
			Handler:    _ABCIListenerService_ListenDeliverTx_Handler,
		},
		{
			MethodName: "ListenCommit",
			Handler:    _ABCIListenerService_ListenCommit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/store/streaming/abci/grpc.proto",
}
