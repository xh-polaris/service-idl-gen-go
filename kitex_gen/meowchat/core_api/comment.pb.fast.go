// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package core_api

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
	user "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/user"
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
	var v user.UserPreview
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.User = &v
	return offset, nil
}

func (x *Comment) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.ReplyName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Comment) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Comments, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Comment) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.UpdateAt, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Comment) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.CreateAt, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *NewCommentReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_NewCommentReq[number], err)
}

func (x *NewCommentReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Id = &tmp
	return offset, err
}

func (x *NewCommentReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Text, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *NewCommentReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Scope, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *NewCommentResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_NewCommentResp[number], err)
}

func (x *NewCommentResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.IsFirst, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *GetCommentsReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetCommentsReq[number], err)
}

func (x *GetCommentsReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetCommentsReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Scope, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetCommentsReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Page, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetCommentsResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetCommentsResp[number], err)
}

func (x *GetCommentsResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Comment
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Comments = append(x.Comments, &v)
	return offset, nil
}

func (x *GetCommentsResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Total, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DeleteCommentReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeleteCommentReq[number], err)
}

func (x *DeleteCommentReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.CommentId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DeleteCommentResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	if x.User == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 3, x.GetUser())
	return offset
}

func (x *Comment) fastWriteField4(buf []byte) (offset int) {
	if x.ReplyName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetReplyName())
	return offset
}

func (x *Comment) fastWriteField5(buf []byte) (offset int) {
	if x.Comments == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 5, x.GetComments())
	return offset
}

func (x *Comment) fastWriteField6(buf []byte) (offset int) {
	if x.UpdateAt == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 6, x.GetUpdateAt())
	return offset
}

func (x *Comment) fastWriteField7(buf []byte) (offset int) {
	if x.CreateAt == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 7, x.GetCreateAt())
	return offset
}

func (x *NewCommentReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *NewCommentReq) fastWriteField1(buf []byte) (offset int) {
	if x.Id == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetId())
	return offset
}

func (x *NewCommentReq) fastWriteField2(buf []byte) (offset int) {
	if x.Text == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetText())
	return offset
}

func (x *NewCommentReq) fastWriteField3(buf []byte) (offset int) {
	if x.Scope == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetScope())
	return offset
}

func (x *NewCommentResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *NewCommentResp) fastWriteField1(buf []byte) (offset int) {
	if !x.IsFirst {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetIsFirst())
	return offset
}

func (x *GetCommentsReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *GetCommentsReq) fastWriteField1(buf []byte) (offset int) {
	if x.Id == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetId())
	return offset
}

func (x *GetCommentsReq) fastWriteField2(buf []byte) (offset int) {
	if x.Scope == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetScope())
	return offset
}

func (x *GetCommentsReq) fastWriteField3(buf []byte) (offset int) {
	if x.Page == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetPage())
	return offset
}

func (x *GetCommentsResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetCommentsResp) fastWriteField1(buf []byte) (offset int) {
	if x.Comments == nil {
		return offset
	}
	for i := range x.GetComments() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetComments()[i])
	}
	return offset
}

func (x *GetCommentsResp) fastWriteField2(buf []byte) (offset int) {
	if x.Total == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetTotal())
	return offset
}

func (x *DeleteCommentReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DeleteCommentReq) fastWriteField1(buf []byte) (offset int) {
	if x.CommentId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetCommentId())
	return offset
}

func (x *DeleteCommentResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
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
	if x.User == nil {
		return n
	}
	n += fastpb.SizeMessage(3, x.GetUser())
	return n
}

func (x *Comment) sizeField4() (n int) {
	if x.ReplyName == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetReplyName())
	return n
}

func (x *Comment) sizeField5() (n int) {
	if x.Comments == 0 {
		return n
	}
	n += fastpb.SizeInt64(5, x.GetComments())
	return n
}

func (x *Comment) sizeField6() (n int) {
	if x.UpdateAt == 0 {
		return n
	}
	n += fastpb.SizeInt64(6, x.GetUpdateAt())
	return n
}

func (x *Comment) sizeField7() (n int) {
	if x.CreateAt == 0 {
		return n
	}
	n += fastpb.SizeInt64(7, x.GetCreateAt())
	return n
}

func (x *NewCommentReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *NewCommentReq) sizeField1() (n int) {
	if x.Id == nil {
		return n
	}
	n += fastpb.SizeString(1, x.GetId())
	return n
}

func (x *NewCommentReq) sizeField2() (n int) {
	if x.Text == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetText())
	return n
}

func (x *NewCommentReq) sizeField3() (n int) {
	if x.Scope == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetScope())
	return n
}

func (x *NewCommentResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *NewCommentResp) sizeField1() (n int) {
	if !x.IsFirst {
		return n
	}
	n += fastpb.SizeBool(1, x.GetIsFirst())
	return n
}

func (x *GetCommentsReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *GetCommentsReq) sizeField1() (n int) {
	if x.Id == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetId())
	return n
}

func (x *GetCommentsReq) sizeField2() (n int) {
	if x.Scope == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetScope())
	return n
}

func (x *GetCommentsReq) sizeField3() (n int) {
	if x.Page == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetPage())
	return n
}

func (x *GetCommentsResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetCommentsResp) sizeField1() (n int) {
	if x.Comments == nil {
		return n
	}
	for i := range x.GetComments() {
		n += fastpb.SizeMessage(1, x.GetComments()[i])
	}
	return n
}

func (x *GetCommentsResp) sizeField2() (n int) {
	if x.Total == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetTotal())
	return n
}

func (x *DeleteCommentReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DeleteCommentReq) sizeField1() (n int) {
	if x.CommentId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetCommentId())
	return n
}

func (x *DeleteCommentResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

var fieldIDToName_Comment = map[int32]string{
	1: "Id",
	2: "Text",
	3: "User",
	4: "ReplyName",
	5: "Comments",
	6: "UpdateAt",
	7: "CreateAt",
}

var fieldIDToName_NewCommentReq = map[int32]string{
	1: "Id",
	2: "Text",
	3: "Scope",
}

var fieldIDToName_NewCommentResp = map[int32]string{
	1: "IsFirst",
}

var fieldIDToName_GetCommentsReq = map[int32]string{
	1: "Id",
	2: "Scope",
	3: "Page",
}

var fieldIDToName_GetCommentsResp = map[int32]string{
	1: "Comments",
	2: "Total",
}

var fieldIDToName_DeleteCommentReq = map[int32]string{
	1: "CommentId",
}

var fieldIDToName_DeleteCommentResp = map[int32]string{}

var _ = user.File_meowchat_user_common_proto
