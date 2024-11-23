// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package action

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *Action) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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

func (x *Action_Like) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Action_Like[number], err)
}

func (x *Action_Like) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Action_Like) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.TargetId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Action_Like) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.TargetType = TargetType(v)
	return offset, nil
}

func (x *Action_Like) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Action_Like) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.CreateAt, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Action_Like) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.IsCancel, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *Action_Share) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Action_Share[number], err)
}

func (x *Action_Share) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Action_Share) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.TargetId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Action_Share) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.TargetType = TargetType(v)
	return offset, nil
}

func (x *Action_Share) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Action_Share) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.CreateAt, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Action_Follow) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Action_Follow[number], err)
}

func (x *Action_Follow) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Action_Follow) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.TargetId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Action_Follow) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.TargetType = TargetType(v)
	return offset, nil
}

func (x *Action_Follow) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Action_Follow) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.CreateAt, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Action_Follow) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.IsCancel, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *Action) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *Action_Like) FastWrite(buf []byte) (offset int) {
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

func (x *Action_Like) fastWriteField1(buf []byte) (offset int) {
	if x.Id == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetId())
	return offset
}

func (x *Action_Like) fastWriteField2(buf []byte) (offset int) {
	if x.TargetId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetTargetId())
	return offset
}

func (x *Action_Like) fastWriteField3(buf []byte) (offset int) {
	if x.TargetType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, int32(x.GetTargetType()))
	return offset
}

func (x *Action_Like) fastWriteField4(buf []byte) (offset int) {
	if x.UserId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetUserId())
	return offset
}

func (x *Action_Like) fastWriteField5(buf []byte) (offset int) {
	if x.CreateAt == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 5, x.GetCreateAt())
	return offset
}

func (x *Action_Like) fastWriteField6(buf []byte) (offset int) {
	if !x.IsCancel {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 6, x.GetIsCancel())
	return offset
}

func (x *Action_Share) FastWrite(buf []byte) (offset int) {
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

func (x *Action_Share) fastWriteField1(buf []byte) (offset int) {
	if x.Id == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetId())
	return offset
}

func (x *Action_Share) fastWriteField2(buf []byte) (offset int) {
	if x.TargetId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetTargetId())
	return offset
}

func (x *Action_Share) fastWriteField3(buf []byte) (offset int) {
	if x.TargetType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, int32(x.GetTargetType()))
	return offset
}

func (x *Action_Share) fastWriteField4(buf []byte) (offset int) {
	if x.UserId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetUserId())
	return offset
}

func (x *Action_Share) fastWriteField5(buf []byte) (offset int) {
	if x.CreateAt == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 5, x.GetCreateAt())
	return offset
}

func (x *Action_Follow) FastWrite(buf []byte) (offset int) {
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

func (x *Action_Follow) fastWriteField1(buf []byte) (offset int) {
	if x.Id == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetId())
	return offset
}

func (x *Action_Follow) fastWriteField2(buf []byte) (offset int) {
	if x.TargetId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetTargetId())
	return offset
}

func (x *Action_Follow) fastWriteField3(buf []byte) (offset int) {
	if x.TargetType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, int32(x.GetTargetType()))
	return offset
}

func (x *Action_Follow) fastWriteField4(buf []byte) (offset int) {
	if x.UserId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetUserId())
	return offset
}

func (x *Action_Follow) fastWriteField5(buf []byte) (offset int) {
	if x.CreateAt == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 5, x.GetCreateAt())
	return offset
}

func (x *Action_Follow) fastWriteField6(buf []byte) (offset int) {
	if !x.IsCancel {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 6, x.GetIsCancel())
	return offset
}

func (x *Action) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *Action_Like) Size() (n int) {
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

func (x *Action_Like) sizeField1() (n int) {
	if x.Id == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetId())
	return n
}

func (x *Action_Like) sizeField2() (n int) {
	if x.TargetId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetTargetId())
	return n
}

func (x *Action_Like) sizeField3() (n int) {
	if x.TargetType == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, int32(x.GetTargetType()))
	return n
}

func (x *Action_Like) sizeField4() (n int) {
	if x.UserId == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetUserId())
	return n
}

func (x *Action_Like) sizeField5() (n int) {
	if x.CreateAt == 0 {
		return n
	}
	n += fastpb.SizeInt64(5, x.GetCreateAt())
	return n
}

func (x *Action_Like) sizeField6() (n int) {
	if !x.IsCancel {
		return n
	}
	n += fastpb.SizeBool(6, x.GetIsCancel())
	return n
}

func (x *Action_Share) Size() (n int) {
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

func (x *Action_Share) sizeField1() (n int) {
	if x.Id == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetId())
	return n
}

func (x *Action_Share) sizeField2() (n int) {
	if x.TargetId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetTargetId())
	return n
}

func (x *Action_Share) sizeField3() (n int) {
	if x.TargetType == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, int32(x.GetTargetType()))
	return n
}

func (x *Action_Share) sizeField4() (n int) {
	if x.UserId == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetUserId())
	return n
}

func (x *Action_Share) sizeField5() (n int) {
	if x.CreateAt == 0 {
		return n
	}
	n += fastpb.SizeInt64(5, x.GetCreateAt())
	return n
}

func (x *Action_Follow) Size() (n int) {
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

func (x *Action_Follow) sizeField1() (n int) {
	if x.Id == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetId())
	return n
}

func (x *Action_Follow) sizeField2() (n int) {
	if x.TargetId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetTargetId())
	return n
}

func (x *Action_Follow) sizeField3() (n int) {
	if x.TargetType == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, int32(x.GetTargetType()))
	return n
}

func (x *Action_Follow) sizeField4() (n int) {
	if x.UserId == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetUserId())
	return n
}

func (x *Action_Follow) sizeField5() (n int) {
	if x.CreateAt == 0 {
		return n
	}
	n += fastpb.SizeInt64(5, x.GetCreateAt())
	return n
}

func (x *Action_Follow) sizeField6() (n int) {
	if !x.IsCancel {
		return n
	}
	n += fastpb.SizeBool(6, x.GetIsCancel())
	return n
}

var fieldIDToName_Action = map[int32]string{}

var fieldIDToName_Action_Like = map[int32]string{
	1: "Id",
	2: "TargetId",
	3: "TargetType",
	4: "UserId",
	5: "CreateAt",
	6: "IsCancel",
}

var fieldIDToName_Action_Share = map[int32]string{
	1: "Id",
	2: "TargetId",
	3: "TargetType",
	4: "UserId",
	5: "CreateAt",
}

var fieldIDToName_Action_Follow = map[int32]string{
	1: "Id",
	2: "TargetId",
	3: "TargetType",
	4: "UserId",
	5: "CreateAt",
	6: "IsCancel",
}
