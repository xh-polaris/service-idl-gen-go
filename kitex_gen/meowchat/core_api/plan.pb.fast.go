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

func (x *Plan) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 10:
		offset, err = x.fastReadField10(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Plan[number], err)
}

func (x *Plan) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Plan) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Name, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Plan) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Description, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Plan) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.PlanType, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Plan) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.StartTime, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Plan) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.Duration, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Plan) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.CatId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Plan) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	var v user.UserPreview
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.User = append(x.User, &v)
	return offset, nil
}

func (x *Plan) fastReadField9(buf []byte, _type int8) (offset int, err error) {
	var v string
	v, offset, err = fastpb.ReadString(buf, _type)
	if err != nil {
		return offset, err
	}
	x.ImageUrl = append(x.ImageUrl, v)
	return offset, err
}

func (x *Plan) fastReadField10(buf []byte, _type int8) (offset int, err error) {
	x.CreateAt, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetPlanPreviewsReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetPlanPreviewsReq[number], err)
}

func (x *GetPlanPreviewsReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.CatId = &tmp
	return offset, err
}

func (x *GetPlanPreviewsReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.OnlyUserId = &tmp
	return offset, err
}

func (x *GetPlanPreviewsReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v basic.PaginationOptions
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.PaginationOption = &v
	return offset, nil
}

func (x *GetPlanPreviewsResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetPlanPreviewsResp[number], err)
}

func (x *GetPlanPreviewsResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Plan
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Plans = append(x.Plans, &v)
	return offset, nil
}

func (x *GetPlanPreviewsResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Total, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetPlanPreviewsResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Token, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetPlanDetailReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetPlanDetailReq[number], err)
}

func (x *GetPlanDetailReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.PlanId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetPlanDetailResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetPlanDetailResp[number], err)
}

func (x *GetPlanDetailResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Plan
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Plan = &v
	return offset, nil
}

func (x *DeletePlanReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeletePlanReq[number], err)
}

func (x *DeletePlanReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.PlanId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DeletePlanResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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

func (x *NewPlanReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_NewPlanReq[number], err)
}

func (x *NewPlanReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Id = &tmp
	return offset, err
}

func (x *NewPlanReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Name = &tmp
	return offset, err
}

func (x *NewPlanReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Description = &tmp
	return offset, err
}

func (x *NewPlanReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.PlanType = &tmp
	return offset, err
}

func (x *NewPlanReq) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.StartTime = &tmp
	return offset, err
}

func (x *NewPlanReq) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Duration = &tmp
	return offset, err
}

func (x *NewPlanReq) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.CatId = &tmp
	return offset, err
}

func (x *NewPlanReq) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	var v string
	v, offset, err = fastpb.ReadString(buf, _type)
	if err != nil {
		return offset, err
	}
	x.ImageUrl = append(x.ImageUrl, v)
	return offset, err
}

func (x *NewPlanResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_NewPlanResp[number], err)
}

func (x *NewPlanResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.PlanId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SearchPlanReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SearchPlanReq[number], err)
}

func (x *SearchPlanReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.CatId = &tmp
	return offset, err
}

func (x *SearchPlanReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.OnlyUserId = &tmp
	return offset, err
}

func (x *SearchPlanReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Keyword = &tmp
	return offset, err
}

func (x *SearchPlanReq) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	var v basic.PaginationOptions
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.PaginationOption = &v
	return offset, nil
}

func (x *SearchPlanResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SearchPlanResp[number], err)
}

func (x *SearchPlanResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Plan
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Plans = append(x.Plans, &v)
	return offset, nil
}

func (x *SearchPlanResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Total, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Plan) FastWrite(buf []byte) (offset int) {
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
	offset += x.fastWriteField10(buf[offset:])
	return offset
}

func (x *Plan) fastWriteField1(buf []byte) (offset int) {
	if x.Id == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetId())
	return offset
}

func (x *Plan) fastWriteField2(buf []byte) (offset int) {
	if x.Name == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetName())
	return offset
}

func (x *Plan) fastWriteField3(buf []byte) (offset int) {
	if x.Description == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetDescription())
	return offset
}

func (x *Plan) fastWriteField4(buf []byte) (offset int) {
	if x.PlanType == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetPlanType())
	return offset
}

func (x *Plan) fastWriteField5(buf []byte) (offset int) {
	if x.StartTime == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetStartTime())
	return offset
}

func (x *Plan) fastWriteField6(buf []byte) (offset int) {
	if x.Duration == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetDuration())
	return offset
}

func (x *Plan) fastWriteField7(buf []byte) (offset int) {
	if x.CatId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 7, x.GetCatId())
	return offset
}

func (x *Plan) fastWriteField8(buf []byte) (offset int) {
	if x.User == nil {
		return offset
	}
	for i := range x.GetUser() {
		offset += fastpb.WriteMessage(buf[offset:], 8, x.GetUser()[i])
	}
	return offset
}

func (x *Plan) fastWriteField9(buf []byte) (offset int) {
	if len(x.ImageUrl) == 0 {
		return offset
	}
	for i := range x.GetImageUrl() {
		offset += fastpb.WriteString(buf[offset:], 9, x.GetImageUrl()[i])
	}
	return offset
}

func (x *Plan) fastWriteField10(buf []byte) (offset int) {
	if x.CreateAt == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 10, x.GetCreateAt())
	return offset
}

func (x *GetPlanPreviewsReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *GetPlanPreviewsReq) fastWriteField1(buf []byte) (offset int) {
	if x.CatId == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetCatId())
	return offset
}

func (x *GetPlanPreviewsReq) fastWriteField2(buf []byte) (offset int) {
	if x.OnlyUserId == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetOnlyUserId())
	return offset
}

func (x *GetPlanPreviewsReq) fastWriteField3(buf []byte) (offset int) {
	if x.PaginationOption == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 3, x.GetPaginationOption())
	return offset
}

func (x *GetPlanPreviewsResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *GetPlanPreviewsResp) fastWriteField1(buf []byte) (offset int) {
	if x.Plans == nil {
		return offset
	}
	for i := range x.GetPlans() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetPlans()[i])
	}
	return offset
}

func (x *GetPlanPreviewsResp) fastWriteField2(buf []byte) (offset int) {
	if x.Total == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetTotal())
	return offset
}

func (x *GetPlanPreviewsResp) fastWriteField3(buf []byte) (offset int) {
	if x.Token == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetToken())
	return offset
}

func (x *GetPlanDetailReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetPlanDetailReq) fastWriteField1(buf []byte) (offset int) {
	if x.PlanId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetPlanId())
	return offset
}

func (x *GetPlanDetailResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetPlanDetailResp) fastWriteField1(buf []byte) (offset int) {
	if x.Plan == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetPlan())
	return offset
}

func (x *DeletePlanReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DeletePlanReq) fastWriteField1(buf []byte) (offset int) {
	if x.PlanId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetPlanId())
	return offset
}

func (x *DeletePlanResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *NewPlanReq) FastWrite(buf []byte) (offset int) {
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

func (x *NewPlanReq) fastWriteField1(buf []byte) (offset int) {
	if x.Id == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetId())
	return offset
}

func (x *NewPlanReq) fastWriteField2(buf []byte) (offset int) {
	if x.Name == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetName())
	return offset
}

func (x *NewPlanReq) fastWriteField3(buf []byte) (offset int) {
	if x.Description == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetDescription())
	return offset
}

func (x *NewPlanReq) fastWriteField4(buf []byte) (offset int) {
	if x.PlanType == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetPlanType())
	return offset
}

func (x *NewPlanReq) fastWriteField5(buf []byte) (offset int) {
	if x.StartTime == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetStartTime())
	return offset
}

func (x *NewPlanReq) fastWriteField6(buf []byte) (offset int) {
	if x.Duration == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetDuration())
	return offset
}

func (x *NewPlanReq) fastWriteField7(buf []byte) (offset int) {
	if x.CatId == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 7, x.GetCatId())
	return offset
}

func (x *NewPlanReq) fastWriteField8(buf []byte) (offset int) {
	if len(x.ImageUrl) == 0 {
		return offset
	}
	for i := range x.GetImageUrl() {
		offset += fastpb.WriteString(buf[offset:], 8, x.GetImageUrl()[i])
	}
	return offset
}

func (x *NewPlanResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *NewPlanResp) fastWriteField1(buf []byte) (offset int) {
	if x.PlanId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetPlanId())
	return offset
}

func (x *SearchPlanReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	return offset
}

func (x *SearchPlanReq) fastWriteField1(buf []byte) (offset int) {
	if x.CatId == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetCatId())
	return offset
}

func (x *SearchPlanReq) fastWriteField3(buf []byte) (offset int) {
	if x.OnlyUserId == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetOnlyUserId())
	return offset
}

func (x *SearchPlanReq) fastWriteField4(buf []byte) (offset int) {
	if x.Keyword == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetKeyword())
	return offset
}

func (x *SearchPlanReq) fastWriteField5(buf []byte) (offset int) {
	if x.PaginationOption == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 5, x.GetPaginationOption())
	return offset
}

func (x *SearchPlanResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *SearchPlanResp) fastWriteField1(buf []byte) (offset int) {
	if x.Plans == nil {
		return offset
	}
	for i := range x.GetPlans() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetPlans()[i])
	}
	return offset
}

func (x *SearchPlanResp) fastWriteField2(buf []byte) (offset int) {
	if x.Total == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetTotal())
	return offset
}

func (x *Plan) Size() (n int) {
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
	n += x.sizeField10()
	return n
}

func (x *Plan) sizeField1() (n int) {
	if x.Id == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetId())
	return n
}

func (x *Plan) sizeField2() (n int) {
	if x.Name == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetName())
	return n
}

func (x *Plan) sizeField3() (n int) {
	if x.Description == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetDescription())
	return n
}

func (x *Plan) sizeField4() (n int) {
	if x.PlanType == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetPlanType())
	return n
}

func (x *Plan) sizeField5() (n int) {
	if x.StartTime == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetStartTime())
	return n
}

func (x *Plan) sizeField6() (n int) {
	if x.Duration == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetDuration())
	return n
}

func (x *Plan) sizeField7() (n int) {
	if x.CatId == "" {
		return n
	}
	n += fastpb.SizeString(7, x.GetCatId())
	return n
}

func (x *Plan) sizeField8() (n int) {
	if x.User == nil {
		return n
	}
	for i := range x.GetUser() {
		n += fastpb.SizeMessage(8, x.GetUser()[i])
	}
	return n
}

func (x *Plan) sizeField9() (n int) {
	if len(x.ImageUrl) == 0 {
		return n
	}
	for i := range x.GetImageUrl() {
		n += fastpb.SizeString(9, x.GetImageUrl()[i])
	}
	return n
}

func (x *Plan) sizeField10() (n int) {
	if x.CreateAt == 0 {
		return n
	}
	n += fastpb.SizeInt64(10, x.GetCreateAt())
	return n
}

func (x *GetPlanPreviewsReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *GetPlanPreviewsReq) sizeField1() (n int) {
	if x.CatId == nil {
		return n
	}
	n += fastpb.SizeString(1, x.GetCatId())
	return n
}

func (x *GetPlanPreviewsReq) sizeField2() (n int) {
	if x.OnlyUserId == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetOnlyUserId())
	return n
}

func (x *GetPlanPreviewsReq) sizeField3() (n int) {
	if x.PaginationOption == nil {
		return n
	}
	n += fastpb.SizeMessage(3, x.GetPaginationOption())
	return n
}

func (x *GetPlanPreviewsResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *GetPlanPreviewsResp) sizeField1() (n int) {
	if x.Plans == nil {
		return n
	}
	for i := range x.GetPlans() {
		n += fastpb.SizeMessage(1, x.GetPlans()[i])
	}
	return n
}

func (x *GetPlanPreviewsResp) sizeField2() (n int) {
	if x.Total == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetTotal())
	return n
}

func (x *GetPlanPreviewsResp) sizeField3() (n int) {
	if x.Token == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetToken())
	return n
}

func (x *GetPlanDetailReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetPlanDetailReq) sizeField1() (n int) {
	if x.PlanId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetPlanId())
	return n
}

func (x *GetPlanDetailResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetPlanDetailResp) sizeField1() (n int) {
	if x.Plan == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetPlan())
	return n
}

func (x *DeletePlanReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DeletePlanReq) sizeField1() (n int) {
	if x.PlanId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetPlanId())
	return n
}

func (x *DeletePlanResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *NewPlanReq) Size() (n int) {
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

func (x *NewPlanReq) sizeField1() (n int) {
	if x.Id == nil {
		return n
	}
	n += fastpb.SizeString(1, x.GetId())
	return n
}

func (x *NewPlanReq) sizeField2() (n int) {
	if x.Name == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetName())
	return n
}

func (x *NewPlanReq) sizeField3() (n int) {
	if x.Description == nil {
		return n
	}
	n += fastpb.SizeString(3, x.GetDescription())
	return n
}

func (x *NewPlanReq) sizeField4() (n int) {
	if x.PlanType == nil {
		return n
	}
	n += fastpb.SizeString(4, x.GetPlanType())
	return n
}

func (x *NewPlanReq) sizeField5() (n int) {
	if x.StartTime == nil {
		return n
	}
	n += fastpb.SizeString(5, x.GetStartTime())
	return n
}

func (x *NewPlanReq) sizeField6() (n int) {
	if x.Duration == nil {
		return n
	}
	n += fastpb.SizeString(6, x.GetDuration())
	return n
}

func (x *NewPlanReq) sizeField7() (n int) {
	if x.CatId == nil {
		return n
	}
	n += fastpb.SizeString(7, x.GetCatId())
	return n
}

func (x *NewPlanReq) sizeField8() (n int) {
	if len(x.ImageUrl) == 0 {
		return n
	}
	for i := range x.GetImageUrl() {
		n += fastpb.SizeString(8, x.GetImageUrl()[i])
	}
	return n
}

func (x *NewPlanResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *NewPlanResp) sizeField1() (n int) {
	if x.PlanId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetPlanId())
	return n
}

func (x *SearchPlanReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	return n
}

func (x *SearchPlanReq) sizeField1() (n int) {
	if x.CatId == nil {
		return n
	}
	n += fastpb.SizeString(1, x.GetCatId())
	return n
}

func (x *SearchPlanReq) sizeField3() (n int) {
	if x.OnlyUserId == nil {
		return n
	}
	n += fastpb.SizeString(3, x.GetOnlyUserId())
	return n
}

func (x *SearchPlanReq) sizeField4() (n int) {
	if x.Keyword == nil {
		return n
	}
	n += fastpb.SizeString(4, x.GetKeyword())
	return n
}

func (x *SearchPlanReq) sizeField5() (n int) {
	if x.PaginationOption == nil {
		return n
	}
	n += fastpb.SizeMessage(5, x.GetPaginationOption())
	return n
}

func (x *SearchPlanResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *SearchPlanResp) sizeField1() (n int) {
	if x.Plans == nil {
		return n
	}
	for i := range x.GetPlans() {
		n += fastpb.SizeMessage(1, x.GetPlans()[i])
	}
	return n
}

func (x *SearchPlanResp) sizeField2() (n int) {
	if x.Total == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetTotal())
	return n
}

var fieldIDToName_Plan = map[int32]string{
	1:  "Id",
	2:  "Name",
	3:  "Description",
	4:  "PlanType",
	5:  "StartTime",
	6:  "Duration",
	7:  "CatId",
	8:  "User",
	9:  "ImageUrl",
	10: "CreateAt",
}

var fieldIDToName_GetPlanPreviewsReq = map[int32]string{
	1: "CatId",
	2: "OnlyUserId",
	3: "PaginationOption",
}

var fieldIDToName_GetPlanPreviewsResp = map[int32]string{
	1: "Plans",
	2: "Total",
	3: "Token",
}

var fieldIDToName_GetPlanDetailReq = map[int32]string{
	1: "PlanId",
}

var fieldIDToName_GetPlanDetailResp = map[int32]string{
	1: "Plan",
}

var fieldIDToName_DeletePlanReq = map[int32]string{
	1: "PlanId",
}

var fieldIDToName_DeletePlanResp = map[int32]string{}

var fieldIDToName_NewPlanReq = map[int32]string{
	1: "Id",
	2: "Name",
	3: "Description",
	4: "PlanType",
	5: "StartTime",
	6: "Duration",
	7: "CatId",
	8: "ImageUrl",
}

var fieldIDToName_NewPlanResp = map[int32]string{
	1: "PlanId",
}

var fieldIDToName_SearchPlanReq = map[int32]string{
	1: "CatId",
	3: "OnlyUserId",
	4: "Keyword",
	5: "PaginationOption",
}

var fieldIDToName_SearchPlanResp = map[int32]string{
	1: "Plans",
	2: "Total",
}

var _ = http.File_http_http_proto
var _ = basic.File_basic_pagination_proto
var _ = user.File_meowchat_user_common_proto