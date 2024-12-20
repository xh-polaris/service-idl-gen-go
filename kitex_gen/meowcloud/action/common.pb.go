// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.19.4
// source: meowcloud/action/common.proto

package action

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

// 目标类型
type TargetType int32

const (
	TargetType_PHOTO   TargetType = 0
	TargetType_ALBUM   TargetType = 1
	TargetType_COMMENT TargetType = 2
	TargetType_USER    TargetType = 3
)

// Enum value maps for TargetType.
var (
	TargetType_name = map[int32]string{
		0: "PHOTO",
		1: "ALBUM",
		2: "COMMENT",
		3: "USER",
	}
	TargetType_value = map[string]int32{
		"PHOTO":   0,
		"ALBUM":   1,
		"COMMENT": 2,
		"USER":    3,
	}
)

func (x TargetType) Enum() *TargetType {
	p := new(TargetType)
	*p = x
	return p
}

func (x TargetType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TargetType) Descriptor() protoreflect.EnumDescriptor {
	return file_meowcloud_action_common_proto_enumTypes[0].Descriptor()
}

func (TargetType) Type() protoreflect.EnumType {
	return &file_meowcloud_action_common_proto_enumTypes[0]
}

func (x TargetType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TargetType.Descriptor instead.
func (TargetType) EnumDescriptor() ([]byte, []int) {
	return file_meowcloud_action_common_proto_rawDescGZIP(), []int{0}
}

type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowcloud_action_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_meowcloud_action_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_meowcloud_action_common_proto_rawDescGZIP(), []int{0}
}

// 点赞
type Action_Like struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TargetId   string     `protobuf:"bytes,2,opt,name=targetId,proto3" json:"targetId,omitempty"`
	TargetType TargetType `protobuf:"varint,3,opt,name=targetType,proto3,enum=meowcloud.action.TargetType" json:"targetType,omitempty"`
	UserId     string     `protobuf:"bytes,4,opt,name=userId,proto3" json:"userId,omitempty"`
	CreateAt   int64      `protobuf:"varint,5,opt,name=createAt,proto3" json:"createAt,omitempty"`
	IsCancel   bool       `protobuf:"varint,6,opt,name=isCancel,proto3" json:"isCancel,omitempty"`
}

func (x *Action_Like) Reset() {
	*x = Action_Like{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowcloud_action_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action_Like) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action_Like) ProtoMessage() {}

func (x *Action_Like) ProtoReflect() protoreflect.Message {
	mi := &file_meowcloud_action_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action_Like.ProtoReflect.Descriptor instead.
func (*Action_Like) Descriptor() ([]byte, []int) {
	return file_meowcloud_action_common_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Action_Like) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Action_Like) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *Action_Like) GetTargetType() TargetType {
	if x != nil {
		return x.TargetType
	}
	return TargetType_PHOTO
}

func (x *Action_Like) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Action_Like) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Action_Like) GetIsCancel() bool {
	if x != nil {
		return x.IsCancel
	}
	return false
}

// 分享
type Action_Share struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TargetId   string     `protobuf:"bytes,2,opt,name=targetId,proto3" json:"targetId,omitempty"`
	TargetType TargetType `protobuf:"varint,3,opt,name=targetType,proto3,enum=meowcloud.action.TargetType" json:"targetType,omitempty"`
	UserId     string     `protobuf:"bytes,4,opt,name=userId,proto3" json:"userId,omitempty"`
	CreateAt   int64      `protobuf:"varint,5,opt,name=createAt,proto3" json:"createAt,omitempty"`
}

func (x *Action_Share) Reset() {
	*x = Action_Share{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowcloud_action_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action_Share) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action_Share) ProtoMessage() {}

func (x *Action_Share) ProtoReflect() protoreflect.Message {
	mi := &file_meowcloud_action_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action_Share.ProtoReflect.Descriptor instead.
func (*Action_Share) Descriptor() ([]byte, []int) {
	return file_meowcloud_action_common_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Action_Share) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Action_Share) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *Action_Share) GetTargetType() TargetType {
	if x != nil {
		return x.TargetType
	}
	return TargetType_PHOTO
}

func (x *Action_Share) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Action_Share) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

// 关注
type Action_Follow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TargetId   string     `protobuf:"bytes,2,opt,name=targetId,proto3" json:"targetId,omitempty"`
	TargetType TargetType `protobuf:"varint,3,opt,name=targetType,proto3,enum=meowcloud.action.TargetType" json:"targetType,omitempty"`
	UserId     string     `protobuf:"bytes,4,opt,name=userId,proto3" json:"userId,omitempty"`
	CreateAt   int64      `protobuf:"varint,5,opt,name=createAt,proto3" json:"createAt,omitempty"`
	IsCancel   bool       `protobuf:"varint,6,opt,name=isCancel,proto3" json:"isCancel,omitempty"`
}

func (x *Action_Follow) Reset() {
	*x = Action_Follow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowcloud_action_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action_Follow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action_Follow) ProtoMessage() {}

func (x *Action_Follow) ProtoReflect() protoreflect.Message {
	mi := &file_meowcloud_action_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action_Follow.ProtoReflect.Descriptor instead.
func (*Action_Follow) Descriptor() ([]byte, []int) {
	return file_meowcloud_action_common_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Action_Follow) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Action_Follow) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *Action_Follow) GetTargetType() TargetType {
	if x != nil {
		return x.TargetType
	}
	return TargetType_PHOTO
}

func (x *Action_Follow) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Action_Follow) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Action_Follow) GetIsCancel() bool {
	if x != nil {
		return x.IsCancel
	}
	return false
}

var File_meowcloud_action_common_proto protoreflect.FileDescriptor

var file_meowcloud_action_common_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0xb8, 0x04, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0xc0, 0x01, 0x0a,
	0x04, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49,
	0x64, 0x12, 0x3c, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x1a,
	0xa5, 0x01, 0x0a, 0x05, 0x53, 0x68, 0x61, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6d, 0x65, 0x6f, 0x77,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x1a, 0xc2, 0x01, 0x0a, 0x06, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x3c,
	0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x2a, 0x39, 0x0a, 0x0a,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x48,
	0x4f, 0x54, 0x4f, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c, 0x42, 0x55, 0x4d, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x08, 0x0a,
	0x04, 0x55, 0x53, 0x45, 0x52, 0x10, 0x03, 0x42, 0x7b, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x78,
	0x68, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2e, 0x69, 0x64, 0x6c, 0x67, 0x65, 0x6e, 0x2e,
	0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x68, 0x2d, 0x70,
	0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x69,
	0x64, 0x6c, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f,
	0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meowcloud_action_common_proto_rawDescOnce sync.Once
	file_meowcloud_action_common_proto_rawDescData = file_meowcloud_action_common_proto_rawDesc
)

func file_meowcloud_action_common_proto_rawDescGZIP() []byte {
	file_meowcloud_action_common_proto_rawDescOnce.Do(func() {
		file_meowcloud_action_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_meowcloud_action_common_proto_rawDescData)
	})
	return file_meowcloud_action_common_proto_rawDescData
}

var file_meowcloud_action_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_meowcloud_action_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_meowcloud_action_common_proto_goTypes = []interface{}{
	(TargetType)(0),       // 0: meowcloud.action.TargetType
	(*Action)(nil),        // 1: meowcloud.action.Action
	(*Action_Like)(nil),   // 2: meowcloud.action.Action.Like
	(*Action_Share)(nil),  // 3: meowcloud.action.Action.Share
	(*Action_Follow)(nil), // 4: meowcloud.action.Action.Follow
}
var file_meowcloud_action_common_proto_depIdxs = []int32{
	0, // 0: meowcloud.action.Action.Like.targetType:type_name -> meowcloud.action.TargetType
	0, // 1: meowcloud.action.Action.Share.targetType:type_name -> meowcloud.action.TargetType
	0, // 2: meowcloud.action.Action.Follow.targetType:type_name -> meowcloud.action.TargetType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_meowcloud_action_common_proto_init() }
func file_meowcloud_action_common_proto_init() {
	if File_meowcloud_action_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_meowcloud_action_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action); i {
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
		file_meowcloud_action_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action_Like); i {
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
		file_meowcloud_action_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action_Share); i {
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
		file_meowcloud_action_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action_Follow); i {
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
			RawDescriptor: file_meowcloud_action_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meowcloud_action_common_proto_goTypes,
		DependencyIndexes: file_meowcloud_action_common_proto_depIdxs,
		EnumInfos:         file_meowcloud_action_common_proto_enumTypes,
		MessageInfos:      file_meowcloud_action_common_proto_msgTypes,
	}.Build()
	File_meowcloud_action_common_proto = out.File
	file_meowcloud_action_common_proto_rawDesc = nil
	file_meowcloud_action_common_proto_goTypes = nil
	file_meowcloud_action_common_proto_depIdxs = nil
}

var _ context.Context
