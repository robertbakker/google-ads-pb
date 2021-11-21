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
// source: google/ads/googleads/v9/enums/price_placeholder_field.proto

package enums

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

// Possible values for Price placeholder fields.
type PricePlaceholderFieldEnum_PricePlaceholderField int32

const (
	// Not specified.
	PricePlaceholderFieldEnum_UNSPECIFIED PricePlaceholderFieldEnum_PricePlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	PricePlaceholderFieldEnum_UNKNOWN PricePlaceholderFieldEnum_PricePlaceholderField = 1
	// Data Type: STRING. The type of your price feed. Must match one of the
	// predefined price feed type exactly.
	PricePlaceholderFieldEnum_TYPE PricePlaceholderFieldEnum_PricePlaceholderField = 2
	// Data Type: STRING. The qualifier of each price. Must match one of the
	// predefined price qualifiers exactly.
	PricePlaceholderFieldEnum_PRICE_QUALIFIER PricePlaceholderFieldEnum_PricePlaceholderField = 3
	// Data Type: URL. Tracking template for the price feed when using Upgraded
	// URLs.
	PricePlaceholderFieldEnum_TRACKING_TEMPLATE PricePlaceholderFieldEnum_PricePlaceholderField = 4
	// Data Type: STRING. Language of the price feed. Must match one of the
	// available available locale codes exactly.
	PricePlaceholderFieldEnum_LANGUAGE PricePlaceholderFieldEnum_PricePlaceholderField = 5
	// Data Type: STRING. Final URL suffix for the price feed when using
	// parallel tracking.
	PricePlaceholderFieldEnum_FINAL_URL_SUFFIX PricePlaceholderFieldEnum_PricePlaceholderField = 6
	// Data Type: STRING. The header of item 1 of the table.
	PricePlaceholderFieldEnum_ITEM_1_HEADER PricePlaceholderFieldEnum_PricePlaceholderField = 100
	// Data Type: STRING. The description of item 1 of the table.
	PricePlaceholderFieldEnum_ITEM_1_DESCRIPTION PricePlaceholderFieldEnum_PricePlaceholderField = 101
	// Data Type: MONEY. The price (money with currency) of item 1 of the table,
	// e.g., 30 USD. The currency must match one of the available currencies.
	PricePlaceholderFieldEnum_ITEM_1_PRICE PricePlaceholderFieldEnum_PricePlaceholderField = 102
	// Data Type: STRING. The price unit of item 1 of the table. Must match one
	// of the predefined price units.
	PricePlaceholderFieldEnum_ITEM_1_UNIT PricePlaceholderFieldEnum_PricePlaceholderField = 103
	// Data Type: URL_LIST. The final URLs of item 1 of the table when using
	// Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_1_FINAL_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 104
	// Data Type: URL_LIST. The final mobile URLs of item 1 of the table when
	// using Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_1_FINAL_MOBILE_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 105
	// Data Type: STRING. The header of item 2 of the table.
	PricePlaceholderFieldEnum_ITEM_2_HEADER PricePlaceholderFieldEnum_PricePlaceholderField = 200
	// Data Type: STRING. The description of item 2 of the table.
	PricePlaceholderFieldEnum_ITEM_2_DESCRIPTION PricePlaceholderFieldEnum_PricePlaceholderField = 201
	// Data Type: MONEY. The price (money with currency) of item 2 of the table,
	// e.g., 30 USD. The currency must match one of the available currencies.
	PricePlaceholderFieldEnum_ITEM_2_PRICE PricePlaceholderFieldEnum_PricePlaceholderField = 202
	// Data Type: STRING. The price unit of item 2 of the table. Must match one
	// of the predefined price units.
	PricePlaceholderFieldEnum_ITEM_2_UNIT PricePlaceholderFieldEnum_PricePlaceholderField = 203
	// Data Type: URL_LIST. The final URLs of item 2 of the table when using
	// Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_2_FINAL_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 204
	// Data Type: URL_LIST. The final mobile URLs of item 2 of the table when
	// using Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_2_FINAL_MOBILE_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 205
	// Data Type: STRING. The header of item 3 of the table.
	PricePlaceholderFieldEnum_ITEM_3_HEADER PricePlaceholderFieldEnum_PricePlaceholderField = 300
	// Data Type: STRING. The description of item 3 of the table.
	PricePlaceholderFieldEnum_ITEM_3_DESCRIPTION PricePlaceholderFieldEnum_PricePlaceholderField = 301
	// Data Type: MONEY. The price (money with currency) of item 3 of the table,
	// e.g., 30 USD. The currency must match one of the available currencies.
	PricePlaceholderFieldEnum_ITEM_3_PRICE PricePlaceholderFieldEnum_PricePlaceholderField = 302
	// Data Type: STRING. The price unit of item 3 of the table. Must match one
	// of the predefined price units.
	PricePlaceholderFieldEnum_ITEM_3_UNIT PricePlaceholderFieldEnum_PricePlaceholderField = 303
	// Data Type: URL_LIST. The final URLs of item 3 of the table when using
	// Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_3_FINAL_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 304
	// Data Type: URL_LIST. The final mobile URLs of item 3 of the table when
	// using Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_3_FINAL_MOBILE_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 305
	// Data Type: STRING. The header of item 4 of the table.
	PricePlaceholderFieldEnum_ITEM_4_HEADER PricePlaceholderFieldEnum_PricePlaceholderField = 400
	// Data Type: STRING. The description of item 4 of the table.
	PricePlaceholderFieldEnum_ITEM_4_DESCRIPTION PricePlaceholderFieldEnum_PricePlaceholderField = 401
	// Data Type: MONEY. The price (money with currency) of item 4 of the table,
	// e.g., 30 USD. The currency must match one of the available currencies.
	PricePlaceholderFieldEnum_ITEM_4_PRICE PricePlaceholderFieldEnum_PricePlaceholderField = 402
	// Data Type: STRING. The price unit of item 4 of the table. Must match one
	// of the predefined price units.
	PricePlaceholderFieldEnum_ITEM_4_UNIT PricePlaceholderFieldEnum_PricePlaceholderField = 403
	// Data Type: URL_LIST. The final URLs of item 4 of the table when using
	// Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_4_FINAL_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 404
	// Data Type: URL_LIST. The final mobile URLs of item 4 of the table when
	// using Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_4_FINAL_MOBILE_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 405
	// Data Type: STRING. The header of item 5 of the table.
	PricePlaceholderFieldEnum_ITEM_5_HEADER PricePlaceholderFieldEnum_PricePlaceholderField = 500
	// Data Type: STRING. The description of item 5 of the table.
	PricePlaceholderFieldEnum_ITEM_5_DESCRIPTION PricePlaceholderFieldEnum_PricePlaceholderField = 501
	// Data Type: MONEY. The price (money with currency) of item 5 of the table,
	// e.g., 30 USD. The currency must match one of the available currencies.
	PricePlaceholderFieldEnum_ITEM_5_PRICE PricePlaceholderFieldEnum_PricePlaceholderField = 502
	// Data Type: STRING. The price unit of item 5 of the table. Must match one
	// of the predefined price units.
	PricePlaceholderFieldEnum_ITEM_5_UNIT PricePlaceholderFieldEnum_PricePlaceholderField = 503
	// Data Type: URL_LIST. The final URLs of item 5 of the table when using
	// Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_5_FINAL_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 504
	// Data Type: URL_LIST. The final mobile URLs of item 5 of the table when
	// using Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_5_FINAL_MOBILE_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 505
	// Data Type: STRING. The header of item 6 of the table.
	PricePlaceholderFieldEnum_ITEM_6_HEADER PricePlaceholderFieldEnum_PricePlaceholderField = 600
	// Data Type: STRING. The description of item 6 of the table.
	PricePlaceholderFieldEnum_ITEM_6_DESCRIPTION PricePlaceholderFieldEnum_PricePlaceholderField = 601
	// Data Type: MONEY. The price (money with currency) of item 6 of the table,
	// e.g., 30 USD. The currency must match one of the available currencies.
	PricePlaceholderFieldEnum_ITEM_6_PRICE PricePlaceholderFieldEnum_PricePlaceholderField = 602
	// Data Type: STRING. The price unit of item 6 of the table. Must match one
	// of the predefined price units.
	PricePlaceholderFieldEnum_ITEM_6_UNIT PricePlaceholderFieldEnum_PricePlaceholderField = 603
	// Data Type: URL_LIST. The final URLs of item 6 of the table when using
	// Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_6_FINAL_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 604
	// Data Type: URL_LIST. The final mobile URLs of item 6 of the table when
	// using Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_6_FINAL_MOBILE_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 605
	// Data Type: STRING. The header of item 7 of the table.
	PricePlaceholderFieldEnum_ITEM_7_HEADER PricePlaceholderFieldEnum_PricePlaceholderField = 700
	// Data Type: STRING. The description of item 7 of the table.
	PricePlaceholderFieldEnum_ITEM_7_DESCRIPTION PricePlaceholderFieldEnum_PricePlaceholderField = 701
	// Data Type: MONEY. The price (money with currency) of item 7 of the table,
	// e.g., 30 USD. The currency must match one of the available currencies.
	PricePlaceholderFieldEnum_ITEM_7_PRICE PricePlaceholderFieldEnum_PricePlaceholderField = 702
	// Data Type: STRING. The price unit of item 7 of the table. Must match one
	// of the predefined price units.
	PricePlaceholderFieldEnum_ITEM_7_UNIT PricePlaceholderFieldEnum_PricePlaceholderField = 703
	// Data Type: URL_LIST. The final URLs of item 7 of the table when using
	// Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_7_FINAL_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 704
	// Data Type: URL_LIST. The final mobile URLs of item 7 of the table when
	// using Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_7_FINAL_MOBILE_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 705
	// Data Type: STRING. The header of item 8 of the table.
	PricePlaceholderFieldEnum_ITEM_8_HEADER PricePlaceholderFieldEnum_PricePlaceholderField = 800
	// Data Type: STRING. The description of item 8 of the table.
	PricePlaceholderFieldEnum_ITEM_8_DESCRIPTION PricePlaceholderFieldEnum_PricePlaceholderField = 801
	// Data Type: MONEY. The price (money with currency) of item 8 of the table,
	// e.g., 30 USD. The currency must match one of the available currencies.
	PricePlaceholderFieldEnum_ITEM_8_PRICE PricePlaceholderFieldEnum_PricePlaceholderField = 802
	// Data Type: STRING. The price unit of item 8 of the table. Must match one
	// of the predefined price units.
	PricePlaceholderFieldEnum_ITEM_8_UNIT PricePlaceholderFieldEnum_PricePlaceholderField = 803
	// Data Type: URL_LIST. The final URLs of item 8 of the table when using
	// Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_8_FINAL_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 804
	// Data Type: URL_LIST. The final mobile URLs of item 8 of the table when
	// using Upgraded URLs.
	PricePlaceholderFieldEnum_ITEM_8_FINAL_MOBILE_URLS PricePlaceholderFieldEnum_PricePlaceholderField = 805
)

// Enum value maps for PricePlaceholderFieldEnum_PricePlaceholderField.
var (
	PricePlaceholderFieldEnum_PricePlaceholderField_name = map[int32]string{
		0:   "UNSPECIFIED",
		1:   "UNKNOWN",
		2:   "TYPE",
		3:   "PRICE_QUALIFIER",
		4:   "TRACKING_TEMPLATE",
		5:   "LANGUAGE",
		6:   "FINAL_URL_SUFFIX",
		100: "ITEM_1_HEADER",
		101: "ITEM_1_DESCRIPTION",
		102: "ITEM_1_PRICE",
		103: "ITEM_1_UNIT",
		104: "ITEM_1_FINAL_URLS",
		105: "ITEM_1_FINAL_MOBILE_URLS",
		200: "ITEM_2_HEADER",
		201: "ITEM_2_DESCRIPTION",
		202: "ITEM_2_PRICE",
		203: "ITEM_2_UNIT",
		204: "ITEM_2_FINAL_URLS",
		205: "ITEM_2_FINAL_MOBILE_URLS",
		300: "ITEM_3_HEADER",
		301: "ITEM_3_DESCRIPTION",
		302: "ITEM_3_PRICE",
		303: "ITEM_3_UNIT",
		304: "ITEM_3_FINAL_URLS",
		305: "ITEM_3_FINAL_MOBILE_URLS",
		400: "ITEM_4_HEADER",
		401: "ITEM_4_DESCRIPTION",
		402: "ITEM_4_PRICE",
		403: "ITEM_4_UNIT",
		404: "ITEM_4_FINAL_URLS",
		405: "ITEM_4_FINAL_MOBILE_URLS",
		500: "ITEM_5_HEADER",
		501: "ITEM_5_DESCRIPTION",
		502: "ITEM_5_PRICE",
		503: "ITEM_5_UNIT",
		504: "ITEM_5_FINAL_URLS",
		505: "ITEM_5_FINAL_MOBILE_URLS",
		600: "ITEM_6_HEADER",
		601: "ITEM_6_DESCRIPTION",
		602: "ITEM_6_PRICE",
		603: "ITEM_6_UNIT",
		604: "ITEM_6_FINAL_URLS",
		605: "ITEM_6_FINAL_MOBILE_URLS",
		700: "ITEM_7_HEADER",
		701: "ITEM_7_DESCRIPTION",
		702: "ITEM_7_PRICE",
		703: "ITEM_7_UNIT",
		704: "ITEM_7_FINAL_URLS",
		705: "ITEM_7_FINAL_MOBILE_URLS",
		800: "ITEM_8_HEADER",
		801: "ITEM_8_DESCRIPTION",
		802: "ITEM_8_PRICE",
		803: "ITEM_8_UNIT",
		804: "ITEM_8_FINAL_URLS",
		805: "ITEM_8_FINAL_MOBILE_URLS",
	}
	PricePlaceholderFieldEnum_PricePlaceholderField_value = map[string]int32{
		"UNSPECIFIED":              0,
		"UNKNOWN":                  1,
		"TYPE":                     2,
		"PRICE_QUALIFIER":          3,
		"TRACKING_TEMPLATE":        4,
		"LANGUAGE":                 5,
		"FINAL_URL_SUFFIX":         6,
		"ITEM_1_HEADER":            100,
		"ITEM_1_DESCRIPTION":       101,
		"ITEM_1_PRICE":             102,
		"ITEM_1_UNIT":              103,
		"ITEM_1_FINAL_URLS":        104,
		"ITEM_1_FINAL_MOBILE_URLS": 105,
		"ITEM_2_HEADER":            200,
		"ITEM_2_DESCRIPTION":       201,
		"ITEM_2_PRICE":             202,
		"ITEM_2_UNIT":              203,
		"ITEM_2_FINAL_URLS":        204,
		"ITEM_2_FINAL_MOBILE_URLS": 205,
		"ITEM_3_HEADER":            300,
		"ITEM_3_DESCRIPTION":       301,
		"ITEM_3_PRICE":             302,
		"ITEM_3_UNIT":              303,
		"ITEM_3_FINAL_URLS":        304,
		"ITEM_3_FINAL_MOBILE_URLS": 305,
		"ITEM_4_HEADER":            400,
		"ITEM_4_DESCRIPTION":       401,
		"ITEM_4_PRICE":             402,
		"ITEM_4_UNIT":              403,
		"ITEM_4_FINAL_URLS":        404,
		"ITEM_4_FINAL_MOBILE_URLS": 405,
		"ITEM_5_HEADER":            500,
		"ITEM_5_DESCRIPTION":       501,
		"ITEM_5_PRICE":             502,
		"ITEM_5_UNIT":              503,
		"ITEM_5_FINAL_URLS":        504,
		"ITEM_5_FINAL_MOBILE_URLS": 505,
		"ITEM_6_HEADER":            600,
		"ITEM_6_DESCRIPTION":       601,
		"ITEM_6_PRICE":             602,
		"ITEM_6_UNIT":              603,
		"ITEM_6_FINAL_URLS":        604,
		"ITEM_6_FINAL_MOBILE_URLS": 605,
		"ITEM_7_HEADER":            700,
		"ITEM_7_DESCRIPTION":       701,
		"ITEM_7_PRICE":             702,
		"ITEM_7_UNIT":              703,
		"ITEM_7_FINAL_URLS":        704,
		"ITEM_7_FINAL_MOBILE_URLS": 705,
		"ITEM_8_HEADER":            800,
		"ITEM_8_DESCRIPTION":       801,
		"ITEM_8_PRICE":             802,
		"ITEM_8_UNIT":              803,
		"ITEM_8_FINAL_URLS":        804,
		"ITEM_8_FINAL_MOBILE_URLS": 805,
	}
)

func (x PricePlaceholderFieldEnum_PricePlaceholderField) Enum() *PricePlaceholderFieldEnum_PricePlaceholderField {
	p := new(PricePlaceholderFieldEnum_PricePlaceholderField)
	*p = x
	return p
}

func (x PricePlaceholderFieldEnum_PricePlaceholderField) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PricePlaceholderFieldEnum_PricePlaceholderField) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v9_enums_price_placeholder_field_proto_enumTypes[0].Descriptor()
}

func (PricePlaceholderFieldEnum_PricePlaceholderField) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v9_enums_price_placeholder_field_proto_enumTypes[0]
}

func (x PricePlaceholderFieldEnum_PricePlaceholderField) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PricePlaceholderFieldEnum_PricePlaceholderField.Descriptor instead.
func (PricePlaceholderFieldEnum_PricePlaceholderField) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v9_enums_price_placeholder_field_proto_rawDescGZIP(), []int{0, 0}
}

// Values for Price placeholder fields.
type PricePlaceholderFieldEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PricePlaceholderFieldEnum) Reset() {
	*x = PricePlaceholderFieldEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v9_enums_price_placeholder_field_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PricePlaceholderFieldEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PricePlaceholderFieldEnum) ProtoMessage() {}

func (x *PricePlaceholderFieldEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v9_enums_price_placeholder_field_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PricePlaceholderFieldEnum.ProtoReflect.Descriptor instead.
func (*PricePlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v9_enums_price_placeholder_field_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v9_enums_price_placeholder_field_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v9_enums_price_placeholder_field_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x09, 0x0a, 0x19, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0xd1, 0x09, 0x0a, 0x15, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x54, 0x59, 0x50, 0x45, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x52,
	0x49, 0x43, 0x45, 0x5f, 0x51, 0x55, 0x41, 0x4c, 0x49, 0x46, 0x49, 0x45, 0x52, 0x10, 0x03, 0x12,
	0x15, 0x0a, 0x11, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x45, 0x4d, 0x50,
	0x4c, 0x41, 0x54, 0x45, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41,
	0x47, 0x45, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x55, 0x52,
	0x4c, 0x5f, 0x53, 0x55, 0x46, 0x46, 0x49, 0x58, 0x10, 0x06, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x54,
	0x45, 0x4d, 0x5f, 0x31, 0x5f, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x10, 0x64, 0x12, 0x16, 0x0a,
	0x12, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x31, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x31, 0x5f,
	0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x66, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x54, 0x45, 0x4d, 0x5f,
	0x31, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x10, 0x67, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x54, 0x45, 0x4d,
	0x5f, 0x31, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x55, 0x52, 0x4c, 0x53, 0x10, 0x68, 0x12,
	0x1c, 0x0a, 0x18, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x31, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f,
	0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x55, 0x52, 0x4c, 0x53, 0x10, 0x69, 0x12, 0x12, 0x0a,
	0x0d, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x32, 0x5f, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x10, 0xc8,
	0x01, 0x12, 0x17, 0x0a, 0x12, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x32, 0x5f, 0x44, 0x45, 0x53, 0x43,
	0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0xc9, 0x01, 0x12, 0x11, 0x0a, 0x0c, 0x49, 0x54,
	0x45, 0x4d, 0x5f, 0x32, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0xca, 0x01, 0x12, 0x10, 0x0a,
	0x0b, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x32, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x10, 0xcb, 0x01, 0x12,
	0x16, 0x0a, 0x11, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x32, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f,
	0x55, 0x52, 0x4c, 0x53, 0x10, 0xcc, 0x01, 0x12, 0x1d, 0x0a, 0x18, 0x49, 0x54, 0x45, 0x4d, 0x5f,
	0x32, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x55,
	0x52, 0x4c, 0x53, 0x10, 0xcd, 0x01, 0x12, 0x12, 0x0a, 0x0d, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x33,
	0x5f, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x10, 0xac, 0x02, 0x12, 0x17, 0x0a, 0x12, 0x49, 0x54,
	0x45, 0x4d, 0x5f, 0x33, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0xad, 0x02, 0x12, 0x11, 0x0a, 0x0c, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x33, 0x5f, 0x50, 0x52,
	0x49, 0x43, 0x45, 0x10, 0xae, 0x02, 0x12, 0x10, 0x0a, 0x0b, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x33,
	0x5f, 0x55, 0x4e, 0x49, 0x54, 0x10, 0xaf, 0x02, 0x12, 0x16, 0x0a, 0x11, 0x49, 0x54, 0x45, 0x4d,
	0x5f, 0x33, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x55, 0x52, 0x4c, 0x53, 0x10, 0xb0, 0x02,
	0x12, 0x1d, 0x0a, 0x18, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x33, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c,
	0x5f, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x55, 0x52, 0x4c, 0x53, 0x10, 0xb1, 0x02, 0x12,
	0x12, 0x0a, 0x0d, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x34, 0x5f, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52,
	0x10, 0x90, 0x03, 0x12, 0x17, 0x0a, 0x12, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x34, 0x5f, 0x44, 0x45,
	0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x91, 0x03, 0x12, 0x11, 0x0a, 0x0c,
	0x49, 0x54, 0x45, 0x4d, 0x5f, 0x34, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0x92, 0x03, 0x12,
	0x10, 0x0a, 0x0b, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x34, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x10, 0x93,
	0x03, 0x12, 0x16, 0x0a, 0x11, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x34, 0x5f, 0x46, 0x49, 0x4e, 0x41,
	0x4c, 0x5f, 0x55, 0x52, 0x4c, 0x53, 0x10, 0x94, 0x03, 0x12, 0x1d, 0x0a, 0x18, 0x49, 0x54, 0x45,
	0x4d, 0x5f, 0x34, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45,
	0x5f, 0x55, 0x52, 0x4c, 0x53, 0x10, 0x95, 0x03, 0x12, 0x12, 0x0a, 0x0d, 0x49, 0x54, 0x45, 0x4d,
	0x5f, 0x35, 0x5f, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x10, 0xf4, 0x03, 0x12, 0x17, 0x0a, 0x12,
	0x49, 0x54, 0x45, 0x4d, 0x5f, 0x35, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0xf5, 0x03, 0x12, 0x11, 0x0a, 0x0c, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x35, 0x5f,
	0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0xf6, 0x03, 0x12, 0x10, 0x0a, 0x0b, 0x49, 0x54, 0x45, 0x4d,
	0x5f, 0x35, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x10, 0xf7, 0x03, 0x12, 0x16, 0x0a, 0x11, 0x49, 0x54,
	0x45, 0x4d, 0x5f, 0x35, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x55, 0x52, 0x4c, 0x53, 0x10,
	0xf8, 0x03, 0x12, 0x1d, 0x0a, 0x18, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x35, 0x5f, 0x46, 0x49, 0x4e,
	0x41, 0x4c, 0x5f, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x55, 0x52, 0x4c, 0x53, 0x10, 0xf9,
	0x03, 0x12, 0x12, 0x0a, 0x0d, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x36, 0x5f, 0x48, 0x45, 0x41, 0x44,
	0x45, 0x52, 0x10, 0xd8, 0x04, 0x12, 0x17, 0x0a, 0x12, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x36, 0x5f,
	0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0xd9, 0x04, 0x12, 0x11,
	0x0a, 0x0c, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x36, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0xda,
	0x04, 0x12, 0x10, 0x0a, 0x0b, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x36, 0x5f, 0x55, 0x4e, 0x49, 0x54,
	0x10, 0xdb, 0x04, 0x12, 0x16, 0x0a, 0x11, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x36, 0x5f, 0x46, 0x49,
	0x4e, 0x41, 0x4c, 0x5f, 0x55, 0x52, 0x4c, 0x53, 0x10, 0xdc, 0x04, 0x12, 0x1d, 0x0a, 0x18, 0x49,
	0x54, 0x45, 0x4d, 0x5f, 0x36, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x4d, 0x4f, 0x42, 0x49,
	0x4c, 0x45, 0x5f, 0x55, 0x52, 0x4c, 0x53, 0x10, 0xdd, 0x04, 0x12, 0x12, 0x0a, 0x0d, 0x49, 0x54,
	0x45, 0x4d, 0x5f, 0x37, 0x5f, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x10, 0xbc, 0x05, 0x12, 0x17,
	0x0a, 0x12, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x37, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0xbd, 0x05, 0x12, 0x11, 0x0a, 0x0c, 0x49, 0x54, 0x45, 0x4d, 0x5f,
	0x37, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10, 0xbe, 0x05, 0x12, 0x10, 0x0a, 0x0b, 0x49, 0x54,
	0x45, 0x4d, 0x5f, 0x37, 0x5f, 0x55, 0x4e, 0x49, 0x54, 0x10, 0xbf, 0x05, 0x12, 0x16, 0x0a, 0x11,
	0x49, 0x54, 0x45, 0x4d, 0x5f, 0x37, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x55, 0x52, 0x4c,
	0x53, 0x10, 0xc0, 0x05, 0x12, 0x1d, 0x0a, 0x18, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x37, 0x5f, 0x46,
	0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x5f, 0x55, 0x52, 0x4c, 0x53,
	0x10, 0xc1, 0x05, 0x12, 0x12, 0x0a, 0x0d, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x38, 0x5f, 0x48, 0x45,
	0x41, 0x44, 0x45, 0x52, 0x10, 0xa0, 0x06, 0x12, 0x17, 0x0a, 0x12, 0x49, 0x54, 0x45, 0x4d, 0x5f,
	0x38, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0xa1, 0x06,
	0x12, 0x11, 0x0a, 0x0c, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x38, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45,
	0x10, 0xa2, 0x06, 0x12, 0x10, 0x0a, 0x0b, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x38, 0x5f, 0x55, 0x4e,
	0x49, 0x54, 0x10, 0xa3, 0x06, 0x12, 0x16, 0x0a, 0x11, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x38, 0x5f,
	0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x55, 0x52, 0x4c, 0x53, 0x10, 0xa4, 0x06, 0x12, 0x1d, 0x0a,
	0x18, 0x49, 0x54, 0x45, 0x4d, 0x5f, 0x38, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x5f, 0x4d, 0x4f,
	0x42, 0x49, 0x4c, 0x45, 0x5f, 0x55, 0x52, 0x4c, 0x53, 0x10, 0xa5, 0x06, 0x42, 0xef, 0x01, 0x0a,
	0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x42, 0x1a, 0x50, 0x72, 0x69, 0x63, 0x65, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x42, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1d, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x2e, 0x56, 0x39, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xca, 0x02, 0x1d, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64,
	0x73, 0x5c, 0x56, 0x39, 0x5c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x39, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v9_enums_price_placeholder_field_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v9_enums_price_placeholder_field_proto_rawDescData = file_google_ads_googleads_v9_enums_price_placeholder_field_proto_rawDesc
)

func file_google_ads_googleads_v9_enums_price_placeholder_field_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v9_enums_price_placeholder_field_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v9_enums_price_placeholder_field_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v9_enums_price_placeholder_field_proto_rawDescData)
	})
	return file_google_ads_googleads_v9_enums_price_placeholder_field_proto_rawDescData
}

var file_google_ads_googleads_v9_enums_price_placeholder_field_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v9_enums_price_placeholder_field_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v9_enums_price_placeholder_field_proto_goTypes = []interface{}{
	(PricePlaceholderFieldEnum_PricePlaceholderField)(0), // 0: google.ads.googleads.v9.enums.PricePlaceholderFieldEnum.PricePlaceholderField
	(*PricePlaceholderFieldEnum)(nil),                    // 1: google.ads.googleads.v9.enums.PricePlaceholderFieldEnum
}
var file_google_ads_googleads_v9_enums_price_placeholder_field_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v9_enums_price_placeholder_field_proto_init() }
func file_google_ads_googleads_v9_enums_price_placeholder_field_proto_init() {
	if File_google_ads_googleads_v9_enums_price_placeholder_field_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v9_enums_price_placeholder_field_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PricePlaceholderFieldEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v9_enums_price_placeholder_field_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v9_enums_price_placeholder_field_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v9_enums_price_placeholder_field_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v9_enums_price_placeholder_field_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v9_enums_price_placeholder_field_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v9_enums_price_placeholder_field_proto = out.File
	file_google_ads_googleads_v9_enums_price_placeholder_field_proto_rawDesc = nil
	file_google_ads_googleads_v9_enums_price_placeholder_field_proto_goTypes = nil
	file_google_ads_googleads_v9_enums_price_placeholder_field_proto_depIdxs = nil
}
