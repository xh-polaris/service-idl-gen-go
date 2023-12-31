// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package core_api

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
	system "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/system"
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

var _ = system.File_meowchat_system_common_proto
