// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
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

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Phone    string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"` // 需要MD5加密
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

func (x *SignUpReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SignUpReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SignUpReq) GetPassword() string {
	if x != nil {
		return x.Password
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

// 批改文本作文的请求
type EssayEvaluateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *EssayEvaluateReq) Reset() {
	*x = EssayEvaluateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_essay_show_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EssayEvaluateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EssayEvaluateReq) ProtoMessage() {}

func (x *EssayEvaluateReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use EssayEvaluateReq.ProtoReflect.Descriptor instead.
func (*EssayEvaluateReq) Descriptor() ([]byte, []int) {
	return file_essay_show_common_proto_rawDescGZIP(), []int{2}
}

func (x *EssayEvaluateReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *EssayEvaluateReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// 批改文本作文的响应
type EssayEvaluateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     int64             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg      string            `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Response *EvaluateResponse `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *EssayEvaluateResp) Reset() {
	*x = EssayEvaluateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_essay_show_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EssayEvaluateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EssayEvaluateResp) ProtoMessage() {}

func (x *EssayEvaluateResp) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use EssayEvaluateResp.ProtoReflect.Descriptor instead.
func (*EssayEvaluateResp) Descriptor() ([]byte, []int) {
	return file_essay_show_common_proto_rawDescGZIP(), []int{3}
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

func (x *EssayEvaluateResp) GetResponse() *EvaluateResponse {
	if x != nil {
		return x.Response
	}
	return nil
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
		mi := &file_essay_show_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEssayEvaluateLogsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEssayEvaluateLogsReq) ProtoMessage() {}

func (x *GetEssayEvaluateLogsReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetEssayEvaluateLogsReq.ProtoReflect.Descriptor instead.
func (*GetEssayEvaluateLogsReq) Descriptor() ([]byte, []int) {
	return file_essay_show_common_proto_rawDescGZIP(), []int{4}
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

	Total     int64               `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Responses []*EvaluateResponse `protobuf:"bytes,2,rep,name=responses,proto3" json:"responses,omitempty"`
}

func (x *GetEssayEvaluateLogsResp) Reset() {
	*x = GetEssayEvaluateLogsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_essay_show_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEssayEvaluateLogsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEssayEvaluateLogsResp) ProtoMessage() {}

func (x *GetEssayEvaluateLogsResp) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetEssayEvaluateLogsResp.ProtoReflect.Descriptor instead.
func (*GetEssayEvaluateLogsResp) Descriptor() ([]byte, []int) {
	return file_essay_show_common_proto_rawDescGZIP(), []int{5}
}

func (x *GetEssayEvaluateLogsResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GetEssayEvaluateLogsResp) GetResponses() []*EvaluateResponse {
	if x != nil {
		return x.Responses
	}
	return nil
}

// beta接口批改的响应
type EvaluateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//	message ModelVersion {
	//	  string name = 1;
	//	  string version = 2;
	//	}
	//
	//	message OverallEvaluation {
	//	  int64 topicRelevanceScore = 1;
	//	  string description = 2;
	//	}
	//
	//	message FluencyEvaluation {
	//	  int64 fluencyScore = 1;
	//	  string fluencyDescription = 2;
	//	}
	//
	//	message WordSentenceEvaluation {
	//	  int64 wordSentenceScore = 1;
	//	  string wordSentenceDescription = 2;
	//	  repeated google.protobuf.Any sentenceEvaluations = 3;
	//	}
	//
	//	message ExpressionEvaluation {
	//	  int64 expressionScore = 1;
	//	  string expressDescription = 2;
	//	  repeated google.protobuf.Any expressionEvaluations = 3;
	//	}
	//
	//	message SuggestionEvaluation {
	//	  string suggestionDescription = 1;
	//	}
	//
	//	message ParagraphEvaluations {
	//	  int64 paragraphIndex = 1;
	//	  string content = 2;
	//	}
	//
	//	message AIEvaluation {
	//	  ModelVersion modelVersion = 1;
	//	  OverallEvaluation overallEvaluation = 2;
	//	  FluencyEvaluation fluencyEvaluation = 3;
	//	  WordSentenceEvaluation wordSentenceEvaluation = 4;
	//	  ExpressionEvaluation expressionEvaluation = 5;
	//	  SuggestionEvaluation suggestionEvaluation = 6;
	//	  repeated ParagraphEvaluations paragraphEvaluations = 7;
	//	}
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *EvaluateResponse) Reset() {
	*x = EvaluateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_essay_show_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvaluateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluateResponse) ProtoMessage() {}

func (x *EvaluateResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use EvaluateResponse.ProtoReflect.Descriptor instead.
func (*EvaluateResponse) Descriptor() ([]byte, []int) {
	return file_essay_show_common_proto_rawDescGZIP(), []int{6}
}

func (x *EvaluateResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

var File_essay_show_common_proto protoreflect.FileDescriptor

var file_essay_show_common_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x73, 0x73, 0x61, 0x79, 0x2f, 0x73, 0x68, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x65, 0x73, 0x73, 0x61, 0x79,
	0x2e, 0x73, 0x68, 0x6f, 0x77, 0x1a, 0x16, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a,
	0x09, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x62, 0x0a, 0x0a, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x22, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x22, 0x42, 0x0a, 0x10, 0x45, 0x73, 0x73, 0x61, 0x79, 0x45, 0x76, 0x61,
	0x6c, 0x75, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x73, 0x0a, 0x11, 0x45, 0x73, 0x73, 0x61,
	0x79, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x38, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x73, 0x73, 0x61, 0x79, 0x2e, 0x73, 0x68,
	0x6f, 0x77, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x61, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x45, 0x73, 0x73, 0x61, 0x79, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74,
	0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x12, 0x46, 0x0a, 0x11, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x11, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x6c, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x45, 0x73, 0x73, 0x61, 0x79, 0x45, 0x76, 0x61, 0x6c,
	0x75, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x3a, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x73, 0x73, 0x61, 0x79, 0x2e, 0x73, 0x68,
	0x6f, 0x77, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x22, 0x28,
	0x0a, 0x10, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x6f, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e,
	0x78, 0x68, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2e, 0x69, 0x64, 0x6c, 0x67, 0x65, 0x6e,
	0x2e, 0x65, 0x73, 0x73, 0x61, 0x79, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x42, 0x0b, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x68, 0x2d, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69,
	0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x69, 0x64, 0x6c, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x65,
	0x73, 0x73, 0x61, 0x79, 0x2f, 0x73, 0x68, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
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

var file_essay_show_common_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_essay_show_common_proto_goTypes = []interface{}{
	(*SignUpReq)(nil),                // 0: essay.show.SignUpReq
	(*SignUpResp)(nil),               // 1: essay.show.SignUpResp
	(*EssayEvaluateReq)(nil),         // 2: essay.show.EssayEvaluateReq
	(*EssayEvaluateResp)(nil),        // 3: essay.show.EssayEvaluateResp
	(*GetEssayEvaluateLogsReq)(nil),  // 4: essay.show.GetEssayEvaluateLogsReq
	(*GetEssayEvaluateLogsResp)(nil), // 5: essay.show.GetEssayEvaluateLogsResp
	(*EvaluateResponse)(nil),         // 6: essay.show.EvaluateResponse
	(*basic.PaginationOptions)(nil),  // 7: basic.PaginationOptions
}
var file_essay_show_common_proto_depIdxs = []int32{
	6, // 0: essay.show.EssayEvaluateResp.response:type_name -> essay.show.EvaluateResponse
	7, // 1: essay.show.GetEssayEvaluateLogsReq.paginationOptions:type_name -> basic.PaginationOptions
	6, // 2: essay.show.GetEssayEvaluateLogsResp.responses:type_name -> essay.show.EvaluateResponse
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
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
		file_essay_show_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_essay_show_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_essay_show_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_essay_show_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvaluateResponse); i {
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
			RawDescriptor: file_essay_show_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
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
