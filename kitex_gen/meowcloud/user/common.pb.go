// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: meowcloud/user/common.proto

package user

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

type UserPreview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Avatar   string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *UserPreview) Reset() {
	*x = UserPreview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowcloud_user_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPreview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPreview) ProtoMessage() {}

func (x *UserPreview) ProtoReflect() protoreflect.Message {
	mi := &file_meowcloud_user_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPreview.ProtoReflect.Descriptor instead.
func (*UserPreview) Descriptor() ([]byte, []int) {
	return file_meowcloud_user_common_proto_rawDescGZIP(), []int{0}
}

func (x *UserPreview) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserPreview) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UserPreview) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Nickname           string            `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Bio                string            `protobuf:"bytes,3,opt,name=bio,proto3" json:"bio,omitempty"`
	Avatar             string            `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	UpdatedAt          int64             `protobuf:"varint,5,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	CreatedAt          int64             `protobuf:"varint,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	DeletedAt          int64             `protobuf:"varint,7,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
	Membership         *User_Membership  `protobuf:"bytes,8,opt,name=membership,proto3" json:"membership,omitempty"`
	TeamCount          int32             `protobuf:"varint,9,opt,name=teamCount,proto3" json:"teamCount,omitempty"`
	TeamIds            []string          `protobuf:"bytes,10,rep,name=teamIds,proto3" json:"teamIds,omitempty"`
	MyAlbumCount       int32             `protobuf:"varint,11,opt,name=myAlbumCount,proto3" json:"myAlbumCount,omitempty"`
	MyAlbumIds         []int32           `protobuf:"varint,12,rep,packed,name=myAlbumIds,proto3" json:"myAlbumIds,omitempty"`
	FollowedAlbumCount int32             `protobuf:"varint,13,opt,name=followedAlbumCount,proto3" json:"followedAlbumCount,omitempty"`
	FollowedAlbumIds   []int32           `protobuf:"varint,14,rep,packed,name=followedAlbumIds,proto3" json:"followedAlbumIds,omitempty"`
	StorageInfo        *User_StorageInfo `protobuf:"bytes,15,opt,name=storageInfo,proto3" json:"storageInfo,omitempty"`
	Points             int32             `protobuf:"varint,16,opt,name=points,proto3" json:"points,omitempty"`
	Achievements       []string          `protobuf:"bytes,17,rep,name=achievements,proto3" json:"achievements,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowcloud_user_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_meowcloud_user_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_meowcloud_user_common_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *User) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *User) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *User) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

func (x *User) GetMembership() *User_Membership {
	if x != nil {
		return x.Membership
	}
	return nil
}

func (x *User) GetTeamCount() int32 {
	if x != nil {
		return x.TeamCount
	}
	return 0
}

func (x *User) GetTeamIds() []string {
	if x != nil {
		return x.TeamIds
	}
	return nil
}

func (x *User) GetMyAlbumCount() int32 {
	if x != nil {
		return x.MyAlbumCount
	}
	return 0
}

func (x *User) GetMyAlbumIds() []int32 {
	if x != nil {
		return x.MyAlbumIds
	}
	return nil
}

func (x *User) GetFollowedAlbumCount() int32 {
	if x != nil {
		return x.FollowedAlbumCount
	}
	return 0
}

func (x *User) GetFollowedAlbumIds() []int32 {
	if x != nil {
		return x.FollowedAlbumIds
	}
	return nil
}

func (x *User) GetStorageInfo() *User_StorageInfo {
	if x != nil {
		return x.StorageInfo
	}
	return nil
}

func (x *User) GetPoints() int32 {
	if x != nil {
		return x.Points
	}
	return 0
}

func (x *User) GetAchievements() []string {
	if x != nil {
		return x.Achievements
	}
	return nil
}

type User_Membership struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberId    string `protobuf:"bytes,1,opt,name=memberId,proto3" json:"memberId,omitempty"`
	MemberLevel int32  `protobuf:"varint,2,opt,name=memberLevel,proto3" json:"memberLevel,omitempty"` // 0: 普通用户, 1: 普通会员, 2: 高级会员
}

func (x *User_Membership) Reset() {
	*x = User_Membership{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowcloud_user_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User_Membership) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User_Membership) ProtoMessage() {}

func (x *User_Membership) ProtoReflect() protoreflect.Message {
	mi := &file_meowcloud_user_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User_Membership.ProtoReflect.Descriptor instead.
func (*User_Membership) Descriptor() ([]byte, []int) {
	return file_meowcloud_user_common_proto_rawDescGZIP(), []int{1, 0}
}

func (x *User_Membership) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *User_Membership) GetMemberLevel() int32 {
	if x != nil {
		return x.MemberLevel
	}
	return 0
}

type User_StorageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvailablePhotos int32 `protobuf:"varint,1,opt,name=availablePhotos,proto3" json:"availablePhotos,omitempty"`
	UsedPhotos      int32 `protobuf:"varint,2,opt,name=usedPhotos,proto3" json:"usedPhotos,omitempty"`
	AvailableMemory int64 `protobuf:"varint,3,opt,name=availableMemory,proto3" json:"availableMemory,omitempty"`
	UsedMemory      int64 `protobuf:"varint,4,opt,name=usedMemory,proto3" json:"usedMemory,omitempty"`
	AvailableAlbums int32 `protobuf:"varint,5,opt,name=availableAlbums,proto3" json:"availableAlbums,omitempty"`
	UsedAlbums      int32 `protobuf:"varint,6,opt,name=usedAlbums,proto3" json:"usedAlbums,omitempty"`
}

func (x *User_StorageInfo) Reset() {
	*x = User_StorageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meowcloud_user_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User_StorageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User_StorageInfo) ProtoMessage() {}

func (x *User_StorageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_meowcloud_user_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User_StorageInfo.ProtoReflect.Descriptor instead.
func (*User_StorageInfo) Descriptor() ([]byte, []int) {
	return file_meowcloud_user_common_proto_rawDescGZIP(), []int{1, 1}
}

func (x *User_StorageInfo) GetAvailablePhotos() int32 {
	if x != nil {
		return x.AvailablePhotos
	}
	return 0
}

func (x *User_StorageInfo) GetUsedPhotos() int32 {
	if x != nil {
		return x.UsedPhotos
	}
	return 0
}

func (x *User_StorageInfo) GetAvailableMemory() int64 {
	if x != nil {
		return x.AvailableMemory
	}
	return 0
}

func (x *User_StorageInfo) GetUsedMemory() int64 {
	if x != nil {
		return x.UsedMemory
	}
	return 0
}

func (x *User_StorageInfo) GetAvailableAlbums() int32 {
	if x != nil {
		return x.AvailableAlbums
	}
	return 0
}

func (x *User_StorageInfo) GetUsedAlbums() int32 {
	if x != nil {
		return x.UsedAlbums
	}
	return 0
}

var File_meowcloud_user_common_proto protoreflect.FileDescriptor

var file_meowcloud_user_common_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6d,
	0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x22, 0x51, 0x0a,
	0x0b, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x22, 0x89, 0x07, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12,
	0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3f, 0x0a, 0x0a, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x0a,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65,
	0x61, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74,
	0x65, 0x61, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d,
	0x49, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x49,
	0x64, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x79, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x79, 0x41, 0x6c, 0x62, 0x75,
	0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x79, 0x41, 0x6c, 0x62, 0x75,
	0x6d, 0x49, 0x64, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a, 0x6d, 0x79, 0x41, 0x6c,
	0x62, 0x75, 0x6d, 0x49, 0x64, 0x73, 0x12, 0x2e, 0x0a, 0x12, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x12, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x41, 0x6c, 0x62, 0x75,
	0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x10, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x49,
	0x64, 0x73, 0x12, 0x42, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x22,
	0x0a, 0x0c, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x11,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x1a, 0x4a, 0x0a, 0x0a, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70,
	0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x1a, 0xeb,
	0x01, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28,
	0x0a, 0x0f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x64,
	0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x75, 0x73,
	0x65, 0x64, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x64, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x64, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x41,
	0x6c, 0x62, 0x75, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x61, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x12, 0x1e, 0x0a, 0x0a,
	0x75, 0x73, 0x65, 0x64, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x75, 0x73, 0x65, 0x64, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x42, 0x77, 0x0a, 0x23,
	0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x68, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2e, 0x69, 0x64,
	0x6c, 0x67, 0x65, 0x6e, 0x2e, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78,
	0x68, 0x2d, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2d, 0x69, 0x64, 0x6c, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x6b, 0x69, 0x74,
	0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meowcloud_user_common_proto_rawDescOnce sync.Once
	file_meowcloud_user_common_proto_rawDescData = file_meowcloud_user_common_proto_rawDesc
)

func file_meowcloud_user_common_proto_rawDescGZIP() []byte {
	file_meowcloud_user_common_proto_rawDescOnce.Do(func() {
		file_meowcloud_user_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_meowcloud_user_common_proto_rawDescData)
	})
	return file_meowcloud_user_common_proto_rawDescData
}

var file_meowcloud_user_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_meowcloud_user_common_proto_goTypes = []interface{}{
	(*UserPreview)(nil),      // 0: meowcloud.user.UserPreview
	(*User)(nil),             // 1: meowcloud.user.User
	(*User_Membership)(nil),  // 2: meowcloud.user.User.Membership
	(*User_StorageInfo)(nil), // 3: meowcloud.user.User.StorageInfo
}
var file_meowcloud_user_common_proto_depIdxs = []int32{
	2, // 0: meowcloud.user.User.membership:type_name -> meowcloud.user.User.Membership
	3, // 1: meowcloud.user.User.storageInfo:type_name -> meowcloud.user.User.StorageInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_meowcloud_user_common_proto_init() }
func file_meowcloud_user_common_proto_init() {
	if File_meowcloud_user_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_meowcloud_user_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPreview); i {
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
		file_meowcloud_user_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_meowcloud_user_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User_Membership); i {
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
		file_meowcloud_user_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User_StorageInfo); i {
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
			RawDescriptor: file_meowcloud_user_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meowcloud_user_common_proto_goTypes,
		DependencyIndexes: file_meowcloud_user_common_proto_depIdxs,
		MessageInfos:      file_meowcloud_user_common_proto_msgTypes,
	}.Build()
	File_meowcloud_user_common_proto = out.File
	file_meowcloud_user_common_proto_rawDesc = nil
	file_meowcloud_user_common_proto_goTypes = nil
	file_meowcloud_user_common_proto_depIdxs = nil
}

var _ context.Context