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

// ManagedPlacementViewServiceClient is the client API for ManagedPlacementViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagedPlacementViewServiceClient interface {
	// Returns the requested Managed Placement view in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetManagedPlacementView(ctx context.Context, in *GetManagedPlacementViewRequest, opts ...grpc.CallOption) (*resources.ManagedPlacementView, error)
}

type managedPlacementViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewManagedPlacementViewServiceClient(cc grpc.ClientConnInterface) ManagedPlacementViewServiceClient {
	return &managedPlacementViewServiceClient{cc}
}

func (c *managedPlacementViewServiceClient) GetManagedPlacementView(ctx context.Context, in *GetManagedPlacementViewRequest, opts ...grpc.CallOption) (*resources.ManagedPlacementView, error) {
	out := new(resources.ManagedPlacementView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.ManagedPlacementViewService/GetManagedPlacementView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagedPlacementViewServiceServer is the server API for ManagedPlacementViewService service.
// All implementations must embed UnimplementedManagedPlacementViewServiceServer
// for forward compatibility
type ManagedPlacementViewServiceServer interface {
	// Returns the requested Managed Placement view in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetManagedPlacementView(context.Context, *GetManagedPlacementViewRequest) (*resources.ManagedPlacementView, error)
	mustEmbedUnimplementedManagedPlacementViewServiceServer()
}

// UnimplementedManagedPlacementViewServiceServer must be embedded to have forward compatible implementations.
type UnimplementedManagedPlacementViewServiceServer struct {
}

func (UnimplementedManagedPlacementViewServiceServer) GetManagedPlacementView(context.Context, *GetManagedPlacementViewRequest) (*resources.ManagedPlacementView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetManagedPlacementView not implemented")
}
func (UnimplementedManagedPlacementViewServiceServer) mustEmbedUnimplementedManagedPlacementViewServiceServer() {
}

// UnsafeManagedPlacementViewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagedPlacementViewServiceServer will
// result in compilation errors.
type UnsafeManagedPlacementViewServiceServer interface {
	mustEmbedUnimplementedManagedPlacementViewServiceServer()
}

func RegisterManagedPlacementViewServiceServer(s grpc.ServiceRegistrar, srv ManagedPlacementViewServiceServer) {
	s.RegisterService(&ManagedPlacementViewService_ServiceDesc, srv)
}

func _ManagedPlacementViewService_GetManagedPlacementView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetManagedPlacementViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagedPlacementViewServiceServer).GetManagedPlacementView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.ManagedPlacementViewService/GetManagedPlacementView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagedPlacementViewServiceServer).GetManagedPlacementView(ctx, req.(*GetManagedPlacementViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ManagedPlacementViewService_ServiceDesc is the grpc.ServiceDesc for ManagedPlacementViewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ManagedPlacementViewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.ManagedPlacementViewService",
	HandlerType: (*ManagedPlacementViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetManagedPlacementView",
			Handler:    _ManagedPlacementViewService_GetManagedPlacementView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/managed_placement_view_service.proto",
}
