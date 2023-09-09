// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package core_api

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *ApplySignedUrlReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ApplySignedUrlReq[number], err)
}

func (x *ApplySignedUrlReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Prefix = &tmp
	return offset, err
}

func (x *ApplySignedUrlReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Suffix = &tmp
	return offset, err
}

func (x *ApplySignedUrlResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ApplySignedUrlResp[number], err)
}

func (x *ApplySignedUrlResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Url, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ApplySignedUrlResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.SessionToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ApplySignedUrlAsCommunityReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ApplySignedUrlAsCommunityReq[number], err)
}

func (x *ApplySignedUrlAsCommunityReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.CommunityId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ApplySignedUrlAsCommunityReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Prefix, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ApplySignedUrlAsCommunityReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Suffix, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ApplySignedUrlAsCommunityResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ApplySignedUrlAsCommunityResp[number], err)
}

func (x *ApplySignedUrlAsCommunityResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Url, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ApplySignedUrlAsCommunityResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.SessionToken, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ApplySignedUrlReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *ApplySignedUrlReq) fastWriteField1(buf []byte) (offset int) {
	if x.Prefix == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetPrefix())
	return offset
}

func (x *ApplySignedUrlReq) fastWriteField2(buf []byte) (offset int) {
	if x.Suffix == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetSuffix())
	return offset
}

func (x *ApplySignedUrlResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *ApplySignedUrlResp) fastWriteField1(buf []byte) (offset int) {
	if x.Url == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetUrl())
	return offset
}

func (x *ApplySignedUrlResp) fastWriteField2(buf []byte) (offset int) {
	if x.SessionToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetSessionToken())
	return offset
}

func (x *ApplySignedUrlAsCommunityReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *ApplySignedUrlAsCommunityReq) fastWriteField1(buf []byte) (offset int) {
	if x.CommunityId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetCommunityId())
	return offset
}

func (x *ApplySignedUrlAsCommunityReq) fastWriteField2(buf []byte) (offset int) {
	if x.Prefix == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetPrefix())
	return offset
}

func (x *ApplySignedUrlAsCommunityReq) fastWriteField3(buf []byte) (offset int) {
	if x.Suffix == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetSuffix())
	return offset
}

func (x *ApplySignedUrlAsCommunityResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *ApplySignedUrlAsCommunityResp) fastWriteField1(buf []byte) (offset int) {
	if x.Url == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetUrl())
	return offset
}

func (x *ApplySignedUrlAsCommunityResp) fastWriteField2(buf []byte) (offset int) {
	if x.SessionToken == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetSessionToken())
	return offset
}

func (x *ApplySignedUrlReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *ApplySignedUrlReq) sizeField1() (n int) {
	if x.Prefix == nil {
		return n
	}
	n += fastpb.SizeString(1, x.GetPrefix())
	return n
}

func (x *ApplySignedUrlReq) sizeField2() (n int) {
	if x.Suffix == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetSuffix())
	return n
}

func (x *ApplySignedUrlResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *ApplySignedUrlResp) sizeField1() (n int) {
	if x.Url == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetUrl())
	return n
}

func (x *ApplySignedUrlResp) sizeField2() (n int) {
	if x.SessionToken == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetSessionToken())
	return n
}

func (x *ApplySignedUrlAsCommunityReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *ApplySignedUrlAsCommunityReq) sizeField1() (n int) {
	if x.CommunityId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetCommunityId())
	return n
}

func (x *ApplySignedUrlAsCommunityReq) sizeField2() (n int) {
	if x.Prefix == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetPrefix())
	return n
}

func (x *ApplySignedUrlAsCommunityReq) sizeField3() (n int) {
	if x.Suffix == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetSuffix())
	return n
}

func (x *ApplySignedUrlAsCommunityResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *ApplySignedUrlAsCommunityResp) sizeField1() (n int) {
	if x.Url == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetUrl())
	return n
}

func (x *ApplySignedUrlAsCommunityResp) sizeField2() (n int) {
	if x.SessionToken == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetSessionToken())
	return n
}

var fieldIDToName_ApplySignedUrlReq = map[int32]string{
	1: "Prefix",
	2: "Suffix",
}

var fieldIDToName_ApplySignedUrlResp = map[int32]string{
	1: "Url",
	2: "SessionToken",
}

var fieldIDToName_ApplySignedUrlAsCommunityReq = map[int32]string{
	1: "CommunityId",
	2: "Prefix",
	3: "Suffix",
}

var fieldIDToName_ApplySignedUrlAsCommunityResp = map[int32]string{
	1: "Url",
	2: "SessionToken",
}
