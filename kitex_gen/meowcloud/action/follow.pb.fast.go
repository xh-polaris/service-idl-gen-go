// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package action

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
	basic "github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *DoFollowReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 254:
		offset, err = x.fastReadField254(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DoFollowReq[number], err)
}

func (x *DoFollowReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.TargetId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DoFollowReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.TargetType = TargetType(v)
	return offset, nil
}

func (x *DoFollowReq) fastReadField254(buf []byte, _type int8) (offset int, err error) {
	var v basic.UserMeta
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.User = &v
	return offset, nil
}

func (x *DoFollowResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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

func (x *CancelFollowReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 254:
		offset, err = x.fastReadField254(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CancelFollowReq[number], err)
}

func (x *CancelFollowReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.TargetId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CancelFollowReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.TargetType = TargetType(v)
	return offset, nil
}

func (x *CancelFollowReq) fastReadField254(buf []byte, _type int8) (offset int, err error) {
	var v basic.UserMeta
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.User = &v
	return offset, nil
}

func (x *CancelFollowResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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

func (x *GetFollowedCountReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetFollowedCountReq[number], err)
}

func (x *GetFollowedCountReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.TargetId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetFollowedCountReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.TargetType = TargetType(v)
	return offset, nil
}

func (x *GetFollowedCountResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetFollowedCountResp[number], err)
}

func (x *GetFollowedCountResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Count, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetFollowedUsersReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetFollowedUsersReq[number], err)
}

func (x *GetFollowedUsersReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.TargetId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetFollowedUsersReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.TargetType = TargetType(v)
	return offset, nil
}

func (x *GetFollowedUsersReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v basic.PaginationOptions
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.PaginationOption = &v
	return offset, nil
}

func (x *GetFollowedUsersResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetFollowedUsersResp[number], err)
}

func (x *GetFollowedUsersResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Action_Follow
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Follows = append(x.Follows, &v)
	return offset, nil
}

func (x *GetFollowedUsersResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Total, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetUserFollowedReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 254:
		offset, err = x.fastReadField254(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetUserFollowedReq[number], err)
}

func (x *GetUserFollowedReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.TargetType = TargetType(v)
	return offset, nil
}

func (x *GetUserFollowedReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v basic.PaginationOptions
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.PaginationOption = &v
	return offset, nil
}

func (x *GetUserFollowedReq) fastReadField254(buf []byte, _type int8) (offset int, err error) {
	var v basic.UserMeta
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.User = &v
	return offset, nil
}

func (x *GetUserFollowedResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetUserFollowedResp[number], err)
}

func (x *GetUserFollowedResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Action_Follow
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Follows = append(x.Follows, &v)
	return offset, nil
}

func (x *GetUserFollowedResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Total, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetFollowedReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 254:
		offset, err = x.fastReadField254(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetFollowedReq[number], err)
}

func (x *GetFollowedReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.TargetId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetFollowedReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.TargetType = TargetType(v)
	return offset, nil
}

func (x *GetFollowedReq) fastReadField254(buf []byte, _type int8) (offset int, err error) {
	var v basic.UserMeta
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.User = &v
	return offset, nil
}

func (x *GetFollowedResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetFollowedResp[number], err)
}

func (x *GetFollowedResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Followed, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *DoFollowReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField254(buf[offset:])
	return offset
}

func (x *DoFollowReq) fastWriteField1(buf []byte) (offset int) {
	if x.TargetId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetTargetId())
	return offset
}

func (x *DoFollowReq) fastWriteField2(buf []byte) (offset int) {
	if x.TargetType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, int32(x.GetTargetType()))
	return offset
}

func (x *DoFollowReq) fastWriteField254(buf []byte) (offset int) {
	if x.User == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 254, x.GetUser())
	return offset
}

func (x *DoFollowResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *CancelFollowReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField254(buf[offset:])
	return offset
}

func (x *CancelFollowReq) fastWriteField1(buf []byte) (offset int) {
	if x.TargetId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetTargetId())
	return offset
}

func (x *CancelFollowReq) fastWriteField2(buf []byte) (offset int) {
	if x.TargetType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, int32(x.GetTargetType()))
	return offset
}

func (x *CancelFollowReq) fastWriteField254(buf []byte) (offset int) {
	if x.User == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 254, x.GetUser())
	return offset
}

func (x *CancelFollowResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *GetFollowedCountReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetFollowedCountReq) fastWriteField1(buf []byte) (offset int) {
	if x.TargetId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetTargetId())
	return offset
}

func (x *GetFollowedCountReq) fastWriteField2(buf []byte) (offset int) {
	if x.TargetType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, int32(x.GetTargetType()))
	return offset
}

func (x *GetFollowedCountResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetFollowedCountResp) fastWriteField1(buf []byte) (offset int) {
	if x.Count == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetCount())
	return offset
}

func (x *GetFollowedUsersReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *GetFollowedUsersReq) fastWriteField1(buf []byte) (offset int) {
	if x.TargetId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetTargetId())
	return offset
}

func (x *GetFollowedUsersReq) fastWriteField2(buf []byte) (offset int) {
	if x.TargetType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, int32(x.GetTargetType()))
	return offset
}

func (x *GetFollowedUsersReq) fastWriteField3(buf []byte) (offset int) {
	if x.PaginationOption == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 3, x.GetPaginationOption())
	return offset
}

func (x *GetFollowedUsersResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetFollowedUsersResp) fastWriteField1(buf []byte) (offset int) {
	if x.Follows == nil {
		return offset
	}
	for i := range x.GetFollows() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetFollows()[i])
	}
	return offset
}

func (x *GetFollowedUsersResp) fastWriteField2(buf []byte) (offset int) {
	if x.Total == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetTotal())
	return offset
}

func (x *GetUserFollowedReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField254(buf[offset:])
	return offset
}

func (x *GetUserFollowedReq) fastWriteField1(buf []byte) (offset int) {
	if x.TargetType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, int32(x.GetTargetType()))
	return offset
}

func (x *GetUserFollowedReq) fastWriteField2(buf []byte) (offset int) {
	if x.PaginationOption == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetPaginationOption())
	return offset
}

func (x *GetUserFollowedReq) fastWriteField254(buf []byte) (offset int) {
	if x.User == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 254, x.GetUser())
	return offset
}

func (x *GetUserFollowedResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetUserFollowedResp) fastWriteField1(buf []byte) (offset int) {
	if x.Follows == nil {
		return offset
	}
	for i := range x.GetFollows() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetFollows()[i])
	}
	return offset
}

func (x *GetUserFollowedResp) fastWriteField2(buf []byte) (offset int) {
	if x.Total == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetTotal())
	return offset
}

func (x *GetFollowedReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField254(buf[offset:])
	return offset
}

func (x *GetFollowedReq) fastWriteField1(buf []byte) (offset int) {
	if x.TargetId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetTargetId())
	return offset
}

func (x *GetFollowedReq) fastWriteField2(buf []byte) (offset int) {
	if x.TargetType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, int32(x.GetTargetType()))
	return offset
}

func (x *GetFollowedReq) fastWriteField254(buf []byte) (offset int) {
	if x.User == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 254, x.GetUser())
	return offset
}

func (x *GetFollowedResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetFollowedResp) fastWriteField1(buf []byte) (offset int) {
	if !x.Followed {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetFollowed())
	return offset
}

func (x *DoFollowReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField254()
	return n
}

func (x *DoFollowReq) sizeField1() (n int) {
	if x.TargetId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetTargetId())
	return n
}

func (x *DoFollowReq) sizeField2() (n int) {
	if x.TargetType == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, int32(x.GetTargetType()))
	return n
}

func (x *DoFollowReq) sizeField254() (n int) {
	if x.User == nil {
		return n
	}
	n += fastpb.SizeMessage(254, x.GetUser())
	return n
}

func (x *DoFollowResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *CancelFollowReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField254()
	return n
}

func (x *CancelFollowReq) sizeField1() (n int) {
	if x.TargetId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetTargetId())
	return n
}

func (x *CancelFollowReq) sizeField2() (n int) {
	if x.TargetType == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, int32(x.GetTargetType()))
	return n
}

func (x *CancelFollowReq) sizeField254() (n int) {
	if x.User == nil {
		return n
	}
	n += fastpb.SizeMessage(254, x.GetUser())
	return n
}

func (x *CancelFollowResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *GetFollowedCountReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetFollowedCountReq) sizeField1() (n int) {
	if x.TargetId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetTargetId())
	return n
}

func (x *GetFollowedCountReq) sizeField2() (n int) {
	if x.TargetType == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, int32(x.GetTargetType()))
	return n
}

func (x *GetFollowedCountResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetFollowedCountResp) sizeField1() (n int) {
	if x.Count == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetCount())
	return n
}

func (x *GetFollowedUsersReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *GetFollowedUsersReq) sizeField1() (n int) {
	if x.TargetId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetTargetId())
	return n
}

func (x *GetFollowedUsersReq) sizeField2() (n int) {
	if x.TargetType == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, int32(x.GetTargetType()))
	return n
}

func (x *GetFollowedUsersReq) sizeField3() (n int) {
	if x.PaginationOption == nil {
		return n
	}
	n += fastpb.SizeMessage(3, x.GetPaginationOption())
	return n
}

func (x *GetFollowedUsersResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetFollowedUsersResp) sizeField1() (n int) {
	if x.Follows == nil {
		return n
	}
	for i := range x.GetFollows() {
		n += fastpb.SizeMessage(1, x.GetFollows()[i])
	}
	return n
}

func (x *GetFollowedUsersResp) sizeField2() (n int) {
	if x.Total == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetTotal())
	return n
}

func (x *GetUserFollowedReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField254()
	return n
}

func (x *GetUserFollowedReq) sizeField1() (n int) {
	if x.TargetType == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, int32(x.GetTargetType()))
	return n
}

func (x *GetUserFollowedReq) sizeField2() (n int) {
	if x.PaginationOption == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetPaginationOption())
	return n
}

func (x *GetUserFollowedReq) sizeField254() (n int) {
	if x.User == nil {
		return n
	}
	n += fastpb.SizeMessage(254, x.GetUser())
	return n
}

func (x *GetUserFollowedResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetUserFollowedResp) sizeField1() (n int) {
	if x.Follows == nil {
		return n
	}
	for i := range x.GetFollows() {
		n += fastpb.SizeMessage(1, x.GetFollows()[i])
	}
	return n
}

func (x *GetUserFollowedResp) sizeField2() (n int) {
	if x.Total == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetTotal())
	return n
}

func (x *GetFollowedReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField254()
	return n
}

func (x *GetFollowedReq) sizeField1() (n int) {
	if x.TargetId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetTargetId())
	return n
}

func (x *GetFollowedReq) sizeField2() (n int) {
	if x.TargetType == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, int32(x.GetTargetType()))
	return n
}

func (x *GetFollowedReq) sizeField254() (n int) {
	if x.User == nil {
		return n
	}
	n += fastpb.SizeMessage(254, x.GetUser())
	return n
}

func (x *GetFollowedResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetFollowedResp) sizeField1() (n int) {
	if !x.Followed {
		return n
	}
	n += fastpb.SizeBool(1, x.GetFollowed())
	return n
}

var fieldIDToName_DoFollowReq = map[int32]string{
	1:   "TargetId",
	2:   "TargetType",
	254: "User",
}

var fieldIDToName_DoFollowResp = map[int32]string{}

var fieldIDToName_CancelFollowReq = map[int32]string{
	1:   "TargetId",
	2:   "TargetType",
	254: "User",
}

var fieldIDToName_CancelFollowResp = map[int32]string{}

var fieldIDToName_GetFollowedCountReq = map[int32]string{
	1: "TargetId",
	2: "TargetType",
}

var fieldIDToName_GetFollowedCountResp = map[int32]string{
	1: "Count",
}

var fieldIDToName_GetFollowedUsersReq = map[int32]string{
	1: "TargetId",
	2: "TargetType",
	3: "PaginationOption",
}

var fieldIDToName_GetFollowedUsersResp = map[int32]string{
	1: "Follows",
	2: "Total",
}

var fieldIDToName_GetUserFollowedReq = map[int32]string{
	1:   "TargetType",
	2:   "PaginationOption",
	254: "User",
}

var fieldIDToName_GetUserFollowedResp = map[int32]string{
	1: "Follows",
	2: "Total",
}

var fieldIDToName_GetFollowedReq = map[int32]string{
	1:   "TargetId",
	2:   "TargetType",
	254: "User",
}

var fieldIDToName_GetFollowedResp = map[int32]string{
	1: "Followed",
}

var _ = basic.File_basic_pagination_proto
var _ = basic.File_basic_user_proto
