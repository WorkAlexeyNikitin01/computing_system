// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: service.proto

package grpc

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

// StoreroomServiceClient is the client API for StoreroomService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StoreroomServiceClient interface {
	GetFromStoreroom(ctx context.Context, in *StoreProductRequest, opts ...grpc.CallOption) (*SProductResponse, error)
}

type storeroomServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreroomServiceClient(cc grpc.ClientConnInterface) StoreroomServiceClient {
	return &storeroomServiceClient{cc}
}

func (c *storeroomServiceClient) GetFromStoreroom(ctx context.Context, in *StoreProductRequest, opts ...grpc.CallOption) (*SProductResponse, error) {
	out := new(SProductResponse)
	err := c.cc.Invoke(ctx, "/storeroom.StoreroomService/GetFromStoreroom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreroomServiceServer is the server API for StoreroomService service.
// All implementations should embed UnimplementedStoreroomServiceServer
// for forward compatibility
type StoreroomServiceServer interface {
	GetFromStoreroom(context.Context, *StoreProductRequest) (*SProductResponse, error)
}

// UnimplementedStoreroomServiceServer should be embedded to have forward compatible implementations.
type UnimplementedStoreroomServiceServer struct {
}

func (UnimplementedStoreroomServiceServer) GetFromStoreroom(context.Context, *StoreProductRequest) (*SProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFromStoreroom not implemented")
}

// UnsafeStoreroomServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoreroomServiceServer will
// result in compilation errors.
type UnsafeStoreroomServiceServer interface {
	mustEmbedUnimplementedStoreroomServiceServer()
}

func RegisterStoreroomServiceServer(s grpc.ServiceRegistrar, srv StoreroomServiceServer) {
	s.RegisterService(&StoreroomService_ServiceDesc, srv)
}

func _StoreroomService_GetFromStoreroom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreroomServiceServer).GetFromStoreroom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storeroom.StoreroomService/GetFromStoreroom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreroomServiceServer).GetFromStoreroom(ctx, req.(*StoreProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StoreroomService_ServiceDesc is the grpc.ServiceDesc for StoreroomService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StoreroomService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "storeroom.StoreroomService",
	HandlerType: (*StoreroomServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFromStoreroom",
			Handler:    _StoreroomService_GetFromStoreroom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
