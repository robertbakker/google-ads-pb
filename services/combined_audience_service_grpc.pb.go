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

// CombinedAudienceServiceClient is the client API for CombinedAudienceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CombinedAudienceServiceClient interface {
	// Returns the requested combined audience in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetCombinedAudience(ctx context.Context, in *GetCombinedAudienceRequest, opts ...grpc.CallOption) (*resources.CombinedAudience, error)
}

type combinedAudienceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCombinedAudienceServiceClient(cc grpc.ClientConnInterface) CombinedAudienceServiceClient {
	return &combinedAudienceServiceClient{cc}
}

func (c *combinedAudienceServiceClient) GetCombinedAudience(ctx context.Context, in *GetCombinedAudienceRequest, opts ...grpc.CallOption) (*resources.CombinedAudience, error) {
	out := new(resources.CombinedAudience)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.CombinedAudienceService/GetCombinedAudience", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CombinedAudienceServiceServer is the server API for CombinedAudienceService service.
// All implementations must embed UnimplementedCombinedAudienceServiceServer
// for forward compatibility
type CombinedAudienceServiceServer interface {
	// Returns the requested combined audience in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetCombinedAudience(context.Context, *GetCombinedAudienceRequest) (*resources.CombinedAudience, error)
	mustEmbedUnimplementedCombinedAudienceServiceServer()
}

// UnimplementedCombinedAudienceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCombinedAudienceServiceServer struct {
}

func (UnimplementedCombinedAudienceServiceServer) GetCombinedAudience(context.Context, *GetCombinedAudienceRequest) (*resources.CombinedAudience, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCombinedAudience not implemented")
}
func (UnimplementedCombinedAudienceServiceServer) mustEmbedUnimplementedCombinedAudienceServiceServer() {
}

// UnsafeCombinedAudienceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CombinedAudienceServiceServer will
// result in compilation errors.
type UnsafeCombinedAudienceServiceServer interface {
	mustEmbedUnimplementedCombinedAudienceServiceServer()
}

func RegisterCombinedAudienceServiceServer(s grpc.ServiceRegistrar, srv CombinedAudienceServiceServer) {
	s.RegisterService(&CombinedAudienceService_ServiceDesc, srv)
}

func _CombinedAudienceService_GetCombinedAudience_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCombinedAudienceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CombinedAudienceServiceServer).GetCombinedAudience(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.CombinedAudienceService/GetCombinedAudience",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CombinedAudienceServiceServer).GetCombinedAudience(ctx, req.(*GetCombinedAudienceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CombinedAudienceService_ServiceDesc is the grpc.ServiceDesc for CombinedAudienceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CombinedAudienceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.CombinedAudienceService",
	HandlerType: (*CombinedAudienceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCombinedAudience",
			Handler:    _CombinedAudienceService_GetCombinedAudience_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/combined_audience_service.proto",
}
