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

func (x *Moment) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Moment[number], err)
}

func (x *Moment) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Moment) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.CreateAt, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Moment) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.CatId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Moment) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	var v string
	v, offset, err = fastpb.ReadString(buf, _type)
	if err != nil {
		return offset, err
	}
	x.Photos = append(x.Photos, v)
	return offset, err
}

func (x *Moment) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Title, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Moment) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.Text, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Moment) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.CommunityId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Moment) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	var v user.UserPreview
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.User = &v
	return offset, nil
}

func (x *GetMomentPreviewsReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetMomentPreviewsReq[number], err)
}

func (x *GetMomentPreviewsReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.CommunityId = &tmp
	return offset, err
}

func (x *GetMomentPreviewsReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadBool(buf, _type)
	x.IsParent = &tmp
	return offset, err
}

func (x *GetMomentPreviewsReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.OnlyUserId = &tmp
	return offset, err
}

func (x *GetMomentPreviewsReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	var v basic.PaginationOptions
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.PaginationOption = &v
	return offset, nil
}

func (x *GetMomentPreviewsResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetMomentPreviewsResp[number], err)
}

func (x *GetMomentPreviewsResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Moment
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Moments = append(x.Moments, &v)
	return offset, nil
}

func (x *GetMomentPreviewsResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Total, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetMomentPreviewsResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Token, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetMomentDetailReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetMomentDetailReq[number], err)
}

func (x *GetMomentDetailReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.MomentId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetMomentDetailResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetMomentDetailResp[number], err)
}

func (x *GetMomentDetailResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Moment
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Moment = &v
	return offset, nil
}

func (x *DeleteMomentReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeleteMomentReq[number], err)
}

func (x *DeleteMomentReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.MomentId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DeleteMomentResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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

func (x *NewMomentReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_NewMomentReq[number], err)
}

func (x *NewMomentReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Id = &tmp
	return offset, err
}

func (x *NewMomentReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Title = &tmp
	return offset, err
}

func (x *NewMomentReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.CatId = &tmp
	return offset, err
}

func (x *NewMomentReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Text = &tmp
	return offset, err
}

func (x *NewMomentReq) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	var v string
	v, offset, err = fastpb.ReadString(buf, _type)
	if err != nil {
		return offset, err
	}
	x.Photos = append(x.Photos, v)
	return offset, err
}

func (x *NewMomentReq) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.CommunityId = &tmp
	return offset, err
}

func (x *NewMomentResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_NewMomentResp[number], err)
}

func (x *NewMomentResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.MomentId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SearchMomentReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SearchMomentReq[number], err)
}

func (x *SearchMomentReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.CommunityId = &tmp
	return offset, err
}

func (x *SearchMomentReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadBool(buf, _type)
	x.IsParent = &tmp
	return offset, err
}

func (x *SearchMomentReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.OnlyUserId = &tmp
	return offset, err
}

func (x *SearchMomentReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Keyword = &tmp
	return offset, err
}

func (x *SearchMomentReq) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	var v basic.PaginationOptions
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.PaginationOption = &v
	return offset, nil
}

func (x *SearchMomentResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SearchMomentResp[number], err)
}

func (x *SearchMomentResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Moment
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Moments = append(x.Moments, &v)
	return offset, nil
}

func (x *SearchMomentResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Total, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Moment) FastWrite(buf []byte) (offset int) {
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

func (x *Moment) fastWriteField1(buf []byte) (offset int) {
	if x.Id == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetId())
	return offset
}

func (x *Moment) fastWriteField2(buf []byte) (offset int) {
	if x.CreateAt == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetCreateAt())
	return offset
}

func (x *Moment) fastWriteField3(buf []byte) (offset int) {
	if x.CatId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetCatId())
	return offset
}

func (x *Moment) fastWriteField4(buf []byte) (offset int) {
	if len(x.Photos) == 0 {
		return offset
	}
	for i := range x.GetPhotos() {
		offset += fastpb.WriteString(buf[offset:], 4, x.GetPhotos()[i])
	}
	return offset
}

func (x *Moment) fastWriteField5(buf []byte) (offset int) {
	if x.Title == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetTitle())
	return offset
}

func (x *Moment) fastWriteField6(buf []byte) (offset int) {
	if x.Text == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetText())
	return offset
}

func (x *Moment) fastWriteField7(buf []byte) (offset int) {
	if x.CommunityId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 7, x.GetCommunityId())
	return offset
}

func (x *Moment) fastWriteField8(buf []byte) (offset int) {
	if x.User == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 8, x.GetUser())
	return offset
}

func (x *GetMomentPreviewsReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *GetMomentPreviewsReq) fastWriteField1(buf []byte) (offset int) {
	if x.CommunityId == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetCommunityId())
	return offset
}

func (x *GetMomentPreviewsReq) fastWriteField2(buf []byte) (offset int) {
	if x.IsParent == nil {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 2, x.GetIsParent())
	return offset
}

func (x *GetMomentPreviewsReq) fastWriteField3(buf []byte) (offset int) {
	if x.OnlyUserId == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetOnlyUserId())
	return offset
}

func (x *GetMomentPreviewsReq) fastWriteField4(buf []byte) (offset int) {
	if x.PaginationOption == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 4, x.GetPaginationOption())
	return offset
}

func (x *GetMomentPreviewsResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *GetMomentPreviewsResp) fastWriteField1(buf []byte) (offset int) {
	if x.Moments == nil {
		return offset
	}
	for i := range x.GetMoments() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetMoments()[i])
	}
	return offset
}

func (x *GetMomentPreviewsResp) fastWriteField2(buf []byte) (offset int) {
	if x.Total == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetTotal())
	return offset
}

func (x *GetMomentPreviewsResp) fastWriteField3(buf []byte) (offset int) {
	if x.Token == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetToken())
	return offset
}

func (x *GetMomentDetailReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetMomentDetailReq) fastWriteField1(buf []byte) (offset int) {
	if x.MomentId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetMomentId())
	return offset
}

func (x *GetMomentDetailResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetMomentDetailResp) fastWriteField1(buf []byte) (offset int) {
	if x.Moment == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetMoment())
	return offset
}

func (x *DeleteMomentReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DeleteMomentReq) fastWriteField1(buf []byte) (offset int) {
	if x.MomentId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetMomentId())
	return offset
}

func (x *DeleteMomentResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *NewMomentReq) FastWrite(buf []byte) (offset int) {
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

func (x *NewMomentReq) fastWriteField1(buf []byte) (offset int) {
	if x.Id == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetId())
	return offset
}

func (x *NewMomentReq) fastWriteField2(buf []byte) (offset int) {
	if x.Title == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetTitle())
	return offset
}

func (x *NewMomentReq) fastWriteField3(buf []byte) (offset int) {
	if x.CatId == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetCatId())
	return offset
}

func (x *NewMomentReq) fastWriteField4(buf []byte) (offset int) {
	if x.Text == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetText())
	return offset
}

func (x *NewMomentReq) fastWriteField5(buf []byte) (offset int) {
	if len(x.Photos) == 0 {
		return offset
	}
	for i := range x.GetPhotos() {
		offset += fastpb.WriteString(buf[offset:], 5, x.GetPhotos()[i])
	}
	return offset
}

func (x *NewMomentReq) fastWriteField6(buf []byte) (offset int) {
	if x.CommunityId == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetCommunityId())
	return offset
}

func (x *NewMomentResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *NewMomentResp) fastWriteField1(buf []byte) (offset int) {
	if x.MomentId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetMomentId())
	return offset
}

func (x *SearchMomentReq) FastWrite(buf []byte) (offset int) {
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

func (x *SearchMomentReq) fastWriteField1(buf []byte) (offset int) {
	if x.CommunityId == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetCommunityId())
	return offset
}

func (x *SearchMomentReq) fastWriteField2(buf []byte) (offset int) {
	if x.IsParent == nil {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 2, x.GetIsParent())
	return offset
}

func (x *SearchMomentReq) fastWriteField3(buf []byte) (offset int) {
	if x.OnlyUserId == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetOnlyUserId())
	return offset
}

func (x *SearchMomentReq) fastWriteField4(buf []byte) (offset int) {
	if x.Keyword == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetKeyword())
	return offset
}

func (x *SearchMomentReq) fastWriteField5(buf []byte) (offset int) {
	if x.PaginationOption == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 5, x.GetPaginationOption())
	return offset
}

func (x *SearchMomentResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *SearchMomentResp) fastWriteField1(buf []byte) (offset int) {
	if x.Moments == nil {
		return offset
	}
	for i := range x.GetMoments() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetMoments()[i])
	}
	return offset
}

func (x *SearchMomentResp) fastWriteField2(buf []byte) (offset int) {
	if x.Total == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetTotal())
	return offset
}

func (x *Moment) Size() (n int) {
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

func (x *Moment) sizeField1() (n int) {
	if x.Id == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetId())
	return n
}

func (x *Moment) sizeField2() (n int) {
	if x.CreateAt == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetCreateAt())
	return n
}

func (x *Moment) sizeField3() (n int) {
	if x.CatId == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetCatId())
	return n
}

func (x *Moment) sizeField4() (n int) {
	if len(x.Photos) == 0 {
		return n
	}
	for i := range x.GetPhotos() {
		n += fastpb.SizeString(4, x.GetPhotos()[i])
	}
	return n
}

func (x *Moment) sizeField5() (n int) {
	if x.Title == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetTitle())
	return n
}

func (x *Moment) sizeField6() (n int) {
	if x.Text == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetText())
	return n
}

func (x *Moment) sizeField7() (n int) {
	if x.CommunityId == "" {
		return n
	}
	n += fastpb.SizeString(7, x.GetCommunityId())
	return n
}

func (x *Moment) sizeField8() (n int) {
	if x.User == nil {
		return n
	}
	n += fastpb.SizeMessage(8, x.GetUser())
	return n
}

func (x *GetMomentPreviewsReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *GetMomentPreviewsReq) sizeField1() (n int) {
	if x.CommunityId == nil {
		return n
	}
	n += fastpb.SizeString(1, x.GetCommunityId())
	return n
}

func (x *GetMomentPreviewsReq) sizeField2() (n int) {
	if x.IsParent == nil {
		return n
	}
	n += fastpb.SizeBool(2, x.GetIsParent())
	return n
}

func (x *GetMomentPreviewsReq) sizeField3() (n int) {
	if x.OnlyUserId == nil {
		return n
	}
	n += fastpb.SizeString(3, x.GetOnlyUserId())
	return n
}

func (x *GetMomentPreviewsReq) sizeField4() (n int) {
	if x.PaginationOption == nil {
		return n
	}
	n += fastpb.SizeMessage(4, x.GetPaginationOption())
	return n
}

func (x *GetMomentPreviewsResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *GetMomentPreviewsResp) sizeField1() (n int) {
	if x.Moments == nil {
		return n
	}
	for i := range x.GetMoments() {
		n += fastpb.SizeMessage(1, x.GetMoments()[i])
	}
	return n
}

func (x *GetMomentPreviewsResp) sizeField2() (n int) {
	if x.Total == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetTotal())
	return n
}

func (x *GetMomentPreviewsResp) sizeField3() (n int) {
	if x.Token == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetToken())
	return n
}

func (x *GetMomentDetailReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetMomentDetailReq) sizeField1() (n int) {
	if x.MomentId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetMomentId())
	return n
}

func (x *GetMomentDetailResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetMomentDetailResp) sizeField1() (n int) {
	if x.Moment == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetMoment())
	return n
}

func (x *DeleteMomentReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DeleteMomentReq) sizeField1() (n int) {
	if x.MomentId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetMomentId())
	return n
}

func (x *DeleteMomentResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *NewMomentReq) Size() (n int) {
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

func (x *NewMomentReq) sizeField1() (n int) {
	if x.Id == nil {
		return n
	}
	n += fastpb.SizeString(1, x.GetId())
	return n
}

func (x *NewMomentReq) sizeField2() (n int) {
	if x.Title == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetTitle())
	return n
}

func (x *NewMomentReq) sizeField3() (n int) {
	if x.CatId == nil {
		return n
	}
	n += fastpb.SizeString(3, x.GetCatId())
	return n
}

func (x *NewMomentReq) sizeField4() (n int) {
	if x.Text == nil {
		return n
	}
	n += fastpb.SizeString(4, x.GetText())
	return n
}

func (x *NewMomentReq) sizeField5() (n int) {
	if len(x.Photos) == 0 {
		return n
	}
	for i := range x.GetPhotos() {
		n += fastpb.SizeString(5, x.GetPhotos()[i])
	}
	return n
}

func (x *NewMomentReq) sizeField6() (n int) {
	if x.CommunityId == nil {
		return n
	}
	n += fastpb.SizeString(6, x.GetCommunityId())
	return n
}

func (x *NewMomentResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *NewMomentResp) sizeField1() (n int) {
	if x.MomentId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetMomentId())
	return n
}

func (x *SearchMomentReq) Size() (n int) {
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

func (x *SearchMomentReq) sizeField1() (n int) {
	if x.CommunityId == nil {
		return n
	}
	n += fastpb.SizeString(1, x.GetCommunityId())
	return n
}

func (x *SearchMomentReq) sizeField2() (n int) {
	if x.IsParent == nil {
		return n
	}
	n += fastpb.SizeBool(2, x.GetIsParent())
	return n
}

func (x *SearchMomentReq) sizeField3() (n int) {
	if x.OnlyUserId == nil {
		return n
	}
	n += fastpb.SizeString(3, x.GetOnlyUserId())
	return n
}

func (x *SearchMomentReq) sizeField4() (n int) {
	if x.Keyword == nil {
		return n
	}
	n += fastpb.SizeString(4, x.GetKeyword())
	return n
}

func (x *SearchMomentReq) sizeField5() (n int) {
	if x.PaginationOption == nil {
		return n
	}
	n += fastpb.SizeMessage(5, x.GetPaginationOption())
	return n
}

func (x *SearchMomentResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *SearchMomentResp) sizeField1() (n int) {
	if x.Moments == nil {
		return n
	}
	for i := range x.GetMoments() {
		n += fastpb.SizeMessage(1, x.GetMoments()[i])
	}
	return n
}

func (x *SearchMomentResp) sizeField2() (n int) {
	if x.Total == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetTotal())
	return n
}

var fieldIDToName_Moment = map[int32]string{
	1: "Id",
	2: "CreateAt",
	3: "CatId",
	4: "Photos",
	5: "Title",
	6: "Text",
	7: "CommunityId",
	8: "User",
}

var fieldIDToName_GetMomentPreviewsReq = map[int32]string{
	1: "CommunityId",
	2: "IsParent",
	3: "OnlyUserId",
	4: "PaginationOption",
}

var fieldIDToName_GetMomentPreviewsResp = map[int32]string{
	1: "Moments",
	2: "Total",
	3: "Token",
}

var fieldIDToName_GetMomentDetailReq = map[int32]string{
	1: "MomentId",
}

var fieldIDToName_GetMomentDetailResp = map[int32]string{
	1: "Moment",
}

var fieldIDToName_DeleteMomentReq = map[int32]string{
	1: "MomentId",
}

var fieldIDToName_DeleteMomentResp = map[int32]string{}

var fieldIDToName_NewMomentReq = map[int32]string{
	1: "Id",
	2: "Title",
	3: "CatId",
	4: "Text",
	5: "Photos",
	6: "CommunityId",
}

var fieldIDToName_NewMomentResp = map[int32]string{
	1: "MomentId",
}

var fieldIDToName_SearchMomentReq = map[int32]string{
	1: "CommunityId",
	2: "IsParent",
	3: "OnlyUserId",
	4: "Keyword",
	5: "PaginationOption",
}

var fieldIDToName_SearchMomentResp = map[int32]string{
	1: "Moments",
	2: "Total",
}

var _ = http.File_http_http_proto
var _ = basic.File_basic_pagination_proto
var _ = user.File_meowchat_user_common_proto
