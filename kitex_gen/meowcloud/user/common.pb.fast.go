// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package user

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *UserPreview) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UserPreview[number], err)
}

func (x *UserPreview) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UserPreview) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Nickname, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UserPreview) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Avatar, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *User) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 11:
		offset, err = x.fastReadField11(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 12:
		offset, err = x.fastReadField12(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 13:
		offset, err = x.fastReadField13(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 14:
		offset, err = x.fastReadField14(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 15:
		offset, err = x.fastReadField15(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 16:
		offset, err = x.fastReadField16(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 17:
		offset, err = x.fastReadField17(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_User[number], err)
}

func (x *User) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *User) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Nickname, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *User) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Bio, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *User) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Avatar, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *User) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.UpdatedAt, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *User) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.CreatedAt, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *User) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.DeletedAt, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *User) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	var v User_Membership
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Membership = &v
	return offset, nil
}

func (x *User) fastReadField9(buf []byte, _type int8) (offset int, err error) {
	x.TeamCount, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *User) fastReadField10(buf []byte, _type int8) (offset int, err error) {
	var v string
	v, offset, err = fastpb.ReadString(buf, _type)
	if err != nil {
		return offset, err
	}
	x.TeamIds = append(x.TeamIds, v)
	return offset, err
}

func (x *User) fastReadField11(buf []byte, _type int8) (offset int, err error) {
	x.MyAlbumCount, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *User) fastReadField12(buf []byte, _type int8) (offset int, err error) {
	offset, err = fastpb.ReadList(buf, _type,
		func(buf []byte, _type int8) (n int, err error) {
			var v int32
			v, offset, err = fastpb.ReadInt32(buf, _type)
			if err != nil {
				return offset, err
			}
			x.MyAlbumIds = append(x.MyAlbumIds, v)
			return offset, err
		})
	return offset, err
}

func (x *User) fastReadField13(buf []byte, _type int8) (offset int, err error) {
	x.FollowedAlbumCount, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *User) fastReadField14(buf []byte, _type int8) (offset int, err error) {
	offset, err = fastpb.ReadList(buf, _type,
		func(buf []byte, _type int8) (n int, err error) {
			var v int32
			v, offset, err = fastpb.ReadInt32(buf, _type)
			if err != nil {
				return offset, err
			}
			x.FollowedAlbumIds = append(x.FollowedAlbumIds, v)
			return offset, err
		})
	return offset, err
}

func (x *User) fastReadField15(buf []byte, _type int8) (offset int, err error) {
	var v User_StorageInfo
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.StorageInfo = &v
	return offset, nil
}

func (x *User) fastReadField16(buf []byte, _type int8) (offset int, err error) {
	x.Points, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *User) fastReadField17(buf []byte, _type int8) (offset int, err error) {
	var v string
	v, offset, err = fastpb.ReadString(buf, _type)
	if err != nil {
		return offset, err
	}
	x.Achievements = append(x.Achievements, v)
	return offset, err
}

func (x *User_Membership) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_User_Membership[number], err)
}

func (x *User_Membership) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.MemberId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *User_Membership) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.MemberLevel, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *User_StorageInfo) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_User_StorageInfo[number], err)
}

func (x *User_StorageInfo) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.AvailablePhotos, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *User_StorageInfo) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UsedPhotos, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *User_StorageInfo) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.AvailableMemory, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *User_StorageInfo) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.UsedMemory, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *User_StorageInfo) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.AvailableAlbums, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *User_StorageInfo) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.UsedAlbums, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *UserPreview) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *UserPreview) fastWriteField1(buf []byte) (offset int) {
	if x.Id == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetId())
	return offset
}

func (x *UserPreview) fastWriteField2(buf []byte) (offset int) {
	if x.Nickname == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetNickname())
	return offset
}

func (x *UserPreview) fastWriteField3(buf []byte) (offset int) {
	if x.Avatar == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetAvatar())
	return offset
}

func (x *User) FastWrite(buf []byte) (offset int) {
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
	offset += x.fastWriteField11(buf[offset:])
	offset += x.fastWriteField12(buf[offset:])
	offset += x.fastWriteField13(buf[offset:])
	offset += x.fastWriteField14(buf[offset:])
	offset += x.fastWriteField15(buf[offset:])
	offset += x.fastWriteField16(buf[offset:])
	offset += x.fastWriteField17(buf[offset:])
	return offset
}

func (x *User) fastWriteField1(buf []byte) (offset int) {
	if x.Id == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetId())
	return offset
}

func (x *User) fastWriteField2(buf []byte) (offset int) {
	if x.Nickname == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetNickname())
	return offset
}

func (x *User) fastWriteField3(buf []byte) (offset int) {
	if x.Bio == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetBio())
	return offset
}

func (x *User) fastWriteField4(buf []byte) (offset int) {
	if x.Avatar == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetAvatar())
	return offset
}

func (x *User) fastWriteField5(buf []byte) (offset int) {
	if x.UpdatedAt == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 5, x.GetUpdatedAt())
	return offset
}

func (x *User) fastWriteField6(buf []byte) (offset int) {
	if x.CreatedAt == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 6, x.GetCreatedAt())
	return offset
}

func (x *User) fastWriteField7(buf []byte) (offset int) {
	if x.DeletedAt == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 7, x.GetDeletedAt())
	return offset
}

func (x *User) fastWriteField8(buf []byte) (offset int) {
	if x.Membership == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 8, x.GetMembership())
	return offset
}

func (x *User) fastWriteField9(buf []byte) (offset int) {
	if x.TeamCount == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 9, x.GetTeamCount())
	return offset
}

func (x *User) fastWriteField10(buf []byte) (offset int) {
	if len(x.TeamIds) == 0 {
		return offset
	}
	for i := range x.GetTeamIds() {
		offset += fastpb.WriteString(buf[offset:], 10, x.GetTeamIds()[i])
	}
	return offset
}

func (x *User) fastWriteField11(buf []byte) (offset int) {
	if x.MyAlbumCount == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 11, x.GetMyAlbumCount())
	return offset
}

func (x *User) fastWriteField12(buf []byte) (offset int) {
	if len(x.MyAlbumIds) == 0 {
		return offset
	}
	offset += fastpb.WriteListPacked(buf[offset:], 12, len(x.GetMyAlbumIds()),
		func(buf []byte, numTagOrKey, numIdxOrVal int32) int {
			offset := 0
			offset += fastpb.WriteInt32(buf[offset:], numTagOrKey, x.GetMyAlbumIds()[numIdxOrVal])
			return offset
		})
	return offset
}

func (x *User) fastWriteField13(buf []byte) (offset int) {
	if x.FollowedAlbumCount == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 13, x.GetFollowedAlbumCount())
	return offset
}

func (x *User) fastWriteField14(buf []byte) (offset int) {
	if len(x.FollowedAlbumIds) == 0 {
		return offset
	}
	offset += fastpb.WriteListPacked(buf[offset:], 14, len(x.GetFollowedAlbumIds()),
		func(buf []byte, numTagOrKey, numIdxOrVal int32) int {
			offset := 0
			offset += fastpb.WriteInt32(buf[offset:], numTagOrKey, x.GetFollowedAlbumIds()[numIdxOrVal])
			return offset
		})
	return offset
}

func (x *User) fastWriteField15(buf []byte) (offset int) {
	if x.StorageInfo == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 15, x.GetStorageInfo())
	return offset
}

func (x *User) fastWriteField16(buf []byte) (offset int) {
	if x.Points == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 16, x.GetPoints())
	return offset
}

func (x *User) fastWriteField17(buf []byte) (offset int) {
	if len(x.Achievements) == 0 {
		return offset
	}
	for i := range x.GetAchievements() {
		offset += fastpb.WriteString(buf[offset:], 17, x.GetAchievements()[i])
	}
	return offset
}

func (x *User_Membership) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *User_Membership) fastWriteField1(buf []byte) (offset int) {
	if x.MemberId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetMemberId())
	return offset
}

func (x *User_Membership) fastWriteField2(buf []byte) (offset int) {
	if x.MemberLevel == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetMemberLevel())
	return offset
}

func (x *User_StorageInfo) FastWrite(buf []byte) (offset int) {
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

func (x *User_StorageInfo) fastWriteField1(buf []byte) (offset int) {
	if x.AvailablePhotos == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetAvailablePhotos())
	return offset
}

func (x *User_StorageInfo) fastWriteField2(buf []byte) (offset int) {
	if x.UsedPhotos == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetUsedPhotos())
	return offset
}

func (x *User_StorageInfo) fastWriteField3(buf []byte) (offset int) {
	if x.AvailableMemory == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetAvailableMemory())
	return offset
}

func (x *User_StorageInfo) fastWriteField4(buf []byte) (offset int) {
	if x.UsedMemory == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.GetUsedMemory())
	return offset
}

func (x *User_StorageInfo) fastWriteField5(buf []byte) (offset int) {
	if x.AvailableAlbums == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 5, x.GetAvailableAlbums())
	return offset
}

func (x *User_StorageInfo) fastWriteField6(buf []byte) (offset int) {
	if x.UsedAlbums == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 6, x.GetUsedAlbums())
	return offset
}

func (x *UserPreview) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *UserPreview) sizeField1() (n int) {
	if x.Id == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetId())
	return n
}

func (x *UserPreview) sizeField2() (n int) {
	if x.Nickname == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetNickname())
	return n
}

func (x *UserPreview) sizeField3() (n int) {
	if x.Avatar == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetAvatar())
	return n
}

func (x *User) Size() (n int) {
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
	n += x.sizeField11()
	n += x.sizeField12()
	n += x.sizeField13()
	n += x.sizeField14()
	n += x.sizeField15()
	n += x.sizeField16()
	n += x.sizeField17()
	return n
}

func (x *User) sizeField1() (n int) {
	if x.Id == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetId())
	return n
}

func (x *User) sizeField2() (n int) {
	if x.Nickname == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetNickname())
	return n
}

func (x *User) sizeField3() (n int) {
	if x.Bio == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetBio())
	return n
}

func (x *User) sizeField4() (n int) {
	if x.Avatar == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetAvatar())
	return n
}

func (x *User) sizeField5() (n int) {
	if x.UpdatedAt == 0 {
		return n
	}
	n += fastpb.SizeInt64(5, x.GetUpdatedAt())
	return n
}

func (x *User) sizeField6() (n int) {
	if x.CreatedAt == 0 {
		return n
	}
	n += fastpb.SizeInt64(6, x.GetCreatedAt())
	return n
}

func (x *User) sizeField7() (n int) {
	if x.DeletedAt == 0 {
		return n
	}
	n += fastpb.SizeInt64(7, x.GetDeletedAt())
	return n
}

func (x *User) sizeField8() (n int) {
	if x.Membership == nil {
		return n
	}
	n += fastpb.SizeMessage(8, x.GetMembership())
	return n
}

func (x *User) sizeField9() (n int) {
	if x.TeamCount == 0 {
		return n
	}
	n += fastpb.SizeInt32(9, x.GetTeamCount())
	return n
}

func (x *User) sizeField10() (n int) {
	if len(x.TeamIds) == 0 {
		return n
	}
	for i := range x.GetTeamIds() {
		n += fastpb.SizeString(10, x.GetTeamIds()[i])
	}
	return n
}

func (x *User) sizeField11() (n int) {
	if x.MyAlbumCount == 0 {
		return n
	}
	n += fastpb.SizeInt32(11, x.GetMyAlbumCount())
	return n
}

func (x *User) sizeField12() (n int) {
	if len(x.MyAlbumIds) == 0 {
		return n
	}
	n += fastpb.SizeListPacked(12, len(x.GetMyAlbumIds()),
		func(numTagOrKey, numIdxOrVal int32) int {
			n := 0
			n += fastpb.SizeInt32(numTagOrKey, x.GetMyAlbumIds()[numIdxOrVal])
			return n
		})
	return n
}

func (x *User) sizeField13() (n int) {
	if x.FollowedAlbumCount == 0 {
		return n
	}
	n += fastpb.SizeInt32(13, x.GetFollowedAlbumCount())
	return n
}

func (x *User) sizeField14() (n int) {
	if len(x.FollowedAlbumIds) == 0 {
		return n
	}
	n += fastpb.SizeListPacked(14, len(x.GetFollowedAlbumIds()),
		func(numTagOrKey, numIdxOrVal int32) int {
			n := 0
			n += fastpb.SizeInt32(numTagOrKey, x.GetFollowedAlbumIds()[numIdxOrVal])
			return n
		})
	return n
}

func (x *User) sizeField15() (n int) {
	if x.StorageInfo == nil {
		return n
	}
	n += fastpb.SizeMessage(15, x.GetStorageInfo())
	return n
}

func (x *User) sizeField16() (n int) {
	if x.Points == 0 {
		return n
	}
	n += fastpb.SizeInt32(16, x.GetPoints())
	return n
}

func (x *User) sizeField17() (n int) {
	if len(x.Achievements) == 0 {
		return n
	}
	for i := range x.GetAchievements() {
		n += fastpb.SizeString(17, x.GetAchievements()[i])
	}
	return n
}

func (x *User_Membership) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *User_Membership) sizeField1() (n int) {
	if x.MemberId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetMemberId())
	return n
}

func (x *User_Membership) sizeField2() (n int) {
	if x.MemberLevel == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetMemberLevel())
	return n
}

func (x *User_StorageInfo) Size() (n int) {
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

func (x *User_StorageInfo) sizeField1() (n int) {
	if x.AvailablePhotos == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetAvailablePhotos())
	return n
}

func (x *User_StorageInfo) sizeField2() (n int) {
	if x.UsedPhotos == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetUsedPhotos())
	return n
}

func (x *User_StorageInfo) sizeField3() (n int) {
	if x.AvailableMemory == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetAvailableMemory())
	return n
}

func (x *User_StorageInfo) sizeField4() (n int) {
	if x.UsedMemory == 0 {
		return n
	}
	n += fastpb.SizeInt64(4, x.GetUsedMemory())
	return n
}

func (x *User_StorageInfo) sizeField5() (n int) {
	if x.AvailableAlbums == 0 {
		return n
	}
	n += fastpb.SizeInt32(5, x.GetAvailableAlbums())
	return n
}

func (x *User_StorageInfo) sizeField6() (n int) {
	if x.UsedAlbums == 0 {
		return n
	}
	n += fastpb.SizeInt32(6, x.GetUsedAlbums())
	return n
}

var fieldIDToName_UserPreview = map[int32]string{
	1: "Id",
	2: "Nickname",
	3: "Avatar",
}

var fieldIDToName_User = map[int32]string{
	1:  "Id",
	2:  "Nickname",
	3:  "Bio",
	4:  "Avatar",
	5:  "UpdatedAt",
	6:  "CreatedAt",
	7:  "DeletedAt",
	8:  "Membership",
	9:  "TeamCount",
	10: "TeamIds",
	11: "MyAlbumCount",
	12: "MyAlbumIds",
	13: "FollowedAlbumCount",
	14: "FollowedAlbumIds",
	15: "StorageInfo",
	16: "Points",
	17: "Achievements",
}

var fieldIDToName_User_Membership = map[int32]string{
	1: "MemberId",
	2: "MemberLevel",
}

var fieldIDToName_User_StorageInfo = map[int32]string{
	1: "AvailablePhotos",
	2: "UsedPhotos",
	3: "AvailableMemory",
	4: "UsedMemory",
	5: "AvailableAlbums",
	6: "UsedAlbums",
}
