// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.19.4
// source: meowchat/core_api/agg.proto

package core_api

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

type PrefetchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid     string  `protobuf:"bytes,1,opt,name=appid,proto3" json:"appid,omitempty"` // 微信小程序AppID
	Token     *string `protobuf:"bytes,2,opt,name=token,proto3,oneof" json:"token,omitempty"`
	Code      *string `protobuf:"bytes,3,opt,name=code,proto3,oneof" json:"code,omitempty"`      // 登录凭证 code
	Timestamp int64   `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // 时间戳
	Path      *string `protobuf:"bytes,5,opt,name=path,proto3,oneof" json:"path,omitempty"`      // 页面路径
	Query     *string `protobuf:"bytes,6,opt,name=query,proto3,oneof" json:"query,omitempty"`    // 页面参数
	Scene     *string `protobuf:"bytes,7,opt,name=scene,proto3,oneof" json:"scene,omitempty"`    // 场景值
}

func (x *PrefetchReq) Reset() {
	*x = PrefetchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_agg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrefetchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrefetchReq) ProtoMessage() {}

func (x *PrefetchReq) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_agg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrefetchReq.ProtoReflect.Descriptor instead.
func (*PrefetchReq) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_agg_proto_rawDescGZIP(), []int{0}
}

func (x *PrefetchReq) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

func (x *PrefetchReq) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

func (x *PrefetchReq) GetCode() string {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ""
}

func (x *PrefetchReq) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *PrefetchReq) GetPath() string {
	if x != nil && x.Path != nil {
		return *x.Path
	}
	return ""
}

func (x *PrefetchReq) GetQuery() string {
	if x != nil && x.Query != nil {
		return *x.Query
	}
	return ""
}

func (x *PrefetchReq) GetScene() string {
	if x != nil && x.Scene != nil {
		return *x.Scene
	}
	return ""
}

// 所有响应都是弱依赖
type PrefetchResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SignInResp              *SignInResp            `protobuf:"bytes,1,opt,name=signInResp,proto3,oneof" json:"signInResp,omitempty"`
	GetUserInfoResp         *GetUserInfoResp       `protobuf:"bytes,2,opt,name=getUserInfoResp,proto3,oneof" json:"getUserInfoResp,omitempty"`
	ListCommunityResp       *ListCommunityResp     `protobuf:"bytes,3,opt,name=listCommunityResp,proto3,oneof" json:"listCommunityResp,omitempty"`
	FirstPostPreviewsResp   *GetPostPreviewsResp   `protobuf:"bytes,4,opt,name=firstPostPreviewsResp,proto3,oneof" json:"firstPostPreviewsResp,omitempty"`
	FirstMomentPreviewsResp *GetMomentPreviewsResp `protobuf:"bytes,5,opt,name=firstMomentPreviewsResp,proto3,oneof" json:"firstMomentPreviewsResp,omitempty"`
	FirstCatPreviewsResp    *GetCatPreviewsResp    `protobuf:"bytes,6,opt,name=firstCatPreviewsResp,proto3,oneof" json:"firstCatPreviewsResp,omitempty"`
	GetNewsResp             *GetNewsResp           `protobuf:"bytes,7,opt,name=getNewsResp,proto3,oneof" json:"getNewsResp,omitempty"`
	Token                   *string                `protobuf:"bytes,8,opt,name=token,proto3,oneof" json:"token,omitempty"`    // 直接返回token
	Timestamp               int64                  `protobuf:"varint,9,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // 时间戳
}

func (x *PrefetchResp) Reset() {
	*x = PrefetchResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_core_api_agg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrefetchResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrefetchResp) ProtoMessage() {}

func (x *PrefetchResp) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_core_api_agg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrefetchResp.ProtoReflect.Descriptor instead.
func (*PrefetchResp) Descriptor() ([]byte, []int) {
	return file_meowchat_core_api_agg_proto_rawDescGZIP(), []int{1}
}

func (x *PrefetchResp) GetSignInResp() *SignInResp {
	if x != nil {
		return x.SignInResp
	}
	return nil
}

func (x *PrefetchResp) GetGetUserInfoResp() *GetUserInfoResp {
	if x != nil {
		return x.GetUserInfoResp
	}
	return nil
}

func (x *PrefetchResp) GetListCommunityResp() *ListCommunityResp {
	if x != nil {
		return x.ListCommunityResp
	}
	return nil
}

func (x *PrefetchResp) GetFirstPostPreviewsResp() *GetPostPreviewsResp {
	if x != nil {
		return x.FirstPostPreviewsResp
	}
	return nil
}

func (x *PrefetchResp) GetFirstMomentPreviewsResp() *GetMomentPreviewsResp {
	if x != nil {
		return x.FirstMomentPreviewsResp
	}
	return nil
}

func (x *PrefetchResp) GetFirstCatPreviewsResp() *GetCatPreviewsResp {
	if x != nil {
		return x.FirstCatPreviewsResp
	}
	return nil
}

func (x *PrefetchResp) GetGetNewsResp() *GetNewsResp {
	if x != nil {
		return x.GetNewsResp
	}
	return nil
}

func (x *PrefetchResp) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

func (x *PrefetchResp) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_meowchat_core_api_agg_proto protoreflect.FileDescriptor

var file_meowchat_core_api_agg_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x67, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6d,
	0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69,
	0x1a, 0x1c, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6d, 0x65,
	0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6d, 0x65,
	0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6d, 0x65, 0x6f, 0x77,
	0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6d, 0x65, 0x6f, 0x77,
	0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4,
	0x01, 0x0a, 0x0b, 0x50, 0x72, 0x65, 0x66, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12, 0x14,
	0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61,
	0x70, 0x70, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x12,
	0x17, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x88, 0x01, 0x01, 0x12,
	0x19, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x73, 0x63,
	0x65, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x05, 0x73, 0x63, 0x65,
	0x6e, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x73, 0x63, 0x65, 0x6e, 0x65, 0x22, 0xcc, 0x06, 0x0a, 0x0c, 0x50, 0x72, 0x65, 0x66, 0x65, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x42, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x49, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x65, 0x6f,
	0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x69, 0x67,
	0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x88, 0x01, 0x01, 0x12, 0x51, 0x0a, 0x0f, 0x67, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x48, 0x01, 0x52, 0x0f, 0x67, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x88, 0x01, 0x01, 0x12, 0x57, 0x0a,
	0x11, 0x6c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x48, 0x02,
	0x52, 0x11, 0x6c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x88, 0x01, 0x01, 0x12, 0x61, 0x0a, 0x15, 0x66, 0x69, 0x72, 0x73, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x48, 0x03, 0x52,
	0x15, 0x66, 0x69, 0x72, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x88, 0x01, 0x01, 0x12, 0x67, 0x0a, 0x17, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6d, 0x65, 0x6f,
	0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x48, 0x04, 0x52, 0x17, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4d, 0x6f, 0x6d,
	0x65, 0x6e, 0x74, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x88,
	0x01, 0x01, 0x12, 0x5e, 0x0a, 0x14, 0x66, 0x69, 0x72, 0x73, 0x74, 0x43, 0x61, 0x74, 0x50, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x50, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x48, 0x05, 0x52, 0x14, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x43, 0x61, 0x74, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x88,
	0x01, 0x01, 0x12, 0x45, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4e,
	0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x48, 0x06, 0x52, 0x0b, 0x67, 0x65, 0x74, 0x4e, 0x65,
	0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x42, 0x18, 0x0a, 0x16, 0x5f,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x42, 0x1a, 0x0a, 0x18, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4d,
	0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x43, 0x61, 0x74, 0x50, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x67,
	0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x79, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x68, 0x70, 0x6f,
	0x6c, 0x61, 0x72, 0x69, 0x73, 0x2e, 0x69, 0x64, 0x6c, 0x67, 0x65, 0x6e, 0x2e, 0x6d, 0x65, 0x6f,
	0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x61, 0x70, 0x69, 0x42, 0x08, 0x41,
	0x67, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x68, 0x2d, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x69, 0x64, 0x6c, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x65,
	0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meowchat_core_api_agg_proto_rawDescOnce sync.Once
	file_meowchat_core_api_agg_proto_rawDescData = file_meowchat_core_api_agg_proto_rawDesc
)

func file_meowchat_core_api_agg_proto_rawDescGZIP() []byte {
	file_meowchat_core_api_agg_proto_rawDescOnce.Do(func() {
		file_meowchat_core_api_agg_proto_rawDescData = protoimpl.X.CompressGZIP(file_meowchat_core_api_agg_proto_rawDescData)
	})
	return file_meowchat_core_api_agg_proto_rawDescData
}

var file_meowchat_core_api_agg_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_meowchat_core_api_agg_proto_goTypes = []interface{}{
	(*PrefetchReq)(nil),           // 0: meowchat.core_api.PrefetchReq
	(*PrefetchResp)(nil),          // 1: meowchat.core_api.PrefetchResp
	(*SignInResp)(nil),            // 2: meowchat.core_api.SignInResp
	(*GetUserInfoResp)(nil),       // 3: meowchat.core_api.GetUserInfoResp
	(*ListCommunityResp)(nil),     // 4: meowchat.core_api.ListCommunityResp
	(*GetPostPreviewsResp)(nil),   // 5: meowchat.core_api.GetPostPreviewsResp
	(*GetMomentPreviewsResp)(nil), // 6: meowchat.core_api.GetMomentPreviewsResp
	(*GetCatPreviewsResp)(nil),    // 7: meowchat.core_api.GetCatPreviewsResp
	(*GetNewsResp)(nil),           // 8: meowchat.core_api.GetNewsResp
}
var file_meowchat_core_api_agg_proto_depIdxs = []int32{
	2, // 0: meowchat.core_api.PrefetchResp.signInResp:type_name -> meowchat.core_api.SignInResp
	3, // 1: meowchat.core_api.PrefetchResp.getUserInfoResp:type_name -> meowchat.core_api.GetUserInfoResp
	4, // 2: meowchat.core_api.PrefetchResp.listCommunityResp:type_name -> meowchat.core_api.ListCommunityResp
	5, // 3: meowchat.core_api.PrefetchResp.firstPostPreviewsResp:type_name -> meowchat.core_api.GetPostPreviewsResp
	6, // 4: meowchat.core_api.PrefetchResp.firstMomentPreviewsResp:type_name -> meowchat.core_api.GetMomentPreviewsResp
	7, // 5: meowchat.core_api.PrefetchResp.firstCatPreviewsResp:type_name -> meowchat.core_api.GetCatPreviewsResp
	8, // 6: meowchat.core_api.PrefetchResp.getNewsResp:type_name -> meowchat.core_api.GetNewsResp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_meowchat_core_api_agg_proto_init() }
func file_meowchat_core_api_agg_proto_init() {
	if File_meowchat_core_api_agg_proto != nil {
		return
	}
	file_meowchat_core_api_user_proto_init()
	file_meowchat_core_api_auth_proto_init()
	file_meowchat_core_api_system_proto_init()
	file_meowchat_core_api_post_proto_init()
	file_meowchat_core_api_moment_proto_init()
	file_meowchat_core_api_collection_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_meowchat_core_api_agg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrefetchReq); i {
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
		file_meowchat_core_api_agg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrefetchResp); i {
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
	file_meowchat_core_api_agg_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_meowchat_core_api_agg_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_meowchat_core_api_agg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meowchat_core_api_agg_proto_goTypes,
		DependencyIndexes: file_meowchat_core_api_agg_proto_depIdxs,
		MessageInfos:      file_meowchat_core_api_agg_proto_msgTypes,
	}.Build()
	File_meowchat_core_api_agg_proto = out.File
	file_meowchat_core_api_agg_proto_rawDesc = nil
	file_meowchat_core_api_agg_proto_goTypes = nil
	file_meowchat_core_api_agg_proto_depIdxs = nil
}

var _ context.Context
