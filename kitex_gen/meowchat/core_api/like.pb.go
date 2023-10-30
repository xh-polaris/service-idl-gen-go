// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: meowchat/core_api/like.proto

package core_api

import (
	context "context"
	user "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/user"
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

type DoLikeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetId   string        `protobuf:"bytes,1,opt,name=targetId,proto3" json:"targetId,omitempty"`
	TargetType user.LikeType `protobuf:"varint,2,opt,name=targetType,proto3,enum=meowchat.user.LikeType" json:"targetType,omitempty"`
}

func (x *DoLikeReq) Reset() {
	*x = DoLikeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_like_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoLikeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoLikeReq) ProtoMessage() {}

func (x *DoLikeReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_like_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoLikeReq.ProtoReflect.Descriptor instead.
func (*DoLikeReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_like_proto_rawDescGZIP(), []int{0}
}

func (x *DoLikeReq) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *DoLikeReq) GetTargetType() user.LikeType {
	if x != nil {
		return x.TargetType
	}
	return user.LikeType(0)
}

type DoLikeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsFirst    bool  `protobuf:"varint,1,opt,name=isFirst,proto3" json:"isFirst,omitempty"`
	GetFishNum int64 `protobuf:"varint,2,opt,name=getFishNum,proto3" json:"getFishNum,omitempty"`
}

func (x *DoLikeResp) Reset() {
	*x = DoLikeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_like_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoLikeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoLikeResp) ProtoMessage() {}

func (x *DoLikeResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_like_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoLikeResp.ProtoReflect.Descriptor instead.
func (*DoLikeResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_like_proto_rawDescGZIP(), []int{1}
}

func (x *DoLikeResp) GetIsFirst() bool {
	if x != nil {
		return x.IsFirst
	}
	return false
}

func (x *DoLikeResp) GetGetFishNum() int64 {
	if x != nil {
		return x.GetFishNum
	}
	return 0
}

type GetUserLikedReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetId   string        `protobuf:"bytes,1,opt,name=targetId,proto3" json:"targetId,omitempty"`
	TargetType user.LikeType `protobuf:"varint,2,opt,name=targetType,proto3,enum=meowchat.user.LikeType" json:"targetType,omitempty"`
}

func (x *GetUserLikedReq) Reset() {
	*x = GetUserLikedReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_like_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserLikedReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserLikedReq) ProtoMessage() {}

func (x *GetUserLikedReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_like_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserLikedReq.ProtoReflect.Descriptor instead.
func (*GetUserLikedReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_like_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserLikedReq) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *GetUserLikedReq) GetTargetType() user.LikeType {
	if x != nil {
		return x.TargetType
	}
	return user.LikeType(0)
}

type GetUserLikedResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Liked bool `protobuf:"varint,1,opt,name=liked,proto3" json:"liked,omitempty"`
}

func (x *GetUserLikedResp) Reset() {
	*x = GetUserLikedResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_like_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserLikedResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserLikedResp) ProtoMessage() {}

func (x *GetUserLikedResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_like_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserLikedResp.ProtoReflect.Descriptor instead.
func (*GetUserLikedResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_like_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserLikedResp) GetLiked() bool {
	if x != nil {
		return x.Liked
	}
	return false
}

type GetLikedCountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetId   string        `protobuf:"bytes,1,opt,name=targetId,proto3" json:"targetId,omitempty"`
	TargetType user.LikeType `protobuf:"varint,2,opt,name=targetType,proto3,enum=meowchat.user.LikeType" json:"targetType,omitempty"`
}

func (x *GetLikedCountReq) Reset() {
	*x = GetLikedCountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_like_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLikedCountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLikedCountReq) ProtoMessage() {}

func (x *GetLikedCountReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_like_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLikedCountReq.ProtoReflect.Descriptor instead.
func (*GetLikedCountReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_like_proto_rawDescGZIP(), []int{4}
}

func (x *GetLikedCountReq) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *GetLikedCountReq) GetTargetType() user.LikeType {
	if x != nil {
		return x.TargetType
	}
	return user.LikeType(0)
}

type GetLikedCountResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetLikedCountResp) Reset() {
	*x = GetLikedCountResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_like_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLikedCountResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLikedCountResp) ProtoMessage() {}

func (x *GetLikedCountResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_like_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLikedCountResp.ProtoReflect.Descriptor instead.
func (*GetLikedCountResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_like_proto_rawDescGZIP(), []int{5}
}

func (x *GetLikedCountResp) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetUserLikesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string        `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	TargetType user.LikeType `protobuf:"varint,2,opt,name=targetType,proto3,enum=meowchat.user.LikeType" json:"targetType,omitempty"`
}

func (x *GetUserLikesReq) Reset() {
	*x = GetUserLikesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_like_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserLikesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserLikesReq) ProtoMessage() {}

func (x *GetUserLikesReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_like_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserLikesReq.ProtoReflect.Descriptor instead.
func (*GetUserLikesReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_like_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserLikesReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetUserLikesReq) GetTargetType() user.LikeType {
	if x != nil {
		return x.TargetType
	}
	return user.LikeType(0)
}

type GetUserLikesResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Likes []*user.Like `protobuf:"bytes,1,rep,name=likes,proto3" json:"likes,omitempty"`
}

func (x *GetUserLikesResp) Reset() {
	*x = GetUserLikesResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_like_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserLikesResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserLikesResp) ProtoMessage() {}

func (x *GetUserLikesResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_like_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserLikesResp.ProtoReflect.Descriptor instead.
func (*GetUserLikesResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_like_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserLikesResp) GetLikes() []*user.Like {
	if x != nil {
		return x.Likes
	}
	return nil
}

type GetLikedUsersReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetId   string        `protobuf:"bytes,1,opt,name=targetId,proto3" json:"targetId,omitempty"`
	TargetType user.LikeType `protobuf:"varint,2,opt,name=targetType,proto3,enum=meowchat.user.LikeType" json:"targetType,omitempty"`
}

func (x *GetLikedUsersReq) Reset() {
	*x = GetLikedUsersReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_like_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLikedUsersReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLikedUsersReq) ProtoMessage() {}

func (x *GetLikedUsersReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_like_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLikedUsersReq.ProtoReflect.Descriptor instead.
func (*GetLikedUsersReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_like_proto_rawDescGZIP(), []int{8}
}

func (x *GetLikedUsersReq) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *GetLikedUsersReq) GetTargetType() user.LikeType {
	if x != nil {
		return x.TargetType
	}
	return user.LikeType(0)
}

type GetLikedUsersResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*user.UserPreview `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *GetLikedUsersResp) Reset() {
	*x = GetLikedUsersResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_like_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLikedUsersResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLikedUsersResp) ProtoMessage() {}

func (x *GetLikedUsersResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_like_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLikedUsersResp.ProtoReflect.Descriptor instead.
func (*GetLikedUsersResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_like_proto_rawDescGZIP(), []int{9}
}

func (x *GetLikedUsersResp) GetUsers() []*user.UserPreview {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_meowchat_core_api_like_proto protoreflect.FileDescriptor

var file_meowchat_core_api_like_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x1a, 0x1a, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a,
	0x09, 0x44, 0x6f, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x6f,
	0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x46, 0x0a, 0x0a, 0x44, 0x6f, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a,
	0x07, 0x69, 0x73, 0x46, 0x69, 0x72, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x69, 0x73, 0x46, 0x69, 0x72, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x65, 0x74, 0x46, 0x69,
	0x73, 0x68, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x67, 0x65, 0x74,
	0x46, 0x69, 0x73, 0x68, 0x4e, 0x75, 0x6d, 0x22, 0x66, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x6f,
	0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x28, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x64, 0x22, 0x67, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x6b, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a,
	0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x0a, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e,
	0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69,
	0x6b, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x29, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x62, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6d,
	0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x6b,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x3d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x6b, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x29, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73,
	0x22, 0x67, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64,
	0x12, 0x37, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x45, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x6b, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x30,
	0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x42, 0x7a, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x68, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69,
	0x73, 0x2e, 0x69, 0x64, 0x6c, 0x67, 0x65, 0x6e, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x61, 0x70, 0x69, 0x42, 0x09, 0x4c, 0x69, 0x6b, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x78, 0x68, 0x2d, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x69, 0x64, 0x6c, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f,
	0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x65, 0x6f, 0x77, 0x63,
	0x68, 0x61, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meowchat_core_api_like_proto_rawDescOnce sync.Once
	file_meowchat_core_api_like_proto_rawDescData = file_meowchat_core_api_like_proto_rawDesc
)

func file_meowchat_core_api_like_proto_rawDescGZIP() []byte {
	file_meowchat_core_api_like_proto_rawDescOnce.Do(func() {
		file_meowchat_core_api_like_proto_rawDescData = protoimpl.X.CompressGZIP(file_meowchat_core_api_like_proto_rawDescData)
	})
	return file_meowchat_core_api_like_proto_rawDescData
}

var file_meowchat_core_api_like_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_meowchat_core_api_like_proto_goTypes = []interface{}{
	(*DoLikeReq)(nil),         // 0: meowchat.core_api.DoLikeReq
	(*DoLikeResp)(nil),        // 1: meowchat.core_api.DoLikeResp
	(*GetUserLikedReq)(nil),   // 2: meowchat.core_api.GetUserLikedReq
	(*GetUserLikedResp)(nil),  // 3: meowchat.core_api.GetUserLikedResp
	(*GetLikedCountReq)(nil),  // 4: meowchat.core_api.GetLikedCountReq
	(*GetLikedCountResp)(nil), // 5: meowchat.core_api.GetLikedCountResp
	(*GetUserLikesReq)(nil),   // 6: meowchat.core_api.GetUserLikesReq
	(*GetUserLikesResp)(nil),  // 7: meowchat.core_api.GetUserLikesResp
	(*GetLikedUsersReq)(nil),  // 8: meowchat.core_api.GetLikedUsersReq
	(*GetLikedUsersResp)(nil), // 9: meowchat.core_api.GetLikedUsersResp
	(user.LikeType)(0),        // 10: meowchat.user.LikeType
	(*user.Like)(nil),         // 11: meowchat.user.Like
	(*user.UserPreview)(nil),  // 12: meowchat.user.UserPreview
}
var file_meowchat_core_api_like_proto_depIdxs = []int32{
	10, // 0: meowchat.core_api.DoLikeReq.targetType:type_name -> meowchat.user.LikeType
	10, // 1: meowchat.core_api.GetUserLikedReq.targetType:type_name -> meowchat.user.LikeType
	10, // 2: meowchat.core_api.GetLikedCountReq.targetType:type_name -> meowchat.user.LikeType
	10, // 3: meowchat.core_api.GetUserLikesReq.targetType:type_name -> meowchat.user.LikeType
	11, // 4: meowchat.core_api.GetUserLikesResp.likes:type_name -> meowchat.user.Like
	10, // 5: meowchat.core_api.GetLikedUsersReq.targetType:type_name -> meowchat.user.LikeType
	12, // 6: meowchat.core_api.GetLikedUsersResp.users:type_name -> meowchat.user.UserPreview
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_meowchat_core_api_like_proto_init() }
func file_meowchat_core_api_like_proto_init() {
	if File_meowchat_core_api_like_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_meowchat_core_api_like_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoLikeReq); i {
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
		file_meowchat_core_api_like_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoLikeResp); i {
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
		file_meowchat_core_api_like_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserLikedReq); i {
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
		file_meowchat_core_api_like_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserLikedResp); i {
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
		file_meowchat_core_api_like_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLikedCountReq); i {
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
		file_meowchat_core_api_like_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLikedCountResp); i {
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
		file_meowchat_core_api_like_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserLikesReq); i {
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
		file_meowchat_core_api_like_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserLikesResp); i {
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
		file_meowchat_core_api_like_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLikedUsersReq); i {
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
		file_meowchat_core_api_like_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLikedUsersResp); i {
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
			RawDescriptor: file_meowchat_core_api_like_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meowchat_core_api_like_proto_goTypes,
		DependencyIndexes: file_meowchat_core_api_like_proto_depIdxs,
		MessageInfos:      file_meowchat_core_api_like_proto_msgTypes,
	}.Build()
	File_meowchat_core_api_like_proto = out.File
	file_meowchat_core_api_like_proto_rawDesc = nil
	file_meowchat_core_api_like_proto_goTypes = nil
	file_meowchat_core_api_like_proto_depIdxs = nil
}

var _ context.Context
