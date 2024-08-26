// Code generated by Kitex v0.10.3. DO NOT EDIT.

package photoservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	content "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowcloud/content"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"UploadPhoto": kitex.NewMethodInfo(
		uploadPhotoHandler,
		newUploadPhotoArgs,
		newUploadPhotoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"UpdatePhoto": kitex.NewMethodInfo(
		updatePhotoHandler,
		newUpdatePhotoArgs,
		newUpdatePhotoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetPhoto": kitex.NewMethodInfo(
		getPhotoHandler,
		newGetPhotoArgs,
		newGetPhotoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"ListPhotos": kitex.NewMethodInfo(
		listPhotosHandler,
		newListPhotosArgs,
		newListPhotosResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"DeletePhoto": kitex.NewMethodInfo(
		deletePhotoHandler,
		newDeletePhotoArgs,
		newDeletePhotoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	photoServiceServiceInfo                = NewServiceInfo()
	photoServiceServiceInfoForClient       = NewServiceInfoForClient()
	photoServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return photoServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return photoServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return photoServiceServiceInfoForClient
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
	serviceName := "PhotoService"
	handlerType := (*content.PhotoService)(nil)
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
		"PackageName": "meowcloud.content",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.10.3",
		Extra:           extra,
	}
	return svcInfo
}

func uploadPhotoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(content.UploadPhotoReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(content.PhotoService).UploadPhoto(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *UploadPhotoArgs:
		success, err := handler.(content.PhotoService).UploadPhoto(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UploadPhotoResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newUploadPhotoArgs() interface{} {
	return &UploadPhotoArgs{}
}

func newUploadPhotoResult() interface{} {
	return &UploadPhotoResult{}
}

type UploadPhotoArgs struct {
	Req *content.UploadPhotoReq
}

func (p *UploadPhotoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(content.UploadPhotoReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UploadPhotoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UploadPhotoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UploadPhotoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UploadPhotoArgs) Unmarshal(in []byte) error {
	msg := new(content.UploadPhotoReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UploadPhotoArgs_Req_DEFAULT *content.UploadPhotoReq

func (p *UploadPhotoArgs) GetReq() *content.UploadPhotoReq {
	if !p.IsSetReq() {
		return UploadPhotoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UploadPhotoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UploadPhotoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UploadPhotoResult struct {
	Success *content.UploadPhotoResp
}

var UploadPhotoResult_Success_DEFAULT *content.UploadPhotoResp

func (p *UploadPhotoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(content.UploadPhotoResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UploadPhotoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UploadPhotoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UploadPhotoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UploadPhotoResult) Unmarshal(in []byte) error {
	msg := new(content.UploadPhotoResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UploadPhotoResult) GetSuccess() *content.UploadPhotoResp {
	if !p.IsSetSuccess() {
		return UploadPhotoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UploadPhotoResult) SetSuccess(x interface{}) {
	p.Success = x.(*content.UploadPhotoResp)
}

func (p *UploadPhotoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UploadPhotoResult) GetResult() interface{} {
	return p.Success
}

func updatePhotoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(content.UpdatePhotoReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(content.PhotoService).UpdatePhoto(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *UpdatePhotoArgs:
		success, err := handler.(content.PhotoService).UpdatePhoto(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdatePhotoResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newUpdatePhotoArgs() interface{} {
	return &UpdatePhotoArgs{}
}

func newUpdatePhotoResult() interface{} {
	return &UpdatePhotoResult{}
}

type UpdatePhotoArgs struct {
	Req *content.UpdatePhotoReq
}

func (p *UpdatePhotoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(content.UpdatePhotoReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdatePhotoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdatePhotoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdatePhotoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UpdatePhotoArgs) Unmarshal(in []byte) error {
	msg := new(content.UpdatePhotoReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdatePhotoArgs_Req_DEFAULT *content.UpdatePhotoReq

func (p *UpdatePhotoArgs) GetReq() *content.UpdatePhotoReq {
	if !p.IsSetReq() {
		return UpdatePhotoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdatePhotoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdatePhotoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdatePhotoResult struct {
	Success *content.UpdatePhotoResp
}

var UpdatePhotoResult_Success_DEFAULT *content.UpdatePhotoResp

func (p *UpdatePhotoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(content.UpdatePhotoResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdatePhotoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdatePhotoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdatePhotoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UpdatePhotoResult) Unmarshal(in []byte) error {
	msg := new(content.UpdatePhotoResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdatePhotoResult) GetSuccess() *content.UpdatePhotoResp {
	if !p.IsSetSuccess() {
		return UpdatePhotoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdatePhotoResult) SetSuccess(x interface{}) {
	p.Success = x.(*content.UpdatePhotoResp)
}

func (p *UpdatePhotoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdatePhotoResult) GetResult() interface{} {
	return p.Success
}

func getPhotoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(content.GetPhotoReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(content.PhotoService).GetPhoto(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetPhotoArgs:
		success, err := handler.(content.PhotoService).GetPhoto(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetPhotoResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetPhotoArgs() interface{} {
	return &GetPhotoArgs{}
}

func newGetPhotoResult() interface{} {
	return &GetPhotoResult{}
}

type GetPhotoArgs struct {
	Req *content.GetPhotoReq
}

func (p *GetPhotoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(content.GetPhotoReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetPhotoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetPhotoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetPhotoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetPhotoArgs) Unmarshal(in []byte) error {
	msg := new(content.GetPhotoReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetPhotoArgs_Req_DEFAULT *content.GetPhotoReq

func (p *GetPhotoArgs) GetReq() *content.GetPhotoReq {
	if !p.IsSetReq() {
		return GetPhotoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetPhotoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetPhotoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetPhotoResult struct {
	Success *content.GetPhotoResp
}

var GetPhotoResult_Success_DEFAULT *content.GetPhotoResp

func (p *GetPhotoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(content.GetPhotoResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetPhotoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetPhotoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetPhotoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetPhotoResult) Unmarshal(in []byte) error {
	msg := new(content.GetPhotoResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPhotoResult) GetSuccess() *content.GetPhotoResp {
	if !p.IsSetSuccess() {
		return GetPhotoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetPhotoResult) SetSuccess(x interface{}) {
	p.Success = x.(*content.GetPhotoResp)
}

func (p *GetPhotoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetPhotoResult) GetResult() interface{} {
	return p.Success
}

func listPhotosHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(content.ListPhotosReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(content.PhotoService).ListPhotos(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ListPhotosArgs:
		success, err := handler.(content.PhotoService).ListPhotos(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ListPhotosResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newListPhotosArgs() interface{} {
	return &ListPhotosArgs{}
}

func newListPhotosResult() interface{} {
	return &ListPhotosResult{}
}

type ListPhotosArgs struct {
	Req *content.ListPhotosReq
}

func (p *ListPhotosArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(content.ListPhotosReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ListPhotosArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ListPhotosArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ListPhotosArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ListPhotosArgs) Unmarshal(in []byte) error {
	msg := new(content.ListPhotosReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ListPhotosArgs_Req_DEFAULT *content.ListPhotosReq

func (p *ListPhotosArgs) GetReq() *content.ListPhotosReq {
	if !p.IsSetReq() {
		return ListPhotosArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ListPhotosArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ListPhotosArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ListPhotosResult struct {
	Success *content.ListPhotosResp
}

var ListPhotosResult_Success_DEFAULT *content.ListPhotosResp

func (p *ListPhotosResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(content.ListPhotosResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ListPhotosResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ListPhotosResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ListPhotosResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ListPhotosResult) Unmarshal(in []byte) error {
	msg := new(content.ListPhotosResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListPhotosResult) GetSuccess() *content.ListPhotosResp {
	if !p.IsSetSuccess() {
		return ListPhotosResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ListPhotosResult) SetSuccess(x interface{}) {
	p.Success = x.(*content.ListPhotosResp)
}

func (p *ListPhotosResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ListPhotosResult) GetResult() interface{} {
	return p.Success
}

func deletePhotoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(content.DeletePhotoReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(content.PhotoService).DeletePhoto(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *DeletePhotoArgs:
		success, err := handler.(content.PhotoService).DeletePhoto(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeletePhotoResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newDeletePhotoArgs() interface{} {
	return &DeletePhotoArgs{}
}

func newDeletePhotoResult() interface{} {
	return &DeletePhotoResult{}
}

type DeletePhotoArgs struct {
	Req *content.DeletePhotoReq
}

func (p *DeletePhotoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(content.DeletePhotoReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeletePhotoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeletePhotoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeletePhotoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *DeletePhotoArgs) Unmarshal(in []byte) error {
	msg := new(content.DeletePhotoReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeletePhotoArgs_Req_DEFAULT *content.DeletePhotoReq

func (p *DeletePhotoArgs) GetReq() *content.DeletePhotoReq {
	if !p.IsSetReq() {
		return DeletePhotoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeletePhotoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DeletePhotoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type DeletePhotoResult struct {
	Success *content.DeletePhotoResp
}

var DeletePhotoResult_Success_DEFAULT *content.DeletePhotoResp

func (p *DeletePhotoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(content.DeletePhotoResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeletePhotoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeletePhotoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeletePhotoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *DeletePhotoResult) Unmarshal(in []byte) error {
	msg := new(content.DeletePhotoResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeletePhotoResult) GetSuccess() *content.DeletePhotoResp {
	if !p.IsSetSuccess() {
		return DeletePhotoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeletePhotoResult) SetSuccess(x interface{}) {
	p.Success = x.(*content.DeletePhotoResp)
}

func (p *DeletePhotoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeletePhotoResult) GetResult() interface{} {
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

func (p *kClient) UploadPhoto(ctx context.Context, Req *content.UploadPhotoReq) (r *content.UploadPhotoResp, err error) {
	var _args UploadPhotoArgs
	_args.Req = Req
	var _result UploadPhotoResult
	if err = p.c.Call(ctx, "UploadPhoto", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdatePhoto(ctx context.Context, Req *content.UpdatePhotoReq) (r *content.UpdatePhotoResp, err error) {
	var _args UpdatePhotoArgs
	_args.Req = Req
	var _result UpdatePhotoResult
	if err = p.c.Call(ctx, "UpdatePhoto", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetPhoto(ctx context.Context, Req *content.GetPhotoReq) (r *content.GetPhotoResp, err error) {
	var _args GetPhotoArgs
	_args.Req = Req
	var _result GetPhotoResult
	if err = p.c.Call(ctx, "GetPhoto", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListPhotos(ctx context.Context, Req *content.ListPhotosReq) (r *content.ListPhotosResp, err error) {
	var _args ListPhotosArgs
	_args.Req = Req
	var _result ListPhotosResult
	if err = p.c.Call(ctx, "ListPhotos", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeletePhoto(ctx context.Context, Req *content.DeletePhotoReq) (r *content.DeletePhotoResp, err error) {
	var _args DeletePhotoArgs
	_args.Req = Req
	var _result DeletePhotoResult
	if err = p.c.Call(ctx, "DeletePhoto", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
