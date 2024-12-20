// 该文件中定义了分享服务需要使用的message

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.19.4
// source: meowcloud/action/share.proto

package action

import (
	context "context"
	basic "github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
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

// 分享请求
type DoShareReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetId   string          `protobuf:"bytes,1,opt,name=targetId,proto3" json:"targetId,omitempty"`                                       // 分享目标id
	TargetType TargetType      `protobuf:"varint,2,opt,name=targetType,proto3,enum=meowcloud.action.TargetType" json:"targetType,omitempty"` // 分享类型(0相片，1相册，2评论)
	User       *basic.UserMeta `protobuf:"bytes,254,opt,name=user,proto3" json:"user,omitempty"`                                             // 用户信息
}

func (x *DoShareReq) Reset() {
	*x = DoShareReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowcloud_action_share_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoShareReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoShareReq) ProtoMessage() {}

func (x *DoShareReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowcloud_action_share_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoShareReq.ProtoReflect.Descriptor instead.
func (*DoShareReq) Descriptor() ([]byte, []int) {
	return file_meowcloud_action_share_proto_rawDescGZIP(), []int{0}
}

func (x *DoShareReq) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *DoShareReq) GetTargetType() TargetType {
	if x != nil {
		return x.TargetType
	}
	return TargetType_PHOTO
}

func (x *DoShareReq) GetUser() *basic.UserMeta {
	if x != nil {
		return x.User
	}
	return nil
}

// 分享响应
type DoShareResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DoShareResp) Reset() {
	*x = DoShareResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowcloud_action_share_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoShareResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoShareResp) ProtoMessage() {}

func (x *DoShareResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowcloud_action_share_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoShareResp.ProtoReflect.Descriptor instead.
func (*DoShareResp) Descriptor() ([]byte, []int) {
	return file_meowcloud_action_share_proto_rawDescGZIP(), []int{1}
}

// 获取分享数请求
type GetSharedCountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetId   string     `protobuf:"bytes,1,opt,name=targetId,proto3" json:"targetId,omitempty"`
	TargetType TargetType `protobuf:"varint,2,opt,name=targetType,proto3,enum=meowcloud.action.TargetType" json:"targetType,omitempty"`
}

func (x *GetSharedCountReq) Reset() {
	*x = GetSharedCountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowcloud_action_share_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSharedCountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSharedCountReq) ProtoMessage() {}

func (x *GetSharedCountReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowcloud_action_share_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSharedCountReq.ProtoReflect.Descriptor instead.
func (*GetSharedCountReq) Descriptor() ([]byte, []int) {
	return file_meowcloud_action_share_proto_rawDescGZIP(), []int{2}
}

func (x *GetSharedCountReq) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *GetSharedCountReq) GetTargetType() TargetType {
	if x != nil {
		return x.TargetType
	}
	return TargetType_PHOTO
}

// 获取分享数响应
type GetSharedCountResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetSharedCountResp) Reset() {
	*x = GetSharedCountResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowcloud_action_share_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSharedCountResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSharedCountResp) ProtoMessage() {}

func (x *GetSharedCountResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowcloud_action_share_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSharedCountResp.ProtoReflect.Descriptor instead.
func (*GetSharedCountResp) Descriptor() ([]byte, []int) {
	return file_meowcloud_action_share_proto_rawDescGZIP(), []int{3}
}

func (x *GetSharedCountResp) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

// 获取分享对象列表
type GetSharedUsersReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetId         string                   `protobuf:"bytes,1,opt,name=targetId,proto3" json:"targetId,omitempty"`
	TargetType       TargetType               `protobuf:"varint,2,opt,name=targetType,proto3,enum=meowcloud.action.TargetType" json:"targetType,omitempty"`
	PaginationOption *basic.PaginationOptions `protobuf:"bytes,3,opt,name=paginationOption,proto3" json:"paginationOption,omitempty"`
}

func (x *GetSharedUsersReq) Reset() {
	*x = GetSharedUsersReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowcloud_action_share_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSharedUsersReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSharedUsersReq) ProtoMessage() {}

func (x *GetSharedUsersReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowcloud_action_share_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSharedUsersReq.ProtoReflect.Descriptor instead.
func (*GetSharedUsersReq) Descriptor() ([]byte, []int) {
	return file_meowcloud_action_share_proto_rawDescGZIP(), []int{4}
}

func (x *GetSharedUsersReq) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *GetSharedUsersReq) GetTargetType() TargetType {
	if x != nil {
		return x.TargetType
	}
	return TargetType_PHOTO
}

func (x *GetSharedUsersReq) GetPaginationOption() *basic.PaginationOptions {
	if x != nil {
		return x.PaginationOption
	}
	return nil
}

// 获取分享对象列表响应
type GetSharedUsersResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shares []*Action_Share `protobuf:"bytes,1,rep,name=shares,proto3" json:"shares,omitempty"`
	Total  int64           `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetSharedUsersResp) Reset() {
	*x = GetSharedUsersResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowcloud_action_share_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSharedUsersResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSharedUsersResp) ProtoMessage() {}

func (x *GetSharedUsersResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowcloud_action_share_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSharedUsersResp.ProtoReflect.Descriptor instead.
func (*GetSharedUsersResp) Descriptor() ([]byte, []int) {
	return file_meowcloud_action_share_proto_rawDescGZIP(), []int{5}
}

func (x *GetSharedUsersResp) GetShares() []*Action_Share {
	if x != nil {
		return x.Shares
	}
	return nil
}

func (x *GetSharedUsersResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

// 获取用户分享过对象请求
type GetUserSharedReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetType       TargetType               `protobuf:"varint,1,opt,name=targetType,proto3,enum=meowcloud.action.TargetType" json:"targetType,omitempty"`
	PaginationOption *basic.PaginationOptions `protobuf:"bytes,2,opt,name=paginationOption,proto3" json:"paginationOption,omitempty"`
	User             *basic.UserMeta          `protobuf:"bytes,254,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserSharedReq) Reset() {
	*x = GetUserSharedReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowcloud_action_share_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserSharedReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserSharedReq) ProtoMessage() {}

func (x *GetUserSharedReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowcloud_action_share_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserSharedReq.ProtoReflect.Descriptor instead.
func (*GetUserSharedReq) Descriptor() ([]byte, []int) {
	return file_meowcloud_action_share_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserSharedReq) GetTargetType() TargetType {
	if x != nil {
		return x.TargetType
	}
	return TargetType_PHOTO
}

func (x *GetUserSharedReq) GetPaginationOption() *basic.PaginationOptions {
	if x != nil {
		return x.PaginationOption
	}
	return nil
}

func (x *GetUserSharedReq) GetUser() *basic.UserMeta {
	if x != nil {
		return x.User
	}
	return nil
}

// 获取用户分享过对象响应
type GetUserSharedResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shares []*Action_Share `protobuf:"bytes,1,rep,name=shares,proto3" json:"shares,omitempty"`
	Total  int64           `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetUserSharedResp) Reset() {
	*x = GetUserSharedResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowcloud_action_share_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserSharedResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserSharedResp) ProtoMessage() {}

func (x *GetUserSharedResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowcloud_action_share_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserSharedResp.ProtoReflect.Descriptor instead.
func (*GetUserSharedResp) Descriptor() ([]byte, []int) {
	return file_meowcloud_action_share_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserSharedResp) GetShares() []*Action_Share {
	if x != nil {
		return x.Shares
	}
	return nil
}

func (x *GetUserSharedResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

// 判断是否分享过
type GetSharedReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetId   string          `protobuf:"bytes,1,opt,name=targetId,proto3" json:"targetId,omitempty"`
	TargetType TargetType      `protobuf:"varint,2,opt,name=targetType,proto3,enum=meowcloud.action.TargetType" json:"targetType,omitempty"`
	User       *basic.UserMeta `protobuf:"bytes,254,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetSharedReq) Reset() {
	*x = GetSharedReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowcloud_action_share_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSharedReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSharedReq) ProtoMessage() {}

func (x *GetSharedReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowcloud_action_share_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSharedReq.ProtoReflect.Descriptor instead.
func (*GetSharedReq) Descriptor() ([]byte, []int) {
	return file_meowcloud_action_share_proto_rawDescGZIP(), []int{8}
}

func (x *GetSharedReq) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *GetSharedReq) GetTargetType() TargetType {
	if x != nil {
		return x.TargetType
	}
	return TargetType_PHOTO
}

func (x *GetSharedReq) GetUser() *basic.UserMeta {
	if x != nil {
		return x.User
	}
	return nil
}

type GetSharedResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shared bool `protobuf:"varint,1,opt,name=shared,proto3" json:"shared,omitempty"`
}

func (x *GetSharedResp) Reset() {
	*x = GetSharedResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowcloud_action_share_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSharedResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSharedResp) ProtoMessage() {}

func (x *GetSharedResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowcloud_action_share_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSharedResp.ProtoReflect.Descriptor instead.
func (*GetSharedResp) Descriptor() ([]byte, []int) {
	return file_meowcloud_action_share_proto_rawDescGZIP(), []int{9}
}

func (x *GetSharedResp) GetShared() bool {
	if x != nil {
		return x.Shared
	}
	return false
}

var File_meowcloud_action_share_proto protoreflect.FileDescriptor

var file_meowcloud_action_share_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x16, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x6d, 0x65, 0x6f, 0x77,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x01, 0x0a, 0x0a, 0x44, 0x6f,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0xfe, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65,
	0x74, 0x61, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x0d, 0x0a, 0x0b, 0x44, 0x6f, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x6d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6d,
	0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x2a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0xb3, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x44, 0x0a, 0x10, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62,
	0x61, 0x73, 0x69, 0x63, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x10, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x62, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x36,
	0x0a, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x52, 0x06,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xbc, 0x01, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x12, 0x3c, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x44, 0x0a, 0x10, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x10, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0xfe, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x61, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x36, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x52, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x8e,
	0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x52, 0x65, 0x71, 0x12,
	0x1a, 0x0a, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x0a, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1c, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0xfe, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0x27, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x42, 0x7a, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e,
	0x78, 0x68, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2e, 0x69, 0x64, 0x6c, 0x67, 0x65, 0x6e,
	0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x0a, 0x53, 0x68, 0x61, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x68, 0x2d, 0x70,
	0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x69,
	0x64, 0x6c, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f,
	0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meowcloud_action_share_proto_rawDescOnce sync.Once
	file_meowcloud_action_share_proto_rawDescData = file_meowcloud_action_share_proto_rawDesc
)

func file_meowcloud_action_share_proto_rawDescGZIP() []byte {
	file_meowcloud_action_share_proto_rawDescOnce.Do(func() {
		file_meowcloud_action_share_proto_rawDescData = protoimpl.X.CompressGZIP(file_meowcloud_action_share_proto_rawDescData)
	})
	return file_meowcloud_action_share_proto_rawDescData
}

var file_meowcloud_action_share_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_meowcloud_action_share_proto_goTypes = []interface{}{
	(*DoShareReq)(nil),              // 0: meowcloud.action.DoShareReq
	(*DoShareResp)(nil),             // 1: meowcloud.action.DoShareResp
	(*GetSharedCountReq)(nil),       // 2: meowcloud.action.GetSharedCountReq
	(*GetSharedCountResp)(nil),      // 3: meowcloud.action.GetSharedCountResp
	(*GetSharedUsersReq)(nil),       // 4: meowcloud.action.GetSharedUsersReq
	(*GetSharedUsersResp)(nil),      // 5: meowcloud.action.GetSharedUsersResp
	(*GetUserSharedReq)(nil),        // 6: meowcloud.action.GetUserSharedReq
	(*GetUserSharedResp)(nil),       // 7: meowcloud.action.GetUserSharedResp
	(*GetSharedReq)(nil),            // 8: meowcloud.action.GetSharedReq
	(*GetSharedResp)(nil),           // 9: meowcloud.action.GetSharedResp
	(TargetType)(0),                 // 10: meowcloud.action.TargetType
	(*basic.UserMeta)(nil),          // 11: basic.UserMeta
	(*basic.PaginationOptions)(nil), // 12: basic.PaginationOptions
	(*Action_Share)(nil),            // 13: meowcloud.action.Action.Share
}
var file_meowcloud_action_share_proto_depIdxs = []int32{
	10, // 0: meowcloud.action.DoShareReq.targetType:type_name -> meowcloud.action.TargetType
	11, // 1: meowcloud.action.DoShareReq.user:type_name -> basic.UserMeta
	10, // 2: meowcloud.action.GetSharedCountReq.targetType:type_name -> meowcloud.action.TargetType
	10, // 3: meowcloud.action.GetSharedUsersReq.targetType:type_name -> meowcloud.action.TargetType
	12, // 4: meowcloud.action.GetSharedUsersReq.paginationOption:type_name -> basic.PaginationOptions
	13, // 5: meowcloud.action.GetSharedUsersResp.shares:type_name -> meowcloud.action.Action.Share
	10, // 6: meowcloud.action.GetUserSharedReq.targetType:type_name -> meowcloud.action.TargetType
	12, // 7: meowcloud.action.GetUserSharedReq.paginationOption:type_name -> basic.PaginationOptions
	11, // 8: meowcloud.action.GetUserSharedReq.user:type_name -> basic.UserMeta
	13, // 9: meowcloud.action.GetUserSharedResp.shares:type_name -> meowcloud.action.Action.Share
	10, // 10: meowcloud.action.GetSharedReq.targetType:type_name -> meowcloud.action.TargetType
	11, // 11: meowcloud.action.GetSharedReq.user:type_name -> basic.UserMeta
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_meowcloud_action_share_proto_init() }
func file_meowcloud_action_share_proto_init() {
	if File_meowcloud_action_share_proto != nil {
		return
	}
	file_meowcloud_action_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_meowcloud_action_share_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoShareReq); i {
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
		file_meowcloud_action_share_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoShareResp); i {
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
		file_meowcloud_action_share_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSharedCountReq); i {
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
		file_meowcloud_action_share_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSharedCountResp); i {
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
		file_meowcloud_action_share_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSharedUsersReq); i {
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
		file_meowcloud_action_share_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSharedUsersResp); i {
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
		file_meowcloud_action_share_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserSharedReq); i {
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
		file_meowcloud_action_share_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserSharedResp); i {
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
		file_meowcloud_action_share_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSharedReq); i {
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
		file_meowcloud_action_share_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSharedResp); i {
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
			RawDescriptor: file_meowcloud_action_share_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meowcloud_action_share_proto_goTypes,
		DependencyIndexes: file_meowcloud_action_share_proto_depIdxs,
		MessageInfos:      file_meowcloud_action_share_proto_msgTypes,
	}.Build()
	File_meowcloud_action_share_proto = out.File
	file_meowcloud_action_share_proto_rawDesc = nil
	file_meowcloud_action_share_proto_goTypes = nil
	file_meowcloud_action_share_proto_depIdxs = nil
}

var _ context.Context
