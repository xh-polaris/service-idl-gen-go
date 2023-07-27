// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package moment

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
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
	x.UserId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SearchField) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SearchField[number], err)
}

func (x *SearchField) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Text = &tmp
	return offset, err
}

func (x *SearchField) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.Title = &tmp
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
	var ov SearchOptions_AllFieldsKey
	x.Type = &ov
	ov.AllFieldsKey, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SearchOptions) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var ov SearchOptions_MultiFieldsKey
	x.Type = &ov
	var v SearchField
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	ov.MultiFieldsKey = &v
	return offset, nil
}

func (x *FilterOptions) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_FilterOptions[number], err)
}

func (x *FilterOptions) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.OnlyUserId = &tmp
	return offset, err
}

func (x *FilterOptions) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.OnlyCommunityId = &tmp
	return offset, err
}

func (x *FilterOptions) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v string
	v, offset, err = fastpb.ReadString(buf, _type)
	if err != nil {
		return offset, err
	}
	x.OnlyCommunityIds = append(x.OnlyCommunityIds, v)
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
	if x.UserId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 8, x.GetUserId())
	return offset
}

func (x *SearchField) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *SearchField) fastWriteField1(buf []byte) (offset int) {
	if x.Text == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetText())
	return offset
}

func (x *SearchField) fastWriteField2(buf []byte) (offset int) {
	if x.Title == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetTitle())
	return offset
}

func (x *SearchOptions) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *SearchOptions) fastWriteField1(buf []byte) (offset int) {
	if x.GetAllFieldsKey() == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetAllFieldsKey())
	return offset
}

func (x *SearchOptions) fastWriteField2(buf []byte) (offset int) {
	if x.GetMultiFieldsKey() == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetMultiFieldsKey())
	return offset
}

func (x *FilterOptions) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *FilterOptions) fastWriteField1(buf []byte) (offset int) {
	if x.OnlyUserId == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetOnlyUserId())
	return offset
}

func (x *FilterOptions) fastWriteField2(buf []byte) (offset int) {
	if x.OnlyCommunityId == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetOnlyCommunityId())
	return offset
}

func (x *FilterOptions) fastWriteField3(buf []byte) (offset int) {
	if len(x.OnlyCommunityIds) == 0 {
		return offset
	}
	for i := range x.GetOnlyCommunityIds() {
		offset += fastpb.WriteString(buf[offset:], 3, x.GetOnlyCommunityIds()[i])
	}
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
	if x.UserId == "" {
		return n
	}
	n += fastpb.SizeString(8, x.GetUserId())
	return n
}

func (x *SearchField) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *SearchField) sizeField1() (n int) {
	if x.Text == nil {
		return n
	}
	n += fastpb.SizeString(1, x.GetText())
	return n
}

func (x *SearchField) sizeField2() (n int) {
	if x.Title == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetTitle())
	return n
}

func (x *SearchOptions) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *SearchOptions) sizeField1() (n int) {
	if x.GetAllFieldsKey() == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetAllFieldsKey())
	return n
}

func (x *SearchOptions) sizeField2() (n int) {
	if x.GetMultiFieldsKey() == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetMultiFieldsKey())
	return n
}

func (x *FilterOptions) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *FilterOptions) sizeField1() (n int) {
	if x.OnlyUserId == nil {
		return n
	}
	n += fastpb.SizeString(1, x.GetOnlyUserId())
	return n
}

func (x *FilterOptions) sizeField2() (n int) {
	if x.OnlyCommunityId == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetOnlyCommunityId())
	return n
}

func (x *FilterOptions) sizeField3() (n int) {
	if len(x.OnlyCommunityIds) == 0 {
		return n
	}
	for i := range x.GetOnlyCommunityIds() {
		n += fastpb.SizeString(3, x.GetOnlyCommunityIds()[i])
	}
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
	8: "UserId",
}

var fieldIDToName_SearchField = map[int32]string{
	1: "Text",
	2: "Title",
}

var fieldIDToName_SearchOptions = map[int32]string{
	1: "AllFieldsKey",
	2: "MultiFieldsKey",
}

var fieldIDToName_FilterOptions = map[int32]string{
	1: "OnlyUserId",
	2: "OnlyCommunityId",
	3: "OnlyCommunityIds",
}
