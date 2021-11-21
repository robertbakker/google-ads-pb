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

// KeywordPlanServiceClient is the client API for KeywordPlanService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeywordPlanServiceClient interface {
	// Returns the requested plan in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetKeywordPlan(ctx context.Context, in *GetKeywordPlanRequest, opts ...grpc.CallOption) (*resources.KeywordPlan, error)
	// Creates, updates, or removes keyword plans. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [KeywordPlanError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [QuotaError]()
	//   [RequestError]()
	//   [ResourceCountLimitExceededError]()
	//   [StringLengthError]()
	MutateKeywordPlans(ctx context.Context, in *MutateKeywordPlansRequest, opts ...grpc.CallOption) (*MutateKeywordPlansResponse, error)
	// Returns the requested Keyword Plan forecast curve.
	// Only the bidding strategy is considered for generating forecast curve.
	// The bidding strategy value specified in the plan is ignored.
	//
	// To generate a forecast at a value specified in the plan, use
	// KeywordPlanService.GenerateForecastMetrics.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [KeywordPlanError]()
	//   [QuotaError]()
	//   [RequestError]()
	GenerateForecastCurve(ctx context.Context, in *GenerateForecastCurveRequest, opts ...grpc.CallOption) (*GenerateForecastCurveResponse, error)
	// Returns a forecast in the form of a time series for the Keyword Plan over
	// the next 52 weeks.
	// (1) Forecasts closer to the current date are generally more accurate than
	// further out.
	//
	// (2) The forecast reflects seasonal trends using current and
	// prior traffic patterns. The forecast period of the plan is ignored.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [KeywordPlanError]()
	//   [QuotaError]()
	//   [RequestError]()
	GenerateForecastTimeSeries(ctx context.Context, in *GenerateForecastTimeSeriesRequest, opts ...grpc.CallOption) (*GenerateForecastTimeSeriesResponse, error)
	// Returns the requested Keyword Plan forecasts.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [KeywordPlanError]()
	//   [QuotaError]()
	//   [RequestError]()
	GenerateForecastMetrics(ctx context.Context, in *GenerateForecastMetricsRequest, opts ...grpc.CallOption) (*GenerateForecastMetricsResponse, error)
	// Returns the requested Keyword Plan historical metrics.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [KeywordPlanError]()
	//   [QuotaError]()
	//   [RequestError]()
	GenerateHistoricalMetrics(ctx context.Context, in *GenerateHistoricalMetricsRequest, opts ...grpc.CallOption) (*GenerateHistoricalMetricsResponse, error)
}

type keywordPlanServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeywordPlanServiceClient(cc grpc.ClientConnInterface) KeywordPlanServiceClient {
	return &keywordPlanServiceClient{cc}
}

func (c *keywordPlanServiceClient) GetKeywordPlan(ctx context.Context, in *GetKeywordPlanRequest, opts ...grpc.CallOption) (*resources.KeywordPlan, error) {
	out := new(resources.KeywordPlan)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.KeywordPlanService/GetKeywordPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keywordPlanServiceClient) MutateKeywordPlans(ctx context.Context, in *MutateKeywordPlansRequest, opts ...grpc.CallOption) (*MutateKeywordPlansResponse, error) {
	out := new(MutateKeywordPlansResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.KeywordPlanService/MutateKeywordPlans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keywordPlanServiceClient) GenerateForecastCurve(ctx context.Context, in *GenerateForecastCurveRequest, opts ...grpc.CallOption) (*GenerateForecastCurveResponse, error) {
	out := new(GenerateForecastCurveResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.KeywordPlanService/GenerateForecastCurve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keywordPlanServiceClient) GenerateForecastTimeSeries(ctx context.Context, in *GenerateForecastTimeSeriesRequest, opts ...grpc.CallOption) (*GenerateForecastTimeSeriesResponse, error) {
	out := new(GenerateForecastTimeSeriesResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.KeywordPlanService/GenerateForecastTimeSeries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keywordPlanServiceClient) GenerateForecastMetrics(ctx context.Context, in *GenerateForecastMetricsRequest, opts ...grpc.CallOption) (*GenerateForecastMetricsResponse, error) {
	out := new(GenerateForecastMetricsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.KeywordPlanService/GenerateForecastMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keywordPlanServiceClient) GenerateHistoricalMetrics(ctx context.Context, in *GenerateHistoricalMetricsRequest, opts ...grpc.CallOption) (*GenerateHistoricalMetricsResponse, error) {
	out := new(GenerateHistoricalMetricsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.KeywordPlanService/GenerateHistoricalMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordPlanServiceServer is the server API for KeywordPlanService service.
// All implementations must embed UnimplementedKeywordPlanServiceServer
// for forward compatibility
type KeywordPlanServiceServer interface {
	// Returns the requested plan in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetKeywordPlan(context.Context, *GetKeywordPlanRequest) (*resources.KeywordPlan, error)
	// Creates, updates, or removes keyword plans. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [FieldError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [KeywordPlanError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [QuotaError]()
	//   [RequestError]()
	//   [ResourceCountLimitExceededError]()
	//   [StringLengthError]()
	MutateKeywordPlans(context.Context, *MutateKeywordPlansRequest) (*MutateKeywordPlansResponse, error)
	// Returns the requested Keyword Plan forecast curve.
	// Only the bidding strategy is considered for generating forecast curve.
	// The bidding strategy value specified in the plan is ignored.
	//
	// To generate a forecast at a value specified in the plan, use
	// KeywordPlanService.GenerateForecastMetrics.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [KeywordPlanError]()
	//   [QuotaError]()
	//   [RequestError]()
	GenerateForecastCurve(context.Context, *GenerateForecastCurveRequest) (*GenerateForecastCurveResponse, error)
	// Returns a forecast in the form of a time series for the Keyword Plan over
	// the next 52 weeks.
	// (1) Forecasts closer to the current date are generally more accurate than
	// further out.
	//
	// (2) The forecast reflects seasonal trends using current and
	// prior traffic patterns. The forecast period of the plan is ignored.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [KeywordPlanError]()
	//   [QuotaError]()
	//   [RequestError]()
	GenerateForecastTimeSeries(context.Context, *GenerateForecastTimeSeriesRequest) (*GenerateForecastTimeSeriesResponse, error)
	// Returns the requested Keyword Plan forecasts.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [KeywordPlanError]()
	//   [QuotaError]()
	//   [RequestError]()
	GenerateForecastMetrics(context.Context, *GenerateForecastMetricsRequest) (*GenerateForecastMetricsResponse, error)
	// Returns the requested Keyword Plan historical metrics.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [KeywordPlanError]()
	//   [QuotaError]()
	//   [RequestError]()
	GenerateHistoricalMetrics(context.Context, *GenerateHistoricalMetricsRequest) (*GenerateHistoricalMetricsResponse, error)
	mustEmbedUnimplementedKeywordPlanServiceServer()
}

// UnimplementedKeywordPlanServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKeywordPlanServiceServer struct {
}

func (UnimplementedKeywordPlanServiceServer) GetKeywordPlan(context.Context, *GetKeywordPlanRequest) (*resources.KeywordPlan, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKeywordPlan not implemented")
}
func (UnimplementedKeywordPlanServiceServer) MutateKeywordPlans(context.Context, *MutateKeywordPlansRequest) (*MutateKeywordPlansResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateKeywordPlans not implemented")
}
func (UnimplementedKeywordPlanServiceServer) GenerateForecastCurve(context.Context, *GenerateForecastCurveRequest) (*GenerateForecastCurveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateForecastCurve not implemented")
}
func (UnimplementedKeywordPlanServiceServer) GenerateForecastTimeSeries(context.Context, *GenerateForecastTimeSeriesRequest) (*GenerateForecastTimeSeriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateForecastTimeSeries not implemented")
}
func (UnimplementedKeywordPlanServiceServer) GenerateForecastMetrics(context.Context, *GenerateForecastMetricsRequest) (*GenerateForecastMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateForecastMetrics not implemented")
}
func (UnimplementedKeywordPlanServiceServer) GenerateHistoricalMetrics(context.Context, *GenerateHistoricalMetricsRequest) (*GenerateHistoricalMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateHistoricalMetrics not implemented")
}
func (UnimplementedKeywordPlanServiceServer) mustEmbedUnimplementedKeywordPlanServiceServer() {}

// UnsafeKeywordPlanServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeywordPlanServiceServer will
// result in compilation errors.
type UnsafeKeywordPlanServiceServer interface {
	mustEmbedUnimplementedKeywordPlanServiceServer()
}

func RegisterKeywordPlanServiceServer(s grpc.ServiceRegistrar, srv KeywordPlanServiceServer) {
	s.RegisterService(&KeywordPlanService_ServiceDesc, srv)
}

func _KeywordPlanService_GetKeywordPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeywordPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanServiceServer).GetKeywordPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.KeywordPlanService/GetKeywordPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanServiceServer).GetKeywordPlan(ctx, req.(*GetKeywordPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeywordPlanService_MutateKeywordPlans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateKeywordPlansRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanServiceServer).MutateKeywordPlans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.KeywordPlanService/MutateKeywordPlans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanServiceServer).MutateKeywordPlans(ctx, req.(*MutateKeywordPlansRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeywordPlanService_GenerateForecastCurve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateForecastCurveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanServiceServer).GenerateForecastCurve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.KeywordPlanService/GenerateForecastCurve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanServiceServer).GenerateForecastCurve(ctx, req.(*GenerateForecastCurveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeywordPlanService_GenerateForecastTimeSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateForecastTimeSeriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanServiceServer).GenerateForecastTimeSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.KeywordPlanService/GenerateForecastTimeSeries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanServiceServer).GenerateForecastTimeSeries(ctx, req.(*GenerateForecastTimeSeriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeywordPlanService_GenerateForecastMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateForecastMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanServiceServer).GenerateForecastMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.KeywordPlanService/GenerateForecastMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanServiceServer).GenerateForecastMetrics(ctx, req.(*GenerateForecastMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeywordPlanService_GenerateHistoricalMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateHistoricalMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanServiceServer).GenerateHistoricalMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.KeywordPlanService/GenerateHistoricalMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanServiceServer).GenerateHistoricalMetrics(ctx, req.(*GenerateHistoricalMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeywordPlanService_ServiceDesc is the grpc.ServiceDesc for KeywordPlanService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeywordPlanService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.KeywordPlanService",
	HandlerType: (*KeywordPlanServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKeywordPlan",
			Handler:    _KeywordPlanService_GetKeywordPlan_Handler,
		},
		{
			MethodName: "MutateKeywordPlans",
			Handler:    _KeywordPlanService_MutateKeywordPlans_Handler,
		},
		{
			MethodName: "GenerateForecastCurve",
			Handler:    _KeywordPlanService_GenerateForecastCurve_Handler,
		},
		{
			MethodName: "GenerateForecastTimeSeries",
			Handler:    _KeywordPlanService_GenerateForecastTimeSeries_Handler,
		},
		{
			MethodName: "GenerateForecastMetrics",
			Handler:    _KeywordPlanService_GenerateForecastMetrics_Handler,
		},
		{
			MethodName: "GenerateHistoricalMetrics",
			Handler:    _KeywordPlanService_GenerateHistoricalMetrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/keyword_plan_service.proto",
}
