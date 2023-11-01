// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package sts

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
	basic "github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *GenCosStsReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GenCosStsReq[number], err)
}

func (x *GenCosStsReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Path, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GenCosStsResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GenCosStsResp[number], err)
}

func (x *GenCosStsResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.SecretId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GenCosStsResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.SecretKey, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GenCosStsResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.SessionToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GenCosStsResp) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.ExpiredTime, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GenCosStsResp) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.StartTime, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GenSignedUrlReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GenSignedUrlReq[number], err)
}

func (x *GenSignedUrlReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.SecretId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GenSignedUrlReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.SecretKey, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GenSignedUrlReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Method, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GenSignedUrlReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Path, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GenSignedUrlResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GenSignedUrlResp[number], err)
}

func (x *GenSignedUrlResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.SignedUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DeleteObjectReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeleteObjectReq[number], err)
}

func (x *DeleteObjectReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Path, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DeleteObjectResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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

func (x *TextCheckReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_TextCheckReq[number], err)
}

func (x *TextCheckReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Text, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *TextCheckReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v basic.UserMeta
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.User = &v
	return offset, nil
}

func (x *TextCheckReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.Scene = Scene(v)
	return offset, nil
}

func (x *TextCheckReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Title = &tmp
	return offset, err
}

func (x *TextCheckResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_TextCheckResp[number], err)
}

func (x *TextCheckResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Pass, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *PhotoCheckReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_PhotoCheckReq[number], err)
}

func (x *PhotoCheckReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v string
	v, offset, err = fastpb.ReadString(buf, _type)
	if err != nil {
		return offset, err
	}
	x.Url = append(x.Url, v)
	return offset, err
}

func (x *PhotoCheckReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v basic.UserMeta
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.User = &v
	return offset, nil
}

func (x *PhotoCheckResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_PhotoCheckResp[number], err)
}

func (x *PhotoCheckResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Pass, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

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
	x.OpenId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SignInResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.UnionId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SignInResp) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.AppId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SetPasswordReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SetPasswordReq[number], err)
}

func (x *SetPasswordReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SetPasswordReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
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

func (x *GenCosStsReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GenCosStsReq) fastWriteField1(buf []byte) (offset int) {
	if x.Path == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetPath())
	return offset
}

func (x *GenCosStsResp) FastWrite(buf []byte) (offset int) {
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

func (x *GenCosStsResp) fastWriteField1(buf []byte) (offset int) {
	if x.SecretId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetSecretId())
	return offset
}

func (x *GenCosStsResp) fastWriteField2(buf []byte) (offset int) {
	if x.SecretKey == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetSecretKey())
	return offset
}

func (x *GenCosStsResp) fastWriteField3(buf []byte) (offset int) {
	if x.SessionToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetSessionToken())
	return offset
}

func (x *GenCosStsResp) fastWriteField4(buf []byte) (offset int) {
	if x.ExpiredTime == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.GetExpiredTime())
	return offset
}

func (x *GenCosStsResp) fastWriteField5(buf []byte) (offset int) {
	if x.StartTime == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 5, x.GetStartTime())
	return offset
}

func (x *GenSignedUrlReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *GenSignedUrlReq) fastWriteField1(buf []byte) (offset int) {
	if x.SecretId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetSecretId())
	return offset
}

func (x *GenSignedUrlReq) fastWriteField2(buf []byte) (offset int) {
	if x.SecretKey == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetSecretKey())
	return offset
}

func (x *GenSignedUrlReq) fastWriteField3(buf []byte) (offset int) {
	if x.Method == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetMethod())
	return offset
}

func (x *GenSignedUrlReq) fastWriteField4(buf []byte) (offset int) {
	if x.Path == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetPath())
	return offset
}

func (x *GenSignedUrlResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GenSignedUrlResp) fastWriteField1(buf []byte) (offset int) {
	if x.SignedUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetSignedUrl())
	return offset
}

func (x *DeleteObjectReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DeleteObjectReq) fastWriteField1(buf []byte) (offset int) {
	if x.Path == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetPath())
	return offset
}

func (x *DeleteObjectResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *TextCheckReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *TextCheckReq) fastWriteField1(buf []byte) (offset int) {
	if x.Text == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetText())
	return offset
}

func (x *TextCheckReq) fastWriteField2(buf []byte) (offset int) {
	if x.User == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetUser())
	return offset
}

func (x *TextCheckReq) fastWriteField3(buf []byte) (offset int) {
	if x.Scene == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, int32(x.GetScene()))
	return offset
}

func (x *TextCheckReq) fastWriteField4(buf []byte) (offset int) {
	if x.Title == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetTitle())
	return offset
}

func (x *TextCheckResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *TextCheckResp) fastWriteField1(buf []byte) (offset int) {
	if !x.Pass {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetPass())
	return offset
}

func (x *PhotoCheckReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *PhotoCheckReq) fastWriteField1(buf []byte) (offset int) {
	if len(x.Url) == 0 {
		return offset
	}
	for i := range x.GetUrl() {
		offset += fastpb.WriteString(buf[offset:], 1, x.GetUrl()[i])
	}
	return offset
}

func (x *PhotoCheckReq) fastWriteField2(buf []byte) (offset int) {
	if x.User == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetUser())
	return offset
}

func (x *PhotoCheckResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *PhotoCheckResp) fastWriteField1(buf []byte) (offset int) {
	if !x.Pass {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetPass())
	return offset
}

func (x *SignInReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
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
	if x.OpenId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetOpenId())
	return offset
}

func (x *SignInResp) fastWriteField3(buf []byte) (offset int) {
	if x.UnionId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetUnionId())
	return offset
}

func (x *SignInResp) fastWriteField4(buf []byte) (offset int) {
	if x.AppId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetAppId())
	return offset
}

func (x *SetPasswordReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *SetPasswordReq) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *SetPasswordReq) fastWriteField2(buf []byte) (offset int) {
	if x.Password == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetPassword())
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

func (x *GenCosStsReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GenCosStsReq) sizeField1() (n int) {
	if x.Path == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetPath())
	return n
}

func (x *GenCosStsResp) Size() (n int) {
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

func (x *GenCosStsResp) sizeField1() (n int) {
	if x.SecretId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetSecretId())
	return n
}

func (x *GenCosStsResp) sizeField2() (n int) {
	if x.SecretKey == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetSecretKey())
	return n
}

func (x *GenCosStsResp) sizeField3() (n int) {
	if x.SessionToken == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetSessionToken())
	return n
}

func (x *GenCosStsResp) sizeField4() (n int) {
	if x.ExpiredTime == 0 {
		return n
	}
	n += fastpb.SizeInt64(4, x.GetExpiredTime())
	return n
}

func (x *GenCosStsResp) sizeField5() (n int) {
	if x.StartTime == 0 {
		return n
	}
	n += fastpb.SizeInt64(5, x.GetStartTime())
	return n
}

func (x *GenSignedUrlReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *GenSignedUrlReq) sizeField1() (n int) {
	if x.SecretId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetSecretId())
	return n
}

func (x *GenSignedUrlReq) sizeField2() (n int) {
	if x.SecretKey == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetSecretKey())
	return n
}

func (x *GenSignedUrlReq) sizeField3() (n int) {
	if x.Method == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetMethod())
	return n
}

func (x *GenSignedUrlReq) sizeField4() (n int) {
	if x.Path == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetPath())
	return n
}

func (x *GenSignedUrlResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GenSignedUrlResp) sizeField1() (n int) {
	if x.SignedUrl == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetSignedUrl())
	return n
}

func (x *DeleteObjectReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DeleteObjectReq) sizeField1() (n int) {
	if x.Path == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetPath())
	return n
}

func (x *DeleteObjectResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *TextCheckReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *TextCheckReq) sizeField1() (n int) {
	if x.Text == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetText())
	return n
}

func (x *TextCheckReq) sizeField2() (n int) {
	if x.User == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetUser())
	return n
}

func (x *TextCheckReq) sizeField3() (n int) {
	if x.Scene == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, int32(x.GetScene()))
	return n
}

func (x *TextCheckReq) sizeField4() (n int) {
	if x.Title == nil {
		return n
	}
	n += fastpb.SizeString(4, x.GetTitle())
	return n
}

func (x *TextCheckResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *TextCheckResp) sizeField1() (n int) {
	if !x.Pass {
		return n
	}
	n += fastpb.SizeBool(1, x.GetPass())
	return n
}

func (x *PhotoCheckReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *PhotoCheckReq) sizeField1() (n int) {
	if len(x.Url) == 0 {
		return n
	}
	for i := range x.GetUrl() {
		n += fastpb.SizeString(1, x.GetUrl()[i])
	}
	return n
}

func (x *PhotoCheckReq) sizeField2() (n int) {
	if x.User == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetUser())
	return n
}

func (x *PhotoCheckResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *PhotoCheckResp) sizeField1() (n int) {
	if !x.Pass {
		return n
	}
	n += fastpb.SizeBool(1, x.GetPass())
	return n
}

func (x *SignInReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
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
	if x.OpenId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetOpenId())
	return n
}

func (x *SignInResp) sizeField3() (n int) {
	if x.UnionId == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetUnionId())
	return n
}

func (x *SignInResp) sizeField4() (n int) {
	if x.AppId == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetAppId())
	return n
}

func (x *SetPasswordReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *SetPasswordReq) sizeField1() (n int) {
	if x.UserId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetUserId())
	return n
}

func (x *SetPasswordReq) sizeField2() (n int) {
	if x.Password == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetPassword())
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

var fieldIDToName_GenCosStsReq = map[int32]string{
	1: "Path",
}

var fieldIDToName_GenCosStsResp = map[int32]string{
	1: "SecretId",
	2: "SecretKey",
	3: "SessionToken",
	4: "ExpiredTime",
	5: "StartTime",
}

var fieldIDToName_GenSignedUrlReq = map[int32]string{
	1: "SecretId",
	2: "SecretKey",
	3: "Method",
	4: "Path",
}

var fieldIDToName_GenSignedUrlResp = map[int32]string{
	1: "SignedUrl",
}

var fieldIDToName_DeleteObjectReq = map[int32]string{
	1: "Path",
}

var fieldIDToName_DeleteObjectResp = map[int32]string{}

var fieldIDToName_TextCheckReq = map[int32]string{
	1: "Text",
	2: "User",
	3: "Scene",
	4: "Title",
}

var fieldIDToName_TextCheckResp = map[int32]string{
	1: "Pass",
}

var fieldIDToName_PhotoCheckReq = map[int32]string{
	1: "Url",
	2: "User",
}

var fieldIDToName_PhotoCheckResp = map[int32]string{
	1: "Pass",
}

var fieldIDToName_SignInReq = map[int32]string{
	1: "AuthType",
	2: "AuthId",
	3: "Password",
	4: "VerifyCode",
}

var fieldIDToName_SignInResp = map[int32]string{
	1: "UserId",
	2: "OpenId",
	3: "UnionId",
	4: "AppId",
}

var fieldIDToName_SetPasswordReq = map[int32]string{
	1: "UserId",
	2: "Password",
}

var fieldIDToName_SetPasswordResp = map[int32]string{}

var fieldIDToName_SendVerifyCodeReq = map[int32]string{
	1: "AuthType",
	2: "AuthId",
}

var fieldIDToName_SendVerifyCodeResp = map[int32]string{}

var _ = basic.File_basic_user_proto
