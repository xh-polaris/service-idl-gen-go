// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.19.4
// source: openapi/core_api/core_api.proto

package core_api

import (
	context "context"
	_ "github.com/xh-polaris/service-idl-gen-go/kitex_gen/http"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_openapi_core_api_core_api_proto protoreflect.FileDescriptor

var file_openapi_core_api_core_api_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x1a, 0x1d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0xf7, 0x02, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x56, 0x0a, 0x06,
	0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x1b, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70,
	0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x11, 0xd2, 0xc1, 0x18, 0x0d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x69, 0x67,
	0x6e, 0x5f, 0x75, 0x70, 0x12, 0x56, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x1b,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x22, 0x11, 0xd2, 0xc1, 0x18, 0x0d, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x6e, 0x12, 0x62, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x0e, 0xca, 0xc1, 0x18, 0x0a, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x66, 0x6f,
	0x12, 0x5b, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x20, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0e, 0xd2,
	0xc1, 0x18, 0x0a, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xa8, 0x04,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x5e, 0x0a, 0x0b, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x4b, 0x65, 0x79, 0x12, 0x20, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x11, 0xd2, 0xc1, 0x18, 0x0d, 0x2f, 0x6b, 0x65, 0x79, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x54, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x73,
	0x12, 0x1c, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1d,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x0c, 0xd2,
	0xc1, 0x18, 0x08, 0x2f, 0x6b, 0x65, 0x79, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x58, 0x0a, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0xd2, 0xc1, 0x18, 0x0b, 0x2f, 0x6b, 0x65, 0x79, 0x2f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x5b, 0x0a, 0x0a, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x4b, 0x65, 0x79, 0x12, 0x1f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x10, 0xca, 0xc1, 0x18, 0x0c, 0x2f, 0x6b, 0x65, 0x79, 0x2f, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x12, 0x5a, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74,
	0x73, 0x12, 0x1f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0e,
	0xd2, 0xc1, 0x18, 0x0a, 0x2f, 0x6b, 0x65, 0x79, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x58,
	0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0xca, 0xc1, 0x18, 0x0b, 0x2f, 0x6b, 0x65,
	0x79, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x32, 0x8c, 0x0b, 0x0a, 0x06, 0x63, 0x68, 0x61,
	0x72, 0x67, 0x65, 0x12, 0x6d, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x73,
	0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x28, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x10, 0xd2, 0xc1, 0x18, 0x0c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x6d, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x73, 0x65,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x28, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x10, 0xd2, 0xc1, 0x18, 0x0c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x6d, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x73, 0x65, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x28, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10,
	0xca, 0xc1, 0x18, 0x0c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x73, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x73, 0x12, 0x26, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x73, 0x65,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x0d, 0xd2, 0xc1, 0x18, 0x09, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x6d, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46,
	0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x28, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x10, 0xd2, 0xc1, 0x18, 0x0c, 0x2f, 0x66, 0x75, 0x6c, 0x6c, 0x2f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x6d, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x75,
	0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x28, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x10, 0xd2, 0xc1, 0x18, 0x0c, 0x2f, 0x66, 0x75, 0x6c, 0x6c, 0x2f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x5f, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x72,
	0x67, 0x69, 0x6e, 0x12, 0x21, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x72,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x10, 0xd2, 0xc1, 0x18, 0x0c, 0x2f, 0x66, 0x75, 0x6c, 0x6c, 0x2f, 0x6d, 0x61,
	0x72, 0x67, 0x69, 0x6e, 0x12, 0x6d, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x75,
	0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x28, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x10, 0xca, 0xc1, 0x18, 0x0c, 0x2f, 0x66, 0x75, 0x6c, 0x6c, 0x2f, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x73, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x12, 0x26, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x46,
	0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x1a, 0x27, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x0d, 0xd2, 0xc1, 0x18, 0x09, 0x2f,
	0x66, 0x75, 0x6c, 0x6c, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x64, 0x0a, 0x10, 0x42, 0x75, 0x79, 0x46,
	0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x25, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x42, 0x75, 0x79, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x0d, 0xd2, 0xc1, 0x18, 0x09, 0x2f, 0x66, 0x75, 0x6c, 0x6c, 0x2f, 0x62, 0x75, 0x79, 0x12, 0x67,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x23, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x61, 0x64, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x14, 0xd2, 0xc1, 0x18, 0x10, 0x2f, 0x67, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x67, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x47, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0xd2, 0xc1, 0x18, 0x10,
	0x2f, 0x67, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x65, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x47, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x12,
	0x20, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x21, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x11, 0xca, 0xc1, 0x18, 0x0d, 0x2f, 0x67, 0x72, 0x61, 0x64, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x67, 0x65, 0x74, 0x32, 0x6c, 0x0a, 0x04, 0x63, 0x61, 0x6c, 0x6c, 0x12,
	0x64, 0x0a, 0x0d, 0x43, 0x61, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x12, 0x22, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x0a, 0xd2, 0xc1, 0x18, 0x06, 0x2f,
	0x63, 0x61, 0x6c, 0x6c, 0x2f, 0x42, 0x7b, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x68, 0x70,
	0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2e, 0x69, 0x64, 0x6c, 0x67, 0x65, 0x6e, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x61, 0x70, 0x69, 0x42, 0x0c, 0x43,
	0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x68, 0x2d, 0x70, 0x6f, 0x6c,
	0x61, 0x72, 0x69, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x69, 0x64, 0x6c,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65,
	0x6e, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_openapi_core_api_core_api_proto_goTypes = []interface{}{
	(*SignUpReq)(nil),              // 0: openapi.core_api.SignUpReq
	(*GetUserInfoReq)(nil),         // 1: openapi.core_api.GetUserInfoReq
	(*SetUserInfoReq)(nil),         // 2: openapi.core_api.SetUserInfoReq
	(*GenerateKeyReq)(nil),         // 3: openapi.core_api.GenerateKeyReq
	(*GetKeysReq)(nil),             // 4: openapi.core_api.GetKeysReq
	(*UpdateKeyReq)(nil),           // 5: openapi.core_api.UpdateKeyReq
	(*RefreshKeyReq)(nil),          // 6: openapi.core_api.RefreshKeyReq
	(*UpdateHostReq)(nil),          // 7: openapi.core_api.UpdateHostReq
	(*DeleteKeyReq)(nil),           // 8: openapi.core_api.DeleteKeyReq
	(*CreateBaseInterfaceReq)(nil), // 9: openapi.core_api.CreateBaseInterfaceReq
	(*UpdateBaseInterfaceReq)(nil), // 10: openapi.core_api.UpdateBaseInterfaceReq
	(*DeleteBaseInterfaceReq)(nil), // 11: openapi.core_api.DeleteBaseInterfaceReq
	(*GetBaseInterfacesReq)(nil),   // 12: openapi.core_api.GetBaseInterfacesReq
	(*CreateFullInterfaceReq)(nil), // 13: openapi.core_api.CreateFullInterfaceReq
	(*UpdateFullInterfaceReq)(nil), // 14: openapi.core_api.UpdateFullInterfaceReq
	(*UpdateMarginReq)(nil),        // 15: openapi.core_api.UpdateMarginReq
	(*DeleteFullInterfaceReq)(nil), // 16: openapi.core_api.DeleteFullInterfaceReq
	(*GetFullInterfacesReq)(nil),   // 17: openapi.core_api.GetFullInterfacesReq
	(*BuyFullInterfaceReq)(nil),    // 18: openapi.core_api.BuyFullInterfaceReq
	(*CreateGradientReq)(nil),      // 19: openapi.core_api.CreateGradientReq
	(*UpdateGradientReq)(nil),      // 20: openapi.core_api.UpdateGradientReq
	(*GetGradientReq)(nil),         // 21: openapi.core_api.GetGradientReq
	(*CallInterfaceReq)(nil),       // 22: openapi.core_api.CallInterfaceReq
	(*SignUpResp)(nil),             // 23: openapi.core_api.SignUpResp
	(*GetUserInfoResp)(nil),        // 24: openapi.core_api.GetUserInfoResp
	(*Response)(nil),               // 25: openapi.core_api.Response
	(*GetKeysResp)(nil),            // 26: openapi.core_api.GetKeysResp
	(*GetBaseInterfacesResp)(nil),  // 27: openapi.core_api.GetBaseInterfacesResp
	(*GetFullInterfacesResp)(nil),  // 28: openapi.core_api.GetFullInterfacesResp
	(*GetGradientResp)(nil),        // 29: openapi.core_api.GetGradientResp
	(*CallInterfaceResp)(nil),      // 30: openapi.core_api.CallInterfaceResp
}
var file_openapi_core_api_core_api_proto_depIdxs = []int32{
	0,  // 0: openapi.core_api.user.SignUp:input_type -> openapi.core_api.SignUpReq
	0,  // 1: openapi.core_api.user.SignIn:input_type -> openapi.core_api.SignUpReq
	1,  // 2: openapi.core_api.user.GetUserInfo:input_type -> openapi.core_api.GetUserInfoReq
	2,  // 3: openapi.core_api.user.SetUserInfo:input_type -> openapi.core_api.SetUserInfoReq
	3,  // 4: openapi.core_api.key.GenerateKey:input_type -> openapi.core_api.GenerateKeyReq
	4,  // 5: openapi.core_api.key.GetKeys:input_type -> openapi.core_api.GetKeysReq
	5,  // 6: openapi.core_api.key.UpdateKey:input_type -> openapi.core_api.UpdateKeyReq
	6,  // 7: openapi.core_api.key.RefreshKey:input_type -> openapi.core_api.RefreshKeyReq
	7,  // 8: openapi.core_api.key.UpdateHosts:input_type -> openapi.core_api.UpdateHostReq
	8,  // 9: openapi.core_api.key.DeleteKey:input_type -> openapi.core_api.DeleteKeyReq
	9,  // 10: openapi.core_api.charge.CreateBaseInterface:input_type -> openapi.core_api.CreateBaseInterfaceReq
	10, // 11: openapi.core_api.charge.UpdateBaseInterface:input_type -> openapi.core_api.UpdateBaseInterfaceReq
	11, // 12: openapi.core_api.charge.DeleteBaseInterface:input_type -> openapi.core_api.DeleteBaseInterfaceReq
	12, // 13: openapi.core_api.charge.GetBaseInterfaces:input_type -> openapi.core_api.GetBaseInterfacesReq
	13, // 14: openapi.core_api.charge.CreateFullInterface:input_type -> openapi.core_api.CreateFullInterfaceReq
	14, // 15: openapi.core_api.charge.UpdateFullInterface:input_type -> openapi.core_api.UpdateFullInterfaceReq
	15, // 16: openapi.core_api.charge.UpdateMargin:input_type -> openapi.core_api.UpdateMarginReq
	16, // 17: openapi.core_api.charge.DeleteFullInterface:input_type -> openapi.core_api.DeleteFullInterfaceReq
	17, // 18: openapi.core_api.charge.GetFullInterfaces:input_type -> openapi.core_api.GetFullInterfacesReq
	18, // 19: openapi.core_api.charge.BuyFullInterface:input_type -> openapi.core_api.BuyFullInterfaceReq
	19, // 20: openapi.core_api.charge.CreateGradient:input_type -> openapi.core_api.CreateGradientReq
	20, // 21: openapi.core_api.charge.UpdateGradient:input_type -> openapi.core_api.UpdateGradientReq
	21, // 22: openapi.core_api.charge.GetGradient:input_type -> openapi.core_api.GetGradientReq
	22, // 23: openapi.core_api.call.CallInterface:input_type -> openapi.core_api.CallInterfaceReq
	23, // 24: openapi.core_api.user.SignUp:output_type -> openapi.core_api.SignUpResp
	23, // 25: openapi.core_api.user.SignIn:output_type -> openapi.core_api.SignUpResp
	24, // 26: openapi.core_api.user.GetUserInfo:output_type -> openapi.core_api.GetUserInfoResp
	25, // 27: openapi.core_api.user.SetUserInfo:output_type -> openapi.core_api.Response
	25, // 28: openapi.core_api.key.GenerateKey:output_type -> openapi.core_api.Response
	26, // 29: openapi.core_api.key.GetKeys:output_type -> openapi.core_api.GetKeysResp
	25, // 30: openapi.core_api.key.UpdateKey:output_type -> openapi.core_api.Response
	25, // 31: openapi.core_api.key.RefreshKey:output_type -> openapi.core_api.Response
	25, // 32: openapi.core_api.key.UpdateHosts:output_type -> openapi.core_api.Response
	25, // 33: openapi.core_api.key.DeleteKey:output_type -> openapi.core_api.Response
	25, // 34: openapi.core_api.charge.CreateBaseInterface:output_type -> openapi.core_api.Response
	25, // 35: openapi.core_api.charge.UpdateBaseInterface:output_type -> openapi.core_api.Response
	25, // 36: openapi.core_api.charge.DeleteBaseInterface:output_type -> openapi.core_api.Response
	27, // 37: openapi.core_api.charge.GetBaseInterfaces:output_type -> openapi.core_api.GetBaseInterfacesResp
	25, // 38: openapi.core_api.charge.CreateFullInterface:output_type -> openapi.core_api.Response
	25, // 39: openapi.core_api.charge.UpdateFullInterface:output_type -> openapi.core_api.Response
	25, // 40: openapi.core_api.charge.UpdateMargin:output_type -> openapi.core_api.Response
	25, // 41: openapi.core_api.charge.DeleteFullInterface:output_type -> openapi.core_api.Response
	28, // 42: openapi.core_api.charge.GetFullInterfaces:output_type -> openapi.core_api.GetFullInterfacesResp
	25, // 43: openapi.core_api.charge.BuyFullInterface:output_type -> openapi.core_api.Response
	25, // 44: openapi.core_api.charge.CreateGradient:output_type -> openapi.core_api.Response
	25, // 45: openapi.core_api.charge.UpdateGradient:output_type -> openapi.core_api.Response
	29, // 46: openapi.core_api.charge.GetGradient:output_type -> openapi.core_api.GetGradientResp
	30, // 47: openapi.core_api.call.CallInterface:output_type -> openapi.core_api.CallInterfaceResp
	24, // [24:48] is the sub-list for method output_type
	0,  // [0:24] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_openapi_core_api_core_api_proto_init() }
func file_openapi_core_api_core_api_proto_init() {
	if File_openapi_core_api_core_api_proto != nil {
		return
	}
	file_openapi_core_api_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_openapi_core_api_core_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   4,
		},
		GoTypes:           file_openapi_core_api_core_api_proto_goTypes,
		DependencyIndexes: file_openapi_core_api_core_api_proto_depIdxs,
	}.Build()
	File_openapi_core_api_core_api_proto = out.File
	file_openapi_core_api_core_api_proto_rawDesc = nil
	file_openapi_core_api_core_api_proto_goTypes = nil
	file_openapi_core_api_core_api_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.12.0. DO NOT EDIT.

type User interface {
	SignUp(ctx context.Context, req *SignUpReq) (res *SignUpResp, err error)
	SignIn(ctx context.Context, req *SignUpReq) (res *SignUpResp, err error)
	GetUserInfo(ctx context.Context, req *GetUserInfoReq) (res *GetUserInfoResp, err error)
	SetUserInfo(ctx context.Context, req *SetUserInfoReq) (res *Response, err error)
}

type Key interface {
	GenerateKey(ctx context.Context, req *GenerateKeyReq) (res *Response, err error)
	GetKeys(ctx context.Context, req *GetKeysReq) (res *GetKeysResp, err error)
	UpdateKey(ctx context.Context, req *UpdateKeyReq) (res *Response, err error)
	RefreshKey(ctx context.Context, req *RefreshKeyReq) (res *Response, err error)
	UpdateHosts(ctx context.Context, req *UpdateHostReq) (res *Response, err error)
	DeleteKey(ctx context.Context, req *DeleteKeyReq) (res *Response, err error)
}

type Charge interface {
	CreateBaseInterface(ctx context.Context, req *CreateBaseInterfaceReq) (res *Response, err error)
	UpdateBaseInterface(ctx context.Context, req *UpdateBaseInterfaceReq) (res *Response, err error)
	DeleteBaseInterface(ctx context.Context, req *DeleteBaseInterfaceReq) (res *Response, err error)
	GetBaseInterfaces(ctx context.Context, req *GetBaseInterfacesReq) (res *GetBaseInterfacesResp, err error)
	CreateFullInterface(ctx context.Context, req *CreateFullInterfaceReq) (res *Response, err error)
	UpdateFullInterface(ctx context.Context, req *UpdateFullInterfaceReq) (res *Response, err error)
	UpdateMargin(ctx context.Context, req *UpdateMarginReq) (res *Response, err error)
	DeleteFullInterface(ctx context.Context, req *DeleteFullInterfaceReq) (res *Response, err error)
	GetFullInterfaces(ctx context.Context, req *GetFullInterfacesReq) (res *GetFullInterfacesResp, err error)
	BuyFullInterface(ctx context.Context, req *BuyFullInterfaceReq) (res *Response, err error)
	CreateGradient(ctx context.Context, req *CreateGradientReq) (res *Response, err error)
	UpdateGradient(ctx context.Context, req *UpdateGradientReq) (res *Response, err error)
	GetGradient(ctx context.Context, req *GetGradientReq) (res *GetGradientResp, err error)
}

type Call interface {
	CallInterface(ctx context.Context, req *CallInterfaceReq) (res *CallInterfaceResp, err error)
}
