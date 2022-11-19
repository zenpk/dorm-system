// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: token.proto

package token

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

// TokenClient is the client API for Token service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TokenClient interface {
	GenAllToken(ctx context.Context, in *GenAllTokenRequest, opts ...grpc.CallOption) (*TokenReply, error)
	GenAccessToken(ctx context.Context, in *GenAccessTokenRequest, opts ...grpc.CallOption) (*TokenReply, error)
}

type tokenClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenClient(cc grpc.ClientConnInterface) TokenClient {
	return &tokenClient{cc}
}

func (c *tokenClient) GenAllToken(ctx context.Context, in *GenAllTokenRequest, opts ...grpc.CallOption) (*TokenReply, error) {
	out := new(TokenReply)
	err := c.cc.Invoke(ctx, "/token.Token/GenAllToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenClient) GenAccessToken(ctx context.Context, in *GenAccessTokenRequest, opts ...grpc.CallOption) (*TokenReply, error) {
	out := new(TokenReply)
	err := c.cc.Invoke(ctx, "/token.Token/GenAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenServer is the server API for Token service.
// All implementations must embed UnimplementedTokenServer
// for forward compatibility
type TokenServer interface {
	GenAllToken(context.Context, *GenAllTokenRequest) (*TokenReply, error)
	GenAccessToken(context.Context, *GenAccessTokenRequest) (*TokenReply, error)
	mustEmbedUnimplementedTokenServer()
}

// UnimplementedTokenServer must be embedded to have forward compatible implementations.
type UnimplementedTokenServer struct {
}

func (UnimplementedTokenServer) GenAllToken(context.Context, *GenAllTokenRequest) (*TokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenAllToken not implemented")
}
func (UnimplementedTokenServer) GenAccessToken(context.Context, *GenAccessTokenRequest) (*TokenReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenAccessToken not implemented")
}
func (UnimplementedTokenServer) mustEmbedUnimplementedTokenServer() {}

// UnsafeTokenServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TokenServer will
// result in compilation errors.
type UnsafeTokenServer interface {
	mustEmbedUnimplementedTokenServer()
}

func RegisterTokenServer(s grpc.ServiceRegistrar, srv TokenServer) {
	s.RegisterService(&Token_ServiceDesc, srv)
}

func _Token_GenAllToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenAllTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServer).GenAllToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/token.Token/GenAllToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServer).GenAllToken(ctx, req.(*GenAllTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Token_GenAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServer).GenAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/token.Token/GenAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServer).GenAccessToken(ctx, req.(*GenAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Token_ServiceDesc is the grpc.ServiceDesc for Token service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Token_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "token.Token",
	HandlerType: (*TokenServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenAllToken",
			Handler:    _Token_GenAllToken_Handler,
		},
		{
			MethodName: "GenAccessToken",
			Handler:    _Token_GenAccessToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "token.proto",
}