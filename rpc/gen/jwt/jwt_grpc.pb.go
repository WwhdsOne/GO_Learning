// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.3
// source: jwt.proto

package jwt_grpc

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
	JwtService_Login_FullMethodName        = "/jwt_grpc.JwtService/Login"
	JwtService_RefreshToken_FullMethodName = "/jwt_grpc.JwtService/RefreshToken"
	JwtService_Logout_FullMethodName       = "/jwt_grpc.JwtService/Logout"
)

// JwtServiceClient is the client API for JwtService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 定义服务
type JwtServiceClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*TokenResponse, error)
	RefreshToken(ctx context.Context, in *RefreshReq, opts ...grpc.CallOption) (*TokenResponse, error)
	Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*Response, error)
}

type jwtServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJwtServiceClient(cc grpc.ClientConnInterface) JwtServiceClient {
	return &jwtServiceClient{cc}
}

func (c *jwtServiceClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*TokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, JwtService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jwtServiceClient) RefreshToken(ctx context.Context, in *RefreshReq, opts ...grpc.CallOption) (*TokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, JwtService_RefreshToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jwtServiceClient) Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, JwtService_Logout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JwtServiceServer is the server API for JwtService service.
// All implementations must embed UnimplementedJwtServiceServer
// for forward compatibility.
//
// 定义服务
type JwtServiceServer interface {
	Login(context.Context, *LoginReq) (*TokenResponse, error)
	RefreshToken(context.Context, *RefreshReq) (*TokenResponse, error)
	Logout(context.Context, *LogoutReq) (*Response, error)
	mustEmbedUnimplementedJwtServiceServer()
}

// UnimplementedJwtServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedJwtServiceServer struct{}

func (UnimplementedJwtServiceServer) Login(context.Context, *LoginReq) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedJwtServiceServer) RefreshToken(context.Context, *RefreshReq) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedJwtServiceServer) Logout(context.Context, *LogoutReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedJwtServiceServer) mustEmbedUnimplementedJwtServiceServer() {}
func (UnimplementedJwtServiceServer) testEmbeddedByValue()                    {}

// UnsafeJwtServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JwtServiceServer will
// result in compilation errors.
type UnsafeJwtServiceServer interface {
	mustEmbedUnimplementedJwtServiceServer()
}

func RegisterJwtServiceServer(s grpc.ServiceRegistrar, srv JwtServiceServer) {
	// If the following call pancis, it indicates UnimplementedJwtServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&JwtService_ServiceDesc, srv)
}

func _JwtService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JwtServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JwtService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JwtServiceServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _JwtService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JwtServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JwtService_RefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JwtServiceServer).RefreshToken(ctx, req.(*RefreshReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _JwtService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JwtServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JwtService_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JwtServiceServer).Logout(ctx, req.(*LogoutReq))
	}
	return interceptor(ctx, in, info, handler)
}

// JwtService_ServiceDesc is the grpc.ServiceDesc for JwtService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JwtService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jwt_grpc.JwtService",
	HandlerType: (*JwtServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _JwtService_Login_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _JwtService_RefreshToken_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _JwtService_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jwt.proto",
}
