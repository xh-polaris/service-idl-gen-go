// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package user

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
	basic "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/basic"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *UserDetail) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UserDetail[number], err)
}

func (x *UserDetail) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UserDetail) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.AvatarUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UserDetail) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Nickname, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UserDetail) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Motto, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetUserReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetUserReq[number], err)
}

func (x *GetUserReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetUserResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetUserResp[number], err)
}

func (x *GetUserResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v basic.UserPreview
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.User = &v
	return offset, nil
}

func (x *GetUserDetailReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetUserDetailReq[number], err)
}

func (x *GetUserDetailReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetUserDetailResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetUserDetailResp[number], err)
}

func (x *GetUserDetailResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v UserDetail
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.User = &v
	return offset, nil
}

func (x *UpdateUserReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UpdateUserReq[number], err)
}

func (x *UpdateUserReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v UserDetail
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.User = &v
	return offset, nil
}

func (x *UpdateUserResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	x.Nickname, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SearchUserReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt64(buf, _type)
	x.Offset = &tmp
	return offset, err
}

func (x *SearchUserReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadInt64(buf, _type)
	x.Limit = &tmp
	return offset, err
}

func (x *SearchUserReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadBool(buf, _type)
	x.Backward = &tmp
	return offset, err
}

func (x *SearchUserReq) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.LastToken = &tmp
	return offset, err
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
	var v basic.UserPreview
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

func (x *UserDetail) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *UserDetail) fastWriteField1(buf []byte) (offset int) {
	if x.Id == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetId())
	return offset
}

func (x *UserDetail) fastWriteField2(buf []byte) (offset int) {
	if x.AvatarUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetAvatarUrl())
	return offset
}

func (x *UserDetail) fastWriteField3(buf []byte) (offset int) {
	if x.Nickname == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetNickname())
	return offset
}

func (x *UserDetail) fastWriteField4(buf []byte) (offset int) {
	if x.Motto == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetMotto())
	return offset
}

func (x *GetUserReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetUserReq) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *GetUserResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetUserResp) fastWriteField1(buf []byte) (offset int) {
	if x.User == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetUser())
	return offset
}

func (x *GetUserDetailReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetUserDetailReq) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *GetUserDetailResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetUserDetailResp) fastWriteField1(buf []byte) (offset int) {
	if x.User == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetUser())
	return offset
}

func (x *UpdateUserReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *UpdateUserReq) fastWriteField1(buf []byte) (offset int) {
	if x.User == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetUser())
	return offset
}

func (x *UpdateUserResp) FastWrite(buf []byte) (offset int) {
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
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	return offset
}

func (x *SearchUserReq) fastWriteField1(buf []byte) (offset int) {
	if x.Nickname == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetNickname())
	return offset
}

func (x *SearchUserReq) fastWriteField2(buf []byte) (offset int) {
	if x.Offset == nil {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetOffset())
	return offset
}

func (x *SearchUserReq) fastWriteField3(buf []byte) (offset int) {
	if x.Limit == nil {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetLimit())
	return offset
}

func (x *SearchUserReq) fastWriteField4(buf []byte) (offset int) {
	if x.Backward == nil {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 4, x.GetBackward())
	return offset
}

func (x *SearchUserReq) fastWriteField5(buf []byte) (offset int) {
	if x.LastToken == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetLastToken())
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

func (x *UserDetail) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *UserDetail) sizeField1() (n int) {
	if x.Id == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetId())
	return n
}

func (x *UserDetail) sizeField2() (n int) {
	if x.AvatarUrl == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetAvatarUrl())
	return n
}

func (x *UserDetail) sizeField3() (n int) {
	if x.Nickname == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetNickname())
	return n
}

func (x *UserDetail) sizeField4() (n int) {
	if x.Motto == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetMotto())
	return n
}

func (x *GetUserReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetUserReq) sizeField1() (n int) {
	if x.UserId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetUserId())
	return n
}

func (x *GetUserResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetUserResp) sizeField1() (n int) {
	if x.User == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetUser())
	return n
}

func (x *GetUserDetailReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetUserDetailReq) sizeField1() (n int) {
	if x.UserId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetUserId())
	return n
}

func (x *GetUserDetailResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetUserDetailResp) sizeField1() (n int) {
	if x.User == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetUser())
	return n
}

func (x *UpdateUserReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *UpdateUserReq) sizeField1() (n int) {
	if x.User == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetUser())
	return n
}

func (x *UpdateUserResp) Size() (n int) {
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
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	return n
}

func (x *SearchUserReq) sizeField1() (n int) {
	if x.Nickname == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetNickname())
	return n
}

func (x *SearchUserReq) sizeField2() (n int) {
	if x.Offset == nil {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetOffset())
	return n
}

func (x *SearchUserReq) sizeField3() (n int) {
	if x.Limit == nil {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetLimit())
	return n
}

func (x *SearchUserReq) sizeField4() (n int) {
	if x.Backward == nil {
		return n
	}
	n += fastpb.SizeBool(4, x.GetBackward())
	return n
}

func (x *SearchUserReq) sizeField5() (n int) {
	if x.LastToken == nil {
		return n
	}
	n += fastpb.SizeString(5, x.GetLastToken())
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

var fieldIDToName_UserDetail = map[int32]string{
	1: "Id",
	2: "AvatarUrl",
	3: "Nickname",
	4: "Motto",
}

var fieldIDToName_GetUserReq = map[int32]string{
	1: "UserId",
}

var fieldIDToName_GetUserResp = map[int32]string{
	1: "User",
}

var fieldIDToName_GetUserDetailReq = map[int32]string{
	1: "UserId",
}

var fieldIDToName_GetUserDetailResp = map[int32]string{
	1: "User",
}

var fieldIDToName_UpdateUserReq = map[int32]string{
	1: "User",
}

var fieldIDToName_UpdateUserResp = map[int32]string{}

var fieldIDToName_SearchUserReq = map[int32]string{
	1: "Nickname",
	2: "Offset",
	3: "Limit",
	4: "Backward",
	5: "LastToken",
}

var fieldIDToName_SearchUserResp = map[int32]string{
	1: "Users",
	2: "Total",
	3: "Token",
}

var _ = basic.File_meowchat_basic_basic_proto
