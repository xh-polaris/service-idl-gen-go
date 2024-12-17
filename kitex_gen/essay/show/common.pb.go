// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.19.4
// source: essay/show/common.proto

package show

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

// 注册请求
type SignUpReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthId     string `protobuf:"bytes,1,opt,name=authId,proto3" json:"authId,omitempty"`
	AuthType   string `protobuf:"bytes,2,opt,name=authType,proto3" json:"authType,omitempty"`
	VerifyCode string `protobuf:"bytes,3,opt,name=verifyCode,proto3" json:"verifyCode,omitempty"`
	Password   string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"` // 需要MD5加密
	Name       string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SignUpReq) Reset() {
	*x = SignUpReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_essay_show_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpReq) ProtoMessage() {}

func (x *SignUpReq) ProtoReflect() protoreflect.Message {
	mi := &file_essay_show_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpReq.ProtoReflect.Descriptor instead.
func (*SignUpReq) Descriptor() ([]byte, []int) {
	return file_essay_show_common_proto_rawDescGZIP(), []int{0}
}

func (x *SignUpReq) GetAuthId() string {
	if x != nil {
		return x.AuthId
	}
	return ""
}

func (x *SignUpReq) GetAuthType() string {
	if x != nil {
		return x.AuthType
	}
	return ""
}

func (x *SignUpReq) GetVerifyCode() string {
	if x != nil {
		return x.VerifyCode
	}
	return ""
}

func (x *SignUpReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SignUpReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 注册响应
type SignUpResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AccessToken  string `protobuf:"bytes,2,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	AccessExpire int64  `protobuf:"varint,3,opt,name=accessExpire,proto3" json:"accessExpire,omitempty"`
}

func (x *SignUpResp) Reset() {
	*x = SignUpResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_essay_show_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpResp) ProtoMessage() {}

func (x *SignUpResp) ProtoReflect() protoreflect.Message {
	mi := &file_essay_show_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpResp.ProtoReflect.Descriptor instead.
func (*SignUpResp) Descriptor() ([]byte, []int) {
	return file_essay_show_common_proto_rawDescGZIP(), []int{1}
}

func (x *SignUpResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SignUpResp) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *SignUpResp) GetAccessExpire() int64 {
	if x != nil {
		return x.AccessExpire
	}
	return 0
}

// 登录请求
type SignInReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthId     string  `protobuf:"bytes,1,opt,name=authId,proto3" json:"authId,omitempty"`
	AuthType   string  `protobuf:"bytes,2,opt,name=authType,proto3" json:"authType,omitempty"`
	VerifyCode *string `protobuf:"bytes,3,opt,name=verifyCode,proto3,oneof" json:"verifyCode,omitempty"`
	Password   *string `protobuf:"bytes,4,opt,name=password,proto3,oneof" json:"password,omitempty"`
}

func (x *SignInReq) Reset() {
	*x = SignInReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_essay_show_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInReq) ProtoMessage() {}

func (x *SignInReq) ProtoReflect() protoreflect.Message {
	mi := &file_essay_show_common_proto_msgTypes[2]
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
	return file_essay_show_common_proto_rawDescGZIP(), []int{2}
}

func (x *SignInReq) GetAuthId() string {
	if x != nil {
		return x.AuthId
	}
	return ""
}

func (x *SignInReq) GetAuthType() string {
	if x != nil {
		return x.AuthType
	}
	return ""
}

func (x *SignInReq) GetVerifyCode() string {
	if x != nil && x.VerifyCode != nil {
		return *x.VerifyCode
	}
	return ""
}

func (x *SignInReq) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

// 登录响应
type SignInResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AccessToken  string `protobuf:"bytes,2,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	AccessExpire int64  `protobuf:"varint,3,opt,name=accessExpire,proto3" json:"accessExpire,omitempty"`
}

func (x *SignInResp) Reset() {
	*x = SignInResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_essay_show_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInResp) ProtoMessage() {}

func (x *SignInResp) ProtoReflect() protoreflect.Message {
	mi := &file_essay_show_common_proto_msgTypes[3]
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
	return file_essay_show_common_proto_rawDescGZIP(), []int{3}
}

func (x *SignInResp) GetId() string {
	if x != nil {
		return x.Id
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

// 获取用户信息请求
type GetUserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUserInfoReq) Reset() {
	*x = GetUserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_essay_show_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoReq) ProtoMessage() {}

func (x *GetUserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_essay_show_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoReq.ProtoReflect.Descriptor instead.
func (*GetUserInfoReq) Descriptor() ([]byte, []int) {
	return file_essay_show_common_proto_rawDescGZIP(), []int{4}
}

// 获取用户信息响应
type GetUserInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int64                    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg     string                   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Payload *GetUserInfoResp_Payload `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *GetUserInfoResp) Reset() {
	*x = GetUserInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_essay_show_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoResp) ProtoMessage() {}

func (x *GetUserInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_essay_show_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoResp.ProtoReflect.Descriptor instead.
func (*GetUserInfoResp) Descriptor() ([]byte, []int) {
	return file_essay_show_common_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserInfoResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetUserInfoResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetUserInfoResp) GetPayload() *GetUserInfoResp_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

// 更新用户信息
type UpdateUserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *UpdateUserInfoReq) Reset() {
	*x = UpdateUserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_essay_show_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserInfoReq) ProtoMessage() {}

func (x *UpdateUserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_essay_show_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserInfoReq.ProtoReflect.Descriptor instead.
func (*UpdateUserInfoReq) Descriptor() ([]byte, []int) {
	return file_essay_show_common_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateUserInfoReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

// 批改作文的请求
type EssayEvaluateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Text  string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Grade int64  `protobuf:"varint,3,opt,name=grade,proto3" json:"grade,omitempty"`
	Ocr   string `protobuf:"bytes,4,opt,name=ocr,proto3" json:"ocr,omitempty"`
}

func (x *EssayEvaluateReq) Reset() {
	*x = EssayEvaluateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_essay_show_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EssayEvaluateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EssayEvaluateReq) ProtoMessage() {}

func (x *EssayEvaluateReq) ProtoReflect() protoreflect.Message {
	mi := &file_essay_show_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EssayEvaluateReq.ProtoReflect.Descriptor instead.
func (*EssayEvaluateReq) Descriptor() ([]byte, []int) {
	return file_essay_show_common_proto_rawDescGZIP(), []int{7}
}

func (x *EssayEvaluateReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *EssayEvaluateReq) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *EssayEvaluateReq) GetGrade() int64 {
	if x != nil {
		return x.Grade
	}
	return 0
}

func (x *EssayEvaluateReq) GetOcr() string {
	if x != nil {
		return x.Ocr
	}
	return ""
}

// 批改作文的响应
type EssayEvaluateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg      string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Response string `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *EssayEvaluateResp) Reset() {
	*x = EssayEvaluateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_essay_show_common_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EssayEvaluateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EssayEvaluateResp) ProtoMessage() {}

func (x *EssayEvaluateResp) ProtoReflect() protoreflect.Message {
	mi := &file_essay_show_common_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EssayEvaluateResp.ProtoReflect.Descriptor instead.
func (*EssayEvaluateResp) Descriptor() ([]byte, []int) {
	return file_essay_show_common_proto_rawDescGZIP(), []int{8}
}

func (x *EssayEvaluateResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *EssayEvaluateResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *EssayEvaluateResp) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

// 获取批改记录请求
type GetEssayEvaluateLogsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaginationOptions *basic.PaginationOptions `protobuf:"bytes,1,opt,name=paginationOptions,proto3" json:"paginationOptions,omitempty"`
}

func (x *GetEssayEvaluateLogsReq) Reset() {
	*x = GetEssayEvaluateLogsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_essay_show_common_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEssayEvaluateLogsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEssayEvaluateLogsReq) ProtoMessage() {}

func (x *GetEssayEvaluateLogsReq) ProtoReflect() protoreflect.Message {
	mi := &file_essay_show_common_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEssayEvaluateLogsReq.ProtoReflect.Descriptor instead.
func (*GetEssayEvaluateLogsReq) Descriptor() ([]byte, []int) {
	return file_essay_show_common_proto_rawDescGZIP(), []int{9}
}

func (x *GetEssayEvaluateLogsReq) GetPaginationOptions() *basic.PaginationOptions {
	if x != nil {
		return x.PaginationOptions
	}
	return nil
}

// 获取批改记录响应
type GetEssayEvaluateLogsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Logs  []*Log `protobuf:"bytes,2,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (x *GetEssayEvaluateLogsResp) Reset() {
	*x = GetEssayEvaluateLogsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_essay_show_common_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEssayEvaluateLogsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEssayEvaluateLogsResp) ProtoMessage() {}

func (x *GetEssayEvaluateLogsResp) ProtoReflect() protoreflect.Message {
	mi := &file_essay_show_common_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEssayEvaluateLogsResp.ProtoReflect.Descriptor instead.
func (*GetEssayEvaluateLogsResp) Descriptor() ([]byte, []int) {
	return file_essay_show_common_proto_rawDescGZIP(), []int{10}
}

func (x *GetEssayEvaluateLogsResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GetEssayEvaluateLogsResp) GetLogs() []*Log {
	if x != nil {
		return x.Logs
	}
	return nil
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Grade      int64  `protobuf:"varint,2,opt,name=grade,proto3" json:"grade,omitempty"`
	Ocr        string `protobuf:"bytes,3,opt,name=ocr,proto3" json:"ocr,omitempty"`
	Response   string `protobuf:"bytes,4,opt,name=response,proto3" json:"response,omitempty"`
	CreateTime int64  `protobuf:"varint,5,opt,name=createTime,proto3" json:"createTime,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_essay_show_common_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_essay_show_common_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_essay_show_common_proto_rawDescGZIP(), []int{11}
}

func (x *Log) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Log) GetGrade() int64 {
	if x != nil {
		return x.Grade
	}
	return 0
}

func (x *Log) GetOcr() string {
	if x != nil {
		return x.Ocr
	}
	return ""
}

func (x *Log) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *Log) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_essay_show_common_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_essay_show_common_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_essay_show_common_proto_rawDescGZIP(), []int{12}
}

func (x *Response) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GetUserInfoResp_Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Count int64  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetUserInfoResp_Payload) Reset() {
	*x = GetUserInfoResp_Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_essay_show_common_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoResp_Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoResp_Payload) ProtoMessage() {}

func (x *GetUserInfoResp_Payload) ProtoReflect() protoreflect.Message {
	mi := &file_essay_show_common_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoResp_Payload.ProtoReflect.Descriptor instead.
func (*GetUserInfoResp_Payload) Descriptor() ([]byte, []int) {
	return file_essay_show_common_proto_rawDescGZIP(), []int{5, 0}
}

func (x *GetUserInfoResp_Payload) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetUserInfoResp_Payload) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_essay_show_common_proto protoreflect.FileDescriptor

var file_essay_show_common_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x73, 0x73, 0x61, 0x79, 0x2f, 0x73, 0x68, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x65, 0x73, 0x73, 0x61, 0x79,
	0x2e, 0x73, 0x68, 0x6f, 0x77, 0x1a, 0x16, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01,
	0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x75, 0x74, 0x68, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x62, 0x0a, 0x0a, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x22, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x22, 0xa1, 0x01, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65,
	0x71, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74,
	0x68, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74,
	0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0a, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x62, 0x0a, 0x0a, 0x53, 0x69, 0x67, 0x6e, 0x49,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x22, 0xab, 0x01,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x3d, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x73, 0x73, 0x61, 0x79,
	0x2e, 0x73, 0x68, 0x6f, 0x77, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x33, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2f, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x64, 0x0a, 0x10,
	0x45, 0x73, 0x73, 0x61, 0x79, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6f, 0x63, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f,
	0x63, 0x72, 0x22, 0x55, 0x0a, 0x11, 0x45, 0x73, 0x73, 0x61, 0x79, 0x45, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x61, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x45, 0x73, 0x73, 0x61, 0x79, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67,
	0x73, 0x52, 0x65, 0x71, 0x12, 0x46, 0x0a, 0x11, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x11, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x55, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x45, 0x73, 0x73, 0x61, 0x79, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65,
	0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x23,
	0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x65,
	0x73, 0x73, 0x61, 0x79, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x6c,
	0x6f, 0x67, 0x73, 0x22, 0x79, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6f, 0x63, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f,
	0x63, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x30,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x42, 0x6f, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x68, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69,
	0x73, 0x2e, 0x69, 0x64, 0x6c, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x73, 0x73, 0x61, 0x79, 0x2e, 0x73,
	0x68, 0x6f, 0x77, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78,
	0x68, 0x2d, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2d, 0x69, 0x64, 0x6c, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x6b, 0x69, 0x74,
	0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x73, 0x73, 0x61, 0x79, 0x2f, 0x73, 0x68, 0x6f,
	0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_essay_show_common_proto_rawDescOnce sync.Once
	file_essay_show_common_proto_rawDescData = file_essay_show_common_proto_rawDesc
)

func file_essay_show_common_proto_rawDescGZIP() []byte {
	file_essay_show_common_proto_rawDescOnce.Do(func() {
		file_essay_show_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_essay_show_common_proto_rawDescData)
	})
	return file_essay_show_common_proto_rawDescData
}

var file_essay_show_common_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_essay_show_common_proto_goTypes = []interface{}{
	(*SignUpReq)(nil),                // 0: essay.show.SignUpReq
	(*SignUpResp)(nil),               // 1: essay.show.SignUpResp
	(*SignInReq)(nil),                // 2: essay.show.SignInReq
	(*SignInResp)(nil),               // 3: essay.show.SignInResp
	(*GetUserInfoReq)(nil),           // 4: essay.show.GetUserInfoReq
	(*GetUserInfoResp)(nil),          // 5: essay.show.GetUserInfoResp
	(*UpdateUserInfoReq)(nil),        // 6: essay.show.UpdateUserInfoReq
	(*EssayEvaluateReq)(nil),         // 7: essay.show.EssayEvaluateReq
	(*EssayEvaluateResp)(nil),        // 8: essay.show.EssayEvaluateResp
	(*GetEssayEvaluateLogsReq)(nil),  // 9: essay.show.GetEssayEvaluateLogsReq
	(*GetEssayEvaluateLogsResp)(nil), // 10: essay.show.GetEssayEvaluateLogsResp
	(*Log)(nil),                      // 11: essay.show.Log
	(*Response)(nil),                 // 12: essay.show.Response
	(*GetUserInfoResp_Payload)(nil),  // 13: essay.show.GetUserInfoResp.Payload
	(*basic.PaginationOptions)(nil),  // 14: basic.PaginationOptions
}
var file_essay_show_common_proto_depIdxs = []int32{
	13, // 0: essay.show.GetUserInfoResp.payload:type_name -> essay.show.GetUserInfoResp.Payload
	14, // 1: essay.show.GetEssayEvaluateLogsReq.paginationOptions:type_name -> basic.PaginationOptions
	11, // 2: essay.show.GetEssayEvaluateLogsResp.logs:type_name -> essay.show.Log
	3,  // [3:3] is the sub-list for method output_type
	3,  // [3:3] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_essay_show_common_proto_init() }
func file_essay_show_common_proto_init() {
	if File_essay_show_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_essay_show_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpReq); i {
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
		file_essay_show_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpResp); i {
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
		file_essay_show_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_essay_show_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_essay_show_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoReq); i {
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
		file_essay_show_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoResp); i {
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
		file_essay_show_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserInfoReq); i {
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
		file_essay_show_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EssayEvaluateReq); i {
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
		file_essay_show_common_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EssayEvaluateResp); i {
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
		file_essay_show_common_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEssayEvaluateLogsReq); i {
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
		file_essay_show_common_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEssayEvaluateLogsResp); i {
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
		file_essay_show_common_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
		file_essay_show_common_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_essay_show_common_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoResp_Payload); i {
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
	file_essay_show_common_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_essay_show_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_essay_show_common_proto_goTypes,
		DependencyIndexes: file_essay_show_common_proto_depIdxs,
		MessageInfos:      file_essay_show_common_proto_msgTypes,
	}.Build()
	File_essay_show_common_proto = out.File
	file_essay_show_common_proto_rawDesc = nil
	file_essay_show_common_proto_goTypes = nil
	file_essay_show_common_proto_depIdxs = nil
}

var _ context.Context
