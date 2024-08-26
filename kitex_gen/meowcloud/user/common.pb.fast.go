// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package user

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
	basic "github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *MeowUser) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_MeowUser[number], err)
}

func (x *MeowUser) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v basic.UserMeta
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.UserMeta = &v
	return offset, nil
}

func (x *MeowUser) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Signature, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *MeowUser) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Avatar, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *MeowUser) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Username, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *MeowUser) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.CreatedTime, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *MeowUser) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.UpdatedTime, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *MeowUser) FastWrite(buf []byte) (offset int) {
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

func (x *MeowUser) fastWriteField1(buf []byte) (offset int) {
	if x.UserMeta == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetUserMeta())
	return offset
}

func (x *MeowUser) fastWriteField2(buf []byte) (offset int) {
	if x.Signature == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetSignature())
	return offset
}

func (x *MeowUser) fastWriteField3(buf []byte) (offset int) {
	if x.Avatar == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetAvatar())
	return offset
}

func (x *MeowUser) fastWriteField4(buf []byte) (offset int) {
	if x.Username == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetUsername())
	return offset
}

func (x *MeowUser) fastWriteField5(buf []byte) (offset int) {
	if x.CreatedTime == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 5, x.GetCreatedTime())
	return offset
}

func (x *MeowUser) fastWriteField6(buf []byte) (offset int) {
	if x.UpdatedTime == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 6, x.GetUpdatedTime())
	return offset
}

func (x *MeowUser) Size() (n int) {
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

func (x *MeowUser) sizeField1() (n int) {
	if x.UserMeta == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetUserMeta())
	return n
}

func (x *MeowUser) sizeField2() (n int) {
	if x.Signature == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetSignature())
	return n
}

func (x *MeowUser) sizeField3() (n int) {
	if x.Avatar == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetAvatar())
	return n
}

func (x *MeowUser) sizeField4() (n int) {
	if x.Username == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetUsername())
	return n
}

func (x *MeowUser) sizeField5() (n int) {
	if x.CreatedTime == 0 {
		return n
	}
	n += fastpb.SizeInt64(5, x.GetCreatedTime())
	return n
}

func (x *MeowUser) sizeField6() (n int) {
	if x.UpdatedTime == 0 {
		return n
	}
	n += fastpb.SizeInt64(6, x.GetUpdatedTime())
	return n
}

var fieldIDToName_MeowUser = map[int32]string{
	1: "UserMeta",
	2: "Signature",
	3: "Avatar",
	4: "Username",
	5: "CreatedTime",
	6: "UpdatedTime",
}

var _ = basic.File_basic_user_proto
