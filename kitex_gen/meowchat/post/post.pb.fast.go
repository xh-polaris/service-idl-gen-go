// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package post

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
	basic "github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *CreatePostReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreatePostReq[number], err)
}

func (x *CreatePostReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Title, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreatePostReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Text, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreatePostReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.CoverUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreatePostReq) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	var v string
	v, offset, err = fastpb.ReadString(buf, _type)
	if err != nil {
		return offset, err
	}
	x.Tags = append(x.Tags, v)
	return offset, err
}

func (x *CreatePostReq) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreatePostResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreatePostResp[number], err)
}

func (x *CreatePostResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.PostId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RetrievePostReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RetrievePostReq[number], err)
}

func (x *RetrievePostReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.PostId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RetrievePostResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RetrievePostResp[number], err)
}

func (x *RetrievePostResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Post
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Post = &v
	return offset, nil
}

func (x *UpdatePostReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UpdatePostReq[number], err)
}

func (x *UpdatePostReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UpdatePostReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Title, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UpdatePostReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Text, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UpdatePostReq) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.CoverUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UpdatePostReq) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	var v string
	v, offset, err = fastpb.ReadString(buf, _type)
	if err != nil {
		return offset, err
	}
	x.Tags = append(x.Tags, v)
	return offset, err
}

func (x *UpdatePostResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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

func (x *DeletePostReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeletePostReq[number], err)
}

func (x *DeletePostReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DeletePostResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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

func (x *ListPostReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ListPostReq[number], err)
}

func (x *ListPostReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v SearchOptions
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.SearchOptions = &v
	return offset, nil
}

func (x *ListPostReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v FilterOptions
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.FilterOptions = &v
	return offset, nil
}

func (x *ListPostReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v basic.PaginationOptions
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.PaginationOptions = &v
	return offset, nil
}

func (x *ListPostResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ListPostResp[number], err)
}

func (x *ListPostResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Post
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Posts = append(x.Posts, &v)
	return offset, nil
}

func (x *ListPostResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Total, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *ListPostResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Token, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CountPostReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CountPostReq[number], err)
}

func (x *CountPostReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v SearchOptions
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.SearchOptions = &v
	return offset, nil
}

func (x *CountPostReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v FilterOptions
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.FilterOptions = &v
	return offset, nil
}

func (x *CountPostResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CountPostResp[number], err)
}

func (x *CountPostResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Total, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *SetOfficialReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SetOfficialReq[number], err)
}

func (x *SetOfficialReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.PostId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SetOfficialReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.IsRemove, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *SetOfficialResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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

func (x *CreatePostReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	return offset
}

func (x *CreatePostReq) fastWriteField2(buf []byte) (offset int) {
	if x.Title == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetTitle())
	return offset
}

func (x *CreatePostReq) fastWriteField3(buf []byte) (offset int) {
	if x.Text == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetText())
	return offset
}

func (x *CreatePostReq) fastWriteField4(buf []byte) (offset int) {
	if x.CoverUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetCoverUrl())
	return offset
}

func (x *CreatePostReq) fastWriteField5(buf []byte) (offset int) {
	if len(x.Tags) == 0 {
		return offset
	}
	for i := range x.GetTags() {
		offset += fastpb.WriteString(buf[offset:], 5, x.GetTags()[i])
	}
	return offset
}

func (x *CreatePostReq) fastWriteField6(buf []byte) (offset int) {
	if x.UserId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetUserId())
	return offset
}

func (x *CreatePostResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *CreatePostResp) fastWriteField1(buf []byte) (offset int) {
	if x.PostId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetPostId())
	return offset
}

func (x *RetrievePostReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *RetrievePostReq) fastWriteField1(buf []byte) (offset int) {
	if x.PostId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetPostId())
	return offset
}

func (x *RetrievePostResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *RetrievePostResp) fastWriteField1(buf []byte) (offset int) {
	if x.Post == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetPost())
	return offset
}

func (x *UpdatePostReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	return offset
}

func (x *UpdatePostReq) fastWriteField1(buf []byte) (offset int) {
	if x.Id == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetId())
	return offset
}

func (x *UpdatePostReq) fastWriteField3(buf []byte) (offset int) {
	if x.Title == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetTitle())
	return offset
}

func (x *UpdatePostReq) fastWriteField4(buf []byte) (offset int) {
	if x.Text == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetText())
	return offset
}

func (x *UpdatePostReq) fastWriteField5(buf []byte) (offset int) {
	if x.CoverUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetCoverUrl())
	return offset
}

func (x *UpdatePostReq) fastWriteField6(buf []byte) (offset int) {
	if len(x.Tags) == 0 {
		return offset
	}
	for i := range x.GetTags() {
		offset += fastpb.WriteString(buf[offset:], 6, x.GetTags()[i])
	}
	return offset
}

func (x *UpdatePostResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *DeletePostReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DeletePostReq) fastWriteField1(buf []byte) (offset int) {
	if x.Id == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetId())
	return offset
}

func (x *DeletePostResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *ListPostReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *ListPostReq) fastWriteField1(buf []byte) (offset int) {
	if x.SearchOptions == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetSearchOptions())
	return offset
}

func (x *ListPostReq) fastWriteField2(buf []byte) (offset int) {
	if x.FilterOptions == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetFilterOptions())
	return offset
}

func (x *ListPostReq) fastWriteField3(buf []byte) (offset int) {
	if x.PaginationOptions == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 3, x.GetPaginationOptions())
	return offset
}

func (x *ListPostResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *ListPostResp) fastWriteField1(buf []byte) (offset int) {
	if x.Posts == nil {
		return offset
	}
	for i := range x.GetPosts() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetPosts()[i])
	}
	return offset
}

func (x *ListPostResp) fastWriteField2(buf []byte) (offset int) {
	if x.Total == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetTotal())
	return offset
}

func (x *ListPostResp) fastWriteField3(buf []byte) (offset int) {
	if x.Token == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetToken())
	return offset
}

func (x *CountPostReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *CountPostReq) fastWriteField1(buf []byte) (offset int) {
	if x.SearchOptions == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetSearchOptions())
	return offset
}

func (x *CountPostReq) fastWriteField2(buf []byte) (offset int) {
	if x.FilterOptions == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetFilterOptions())
	return offset
}

func (x *CountPostResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *CountPostResp) fastWriteField1(buf []byte) (offset int) {
	if x.Total == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetTotal())
	return offset
}

func (x *SetOfficialReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *SetOfficialReq) fastWriteField1(buf []byte) (offset int) {
	if x.PostId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetPostId())
	return offset
}

func (x *SetOfficialReq) fastWriteField2(buf []byte) (offset int) {
	if !x.IsRemove {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 2, x.GetIsRemove())
	return offset
}

func (x *SetOfficialResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *CreatePostReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	return n
}

func (x *CreatePostReq) sizeField2() (n int) {
	if x.Title == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetTitle())
	return n
}

func (x *CreatePostReq) sizeField3() (n int) {
	if x.Text == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetText())
	return n
}

func (x *CreatePostReq) sizeField4() (n int) {
	if x.CoverUrl == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetCoverUrl())
	return n
}

func (x *CreatePostReq) sizeField5() (n int) {
	if len(x.Tags) == 0 {
		return n
	}
	for i := range x.GetTags() {
		n += fastpb.SizeString(5, x.GetTags()[i])
	}
	return n
}

func (x *CreatePostReq) sizeField6() (n int) {
	if x.UserId == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetUserId())
	return n
}

func (x *CreatePostResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *CreatePostResp) sizeField1() (n int) {
	if x.PostId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetPostId())
	return n
}

func (x *RetrievePostReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *RetrievePostReq) sizeField1() (n int) {
	if x.PostId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetPostId())
	return n
}

func (x *RetrievePostResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *RetrievePostResp) sizeField1() (n int) {
	if x.Post == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetPost())
	return n
}

func (x *UpdatePostReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	return n
}

func (x *UpdatePostReq) sizeField1() (n int) {
	if x.Id == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetId())
	return n
}

func (x *UpdatePostReq) sizeField3() (n int) {
	if x.Title == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetTitle())
	return n
}

func (x *UpdatePostReq) sizeField4() (n int) {
	if x.Text == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetText())
	return n
}

func (x *UpdatePostReq) sizeField5() (n int) {
	if x.CoverUrl == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetCoverUrl())
	return n
}

func (x *UpdatePostReq) sizeField6() (n int) {
	if len(x.Tags) == 0 {
		return n
	}
	for i := range x.GetTags() {
		n += fastpb.SizeString(6, x.GetTags()[i])
	}
	return n
}

func (x *UpdatePostResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *DeletePostReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DeletePostReq) sizeField1() (n int) {
	if x.Id == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetId())
	return n
}

func (x *DeletePostResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *ListPostReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *ListPostReq) sizeField1() (n int) {
	if x.SearchOptions == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetSearchOptions())
	return n
}

func (x *ListPostReq) sizeField2() (n int) {
	if x.FilterOptions == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetFilterOptions())
	return n
}

func (x *ListPostReq) sizeField3() (n int) {
	if x.PaginationOptions == nil {
		return n
	}
	n += fastpb.SizeMessage(3, x.GetPaginationOptions())
	return n
}

func (x *ListPostResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *ListPostResp) sizeField1() (n int) {
	if x.Posts == nil {
		return n
	}
	for i := range x.GetPosts() {
		n += fastpb.SizeMessage(1, x.GetPosts()[i])
	}
	return n
}

func (x *ListPostResp) sizeField2() (n int) {
	if x.Total == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetTotal())
	return n
}

func (x *ListPostResp) sizeField3() (n int) {
	if x.Token == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetToken())
	return n
}

func (x *CountPostReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *CountPostReq) sizeField1() (n int) {
	if x.SearchOptions == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetSearchOptions())
	return n
}

func (x *CountPostReq) sizeField2() (n int) {
	if x.FilterOptions == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetFilterOptions())
	return n
}

func (x *CountPostResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *CountPostResp) sizeField1() (n int) {
	if x.Total == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetTotal())
	return n
}

func (x *SetOfficialReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *SetOfficialReq) sizeField1() (n int) {
	if x.PostId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetPostId())
	return n
}

func (x *SetOfficialReq) sizeField2() (n int) {
	if !x.IsRemove {
		return n
	}
	n += fastpb.SizeBool(2, x.GetIsRemove())
	return n
}

func (x *SetOfficialResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

var fieldIDToName_CreatePostReq = map[int32]string{
	2: "Title",
	3: "Text",
	4: "CoverUrl",
	5: "Tags",
	6: "UserId",
}

var fieldIDToName_CreatePostResp = map[int32]string{
	1: "PostId",
}

var fieldIDToName_RetrievePostReq = map[int32]string{
	1: "PostId",
}

var fieldIDToName_RetrievePostResp = map[int32]string{
	1: "Post",
}

var fieldIDToName_UpdatePostReq = map[int32]string{
	1: "Id",
	3: "Title",
	4: "Text",
	5: "CoverUrl",
	6: "Tags",
}

var fieldIDToName_UpdatePostResp = map[int32]string{}

var fieldIDToName_DeletePostReq = map[int32]string{
	1: "Id",
}

var fieldIDToName_DeletePostResp = map[int32]string{}

var fieldIDToName_ListPostReq = map[int32]string{
	1: "SearchOptions",
	2: "FilterOptions",
	3: "PaginationOptions",
}

var fieldIDToName_ListPostResp = map[int32]string{
	1: "Posts",
	2: "Total",
	3: "Token",
}

var fieldIDToName_CountPostReq = map[int32]string{
	1: "SearchOptions",
	2: "FilterOptions",
}

var fieldIDToName_CountPostResp = map[int32]string{
	1: "Total",
}

var fieldIDToName_SetOfficialReq = map[int32]string{
	1: "PostId",
	2: "IsRemove",
}

var fieldIDToName_SetOfficialResp = map[int32]string{}

var _ = basic.File_basic_pagination_proto
