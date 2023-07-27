// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: meowchat/user/user.proto

package user

import (
	context "context"
	_ "github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
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

type GetUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetUserReq) Reset() {
	*x = GetUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserReq) ProtoMessage() {}

func (x *GetUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserReq.ProtoReflect.Descriptor instead.
func (*GetUserReq) Descriptor() ([]byte, []int) {
	return file_meowchat_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetUserResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *UserPreview `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserResp) Reset() {
	*x = GetUserResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResp) ProtoMessage() {}

func (x *GetUserResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResp.ProtoReflect.Descriptor instead.
func (*GetUserResp) Descriptor() ([]byte, []int) {
	return file_meowchat_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserResp) GetUser() *UserPreview {
	if x != nil {
		return x.User
	}
	return nil
}

type GetUserDetailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetUserDetailReq) Reset() {
	*x = GetUserDetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserDetailReq) ProtoMessage() {}

func (x *GetUserDetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_user_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserDetailReq.ProtoReflect.Descriptor instead.
func (*GetUserDetailReq) Descriptor() ([]byte, []int) {
	return file_meowchat_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserDetailReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetUserDetailResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *UserDetail `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserDetailResp) Reset() {
	*x = GetUserDetailResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_user_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserDetailResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserDetailResp) ProtoMessage() {}

func (x *GetUserDetailResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_user_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserDetailResp.ProtoReflect.Descriptor instead.
func (*GetUserDetailResp) Descriptor() ([]byte, []int) {
	return file_meowchat_user_user_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserDetailResp) GetUser() *UserDetail {
	if x != nil {
		return x.User
	}
	return nil
}

type UpdateUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 如果不更新头像或昵称或其他东西，对应字段留空
	User *UserDetail `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UpdateUserReq) Reset() {
	*x = UpdateUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_user_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserReq) ProtoMessage() {}

func (x *UpdateUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_user_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserReq.ProtoReflect.Descriptor instead.
func (*UpdateUserReq) Descriptor() ([]byte, []int) {
	return file_meowchat_user_user_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateUserReq) GetUser() *UserDetail {
	if x != nil {
		return x.User
	}
	return nil
}

type UpdateUserResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateUserResp) Reset() {
	*x = UpdateUserResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_user_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserResp) ProtoMessage() {}

func (x *UpdateUserResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_user_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserResp.ProtoReflect.Descriptor instead.
func (*UpdateUserResp) Descriptor() ([]byte, []int) {
	return file_meowchat_user_user_proto_rawDescGZIP(), []int{5}
}

type SearchUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nickname  string  `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Offset    *int64  `protobuf:"varint,2,opt,name=offset,proto3,oneof" json:"offset,omitempty"`
	Limit     *int64  `protobuf:"varint,3,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	Backward  *bool   `protobuf:"varint,4,opt,name=backward,proto3,oneof" json:"backward,omitempty"`
	LastToken *string `protobuf:"bytes,5,opt,name=lastToken,proto3,oneof" json:"lastToken,omitempty"`
}

func (x *SearchUserReq) Reset() {
	*x = SearchUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_user_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUserReq) ProtoMessage() {}

func (x *SearchUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_user_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUserReq.ProtoReflect.Descriptor instead.
func (*SearchUserReq) Descriptor() ([]byte, []int) {
	return file_meowchat_user_user_proto_rawDescGZIP(), []int{6}
}

func (x *SearchUserReq) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *SearchUserReq) GetOffset() int64 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

func (x *SearchUserReq) GetLimit() int64 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *SearchUserReq) GetBackward() bool {
	if x != nil && x.Backward != nil {
		return *x.Backward
	}
	return false
}

func (x *SearchUserReq) GetLastToken() string {
	if x != nil && x.LastToken != nil {
		return *x.LastToken
	}
	return ""
}

type SearchUserResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*UserPreview `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Total int64          `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Token string         `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *SearchUserResp) Reset() {
	*x = SearchUserResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_user_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchUserResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchUserResp) ProtoMessage() {}

func (x *SearchUserResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_user_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchUserResp.ProtoReflect.Descriptor instead.
func (*SearchUserResp) Descriptor() ([]byte, []int) {
	return file_meowchat_user_user_proto_rawDescGZIP(), []int{7}
}

func (x *SearchUserResp) GetUsers() []*UserPreview {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *SearchUserResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SearchUserResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_meowchat_user_user_proto protoreflect.FileDescriptor

var file_meowchat_user_user_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x65, 0x6f, 0x77,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x10, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x6d, 0x65, 0x6f,
	0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3d, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x65, 0x6f,
	0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x2a, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2d, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x65,
	0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x3e, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x2d, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x65,
	0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x10, 0x0a, 0x0e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0xd7,
	0x01, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x62, 0x61, 0x63, 0x6b, 0x77, 0x61, 0x72, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x08, 0x62, 0x61, 0x63, 0x6b, 0x77, 0x61,
	0x72, 0x64, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x09, 0x6c, 0x61, 0x73, 0x74,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x77, 0x61, 0x72, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6c,
	0x61, 0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x6e, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x30, 0x0a, 0x05, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x65, 0x6f, 0x77,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xb9, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1a,
	0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x52, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1f, 0x2e, 0x6d, 0x65,
	0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x6d,
	0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x49,
	0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x6d,
	0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x6d, 0x65, 0x6f,
	0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x49, 0x0a, 0x0a, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x78, 0x68, 0x2d, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x69, 0x64, 0x6c, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f,
	0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x65, 0x6f, 0x77, 0x63,
	0x68, 0x61, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meowchat_user_user_proto_rawDescOnce sync.Once
	file_meowchat_user_user_proto_rawDescData = file_meowchat_user_user_proto_rawDesc
)

func file_meowchat_user_user_proto_rawDescGZIP() []byte {
	file_meowchat_user_user_proto_rawDescOnce.Do(func() {
		file_meowchat_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_meowchat_user_user_proto_rawDescData)
	})
	return file_meowchat_user_user_proto_rawDescData
}

var file_meowchat_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_meowchat_user_user_proto_goTypes = []interface{}{
	(*GetUserReq)(nil),        // 0: meowchat.user.GetUserReq
	(*GetUserResp)(nil),       // 1: meowchat.user.GetUserResp
	(*GetUserDetailReq)(nil),  // 2: meowchat.user.GetUserDetailReq
	(*GetUserDetailResp)(nil), // 3: meowchat.user.GetUserDetailResp
	(*UpdateUserReq)(nil),     // 4: meowchat.user.UpdateUserReq
	(*UpdateUserResp)(nil),    // 5: meowchat.user.UpdateUserResp
	(*SearchUserReq)(nil),     // 6: meowchat.user.SearchUserReq
	(*SearchUserResp)(nil),    // 7: meowchat.user.SearchUserResp
	(*UserPreview)(nil),       // 8: meowchat.user.UserPreview
	(*UserDetail)(nil),        // 9: meowchat.user.UserDetail
}
var file_meowchat_user_user_proto_depIdxs = []int32{
	8, // 0: meowchat.user.GetUserResp.user:type_name -> meowchat.user.UserPreview
	9, // 1: meowchat.user.GetUserDetailResp.user:type_name -> meowchat.user.UserDetail
	9, // 2: meowchat.user.UpdateUserReq.user:type_name -> meowchat.user.UserDetail
	8, // 3: meowchat.user.SearchUserResp.users:type_name -> meowchat.user.UserPreview
	0, // 4: meowchat.user.UserService.GetUser:input_type -> meowchat.user.GetUserReq
	2, // 5: meowchat.user.UserService.GetUserDetail:input_type -> meowchat.user.GetUserDetailReq
	4, // 6: meowchat.user.UserService.UpdateUser:input_type -> meowchat.user.UpdateUserReq
	6, // 7: meowchat.user.UserService.SearchUser:input_type -> meowchat.user.SearchUserReq
	1, // 8: meowchat.user.UserService.GetUser:output_type -> meowchat.user.GetUserResp
	3, // 9: meowchat.user.UserService.GetUserDetail:output_type -> meowchat.user.GetUserDetailResp
	5, // 10: meowchat.user.UserService.UpdateUser:output_type -> meowchat.user.UpdateUserResp
	7, // 11: meowchat.user.UserService.SearchUser:output_type -> meowchat.user.SearchUserResp
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_meowchat_user_user_proto_init() }
func file_meowchat_user_user_proto_init() {
	if File_meowchat_user_user_proto != nil {
		return
	}
	file_meowchat_user_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_meowchat_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserReq); i {
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
		file_meowchat_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserResp); i {
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
		file_meowchat_user_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserDetailReq); i {
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
		file_meowchat_user_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserDetailResp); i {
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
		file_meowchat_user_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserReq); i {
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
		file_meowchat_user_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserResp); i {
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
		file_meowchat_user_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchUserReq); i {
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
		file_meowchat_user_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchUserResp); i {
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
	file_meowchat_user_user_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_meowchat_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_meowchat_user_user_proto_goTypes,
		DependencyIndexes: file_meowchat_user_user_proto_depIdxs,
		MessageInfos:      file_meowchat_user_user_proto_msgTypes,
	}.Build()
	File_meowchat_user_user_proto = out.File
	file_meowchat_user_user_proto_rawDesc = nil
	file_meowchat_user_user_proto_goTypes = nil
	file_meowchat_user_user_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.6.2. DO NOT EDIT.

type UserService interface {
	GetUser(ctx context.Context, req *GetUserReq) (res *GetUserResp, err error)
	GetUserDetail(ctx context.Context, req *GetUserDetailReq) (res *GetUserDetailResp, err error)
	UpdateUser(ctx context.Context, req *UpdateUserReq) (res *UpdateUserResp, err error)
	SearchUser(ctx context.Context, req *SearchUserReq) (res *SearchUserResp, err error)
}
