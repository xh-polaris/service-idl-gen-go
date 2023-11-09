// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: meowchat/system/common.proto

package system

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

type RoleType int32

const (
	RoleType_TypeUnknown        RoleType = 0
	RoleType_TypeNormalUser     RoleType = 1
	RoleType_TypeCommunityAdmin RoleType = 2
	RoleType_TypeSuperAdmin     RoleType = 3
	RoleType_TypeDeveloper      RoleType = 4
)

// Enum value maps for RoleType.
var (
	RoleType_name = map[int32]string{
		0: "TypeUnknown",
		1: "TypeNormalUser",
		2: "TypeCommunityAdmin",
		3: "TypeSuperAdmin",
		4: "TypeDeveloper",
	}
	RoleType_value = map[string]int32{
		"TypeUnknown":        0,
		"TypeNormalUser":     1,
		"TypeCommunityAdmin": 2,
		"TypeSuperAdmin":     3,
		"TypeDeveloper":      4,
	}
)

func (x RoleType) Enum() *RoleType {
	p := new(RoleType)
	*p = x
	return p
}

func (x RoleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoleType) Descriptor() protoreflect.EnumDescriptor {
	return file_meowchat_system_common_proto_enumTypes[0].Descriptor()
}

func (RoleType) Type() protoreflect.EnumType {
	return &file_meowchat_system_common_proto_enumTypes[0]
}

func (x RoleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoleType.Descriptor instead.
func (RoleType) EnumDescriptor() ([]byte, []int) {
	return file_meowchat_system_common_proto_rawDescGZIP(), []int{0}
}

type ApplyStatus int32

const (
	ApplyStatus_StatusUnknown     ApplyStatus = 0
	ApplyStatus_StatusUnprocessed ApplyStatus = 1
	ApplyStatus_StatusAccepted    ApplyStatus = 2
	ApplyStatus_StatusRejected    ApplyStatus = 3
)

// Enum value maps for ApplyStatus.
var (
	ApplyStatus_name = map[int32]string{
		0: "StatusUnknown",
		1: "StatusUnprocessed",
		2: "StatusAccepted",
		3: "StatusRejected",
	}
	ApplyStatus_value = map[string]int32{
		"StatusUnknown":     0,
		"StatusUnprocessed": 1,
		"StatusAccepted":    2,
		"StatusRejected":    3,
	}
)

func (x ApplyStatus) Enum() *ApplyStatus {
	p := new(ApplyStatus)
	*p = x
	return p
}

func (x ApplyStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApplyStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_meowchat_system_common_proto_enumTypes[1].Descriptor()
}

func (ApplyStatus) Type() protoreflect.EnumType {
	return &file_meowchat_system_common_proto_enumTypes[1]
}

func (x ApplyStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApplyStatus.Descriptor instead.
func (ApplyStatus) EnumDescriptor() ([]byte, []int) {
	return file_meowchat_system_common_proto_rawDescGZIP(), []int{1}
}

type NotificationType int32

const (
	NotificationType_TypeOther          NotificationType = 0
	NotificationType_TypeContentLike    NotificationType = 1
	NotificationType_TypeUserLike       NotificationType = 2
	NotificationType_TypeContentComment NotificationType = 3
)

// Enum value maps for NotificationType.
var (
	NotificationType_name = map[int32]string{
		0: "TypeOther",
		1: "TypeContentLike",
		2: "TypeUserLike",
		3: "TypeContentComment",
	}
	NotificationType_value = map[string]int32{
		"TypeOther":          0,
		"TypeContentLike":    1,
		"TypeUserLike":       2,
		"TypeContentComment": 3,
	}
)

func (x NotificationType) Enum() *NotificationType {
	p := new(NotificationType)
	*p = x
	return p
}

func (x NotificationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationType) Descriptor() protoreflect.EnumDescriptor {
	return file_meowchat_system_common_proto_enumTypes[2].Descriptor()
}

func (NotificationType) Type() protoreflect.EnumType {
	return &file_meowchat_system_common_proto_enumTypes[2]
}

func (x NotificationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationType.Descriptor instead.
func (NotificationType) EnumDescriptor() ([]byte, []int) {
	return file_meowchat_system_common_proto_rawDescGZIP(), []int{2}
}

type Notice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CommunityId string `protobuf:"bytes,2,opt,name=communityId,proto3" json:"communityId,omitempty"`
	Text        string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	CreateAt    int64  `protobuf:"varint,4,opt,name=createAt,proto3" json:"createAt,omitempty"`
	UpdateAt    int64  `protobuf:"varint,5,opt,name=updateAt,proto3" json:"updateAt,omitempty"`
}

func (x *Notice) Reset() {
	*x = Notice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_system_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notice) ProtoMessage() {}

func (x *Notice) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_system_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notice.ProtoReflect.Descriptor instead.
func (*Notice) Descriptor() ([]byte, []int) {
	return file_meowchat_system_common_proto_rawDescGZIP(), []int{0}
}

func (x *Notice) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Notice) GetCommunityId() string {
	if x != nil {
		return x.CommunityId
	}
	return ""
}

func (x *Notice) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Notice) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *Notice) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

type News struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CommunityId string `protobuf:"bytes,2,opt,name=communityId,proto3" json:"communityId,omitempty"`
	ImageUrl    string `protobuf:"bytes,3,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	LinkUrl     string `protobuf:"bytes,4,opt,name=linkUrl,proto3" json:"linkUrl,omitempty"`
	Type        string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	IsPublic    int64  `protobuf:"varint,6,opt,name=isPublic,proto3" json:"isPublic,omitempty"`
}

func (x *News) Reset() {
	*x = News{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_system_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *News) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*News) ProtoMessage() {}

func (x *News) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_system_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use News.ProtoReflect.Descriptor instead.
func (*News) Descriptor() ([]byte, []int) {
	return file_meowchat_system_common_proto_rawDescGZIP(), []int{1}
}

func (x *News) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *News) GetCommunityId() string {
	if x != nil {
		return x.CommunityId
	}
	return ""
}

func (x *News) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *News) GetLinkUrl() string {
	if x != nil {
		return x.LinkUrl
	}
	return ""
}

func (x *News) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *News) GetIsPublic() int64 {
	if x != nil {
		return x.IsPublic
	}
	return 0
}

type Admin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CommunityId string `protobuf:"bytes,2,opt,name=communityId,proto3" json:"communityId,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Title       string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Phone       string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	Wechat      string `protobuf:"bytes,6,opt,name=wechat,proto3" json:"wechat,omitempty"`
	AvatarUrl   string `protobuf:"bytes,7,opt,name=avatarUrl,proto3" json:"avatarUrl,omitempty"`
}

func (x *Admin) Reset() {
	*x = Admin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_system_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Admin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Admin) ProtoMessage() {}

func (x *Admin) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_system_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Admin.ProtoReflect.Descriptor instead.
func (*Admin) Descriptor() ([]byte, []int) {
	return file_meowchat_system_common_proto_rawDescGZIP(), []int{2}
}

func (x *Admin) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Admin) GetCommunityId() string {
	if x != nil {
		return x.CommunityId
	}
	return ""
}

func (x *Admin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Admin) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Admin) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Admin) GetWechat() string {
	if x != nil {
		return x.Wechat
	}
	return ""
}

func (x *Admin) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

type Community struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ParentId string `protobuf:"bytes,3,opt,name=parentId,proto3" json:"parentId,omitempty"`
}

func (x *Community) Reset() {
	*x = Community{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_system_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Community) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Community) ProtoMessage() {}

func (x *Community) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_system_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Community.ProtoReflect.Descriptor instead.
func (*Community) Descriptor() ([]byte, []int) {
	return file_meowchat_system_common_proto_rawDescGZIP(), []int{3}
}

func (x *Community) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Community) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Community) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleType    RoleType `protobuf:"varint,1,opt,name=roleType,proto3,enum=meowchat.system.RoleType" json:"roleType,omitempty"`
	CommunityId *string  `protobuf:"bytes,2,opt,name=communityId,proto3,oneof" json:"communityId,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_system_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_system_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_meowchat_system_common_proto_rawDescGZIP(), []int{4}
}

func (x *Role) GetRoleType() RoleType {
	if x != nil {
		return x.RoleType
	}
	return RoleType_TypeUnknown
}

func (x *Role) GetCommunityId() string {
	if x != nil && x.CommunityId != nil {
		return *x.CommunityId
	}
	return ""
}

type Apply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplyId     string `protobuf:"bytes,1,opt,name=applyId,proto3" json:"applyId,omitempty"`
	ApplicantId string `protobuf:"bytes,2,opt,name=applicantId,proto3" json:"applicantId,omitempty"`
	CommunityId string `protobuf:"bytes,4,opt,name=communityId,proto3" json:"communityId,omitempty"`
}

func (x *Apply) Reset() {
	*x = Apply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_system_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Apply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Apply) ProtoMessage() {}

func (x *Apply) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_system_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Apply.ProtoReflect.Descriptor instead.
func (*Apply) Descriptor() ([]byte, []int) {
	return file_meowchat_system_common_proto_rawDescGZIP(), []int{5}
}

func (x *Apply) GetApplyId() string {
	if x != nil {
		return x.ApplyId
	}
	return ""
}

func (x *Apply) GetApplicantId() string {
	if x != nil {
		return x.ApplicantId
	}
	return ""
}

func (x *Apply) GetCommunityId() string {
	if x != nil {
		return x.CommunityId
	}
	return ""
}

type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotificationId string           `protobuf:"bytes,1,opt,name=notificationId,proto3" json:"notificationId,omitempty"`
	UserId         string           `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	ParenId        string           `protobuf:"bytes,3,opt,name=parenId,proto3" json:"parenId,omitempty"`
	Type           NotificationType `protobuf:"varint,4,opt,name=type,proto3,enum=meowchat.system.NotificationType" json:"type,omitempty"`
	Text           string           `protobuf:"bytes,5,opt,name=text,proto3" json:"text,omitempty"`
	CreateAt       int64            `protobuf:"varint,6,opt,name=createAt,proto3" json:"createAt,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowchat_system_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_meowchat_system_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_meowchat_system_common_proto_rawDescGZIP(), []int{6}
}

func (x *Notification) GetNotificationId() string {
	if x != nil {
		return x.NotificationId
	}
	return ""
}

func (x *Notification) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Notification) GetParenId() string {
	if x != nil {
		return x.ParenId
	}
	return ""
}

func (x *Notification) GetType() NotificationType {
	if x != nil {
		return x.Type
	}
	return NotificationType_TypeOther
}

func (x *Notification) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Notification) GetCreateAt() int64 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

var File_meowchat_system_common_proto protoreflect.FileDescriptor

var file_meowchat_system_common_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x22,
	0x86, 0x01, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x22, 0x9e, 0x01, 0x0a, 0x04, 0x4e, 0x65, 0x77,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12,
	0x18, 0x0a, 0x07, 0x6c, 0x69, 0x6e, 0x6b, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6c, 0x69, 0x6e, 0x6b, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22, 0xaf, 0x01, 0x0a, 0x05, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x22, 0x4b, 0x0a, 0x09, 0x43,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x74, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65,
	0x12, 0x35, 0x0a, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x72,
	0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0e,
	0x0a, 0x0c, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x49, 0x64, 0x22, 0x65,
	0x0a, 0x05, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x6c, 0x79,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x22, 0xcf, 0x01, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x49, 0x64,
	0x12, 0x35, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21,
	0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x2a, 0x6e, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x79, 0x70, 0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x6f, 0x72, 0x6d,
	0x61, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x79, 0x70, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x10, 0x02,
	0x12, 0x12, 0x0a, 0x0e, 0x54, 0x79, 0x70, 0x65, 0x53, 0x75, 0x70, 0x65, 0x72, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x79, 0x70, 0x65, 0x44, 0x65, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x65, 0x72, 0x10, 0x04, 0x2a, 0x5f, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x55, 0x6e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x10, 0x01,
	0x12, 0x12, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x65, 0x64, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10, 0x03, 0x2a, 0x60, 0x0a, 0x10, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09,
	0x54, 0x79, 0x70, 0x65, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x54,
	0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x10, 0x01,
	0x12, 0x10, 0x0a, 0x0c, 0x54, 0x79, 0x70, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x6b, 0x65,
	0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x03, 0x42, 0x79, 0x0a, 0x24, 0x63, 0x6f,
	0x6d, 0x2e, 0x78, 0x68, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2e, 0x69, 0x64, 0x6c, 0x67,
	0x65, 0x6e, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x68,
	0x2d, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2d, 0x69, 0x64, 0x6c, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x6b, 0x69, 0x74, 0x65,
	0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meowchat_system_common_proto_rawDescOnce sync.Once
	file_meowchat_system_common_proto_rawDescData = file_meowchat_system_common_proto_rawDesc
)

func file_meowchat_system_common_proto_rawDescGZIP() []byte {
	file_meowchat_system_common_proto_rawDescOnce.Do(func() {
		file_meowchat_system_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_meowchat_system_common_proto_rawDescData)
	})
	return file_meowchat_system_common_proto_rawDescData
}

var file_meowchat_system_common_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_meowchat_system_common_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_meowchat_system_common_proto_goTypes = []interface{}{
	(RoleType)(0),         // 0: meowchat.system.RoleType
	(ApplyStatus)(0),      // 1: meowchat.system.ApplyStatus
	(NotificationType)(0), // 2: meowchat.system.NotificationType
	(*Notice)(nil),        // 3: meowchat.system.Notice
	(*News)(nil),          // 4: meowchat.system.News
	(*Admin)(nil),         // 5: meowchat.system.Admin
	(*Community)(nil),     // 6: meowchat.system.Community
	(*Role)(nil),          // 7: meowchat.system.Role
	(*Apply)(nil),         // 8: meowchat.system.Apply
	(*Notification)(nil),  // 9: meowchat.system.Notification
}
var file_meowchat_system_common_proto_depIdxs = []int32{
	0, // 0: meowchat.system.Role.roleType:type_name -> meowchat.system.RoleType
	2, // 1: meowchat.system.Notification.type:type_name -> meowchat.system.NotificationType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_meowchat_system_common_proto_init() }
func file_meowchat_system_common_proto_init() {
	if File_meowchat_system_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_meowchat_system_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notice); i {
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
		file_meowchat_system_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*News); i {
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
		file_meowchat_system_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Admin); i {
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
		file_meowchat_system_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Community); i {
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
		file_meowchat_system_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_meowchat_system_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Apply); i {
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
		file_meowchat_system_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notification); i {
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
	file_meowchat_system_common_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_meowchat_system_common_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meowchat_system_common_proto_goTypes,
		DependencyIndexes: file_meowchat_system_common_proto_depIdxs,
		EnumInfos:         file_meowchat_system_common_proto_enumTypes,
		MessageInfos:      file_meowchat_system_common_proto_msgTypes,
	}.Build()
	File_meowchat_system_common_proto = out.File
	file_meowchat_system_common_proto_rawDesc = nil
	file_meowchat_system_common_proto_goTypes = nil
	file_meowchat_system_common_proto_depIdxs = nil
}

var _ context.Context
