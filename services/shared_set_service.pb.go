// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.3
// source: google/ads/googleads/v9/services/shared_set_service.proto

package services

import (
	enums "github.com/shenzhencenter/google-ads-pb/enums"
	resources "github.com/shenzhencenter/google-ads-pb/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for [SharedSetService.GetSharedSet][google.ads.googleads.v9.services.SharedSetService.GetSharedSet].
type GetSharedSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the shared set to fetch.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *GetSharedSetRequest) Reset() {
	*x = GetSharedSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v9_services_shared_set_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSharedSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSharedSetRequest) ProtoMessage() {}

func (x *GetSharedSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v9_services_shared_set_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSharedSetRequest.ProtoReflect.Descriptor instead.
func (*GetSharedSetRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v9_services_shared_set_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetSharedSetRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

// Request message for [SharedSetService.MutateSharedSets][google.ads.googleads.v9.services.SharedSetService.MutateSharedSets].
type MutateSharedSetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The ID of the customer whose shared sets are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual shared sets.
	Operations []*SharedSetOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	// The response content type setting. Determines whether the mutable resource
	// or just the resource name should be returned post mutation.
	ResponseContentType enums.ResponseContentTypeEnum_ResponseContentType `protobuf:"varint,5,opt,name=response_content_type,json=responseContentType,proto3,enum=google.ads.googleads.v9.enums.ResponseContentTypeEnum_ResponseContentType" json:"response_content_type,omitempty"`
}

func (x *MutateSharedSetsRequest) Reset() {
	*x = MutateSharedSetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v9_services_shared_set_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateSharedSetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateSharedSetsRequest) ProtoMessage() {}

func (x *MutateSharedSetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v9_services_shared_set_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateSharedSetsRequest.ProtoReflect.Descriptor instead.
func (*MutateSharedSetsRequest) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v9_services_shared_set_service_proto_rawDescGZIP(), []int{1}
}

func (x *MutateSharedSetsRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateSharedSetsRequest) GetOperations() []*SharedSetOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *MutateSharedSetsRequest) GetPartialFailure() bool {
	if x != nil {
		return x.PartialFailure
	}
	return false
}

func (x *MutateSharedSetsRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

func (x *MutateSharedSetsRequest) GetResponseContentType() enums.ResponseContentTypeEnum_ResponseContentType {
	if x != nil {
		return x.ResponseContentType
	}
	return enums.ResponseContentTypeEnum_UNSPECIFIED
}

// A single operation (create, update, remove) on an shared set.
type SharedSetOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are assignable to Operation:
	//	*SharedSetOperation_Create
	//	*SharedSetOperation_Update
	//	*SharedSetOperation_Remove
	Operation isSharedSetOperation_Operation `protobuf_oneof:"operation"`
}

func (x *SharedSetOperation) Reset() {
	*x = SharedSetOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v9_services_shared_set_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SharedSetOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SharedSetOperation) ProtoMessage() {}

func (x *SharedSetOperation) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v9_services_shared_set_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SharedSetOperation.ProtoReflect.Descriptor instead.
func (*SharedSetOperation) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v9_services_shared_set_service_proto_rawDescGZIP(), []int{2}
}

func (x *SharedSetOperation) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (m *SharedSetOperation) GetOperation() isSharedSetOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *SharedSetOperation) GetCreate() *resources.SharedSet {
	if x, ok := x.GetOperation().(*SharedSetOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (x *SharedSetOperation) GetUpdate() *resources.SharedSet {
	if x, ok := x.GetOperation().(*SharedSetOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (x *SharedSetOperation) GetRemove() string {
	if x, ok := x.GetOperation().(*SharedSetOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

type isSharedSetOperation_Operation interface {
	isSharedSetOperation_Operation()
}

type SharedSetOperation_Create struct {
	// Create operation: No resource name is expected for the new shared set.
	Create *resources.SharedSet `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type SharedSetOperation_Update struct {
	// Update operation: The shared set is expected to have a valid resource
	// name.
	Update *resources.SharedSet `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type SharedSetOperation_Remove struct {
	// Remove operation: A resource name for the removed shared set is expected,
	// in this format:
	//
	// `customers/{customer_id}/sharedSets/{shared_set_id}`
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*SharedSetOperation_Create) isSharedSetOperation_Operation() {}

func (*SharedSetOperation_Update) isSharedSetOperation_Operation() {}

func (*SharedSetOperation_Remove) isSharedSetOperation_Operation() {}

// Response message for a shared set mutate.
type MutateSharedSetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results []*MutateSharedSetResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *MutateSharedSetsResponse) Reset() {
	*x = MutateSharedSetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v9_services_shared_set_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateSharedSetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateSharedSetsResponse) ProtoMessage() {}

func (x *MutateSharedSetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v9_services_shared_set_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateSharedSetsResponse.ProtoReflect.Descriptor instead.
func (*MutateSharedSetsResponse) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v9_services_shared_set_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateSharedSetsResponse) GetPartialFailureError() *status.Status {
	if x != nil {
		return x.PartialFailureError
	}
	return nil
}

func (x *MutateSharedSetsResponse) GetResults() []*MutateSharedSetResult {
	if x != nil {
		return x.Results
	}
	return nil
}

// The result for the shared set mutate.
type MutateSharedSetResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The mutated shared set with only mutable fields after mutate. The field
	// will only be returned when response_content_type is set to
	// "MUTABLE_RESOURCE".
	SharedSet *resources.SharedSet `protobuf:"bytes,2,opt,name=shared_set,json=sharedSet,proto3" json:"shared_set,omitempty"`
}

func (x *MutateSharedSetResult) Reset() {
	*x = MutateSharedSetResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v9_services_shared_set_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateSharedSetResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateSharedSetResult) ProtoMessage() {}

func (x *MutateSharedSetResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v9_services_shared_set_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateSharedSetResult.ProtoReflect.Descriptor instead.
func (*MutateSharedSetResult) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v9_services_shared_set_service_proto_rawDescGZIP(), []int{4}
}

func (x *MutateSharedSetResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *MutateSharedSetResult) GetSharedSet() *resources.SharedSet {
	if x != nil {
		return x.SharedSet
	}
	return nil
}

var File_google_ads_googleads_v9_services_shared_set_service_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v9_services_shared_set_service_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x39, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x39, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x39, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x5f, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x4f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2a, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x24,
	0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x53, 0x65, 0x74, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0xe8, 0x02, 0x0a, 0x17, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24,
	0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x59, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x39, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x53, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x27, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75,
	0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61,
	0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x7e, 0x0a,
	0x15, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x13, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x88, 0x02,
	0x0a, 0x12, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d,
	0x61, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73,
	0x6b, 0x12, 0x46, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x48,
	0x00, 0x52, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x46, 0x0a, 0x06, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x39, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x48, 0x00, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x18, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb5, 0x01, 0x0a, 0x18, 0x4d, 0x75, 0x74,
	0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x15, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c,
	0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x13, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61,
	0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x51, 0x0a,
	0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x22, 0x89, 0x01, 0x0a, 0x15, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x4b, 0x0a, 0x0a, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65,
	0x74, 0x52, 0x09, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x32, 0xf6, 0x03, 0x0a,
	0x10, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xb9, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53,
	0x65, 0x74, 0x12, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x39, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x22, 0x44, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x12,
	0x2c, 0x2f, 0x76, 0x39, 0x2f, 0x7b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x3d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0xda, 0x41, 0x0d,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0xde, 0x01,
	0x0a, 0x10, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65,
	0x74, 0x73, 0x12, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x53, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x53, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x34, 0x22, 0x2f, 0x2f, 0x76, 0x39, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73,
	0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d, 0x2a, 0x7d,
	0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x73, 0x3a, 0x6d, 0x75, 0x74, 0x61,
	0x74, 0x65, 0x3a, 0x01, 0x2a, 0xda, 0x41, 0x16, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x2c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x45,
	0xca, 0x41, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x27, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x64,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0xfc, 0x01, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x15,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56,
	0x39, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x20, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x5c, 0x56, 0x39, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x24,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x39, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v9_services_shared_set_service_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v9_services_shared_set_service_proto_rawDescData = file_google_ads_googleads_v9_services_shared_set_service_proto_rawDesc
)

func file_google_ads_googleads_v9_services_shared_set_service_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v9_services_shared_set_service_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v9_services_shared_set_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v9_services_shared_set_service_proto_rawDescData)
	})
	return file_google_ads_googleads_v9_services_shared_set_service_proto_rawDescData
}

var file_google_ads_googleads_v9_services_shared_set_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_ads_googleads_v9_services_shared_set_service_proto_goTypes = []interface{}{
	(*GetSharedSetRequest)(nil),                            // 0: google.ads.googleads.v9.services.GetSharedSetRequest
	(*MutateSharedSetsRequest)(nil),                        // 1: google.ads.googleads.v9.services.MutateSharedSetsRequest
	(*SharedSetOperation)(nil),                             // 2: google.ads.googleads.v9.services.SharedSetOperation
	(*MutateSharedSetsResponse)(nil),                       // 3: google.ads.googleads.v9.services.MutateSharedSetsResponse
	(*MutateSharedSetResult)(nil),                          // 4: google.ads.googleads.v9.services.MutateSharedSetResult
	(enums.ResponseContentTypeEnum_ResponseContentType)(0), // 5: google.ads.googleads.v9.enums.ResponseContentTypeEnum.ResponseContentType
	(*fieldmaskpb.FieldMask)(nil),                          // 6: google.protobuf.FieldMask
	(*resources.SharedSet)(nil),                            // 7: google.ads.googleads.v9.resources.SharedSet
	(*status.Status)(nil),                                  // 8: google.rpc.Status
}
var file_google_ads_googleads_v9_services_shared_set_service_proto_depIdxs = []int32{
	2,  // 0: google.ads.googleads.v9.services.MutateSharedSetsRequest.operations:type_name -> google.ads.googleads.v9.services.SharedSetOperation
	5,  // 1: google.ads.googleads.v9.services.MutateSharedSetsRequest.response_content_type:type_name -> google.ads.googleads.v9.enums.ResponseContentTypeEnum.ResponseContentType
	6,  // 2: google.ads.googleads.v9.services.SharedSetOperation.update_mask:type_name -> google.protobuf.FieldMask
	7,  // 3: google.ads.googleads.v9.services.SharedSetOperation.create:type_name -> google.ads.googleads.v9.resources.SharedSet
	7,  // 4: google.ads.googleads.v9.services.SharedSetOperation.update:type_name -> google.ads.googleads.v9.resources.SharedSet
	8,  // 5: google.ads.googleads.v9.services.MutateSharedSetsResponse.partial_failure_error:type_name -> google.rpc.Status
	4,  // 6: google.ads.googleads.v9.services.MutateSharedSetsResponse.results:type_name -> google.ads.googleads.v9.services.MutateSharedSetResult
	7,  // 7: google.ads.googleads.v9.services.MutateSharedSetResult.shared_set:type_name -> google.ads.googleads.v9.resources.SharedSet
	0,  // 8: google.ads.googleads.v9.services.SharedSetService.GetSharedSet:input_type -> google.ads.googleads.v9.services.GetSharedSetRequest
	1,  // 9: google.ads.googleads.v9.services.SharedSetService.MutateSharedSets:input_type -> google.ads.googleads.v9.services.MutateSharedSetsRequest
	7,  // 10: google.ads.googleads.v9.services.SharedSetService.GetSharedSet:output_type -> google.ads.googleads.v9.resources.SharedSet
	3,  // 11: google.ads.googleads.v9.services.SharedSetService.MutateSharedSets:output_type -> google.ads.googleads.v9.services.MutateSharedSetsResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v9_services_shared_set_service_proto_init() }
func file_google_ads_googleads_v9_services_shared_set_service_proto_init() {
	if File_google_ads_googleads_v9_services_shared_set_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v9_services_shared_set_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSharedSetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_ads_googleads_v9_services_shared_set_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateSharedSetsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_ads_googleads_v9_services_shared_set_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SharedSetOperation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_ads_googleads_v9_services_shared_set_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateSharedSetsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_ads_googleads_v9_services_shared_set_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateSharedSetResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_google_ads_googleads_v9_services_shared_set_service_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*SharedSetOperation_Create)(nil),
		(*SharedSetOperation_Update)(nil),
		(*SharedSetOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v9_services_shared_set_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_ads_googleads_v9_services_shared_set_service_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v9_services_shared_set_service_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v9_services_shared_set_service_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v9_services_shared_set_service_proto = out.File
	file_google_ads_googleads_v9_services_shared_set_service_proto_rawDesc = nil
	file_google_ads_googleads_v9_services_shared_set_service_proto_goTypes = nil
	file_google_ads_googleads_v9_services_shared_set_service_proto_depIdxs = nil
}
