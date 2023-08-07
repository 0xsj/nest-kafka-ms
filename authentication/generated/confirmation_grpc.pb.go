// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: confirmation.proto

package __

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

const (
	ConfirmationService_SendPhoneNumberConfirmation_FullMethodName   = "/confirmation.ConfirmationService/SendPhoneNumberConfirmation"
	ConfirmationService_VerifyPhoneNumberConfirmation_FullMethodName = "/confirmation.ConfirmationService/VerifyPhoneNumberConfirmation"
)

// ConfirmationServiceClient is the client API for ConfirmationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfirmationServiceClient interface {
	SendPhoneNumberConfirmation(ctx context.Context, in *SendPhoneNumberConfirmationRequest, opts ...grpc.CallOption) (*SendPhoneNumberConfirmationResponse, error)
	VerifyPhoneNumberConfirmation(ctx context.Context, in *VerifyPhoneNumberConfirmationRequest, opts ...grpc.CallOption) (*VerifyPhoneNumberConfirmationResponse, error)
}

type confirmationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConfirmationServiceClient(cc grpc.ClientConnInterface) ConfirmationServiceClient {
	return &confirmationServiceClient{cc}
}

func (c *confirmationServiceClient) SendPhoneNumberConfirmation(ctx context.Context, in *SendPhoneNumberConfirmationRequest, opts ...grpc.CallOption) (*SendPhoneNumberConfirmationResponse, error) {
	out := new(SendPhoneNumberConfirmationResponse)
	err := c.cc.Invoke(ctx, ConfirmationService_SendPhoneNumberConfirmation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *confirmationServiceClient) VerifyPhoneNumberConfirmation(ctx context.Context, in *VerifyPhoneNumberConfirmationRequest, opts ...grpc.CallOption) (*VerifyPhoneNumberConfirmationResponse, error) {
	out := new(VerifyPhoneNumberConfirmationResponse)
	err := c.cc.Invoke(ctx, ConfirmationService_VerifyPhoneNumberConfirmation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfirmationServiceServer is the server API for ConfirmationService service.
// All implementations must embed UnimplementedConfirmationServiceServer
// for forward compatibility
type ConfirmationServiceServer interface {
	SendPhoneNumberConfirmation(context.Context, *SendPhoneNumberConfirmationRequest) (*SendPhoneNumberConfirmationResponse, error)
	VerifyPhoneNumberConfirmation(context.Context, *VerifyPhoneNumberConfirmationRequest) (*VerifyPhoneNumberConfirmationResponse, error)
	mustEmbedUnimplementedConfirmationServiceServer()
}

// UnimplementedConfirmationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConfirmationServiceServer struct {
}

func (UnimplementedConfirmationServiceServer) SendPhoneNumberConfirmation(context.Context, *SendPhoneNumberConfirmationRequest) (*SendPhoneNumberConfirmationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPhoneNumberConfirmation not implemented")
}
func (UnimplementedConfirmationServiceServer) VerifyPhoneNumberConfirmation(context.Context, *VerifyPhoneNumberConfirmationRequest) (*VerifyPhoneNumberConfirmationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyPhoneNumberConfirmation not implemented")
}
func (UnimplementedConfirmationServiceServer) mustEmbedUnimplementedConfirmationServiceServer() {}

// UnsafeConfirmationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfirmationServiceServer will
// result in compilation errors.
type UnsafeConfirmationServiceServer interface {
	mustEmbedUnimplementedConfirmationServiceServer()
}

func RegisterConfirmationServiceServer(s grpc.ServiceRegistrar, srv ConfirmationServiceServer) {
	s.RegisterService(&ConfirmationService_ServiceDesc, srv)
}

func _ConfirmationService_SendPhoneNumberConfirmation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendPhoneNumberConfirmationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfirmationServiceServer).SendPhoneNumberConfirmation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfirmationService_SendPhoneNumberConfirmation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfirmationServiceServer).SendPhoneNumberConfirmation(ctx, req.(*SendPhoneNumberConfirmationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfirmationService_VerifyPhoneNumberConfirmation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyPhoneNumberConfirmationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfirmationServiceServer).VerifyPhoneNumberConfirmation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfirmationService_VerifyPhoneNumberConfirmation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfirmationServiceServer).VerifyPhoneNumberConfirmation(ctx, req.(*VerifyPhoneNumberConfirmationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConfirmationService_ServiceDesc is the grpc.ServiceDesc for ConfirmationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConfirmationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "confirmation.ConfirmationService",
	HandlerType: (*ConfirmationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendPhoneNumberConfirmation",
			Handler:    _ConfirmationService_SendPhoneNumberConfirmation_Handler,
		},
		{
			MethodName: "VerifyPhoneNumberConfirmation",
			Handler:    _ConfirmationService_VerifyPhoneNumberConfirmation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "confirmation.proto",
}