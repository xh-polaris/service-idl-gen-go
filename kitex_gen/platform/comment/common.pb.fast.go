// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package comment

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *Comment) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Comment[number], err)
}

func (x *Comment) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Comment) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Text, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Comment) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.AuthorId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Comment) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.ReplyTo, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Comment) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Type, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Comment) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.ParentId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Comment) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.UpdateAt, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Comment) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	x.CreateAt, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Comment) FastWrite(buf []byte) (offset int) {
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

func (x *Comment) fastWriteField1(buf []byte) (offset int) {
	if x.Id == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetId())
	return offset
}

func (x *Comment) fastWriteField2(buf []byte) (offset int) {
	if x.Text == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetText())
	return offset
}

func (x *Comment) fastWriteField3(buf []byte) (offset int) {
	if x.AuthorId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetAuthorId())
	return offset
}

func (x *Comment) fastWriteField4(buf []byte) (offset int) {
	if x.ReplyTo == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetReplyTo())
	return offset
}

func (x *Comment) fastWriteField5(buf []byte) (offset int) {
	if x.Type == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetType())
	return offset
}

func (x *Comment) fastWriteField6(buf []byte) (offset int) {
	if x.ParentId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetParentId())
	return offset
}

func (x *Comment) fastWriteField7(buf []byte) (offset int) {
	if x.UpdateAt == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 7, x.GetUpdateAt())
	return offset
}

func (x *Comment) fastWriteField8(buf []byte) (offset int) {
	if x.CreateAt == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 8, x.GetCreateAt())
	return offset
}

func (x *Comment) Size() (n int) {
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

func (x *Comment) sizeField1() (n int) {
	if x.Id == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetId())
	return n
}

func (x *Comment) sizeField2() (n int) {
	if x.Text == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetText())
	return n
}

func (x *Comment) sizeField3() (n int) {
	if x.AuthorId == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetAuthorId())
	return n
}

func (x *Comment) sizeField4() (n int) {
	if x.ReplyTo == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetReplyTo())
	return n
}

func (x *Comment) sizeField5() (n int) {
	if x.Type == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetType())
	return n
}

func (x *Comment) sizeField6() (n int) {
	if x.ParentId == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetParentId())
	return n
}

func (x *Comment) sizeField7() (n int) {
	if x.UpdateAt == 0 {
		return n
	}
	n += fastpb.SizeInt64(7, x.GetUpdateAt())
	return n
}

func (x *Comment) sizeField8() (n int) {
	if x.CreateAt == 0 {
		return n
	}
	n += fastpb.SizeInt64(8, x.GetCreateAt())
	return n
}

var fieldIDToName_Comment = map[int32]string{
	1: "Id",
	2: "Text",
	3: "AuthorId",
	4: "ReplyTo",
	5: "Type",
	6: "ParentId",
	7: "UpdateAt",
	8: "CreateAt",
}