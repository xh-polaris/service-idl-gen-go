// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package system

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *Notice) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Notice[number], err)
}

func (x *Notice) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Notice) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.CommunityId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Notice) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Text, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Notice) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.CreateAt, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Notice) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.UpdateAt, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *News) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_News[number], err)
}

func (x *News) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *News) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.CommunityId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *News) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ImageUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *News) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.LinkUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *News) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Type, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Admin) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Admin[number], err)
}

func (x *Admin) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Admin) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.CommunityId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Admin) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Name, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Admin) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Title, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Admin) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Phone, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Admin) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.Wechat, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Admin) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.AvatarUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Community) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Community[number], err)
}

func (x *Community) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Community) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Name, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Community) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ParentId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Role) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Role[number], err)
}

func (x *Role) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.RoleType = RoleType(v)
	return offset, nil
}

func (x *Role) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	tmp, offset, err := fastpb.ReadString(buf, _type)
	x.CommunityId = &tmp
	return offset, err
}

func (x *Apply) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Apply[number], err)
}

func (x *Apply) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ApplyId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Apply) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.ApplicantId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Apply) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.CommunityId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Notice) FastWrite(buf []byte) (offset int) {
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

func (x *Notice) fastWriteField1(buf []byte) (offset int) {
	if x.Id == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetId())
	return offset
}

func (x *Notice) fastWriteField2(buf []byte) (offset int) {
	if x.CommunityId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetCommunityId())
	return offset
}

func (x *Notice) fastWriteField3(buf []byte) (offset int) {
	if x.Text == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetText())
	return offset
}

func (x *Notice) fastWriteField4(buf []byte) (offset int) {
	if x.CreateAt == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.GetCreateAt())
	return offset
}

func (x *Notice) fastWriteField5(buf []byte) (offset int) {
	if x.UpdateAt == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 5, x.GetUpdateAt())
	return offset
}

func (x *News) FastWrite(buf []byte) (offset int) {
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

func (x *News) fastWriteField1(buf []byte) (offset int) {
	if x.Id == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetId())
	return offset
}

func (x *News) fastWriteField2(buf []byte) (offset int) {
	if x.CommunityId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetCommunityId())
	return offset
}

func (x *News) fastWriteField3(buf []byte) (offset int) {
	if x.ImageUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetImageUrl())
	return offset
}

func (x *News) fastWriteField4(buf []byte) (offset int) {
	if x.LinkUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetLinkUrl())
	return offset
}

func (x *News) fastWriteField5(buf []byte) (offset int) {
	if x.Type == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetType())
	return offset
}

func (x *Admin) FastWrite(buf []byte) (offset int) {
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

func (x *Admin) fastWriteField1(buf []byte) (offset int) {
	if x.Id == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetId())
	return offset
}

func (x *Admin) fastWriteField2(buf []byte) (offset int) {
	if x.CommunityId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetCommunityId())
	return offset
}

func (x *Admin) fastWriteField3(buf []byte) (offset int) {
	if x.Name == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetName())
	return offset
}

func (x *Admin) fastWriteField4(buf []byte) (offset int) {
	if x.Title == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetTitle())
	return offset
}

func (x *Admin) fastWriteField5(buf []byte) (offset int) {
	if x.Phone == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetPhone())
	return offset
}

func (x *Admin) fastWriteField6(buf []byte) (offset int) {
	if x.Wechat == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetWechat())
	return offset
}

func (x *Admin) fastWriteField7(buf []byte) (offset int) {
	if x.AvatarUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 7, x.GetAvatarUrl())
	return offset
}

func (x *Community) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *Community) fastWriteField1(buf []byte) (offset int) {
	if x.Id == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetId())
	return offset
}

func (x *Community) fastWriteField2(buf []byte) (offset int) {
	if x.Name == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetName())
	return offset
}

func (x *Community) fastWriteField3(buf []byte) (offset int) {
	if x.ParentId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetParentId())
	return offset
}

func (x *Role) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *Role) fastWriteField1(buf []byte) (offset int) {
	if x.RoleType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, int32(x.GetRoleType()))
	return offset
}

func (x *Role) fastWriteField2(buf []byte) (offset int) {
	if x.CommunityId == nil {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetCommunityId())
	return offset
}

func (x *Apply) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *Apply) fastWriteField1(buf []byte) (offset int) {
	if x.ApplyId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetApplyId())
	return offset
}

func (x *Apply) fastWriteField2(buf []byte) (offset int) {
	if x.ApplicantId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetApplicantId())
	return offset
}

func (x *Apply) fastWriteField4(buf []byte) (offset int) {
	if x.CommunityId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetCommunityId())
	return offset
}

func (x *Notice) Size() (n int) {
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

func (x *Notice) sizeField1() (n int) {
	if x.Id == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetId())
	return n
}

func (x *Notice) sizeField2() (n int) {
	if x.CommunityId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetCommunityId())
	return n
}

func (x *Notice) sizeField3() (n int) {
	if x.Text == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetText())
	return n
}

func (x *Notice) sizeField4() (n int) {
	if x.CreateAt == 0 {
		return n
	}
	n += fastpb.SizeInt64(4, x.GetCreateAt())
	return n
}

func (x *Notice) sizeField5() (n int) {
	if x.UpdateAt == 0 {
		return n
	}
	n += fastpb.SizeInt64(5, x.GetUpdateAt())
	return n
}

func (x *News) Size() (n int) {
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

func (x *News) sizeField1() (n int) {
	if x.Id == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetId())
	return n
}

func (x *News) sizeField2() (n int) {
	if x.CommunityId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetCommunityId())
	return n
}

func (x *News) sizeField3() (n int) {
	if x.ImageUrl == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetImageUrl())
	return n
}

func (x *News) sizeField4() (n int) {
	if x.LinkUrl == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetLinkUrl())
	return n
}

func (x *News) sizeField5() (n int) {
	if x.Type == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetType())
	return n
}

func (x *Admin) Size() (n int) {
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

func (x *Admin) sizeField1() (n int) {
	if x.Id == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetId())
	return n
}

func (x *Admin) sizeField2() (n int) {
	if x.CommunityId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetCommunityId())
	return n
}

func (x *Admin) sizeField3() (n int) {
	if x.Name == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetName())
	return n
}

func (x *Admin) sizeField4() (n int) {
	if x.Title == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetTitle())
	return n
}

func (x *Admin) sizeField5() (n int) {
	if x.Phone == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetPhone())
	return n
}

func (x *Admin) sizeField6() (n int) {
	if x.Wechat == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetWechat())
	return n
}

func (x *Admin) sizeField7() (n int) {
	if x.AvatarUrl == "" {
		return n
	}
	n += fastpb.SizeString(7, x.GetAvatarUrl())
	return n
}

func (x *Community) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *Community) sizeField1() (n int) {
	if x.Id == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetId())
	return n
}

func (x *Community) sizeField2() (n int) {
	if x.Name == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetName())
	return n
}

func (x *Community) sizeField3() (n int) {
	if x.ParentId == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetParentId())
	return n
}

func (x *Role) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *Role) sizeField1() (n int) {
	if x.RoleType == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, int32(x.GetRoleType()))
	return n
}

func (x *Role) sizeField2() (n int) {
	if x.CommunityId == nil {
		return n
	}
	n += fastpb.SizeString(2, x.GetCommunityId())
	return n
}

func (x *Apply) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField4()
	return n
}

func (x *Apply) sizeField1() (n int) {
	if x.ApplyId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetApplyId())
	return n
}

func (x *Apply) sizeField2() (n int) {
	if x.ApplicantId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetApplicantId())
	return n
}

func (x *Apply) sizeField4() (n int) {
	if x.CommunityId == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetCommunityId())
	return n
}

var fieldIDToName_Notice = map[int32]string{
	1: "Id",
	2: "CommunityId",
	3: "Text",
	4: "CreateAt",
	5: "UpdateAt",
}

var fieldIDToName_News = map[int32]string{
	1: "Id",
	2: "CommunityId",
	3: "ImageUrl",
	4: "LinkUrl",
	5: "Type",
}

var fieldIDToName_Admin = map[int32]string{
	1: "Id",
	2: "CommunityId",
	3: "Name",
	4: "Title",
	5: "Phone",
	6: "Wechat",
	7: "AvatarUrl",
}

var fieldIDToName_Community = map[int32]string{
	1: "Id",
	2: "Name",
	3: "ParentId",
}

var fieldIDToName_Role = map[int32]string{
	1: "RoleType",
	2: "CommunityId",
}

var fieldIDToName_Apply = map[int32]string{
	1: "ApplyId",
	2: "ApplicantId",
	4: "CommunityId",
}
