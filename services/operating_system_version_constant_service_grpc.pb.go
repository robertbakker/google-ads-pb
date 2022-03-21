// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package services

import (
	context "context"
	resources "github.com/robertbakker/google-ads-pb/resources"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OperatingSystemVersionConstantServiceClient is the client API for OperatingSystemVersionConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OperatingSystemVersionConstantServiceClient interface {
	// Returns the requested OS version constant in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetOperatingSystemVersionConstant(ctx context.Context, in *GetOperatingSystemVersionConstantRequest, opts ...grpc.CallOption) (*resources.OperatingSystemVersionConstant, error)
}

type operatingSystemVersionConstantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOperatingSystemVersionConstantServiceClient(cc grpc.ClientConnInterface) OperatingSystemVersionConstantServiceClient {
	return &operatingSystemVersionConstantServiceClient{cc}
}

func (c *operatingSystemVersionConstantServiceClient) GetOperatingSystemVersionConstant(ctx context.Context, in *GetOperatingSystemVersionConstantRequest, opts ...grpc.CallOption) (*resources.OperatingSystemVersionConstant, error) {
	out := new(resources.OperatingSystemVersionConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.OperatingSystemVersionConstantService/GetOperatingSystemVersionConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperatingSystemVersionConstantServiceServer is the server API for OperatingSystemVersionConstantService service.
// All implementations must embed UnimplementedOperatingSystemVersionConstantServiceServer
// for forward compatibility
type OperatingSystemVersionConstantServiceServer interface {
	// Returns the requested OS version constant in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetOperatingSystemVersionConstant(context.Context, *GetOperatingSystemVersionConstantRequest) (*resources.OperatingSystemVersionConstant, error)
	mustEmbedUnimplementedOperatingSystemVersionConstantServiceServer()
}

// UnimplementedOperatingSystemVersionConstantServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOperatingSystemVersionConstantServiceServer struct {
}

func (UnimplementedOperatingSystemVersionConstantServiceServer) GetOperatingSystemVersionConstant(context.Context, *GetOperatingSystemVersionConstantRequest) (*resources.OperatingSystemVersionConstant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOperatingSystemVersionConstant not implemented")
}
func (UnimplementedOperatingSystemVersionConstantServiceServer) mustEmbedUnimplementedOperatingSystemVersionConstantServiceServer() {
}

// UnsafeOperatingSystemVersionConstantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OperatingSystemVersionConstantServiceServer will
// result in compilation errors.
type UnsafeOperatingSystemVersionConstantServiceServer interface {
	mustEmbedUnimplementedOperatingSystemVersionConstantServiceServer()
}

func RegisterOperatingSystemVersionConstantServiceServer(s grpc.ServiceRegistrar, srv OperatingSystemVersionConstantServiceServer) {
	s.RegisterService(&OperatingSystemVersionConstantService_ServiceDesc, srv)
}

func _OperatingSystemVersionConstantService_GetOperatingSystemVersionConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperatingSystemVersionConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatingSystemVersionConstantServiceServer).GetOperatingSystemVersionConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.OperatingSystemVersionConstantService/GetOperatingSystemVersionConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatingSystemVersionConstantServiceServer).GetOperatingSystemVersionConstant(ctx, req.(*GetOperatingSystemVersionConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OperatingSystemVersionConstantService_ServiceDesc is the grpc.ServiceDesc for OperatingSystemVersionConstantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OperatingSystemVersionConstantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.OperatingSystemVersionConstantService",
	HandlerType: (*OperatingSystemVersionConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOperatingSystemVersionConstant",
			Handler:    _OperatingSystemVersionConstantService_GetOperatingSystemVersionConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/operating_system_version_constant_service.proto",
}
