// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package services

import (
	context "context"
	resources "github.com/shenzhencenter/google-ads-pb/resources"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MobileAppCategoryConstantServiceClient is the client API for MobileAppCategoryConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MobileAppCategoryConstantServiceClient interface {
	// Returns the requested mobile app category constant.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetMobileAppCategoryConstant(ctx context.Context, in *GetMobileAppCategoryConstantRequest, opts ...grpc.CallOption) (*resources.MobileAppCategoryConstant, error)
}

type mobileAppCategoryConstantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMobileAppCategoryConstantServiceClient(cc grpc.ClientConnInterface) MobileAppCategoryConstantServiceClient {
	return &mobileAppCategoryConstantServiceClient{cc}
}

func (c *mobileAppCategoryConstantServiceClient) GetMobileAppCategoryConstant(ctx context.Context, in *GetMobileAppCategoryConstantRequest, opts ...grpc.CallOption) (*resources.MobileAppCategoryConstant, error) {
	out := new(resources.MobileAppCategoryConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.MobileAppCategoryConstantService/GetMobileAppCategoryConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MobileAppCategoryConstantServiceServer is the server API for MobileAppCategoryConstantService service.
// All implementations must embed UnimplementedMobileAppCategoryConstantServiceServer
// for forward compatibility
type MobileAppCategoryConstantServiceServer interface {
	// Returns the requested mobile app category constant.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetMobileAppCategoryConstant(context.Context, *GetMobileAppCategoryConstantRequest) (*resources.MobileAppCategoryConstant, error)
	mustEmbedUnimplementedMobileAppCategoryConstantServiceServer()
}

// UnimplementedMobileAppCategoryConstantServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMobileAppCategoryConstantServiceServer struct {
}

func (UnimplementedMobileAppCategoryConstantServiceServer) GetMobileAppCategoryConstant(context.Context, *GetMobileAppCategoryConstantRequest) (*resources.MobileAppCategoryConstant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMobileAppCategoryConstant not implemented")
}
func (UnimplementedMobileAppCategoryConstantServiceServer) mustEmbedUnimplementedMobileAppCategoryConstantServiceServer() {
}

// UnsafeMobileAppCategoryConstantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MobileAppCategoryConstantServiceServer will
// result in compilation errors.
type UnsafeMobileAppCategoryConstantServiceServer interface {
	mustEmbedUnimplementedMobileAppCategoryConstantServiceServer()
}

func RegisterMobileAppCategoryConstantServiceServer(s grpc.ServiceRegistrar, srv MobileAppCategoryConstantServiceServer) {
	s.RegisterService(&MobileAppCategoryConstantService_ServiceDesc, srv)
}

func _MobileAppCategoryConstantService_GetMobileAppCategoryConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMobileAppCategoryConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileAppCategoryConstantServiceServer).GetMobileAppCategoryConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.MobileAppCategoryConstantService/GetMobileAppCategoryConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileAppCategoryConstantServiceServer).GetMobileAppCategoryConstant(ctx, req.(*GetMobileAppCategoryConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MobileAppCategoryConstantService_ServiceDesc is the grpc.ServiceDesc for MobileAppCategoryConstantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MobileAppCategoryConstantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.MobileAppCategoryConstantService",
	HandlerType: (*MobileAppCategoryConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMobileAppCategoryConstant",
			Handler:    _MobileAppCategoryConstantService_GetMobileAppCategoryConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/mobile_app_category_constant_service.proto",
}
