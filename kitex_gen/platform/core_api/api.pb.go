// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: platform/api/api.proto

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

var File_platform_api_api_proto protoreflect.FileDescriptor

var file_platform_api_api_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x0f, 0x68, 0x74, 0x74,
	0x70, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x80,
	0x01, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x78, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0xd2, 0xc1, 0x18, 0x16, 0x2f, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x32, 0xc9, 0x02, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x58, 0x0a, 0x06, 0x53, 0x69,
	0x67, 0x6e, 0x49, 0x6e, 0x12, 0x1c, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52,
	0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x11, 0xd2, 0xc1, 0x18, 0x0d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x69, 0x67,
	0x6e, 0x5f, 0x69, 0x6e, 0x12, 0x79, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x1a, 0xd2, 0xc1, 0x18, 0x16, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73,
	0x65, 0x6e, 0x64, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x6c, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21,
	0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x1a, 0x22, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x22, 0x16, 0xd2, 0xc1, 0x18, 0x12, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x7e, 0x0a,
	0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x68, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2e, 0x69,
	0x64, 0x6c, 0x67, 0x65, 0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x42, 0x0c, 0x43, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x68, 0x2d, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x69, 0x64, 0x6c, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67,
	0x6f, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_platform_api_api_proto_goTypes = []interface{}{
	(*ReportEventRequest)(nil),  // 0: platform.core_api.ReportEventRequest
	(*SignInReq)(nil),           // 1: platform.core_api.SignInReq
	(*SendVerifyCodeReq)(nil),   // 2: platform.core_api.SendVerifyCodeReq
	(*SetPasswordReq)(nil),      // 3: platform.core_api.SetPasswordReq
	(*ReportEventResponse)(nil), // 4: platform.core_api.ReportEventResponse
	(*SignInResp)(nil),          // 5: platform.core_api.SignInResp
	(*SendVerifyCodeResp)(nil),  // 6: platform.core_api.SendVerifyCodeResp
	(*SetPasswordResp)(nil),     // 7: platform.core_api.SetPasswordResp
}
var file_platform_api_api_proto_depIdxs = []int32{
	0, // 0: platform.core_api.data.ReportEvent:input_type -> platform.core_api.ReportEventRequest
	1, // 1: platform.core_api.auth.SignIn:input_type -> platform.core_api.SignInReq
	2, // 2: platform.core_api.auth.SendVerifyCode:input_type -> platform.core_api.SendVerifyCodeReq
	3, // 3: platform.core_api.auth.SetPassword:input_type -> platform.core_api.SetPasswordReq
	4, // 4: platform.core_api.data.ReportEvent:output_type -> platform.core_api.ReportEventResponse
	5, // 5: platform.core_api.auth.SignIn:output_type -> platform.core_api.SignInResp
	6, // 6: platform.core_api.auth.SendVerifyCode:output_type -> platform.core_api.SendVerifyCodeResp
	7, // 7: platform.core_api.auth.SetPassword:output_type -> platform.core_api.SetPasswordResp
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_platform_api_api_proto_init() }
func file_platform_api_api_proto_init() {
	if File_platform_api_api_proto != nil {
		return
	}
	file_platform_api_data_proto_init()
	file_platform_api_auth_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_platform_api_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_platform_api_api_proto_goTypes,
		DependencyIndexes: file_platform_api_api_proto_depIdxs,
	}.Build()
	File_platform_api_api_proto = out.File
	file_platform_api_api_proto_rawDesc = nil
	file_platform_api_api_proto_goTypes = nil
	file_platform_api_api_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.11.3. DO NOT EDIT.

type Data interface {
	ReportEvent(ctx context.Context, req *ReportEventRequest) (res *ReportEventResponse, err error)
}

type Auth interface {
	SignIn(ctx context.Context, req *SignInReq) (res *SignInResp, err error)
	SendVerifyCode(ctx context.Context, req *SendVerifyCodeReq) (res *SendVerifyCodeResp, err error)
	SetPassword(ctx context.Context, req *SetPasswordReq) (res *SetPasswordResp, err error)
}
