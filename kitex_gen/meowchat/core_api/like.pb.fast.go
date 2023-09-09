// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package core_api

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
	user "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/user"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *DoLikeReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DoLikeReq[number], err)
}

func (x *DoLikeReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.TargetId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DoLikeReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.TargetType = user.LikeType(v)
	return offset, nil
}

func (x *DoLikeResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DoLikeResp[number], err)
}

func (x *DoLikeResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.IsFirst, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *GetUserLikedReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetUserLikedReq[number], err)
}

func (x *GetUserLikedReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.TargetId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetUserLikedReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.TargetType = user.LikeType(v)
	return offset, nil
}

func (x *GetUserLikedResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetUserLikedResp[number], err)
}

func (x *GetUserLikedResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Liked, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *GetLikedCountReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetLikedCountReq[number], err)
}

func (x *GetLikedCountReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.TargetId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetLikedCountReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.TargetType = user.LikeType(v)
	return offset, nil
}

func (x *GetLikedCountResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetLikedCountResp[number], err)
}

func (x *GetLikedCountResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Count, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetUserLikesReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetUserLikesReq[number], err)
}

func (x *GetUserLikesReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetUserLikesReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.TargetType = user.LikeType(v)
	return offset, nil
}

func (x *GetUserLikesResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetUserLikesResp[number], err)
}

func (x *GetUserLikesResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v user.Like
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Likes = append(x.Likes, &v)
	return offset, nil
}

func (x *GetLikedUsersReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetLikedUsersReq[number], err)
}

func (x *GetLikedUsersReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.TargetId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetLikedUsersReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.TargetType = user.LikeType(v)
	return offset, nil
}

func (x *GetLikedUsersResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetLikedUsersResp[number], err)
}

func (x *GetLikedUsersResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v user.UserPreview
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Users = append(x.Users, &v)
	return offset, nil
}

func (x *DoLikeReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *DoLikeReq) fastWriteField1(buf []byte) (offset int) {
	if x.TargetId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetTargetId())
	return offset
}

func (x *DoLikeReq) fastWriteField2(buf []byte) (offset int) {
	if x.TargetType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, int32(x.GetTargetType()))
	return offset
}

func (x *DoLikeResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DoLikeResp) fastWriteField1(buf []byte) (offset int) {
	if !x.IsFirst {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetIsFirst())
	return offset
}

func (x *GetUserLikedReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetUserLikedReq) fastWriteField1(buf []byte) (offset int) {
	if x.TargetId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetTargetId())
	return offset
}

func (x *GetUserLikedReq) fastWriteField2(buf []byte) (offset int) {
	if x.TargetType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, int32(x.GetTargetType()))
	return offset
}

func (x *GetUserLikedResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetUserLikedResp) fastWriteField1(buf []byte) (offset int) {
	if !x.Liked {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetLiked())
	return offset
}

func (x *GetLikedCountReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetLikedCountReq) fastWriteField1(buf []byte) (offset int) {
	if x.TargetId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetTargetId())
	return offset
}

func (x *GetLikedCountReq) fastWriteField2(buf []byte) (offset int) {
	if x.TargetType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, int32(x.GetTargetType()))
	return offset
}

func (x *GetLikedCountResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetLikedCountResp) fastWriteField1(buf []byte) (offset int) {
	if x.Count == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetCount())
	return offset
}

func (x *GetUserLikesReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetUserLikesReq) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *GetUserLikesReq) fastWriteField2(buf []byte) (offset int) {
	if x.TargetType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, int32(x.GetTargetType()))
	return offset
}

func (x *GetUserLikesResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetUserLikesResp) fastWriteField1(buf []byte) (offset int) {
	if x.Likes == nil {
		return offset
	}
	for i := range x.GetLikes() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetLikes()[i])
	}
	return offset
}

func (x *GetLikedUsersReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetLikedUsersReq) fastWriteField1(buf []byte) (offset int) {
	if x.TargetId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetTargetId())
	return offset
}

func (x *GetLikedUsersReq) fastWriteField2(buf []byte) (offset int) {
	if x.TargetType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, int32(x.GetTargetType()))
	return offset
}

func (x *GetLikedUsersResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetLikedUsersResp) fastWriteField1(buf []byte) (offset int) {
	if x.Users == nil {
		return offset
	}
	for i := range x.GetUsers() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetUsers()[i])
	}
	return offset
}

func (x *DoLikeReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *DoLikeReq) sizeField1() (n int) {
	if x.TargetId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetTargetId())
	return n
}

func (x *DoLikeReq) sizeField2() (n int) {
	if x.TargetType == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, int32(x.GetTargetType()))
	return n
}

func (x *DoLikeResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DoLikeResp) sizeField1() (n int) {
	if !x.IsFirst {
		return n
	}
	n += fastpb.SizeBool(1, x.GetIsFirst())
	return n
}

func (x *GetUserLikedReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetUserLikedReq) sizeField1() (n int) {
	if x.TargetId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetTargetId())
	return n
}

func (x *GetUserLikedReq) sizeField2() (n int) {
	if x.TargetType == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, int32(x.GetTargetType()))
	return n
}

func (x *GetUserLikedResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetUserLikedResp) sizeField1() (n int) {
	if !x.Liked {
		return n
	}
	n += fastpb.SizeBool(1, x.GetLiked())
	return n
}

func (x *GetLikedCountReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetLikedCountReq) sizeField1() (n int) {
	if x.TargetId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetTargetId())
	return n
}

func (x *GetLikedCountReq) sizeField2() (n int) {
	if x.TargetType == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, int32(x.GetTargetType()))
	return n
}

func (x *GetLikedCountResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetLikedCountResp) sizeField1() (n int) {
	if x.Count == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetCount())
	return n
}

func (x *GetUserLikesReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetUserLikesReq) sizeField1() (n int) {
	if x.UserId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetUserId())
	return n
}

func (x *GetUserLikesReq) sizeField2() (n int) {
	if x.TargetType == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, int32(x.GetTargetType()))
	return n
}

func (x *GetUserLikesResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetUserLikesResp) sizeField1() (n int) {
	if x.Likes == nil {
		return n
	}
	for i := range x.GetLikes() {
		n += fastpb.SizeMessage(1, x.GetLikes()[i])
	}
	return n
}

func (x *GetLikedUsersReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetLikedUsersReq) sizeField1() (n int) {
	if x.TargetId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetTargetId())
	return n
}

func (x *GetLikedUsersReq) sizeField2() (n int) {
	if x.TargetType == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, int32(x.GetTargetType()))
	return n
}

func (x *GetLikedUsersResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetLikedUsersResp) sizeField1() (n int) {
	if x.Users == nil {
		return n
	}
	for i := range x.GetUsers() {
		n += fastpb.SizeMessage(1, x.GetUsers()[i])
	}
	return n
}

var fieldIDToName_DoLikeReq = map[int32]string{
	1: "TargetId",
	2: "TargetType",
}

var fieldIDToName_DoLikeResp = map[int32]string{
	1: "IsFirst",
}

var fieldIDToName_GetUserLikedReq = map[int32]string{
	1: "TargetId",
	2: "TargetType",
}

var fieldIDToName_GetUserLikedResp = map[int32]string{
	1: "Liked",
}

var fieldIDToName_GetLikedCountReq = map[int32]string{
	1: "TargetId",
	2: "TargetType",
}

var fieldIDToName_GetLikedCountResp = map[int32]string{
	1: "Count",
}

var fieldIDToName_GetUserLikesReq = map[int32]string{
	1: "UserId",
	2: "TargetType",
}

var fieldIDToName_GetUserLikesResp = map[int32]string{
	1: "Likes",
}

var fieldIDToName_GetLikedUsersReq = map[int32]string{
	1: "TargetId",
	2: "TargetType",
}

var fieldIDToName_GetLikedUsersResp = map[int32]string{
	1: "Users",
}

var _ = user.File_meowchat_user_common_proto
