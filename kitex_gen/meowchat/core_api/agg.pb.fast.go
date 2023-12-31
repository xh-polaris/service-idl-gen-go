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

func (x *PrefetchReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_PrefetchReq[number], err)
}

func (x *PrefetchReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Appid, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PrefetchReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Token = &tmp
	return offset, err
}

func (x *PrefetchReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Code = &tmp
	return offset, err
}

func (x *PrefetchReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Timestamp, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *PrefetchReq) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Path = &tmp
	return offset, err
}

func (x *PrefetchReq) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Query = &tmp
	return offset, err
}

func (x *PrefetchReq) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Scene = &tmp
	return offset, err
}

func (x *PrefetchResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_PrefetchResp[number], err)
}

func (x *PrefetchResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v SignInResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.SignInResp = &v
	return offset, nil
}

func (x *PrefetchResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v GetUserInfoResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.GetUserInfoResp = &v
	return offset, nil
}

func (x *PrefetchResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v ListCommunityResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.ListCommunityResp = &v
	return offset, nil
}

func (x *PrefetchResp) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	var v GetPostPreviewsResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.FirstPostPreviewsResp = &v
	return offset, nil
}

func (x *PrefetchResp) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	var v GetMomentPreviewsResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.FirstMomentPreviewsResp = &v
	return offset, nil
}

func (x *PrefetchResp) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	var v GetCatPreviewsResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.FirstCatPreviewsResp = &v
	return offset, nil
}

func (x *PrefetchResp) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	var v GetNewsResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.GetNewsResp = &v
	return offset, nil
}

func (x *PrefetchResp) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Token = &tmp
	return offset, err
}

func (x *PrefetchResp) fastReadField9(buf []byte, _type int8) (offset int, err error) {
	x.Timestamp, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *PrefetchReq) FastWrite(buf []byte) (offset int) {
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

func (x *PrefetchReq) fastWriteField1(buf []byte) (offset int) {
	if x.Appid == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetAppid())
	return offset
}

func (x *PrefetchReq) fastWriteField2(buf []byte) (offset int) {
	if x.Token == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetToken())
	return offset
}

func (x *PrefetchReq) fastWriteField3(buf []byte) (offset int) {
	if x.Code == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetCode())
	return offset
}

func (x *PrefetchReq) fastWriteField4(buf []byte) (offset int) {
	if x.Timestamp == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.GetTimestamp())
	return offset
}

func (x *PrefetchReq) fastWriteField5(buf []byte) (offset int) {
	if x.Path == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetPath())
	return offset
}

func (x *PrefetchReq) fastWriteField6(buf []byte) (offset int) {
	if x.Query == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetQuery())
	return offset
}

func (x *PrefetchReq) fastWriteField7(buf []byte) (offset int) {
	if x.Scene == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 7, x.GetScene())
	return offset
}

func (x *PrefetchResp) FastWrite(buf []byte) (offset int) {
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
	return offset
}

func (x *PrefetchResp) fastWriteField1(buf []byte) (offset int) {
	if x.SignInResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetSignInResp())
	return offset
}

func (x *PrefetchResp) fastWriteField2(buf []byte) (offset int) {
	if x.GetUserInfoResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetGetUserInfoResp())
	return offset
}

func (x *PrefetchResp) fastWriteField3(buf []byte) (offset int) {
	if x.ListCommunityResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 3, x.GetListCommunityResp())
	return offset
}

func (x *PrefetchResp) fastWriteField4(buf []byte) (offset int) {
	if x.FirstPostPreviewsResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 4, x.GetFirstPostPreviewsResp())
	return offset
}

func (x *PrefetchResp) fastWriteField5(buf []byte) (offset int) {
	if x.FirstMomentPreviewsResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 5, x.GetFirstMomentPreviewsResp())
	return offset
}

func (x *PrefetchResp) fastWriteField6(buf []byte) (offset int) {
	if x.FirstCatPreviewsResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 6, x.GetFirstCatPreviewsResp())
	return offset
}

func (x *PrefetchResp) fastWriteField7(buf []byte) (offset int) {
	if x.GetNewsResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 7, x.GetGetNewsResp())
	return offset
}

func (x *PrefetchResp) fastWriteField8(buf []byte) (offset int) {
	if x.Token == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 8, x.GetToken())
	return offset
}

func (x *PrefetchResp) fastWriteField9(buf []byte) (offset int) {
	if x.Timestamp == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 9, x.GetTimestamp())
	return offset
}

func (x *PrefetchReq) Size() (n int) {
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

func (x *PrefetchReq) sizeField1() (n int) {
	if x.Appid == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetAppid())
	return n
}

func (x *PrefetchReq) sizeField2() (n int) {
	if x.Token == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetToken())
	return n
}

func (x *PrefetchReq) sizeField3() (n int) {
	if x.Code == nil {
		return n
	}
	n += fastpb.SizeString(3, x.GetCode())
	return n
}

func (x *PrefetchReq) sizeField4() (n int) {
	if x.Timestamp == 0 {
		return n
	}
	n += fastpb.SizeInt64(4, x.GetTimestamp())
	return n
}

func (x *PrefetchReq) sizeField5() (n int) {
	if x.Path == nil {
		return n
	}
	n += fastpb.SizeString(5, x.GetPath())
	return n
}

func (x *PrefetchReq) sizeField6() (n int) {
	if x.Query == nil {
		return n
	}
	n += fastpb.SizeString(6, x.GetQuery())
	return n
}

func (x *PrefetchReq) sizeField7() (n int) {
	if x.Scene == nil {
		return n
	}
	n += fastpb.SizeString(7, x.GetScene())
	return n
}

func (x *PrefetchResp) Size() (n int) {
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
	return n
}

func (x *PrefetchResp) sizeField1() (n int) {
	if x.SignInResp == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetSignInResp())
	return n
}

func (x *PrefetchResp) sizeField2() (n int) {
	if x.GetUserInfoResp == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetGetUserInfoResp())
	return n
}

func (x *PrefetchResp) sizeField3() (n int) {
	if x.ListCommunityResp == nil {
		return n
	}
	n += fastpb.SizeMessage(3, x.GetListCommunityResp())
	return n
}

func (x *PrefetchResp) sizeField4() (n int) {
	if x.FirstPostPreviewsResp == nil {
		return n
	}
	n += fastpb.SizeMessage(4, x.GetFirstPostPreviewsResp())
	return n
}

func (x *PrefetchResp) sizeField5() (n int) {
	if x.FirstMomentPreviewsResp == nil {
		return n
	}
	n += fastpb.SizeMessage(5, x.GetFirstMomentPreviewsResp())
	return n
}

func (x *PrefetchResp) sizeField6() (n int) {
	if x.FirstCatPreviewsResp == nil {
		return n
	}
	n += fastpb.SizeMessage(6, x.GetFirstCatPreviewsResp())
	return n
}

func (x *PrefetchResp) sizeField7() (n int) {
	if x.GetNewsResp == nil {
		return n
	}
	n += fastpb.SizeMessage(7, x.GetGetNewsResp())
	return n
}

func (x *PrefetchResp) sizeField8() (n int) {
	if x.Token == nil {
		return n
	}
	n += fastpb.SizeString(8, x.GetToken())
	return n
}

func (x *PrefetchResp) sizeField9() (n int) {
	if x.Timestamp == 0 {
		return n
	}
	n += fastpb.SizeInt64(9, x.GetTimestamp())
	return n
}

var fieldIDToName_PrefetchReq = map[int32]string{
	1: "Appid",
	2: "Token",
	3: "Code",
	4: "Timestamp",
	5: "Path",
	6: "Query",
	7: "Scene",
}

var fieldIDToName_PrefetchResp = map[int32]string{
	1: "SignInResp",
	2: "GetUserInfoResp",
	3: "ListCommunityResp",
	4: "FirstPostPreviewsResp",
	5: "FirstMomentPreviewsResp",
	6: "FirstCatPreviewsResp",
	7: "GetNewsResp",
	8: "Token",
	9: "Timestamp",
}
