// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/signingManagement.proto

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

// SigningServiceClient is the client API for SigningService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SigningServiceClient interface {
	SignTxn(ctx context.Context, in *SignatureRequest, opts ...grpc.CallOption) (*SignatureResponse, error)
	BatchSignTxn(ctx context.Context, in *BatchSignatureRequest, opts ...grpc.CallOption) (*BatchSignatureResponse, error)
	GenerateNewKey(ctx context.Context, in *KeyManagementRequest, opts ...grpc.CallOption) (*KeyManagementResponse, error)
	DeleteKey(ctx context.Context, in *KeyManagementRequest, opts ...grpc.CallOption) (*KeyManagementResponse, error)
	GetKey(ctx context.Context, in *KeyManagementRequest, opts ...grpc.CallOption) (*KeyManagementResponse, error)
}

type signingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSigningServiceClient(cc grpc.ClientConnInterface) SigningServiceClient {
	return &signingServiceClient{cc}
}

func (c *signingServiceClient) SignTxn(ctx context.Context, in *SignatureRequest, opts ...grpc.CallOption) (*SignatureResponse, error) {
	out := new(SignatureResponse)
	err := c.cc.Invoke(ctx, "/SigningService/SignTxn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signingServiceClient) BatchSignTxn(ctx context.Context, in *BatchSignatureRequest, opts ...grpc.CallOption) (*BatchSignatureResponse, error) {
	out := new(BatchSignatureResponse)
	err := c.cc.Invoke(ctx, "/SigningService/BatchSignTxn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signingServiceClient) GenerateNewKey(ctx context.Context, in *KeyManagementRequest, opts ...grpc.CallOption) (*KeyManagementResponse, error) {
	out := new(KeyManagementResponse)
	err := c.cc.Invoke(ctx, "/SigningService/GenerateNewKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signingServiceClient) DeleteKey(ctx context.Context, in *KeyManagementRequest, opts ...grpc.CallOption) (*KeyManagementResponse, error) {
	out := new(KeyManagementResponse)
	err := c.cc.Invoke(ctx, "/SigningService/DeleteKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signingServiceClient) GetKey(ctx context.Context, in *KeyManagementRequest, opts ...grpc.CallOption) (*KeyManagementResponse, error) {
	out := new(KeyManagementResponse)
	err := c.cc.Invoke(ctx, "/SigningService/GetKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SigningServiceServer is the server API for SigningService service.
// All implementations must embed UnimplementedSigningServiceServer
// for forward compatibility
type SigningServiceServer interface {
	SignTxn(context.Context, *SignatureRequest) (*SignatureResponse, error)
	BatchSignTxn(context.Context, *BatchSignatureRequest) (*BatchSignatureResponse, error)
	GenerateNewKey(context.Context, *KeyManagementRequest) (*KeyManagementResponse, error)
	DeleteKey(context.Context, *KeyManagementRequest) (*KeyManagementResponse, error)
	GetKey(context.Context, *KeyManagementRequest) (*KeyManagementResponse, error)
	mustEmbedUnimplementedSigningServiceServer()
}

// UnimplementedSigningServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSigningServiceServer struct {
}

func (UnimplementedSigningServiceServer) SignTxn(context.Context, *SignatureRequest) (*SignatureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignTxn not implemented")
}
func (UnimplementedSigningServiceServer) BatchSignTxn(context.Context, *BatchSignatureRequest) (*BatchSignatureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchSignTxn not implemented")
}
func (UnimplementedSigningServiceServer) GenerateNewKey(context.Context, *KeyManagementRequest) (*KeyManagementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateNewKey not implemented")
}
func (UnimplementedSigningServiceServer) DeleteKey(context.Context, *KeyManagementRequest) (*KeyManagementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteKey not implemented")
}
func (UnimplementedSigningServiceServer) GetKey(context.Context, *KeyManagementRequest) (*KeyManagementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKey not implemented")
}
func (UnimplementedSigningServiceServer) mustEmbedUnimplementedSigningServiceServer() {}

// UnsafeSigningServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SigningServiceServer will
// result in compilation errors.
type UnsafeSigningServiceServer interface {
	mustEmbedUnimplementedSigningServiceServer()
}

func RegisterSigningServiceServer(s grpc.ServiceRegistrar, srv SigningServiceServer) {
	s.RegisterService(&SigningService_ServiceDesc, srv)
}

func _SigningService_SignTxn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SigningServiceServer).SignTxn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SigningService/SignTxn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SigningServiceServer).SignTxn(ctx, req.(*SignatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SigningService_BatchSignTxn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchSignatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SigningServiceServer).BatchSignTxn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SigningService/BatchSignTxn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SigningServiceServer).BatchSignTxn(ctx, req.(*BatchSignatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SigningService_GenerateNewKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyManagementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SigningServiceServer).GenerateNewKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SigningService/GenerateNewKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SigningServiceServer).GenerateNewKey(ctx, req.(*KeyManagementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SigningService_DeleteKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyManagementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SigningServiceServer).DeleteKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SigningService/DeleteKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SigningServiceServer).DeleteKey(ctx, req.(*KeyManagementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SigningService_GetKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyManagementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SigningServiceServer).GetKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SigningService/GetKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SigningServiceServer).GetKey(ctx, req.(*KeyManagementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SigningService_ServiceDesc is the grpc.ServiceDesc for SigningService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SigningService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SigningService",
	HandlerType: (*SigningServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignTxn",
			Handler:    _SigningService_SignTxn_Handler,
		},
		{
			MethodName: "BatchSignTxn",
			Handler:    _SigningService_BatchSignTxn_Handler,
		},
		{
			MethodName: "GenerateNewKey",
			Handler:    _SigningService_GenerateNewKey_Handler,
		},
		{
			MethodName: "DeleteKey",
			Handler:    _SigningService_DeleteKey_Handler,
		},
		{
			MethodName: "GetKey",
			Handler:    _SigningService_GetKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/signingManagement.proto",
}

// VerificationServiceClient is the client API for VerificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VerificationServiceClient interface {
	Verify(ctx context.Context, in *SignatureVerificationRequest, opts ...grpc.CallOption) (*SignatureVerificationResponse, error)
}

type verificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVerificationServiceClient(cc grpc.ClientConnInterface) VerificationServiceClient {
	return &verificationServiceClient{cc}
}

func (c *verificationServiceClient) Verify(ctx context.Context, in *SignatureVerificationRequest, opts ...grpc.CallOption) (*SignatureVerificationResponse, error) {
	out := new(SignatureVerificationResponse)
	err := c.cc.Invoke(ctx, "/VerificationService/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationServiceServer is the server API for VerificationService service.
// All implementations must embed UnimplementedVerificationServiceServer
// for forward compatibility
type VerificationServiceServer interface {
	Verify(context.Context, *SignatureVerificationRequest) (*SignatureVerificationResponse, error)
	mustEmbedUnimplementedVerificationServiceServer()
}

// UnimplementedVerificationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVerificationServiceServer struct {
}

func (UnimplementedVerificationServiceServer) Verify(context.Context, *SignatureVerificationRequest) (*SignatureVerificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedVerificationServiceServer) mustEmbedUnimplementedVerificationServiceServer() {}

// UnsafeVerificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VerificationServiceServer will
// result in compilation errors.
type UnsafeVerificationServiceServer interface {
	mustEmbedUnimplementedVerificationServiceServer()
}

func RegisterVerificationServiceServer(s grpc.ServiceRegistrar, srv VerificationServiceServer) {
	s.RegisterService(&VerificationService_ServiceDesc, srv)
}

func _VerificationService_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignatureVerificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerificationServiceServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VerificationService/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerificationServiceServer).Verify(ctx, req.(*SignatureVerificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VerificationService_ServiceDesc is the grpc.ServiceDesc for VerificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VerificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "VerificationService",
	HandlerType: (*VerificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Verify",
			Handler:    _VerificationService_Verify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/signingManagement.proto",
}
