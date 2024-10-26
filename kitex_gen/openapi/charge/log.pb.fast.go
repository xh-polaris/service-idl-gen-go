// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package charge

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
	basic "github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *CreateLogReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 9:
		offset, err = x.fastReadField9(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreateLogReq[number], err)
}

func (x *CreateLogReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.MarginId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateLogReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.FullInterfaceId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateLogReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateLogReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.KeyId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateLogReq) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.Status = LogStatus(v)
	return offset, nil
}

func (x *CreateLogReq) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.Info, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateLogReq) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.Count, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CreateLogReq) fastReadField9(buf []byte, _type int8) (offset int, err error) {
	x.Timestamp, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CreateLogResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreateLogResp[number], err)
}

func (x *CreateLogResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Done, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *CreateLogResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Msg, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetLogReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetLogReq[number], err)
}

func (x *GetLogReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.FullInterfaceId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetLogReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v basic.PaginationOptions
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.PaginationOptions = &v
	return offset, nil
}

func (x *GetLogResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetLogResp[number], err)
}

func (x *GetLogResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Log
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Logs = append(x.Logs, &v)
	return offset, nil
}

func (x *GetLogResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Total, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CreateLogReq) FastWrite(buf []byte) (offset int) {
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
	offset += x.fastWriteField9(buf[offset:])
	return offset
}

func (x *CreateLogReq) fastWriteField1(buf []byte) (offset int) {
	if x.MarginId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetMarginId())
	return offset
}

func (x *CreateLogReq) fastWriteField2(buf []byte) (offset int) {
	if x.FullInterfaceId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetFullInterfaceId())
	return offset
}

func (x *CreateLogReq) fastWriteField3(buf []byte) (offset int) {
	if x.UserId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetUserId())
	return offset
}

func (x *CreateLogReq) fastWriteField4(buf []byte) (offset int) {
	if x.KeyId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetKeyId())
	return offset
}

func (x *CreateLogReq) fastWriteField5(buf []byte) (offset int) {
	if x.Status == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 5, int32(x.GetStatus()))
	return offset
}

func (x *CreateLogReq) fastWriteField6(buf []byte) (offset int) {
	if x.Info == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetInfo())
	return offset
}

func (x *CreateLogReq) fastWriteField7(buf []byte) (offset int) {
	if x.Count == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 7, x.GetCount())
	return offset
}

func (x *CreateLogReq) fastWriteField9(buf []byte) (offset int) {
	if x.Timestamp == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 9, x.GetTimestamp())
	return offset
}

func (x *CreateLogResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *CreateLogResp) fastWriteField1(buf []byte) (offset int) {
	if !x.Done {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetDone())
	return offset
}

func (x *CreateLogResp) fastWriteField2(buf []byte) (offset int) {
	if x.Msg == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetMsg())
	return offset
}

func (x *GetLogReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetLogReq) fastWriteField1(buf []byte) (offset int) {
	if x.FullInterfaceId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetFullInterfaceId())
	return offset
}

func (x *GetLogReq) fastWriteField2(buf []byte) (offset int) {
	if x.PaginationOptions == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetPaginationOptions())
	return offset
}

func (x *GetLogResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetLogResp) fastWriteField1(buf []byte) (offset int) {
	if x.Logs == nil {
		return offset
	}
	for i := range x.GetLogs() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetLogs()[i])
	}
	return offset
}

func (x *GetLogResp) fastWriteField2(buf []byte) (offset int) {
	if x.Total == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetTotal())
	return offset
}

func (x *CreateLogReq) Size() (n int) {
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
	n += x.sizeField9()
	return n
}

func (x *CreateLogReq) sizeField1() (n int) {
	if x.MarginId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetMarginId())
	return n
}

func (x *CreateLogReq) sizeField2() (n int) {
	if x.FullInterfaceId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetFullInterfaceId())
	return n
}

func (x *CreateLogReq) sizeField3() (n int) {
	if x.UserId == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetUserId())
	return n
}

func (x *CreateLogReq) sizeField4() (n int) {
	if x.KeyId == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetKeyId())
	return n
}

func (x *CreateLogReq) sizeField5() (n int) {
	if x.Status == 0 {
		return n
	}
	n += fastpb.SizeInt32(5, int32(x.GetStatus()))
	return n
}

func (x *CreateLogReq) sizeField6() (n int) {
	if x.Info == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetInfo())
	return n
}

func (x *CreateLogReq) sizeField7() (n int) {
	if x.Count == 0 {
		return n
	}
	n += fastpb.SizeInt64(7, x.GetCount())
	return n
}

func (x *CreateLogReq) sizeField9() (n int) {
	if x.Timestamp == 0 {
		return n
	}
	n += fastpb.SizeInt64(9, x.GetTimestamp())
	return n
}

func (x *CreateLogResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *CreateLogResp) sizeField1() (n int) {
	if !x.Done {
		return n
	}
	n += fastpb.SizeBool(1, x.GetDone())
	return n
}

func (x *CreateLogResp) sizeField2() (n int) {
	if x.Msg == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetMsg())
	return n
}

func (x *GetLogReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetLogReq) sizeField1() (n int) {
	if x.FullInterfaceId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetFullInterfaceId())
	return n
}

func (x *GetLogReq) sizeField2() (n int) {
	if x.PaginationOptions == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetPaginationOptions())
	return n
}

func (x *GetLogResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetLogResp) sizeField1() (n int) {
	if x.Logs == nil {
		return n
	}
	for i := range x.GetLogs() {
		n += fastpb.SizeMessage(1, x.GetLogs()[i])
	}
	return n
}

func (x *GetLogResp) sizeField2() (n int) {
	if x.Total == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetTotal())
	return n
}

var fieldIDToName_CreateLogReq = map[int32]string{
	1: "MarginId",
	2: "FullInterfaceId",
	3: "UserId",
	4: "KeyId",
	5: "Status",
	6: "Info",
	7: "Count",
	9: "Timestamp",
}

var fieldIDToName_CreateLogResp = map[int32]string{
	1: "Done",
	2: "Msg",
}

var fieldIDToName_GetLogReq = map[int32]string{
	1: "FullInterfaceId",
	2: "PaginationOptions",
}

var fieldIDToName_GetLogResp = map[int32]string{
	1: "Logs",
	2: "Total",
}

var _ = basic.File_basic_pagination_proto
