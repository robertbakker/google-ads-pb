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
// source: google/ads/googleads/v9/errors/feed_item_set_error.proto

package errors

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Enum describing possible feed item set errors.
type FeedItemSetErrorEnum_FeedItemSetError int32

const (
	// Enum unspecified.
	FeedItemSetErrorEnum_UNSPECIFIED FeedItemSetErrorEnum_FeedItemSetError = 0
	// The received error code is not known in this version.
	FeedItemSetErrorEnum_UNKNOWN FeedItemSetErrorEnum_FeedItemSetError = 1
	// The given ID refers to a removed FeedItemSet.
	FeedItemSetErrorEnum_FEED_ITEM_SET_REMOVED FeedItemSetErrorEnum_FeedItemSetError = 2
	// The dynamic filter of a feed item set cannot be cleared on UPDATE if it
	// exists. A set is either static or dynamic once added, and that cannot
	// change.
	FeedItemSetErrorEnum_CANNOT_CLEAR_DYNAMIC_FILTER FeedItemSetErrorEnum_FeedItemSetError = 3
	// The dynamic filter of a feed item set cannot be created on UPDATE if it
	// does not exist. A set is either static or dynamic once added, and that
	// cannot change.
	FeedItemSetErrorEnum_CANNOT_CREATE_DYNAMIC_FILTER FeedItemSetErrorEnum_FeedItemSetError = 4
	// FeedItemSets can only be made for location or affiliate location feeds.
	FeedItemSetErrorEnum_INVALID_FEED_TYPE FeedItemSetErrorEnum_FeedItemSetError = 5
	// FeedItemSets duplicate name. Name should be unique within an account.
	FeedItemSetErrorEnum_DUPLICATE_NAME FeedItemSetErrorEnum_FeedItemSetError = 6
	// The feed type of the parent Feed is not compatible with the type of
	// dynamic filter being set. For example, you can only set
	// dynamic_location_set_filter for LOCATION feed item sets.
	FeedItemSetErrorEnum_WRONG_DYNAMIC_FILTER_FOR_FEED_TYPE FeedItemSetErrorEnum_FeedItemSetError = 7
	// Chain ID specified for AffiliateLocationFeedData is invalid.
	FeedItemSetErrorEnum_DYNAMIC_FILTER_INVALID_CHAIN_IDS FeedItemSetErrorEnum_FeedItemSetError = 8
)

// Enum value maps for FeedItemSetErrorEnum_FeedItemSetError.
var (
	FeedItemSetErrorEnum_FeedItemSetError_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "FEED_ITEM_SET_REMOVED",
		3: "CANNOT_CLEAR_DYNAMIC_FILTER",
		4: "CANNOT_CREATE_DYNAMIC_FILTER",
		5: "INVALID_FEED_TYPE",
		6: "DUPLICATE_NAME",
		7: "WRONG_DYNAMIC_FILTER_FOR_FEED_TYPE",
		8: "DYNAMIC_FILTER_INVALID_CHAIN_IDS",
	}
	FeedItemSetErrorEnum_FeedItemSetError_value = map[string]int32{
		"UNSPECIFIED":                        0,
		"UNKNOWN":                            1,
		"FEED_ITEM_SET_REMOVED":              2,
		"CANNOT_CLEAR_DYNAMIC_FILTER":        3,
		"CANNOT_CREATE_DYNAMIC_FILTER":       4,
		"INVALID_FEED_TYPE":                  5,
		"DUPLICATE_NAME":                     6,
		"WRONG_DYNAMIC_FILTER_FOR_FEED_TYPE": 7,
		"DYNAMIC_FILTER_INVALID_CHAIN_IDS":   8,
	}
)

func (x FeedItemSetErrorEnum_FeedItemSetError) Enum() *FeedItemSetErrorEnum_FeedItemSetError {
	p := new(FeedItemSetErrorEnum_FeedItemSetError)
	*p = x
	return p
}

func (x FeedItemSetErrorEnum_FeedItemSetError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeedItemSetErrorEnum_FeedItemSetError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v9_errors_feed_item_set_error_proto_enumTypes[0].Descriptor()
}

func (FeedItemSetErrorEnum_FeedItemSetError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v9_errors_feed_item_set_error_proto_enumTypes[0]
}

func (x FeedItemSetErrorEnum_FeedItemSetError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeedItemSetErrorEnum_FeedItemSetError.Descriptor instead.
func (FeedItemSetErrorEnum_FeedItemSetError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v9_errors_feed_item_set_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible feed item set errors.
type FeedItemSetErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FeedItemSetErrorEnum) Reset() {
	*x = FeedItemSetErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v9_errors_feed_item_set_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedItemSetErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedItemSetErrorEnum) ProtoMessage() {}

func (x *FeedItemSetErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v9_errors_feed_item_set_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedItemSetErrorEnum.ProtoReflect.Descriptor instead.
func (*FeedItemSetErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v9_errors_feed_item_set_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v9_errors_feed_item_set_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v9_errors_feed_item_set_error_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x39, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x02, 0x0a, 0x14, 0x46, 0x65, 0x65,
	0x64, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75,
	0x6d, 0x22, 0x87, 0x02, 0x0a, 0x10, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x65,
	0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x46, 0x45, 0x45, 0x44, 0x5f, 0x49, 0x54, 0x45,
	0x4d, 0x5f, 0x53, 0x45, 0x54, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x02, 0x12,
	0x1f, 0x0a, 0x1b, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x4c, 0x45, 0x41, 0x52, 0x5f,
	0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x10, 0x03,
	0x12, 0x20, 0x0a, 0x1c, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x5f, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52,
	0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x46, 0x45,
	0x45, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x55, 0x50,
	0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x06, 0x12, 0x26, 0x0a,
	0x22, 0x57, 0x52, 0x4f, 0x4e, 0x47, 0x5f, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x5f, 0x46,
	0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x46, 0x45, 0x45, 0x44, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x10, 0x07, 0x12, 0x24, 0x0a, 0x20, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43,
	0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f,
	0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x49, 0x44, 0x53, 0x10, 0x08, 0x42, 0xf0, 0x01, 0x0a, 0x22,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x42, 0x15, 0x46, 0x65, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x74, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x39, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56,
	0x39, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c,
	0x56, 0x39, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x56, 0x39, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v9_errors_feed_item_set_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v9_errors_feed_item_set_error_proto_rawDescData = file_google_ads_googleads_v9_errors_feed_item_set_error_proto_rawDesc
)

func file_google_ads_googleads_v9_errors_feed_item_set_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v9_errors_feed_item_set_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v9_errors_feed_item_set_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v9_errors_feed_item_set_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v9_errors_feed_item_set_error_proto_rawDescData
}

var file_google_ads_googleads_v9_errors_feed_item_set_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v9_errors_feed_item_set_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v9_errors_feed_item_set_error_proto_goTypes = []interface{}{
	(FeedItemSetErrorEnum_FeedItemSetError)(0), // 0: google.ads.googleads.v9.errors.FeedItemSetErrorEnum.FeedItemSetError
	(*FeedItemSetErrorEnum)(nil),               // 1: google.ads.googleads.v9.errors.FeedItemSetErrorEnum
}
var file_google_ads_googleads_v9_errors_feed_item_set_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v9_errors_feed_item_set_error_proto_init() }
func file_google_ads_googleads_v9_errors_feed_item_set_error_proto_init() {
	if File_google_ads_googleads_v9_errors_feed_item_set_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v9_errors_feed_item_set_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedItemSetErrorEnum); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v9_errors_feed_item_set_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v9_errors_feed_item_set_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v9_errors_feed_item_set_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v9_errors_feed_item_set_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v9_errors_feed_item_set_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v9_errors_feed_item_set_error_proto = out.File
	file_google_ads_googleads_v9_errors_feed_item_set_error_proto_rawDesc = nil
	file_google_ads_googleads_v9_errors_feed_item_set_error_proto_goTypes = nil
	file_google_ads_googleads_v9_errors_feed_item_set_error_proto_depIdxs = nil
}
