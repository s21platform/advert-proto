// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.1
// source: advert.proto

package advert_proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AdvertService_GetAdvert_FullMethodName    = "/AdvertService/GetAdvert"
	AdvertService_CreateAdvert_FullMethodName = "/AdvertService/CreateAdvert"
)

// AdvertServiceClient is the client API for AdvertService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdvertServiceClient interface {
	GetAdvert(ctx context.Context, in *GetAdvertIn, opts ...grpc.CallOption) (*GetAdvertOut, error)
	CreateAdvert(ctx context.Context, in *CreateAdvertIn, opts ...grpc.CallOption) (*CreateAdvertOut, error)
}

type advertServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdvertServiceClient(cc grpc.ClientConnInterface) AdvertServiceClient {
	return &advertServiceClient{cc}
}

func (c *advertServiceClient) GetAdvert(ctx context.Context, in *GetAdvertIn, opts ...grpc.CallOption) (*GetAdvertOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAdvertOut)
	err := c.cc.Invoke(ctx, AdvertService_GetAdvert_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertServiceClient) CreateAdvert(ctx context.Context, in *CreateAdvertIn, opts ...grpc.CallOption) (*CreateAdvertOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateAdvertOut)
	err := c.cc.Invoke(ctx, AdvertService_CreateAdvert_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdvertServiceServer is the server API for AdvertService service.
// All implementations must embed UnimplementedAdvertServiceServer
// for forward compatibility.
type AdvertServiceServer interface {
	GetAdvert(context.Context, *GetAdvertIn) (*GetAdvertOut, error)
	CreateAdvert(context.Context, *CreateAdvertIn) (*CreateAdvertOut, error)
	mustEmbedUnimplementedAdvertServiceServer()
}

// UnimplementedAdvertServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAdvertServiceServer struct{}

func (UnimplementedAdvertServiceServer) GetAdvert(context.Context, *GetAdvertIn) (*GetAdvertOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdvert not implemented")
}
func (UnimplementedAdvertServiceServer) CreateAdvert(context.Context, *CreateAdvertIn) (*CreateAdvertOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAdvert not implemented")
}
func (UnimplementedAdvertServiceServer) mustEmbedUnimplementedAdvertServiceServer() {}
func (UnimplementedAdvertServiceServer) testEmbeddedByValue()                       {}

// UnsafeAdvertServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdvertServiceServer will
// result in compilation errors.
type UnsafeAdvertServiceServer interface {
	mustEmbedUnimplementedAdvertServiceServer()
}

func RegisterAdvertServiceServer(s grpc.ServiceRegistrar, srv AdvertServiceServer) {
	// If the following call pancis, it indicates UnimplementedAdvertServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AdvertService_ServiceDesc, srv)
}

func _AdvertService_GetAdvert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdvertIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertServiceServer).GetAdvert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdvertService_GetAdvert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertServiceServer).GetAdvert(ctx, req.(*GetAdvertIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdvertService_CreateAdvert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAdvertIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertServiceServer).CreateAdvert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdvertService_CreateAdvert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertServiceServer).CreateAdvert(ctx, req.(*CreateAdvertIn))
	}
	return interceptor(ctx, in, info, handler)
}

// AdvertService_ServiceDesc is the grpc.ServiceDesc for AdvertService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdvertService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AdvertService",
	HandlerType: (*AdvertServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdvert",
			Handler:    _AdvertService_GetAdvert_Handler,
		},
		{
			MethodName: "CreateAdvert",
			Handler:    _AdvertService_CreateAdvert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "advert.proto",
}
