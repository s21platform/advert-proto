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
	AdvertService_GetAdverts_FullMethodName   = "/AdvertService/GetAdverts"
	AdvertService_CreateAdvert_FullMethodName = "/AdvertService/CreateAdvert"
	AdvertService_CancelAdvert_FullMethodName = "/AdvertService/CancelAdvert"
)

// AdvertServiceClient is the client API for AdvertService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdvertServiceClient interface {
	GetAdverts(ctx context.Context, in *AdvertEmpty, opts ...grpc.CallOption) (*GetAdvertsOut, error)
	CreateAdvert(ctx context.Context, in *CreateAdvertIn, opts ...grpc.CallOption) (*AdvertEmpty, error)
	CancelAdvert(ctx context.Context, in *CancelAdvertIn, opts ...grpc.CallOption) (*AdvertEmpty, error)
}

type advertServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdvertServiceClient(cc grpc.ClientConnInterface) AdvertServiceClient {
	return &advertServiceClient{cc}
}

func (c *advertServiceClient) GetAdverts(ctx context.Context, in *AdvertEmpty, opts ...grpc.CallOption) (*GetAdvertsOut, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAdvertsOut)
	err := c.cc.Invoke(ctx, AdvertService_GetAdverts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertServiceClient) CreateAdvert(ctx context.Context, in *CreateAdvertIn, opts ...grpc.CallOption) (*AdvertEmpty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdvertEmpty)
	err := c.cc.Invoke(ctx, AdvertService_CreateAdvert_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *advertServiceClient) CancelAdvert(ctx context.Context, in *CancelAdvertIn, opts ...grpc.CallOption) (*AdvertEmpty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdvertEmpty)
	err := c.cc.Invoke(ctx, AdvertService_CancelAdvert_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdvertServiceServer is the server API for AdvertService service.
// All implementations must embed UnimplementedAdvertServiceServer
// for forward compatibility.
type AdvertServiceServer interface {
	GetAdverts(context.Context, *AdvertEmpty) (*GetAdvertsOut, error)
	CreateAdvert(context.Context, *CreateAdvertIn) (*AdvertEmpty, error)
	CancelAdvert(context.Context, *CancelAdvertIn) (*AdvertEmpty, error)
	mustEmbedUnimplementedAdvertServiceServer()
}

// UnimplementedAdvertServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAdvertServiceServer struct{}

func (UnimplementedAdvertServiceServer) GetAdverts(context.Context, *AdvertEmpty) (*GetAdvertsOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdverts not implemented")
}
func (UnimplementedAdvertServiceServer) CreateAdvert(context.Context, *CreateAdvertIn) (*AdvertEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAdvert not implemented")
}
func (UnimplementedAdvertServiceServer) CancelAdvert(context.Context, *CancelAdvertIn) (*AdvertEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelAdvert not implemented")
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

func _AdvertService_GetAdverts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdvertEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertServiceServer).GetAdverts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdvertService_GetAdverts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertServiceServer).GetAdverts(ctx, req.(*AdvertEmpty))
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

func _AdvertService_CancelAdvert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelAdvertIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdvertServiceServer).CancelAdvert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdvertService_CancelAdvert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdvertServiceServer).CancelAdvert(ctx, req.(*CancelAdvertIn))
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
			MethodName: "GetAdverts",
			Handler:    _AdvertService_GetAdverts_Handler,
		},
		{
			MethodName: "CreateAdvert",
			Handler:    _AdvertService_CreateAdvert_Handler,
		},
		{
			MethodName: "CancelAdvert",
			Handler:    _AdvertService_CancelAdvert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "advert.proto",
}
