// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: meowchat/core_api/auth.proto

package core_api

import (
	context "context"
	basic "github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
	_ "github.com/xh-polaris/service-idl-gen-go/kitex_gen/http"
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

type AuthType int32

const (
	AuthType_Wechat AuthType = 0
	AuthType_Email  AuthType = 1
	AuthType_Phone  AuthType = 2
)

// Enum value maps for AuthType.
var (
	AuthType_name = map[int32]string{
		0: "Wechat",
		1: "Email",
		2: "Phone",
	}
	AuthType_value = map[string]int32{
		"Wechat": 0,
		"Email":  1,
		"Phone":  2,
	}
)

func (x AuthType) Enum() *AuthType {
	p := new(AuthType)
	*p = x
	return p
}

func (x AuthType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthType) Descriptor() protoreflect.EnumDescriptor {
	return file_meowchat_core_api_auth_proto_enumTypes[0].Descriptor()
}

func (AuthType) Type() protoreflect.EnumType {
	return &file_meowchat_core_api_auth_proto_enumTypes[0]
}

func (x AuthType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthType.Descriptor instead.
func (AuthType) EnumDescriptor() ([]byte, []int) {
	return file_meowchat_core_api_auth_proto_rawDescGZIP(), []int{0}
}

type SignInReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthType string   `protobuf:"bytes,1,opt,name=authType,proto3" json:"authType,omitempty"`
	AuthId   string   `protobuf:"bytes,2,opt,name=authId,proto3" json:"authId,omitempty"`
	Password *string  `protobuf:"bytes,3,opt,name=password,proto3,oneof" json:"password,omitempty"`
	Params   []string `protobuf:"bytes,4,rep,name=params,proto3" json:"params,omitempty"`
}

func (x *SignInReq) Reset() {
	*x = SignInReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInReq) ProtoMessage() {}

func (x *SignInReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInReq.ProtoReflect.Descriptor instead.
func (*SignInReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_auth_proto_rawDescGZIP(), []int{0}
}

func (x *SignInReq) GetAuthType() string {
	if x != nil {
		return x.AuthType
	}
	return ""
}

func (x *SignInReq) GetAuthId() string {
	if x != nil {
		return x.AuthId
	}
	return ""
}

func (x *SignInReq) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *SignInReq) GetParams() []string {
	if x != nil {
		return x.Params
	}
	return nil
}

type SignInResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	AccessToken  string `protobuf:"bytes,2,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	AccessExpire int64  `protobuf:"varint,3,opt,name=accessExpire,proto3" json:"accessExpire,omitempty"`
}

func (x *SignInResp) Reset() {
	*x = SignInResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInResp) ProtoMessage() {}

func (x *SignInResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInResp.ProtoReflect.Descriptor instead.
func (*SignInResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_auth_proto_rawDescGZIP(), []int{1}
}

func (x *SignInResp) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SignInResp) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *SignInResp) GetAccessExpire() int64 {
	if x != nil {
		return x.AccessExpire
	}
	return 0
}

type SetPasswordReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string          `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	User     *basic.UserMeta `protobuf:"bytes,255,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *SetPasswordReq) Reset() {
	*x = SetPasswordReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPasswordReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPasswordReq) ProtoMessage() {}

func (x *SetPasswordReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPasswordReq.ProtoReflect.Descriptor instead.
func (*SetPasswordReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_auth_proto_rawDescGZIP(), []int{2}
}

func (x *SetPasswordReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SetPasswordReq) GetUser() *basic.UserMeta {
	if x != nil {
		return x.User
	}
	return nil
}

type SetPasswordResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetPasswordResp) Reset() {
	*x = SetPasswordResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPasswordResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPasswordResp) ProtoMessage() {}

func (x *SetPasswordResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPasswordResp.ProtoReflect.Descriptor instead.
func (*SetPasswordResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_auth_proto_rawDescGZIP(), []int{3}
}

type SendVerifyCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthType string `protobuf:"bytes,1,opt,name=authType,proto3" json:"authType,omitempty"`
	AuthId   string `protobuf:"bytes,2,opt,name=authId,proto3" json:"authId,omitempty"`
}

func (x *SendVerifyCodeReq) Reset() {
	*x = SendVerifyCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendVerifyCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendVerifyCodeReq) ProtoMessage() {}

func (x *SendVerifyCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendVerifyCodeReq.ProtoReflect.Descriptor instead.
func (*SendVerifyCodeReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_auth_proto_rawDescGZIP(), []int{4}
}

func (x *SendVerifyCodeReq) GetAuthType() string {
	if x != nil {
		return x.AuthType
	}
	return ""
}

func (x *SendVerifyCodeReq) GetAuthId() string {
	if x != nil {
		return x.AuthId
	}
	return ""
}

type SendVerifyCodeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendVerifyCodeResp) Reset() {
	*x = SendVerifyCodeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendVerifyCodeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendVerifyCodeResp) ProtoMessage() {}

func (x *SendVerifyCodeResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendVerifyCodeResp.ProtoReflect.Descriptor instead.
func (*SendVerifyCodeResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_auth_proto_rawDescGZIP(), []int{5}
}

var File_meowchat_core_api_auth_proto protoreflect.FileDescriptor

var file_meowchat_core_api_auth_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x1a, 0x0f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x10, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52,
	0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x6a, 0x0a, 0x0a,
	0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x22, 0x65, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x37, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0xff,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x42, 0x11, 0xba, 0xbb, 0x18, 0x0d, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0x11, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x47, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x53,
	0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x2a, 0x2c, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a,
	0x06, 0x57, 0x65, 0x63, 0x68, 0x61, 0x74, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x10, 0x02, 0x42,
	0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x68,
	0x2d, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2d, 0x69, 0x64, 0x6c, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x6b, 0x69, 0x74, 0x65,
	0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meowchat_core_api_auth_proto_rawDescOnce sync.Once
	file_meowchat_core_api_auth_proto_rawDescData = file_meowchat_core_api_auth_proto_rawDesc
)

func file_meowchat_core_api_auth_proto_rawDescGZIP() []byte {
	file_meowchat_core_api_auth_proto_rawDescOnce.Do(func() {
		file_meowchat_core_api_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_meowchat_core_api_auth_proto_rawDescData)
	})
	return file_meowchat_core_api_auth_proto_rawDescData
}

var file_meowchat_core_api_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_meowchat_core_api_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_meowchat_core_api_auth_proto_goTypes = []interface{}{
	(AuthType)(0),              // 0: meowchat.core_api.AuthType
	(*SignInReq)(nil),          // 1: meowchat.core_api.SignInReq
	(*SignInResp)(nil),         // 2: meowchat.core_api.SignInResp
	(*SetPasswordReq)(nil),     // 3: meowchat.core_api.SetPasswordReq
	(*SetPasswordResp)(nil),    // 4: meowchat.core_api.SetPasswordResp
	(*SendVerifyCodeReq)(nil),  // 5: meowchat.core_api.SendVerifyCodeReq
	(*SendVerifyCodeResp)(nil), // 6: meowchat.core_api.SendVerifyCodeResp
	(*basic.UserMeta)(nil),     // 7: basic.UserMeta
}
var file_meowchat_core_api_auth_proto_depIdxs = []int32{
	7, // 0: meowchat.core_api.SetPasswordReq.user:type_name -> basic.UserMeta
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_meowchat_core_api_auth_proto_init() }
func file_meowchat_core_api_auth_proto_init() {
	if File_meowchat_core_api_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_meowchat_core_api_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInReq); i {
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
		file_meowchat_core_api_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInResp); i {
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
		file_meowchat_core_api_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPasswordReq); i {
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
		file_meowchat_core_api_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPasswordResp); i {
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
		file_meowchat_core_api_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendVerifyCodeReq); i {
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
		file_meowchat_core_api_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendVerifyCodeResp); i {
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
	file_meowchat_core_api_auth_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_meowchat_core_api_auth_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meowchat_core_api_auth_proto_goTypes,
		DependencyIndexes: file_meowchat_core_api_auth_proto_depIdxs,
		EnumInfos:         file_meowchat_core_api_auth_proto_enumTypes,
		MessageInfos:      file_meowchat_core_api_auth_proto_msgTypes,
	}.Build()
	File_meowchat_core_api_auth_proto = out.File
	file_meowchat_core_api_auth_proto_rawDesc = nil
	file_meowchat_core_api_auth_proto_goTypes = nil
	file_meowchat_core_api_auth_proto_depIdxs = nil
}

var _ context.Context
