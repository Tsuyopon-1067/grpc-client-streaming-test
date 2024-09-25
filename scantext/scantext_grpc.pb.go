// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: scantext/scantext.proto

package grpc_client_streaming_test

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Sender_SendText_FullMethodName = "/scantext.Sender/SendText"
)

// SenderClient is the client API for Sender service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SenderClient interface {
	SendText(ctx context.Context, in *ScanText, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type senderClient struct {
	cc grpc.ClientConnInterface
}

func NewSenderClient(cc grpc.ClientConnInterface) SenderClient {
	return &senderClient{cc}
}

func (c *senderClient) SendText(ctx context.Context, in *ScanText, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Sender_SendText_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SenderServer is the server API for Sender service.
// All implementations must embed UnimplementedSenderServer
// for forward compatibility.
type SenderServer interface {
	SendText(context.Context, *ScanText) (*emptypb.Empty, error)
	mustEmbedUnimplementedSenderServer()
}

// UnimplementedSenderServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSenderServer struct{}

func (UnimplementedSenderServer) SendText(context.Context, *ScanText) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendText not implemented")
}
func (UnimplementedSenderServer) mustEmbedUnimplementedSenderServer() {}
func (UnimplementedSenderServer) testEmbeddedByValue()                {}

// UnsafeSenderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SenderServer will
// result in compilation errors.
type UnsafeSenderServer interface {
	mustEmbedUnimplementedSenderServer()
}

func RegisterSenderServer(s grpc.ServiceRegistrar, srv SenderServer) {
	// If the following call pancis, it indicates UnimplementedSenderServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Sender_ServiceDesc, srv)
}

func _Sender_SendText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScanText)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SenderServer).SendText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sender_SendText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SenderServer).SendText(ctx, req.(*ScanText))
	}
	return interceptor(ctx, in, info, handler)
}

// Sender_ServiceDesc is the grpc.ServiceDesc for Sender service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sender_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scantext.Sender",
	HandlerType: (*SenderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendText",
			Handler:    _Sender_SendText_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scantext/scantext.proto",
}
