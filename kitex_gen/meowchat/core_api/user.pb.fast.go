// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package core_api

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
	basic "github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
	system "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/system"
	user "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/user"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *User) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 7:
		offset, err = x.fastReadField7(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 8:
		offset, err = x.fastReadField8(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 9:
		offset, err = x.fastReadField9(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 10:
		offset, err = x.fastReadField10(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 11:
		offset, err = x.fastReadField11(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_User[number], err)
}

func (x *User) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *User) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Nickname, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *User) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.AvatarUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *User) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Motto = &tmp
	return offset, err
}

func (x *User) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt64(buf, _type)
	x.Follower = &tmp
	return offset, err
}

func (x *User) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt64(buf, _type)
	x.Following = &tmp
	return offset, err
}

func (x *User) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt64(buf, _type)
	x.Article = &tmp
	return offset, err
}

func (x *User) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt64(buf, _type)
	x.Like = &tmp
	return offset, err
}

func (x *User) fastReadField9(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadBool(buf, _type)
	x.IsFollowing = &tmp
	return offset, err
}

func (x *User) fastReadField10(buf []byte, _type int8) (offset int, err error) {
	var v system.Role
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Roles = append(x.Roles, &v)
	return offset, nil
}

func (x *User) fastReadField11(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadBool(buf, _type)
	x.EnableDebug = &tmp
	return offset, err
}

func (x *GetUserInfoReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetUserInfoReq[number], err)
}

func (x *GetUserInfoReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.UserId = &tmp
	return offset, err
}

func (x *GetUserInfoResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetUserInfoResp[number], err)
}

func (x *GetUserInfoResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v User
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.User = &v
	return offset, nil
}

func (x *UpdateUserInfoReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 3:
		offset, err = x.fastReadField3(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UpdateUserInfoReq[number], err)
}

func (x *UpdateUserInfoReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.AvatarUrl = &tmp
	return offset, err
}

func (x *UpdateUserInfoReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Nickname = &tmp
	return offset, err
}

func (x *UpdateUserInfoReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Motto = &tmp
	return offset, err
}

func (x *UpdateUserInfoResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
}

func (x *SearchUserReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SearchUserReq[number], err)
}

func (x *SearchUserReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Keyword, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SearchUserReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v basic.PaginationOptions
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.PaginationOption = &v
	return offset, nil
}

func (x *SearchUserResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 3:
		offset, err = x.fastReadField3(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SearchUserResp[number], err)
}

func (x *SearchUserResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v User
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Users = append(x.Users, &v)
	return offset, nil
}

func (x *SearchUserResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Total, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *SearchUserResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Token, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CheckInReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
}

func (x *CheckInResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 3:
		offset, err = x.fastReadField3(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CheckInResp[number], err)
}

func (x *CheckInResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.IsFirst, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *CheckInResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.GetFishTimes, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CheckInResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.GetFishNum, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *User) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	offset += x.fastWriteField8(buf[offset:])
	offset += x.fastWriteField9(buf[offset:])
	offset += x.fastWriteField10(buf[offset:])
	offset += x.fastWriteField11(buf[offset:])
	return offset
}

func (x *User) fastWriteField1(buf []byte) (offset int) {
	if x.Id == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetId())
	return offset
}

func (x *User) fastWriteField2(buf []byte) (offset int) {
	if x.Nickname == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetNickname())
	return offset
}

func (x *User) fastWriteField3(buf []byte) (offset int) {
	if x.AvatarUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetAvatarUrl())
	return offset
}

func (x *User) fastWriteField4(buf []byte) (offset int) {
	if x.Motto == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetMotto())
	return offset
}

func (x *User) fastWriteField5(buf []byte) (offset int) {
	if x.Follower == nil {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 5, x.GetFollower())
	return offset
}

func (x *User) fastWriteField6(buf []byte) (offset int) {
	if x.Following == nil {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 6, x.GetFollowing())
	return offset
}

func (x *User) fastWriteField7(buf []byte) (offset int) {
	if x.Article == nil {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 7, x.GetArticle())
	return offset
}

func (x *User) fastWriteField8(buf []byte) (offset int) {
	if x.Like == nil {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 8, x.GetLike())
	return offset
}

func (x *User) fastWriteField9(buf []byte) (offset int) {
	if x.IsFollowing == nil {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 9, x.GetIsFollowing())
	return offset
}

func (x *User) fastWriteField10(buf []byte) (offset int) {
	if x.Roles == nil {
		return offset
	}
	for i := range x.GetRoles() {
		offset += fastpb.WriteMessage(buf[offset:], 10, x.GetRoles()[i])
	}
	return offset
}

func (x *User) fastWriteField11(buf []byte) (offset int) {
	if x.EnableDebug == nil {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 11, x.GetEnableDebug())
	return offset
}

func (x *GetUserInfoReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetUserInfoReq) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *GetUserInfoResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetUserInfoResp) fastWriteField1(buf []byte) (offset int) {
	if x.User == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetUser())
	return offset
}

func (x *UpdateUserInfoReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *UpdateUserInfoReq) fastWriteField1(buf []byte) (offset int) {
	if x.AvatarUrl == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetAvatarUrl())
	return offset
}

func (x *UpdateUserInfoReq) fastWriteField2(buf []byte) (offset int) {
	if x.Nickname == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetNickname())
	return offset
}

func (x *UpdateUserInfoReq) fastWriteField3(buf []byte) (offset int) {
	if x.Motto == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetMotto())
	return offset
}

func (x *UpdateUserInfoResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *SearchUserReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *SearchUserReq) fastWriteField1(buf []byte) (offset int) {
	if x.Keyword == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetKeyword())
	return offset
}

func (x *SearchUserReq) fastWriteField2(buf []byte) (offset int) {
	if x.PaginationOption == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetPaginationOption())
	return offset
}

func (x *SearchUserResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *SearchUserResp) fastWriteField1(buf []byte) (offset int) {
	if x.Users == nil {
		return offset
	}
	for i := range x.GetUsers() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetUsers()[i])
	}
	return offset
}

func (x *SearchUserResp) fastWriteField2(buf []byte) (offset int) {
	if x.Total == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetTotal())
	return offset
}

func (x *SearchUserResp) fastWriteField3(buf []byte) (offset int) {
	if x.Token == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetToken())
	return offset
}

func (x *CheckInReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *CheckInResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *CheckInResp) fastWriteField1(buf []byte) (offset int) {
	if !x.IsFirst {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetIsFirst())
	return offset
}

func (x *CheckInResp) fastWriteField2(buf []byte) (offset int) {
	if x.GetFishTimes == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetGetFishTimes())
	return offset
}

func (x *CheckInResp) fastWriteField3(buf []byte) (offset int) {
	if x.GetFishNum == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetGetFishNum())
	return offset
}

func (x *User) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	n += x.sizeField8()
	n += x.sizeField9()
	n += x.sizeField10()
	n += x.sizeField11()
	return n
}

func (x *User) sizeField1() (n int) {
	if x.Id == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetId())
	return n
}

func (x *User) sizeField2() (n int) {
	if x.Nickname == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetNickname())
	return n
}

func (x *User) sizeField3() (n int) {
	if x.AvatarUrl == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetAvatarUrl())
	return n
}

func (x *User) sizeField4() (n int) {
	if x.Motto == nil {
		return n
	}
	n += fastpb.SizeString(4, x.GetMotto())
	return n
}

func (x *User) sizeField5() (n int) {
	if x.Follower == nil {
		return n
	}
	n += fastpb.SizeInt64(5, x.GetFollower())
	return n
}

func (x *User) sizeField6() (n int) {
	if x.Following == nil {
		return n
	}
	n += fastpb.SizeInt64(6, x.GetFollowing())
	return n
}

func (x *User) sizeField7() (n int) {
	if x.Article == nil {
		return n
	}
	n += fastpb.SizeInt64(7, x.GetArticle())
	return n
}

func (x *User) sizeField8() (n int) {
	if x.Like == nil {
		return n
	}
	n += fastpb.SizeInt64(8, x.GetLike())
	return n
}

func (x *User) sizeField9() (n int) {
	if x.IsFollowing == nil {
		return n
	}
	n += fastpb.SizeBool(9, x.GetIsFollowing())
	return n
}

func (x *User) sizeField10() (n int) {
	if x.Roles == nil {
		return n
	}
	for i := range x.GetRoles() {
		n += fastpb.SizeMessage(10, x.GetRoles()[i])
	}
	return n
}

func (x *User) sizeField11() (n int) {
	if x.EnableDebug == nil {
		return n
	}
	n += fastpb.SizeBool(11, x.GetEnableDebug())
	return n
}

func (x *GetUserInfoReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetUserInfoReq) sizeField1() (n int) {
	if x.UserId == nil {
		return n
	}
	n += fastpb.SizeString(1, x.GetUserId())
	return n
}

func (x *GetUserInfoResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetUserInfoResp) sizeField1() (n int) {
	if x.User == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetUser())
	return n
}

func (x *UpdateUserInfoReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *UpdateUserInfoReq) sizeField1() (n int) {
	if x.AvatarUrl == nil {
		return n
	}
	n += fastpb.SizeString(1, x.GetAvatarUrl())
	return n
}

func (x *UpdateUserInfoReq) sizeField2() (n int) {
	if x.Nickname == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetNickname())
	return n
}

func (x *UpdateUserInfoReq) sizeField3() (n int) {
	if x.Motto == nil {
		return n
	}
	n += fastpb.SizeString(3, x.GetMotto())
	return n
}

func (x *UpdateUserInfoResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *SearchUserReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *SearchUserReq) sizeField1() (n int) {
	if x.Keyword == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetKeyword())
	return n
}

func (x *SearchUserReq) sizeField2() (n int) {
	if x.PaginationOption == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetPaginationOption())
	return n
}

func (x *SearchUserResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *SearchUserResp) sizeField1() (n int) {
	if x.Users == nil {
		return n
	}
	for i := range x.GetUsers() {
		n += fastpb.SizeMessage(1, x.GetUsers()[i])
	}
	return n
}

func (x *SearchUserResp) sizeField2() (n int) {
	if x.Total == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetTotal())
	return n
}

func (x *SearchUserResp) sizeField3() (n int) {
	if x.Token == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetToken())
	return n
}

func (x *CheckInReq) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *CheckInResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *CheckInResp) sizeField1() (n int) {
	if !x.IsFirst {
		return n
	}
	n += fastpb.SizeBool(1, x.GetIsFirst())
	return n
}

func (x *CheckInResp) sizeField2() (n int) {
	if x.GetFishTimes == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetGetFishTimes())
	return n
}

func (x *CheckInResp) sizeField3() (n int) {
	if x.GetFishNum == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetGetFishNum())
	return n
}

var fieldIDToName_User = map[int32]string{
	1:  "Id",
	2:  "Nickname",
	3:  "AvatarUrl",
	4:  "Motto",
	5:  "Follower",
	6:  "Following",
	7:  "Article",
	8:  "Like",
	9:  "IsFollowing",
	10: "Roles",
	11: "EnableDebug",
}

var fieldIDToName_GetUserInfoReq = map[int32]string{
	1: "UserId",
}

var fieldIDToName_GetUserInfoResp = map[int32]string{
	1: "User",
}

var fieldIDToName_UpdateUserInfoReq = map[int32]string{
	1: "AvatarUrl",
	2: "Nickname",
	3: "Motto",
}

var fieldIDToName_UpdateUserInfoResp = map[int32]string{}

var fieldIDToName_SearchUserReq = map[int32]string{
	1: "Keyword",
	2: "PaginationOption",
}

var fieldIDToName_SearchUserResp = map[int32]string{
	1: "Users",
	2: "Total",
	3: "Token",
}

var fieldIDToName_CheckInReq = map[int32]string{}

var fieldIDToName_CheckInResp = map[int32]string{
	1: "IsFirst",
	2: "GetFishTimes",
	3: "GetFishNum",
}

var _ = basic.File_basic_pagination_proto
var _ = user.File_meowchat_user_common_proto
var _ = system.File_meowchat_system_common_proto
