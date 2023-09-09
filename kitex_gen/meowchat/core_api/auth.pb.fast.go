// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package core_api

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
	basic "github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *SignInReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SignInReq[number], err)
}

func (x *SignInReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.AuthType, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SignInReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.AuthId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SignInReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Password = &tmp
	return offset, err
}

func (x *SignInReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.VerifyCode = &tmp
	return offset, err
}

func (x *SignInReq) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.AppId = basic.APP(v)
	return offset, nil
}

func (x *SignInReq) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.DeviceId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SignInResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SignInResp[number], err)
}

func (x *SignInResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SignInResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.AccessToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SignInResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.AccessExpire, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *SignInResp) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.IsFirst, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *SetPasswordReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SetPasswordReq[number], err)
}

func (x *SetPasswordReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Password, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SetPasswordResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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

func (x *SendVerifyCodeReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SendVerifyCodeReq[number], err)
}

func (x *SendVerifyCodeReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.AuthType, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SendVerifyCodeReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.AuthId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SendVerifyCodeResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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

func (x *SignInReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	return offset
}

func (x *SignInReq) fastWriteField1(buf []byte) (offset int) {
	if x.AuthType == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetAuthType())
	return offset
}

func (x *SignInReq) fastWriteField2(buf []byte) (offset int) {
	if x.AuthId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetAuthId())
	return offset
}

func (x *SignInReq) fastWriteField3(buf []byte) (offset int) {
	if x.Password == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetPassword())
	return offset
}

func (x *SignInReq) fastWriteField4(buf []byte) (offset int) {
	if x.VerifyCode == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetVerifyCode())
	return offset
}

func (x *SignInReq) fastWriteField5(buf []byte) (offset int) {
	if x.AppId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 5, int32(x.GetAppId()))
	return offset
}

func (x *SignInReq) fastWriteField6(buf []byte) (offset int) {
	if x.DeviceId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetDeviceId())
	return offset
}

func (x *SignInResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *SignInResp) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *SignInResp) fastWriteField2(buf []byte) (offset int) {
	if x.AccessToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetAccessToken())
	return offset
}

func (x *SignInResp) fastWriteField3(buf []byte) (offset int) {
	if x.AccessExpire == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetAccessExpire())
	return offset
}

func (x *SignInResp) fastWriteField4(buf []byte) (offset int) {
	if !x.IsFirst {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 4, x.GetIsFirst())
	return offset
}

func (x *SetPasswordReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *SetPasswordReq) fastWriteField1(buf []byte) (offset int) {
	if x.Password == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetPassword())
	return offset
}

func (x *SetPasswordResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *SendVerifyCodeReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *SendVerifyCodeReq) fastWriteField1(buf []byte) (offset int) {
	if x.AuthType == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetAuthType())
	return offset
}

func (x *SendVerifyCodeReq) fastWriteField2(buf []byte) (offset int) {
	if x.AuthId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetAuthId())
	return offset
}

func (x *SendVerifyCodeResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *SignInReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	return n
}

func (x *SignInReq) sizeField1() (n int) {
	if x.AuthType == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetAuthType())
	return n
}

func (x *SignInReq) sizeField2() (n int) {
	if x.AuthId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetAuthId())
	return n
}

func (x *SignInReq) sizeField3() (n int) {
	if x.Password == nil {
		return n
	}
	n += fastpb.SizeString(3, x.GetPassword())
	return n
}

func (x *SignInReq) sizeField4() (n int) {
	if x.VerifyCode == nil {
		return n
	}
	n += fastpb.SizeString(4, x.GetVerifyCode())
	return n
}

func (x *SignInReq) sizeField5() (n int) {
	if x.AppId == 0 {
		return n
	}
	n += fastpb.SizeInt32(5, int32(x.GetAppId()))
	return n
}

func (x *SignInReq) sizeField6() (n int) {
	if x.DeviceId == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetDeviceId())
	return n
}

func (x *SignInResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *SignInResp) sizeField1() (n int) {
	if x.UserId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetUserId())
	return n
}

func (x *SignInResp) sizeField2() (n int) {
	if x.AccessToken == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetAccessToken())
	return n
}

func (x *SignInResp) sizeField3() (n int) {
	if x.AccessExpire == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetAccessExpire())
	return n
}

func (x *SignInResp) sizeField4() (n int) {
	if !x.IsFirst {
		return n
	}
	n += fastpb.SizeBool(4, x.GetIsFirst())
	return n
}

func (x *SetPasswordReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *SetPasswordReq) sizeField1() (n int) {
	if x.Password == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetPassword())
	return n
}

func (x *SetPasswordResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *SendVerifyCodeReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *SendVerifyCodeReq) sizeField1() (n int) {
	if x.AuthType == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetAuthType())
	return n
}

func (x *SendVerifyCodeReq) sizeField2() (n int) {
	if x.AuthId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetAuthId())
	return n
}

func (x *SendVerifyCodeResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

var fieldIDToName_SignInReq = map[int32]string{
	1: "AuthType",
	2: "AuthId",
	3: "Password",
	4: "VerifyCode",
	5: "AppId",
	6: "DeviceId",
}

var fieldIDToName_SignInResp = map[int32]string{
	1: "UserId",
	2: "AccessToken",
	3: "AccessExpire",
	4: "IsFirst",
}

var fieldIDToName_SetPasswordReq = map[int32]string{
	1: "Password",
}

var fieldIDToName_SetPasswordResp = map[int32]string{}

var fieldIDToName_SendVerifyCodeReq = map[int32]string{
	1: "AuthType",
	2: "AuthId",
}

var fieldIDToName_SendVerifyCodeResp = map[int32]string{}

var _ = basic.File_basic_app_proto
