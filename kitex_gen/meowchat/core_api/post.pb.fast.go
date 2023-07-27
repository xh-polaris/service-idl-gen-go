// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package core_api

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
	basic "github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
	http "github.com/xh-polaris/service-idl-gen-go/kitex_gen/http"
	user "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/user"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *Post) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 10:
		offset, err = x.fastReadField10(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 11:
		offset, err = x.fastReadField11(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Post[number], err)
}

func (x *Post) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Post) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.CreateAt, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Post) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.UpdateAt, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Post) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Title, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Post) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.Text, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Post) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.CoverUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Post) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	var v string
	v, offset, err = fastpb.ReadString(buf, _type)
	if err != nil {
		return offset, err
	}
	x.Tags = append(x.Tags, v)
	return offset, err
}

func (x *Post) fastReadField9(buf []byte, _type int8) (offset int, err error) {
	var v user.UserPreview
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.User = &v
	return offset, nil
}

func (x *Post) fastReadField10(buf []byte, _type int8) (offset int, err error) {
	x.IsOfficial, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *Post) fastReadField11(buf []byte, _type int8) (offset int, err error) {
	x.Likes, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *SearchOptions) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SearchOptions[number], err)
}

func (x *SearchOptions) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Key = &tmp
	return offset, err
}

func (x *SearchOptions) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Title = &tmp
	return offset, err
}

func (x *SearchOptions) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Text = &tmp
	return offset, err
}

func (x *SearchOptions) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Tag = &tmp
	return offset, err
}

func (x *GetPostPreviewsReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetPostPreviewsReq[number], err)
}

func (x *GetPostPreviewsReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v basic.PaginationOptions
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.PaginationOption = &v
	return offset, nil
}

func (x *GetPostPreviewsReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadBool(buf, _type)
	x.OnlyOfficial = &tmp
	return offset, err
}

func (x *GetPostPreviewsReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.OnlyUserId = &tmp
	return offset, err
}

func (x *GetPostPreviewsReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	var v SearchOptions
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.SearchOptions = &v
	return offset, nil
}

func (x *GetPostPreviewsResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetPostPreviewsResp[number], err)
}

func (x *GetPostPreviewsResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Post
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Posts = append(x.Posts, &v)
	return offset, nil
}

func (x *GetPostPreviewsResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Total, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetPostPreviewsResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Token, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetPostDetailReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetPostDetailReq[number], err)
}

func (x *GetPostDetailReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.PostId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetPostDetailResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetPostDetailResp[number], err)
}

func (x *GetPostDetailResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Post
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Post = &v
	return offset, nil
}

func (x *NewPostReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_NewPostReq[number], err)
}

func (x *NewPostReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Id = &tmp
	return offset, err
}

func (x *NewPostReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Title, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *NewPostReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Text, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *NewPostReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.CoverUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *NewPostReq) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	var v string
	v, offset, err = fastpb.ReadString(buf, _type)
	if err != nil {
		return offset, err
	}
	x.Tags = append(x.Tags, v)
	return offset, err
}

func (x *NewPostResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_NewPostResp[number], err)
}

func (x *NewPostResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.PostId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
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
	tmp, offset, err := fastpb.ReadBool(buf, _type)
	x.IsRemove = &tmp
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

func (x *Post) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	offset += x.fastWriteField8(buf[offset:])
	offset += x.fastWriteField9(buf[offset:])
	offset += x.fastWriteField10(buf[offset:])
	offset += x.fastWriteField11(buf[offset:])
	return offset
}

func (x *Post) fastWriteField1(buf []byte) (offset int) {
	if x.Id == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetId())
	return offset
}

func (x *Post) fastWriteField2(buf []byte) (offset int) {
	if x.CreateAt == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetCreateAt())
	return offset
}

func (x *Post) fastWriteField3(buf []byte) (offset int) {
	if x.UpdateAt == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetUpdateAt())
	return offset
}

func (x *Post) fastWriteField5(buf []byte) (offset int) {
	if x.Title == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetTitle())
	return offset
}

func (x *Post) fastWriteField6(buf []byte) (offset int) {
	if x.Text == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetText())
	return offset
}

func (x *Post) fastWriteField7(buf []byte) (offset int) {
	if x.CoverUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 7, x.GetCoverUrl())
	return offset
}

func (x *Post) fastWriteField8(buf []byte) (offset int) {
	if len(x.Tags) == 0 {
		return offset
	}
	for i := range x.GetTags() {
		offset += fastpb.WriteString(buf[offset:], 8, x.GetTags()[i])
	}
	return offset
}

func (x *Post) fastWriteField9(buf []byte) (offset int) {
	if x.User == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 9, x.GetUser())
	return offset
}

func (x *Post) fastWriteField10(buf []byte) (offset int) {
	if !x.IsOfficial {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 10, x.GetIsOfficial())
	return offset
}

func (x *Post) fastWriteField11(buf []byte) (offset int) {
	if x.Likes == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 11, x.GetLikes())
	return offset
}

func (x *SearchOptions) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *SearchOptions) fastWriteField1(buf []byte) (offset int) {
	if x.Key == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetKey())
	return offset
}

func (x *SearchOptions) fastWriteField2(buf []byte) (offset int) {
	if x.Title == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetTitle())
	return offset
}

func (x *SearchOptions) fastWriteField3(buf []byte) (offset int) {
	if x.Text == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetText())
	return offset
}

func (x *SearchOptions) fastWriteField4(buf []byte) (offset int) {
	if x.Tag == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetTag())
	return offset
}

func (x *GetPostPreviewsReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *GetPostPreviewsReq) fastWriteField1(buf []byte) (offset int) {
	if x.PaginationOption == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetPaginationOption())
	return offset
}

func (x *GetPostPreviewsReq) fastWriteField2(buf []byte) (offset int) {
	if x.OnlyOfficial == nil {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 2, x.GetOnlyOfficial())
	return offset
}

func (x *GetPostPreviewsReq) fastWriteField3(buf []byte) (offset int) {
	if x.OnlyUserId == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetOnlyUserId())
	return offset
}

func (x *GetPostPreviewsReq) fastWriteField4(buf []byte) (offset int) {
	if x.SearchOptions == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 4, x.GetSearchOptions())
	return offset
}

func (x *GetPostPreviewsResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *GetPostPreviewsResp) fastWriteField1(buf []byte) (offset int) {
	if x.Posts == nil {
		return offset
	}
	for i := range x.GetPosts() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetPosts()[i])
	}
	return offset
}

func (x *GetPostPreviewsResp) fastWriteField2(buf []byte) (offset int) {
	if x.Total == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetTotal())
	return offset
}

func (x *GetPostPreviewsResp) fastWriteField3(buf []byte) (offset int) {
	if x.Token == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetToken())
	return offset
}

func (x *GetPostDetailReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetPostDetailReq) fastWriteField1(buf []byte) (offset int) {
	if x.PostId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetPostId())
	return offset
}

func (x *GetPostDetailResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetPostDetailResp) fastWriteField1(buf []byte) (offset int) {
	if x.Post == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetPost())
	return offset
}

func (x *NewPostReq) FastWrite(buf []byte) (offset int) {
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

func (x *NewPostReq) fastWriteField1(buf []byte) (offset int) {
	if x.Id == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetId())
	return offset
}

func (x *NewPostReq) fastWriteField2(buf []byte) (offset int) {
	if x.Title == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetTitle())
	return offset
}

func (x *NewPostReq) fastWriteField3(buf []byte) (offset int) {
	if x.Text == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetText())
	return offset
}

func (x *NewPostReq) fastWriteField4(buf []byte) (offset int) {
	if x.CoverUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetCoverUrl())
	return offset
}

func (x *NewPostReq) fastWriteField5(buf []byte) (offset int) {
	if len(x.Tags) == 0 {
		return offset
	}
	for i := range x.GetTags() {
		offset += fastpb.WriteString(buf[offset:], 5, x.GetTags()[i])
	}
	return offset
}

func (x *NewPostResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *NewPostResp) fastWriteField1(buf []byte) (offset int) {
	if x.PostId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetPostId())
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
	if x.IsRemove == nil {
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

func (x *Post) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	n += x.sizeField8()
	n += x.sizeField9()
	n += x.sizeField10()
	n += x.sizeField11()
	return n
}

func (x *Post) sizeField1() (n int) {
	if x.Id == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetId())
	return n
}

func (x *Post) sizeField2() (n int) {
	if x.CreateAt == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetCreateAt())
	return n
}

func (x *Post) sizeField3() (n int) {
	if x.UpdateAt == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetUpdateAt())
	return n
}

func (x *Post) sizeField5() (n int) {
	if x.Title == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetTitle())
	return n
}

func (x *Post) sizeField6() (n int) {
	if x.Text == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetText())
	return n
}

func (x *Post) sizeField7() (n int) {
	if x.CoverUrl == "" {
		return n
	}
	n += fastpb.SizeString(7, x.GetCoverUrl())
	return n
}

func (x *Post) sizeField8() (n int) {
	if len(x.Tags) == 0 {
		return n
	}
	for i := range x.GetTags() {
		n += fastpb.SizeString(8, x.GetTags()[i])
	}
	return n
}

func (x *Post) sizeField9() (n int) {
	if x.User == nil {
		return n
	}
	n += fastpb.SizeMessage(9, x.GetUser())
	return n
}

func (x *Post) sizeField10() (n int) {
	if !x.IsOfficial {
		return n
	}
	n += fastpb.SizeBool(10, x.GetIsOfficial())
	return n
}

func (x *Post) sizeField11() (n int) {
	if x.Likes == 0 {
		return n
	}
	n += fastpb.SizeInt64(11, x.GetLikes())
	return n
}

func (x *SearchOptions) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *SearchOptions) sizeField1() (n int) {
	if x.Key == nil {
		return n
	}
	n += fastpb.SizeString(1, x.GetKey())
	return n
}

func (x *SearchOptions) sizeField2() (n int) {
	if x.Title == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetTitle())
	return n
}

func (x *SearchOptions) sizeField3() (n int) {
	if x.Text == nil {
		return n
	}
	n += fastpb.SizeString(3, x.GetText())
	return n
}

func (x *SearchOptions) sizeField4() (n int) {
	if x.Tag == nil {
		return n
	}
	n += fastpb.SizeString(4, x.GetTag())
	return n
}

func (x *GetPostPreviewsReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *GetPostPreviewsReq) sizeField1() (n int) {
	if x.PaginationOption == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetPaginationOption())
	return n
}

func (x *GetPostPreviewsReq) sizeField2() (n int) {
	if x.OnlyOfficial == nil {
		return n
	}
	n += fastpb.SizeBool(2, x.GetOnlyOfficial())
	return n
}

func (x *GetPostPreviewsReq) sizeField3() (n int) {
	if x.OnlyUserId == nil {
		return n
	}
	n += fastpb.SizeString(3, x.GetOnlyUserId())
	return n
}

func (x *GetPostPreviewsReq) sizeField4() (n int) {
	if x.SearchOptions == nil {
		return n
	}
	n += fastpb.SizeMessage(4, x.GetSearchOptions())
	return n
}

func (x *GetPostPreviewsResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *GetPostPreviewsResp) sizeField1() (n int) {
	if x.Posts == nil {
		return n
	}
	for i := range x.GetPosts() {
		n += fastpb.SizeMessage(1, x.GetPosts()[i])
	}
	return n
}

func (x *GetPostPreviewsResp) sizeField2() (n int) {
	if x.Total == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetTotal())
	return n
}

func (x *GetPostPreviewsResp) sizeField3() (n int) {
	if x.Token == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetToken())
	return n
}

func (x *GetPostDetailReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetPostDetailReq) sizeField1() (n int) {
	if x.PostId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetPostId())
	return n
}

func (x *GetPostDetailResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetPostDetailResp) sizeField1() (n int) {
	if x.Post == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetPost())
	return n
}

func (x *NewPostReq) Size() (n int) {
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

func (x *NewPostReq) sizeField1() (n int) {
	if x.Id == nil {
		return n
	}
	n += fastpb.SizeString(1, x.GetId())
	return n
}

func (x *NewPostReq) sizeField2() (n int) {
	if x.Title == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetTitle())
	return n
}

func (x *NewPostReq) sizeField3() (n int) {
	if x.Text == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetText())
	return n
}

func (x *NewPostReq) sizeField4() (n int) {
	if x.CoverUrl == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetCoverUrl())
	return n
}

func (x *NewPostReq) sizeField5() (n int) {
	if len(x.Tags) == 0 {
		return n
	}
	for i := range x.GetTags() {
		n += fastpb.SizeString(5, x.GetTags()[i])
	}
	return n
}

func (x *NewPostResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *NewPostResp) sizeField1() (n int) {
	if x.PostId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetPostId())
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
	if x.IsRemove == nil {
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

var fieldIDToName_Post = map[int32]string{
	1:  "Id",
	2:  "CreateAt",
	3:  "UpdateAt",
	5:  "Title",
	6:  "Text",
	7:  "CoverUrl",
	8:  "Tags",
	9:  "User",
	10: "IsOfficial",
	11: "Likes",
}

var fieldIDToName_SearchOptions = map[int32]string{
	1: "Key",
	2: "Title",
	3: "Text",
	4: "Tag",
}

var fieldIDToName_GetPostPreviewsReq = map[int32]string{
	1: "PaginationOption",
	2: "OnlyOfficial",
	3: "OnlyUserId",
	4: "SearchOptions",
}

var fieldIDToName_GetPostPreviewsResp = map[int32]string{
	1: "Posts",
	2: "Total",
	3: "Token",
}

var fieldIDToName_GetPostDetailReq = map[int32]string{
	1: "PostId",
}

var fieldIDToName_GetPostDetailResp = map[int32]string{
	1: "Post",
}

var fieldIDToName_NewPostReq = map[int32]string{
	1: "Id",
	2: "Title",
	3: "Text",
	4: "CoverUrl",
	5: "Tags",
}

var fieldIDToName_NewPostResp = map[int32]string{
	1: "PostId",
}

var fieldIDToName_DeletePostReq = map[int32]string{
	1: "Id",
}

var fieldIDToName_DeletePostResp = map[int32]string{}

var fieldIDToName_SetOfficialReq = map[int32]string{
	1: "PostId",
	2: "IsRemove",
}

var fieldIDToName_SetOfficialResp = map[int32]string{}

var _ = http.File_http_http_proto
var _ = basic.File_basic_pagination_proto
var _ = user.File_meowchat_user_common_proto
