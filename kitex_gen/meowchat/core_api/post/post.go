// Code generated by Kitex v0.6.2. DO NOT EDIT.

package post

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	core_api "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/core_api"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return postServiceInfo
}

var postServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "post"
	handlerType := (*core_api.Post)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetPostPreviews": kitex.NewMethodInfo(getPostPreviewsHandler, newGetPostPreviewsArgs, newGetPostPreviewsResult, false),
		"GetPostDetail":   kitex.NewMethodInfo(getPostDetailHandler, newGetPostDetailArgs, newGetPostDetailResult, false),
		"NewPost":         kitex.NewMethodInfo(newPostHandler, newNewPostArgs, newNewPostResult, false),
		"DeletePost":      kitex.NewMethodInfo(deletePostHandler, newDeletePostArgs, newDeletePostResult, false),
		"SetOfficial":     kitex.NewMethodInfo(setOfficialHandler, newSetOfficialArgs, newSetOfficialResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "meowchat.core_api",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.6.2",
		Extra:           extra,
	}
	return svcInfo
}

func getPostPreviewsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.GetPostPreviewsReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Post).GetPostPreviews(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetPostPreviewsArgs:
		success, err := handler.(core_api.Post).GetPostPreviews(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetPostPreviewsResult)
		realResult.Success = success
	}
	return nil
}
func newGetPostPreviewsArgs() interface{} {
	return &GetPostPreviewsArgs{}
}

func newGetPostPreviewsResult() interface{} {
	return &GetPostPreviewsResult{}
}

type GetPostPreviewsArgs struct {
	Req *core_api.GetPostPreviewsReq
}

func (p *GetPostPreviewsArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.GetPostPreviewsReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetPostPreviewsArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetPostPreviewsArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetPostPreviewsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetPostPreviewsArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetPostPreviewsArgs) Unmarshal(in []byte) error {
	msg := new(core_api.GetPostPreviewsReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetPostPreviewsArgs_Req_DEFAULT *core_api.GetPostPreviewsReq

func (p *GetPostPreviewsArgs) GetReq() *core_api.GetPostPreviewsReq {
	if !p.IsSetReq() {
		return GetPostPreviewsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetPostPreviewsArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetPostPreviewsArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetPostPreviewsResult struct {
	Success *core_api.GetPostPreviewsResp
}

var GetPostPreviewsResult_Success_DEFAULT *core_api.GetPostPreviewsResp

func (p *GetPostPreviewsResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.GetPostPreviewsResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetPostPreviewsResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetPostPreviewsResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetPostPreviewsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetPostPreviewsResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetPostPreviewsResult) Unmarshal(in []byte) error {
	msg := new(core_api.GetPostPreviewsResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPostPreviewsResult) GetSuccess() *core_api.GetPostPreviewsResp {
	if !p.IsSetSuccess() {
		return GetPostPreviewsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetPostPreviewsResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.GetPostPreviewsResp)
}

func (p *GetPostPreviewsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetPostPreviewsResult) GetResult() interface{} {
	return p.Success
}

func getPostDetailHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.GetPostDetailReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Post).GetPostDetail(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetPostDetailArgs:
		success, err := handler.(core_api.Post).GetPostDetail(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetPostDetailResult)
		realResult.Success = success
	}
	return nil
}
func newGetPostDetailArgs() interface{} {
	return &GetPostDetailArgs{}
}

func newGetPostDetailResult() interface{} {
	return &GetPostDetailResult{}
}

type GetPostDetailArgs struct {
	Req *core_api.GetPostDetailReq
}

func (p *GetPostDetailArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.GetPostDetailReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetPostDetailArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetPostDetailArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetPostDetailArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetPostDetailArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetPostDetailArgs) Unmarshal(in []byte) error {
	msg := new(core_api.GetPostDetailReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetPostDetailArgs_Req_DEFAULT *core_api.GetPostDetailReq

func (p *GetPostDetailArgs) GetReq() *core_api.GetPostDetailReq {
	if !p.IsSetReq() {
		return GetPostDetailArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetPostDetailArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetPostDetailArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetPostDetailResult struct {
	Success *core_api.GetPostDetailResp
}

var GetPostDetailResult_Success_DEFAULT *core_api.GetPostDetailResp

func (p *GetPostDetailResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.GetPostDetailResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetPostDetailResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetPostDetailResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetPostDetailResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetPostDetailResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetPostDetailResult) Unmarshal(in []byte) error {
	msg := new(core_api.GetPostDetailResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPostDetailResult) GetSuccess() *core_api.GetPostDetailResp {
	if !p.IsSetSuccess() {
		return GetPostDetailResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetPostDetailResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.GetPostDetailResp)
}

func (p *GetPostDetailResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetPostDetailResult) GetResult() interface{} {
	return p.Success
}

func newPostHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.NewPostReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Post).NewPost(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *NewPostArgs:
		success, err := handler.(core_api.Post).NewPost(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*NewPostResult)
		realResult.Success = success
	}
	return nil
}
func newNewPostArgs() interface{} {
	return &NewPostArgs{}
}

func newNewPostResult() interface{} {
	return &NewPostResult{}
}

type NewPostArgs struct {
	Req *core_api.NewPostReq
}

func (p *NewPostArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.NewPostReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *NewPostArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *NewPostArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *NewPostArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in NewPostArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *NewPostArgs) Unmarshal(in []byte) error {
	msg := new(core_api.NewPostReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var NewPostArgs_Req_DEFAULT *core_api.NewPostReq

func (p *NewPostArgs) GetReq() *core_api.NewPostReq {
	if !p.IsSetReq() {
		return NewPostArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *NewPostArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *NewPostArgs) GetFirstArgument() interface{} {
	return p.Req
}

type NewPostResult struct {
	Success *core_api.NewPostResp
}

var NewPostResult_Success_DEFAULT *core_api.NewPostResp

func (p *NewPostResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.NewPostResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *NewPostResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *NewPostResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *NewPostResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in NewPostResult")
	}
	return proto.Marshal(p.Success)
}

func (p *NewPostResult) Unmarshal(in []byte) error {
	msg := new(core_api.NewPostResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *NewPostResult) GetSuccess() *core_api.NewPostResp {
	if !p.IsSetSuccess() {
		return NewPostResult_Success_DEFAULT
	}
	return p.Success
}

func (p *NewPostResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.NewPostResp)
}

func (p *NewPostResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *NewPostResult) GetResult() interface{} {
	return p.Success
}

func deletePostHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.DeletePostReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Post).DeletePost(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *DeletePostArgs:
		success, err := handler.(core_api.Post).DeletePost(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeletePostResult)
		realResult.Success = success
	}
	return nil
}
func newDeletePostArgs() interface{} {
	return &DeletePostArgs{}
}

func newDeletePostResult() interface{} {
	return &DeletePostResult{}
}

type DeletePostArgs struct {
	Req *core_api.DeletePostReq
}

func (p *DeletePostArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.DeletePostReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeletePostArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeletePostArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeletePostArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeletePostArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *DeletePostArgs) Unmarshal(in []byte) error {
	msg := new(core_api.DeletePostReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeletePostArgs_Req_DEFAULT *core_api.DeletePostReq

func (p *DeletePostArgs) GetReq() *core_api.DeletePostReq {
	if !p.IsSetReq() {
		return DeletePostArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeletePostArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DeletePostArgs) GetFirstArgument() interface{} {
	return p.Req
}

type DeletePostResult struct {
	Success *core_api.DeletePostResp
}

var DeletePostResult_Success_DEFAULT *core_api.DeletePostResp

func (p *DeletePostResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.DeletePostResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeletePostResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeletePostResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeletePostResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeletePostResult")
	}
	return proto.Marshal(p.Success)
}

func (p *DeletePostResult) Unmarshal(in []byte) error {
	msg := new(core_api.DeletePostResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeletePostResult) GetSuccess() *core_api.DeletePostResp {
	if !p.IsSetSuccess() {
		return DeletePostResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeletePostResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.DeletePostResp)
}

func (p *DeletePostResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeletePostResult) GetResult() interface{} {
	return p.Success
}

func setOfficialHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.SetOfficialReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Post).SetOfficial(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *SetOfficialArgs:
		success, err := handler.(core_api.Post).SetOfficial(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SetOfficialResult)
		realResult.Success = success
	}
	return nil
}
func newSetOfficialArgs() interface{} {
	return &SetOfficialArgs{}
}

func newSetOfficialResult() interface{} {
	return &SetOfficialResult{}
}

type SetOfficialArgs struct {
	Req *core_api.SetOfficialReq
}

func (p *SetOfficialArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.SetOfficialReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *SetOfficialArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *SetOfficialArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *SetOfficialArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SetOfficialArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *SetOfficialArgs) Unmarshal(in []byte) error {
	msg := new(core_api.SetOfficialReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SetOfficialArgs_Req_DEFAULT *core_api.SetOfficialReq

func (p *SetOfficialArgs) GetReq() *core_api.SetOfficialReq {
	if !p.IsSetReq() {
		return SetOfficialArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetOfficialArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SetOfficialArgs) GetFirstArgument() interface{} {
	return p.Req
}

type SetOfficialResult struct {
	Success *core_api.SetOfficialResp
}

var SetOfficialResult_Success_DEFAULT *core_api.SetOfficialResp

func (p *SetOfficialResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.SetOfficialResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *SetOfficialResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *SetOfficialResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *SetOfficialResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SetOfficialResult")
	}
	return proto.Marshal(p.Success)
}

func (p *SetOfficialResult) Unmarshal(in []byte) error {
	msg := new(core_api.SetOfficialResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetOfficialResult) GetSuccess() *core_api.SetOfficialResp {
	if !p.IsSetSuccess() {
		return SetOfficialResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetOfficialResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.SetOfficialResp)
}

func (p *SetOfficialResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetOfficialResult) GetResult() interface{} {
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

func (p *kClient) GetPostPreviews(ctx context.Context, Req *core_api.GetPostPreviewsReq) (r *core_api.GetPostPreviewsResp, err error) {
	var _args GetPostPreviewsArgs
	_args.Req = Req
	var _result GetPostPreviewsResult
	if err = p.c.Call(ctx, "GetPostPreviews", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetPostDetail(ctx context.Context, Req *core_api.GetPostDetailReq) (r *core_api.GetPostDetailResp, err error) {
	var _args GetPostDetailArgs
	_args.Req = Req
	var _result GetPostDetailResult
	if err = p.c.Call(ctx, "GetPostDetail", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) NewPost(ctx context.Context, Req *core_api.NewPostReq) (r *core_api.NewPostResp, err error) {
	var _args NewPostArgs
	_args.Req = Req
	var _result NewPostResult
	if err = p.c.Call(ctx, "NewPost", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeletePost(ctx context.Context, Req *core_api.DeletePostReq) (r *core_api.DeletePostResp, err error) {
	var _args DeletePostArgs
	_args.Req = Req
	var _result DeletePostResult
	if err = p.c.Call(ctx, "DeletePost", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SetOfficial(ctx context.Context, Req *core_api.SetOfficialReq) (r *core_api.SetOfficialResp, err error) {
	var _args SetOfficialArgs
	_args.Req = Req
	var _result SetOfficialResult
	if err = p.c.Call(ctx, "SetOfficial", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
