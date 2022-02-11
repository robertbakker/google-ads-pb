// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package clients

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	resourcespb "github.com/shenzhencenter/google-ads-pb/resources"
	servicespb "github.com/shenzhencenter/google-ads-pb/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newCustomInterestClientHook clientHook

// CustomInterestCallOptions contains the retry settings for each method of CustomInterestClient.
type CustomInterestCallOptions struct {
	GetCustomInterest []gax.CallOption
	MutateCustomInterests []gax.CallOption
}

func defaultCustomInterestGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("googleads.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("googleads.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://googleads.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCustomInterestCallOptions() *CustomInterestCallOptions {
	return &CustomInterestCallOptions{
		GetCustomInterest: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		MutateCustomInterests: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalCustomInterestClient is an interface that defines the methods availaible from Google Ads API.
type internalCustomInterestClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetCustomInterest(context.Context, *servicespb.GetCustomInterestRequest, ...gax.CallOption) (*resourcespb.CustomInterest, error)
	MutateCustomInterests(context.Context, *servicespb.MutateCustomInterestsRequest, ...gax.CallOption) (*servicespb.MutateCustomInterestsResponse, error)
}

// CustomInterestClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage custom interests.
type CustomInterestClient struct {
	// The internal transport-dependent client.
	internalClient internalCustomInterestClient

	// The call options for this service.
	CallOptions *CustomInterestCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CustomInterestClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CustomInterestClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *CustomInterestClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetCustomInterest returns the requested custom interest in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *CustomInterestClient) GetCustomInterest(ctx context.Context, req *servicespb.GetCustomInterestRequest, opts ...gax.CallOption) (*resourcespb.CustomInterest, error) {
	return c.internalClient.GetCustomInterest(ctx, req, opts...)
}

// MutateCustomInterests creates or updates custom interests. Operation statuses are returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// CriterionError (at )
// CustomInterestError (at )
// HeaderError (at )
// InternalError (at )
// MutateError (at )
// PolicyViolationError (at )
// QuotaError (at )
// RequestError (at )
// StringLengthError (at )
func (c *CustomInterestClient) MutateCustomInterests(ctx context.Context, req *servicespb.MutateCustomInterestsRequest, opts ...gax.CallOption) (*servicespb.MutateCustomInterestsResponse, error) {
	return c.internalClient.MutateCustomInterests(ctx, req, opts...)
}

// customInterestGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type customInterestGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing CustomInterestClient
	CallOptions **CustomInterestCallOptions

	// The gRPC API client.
	customInterestClient servicespb.CustomInterestServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewCustomInterestClient creates a new custom interest service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage custom interests.
func NewCustomInterestClient(ctx context.Context, opts ...option.ClientOption) (*CustomInterestClient, error) {
	clientOpts := defaultCustomInterestGRPCClientOptions()
	if newCustomInterestClientHook != nil {
		hookOpts, err := newCustomInterestClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := CustomInterestClient{CallOptions: defaultCustomInterestCallOptions()}

	c := &customInterestGRPCClient{
		connPool:    connPool,
		disableDeadlines: disableDeadlines,
		customInterestClient: servicespb.NewCustomInterestServiceClient(connPool),
		CallOptions: &client.CallOptions,

	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *customInterestGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *customInterestGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *customInterestGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *customInterestGRPCClient) GetCustomInterest(ctx context.Context, req *servicespb.GetCustomInterestRequest, opts ...gax.CallOption) (*resourcespb.CustomInterest, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000 * time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetCustomInterest[0:len((*c.CallOptions).GetCustomInterest):len((*c.CallOptions).GetCustomInterest)], opts...)
	var resp *resourcespb.CustomInterest
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.customInterestClient.GetCustomInterest(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *customInterestGRPCClient) MutateCustomInterests(ctx context.Context, req *servicespb.MutateCustomInterestsRequest, opts ...gax.CallOption) (*servicespb.MutateCustomInterestsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000 * time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateCustomInterests[0:len((*c.CallOptions).MutateCustomInterests):len((*c.CallOptions).MutateCustomInterests)], opts...)
	var resp *servicespb.MutateCustomInterestsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.customInterestClient.MutateCustomInterests(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
