// 该文件定义了用户行为服务

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: meowcloud/action/action.proto

package action

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

var File_meowcloud_action_action_proto protoreflect.FileDescriptor

var file_meowcloud_action_action_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x0f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x6d,
	0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x6d, 0x65,
	0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xfb, 0x04, 0x0a, 0x04,
	0x6c, 0x69, 0x6b, 0x65, 0x12, 0x56, 0x0a, 0x06, 0x44, 0x6f, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x1b,
	0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x44, 0x6f, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x6d, 0x65,
	0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44,
	0x6f, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x11, 0xd2, 0xc1, 0x18, 0x0d, 0x2f,
	0x6c, 0x69, 0x6b, 0x65, 0x2f, 0x64, 0x6f, 0x5f, 0x6c, 0x69, 0x6b, 0x65, 0x12, 0x66, 0x0a, 0x0a,
	0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x1f, 0x2e, 0x6d, 0x65, 0x6f,
	0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x6d, 0x65,
	0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x15, 0xd2,
	0xc1, 0x18, 0x11, 0x2f, 0x6c, 0x69, 0x6b, 0x65, 0x2f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x5f,
	0x6c, 0x69, 0x6b, 0x65, 0x12, 0x6d, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x64,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6b, 0x65,
	0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x6d, 0x65, 0x6f, 0x77,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x6b, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x13,
	0xca, 0xc1, 0x18, 0x0f, 0x2f, 0x6c, 0x69, 0x6b, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x73, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x12, 0x22, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x6b, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x19, 0xca,
	0xc1, 0x18, 0x15, 0x2f, 0x6c, 0x69, 0x6b, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x6c, 0x69, 0x6b,
	0x65, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x6f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x12, 0x21, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x6d, 0x65,
	0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x18, 0xca, 0xc1, 0x18, 0x14, 0x2f, 0x6c, 0x69, 0x6b, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6b, 0x65, 0x64, 0x12, 0x5e, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x6b, 0x65, 0x64, 0x12, 0x1d, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6b, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x13, 0xca, 0xc1, 0x18, 0x0f, 0x2f, 0x6c, 0x69, 0x6b, 0x65, 0x2f,
	0x67, 0x65, 0x74, 0x5f, 0x6c, 0x69, 0x6b, 0x65, 0x64, 0x32, 0xac, 0x04, 0x0a, 0x05, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x12, 0x5b, 0x0a, 0x07, 0x44, 0x6f, 0x53, 0x68, 0x61, 0x72, 0x65, 0x12, 0x1c,
	0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x44, 0x6f, 0x53, 0x68, 0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x6d,
	0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x44, 0x6f, 0x53, 0x68, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x13, 0xd2, 0xc1, 0x18,
	0x0f, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2f, 0x64, 0x6f, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x12, 0x71, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x23, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x14, 0xca,
	0xc1, 0x18, 0x10, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x78, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x23, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e, 0x6d, 0x65, 0x6f,
	0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x1b, 0xca, 0xc1, 0x18, 0x17, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2f, 0x67, 0x65, 0x74,
	0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x74, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x12, 0x22,
	0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x52,
	0x65, 0x71, 0x1a, 0x23, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1a, 0xca, 0xc1, 0x18, 0x16, 0x2f, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x12, 0x63, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x12, 0x1e, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x1a, 0x1f, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x15, 0xca, 0xc1, 0x18, 0x11, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2f, 0x67, 0x65,
	0x74, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x32, 0xc7, 0x05, 0x0a, 0x06, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x12, 0x60, 0x0a, 0x08, 0x44, 0x6f, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12,
	0x1d, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x44, 0x6f, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x1a, 0x1e,
	0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x44, 0x6f, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x22, 0x15,
	0xd2, 0xc1, 0x18, 0x11, 0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2f, 0x64, 0x6f, 0x5f, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x70, 0x0a, 0x0c, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x21, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x22, 0x19, 0xd2, 0xc1,
	0x18, 0x15, 0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x78, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x2e, 0x6d, 0x65,
	0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47,
	0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x26, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x15, 0xca, 0xc1, 0x18, 0x11,
	0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x81, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x25, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x26, 0x2e,
	0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1e, 0xca, 0xc1, 0x18, 0x1a, 0x2f, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x7d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x12, 0x24, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x25,
	0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1d, 0xca, 0xc1, 0x18, 0x19, 0x2f, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x64, 0x12, 0x6c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x64, 0x12, 0x20, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x18, 0xca, 0xc1, 0x18, 0x14, 0x2f, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x42, 0x7b, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x68, 0x70, 0x6f, 0x6c, 0x61,
	0x72, 0x69, 0x73, 0x2e, 0x69, 0x64, 0x6c, 0x67, 0x65, 0x6e, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0b, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x68, 0x2d, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69,
	0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x69, 0x64, 0x6c, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x6d,
	0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_meowcloud_action_action_proto_goTypes = []interface{}{
	(*DoLikeReq)(nil),            // 0: meowcloud.action.DoLikeReq
	(*CancelLikeReq)(nil),        // 1: meowcloud.action.CancelLikeReq
	(*GetLikedCountReq)(nil),     // 2: meowcloud.action.GetLikedCountReq
	(*GetLikedUsersReq)(nil),     // 3: meowcloud.action.GetLikedUsersReq
	(*GetUserLikedReq)(nil),      // 4: meowcloud.action.GetUserLikedReq
	(*GetLikedReq)(nil),          // 5: meowcloud.action.GetLikedReq
	(*DoShareReq)(nil),           // 6: meowcloud.action.DoShareReq
	(*GetSharedCountReq)(nil),    // 7: meowcloud.action.GetSharedCountReq
	(*GetSharedUsersReq)(nil),    // 8: meowcloud.action.GetSharedUsersReq
	(*GetUserSharedReq)(nil),     // 9: meowcloud.action.GetUserSharedReq
	(*GetSharedReq)(nil),         // 10: meowcloud.action.GetSharedReq
	(*DoFollowReq)(nil),          // 11: meowcloud.action.DoFollowReq
	(*CancelFollowReq)(nil),      // 12: meowcloud.action.CancelFollowReq
	(*GetFollowedCountReq)(nil),  // 13: meowcloud.action.GetFollowedCountReq
	(*GetFollowedUsersReq)(nil),  // 14: meowcloud.action.GetFollowedUsersReq
	(*GetUserFollowedReq)(nil),   // 15: meowcloud.action.GetUserFollowedReq
	(*GetFollowedReq)(nil),       // 16: meowcloud.action.GetFollowedReq
	(*DoLikeResp)(nil),           // 17: meowcloud.action.DoLikeResp
	(*CancelLikeResp)(nil),       // 18: meowcloud.action.CancelLikeResp
	(*GetLikedCountResp)(nil),    // 19: meowcloud.action.GetLikedCountResp
	(*GetLikedUsersResp)(nil),    // 20: meowcloud.action.GetLikedUsersResp
	(*GetUserLikedResp)(nil),     // 21: meowcloud.action.GetUserLikedResp
	(*GetLikedResp)(nil),         // 22: meowcloud.action.GetLikedResp
	(*DoShareResp)(nil),          // 23: meowcloud.action.DoShareResp
	(*GetSharedCountResp)(nil),   // 24: meowcloud.action.GetSharedCountResp
	(*GetSharedUsersResp)(nil),   // 25: meowcloud.action.GetSharedUsersResp
	(*GetUserSharedResp)(nil),    // 26: meowcloud.action.GetUserSharedResp
	(*GetSharedResp)(nil),        // 27: meowcloud.action.GetSharedResp
	(*DoFollowResp)(nil),         // 28: meowcloud.action.DoFollowResp
	(*CancelFollowResp)(nil),     // 29: meowcloud.action.CancelFollowResp
	(*GetFollowedCountResp)(nil), // 30: meowcloud.action.GetFollowedCountResp
	(*GetFollowedUsersResp)(nil), // 31: meowcloud.action.GetFollowedUsersResp
	(*GetUserFollowedResp)(nil),  // 32: meowcloud.action.GetUserFollowedResp
	(*GetFollowedResp)(nil),      // 33: meowcloud.action.GetFollowedResp
}
var file_meowcloud_action_action_proto_depIdxs = []int32{
	0,  // 0: meowcloud.action.like.DoLike:input_type -> meowcloud.action.DoLikeReq
	1,  // 1: meowcloud.action.like.CancelLike:input_type -> meowcloud.action.CancelLikeReq
	2,  // 2: meowcloud.action.like.GetLikedCount:input_type -> meowcloud.action.GetLikedCountReq
	3,  // 3: meowcloud.action.like.GetLikedUsers:input_type -> meowcloud.action.GetLikedUsersReq
	4,  // 4: meowcloud.action.like.GetUserLiked:input_type -> meowcloud.action.GetUserLikedReq
	5,  // 5: meowcloud.action.like.GetLiked:input_type -> meowcloud.action.GetLikedReq
	6,  // 6: meowcloud.action.share.DoShare:input_type -> meowcloud.action.DoShareReq
	7,  // 7: meowcloud.action.share.GetSharedCount:input_type -> meowcloud.action.GetSharedCountReq
	8,  // 8: meowcloud.action.share.GetSharedUsers:input_type -> meowcloud.action.GetSharedUsersReq
	9,  // 9: meowcloud.action.share.GetUserShared:input_type -> meowcloud.action.GetUserSharedReq
	10, // 10: meowcloud.action.share.GetShared:input_type -> meowcloud.action.GetSharedReq
	11, // 11: meowcloud.action.follow.DoFollow:input_type -> meowcloud.action.DoFollowReq
	12, // 12: meowcloud.action.follow.CancelFollow:input_type -> meowcloud.action.CancelFollowReq
	13, // 13: meowcloud.action.follow.GetFollowedCount:input_type -> meowcloud.action.GetFollowedCountReq
	14, // 14: meowcloud.action.follow.GetFollowedUsers:input_type -> meowcloud.action.GetFollowedUsersReq
	15, // 15: meowcloud.action.follow.GetUserFollowed:input_type -> meowcloud.action.GetUserFollowedReq
	16, // 16: meowcloud.action.follow.GetFollowed:input_type -> meowcloud.action.GetFollowedReq
	17, // 17: meowcloud.action.like.DoLike:output_type -> meowcloud.action.DoLikeResp
	18, // 18: meowcloud.action.like.CancelLike:output_type -> meowcloud.action.CancelLikeResp
	19, // 19: meowcloud.action.like.GetLikedCount:output_type -> meowcloud.action.GetLikedCountResp
	20, // 20: meowcloud.action.like.GetLikedUsers:output_type -> meowcloud.action.GetLikedUsersResp
	21, // 21: meowcloud.action.like.GetUserLiked:output_type -> meowcloud.action.GetUserLikedResp
	22, // 22: meowcloud.action.like.GetLiked:output_type -> meowcloud.action.GetLikedResp
	23, // 23: meowcloud.action.share.DoShare:output_type -> meowcloud.action.DoShareResp
	24, // 24: meowcloud.action.share.GetSharedCount:output_type -> meowcloud.action.GetSharedCountResp
	25, // 25: meowcloud.action.share.GetSharedUsers:output_type -> meowcloud.action.GetSharedUsersResp
	26, // 26: meowcloud.action.share.GetUserShared:output_type -> meowcloud.action.GetUserSharedResp
	27, // 27: meowcloud.action.share.GetShared:output_type -> meowcloud.action.GetSharedResp
	28, // 28: meowcloud.action.follow.DoFollow:output_type -> meowcloud.action.DoFollowResp
	29, // 29: meowcloud.action.follow.CancelFollow:output_type -> meowcloud.action.CancelFollowResp
	30, // 30: meowcloud.action.follow.GetFollowedCount:output_type -> meowcloud.action.GetFollowedCountResp
	31, // 31: meowcloud.action.follow.GetFollowedUsers:output_type -> meowcloud.action.GetFollowedUsersResp
	32, // 32: meowcloud.action.follow.GetUserFollowed:output_type -> meowcloud.action.GetUserFollowedResp
	33, // 33: meowcloud.action.follow.GetFollowed:output_type -> meowcloud.action.GetFollowedResp
	17, // [17:34] is the sub-list for method output_type
	0,  // [0:17] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_meowcloud_action_action_proto_init() }
func file_meowcloud_action_action_proto_init() {
	if File_meowcloud_action_action_proto != nil {
		return
	}
	file_meowcloud_action_like_proto_init()
	file_meowcloud_action_share_proto_init()
	file_meowcloud_action_follow_proto_init()
	file_meowcloud_action_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_meowcloud_action_action_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_meowcloud_action_action_proto_goTypes,
		DependencyIndexes: file_meowcloud_action_action_proto_depIdxs,
	}.Build()
	File_meowcloud_action_action_proto = out.File
	file_meowcloud_action_action_proto_rawDesc = nil
	file_meowcloud_action_action_proto_goTypes = nil
	file_meowcloud_action_action_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.10.3. DO NOT EDIT.

type Like interface {
	DoLike(ctx context.Context, req *DoLikeReq) (res *DoLikeResp, err error)
	CancelLike(ctx context.Context, req *CancelLikeReq) (res *CancelLikeResp, err error)
	GetLikedCount(ctx context.Context, req *GetLikedCountReq) (res *GetLikedCountResp, err error)
	GetLikedUsers(ctx context.Context, req *GetLikedUsersReq) (res *GetLikedUsersResp, err error)
	GetUserLiked(ctx context.Context, req *GetUserLikedReq) (res *GetUserLikedResp, err error)
	GetLiked(ctx context.Context, req *GetLikedReq) (res *GetLikedResp, err error)
}

type Share interface {
	DoShare(ctx context.Context, req *DoShareReq) (res *DoShareResp, err error)
	GetSharedCount(ctx context.Context, req *GetSharedCountReq) (res *GetSharedCountResp, err error)
	GetSharedUsers(ctx context.Context, req *GetSharedUsersReq) (res *GetSharedUsersResp, err error)
	GetUserShared(ctx context.Context, req *GetUserSharedReq) (res *GetUserSharedResp, err error)
	GetShared(ctx context.Context, req *GetSharedReq) (res *GetSharedResp, err error)
}

type Follow interface {
	DoFollow(ctx context.Context, req *DoFollowReq) (res *DoFollowResp, err error)
	CancelFollow(ctx context.Context, req *CancelFollowReq) (res *CancelFollowResp, err error)
	GetFollowedCount(ctx context.Context, req *GetFollowedCountReq) (res *GetFollowedCountResp, err error)
	GetFollowedUsers(ctx context.Context, req *GetFollowedUsersReq) (res *GetFollowedUsersResp, err error)
	GetUserFollowed(ctx context.Context, req *GetUserFollowedReq) (res *GetUserFollowedResp, err error)
	GetFollowed(ctx context.Context, req *GetFollowedReq) (res *GetFollowedResp, err error)
}
