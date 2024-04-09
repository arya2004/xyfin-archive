// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: service_xyfin.proto

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

// XyfinClient is the client API for Xyfin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type XyfinClient interface {
	Createuser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	Loginuser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
}

type xyfinClient struct {
	cc grpc.ClientConnInterface
}

func NewXyfinClient(cc grpc.ClientConnInterface) XyfinClient {
	return &xyfinClient{cc}
}

func (c *xyfinClient) Createuser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/pb.Xyfin/Createuser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xyfinClient) Loginuser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, "/pb.Xyfin/Loginuser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// XyfinServer is the server API for Xyfin service.
// All implementations must embed UnimplementedXyfinServer
// for forward compatibility
type XyfinServer interface {
	Createuser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	Loginuser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	mustEmbedUnimplementedXyfinServer()
}

// UnimplementedXyfinServer must be embedded to have forward compatible implementations.
type UnimplementedXyfinServer struct {
}

func (UnimplementedXyfinServer) Createuser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Createuser not implemented")
}
func (UnimplementedXyfinServer) Loginuser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Loginuser not implemented")
}
func (UnimplementedXyfinServer) mustEmbedUnimplementedXyfinServer() {}

// UnsafeXyfinServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to XyfinServer will
// result in compilation errors.
type UnsafeXyfinServer interface {
	mustEmbedUnimplementedXyfinServer()
}

func RegisterXyfinServer(s grpc.ServiceRegistrar, srv XyfinServer) {
	s.RegisterService(&Xyfin_ServiceDesc, srv)
}

func _Xyfin_Createuser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XyfinServer).Createuser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Xyfin/Createuser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XyfinServer).Createuser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Xyfin_Loginuser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XyfinServer).Loginuser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Xyfin/Loginuser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XyfinServer).Loginuser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Xyfin_ServiceDesc is the grpc.ServiceDesc for Xyfin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Xyfin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Xyfin",
	HandlerType: (*XyfinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Createuser",
			Handler:    _Xyfin_Createuser_Handler,
		},
		{
			MethodName: "Loginuser",
			Handler:    _Xyfin_Loginuser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_xyfin.proto",
}