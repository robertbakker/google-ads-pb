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

// ConversionValueRuleSetServiceClient is the client API for ConversionValueRuleSetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConversionValueRuleSetServiceClient interface {
	// Returns the requested conversion value rule set.
	GetConversionValueRuleSet(ctx context.Context, in *GetConversionValueRuleSetRequest, opts ...grpc.CallOption) (*resources.ConversionValueRuleSet, error)
	// Creates, updates or removes conversion value rule sets. Operation statuses
	// are returned.
	MutateConversionValueRuleSets(ctx context.Context, in *MutateConversionValueRuleSetsRequest, opts ...grpc.CallOption) (*MutateConversionValueRuleSetsResponse, error)
}

type conversionValueRuleSetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConversionValueRuleSetServiceClient(cc grpc.ClientConnInterface) ConversionValueRuleSetServiceClient {
	return &conversionValueRuleSetServiceClient{cc}
}

func (c *conversionValueRuleSetServiceClient) GetConversionValueRuleSet(ctx context.Context, in *GetConversionValueRuleSetRequest, opts ...grpc.CallOption) (*resources.ConversionValueRuleSet, error) {
	out := new(resources.ConversionValueRuleSet)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.ConversionValueRuleSetService/GetConversionValueRuleSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversionValueRuleSetServiceClient) MutateConversionValueRuleSets(ctx context.Context, in *MutateConversionValueRuleSetsRequest, opts ...grpc.CallOption) (*MutateConversionValueRuleSetsResponse, error) {
	out := new(MutateConversionValueRuleSetsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.ConversionValueRuleSetService/MutateConversionValueRuleSets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConversionValueRuleSetServiceServer is the server API for ConversionValueRuleSetService service.
// All implementations must embed UnimplementedConversionValueRuleSetServiceServer
// for forward compatibility
type ConversionValueRuleSetServiceServer interface {
	// Returns the requested conversion value rule set.
	GetConversionValueRuleSet(context.Context, *GetConversionValueRuleSetRequest) (*resources.ConversionValueRuleSet, error)
	// Creates, updates or removes conversion value rule sets. Operation statuses
	// are returned.
	MutateConversionValueRuleSets(context.Context, *MutateConversionValueRuleSetsRequest) (*MutateConversionValueRuleSetsResponse, error)
	mustEmbedUnimplementedConversionValueRuleSetServiceServer()
}

// UnimplementedConversionValueRuleSetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConversionValueRuleSetServiceServer struct {
}

func (UnimplementedConversionValueRuleSetServiceServer) GetConversionValueRuleSet(context.Context, *GetConversionValueRuleSetRequest) (*resources.ConversionValueRuleSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConversionValueRuleSet not implemented")
}
func (UnimplementedConversionValueRuleSetServiceServer) MutateConversionValueRuleSets(context.Context, *MutateConversionValueRuleSetsRequest) (*MutateConversionValueRuleSetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateConversionValueRuleSets not implemented")
}
func (UnimplementedConversionValueRuleSetServiceServer) mustEmbedUnimplementedConversionValueRuleSetServiceServer() {
}

// UnsafeConversionValueRuleSetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConversionValueRuleSetServiceServer will
// result in compilation errors.
type UnsafeConversionValueRuleSetServiceServer interface {
	mustEmbedUnimplementedConversionValueRuleSetServiceServer()
}

func RegisterConversionValueRuleSetServiceServer(s grpc.ServiceRegistrar, srv ConversionValueRuleSetServiceServer) {
	s.RegisterService(&ConversionValueRuleSetService_ServiceDesc, srv)
}

func _ConversionValueRuleSetService_GetConversionValueRuleSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConversionValueRuleSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversionValueRuleSetServiceServer).GetConversionValueRuleSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.ConversionValueRuleSetService/GetConversionValueRuleSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversionValueRuleSetServiceServer).GetConversionValueRuleSet(ctx, req.(*GetConversionValueRuleSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConversionValueRuleSetService_MutateConversionValueRuleSets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateConversionValueRuleSetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversionValueRuleSetServiceServer).MutateConversionValueRuleSets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.ConversionValueRuleSetService/MutateConversionValueRuleSets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversionValueRuleSetServiceServer).MutateConversionValueRuleSets(ctx, req.(*MutateConversionValueRuleSetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConversionValueRuleSetService_ServiceDesc is the grpc.ServiceDesc for ConversionValueRuleSetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConversionValueRuleSetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.ConversionValueRuleSetService",
	HandlerType: (*ConversionValueRuleSetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConversionValueRuleSet",
			Handler:    _ConversionValueRuleSetService_GetConversionValueRuleSet_Handler,
		},
		{
			MethodName: "MutateConversionValueRuleSets",
			Handler:    _ConversionValueRuleSetService_MutateConversionValueRuleSets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/conversion_value_rule_set_service.proto",
}