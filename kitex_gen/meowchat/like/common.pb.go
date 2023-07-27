// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: meowchat/like/common.proto

package like

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

type Like struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetId     string `protobuf:"bytes,1,opt,name=targetId,proto3" json:"targetId,omitempty"`
	AssociatedId string `protobuf:"bytes,2,opt,name=associatedId,proto3" json:"associatedId,omitempty"`
}

func (x *Like) Reset() {
	*x = Like{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_like_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Like) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Like) ProtoMessage() {}

func (x *Like) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_like_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Like.ProtoReflect.Descriptor instead.
func (*Like) Descriptor() ([]byte, []int) {
	return file_meowchat_like_common_proto_rawDescGZIP(), []int{0}
}

func (x *Like) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *Like) GetAssociatedId() string {
	if x != nil {
		return x.AssociatedId
	}
	return ""
}

type ItemScore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Score int64  `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *ItemScore) Reset() {
	*x = ItemScore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_like_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemScore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemScore) ProtoMessage() {}

func (x *ItemScore) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_like_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemScore.ProtoReflect.Descriptor instead.
func (*ItemScore) Descriptor() ([]byte, []int) {
	return file_meowchat_like_common_proto_rawDescGZIP(), []int{1}
}

func (x *ItemScore) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ItemScore) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

type CatPop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CatId      string `protobuf:"bytes,1,opt,name=catId,proto3" json:"catId,omitempty"`
	Popularity int64  `protobuf:"varint,2,opt,name=popularity,proto3" json:"popularity,omitempty"`
}

func (x *CatPop) Reset() {
	*x = CatPop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_like_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CatPop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CatPop) ProtoMessage() {}

func (x *CatPop) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_like_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CatPop.ProtoReflect.Descriptor instead.
func (*CatPop) Descriptor() ([]byte, []int) {
	return file_meowchat_like_common_proto_rawDescGZIP(), []int{2}
}

func (x *CatPop) GetCatId() string {
	if x != nil {
		return x.CatId
	}
	return ""
}

func (x *CatPop) GetPopularity() int64 {
	if x != nil {
		return x.Popularity
	}
	return 0
}

var File_meowchat_like_common_proto protoreflect.FileDescriptor

var file_meowchat_like_common_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x6c, 0x69, 0x6b, 0x65, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x65,
	0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x22, 0x46, 0x0a, 0x04, 0x4c,
	0x69, 0x6b, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12,
	0x22, 0x0a, 0x0c, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65,
	0x64, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x09, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x3e, 0x0a, 0x06, 0x43, 0x61, 0x74, 0x50, 0x6f, 0x70,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x63, 0x61, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x6f, 0x70, 0x75,
	0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x68, 0x2d, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x69, 0x64, 0x6c, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x67, 0x6f, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x65, 0x6f,
	0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x6c, 0x69, 0x6b, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_meowchat_like_common_proto_rawDescOnce sync.Once
	file_meowchat_like_common_proto_rawDescData = file_meowchat_like_common_proto_rawDesc
)

func file_meowchat_like_common_proto_rawDescGZIP() []byte {
	file_meowchat_like_common_proto_rawDescOnce.Do(func() {
		file_meowchat_like_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_meowchat_like_common_proto_rawDescData)
	})
	return file_meowchat_like_common_proto_rawDescData
}

var file_meowchat_like_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_meowchat_like_common_proto_goTypes = []interface{}{
	(*Like)(nil),      // 0: meowchat.like.Like
	(*ItemScore)(nil), // 1: meowchat.like.ItemScore
	(*CatPop)(nil),    // 2: meowchat.like.CatPop
}
var file_meowchat_like_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_meowchat_like_common_proto_init() }
func file_meowchat_like_common_proto_init() {
	if File_meowchat_like_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_meowchat_like_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Like); i {
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
		file_meowchat_like_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemScore); i {
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
		file_meowchat_like_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CatPop); i {
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
			RawDescriptor: file_meowchat_like_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meowchat_like_common_proto_goTypes,
		DependencyIndexes: file_meowchat_like_common_proto_depIdxs,
		MessageInfos:      file_meowchat_like_common_proto_msgTypes,
	}.Build()
	File_meowchat_like_common_proto = out.File
	file_meowchat_like_common_proto_rawDesc = nil
	file_meowchat_like_common_proto_goTypes = nil
	file_meowchat_like_common_proto_depIdxs = nil
}

var _ context.Context
