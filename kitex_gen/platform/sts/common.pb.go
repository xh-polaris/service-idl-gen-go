// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: platform/sts/common.proto

package sts

import (
	context "context"
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

type Scene int32

const (
	Scene_Unknown Scene = 0
	Scene_Profile Scene = 1
	Scene_Comment Scene = 2
	Scene_Forum   Scene = 3
)

// Enum value maps for Scene.
var (
	Scene_name = map[int32]string{
		0: "Unknown",
		1: "Profile",
		2: "Comment",
		3: "Forum",
	}
	Scene_value = map[string]int32{
		"Unknown": 0,
		"Profile": 1,
		"Comment": 2,
		"Forum":   3,
	}
)

func (x Scene) Enum() *Scene {
	p := new(Scene)
	*p = x
	return p
}

func (x Scene) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Scene) Descriptor() protoreflect.EnumDescriptor {
	return file_platform_sts_common_proto_enumTypes[0].Descriptor()
}

func (Scene) Type() protoreflect.EnumType {
	return &file_platform_sts_common_proto_enumTypes[0]
}

func (x Scene) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Scene.Descriptor instead.
func (Scene) EnumDescriptor() ([]byte, []int) {
	return file_platform_sts_common_proto_rawDescGZIP(), []int{0}
}

var File_platform_sts_common_proto protoreflect.FileDescriptor

var file_platform_sts_common_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x73, 0x74, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x74, 0x73, 0x2a, 0x39, 0x0a, 0x05, 0x53, 0x63, 0x65,
	0x6e, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x6f, 0x72,
	0x75, 0x6d, 0x10, 0x03, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x78, 0x68, 0x2d, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x69, 0x64, 0x6c, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f,
	0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x73, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_platform_sts_common_proto_rawDescOnce sync.Once
	file_platform_sts_common_proto_rawDescData = file_platform_sts_common_proto_rawDesc
)

func file_platform_sts_common_proto_rawDescGZIP() []byte {
	file_platform_sts_common_proto_rawDescOnce.Do(func() {
		file_platform_sts_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_platform_sts_common_proto_rawDescData)
	})
	return file_platform_sts_common_proto_rawDescData
}

var file_platform_sts_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_platform_sts_common_proto_goTypes = []interface{}{
	(Scene)(0), // 0: platform.sts.Scene
}
var file_platform_sts_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_platform_sts_common_proto_init() }
func file_platform_sts_common_proto_init() {
	if File_platform_sts_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_platform_sts_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_platform_sts_common_proto_goTypes,
		DependencyIndexes: file_platform_sts_common_proto_depIdxs,
		EnumInfos:         file_platform_sts_common_proto_enumTypes,
	}.Build()
	File_platform_sts_common_proto = out.File
	file_platform_sts_common_proto_rawDesc = nil
	file_platform_sts_common_proto_goTypes = nil
	file_platform_sts_common_proto_depIdxs = nil
}

var _ context.Context
