// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthHandlerClient is the client API for AuthHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthHandlerClient interface {
	Create(ctx context.Context, in *CreateSession, opts ...grpc.CallOption) (*SessionValue, error)
	GetUser(ctx context.Context, in *SessionValue, opts ...grpc.CallOption) (*UserValue, error)
	Delete(ctx context.Context, in *SessionValue, opts ...grpc.CallOption) (*empty.Empty, error)
}

type authHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthHandlerClient(cc grpc.ClientConnInterface) AuthHandlerClient {
	return &authHandlerClient{cc}
}

func (c *authHandlerClient) Create(ctx context.Context, in *CreateSession, opts ...grpc.CallOption) (*SessionValue, error) {
	out := new(SessionValue)
	err := c.cc.Invoke(ctx, "/auth.AuthHandler/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authHandlerClient) GetUser(ctx context.Context, in *SessionValue, opts ...grpc.CallOption) (*UserValue, error) {
	out := new(UserValue)
	err := c.cc.Invoke(ctx, "/auth.AuthHandler/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authHandlerClient) Delete(ctx context.Context, in *SessionValue, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/auth.AuthHandler/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthHandlerServer is the server API for AuthHandler service.
// All implementations should embed UnimplementedAuthHandlerServer
// for forward compatibility
type AuthHandlerServer interface {
	Create(context.Context, *CreateSession) (*SessionValue, error)
	GetUser(context.Context, *SessionValue) (*UserValue, error)
	Delete(context.Context, *SessionValue) (*empty.Empty, error)
}

// UnimplementedAuthHandlerServer should be embedded to have forward compatible implementations.
type UnimplementedAuthHandlerServer struct {
}

func (UnimplementedAuthHandlerServer) Create(context.Context, *CreateSession) (*SessionValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAuthHandlerServer) GetUser(context.Context, *SessionValue) (*UserValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedAuthHandlerServer) Delete(context.Context, *SessionValue) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeAuthHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthHandlerServer will
// result in compilation errors.
type UnsafeAuthHandlerServer interface {
	mustEmbedUnimplementedAuthHandlerServer()
}

func RegisterAuthHandlerServer(s grpc.ServiceRegistrar, srv AuthHandlerServer) {
	s.RegisterService(&AuthHandler_ServiceDesc, srv)
}

func _AuthHandler_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSession)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthHandlerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthHandler/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthHandlerServer).Create(ctx, req.(*CreateSession))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthHandler_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthHandlerServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthHandler/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthHandlerServer).GetUser(ctx, req.(*SessionValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthHandler_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthHandlerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthHandler/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthHandlerServer).Delete(ctx, req.(*SessionValue))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthHandler_ServiceDesc is the grpc.ServiceDesc for AuthHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthHandler",
	HandlerType: (*AuthHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AuthHandler_Create_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _AuthHandler_GetUser_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AuthHandler_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/proto/session.proto",
}