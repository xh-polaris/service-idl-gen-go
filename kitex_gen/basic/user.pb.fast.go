// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package basic

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *UserMeta) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UserMeta[number], err)
}

func (x *UserMeta) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UserMeta) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.AppId = APP(v)
	return offset, nil
}

func (x *UserMeta) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.DeviceId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UserMeta) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.SessionUserId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UserMeta) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.SessionAppId = APP(v)
	return offset, nil
}

func (x *UserMeta) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.SessionDeviceId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UserMeta) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.IsLogin, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *UserMeta) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	var v WechatUserMeta
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.WechatUserMeta = &v
	return offset, nil
}

func (x *Extra) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Extra[number], err)
}

func (x *Extra) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ClientIP, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Extra) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Language, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Extra) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Resolution, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Extra) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.DevicePlatform, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Extra) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.DeviceBrand, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Extra) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.DeviceId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Extra) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.OperateSystem, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *WechatUserMeta) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_WechatUserMeta[number], err)
}

func (x *WechatUserMeta) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.AppId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *WechatUserMeta) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.OpenId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *WechatUserMeta) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.PlatformId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *WechatUserMeta) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.UnionId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UserMeta) FastWrite(buf []byte) (offset int) {
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
	return offset
}

func (x *UserMeta) fastWriteField1(buf []byte) (offset int) {
	if x.UserId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetUserId())
	return offset
}

func (x *UserMeta) fastWriteField2(buf []byte) (offset int) {
	if x.AppId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, int32(x.GetAppId()))
	return offset
}

func (x *UserMeta) fastWriteField3(buf []byte) (offset int) {
	if x.DeviceId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetDeviceId())
	return offset
}

func (x *UserMeta) fastWriteField4(buf []byte) (offset int) {
	if x.SessionUserId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetSessionUserId())
	return offset
}

func (x *UserMeta) fastWriteField5(buf []byte) (offset int) {
	if x.SessionAppId == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 5, int32(x.GetSessionAppId()))
	return offset
}

func (x *UserMeta) fastWriteField6(buf []byte) (offset int) {
	if x.SessionDeviceId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetSessionDeviceId())
	return offset
}

func (x *UserMeta) fastWriteField7(buf []byte) (offset int) {
	if !x.IsLogin {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 7, x.GetIsLogin())
	return offset
}

func (x *UserMeta) fastWriteField8(buf []byte) (offset int) {
	if x.WechatUserMeta == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 8, x.GetWechatUserMeta())
	return offset
}

func (x *Extra) FastWrite(buf []byte) (offset int) {
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
	return offset
}

func (x *Extra) fastWriteField1(buf []byte) (offset int) {
	if x.ClientIP == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetClientIP())
	return offset
}

func (x *Extra) fastWriteField2(buf []byte) (offset int) {
	if x.Language == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetLanguage())
	return offset
}

func (x *Extra) fastWriteField3(buf []byte) (offset int) {
	if x.Resolution == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetResolution())
	return offset
}

func (x *Extra) fastWriteField4(buf []byte) (offset int) {
	if x.DevicePlatform == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetDevicePlatform())
	return offset
}

func (x *Extra) fastWriteField5(buf []byte) (offset int) {
	if x.DeviceBrand == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetDeviceBrand())
	return offset
}

func (x *Extra) fastWriteField6(buf []byte) (offset int) {
	if x.DeviceId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetDeviceId())
	return offset
}

func (x *Extra) fastWriteField7(buf []byte) (offset int) {
	if x.OperateSystem == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 7, x.GetOperateSystem())
	return offset
}

func (x *WechatUserMeta) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *WechatUserMeta) fastWriteField1(buf []byte) (offset int) {
	if x.AppId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetAppId())
	return offset
}

func (x *WechatUserMeta) fastWriteField2(buf []byte) (offset int) {
	if x.OpenId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetOpenId())
	return offset
}

func (x *WechatUserMeta) fastWriteField3(buf []byte) (offset int) {
	if x.PlatformId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetPlatformId())
	return offset
}

func (x *WechatUserMeta) fastWriteField4(buf []byte) (offset int) {
	if x.UnionId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetUnionId())
	return offset
}

func (x *UserMeta) Size() (n int) {
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
	return n
}

func (x *UserMeta) sizeField1() (n int) {
	if x.UserId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetUserId())
	return n
}

func (x *UserMeta) sizeField2() (n int) {
	if x.AppId == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, int32(x.GetAppId()))
	return n
}

func (x *UserMeta) sizeField3() (n int) {
	if x.DeviceId == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetDeviceId())
	return n
}

func (x *UserMeta) sizeField4() (n int) {
	if x.SessionUserId == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetSessionUserId())
	return n
}

func (x *UserMeta) sizeField5() (n int) {
	if x.SessionAppId == 0 {
		return n
	}
	n += fastpb.SizeInt32(5, int32(x.GetSessionAppId()))
	return n
}

func (x *UserMeta) sizeField6() (n int) {
	if x.SessionDeviceId == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetSessionDeviceId())
	return n
}

func (x *UserMeta) sizeField7() (n int) {
	if !x.IsLogin {
		return n
	}
	n += fastpb.SizeBool(7, x.GetIsLogin())
	return n
}

func (x *UserMeta) sizeField8() (n int) {
	if x.WechatUserMeta == nil {
		return n
	}
	n += fastpb.SizeMessage(8, x.GetWechatUserMeta())
	return n
}

func (x *Extra) Size() (n int) {
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
	return n
}

func (x *Extra) sizeField1() (n int) {
	if x.ClientIP == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetClientIP())
	return n
}

func (x *Extra) sizeField2() (n int) {
	if x.Language == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetLanguage())
	return n
}

func (x *Extra) sizeField3() (n int) {
	if x.Resolution == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetResolution())
	return n
}

func (x *Extra) sizeField4() (n int) {
	if x.DevicePlatform == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetDevicePlatform())
	return n
}

func (x *Extra) sizeField5() (n int) {
	if x.DeviceBrand == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetDeviceBrand())
	return n
}

func (x *Extra) sizeField6() (n int) {
	if x.DeviceId == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetDeviceId())
	return n
}

func (x *Extra) sizeField7() (n int) {
	if x.OperateSystem == "" {
		return n
	}
	n += fastpb.SizeString(7, x.GetOperateSystem())
	return n
}

func (x *WechatUserMeta) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *WechatUserMeta) sizeField1() (n int) {
	if x.AppId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetAppId())
	return n
}

func (x *WechatUserMeta) sizeField2() (n int) {
	if x.OpenId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetOpenId())
	return n
}

func (x *WechatUserMeta) sizeField3() (n int) {
	if x.PlatformId == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetPlatformId())
	return n
}

func (x *WechatUserMeta) sizeField4() (n int) {
	if x.UnionId == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetUnionId())
	return n
}

var fieldIDToName_UserMeta = map[int32]string{
	1: "UserId",
	2: "AppId",
	3: "DeviceId",
	4: "SessionUserId",
	5: "SessionAppId",
	6: "SessionDeviceId",
	7: "IsLogin",
	8: "WechatUserMeta",
}

var fieldIDToName_Extra = map[int32]string{
	1: "ClientIP",
	2: "Language",
	3: "Resolution",
	4: "DevicePlatform",
	5: "DeviceBrand",
	6: "DeviceId",
	7: "OperateSystem",
}

var fieldIDToName_WechatUserMeta = map[int32]string{
	1: "AppId",
	2: "OpenId",
	3: "PlatformId",
	4: "UnionId",
}
