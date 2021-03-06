// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package doubletype

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

// DoubleTypeClient is the client API for DoubleType service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DoubleTypeClient interface {
	Average(ctx context.Context, in *NumberList, opts ...grpc.CallOption) (*Result, error)
}

type doubleTypeClient struct {
	cc grpc.ClientConnInterface
}

func NewDoubleTypeClient(cc grpc.ClientConnInterface) DoubleTypeClient {
	return &doubleTypeClient{cc}
}

func (c *doubleTypeClient) Average(ctx context.Context, in *NumberList, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/DoubleType/Average", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DoubleTypeServer is the server API for DoubleType service.
// All implementations must embed UnimplementedDoubleTypeServer
// for forward compatibility
type DoubleTypeServer interface {
	Average(context.Context, *NumberList) (*Result, error)
	mustEmbedUnimplementedDoubleTypeServer()
}

// UnimplementedDoubleTypeServer must be embedded to have forward compatible implementations.
type UnimplementedDoubleTypeServer struct {
}

func (UnimplementedDoubleTypeServer) Average(context.Context, *NumberList) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Average not implemented")
}
func (UnimplementedDoubleTypeServer) mustEmbedUnimplementedDoubleTypeServer() {}

// UnsafeDoubleTypeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DoubleTypeServer will
// result in compilation errors.
type UnsafeDoubleTypeServer interface {
	mustEmbedUnimplementedDoubleTypeServer()
}

func RegisterDoubleTypeServer(s grpc.ServiceRegistrar, srv DoubleTypeServer) {
	s.RegisterService(&DoubleType_ServiceDesc, srv)
}

func _DoubleType_Average_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NumberList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DoubleTypeServer).Average(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DoubleType/Average",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DoubleTypeServer).Average(ctx, req.(*NumberList))
	}
	return interceptor(ctx, in, info, handler)
}

// DoubleType_ServiceDesc is the grpc.ServiceDesc for DoubleType service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DoubleType_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DoubleType",
	HandlerType: (*DoubleTypeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Average",
			Handler:    _DoubleType_Average_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto/doubletype.proto",
}
