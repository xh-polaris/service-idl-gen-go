// Code generated by Kitex v0.8.0. DO NOT EDIT.

package comment

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	core_api "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/core_api"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return commentServiceInfo
}

var commentServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "comment"
	handlerType := (*core_api.Comment)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetComments":   kitex.NewMethodInfo(getCommentsHandler, newGetCommentsArgs, newGetCommentsResult, false),
		"NewComment":    kitex.NewMethodInfo(newCommentHandler, newNewCommentArgs, newNewCommentResult, false),
		"DeleteComment": kitex.NewMethodInfo(deleteCommentHandler, newDeleteCommentArgs, newDeleteCommentResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "meowchat.core_api",
		"ServiceFilePath": ``,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.8.0",
		Extra:           extra,
	}
	return svcInfo
}

func getCommentsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.GetCommentsReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Comment).GetComments(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetCommentsArgs:
		success, err := handler.(core_api.Comment).GetComments(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetCommentsResult)
		realResult.Success = success
	}
	return nil
}
func newGetCommentsArgs() interface{} {
	return &GetCommentsArgs{}
}

func newGetCommentsResult() interface{} {
	return &GetCommentsResult{}
}

type GetCommentsArgs struct {
	Req *core_api.GetCommentsReq
}

func (p *GetCommentsArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.GetCommentsReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetCommentsArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetCommentsArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetCommentsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetCommentsArgs) Unmarshal(in []byte) error {
	msg := new(core_api.GetCommentsReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetCommentsArgs_Req_DEFAULT *core_api.GetCommentsReq

func (p *GetCommentsArgs) GetReq() *core_api.GetCommentsReq {
	if !p.IsSetReq() {
		return GetCommentsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetCommentsArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetCommentsArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetCommentsResult struct {
	Success *core_api.GetCommentsResp
}

var GetCommentsResult_Success_DEFAULT *core_api.GetCommentsResp

func (p *GetCommentsResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.GetCommentsResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetCommentsResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetCommentsResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetCommentsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetCommentsResult) Unmarshal(in []byte) error {
	msg := new(core_api.GetCommentsResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetCommentsResult) GetSuccess() *core_api.GetCommentsResp {
	if !p.IsSetSuccess() {
		return GetCommentsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetCommentsResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.GetCommentsResp)
}

func (p *GetCommentsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetCommentsResult) GetResult() interface{} {
	return p.Success
}

func newCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.NewCommentReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Comment).NewComment(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *NewCommentArgs:
		success, err := handler.(core_api.Comment).NewComment(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*NewCommentResult)
		realResult.Success = success
	}
	return nil
}
func newNewCommentArgs() interface{} {
	return &NewCommentArgs{}
}

func newNewCommentResult() interface{} {
	return &NewCommentResult{}
}

type NewCommentArgs struct {
	Req *core_api.NewCommentReq
}

func (p *NewCommentArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.NewCommentReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *NewCommentArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *NewCommentArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *NewCommentArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *NewCommentArgs) Unmarshal(in []byte) error {
	msg := new(core_api.NewCommentReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var NewCommentArgs_Req_DEFAULT *core_api.NewCommentReq

func (p *NewCommentArgs) GetReq() *core_api.NewCommentReq {
	if !p.IsSetReq() {
		return NewCommentArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *NewCommentArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *NewCommentArgs) GetFirstArgument() interface{} {
	return p.Req
}

type NewCommentResult struct {
	Success *core_api.NewCommentResp
}

var NewCommentResult_Success_DEFAULT *core_api.NewCommentResp

func (p *NewCommentResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.NewCommentResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *NewCommentResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *NewCommentResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *NewCommentResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *NewCommentResult) Unmarshal(in []byte) error {
	msg := new(core_api.NewCommentResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *NewCommentResult) GetSuccess() *core_api.NewCommentResp {
	if !p.IsSetSuccess() {
		return NewCommentResult_Success_DEFAULT
	}
	return p.Success
}

func (p *NewCommentResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.NewCommentResp)
}

func (p *NewCommentResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *NewCommentResult) GetResult() interface{} {
	return p.Success
}

func deleteCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.DeleteCommentReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Comment).DeleteComment(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *DeleteCommentArgs:
		success, err := handler.(core_api.Comment).DeleteComment(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteCommentResult)
		realResult.Success = success
	}
	return nil
}
func newDeleteCommentArgs() interface{} {
	return &DeleteCommentArgs{}
}

func newDeleteCommentResult() interface{} {
	return &DeleteCommentResult{}
}

type DeleteCommentArgs struct {
	Req *core_api.DeleteCommentReq
}

func (p *DeleteCommentArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.DeleteCommentReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeleteCommentArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeleteCommentArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeleteCommentArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteCommentArgs) Unmarshal(in []byte) error {
	msg := new(core_api.DeleteCommentReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteCommentArgs_Req_DEFAULT *core_api.DeleteCommentReq

func (p *DeleteCommentArgs) GetReq() *core_api.DeleteCommentReq {
	if !p.IsSetReq() {
		return DeleteCommentArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteCommentArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DeleteCommentArgs) GetFirstArgument() interface{} {
	return p.Req
}

type DeleteCommentResult struct {
	Success *core_api.DeleteCommentResp
}

var DeleteCommentResult_Success_DEFAULT *core_api.DeleteCommentResp

func (p *DeleteCommentResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.DeleteCommentResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeleteCommentResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeleteCommentResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeleteCommentResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteCommentResult) Unmarshal(in []byte) error {
	msg := new(core_api.DeleteCommentResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteCommentResult) GetSuccess() *core_api.DeleteCommentResp {
	if !p.IsSetSuccess() {
		return DeleteCommentResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteCommentResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.DeleteCommentResp)
}

func (p *DeleteCommentResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteCommentResult) GetResult() interface{} {
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

func (p *kClient) GetComments(ctx context.Context, Req *core_api.GetCommentsReq) (r *core_api.GetCommentsResp, err error) {
	var _args GetCommentsArgs
	_args.Req = Req
	var _result GetCommentsResult
	if err = p.c.Call(ctx, "GetComments", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) NewComment(ctx context.Context, Req *core_api.NewCommentReq) (r *core_api.NewCommentResp, err error) {
	var _args NewCommentArgs
	_args.Req = Req
	var _result NewCommentResult
	if err = p.c.Call(ctx, "NewComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteComment(ctx context.Context, Req *core_api.DeleteCommentReq) (r *core_api.DeleteCommentResp, err error) {
	var _args DeleteCommentArgs
	_args.Req = Req
	var _result DeleteCommentResult
	if err = p.c.Call(ctx, "DeleteComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
