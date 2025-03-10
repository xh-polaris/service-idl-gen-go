// Code generated by Kitex v0.12.1. DO NOT EDIT.

package collection

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	core_api "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/core_api"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"GetCatPreviews": kitex.NewMethodInfo(
		getCatPreviewsHandler,
		newGetCatPreviewsArgs,
		newGetCatPreviewsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetCatDetail": kitex.NewMethodInfo(
		getCatDetailHandler,
		newGetCatDetailArgs,
		newGetCatDetailResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"NewCat": kitex.NewMethodInfo(
		newCatHandler,
		newNewCatArgs,
		newNewCatResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"DeleteCat": kitex.NewMethodInfo(
		deleteCatHandler,
		newDeleteCatArgs,
		newDeleteCatResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"SearchCat": kitex.NewMethodInfo(
		searchCatHandler,
		newSearchCatArgs,
		newSearchCatResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"CreateImage": kitex.NewMethodInfo(
		createImageHandler,
		newCreateImageArgs,
		newCreateImageResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"DeleteImage": kitex.NewMethodInfo(
		deleteImageHandler,
		newDeleteImageArgs,
		newDeleteImageResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetImageByCat": kitex.NewMethodInfo(
		getImageByCatHandler,
		newGetImageByCatArgs,
		newGetImageByCatResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	collectionServiceInfo                = NewServiceInfo()
	collectionServiceInfoForClient       = NewServiceInfoForClient()
	collectionServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return collectionServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return collectionServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return collectionServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "collection"
	handlerType := (*core_api.Collection)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "meowchat.core_api",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.12.1",
		Extra:           extra,
	}
	return svcInfo
}

func getCatPreviewsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.GetCatPreviewsReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Collection).GetCatPreviews(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetCatPreviewsArgs:
		success, err := handler.(core_api.Collection).GetCatPreviews(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetCatPreviewsResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetCatPreviewsArgs() interface{} {
	return &GetCatPreviewsArgs{}
}

func newGetCatPreviewsResult() interface{} {
	return &GetCatPreviewsResult{}
}

type GetCatPreviewsArgs struct {
	Req *core_api.GetCatPreviewsReq
}

func (p *GetCatPreviewsArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.GetCatPreviewsReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetCatPreviewsArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetCatPreviewsArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetCatPreviewsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetCatPreviewsArgs) Unmarshal(in []byte) error {
	msg := new(core_api.GetCatPreviewsReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetCatPreviewsArgs_Req_DEFAULT *core_api.GetCatPreviewsReq

func (p *GetCatPreviewsArgs) GetReq() *core_api.GetCatPreviewsReq {
	if !p.IsSetReq() {
		return GetCatPreviewsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetCatPreviewsArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetCatPreviewsArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetCatPreviewsResult struct {
	Success *core_api.GetCatPreviewsResp
}

var GetCatPreviewsResult_Success_DEFAULT *core_api.GetCatPreviewsResp

func (p *GetCatPreviewsResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.GetCatPreviewsResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetCatPreviewsResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetCatPreviewsResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetCatPreviewsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetCatPreviewsResult) Unmarshal(in []byte) error {
	msg := new(core_api.GetCatPreviewsResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetCatPreviewsResult) GetSuccess() *core_api.GetCatPreviewsResp {
	if !p.IsSetSuccess() {
		return GetCatPreviewsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetCatPreviewsResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.GetCatPreviewsResp)
}

func (p *GetCatPreviewsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetCatPreviewsResult) GetResult() interface{} {
	return p.Success
}

func getCatDetailHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.GetCatDetailReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Collection).GetCatDetail(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetCatDetailArgs:
		success, err := handler.(core_api.Collection).GetCatDetail(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetCatDetailResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetCatDetailArgs() interface{} {
	return &GetCatDetailArgs{}
}

func newGetCatDetailResult() interface{} {
	return &GetCatDetailResult{}
}

type GetCatDetailArgs struct {
	Req *core_api.GetCatDetailReq
}

func (p *GetCatDetailArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.GetCatDetailReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetCatDetailArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetCatDetailArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetCatDetailArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetCatDetailArgs) Unmarshal(in []byte) error {
	msg := new(core_api.GetCatDetailReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetCatDetailArgs_Req_DEFAULT *core_api.GetCatDetailReq

func (p *GetCatDetailArgs) GetReq() *core_api.GetCatDetailReq {
	if !p.IsSetReq() {
		return GetCatDetailArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetCatDetailArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetCatDetailArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetCatDetailResult struct {
	Success *core_api.GetCatDetailResp
}

var GetCatDetailResult_Success_DEFAULT *core_api.GetCatDetailResp

func (p *GetCatDetailResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.GetCatDetailResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetCatDetailResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetCatDetailResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetCatDetailResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetCatDetailResult) Unmarshal(in []byte) error {
	msg := new(core_api.GetCatDetailResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetCatDetailResult) GetSuccess() *core_api.GetCatDetailResp {
	if !p.IsSetSuccess() {
		return GetCatDetailResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetCatDetailResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.GetCatDetailResp)
}

func (p *GetCatDetailResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetCatDetailResult) GetResult() interface{} {
	return p.Success
}

func newCatHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.NewCatReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Collection).NewCat(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *NewCatArgs:
		success, err := handler.(core_api.Collection).NewCat(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*NewCatResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newNewCatArgs() interface{} {
	return &NewCatArgs{}
}

func newNewCatResult() interface{} {
	return &NewCatResult{}
}

type NewCatArgs struct {
	Req *core_api.NewCatReq
}

func (p *NewCatArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.NewCatReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *NewCatArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *NewCatArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *NewCatArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *NewCatArgs) Unmarshal(in []byte) error {
	msg := new(core_api.NewCatReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var NewCatArgs_Req_DEFAULT *core_api.NewCatReq

func (p *NewCatArgs) GetReq() *core_api.NewCatReq {
	if !p.IsSetReq() {
		return NewCatArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *NewCatArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *NewCatArgs) GetFirstArgument() interface{} {
	return p.Req
}

type NewCatResult struct {
	Success *core_api.NewCatResp
}

var NewCatResult_Success_DEFAULT *core_api.NewCatResp

func (p *NewCatResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.NewCatResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *NewCatResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *NewCatResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *NewCatResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *NewCatResult) Unmarshal(in []byte) error {
	msg := new(core_api.NewCatResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *NewCatResult) GetSuccess() *core_api.NewCatResp {
	if !p.IsSetSuccess() {
		return NewCatResult_Success_DEFAULT
	}
	return p.Success
}

func (p *NewCatResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.NewCatResp)
}

func (p *NewCatResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *NewCatResult) GetResult() interface{} {
	return p.Success
}

func deleteCatHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.DeleteCatReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Collection).DeleteCat(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *DeleteCatArgs:
		success, err := handler.(core_api.Collection).DeleteCat(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteCatResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newDeleteCatArgs() interface{} {
	return &DeleteCatArgs{}
}

func newDeleteCatResult() interface{} {
	return &DeleteCatResult{}
}

type DeleteCatArgs struct {
	Req *core_api.DeleteCatReq
}

func (p *DeleteCatArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.DeleteCatReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeleteCatArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeleteCatArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeleteCatArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteCatArgs) Unmarshal(in []byte) error {
	msg := new(core_api.DeleteCatReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteCatArgs_Req_DEFAULT *core_api.DeleteCatReq

func (p *DeleteCatArgs) GetReq() *core_api.DeleteCatReq {
	if !p.IsSetReq() {
		return DeleteCatArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteCatArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DeleteCatArgs) GetFirstArgument() interface{} {
	return p.Req
}

type DeleteCatResult struct {
	Success *core_api.DeleteCatResp
}

var DeleteCatResult_Success_DEFAULT *core_api.DeleteCatResp

func (p *DeleteCatResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.DeleteCatResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeleteCatResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeleteCatResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeleteCatResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteCatResult) Unmarshal(in []byte) error {
	msg := new(core_api.DeleteCatResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteCatResult) GetSuccess() *core_api.DeleteCatResp {
	if !p.IsSetSuccess() {
		return DeleteCatResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteCatResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.DeleteCatResp)
}

func (p *DeleteCatResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteCatResult) GetResult() interface{} {
	return p.Success
}

func searchCatHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.SearchCatReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Collection).SearchCat(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *SearchCatArgs:
		success, err := handler.(core_api.Collection).SearchCat(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SearchCatResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newSearchCatArgs() interface{} {
	return &SearchCatArgs{}
}

func newSearchCatResult() interface{} {
	return &SearchCatResult{}
}

type SearchCatArgs struct {
	Req *core_api.SearchCatReq
}

func (p *SearchCatArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.SearchCatReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *SearchCatArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *SearchCatArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *SearchCatArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *SearchCatArgs) Unmarshal(in []byte) error {
	msg := new(core_api.SearchCatReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SearchCatArgs_Req_DEFAULT *core_api.SearchCatReq

func (p *SearchCatArgs) GetReq() *core_api.SearchCatReq {
	if !p.IsSetReq() {
		return SearchCatArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SearchCatArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SearchCatArgs) GetFirstArgument() interface{} {
	return p.Req
}

type SearchCatResult struct {
	Success *core_api.SearchCatResp
}

var SearchCatResult_Success_DEFAULT *core_api.SearchCatResp

func (p *SearchCatResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.SearchCatResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *SearchCatResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *SearchCatResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *SearchCatResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *SearchCatResult) Unmarshal(in []byte) error {
	msg := new(core_api.SearchCatResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SearchCatResult) GetSuccess() *core_api.SearchCatResp {
	if !p.IsSetSuccess() {
		return SearchCatResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SearchCatResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.SearchCatResp)
}

func (p *SearchCatResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SearchCatResult) GetResult() interface{} {
	return p.Success
}

func createImageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.CreateImageReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Collection).CreateImage(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *CreateImageArgs:
		success, err := handler.(core_api.Collection).CreateImage(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateImageResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newCreateImageArgs() interface{} {
	return &CreateImageArgs{}
}

func newCreateImageResult() interface{} {
	return &CreateImageResult{}
}

type CreateImageArgs struct {
	Req *core_api.CreateImageReq
}

func (p *CreateImageArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.CreateImageReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CreateImageArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CreateImageArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CreateImageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CreateImageArgs) Unmarshal(in []byte) error {
	msg := new(core_api.CreateImageReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateImageArgs_Req_DEFAULT *core_api.CreateImageReq

func (p *CreateImageArgs) GetReq() *core_api.CreateImageReq {
	if !p.IsSetReq() {
		return CreateImageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateImageArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CreateImageArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CreateImageResult struct {
	Success *core_api.CreateImageResp
}

var CreateImageResult_Success_DEFAULT *core_api.CreateImageResp

func (p *CreateImageResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.CreateImageResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CreateImageResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CreateImageResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CreateImageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CreateImageResult) Unmarshal(in []byte) error {
	msg := new(core_api.CreateImageResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateImageResult) GetSuccess() *core_api.CreateImageResp {
	if !p.IsSetSuccess() {
		return CreateImageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateImageResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.CreateImageResp)
}

func (p *CreateImageResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CreateImageResult) GetResult() interface{} {
	return p.Success
}

func deleteImageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.DeleteImageReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Collection).DeleteImage(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *DeleteImageArgs:
		success, err := handler.(core_api.Collection).DeleteImage(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteImageResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newDeleteImageArgs() interface{} {
	return &DeleteImageArgs{}
}

func newDeleteImageResult() interface{} {
	return &DeleteImageResult{}
}

type DeleteImageArgs struct {
	Req *core_api.DeleteImageReq
}

func (p *DeleteImageArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.DeleteImageReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeleteImageArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeleteImageArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeleteImageArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteImageArgs) Unmarshal(in []byte) error {
	msg := new(core_api.DeleteImageReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteImageArgs_Req_DEFAULT *core_api.DeleteImageReq

func (p *DeleteImageArgs) GetReq() *core_api.DeleteImageReq {
	if !p.IsSetReq() {
		return DeleteImageArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteImageArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DeleteImageArgs) GetFirstArgument() interface{} {
	return p.Req
}

type DeleteImageResult struct {
	Success *core_api.DeleteImageResp
}

var DeleteImageResult_Success_DEFAULT *core_api.DeleteImageResp

func (p *DeleteImageResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.DeleteImageResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeleteImageResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeleteImageResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeleteImageResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteImageResult) Unmarshal(in []byte) error {
	msg := new(core_api.DeleteImageResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteImageResult) GetSuccess() *core_api.DeleteImageResp {
	if !p.IsSetSuccess() {
		return DeleteImageResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteImageResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.DeleteImageResp)
}

func (p *DeleteImageResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteImageResult) GetResult() interface{} {
	return p.Success
}

func getImageByCatHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.GetImageByCatReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Collection).GetImageByCat(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetImageByCatArgs:
		success, err := handler.(core_api.Collection).GetImageByCat(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetImageByCatResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetImageByCatArgs() interface{} {
	return &GetImageByCatArgs{}
}

func newGetImageByCatResult() interface{} {
	return &GetImageByCatResult{}
}

type GetImageByCatArgs struct {
	Req *core_api.GetImageByCatReq
}

func (p *GetImageByCatArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.GetImageByCatReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetImageByCatArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetImageByCatArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetImageByCatArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetImageByCatArgs) Unmarshal(in []byte) error {
	msg := new(core_api.GetImageByCatReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetImageByCatArgs_Req_DEFAULT *core_api.GetImageByCatReq

func (p *GetImageByCatArgs) GetReq() *core_api.GetImageByCatReq {
	if !p.IsSetReq() {
		return GetImageByCatArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetImageByCatArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetImageByCatArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetImageByCatResult struct {
	Success *core_api.GetImageByCatResp
}

var GetImageByCatResult_Success_DEFAULT *core_api.GetImageByCatResp

func (p *GetImageByCatResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.GetImageByCatResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetImageByCatResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetImageByCatResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetImageByCatResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetImageByCatResult) Unmarshal(in []byte) error {
	msg := new(core_api.GetImageByCatResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetImageByCatResult) GetSuccess() *core_api.GetImageByCatResp {
	if !p.IsSetSuccess() {
		return GetImageByCatResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetImageByCatResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.GetImageByCatResp)
}

func (p *GetImageByCatResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetImageByCatResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetCatPreviews(ctx context.Context, Req *core_api.GetCatPreviewsReq) (r *core_api.GetCatPreviewsResp, err error) {
	var _args GetCatPreviewsArgs
	_args.Req = Req
	var _result GetCatPreviewsResult
	if err = p.c.Call(ctx, "GetCatPreviews", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetCatDetail(ctx context.Context, Req *core_api.GetCatDetailReq) (r *core_api.GetCatDetailResp, err error) {
	var _args GetCatDetailArgs
	_args.Req = Req
	var _result GetCatDetailResult
	if err = p.c.Call(ctx, "GetCatDetail", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) NewCat(ctx context.Context, Req *core_api.NewCatReq) (r *core_api.NewCatResp, err error) {
	var _args NewCatArgs
	_args.Req = Req
	var _result NewCatResult
	if err = p.c.Call(ctx, "NewCat", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteCat(ctx context.Context, Req *core_api.DeleteCatReq) (r *core_api.DeleteCatResp, err error) {
	var _args DeleteCatArgs
	_args.Req = Req
	var _result DeleteCatResult
	if err = p.c.Call(ctx, "DeleteCat", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SearchCat(ctx context.Context, Req *core_api.SearchCatReq) (r *core_api.SearchCatResp, err error) {
	var _args SearchCatArgs
	_args.Req = Req
	var _result SearchCatResult
	if err = p.c.Call(ctx, "SearchCat", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateImage(ctx context.Context, Req *core_api.CreateImageReq) (r *core_api.CreateImageResp, err error) {
	var _args CreateImageArgs
	_args.Req = Req
	var _result CreateImageResult
	if err = p.c.Call(ctx, "CreateImage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteImage(ctx context.Context, Req *core_api.DeleteImageReq) (r *core_api.DeleteImageResp, err error) {
	var _args DeleteImageArgs
	_args.Req = Req
	var _result DeleteImageResult
	if err = p.c.Call(ctx, "DeleteImage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetImageByCat(ctx context.Context, Req *core_api.GetImageByCatReq) (r *core_api.GetImageByCatResp, err error) {
	var _args GetImageByCatArgs
	_args.Req = Req
	var _result GetImageByCatResult
	if err = p.c.Call(ctx, "GetImageByCat", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
