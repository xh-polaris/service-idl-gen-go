// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.19.4
// source: openapi/charge/charge.proto

package charge

import (
	context "context"
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

var File_openapi_charge_charge_proto protoreflect.FileDescriptor

var file_openapi_charge_charge_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65,
	0x2f, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x1a, 0x1e, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2f, 0x6c, 0x6f,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa4, 0x0f, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x72,
	0x67, 0x65, 0x12, 0x66, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x73, 0x65,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x26, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x27, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72,
	0x67, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x66, 0x0a, 0x13, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x12, 0x26, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72,
	0x67, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x66, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x73, 0x65,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x26, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x27, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72,
	0x67, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x5d, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x23,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68,
	0x61, 0x72, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x66, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x12, 0x26, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67,
	0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x66, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6c, 0x6c, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x26, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x27, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67,
	0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x66, 0x0a, 0x13, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x12, 0x26, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67,
	0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x5d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x23, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x46,
	0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x66, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x26, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x46,
	0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x27, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x46, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x51, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x12, 0x1f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x51, 0x0a, 0x0c, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x12, 0x1f, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x48,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x12, 0x1c, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61,
	0x72, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x51, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x12, 0x1f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x8a, 0x01, 0x0a, 0x1f,
	0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x41, 0x6e, 0x64, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x46, 0x6f, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12,
	0x32, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x41, 0x6e, 0x64, 0x42, 0x61, 0x73, 0x65, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x46, 0x6f, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x1a, 0x33, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68,
	0x61, 0x72, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x41, 0x6e, 0x64, 0x42,
	0x61, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x46, 0x6f, 0x72, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x57, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x47, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x47, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x57, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x61, 0x64, 0x69,
	0x65, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68,
	0x61, 0x72, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x61, 0x64, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72,
	0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x4e, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x47, 0x72, 0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72,
	0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72,
	0x61, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x48, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x48, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f,
	0x67, 0x12, 0x1c, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72,
	0x67, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x1a,
	0x1d, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3f,
	0x0a, 0x06, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x19, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67,
	0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68,
	0x61, 0x72, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x5d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x54,
	0x78, 0x49, 0x64, 0x12, 0x23, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68,
	0x61, 0x72, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x79, 0x54, 0x78, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x42, 0x79, 0x54, 0x78, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x42, 0x77,
	0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x68, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2e,
	0x69, 0x64, 0x6c, 0x67, 0x65, 0x6e, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x68, 0x61, 0x72, 0x67, 0x65, 0x42, 0x0b, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x78, 0x68, 0x2d, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2d, 0x69, 0x64, 0x6c, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x6b,
	0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_openapi_charge_charge_proto_goTypes = []interface{}{
	(*CreateBaseInterfaceReq)(nil),              // 0: openapi.charge.CreateBaseInterfaceReq
	(*UpdateBaseInterfaceReq)(nil),              // 1: openapi.charge.UpdateBaseInterfaceReq
	(*DeleteBaseInterfaceReq)(nil),              // 2: openapi.charge.DeleteBaseInterfaceReq
	(*GetBaseInterfaceReq)(nil),                 // 3: openapi.charge.GetBaseInterfaceReq
	(*CreateFullInterfaceReq)(nil),              // 4: openapi.charge.CreateFullInterfaceReq
	(*UpdateFullInterfaceReq)(nil),              // 5: openapi.charge.UpdateFullInterfaceReq
	(*DeleteFullInterfaceReq)(nil),              // 6: openapi.charge.DeleteFullInterfaceReq
	(*GetFullInterfaceReq)(nil),                 // 7: openapi.charge.GetFullInterfaceReq
	(*GetOneFullInterfaceReq)(nil),              // 8: openapi.charge.GetOneFullInterfaceReq
	(*CreateMarginReq)(nil),                     // 9: openapi.charge.CreateMarginReq
	(*UpdateMarginReq)(nil),                     // 10: openapi.charge.UpdateMarginReq
	(*GetMarginReq)(nil),                        // 11: openapi.charge.GetMarginReq
	(*DeleteMarginReq)(nil),                     // 12: openapi.charge.DeleteMarginReq
	(*GetFullAndBaseInterfaceForCheckReq)(nil),  // 13: openapi.charge.GetFullAndBaseInterfaceForCheckReq
	(*CreateGradientReq)(nil),                   // 14: openapi.charge.CreateGradientReq
	(*UpdateGradientReq)(nil),                   // 15: openapi.charge.UpdateGradientReq
	(*GetGradientReq)(nil),                      // 16: openapi.charge.GetGradientReq
	(*GetAmountReq)(nil),                        // 17: openapi.charge.GetAmountReq
	(*CreateLogReq)(nil),                        // 18: openapi.charge.CreateLogReq
	(*GetLogReq)(nil),                           // 19: openapi.charge.GetLogReq
	(*GetAccountByTxIdReq)(nil),                 // 20: openapi.charge.GetAccountByTxIdReq
	(*CreateBaseInterfaceResp)(nil),             // 21: openapi.charge.CreateBaseInterfaceResp
	(*UpdateBaseInterfaceResp)(nil),             // 22: openapi.charge.UpdateBaseInterfaceResp
	(*DeleteBaseInterfaceResp)(nil),             // 23: openapi.charge.DeleteBaseInterfaceResp
	(*GetBaseInterfaceResp)(nil),                // 24: openapi.charge.GetBaseInterfaceResp
	(*CreateFullInterfaceResp)(nil),             // 25: openapi.charge.CreateFullInterfaceResp
	(*UpdateFullInterfaceResp)(nil),             // 26: openapi.charge.UpdateFullInterfaceResp
	(*DeleteFullInterfaceResp)(nil),             // 27: openapi.charge.DeleteFullInterfaceResp
	(*GetFullInterfaceResp)(nil),                // 28: openapi.charge.GetFullInterfaceResp
	(*GetOneFullInterfaceResp)(nil),             // 29: openapi.charge.GetOneFullInterfaceResp
	(*CreateMarginResp)(nil),                    // 30: openapi.charge.CreateMarginResp
	(*UpdateMarginResp)(nil),                    // 31: openapi.charge.UpdateMarginResp
	(*GetMarginResp)(nil),                       // 32: openapi.charge.GetMarginResp
	(*DeleteMarginResp)(nil),                    // 33: openapi.charge.DeleteMarginResp
	(*GetFullAndBaseInterfaceForCheckResp)(nil), // 34: openapi.charge.GetFullAndBaseInterfaceForCheckResp
	(*CreateGradientResp)(nil),                  // 35: openapi.charge.CreateGradientResp
	(*UpdateGradientResp)(nil),                  // 36: openapi.charge.UpdateGradientResp
	(*GetGradientResp)(nil),                     // 37: openapi.charge.GetGradientResp
	(*GetAmountResp)(nil),                       // 38: openapi.charge.GetAmountResp
	(*CreateLogResp)(nil),                       // 39: openapi.charge.CreateLogResp
	(*GetLogResp)(nil),                          // 40: openapi.charge.GetLogResp
	(*GetAccountByTxIdResp)(nil),                // 41: openapi.charge.GetAccountByTxIdResp
}
var file_openapi_charge_charge_proto_depIdxs = []int32{
	0,  // 0: openapi.charge.charge.CreateBaseInterface:input_type -> openapi.charge.CreateBaseInterfaceReq
	1,  // 1: openapi.charge.charge.UpdateBaseInterface:input_type -> openapi.charge.UpdateBaseInterfaceReq
	2,  // 2: openapi.charge.charge.DeleteBaseInterface:input_type -> openapi.charge.DeleteBaseInterfaceReq
	3,  // 3: openapi.charge.charge.GetBaseInterface:input_type -> openapi.charge.GetBaseInterfaceReq
	4,  // 4: openapi.charge.charge.CreateFullInterface:input_type -> openapi.charge.CreateFullInterfaceReq
	5,  // 5: openapi.charge.charge.UpdateFullInterface:input_type -> openapi.charge.UpdateFullInterfaceReq
	6,  // 6: openapi.charge.charge.DeleteFullInterface:input_type -> openapi.charge.DeleteFullInterfaceReq
	7,  // 7: openapi.charge.charge.GetFullInterface:input_type -> openapi.charge.GetFullInterfaceReq
	8,  // 8: openapi.charge.charge.GetOneFullInterface:input_type -> openapi.charge.GetOneFullInterfaceReq
	9,  // 9: openapi.charge.charge.CreateMargin:input_type -> openapi.charge.CreateMarginReq
	10, // 10: openapi.charge.charge.UpdateMargin:input_type -> openapi.charge.UpdateMarginReq
	11, // 11: openapi.charge.charge.GetMargin:input_type -> openapi.charge.GetMarginReq
	12, // 12: openapi.charge.charge.DeleteMargin:input_type -> openapi.charge.DeleteMarginReq
	13, // 13: openapi.charge.charge.GetFullAndBaseInterfaceForCheck:input_type -> openapi.charge.GetFullAndBaseInterfaceForCheckReq
	14, // 14: openapi.charge.charge.CreateGradient:input_type -> openapi.charge.CreateGradientReq
	15, // 15: openapi.charge.charge.UpdateGradient:input_type -> openapi.charge.UpdateGradientReq
	16, // 16: openapi.charge.charge.GetGradient:input_type -> openapi.charge.GetGradientReq
	17, // 17: openapi.charge.charge.GetAmount:input_type -> openapi.charge.GetAmountReq
	18, // 18: openapi.charge.charge.CreateLog:input_type -> openapi.charge.CreateLogReq
	19, // 19: openapi.charge.charge.GetLog:input_type -> openapi.charge.GetLogReq
	20, // 20: openapi.charge.charge.GetAccountByTxId:input_type -> openapi.charge.GetAccountByTxIdReq
	21, // 21: openapi.charge.charge.CreateBaseInterface:output_type -> openapi.charge.CreateBaseInterfaceResp
	22, // 22: openapi.charge.charge.UpdateBaseInterface:output_type -> openapi.charge.UpdateBaseInterfaceResp
	23, // 23: openapi.charge.charge.DeleteBaseInterface:output_type -> openapi.charge.DeleteBaseInterfaceResp
	24, // 24: openapi.charge.charge.GetBaseInterface:output_type -> openapi.charge.GetBaseInterfaceResp
	25, // 25: openapi.charge.charge.CreateFullInterface:output_type -> openapi.charge.CreateFullInterfaceResp
	26, // 26: openapi.charge.charge.UpdateFullInterface:output_type -> openapi.charge.UpdateFullInterfaceResp
	27, // 27: openapi.charge.charge.DeleteFullInterface:output_type -> openapi.charge.DeleteFullInterfaceResp
	28, // 28: openapi.charge.charge.GetFullInterface:output_type -> openapi.charge.GetFullInterfaceResp
	29, // 29: openapi.charge.charge.GetOneFullInterface:output_type -> openapi.charge.GetOneFullInterfaceResp
	30, // 30: openapi.charge.charge.CreateMargin:output_type -> openapi.charge.CreateMarginResp
	31, // 31: openapi.charge.charge.UpdateMargin:output_type -> openapi.charge.UpdateMarginResp
	32, // 32: openapi.charge.charge.GetMargin:output_type -> openapi.charge.GetMarginResp
	33, // 33: openapi.charge.charge.DeleteMargin:output_type -> openapi.charge.DeleteMarginResp
	34, // 34: openapi.charge.charge.GetFullAndBaseInterfaceForCheck:output_type -> openapi.charge.GetFullAndBaseInterfaceForCheckResp
	35, // 35: openapi.charge.charge.CreateGradient:output_type -> openapi.charge.CreateGradientResp
	36, // 36: openapi.charge.charge.UpdateGradient:output_type -> openapi.charge.UpdateGradientResp
	37, // 37: openapi.charge.charge.GetGradient:output_type -> openapi.charge.GetGradientResp
	38, // 38: openapi.charge.charge.GetAmount:output_type -> openapi.charge.GetAmountResp
	39, // 39: openapi.charge.charge.CreateLog:output_type -> openapi.charge.CreateLogResp
	40, // 40: openapi.charge.charge.GetLog:output_type -> openapi.charge.GetLogResp
	41, // 41: openapi.charge.charge.GetAccountByTxId:output_type -> openapi.charge.GetAccountByTxIdResp
	21, // [21:42] is the sub-list for method output_type
	0,  // [0:21] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_openapi_charge_charge_proto_init() }
func file_openapi_charge_charge_proto_init() {
	if File_openapi_charge_charge_proto != nil {
		return
	}
	file_openapi_charge_interface_proto_init()
	file_openapi_charge_log_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_openapi_charge_charge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_openapi_charge_charge_proto_goTypes,
		DependencyIndexes: file_openapi_charge_charge_proto_depIdxs,
	}.Build()
	File_openapi_charge_charge_proto = out.File
	file_openapi_charge_charge_proto_rawDesc = nil
	file_openapi_charge_charge_proto_goTypes = nil
	file_openapi_charge_charge_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.12.1. DO NOT EDIT.

type Charge interface {
	CreateBaseInterface(ctx context.Context, req *CreateBaseInterfaceReq) (res *CreateBaseInterfaceResp, err error)
	UpdateBaseInterface(ctx context.Context, req *UpdateBaseInterfaceReq) (res *UpdateBaseInterfaceResp, err error)
	DeleteBaseInterface(ctx context.Context, req *DeleteBaseInterfaceReq) (res *DeleteBaseInterfaceResp, err error)
	GetBaseInterface(ctx context.Context, req *GetBaseInterfaceReq) (res *GetBaseInterfaceResp, err error)
	CreateFullInterface(ctx context.Context, req *CreateFullInterfaceReq) (res *CreateFullInterfaceResp, err error)
	UpdateFullInterface(ctx context.Context, req *UpdateFullInterfaceReq) (res *UpdateFullInterfaceResp, err error)
	DeleteFullInterface(ctx context.Context, req *DeleteFullInterfaceReq) (res *DeleteFullInterfaceResp, err error)
	GetFullInterface(ctx context.Context, req *GetFullInterfaceReq) (res *GetFullInterfaceResp, err error)
	GetOneFullInterface(ctx context.Context, req *GetOneFullInterfaceReq) (res *GetOneFullInterfaceResp, err error)
	CreateMargin(ctx context.Context, req *CreateMarginReq) (res *CreateMarginResp, err error)
	UpdateMargin(ctx context.Context, req *UpdateMarginReq) (res *UpdateMarginResp, err error)
	GetMargin(ctx context.Context, req *GetMarginReq) (res *GetMarginResp, err error)
	DeleteMargin(ctx context.Context, req *DeleteMarginReq) (res *DeleteMarginResp, err error)
	GetFullAndBaseInterfaceForCheck(ctx context.Context, req *GetFullAndBaseInterfaceForCheckReq) (res *GetFullAndBaseInterfaceForCheckResp, err error)
	CreateGradient(ctx context.Context, req *CreateGradientReq) (res *CreateGradientResp, err error)
	UpdateGradient(ctx context.Context, req *UpdateGradientReq) (res *UpdateGradientResp, err error)
	GetGradient(ctx context.Context, req *GetGradientReq) (res *GetGradientResp, err error)
	GetAmount(ctx context.Context, req *GetAmountReq) (res *GetAmountResp, err error)
	CreateLog(ctx context.Context, req *CreateLogReq) (res *CreateLogResp, err error)
	GetLog(ctx context.Context, req *GetLogReq) (res *GetLogResp, err error)
	GetAccountByTxId(ctx context.Context, req *GetAccountByTxIdReq) (res *GetAccountByTxIdResp, err error)
}
