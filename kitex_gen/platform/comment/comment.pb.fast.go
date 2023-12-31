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

func (x *CreateCommentReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreateCommentReq[number], err)
}

func (x *CreateCommentReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Text, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateCommentReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.FirstLevelId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateCommentReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.AuthorId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateCommentReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.ReplyTo, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateCommentReq) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.Type = CommentType(v)
	return offset, nil
}

func (x *CreateCommentReq) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.ParentId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateCommentResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreateCommentResp[number], err)
}

func (x *CreateCommentResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateCommentResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.GetFish, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *CreateCommentResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.GetFishTimes, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *UpdateCommentReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UpdateCommentReq[number], err)
}

func (x *UpdateCommentReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UpdateCommentReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Text, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UpdateCommentResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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

func (x *DeleteCommentByIdReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeleteCommentByIdReq[number], err)
}

func (x *DeleteCommentByIdReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DeleteCommentByIdResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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

func (x *ListCommentByParentReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ListCommentByParentReq[number], err)
}

func (x *ListCommentByParentReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.Type = CommentType(v)
	return offset, nil
}

func (x *ListCommentByParentReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ListCommentByParentReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Skip, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *ListCommentByParentReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Limit, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *ListCommentByParentReq) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadBool(buf, _type)
	x.OnlyFirstLevel = &tmp
	return offset, err
}

func (x *ListCommentByParentResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ListCommentByParentResp[number], err)
}

func (x *ListCommentByParentResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Comment
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Comments = append(x.Comments, &v)
	return offset, nil
}

func (x *ListCommentByParentResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Total, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CountCommentByParentReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CountCommentByParentReq[number], err)
}

func (x *CountCommentByParentReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.Type = CommentType(v)
	return offset, nil
}

func (x *CountCommentByParentReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ParentId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CountCommentByParentReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadBool(buf, _type)
	x.OnlyFirstLevel = &tmp
	return offset, err
}

func (x *CountCommentByParentResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CountCommentByParentResp[number], err)
}

func (x *CountCommentByParentResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Total, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *RetrieveCommentByIdReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RetrieveCommentByIdReq[number], err)
}

func (x *RetrieveCommentByIdReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RetrieveCommentByIdResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RetrieveCommentByIdResp[number], err)
}

func (x *RetrieveCommentByIdResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Comment
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Comment = &v
	return offset, nil
}

func (x *ListCommentByAuthorIdAndTypeReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ListCommentByAuthorIdAndTypeReq[number], err)
}

func (x *ListCommentByAuthorIdAndTypeReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.AuthorId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ListCommentByAuthorIdAndTypeReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.Type = CommentType(v)
	return offset, nil
}

func (x *ListCommentByAuthorIdAndTypeReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Skip, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *ListCommentByAuthorIdAndTypeReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Limit, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *ListCommentByAuthorIdAndTypeResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ListCommentByAuthorIdAndTypeResp[number], err)
}

func (x *ListCommentByAuthorIdAndTypeResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Comment
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Comments = append(x.Comments, &v)
	return offset, nil
}

func (x *ListCommentByAuthorIdAndTypeResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Total, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *ListCommentByReplyToAndTypeReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ListCommentByReplyToAndTypeReq[number], err)
}

func (x *ListCommentByReplyToAndTypeReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ReplyTo, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ListCommentByReplyToAndTypeReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.Type = CommentType(v)
	return offset, nil
}

func (x *ListCommentByReplyToAndTypeReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Skip, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *ListCommentByReplyToAndTypeReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Limit, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *ListCommentByReplyToAndTypeResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ListCommentByReplyToAndTypeResp[number], err)
}

func (x *ListCommentByReplyToAndTypeResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Comment
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Comments = append(x.Comments, &v)
	return offset, nil
}

func (x *ListCommentByReplyToAndTypeResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Total, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CreateCommentReq) FastWrite(buf []byte) (offset int) {
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

func (x *CreateCommentReq) fastWriteField1(buf []byte) (offset int) {
	if x.Text == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetText())
	return offset
}

func (x *CreateCommentReq) fastWriteField2(buf []byte) (offset int) {
	if x.FirstLevelId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetFirstLevelId())
	return offset
}

func (x *CreateCommentReq) fastWriteField3(buf []byte) (offset int) {
	if x.AuthorId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetAuthorId())
	return offset
}

func (x *CreateCommentReq) fastWriteField4(buf []byte) (offset int) {
	if x.ReplyTo == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetReplyTo())
	return offset
}

func (x *CreateCommentReq) fastWriteField5(buf []byte) (offset int) {
	if x.Type == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 5, int32(x.GetType()))
	return offset
}

func (x *CreateCommentReq) fastWriteField6(buf []byte) (offset int) {
	if x.ParentId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetParentId())
	return offset
}

func (x *CreateCommentResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *CreateCommentResp) fastWriteField1(buf []byte) (offset int) {
	if x.Id == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetId())
	return offset
}

func (x *CreateCommentResp) fastWriteField2(buf []byte) (offset int) {
	if !x.GetFish {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 2, x.GetGetFish())
	return offset
}

func (x *CreateCommentResp) fastWriteField3(buf []byte) (offset int) {
	if x.GetFishTimes == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetGetFishTimes())
	return offset
}

func (x *UpdateCommentReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *UpdateCommentReq) fastWriteField1(buf []byte) (offset int) {
	if x.Id == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetId())
	return offset
}

func (x *UpdateCommentReq) fastWriteField2(buf []byte) (offset int) {
	if x.Text == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetText())
	return offset
}

func (x *UpdateCommentResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *DeleteCommentByIdReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DeleteCommentByIdReq) fastWriteField1(buf []byte) (offset int) {
	if x.Id == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetId())
	return offset
}

func (x *DeleteCommentByIdResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *ListCommentByParentReq) FastWrite(buf []byte) (offset int) {
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

func (x *ListCommentByParentReq) fastWriteField1(buf []byte) (offset int) {
	if x.Type == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, int32(x.GetType()))
	return offset
}

func (x *ListCommentByParentReq) fastWriteField2(buf []byte) (offset int) {
	if x.Id == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetId())
	return offset
}

func (x *ListCommentByParentReq) fastWriteField3(buf []byte) (offset int) {
	if x.Skip == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetSkip())
	return offset
}

func (x *ListCommentByParentReq) fastWriteField4(buf []byte) (offset int) {
	if x.Limit == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.GetLimit())
	return offset
}

func (x *ListCommentByParentReq) fastWriteField5(buf []byte) (offset int) {
	if x.OnlyFirstLevel == nil {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 5, x.GetOnlyFirstLevel())
	return offset
}

func (x *ListCommentByParentResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *ListCommentByParentResp) fastWriteField1(buf []byte) (offset int) {
	if x.Comments == nil {
		return offset
	}
	for i := range x.GetComments() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetComments()[i])
	}
	return offset
}

func (x *ListCommentByParentResp) fastWriteField2(buf []byte) (offset int) {
	if x.Total == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetTotal())
	return offset
}

func (x *CountCommentByParentReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *CountCommentByParentReq) fastWriteField1(buf []byte) (offset int) {
	if x.Type == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, int32(x.GetType()))
	return offset
}

func (x *CountCommentByParentReq) fastWriteField2(buf []byte) (offset int) {
	if x.ParentId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetParentId())
	return offset
}

func (x *CountCommentByParentReq) fastWriteField3(buf []byte) (offset int) {
	if x.OnlyFirstLevel == nil {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 3, x.GetOnlyFirstLevel())
	return offset
}

func (x *CountCommentByParentResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *CountCommentByParentResp) fastWriteField1(buf []byte) (offset int) {
	if x.Total == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetTotal())
	return offset
}

func (x *RetrieveCommentByIdReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *RetrieveCommentByIdReq) fastWriteField1(buf []byte) (offset int) {
	if x.Id == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetId())
	return offset
}

func (x *RetrieveCommentByIdResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *RetrieveCommentByIdResp) fastWriteField1(buf []byte) (offset int) {
	if x.Comment == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetComment())
	return offset
}

func (x *ListCommentByAuthorIdAndTypeReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *ListCommentByAuthorIdAndTypeReq) fastWriteField1(buf []byte) (offset int) {
	if x.AuthorId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetAuthorId())
	return offset
}

func (x *ListCommentByAuthorIdAndTypeReq) fastWriteField2(buf []byte) (offset int) {
	if x.Type == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, int32(x.GetType()))
	return offset
}

func (x *ListCommentByAuthorIdAndTypeReq) fastWriteField3(buf []byte) (offset int) {
	if x.Skip == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetSkip())
	return offset
}

func (x *ListCommentByAuthorIdAndTypeReq) fastWriteField4(buf []byte) (offset int) {
	if x.Limit == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.GetLimit())
	return offset
}

func (x *ListCommentByAuthorIdAndTypeResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *ListCommentByAuthorIdAndTypeResp) fastWriteField1(buf []byte) (offset int) {
	if x.Comments == nil {
		return offset
	}
	for i := range x.GetComments() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetComments()[i])
	}
	return offset
}

func (x *ListCommentByAuthorIdAndTypeResp) fastWriteField2(buf []byte) (offset int) {
	if x.Total == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetTotal())
	return offset
}

func (x *ListCommentByReplyToAndTypeReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *ListCommentByReplyToAndTypeReq) fastWriteField1(buf []byte) (offset int) {
	if x.ReplyTo == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetReplyTo())
	return offset
}

func (x *ListCommentByReplyToAndTypeReq) fastWriteField2(buf []byte) (offset int) {
	if x.Type == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, int32(x.GetType()))
	return offset
}

func (x *ListCommentByReplyToAndTypeReq) fastWriteField3(buf []byte) (offset int) {
	if x.Skip == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetSkip())
	return offset
}

func (x *ListCommentByReplyToAndTypeReq) fastWriteField4(buf []byte) (offset int) {
	if x.Limit == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.GetLimit())
	return offset
}

func (x *ListCommentByReplyToAndTypeResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *ListCommentByReplyToAndTypeResp) fastWriteField1(buf []byte) (offset int) {
	if x.Comments == nil {
		return offset
	}
	for i := range x.GetComments() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetComments()[i])
	}
	return offset
}

func (x *ListCommentByReplyToAndTypeResp) fastWriteField2(buf []byte) (offset int) {
	if x.Total == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetTotal())
	return offset
}

func (x *CreateCommentReq) Size() (n int) {
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

func (x *CreateCommentReq) sizeField1() (n int) {
	if x.Text == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetText())
	return n
}

func (x *CreateCommentReq) sizeField2() (n int) {
	if x.FirstLevelId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetFirstLevelId())
	return n
}

func (x *CreateCommentReq) sizeField3() (n int) {
	if x.AuthorId == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetAuthorId())
	return n
}

func (x *CreateCommentReq) sizeField4() (n int) {
	if x.ReplyTo == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetReplyTo())
	return n
}

func (x *CreateCommentReq) sizeField5() (n int) {
	if x.Type == 0 {
		return n
	}
	n += fastpb.SizeInt32(5, int32(x.GetType()))
	return n
}

func (x *CreateCommentReq) sizeField6() (n int) {
	if x.ParentId == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetParentId())
	return n
}

func (x *CreateCommentResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *CreateCommentResp) sizeField1() (n int) {
	if x.Id == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetId())
	return n
}

func (x *CreateCommentResp) sizeField2() (n int) {
	if !x.GetFish {
		return n
	}
	n += fastpb.SizeBool(2, x.GetGetFish())
	return n
}

func (x *CreateCommentResp) sizeField3() (n int) {
	if x.GetFishTimes == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetGetFishTimes())
	return n
}

func (x *UpdateCommentReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *UpdateCommentReq) sizeField1() (n int) {
	if x.Id == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetId())
	return n
}

func (x *UpdateCommentReq) sizeField2() (n int) {
	if x.Text == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetText())
	return n
}

func (x *UpdateCommentResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *DeleteCommentByIdReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DeleteCommentByIdReq) sizeField1() (n int) {
	if x.Id == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetId())
	return n
}

func (x *DeleteCommentByIdResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *ListCommentByParentReq) Size() (n int) {
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

func (x *ListCommentByParentReq) sizeField1() (n int) {
	if x.Type == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, int32(x.GetType()))
	return n
}

func (x *ListCommentByParentReq) sizeField2() (n int) {
	if x.Id == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetId())
	return n
}

func (x *ListCommentByParentReq) sizeField3() (n int) {
	if x.Skip == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetSkip())
	return n
}

func (x *ListCommentByParentReq) sizeField4() (n int) {
	if x.Limit == 0 {
		return n
	}
	n += fastpb.SizeInt64(4, x.GetLimit())
	return n
}

func (x *ListCommentByParentReq) sizeField5() (n int) {
	if x.OnlyFirstLevel == nil {
		return n
	}
	n += fastpb.SizeBool(5, x.GetOnlyFirstLevel())
	return n
}

func (x *ListCommentByParentResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *ListCommentByParentResp) sizeField1() (n int) {
	if x.Comments == nil {
		return n
	}
	for i := range x.GetComments() {
		n += fastpb.SizeMessage(1, x.GetComments()[i])
	}
	return n
}

func (x *ListCommentByParentResp) sizeField2() (n int) {
	if x.Total == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetTotal())
	return n
}

func (x *CountCommentByParentReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *CountCommentByParentReq) sizeField1() (n int) {
	if x.Type == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, int32(x.GetType()))
	return n
}

func (x *CountCommentByParentReq) sizeField2() (n int) {
	if x.ParentId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetParentId())
	return n
}

func (x *CountCommentByParentReq) sizeField3() (n int) {
	if x.OnlyFirstLevel == nil {
		return n
	}
	n += fastpb.SizeBool(3, x.GetOnlyFirstLevel())
	return n
}

func (x *CountCommentByParentResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *CountCommentByParentResp) sizeField1() (n int) {
	if x.Total == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetTotal())
	return n
}

func (x *RetrieveCommentByIdReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *RetrieveCommentByIdReq) sizeField1() (n int) {
	if x.Id == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetId())
	return n
}

func (x *RetrieveCommentByIdResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *RetrieveCommentByIdResp) sizeField1() (n int) {
	if x.Comment == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetComment())
	return n
}

func (x *ListCommentByAuthorIdAndTypeReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *ListCommentByAuthorIdAndTypeReq) sizeField1() (n int) {
	if x.AuthorId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetAuthorId())
	return n
}

func (x *ListCommentByAuthorIdAndTypeReq) sizeField2() (n int) {
	if x.Type == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, int32(x.GetType()))
	return n
}

func (x *ListCommentByAuthorIdAndTypeReq) sizeField3() (n int) {
	if x.Skip == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetSkip())
	return n
}

func (x *ListCommentByAuthorIdAndTypeReq) sizeField4() (n int) {
	if x.Limit == 0 {
		return n
	}
	n += fastpb.SizeInt64(4, x.GetLimit())
	return n
}

func (x *ListCommentByAuthorIdAndTypeResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *ListCommentByAuthorIdAndTypeResp) sizeField1() (n int) {
	if x.Comments == nil {
		return n
	}
	for i := range x.GetComments() {
		n += fastpb.SizeMessage(1, x.GetComments()[i])
	}
	return n
}

func (x *ListCommentByAuthorIdAndTypeResp) sizeField2() (n int) {
	if x.Total == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetTotal())
	return n
}

func (x *ListCommentByReplyToAndTypeReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *ListCommentByReplyToAndTypeReq) sizeField1() (n int) {
	if x.ReplyTo == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetReplyTo())
	return n
}

func (x *ListCommentByReplyToAndTypeReq) sizeField2() (n int) {
	if x.Type == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, int32(x.GetType()))
	return n
}

func (x *ListCommentByReplyToAndTypeReq) sizeField3() (n int) {
	if x.Skip == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetSkip())
	return n
}

func (x *ListCommentByReplyToAndTypeReq) sizeField4() (n int) {
	if x.Limit == 0 {
		return n
	}
	n += fastpb.SizeInt64(4, x.GetLimit())
	return n
}

func (x *ListCommentByReplyToAndTypeResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *ListCommentByReplyToAndTypeResp) sizeField1() (n int) {
	if x.Comments == nil {
		return n
	}
	for i := range x.GetComments() {
		n += fastpb.SizeMessage(1, x.GetComments()[i])
	}
	return n
}

func (x *ListCommentByReplyToAndTypeResp) sizeField2() (n int) {
	if x.Total == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetTotal())
	return n
}

var fieldIDToName_CreateCommentReq = map[int32]string{
	1: "Text",
	2: "FirstLevelId",
	3: "AuthorId",
	4: "ReplyTo",
	5: "Type",
	6: "ParentId",
}

var fieldIDToName_CreateCommentResp = map[int32]string{
	1: "Id",
	2: "GetFish",
	3: "GetFishTimes",
}

var fieldIDToName_UpdateCommentReq = map[int32]string{
	1: "Id",
	2: "Text",
}

var fieldIDToName_UpdateCommentResp = map[int32]string{}

var fieldIDToName_DeleteCommentByIdReq = map[int32]string{
	1: "Id",
}

var fieldIDToName_DeleteCommentByIdResp = map[int32]string{}

var fieldIDToName_ListCommentByParentReq = map[int32]string{
	1: "Type",
	2: "Id",
	3: "Skip",
	4: "Limit",
	5: "OnlyFirstLevel",
}

var fieldIDToName_ListCommentByParentResp = map[int32]string{
	1: "Comments",
	2: "Total",
}

var fieldIDToName_CountCommentByParentReq = map[int32]string{
	1: "Type",
	2: "ParentId",
	3: "OnlyFirstLevel",
}

var fieldIDToName_CountCommentByParentResp = map[int32]string{
	1: "Total",
}

var fieldIDToName_RetrieveCommentByIdReq = map[int32]string{
	1: "Id",
}

var fieldIDToName_RetrieveCommentByIdResp = map[int32]string{
	1: "Comment",
}

var fieldIDToName_ListCommentByAuthorIdAndTypeReq = map[int32]string{
	1: "AuthorId",
	2: "Type",
	3: "Skip",
	4: "Limit",
}

var fieldIDToName_ListCommentByAuthorIdAndTypeResp = map[int32]string{
	1: "Comments",
	2: "Total",
}

var fieldIDToName_ListCommentByReplyToAndTypeReq = map[int32]string{
	1: "ReplyTo",
	2: "Type",
	3: "Skip",
	4: "Limit",
}

var fieldIDToName_ListCommentByReplyToAndTypeResp = map[int32]string{
	1: "Comments",
	2: "Total",
}
