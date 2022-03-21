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

// DisplayKeywordViewServiceClient is the client API for DisplayKeywordViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DisplayKeywordViewServiceClient interface {
	// Returns the requested display keyword view in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetDisplayKeywordView(ctx context.Context, in *GetDisplayKeywordViewRequest, opts ...grpc.CallOption) (*resources.DisplayKeywordView, error)
}

type displayKeywordViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDisplayKeywordViewServiceClient(cc grpc.ClientConnInterface) DisplayKeywordViewServiceClient {
	return &displayKeywordViewServiceClient{cc}
}

func (c *displayKeywordViewServiceClient) GetDisplayKeywordView(ctx context.Context, in *GetDisplayKeywordViewRequest, opts ...grpc.CallOption) (*resources.DisplayKeywordView, error) {
	out := new(resources.DisplayKeywordView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.DisplayKeywordViewService/GetDisplayKeywordView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DisplayKeywordViewServiceServer is the server API for DisplayKeywordViewService service.
// All implementations must embed UnimplementedDisplayKeywordViewServiceServer
// for forward compatibility
type DisplayKeywordViewServiceServer interface {
	// Returns the requested display keyword view in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetDisplayKeywordView(context.Context, *GetDisplayKeywordViewRequest) (*resources.DisplayKeywordView, error)
	mustEmbedUnimplementedDisplayKeywordViewServiceServer()
}

// UnimplementedDisplayKeywordViewServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDisplayKeywordViewServiceServer struct {
}

func (UnimplementedDisplayKeywordViewServiceServer) GetDisplayKeywordView(context.Context, *GetDisplayKeywordViewRequest) (*resources.DisplayKeywordView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDisplayKeywordView not implemented")
}
func (UnimplementedDisplayKeywordViewServiceServer) mustEmbedUnimplementedDisplayKeywordViewServiceServer() {
}

// UnsafeDisplayKeywordViewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DisplayKeywordViewServiceServer will
// result in compilation errors.
type UnsafeDisplayKeywordViewServiceServer interface {
	mustEmbedUnimplementedDisplayKeywordViewServiceServer()
}

func RegisterDisplayKeywordViewServiceServer(s grpc.ServiceRegistrar, srv DisplayKeywordViewServiceServer) {
	s.RegisterService(&DisplayKeywordViewService_ServiceDesc, srv)
}

func _DisplayKeywordViewService_GetDisplayKeywordView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDisplayKeywordViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisplayKeywordViewServiceServer).GetDisplayKeywordView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.DisplayKeywordViewService/GetDisplayKeywordView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisplayKeywordViewServiceServer).GetDisplayKeywordView(ctx, req.(*GetDisplayKeywordViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DisplayKeywordViewService_ServiceDesc is the grpc.ServiceDesc for DisplayKeywordViewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DisplayKeywordViewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.DisplayKeywordViewService",
	HandlerType: (*DisplayKeywordViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDisplayKeywordView",
			Handler:    _DisplayKeywordViewService_GetDisplayKeywordView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/display_keyword_view_service.proto",
}
