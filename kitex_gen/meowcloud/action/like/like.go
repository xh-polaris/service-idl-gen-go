// Code generated by Kitex v0.10.3. DO NOT EDIT.

package like

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	action "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowcloud/action"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"DoLike": kitex.NewMethodInfo(
		doLikeHandler,
		newDoLikeArgs,
		newDoLikeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"CancelLike": kitex.NewMethodInfo(
		cancelLikeHandler,
		newCancelLikeArgs,
		newCancelLikeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetLikedCount": kitex.NewMethodInfo(
		getLikedCountHandler,
		newGetLikedCountArgs,
		newGetLikedCountResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetLikedUsers": kitex.NewMethodInfo(
		getLikedUsersHandler,
		newGetLikedUsersArgs,
		newGetLikedUsersResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetUserLiked": kitex.NewMethodInfo(
		getUserLikedHandler,
		newGetUserLikedArgs,
		newGetUserLikedResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetLiked": kitex.NewMethodInfo(
		getLikedHandler,
		newGetLikedArgs,
		newGetLikedResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	likeServiceInfo                = NewServiceInfo()
	likeServiceInfoForClient       = NewServiceInfoForClient()
	likeServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return likeServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return likeServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return likeServiceInfoForClient
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
	serviceName := "like"
	handlerType := (*action.Like)(nil)
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
		"PackageName": "meowcloud.action",
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

func doLikeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(action.DoLikeReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(action.Like).DoLike(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *DoLikeArgs:
		success, err := handler.(action.Like).DoLike(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DoLikeResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newDoLikeArgs() interface{} {
	return &DoLikeArgs{}
}

func newDoLikeResult() interface{} {
	return &DoLikeResult{}
}

type DoLikeArgs struct {
	Req *action.DoLikeReq
}

func (p *DoLikeArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(action.DoLikeReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DoLikeArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DoLikeArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DoLikeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *DoLikeArgs) Unmarshal(in []byte) error {
	msg := new(action.DoLikeReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DoLikeArgs_Req_DEFAULT *action.DoLikeReq

func (p *DoLikeArgs) GetReq() *action.DoLikeReq {
	if !p.IsSetReq() {
		return DoLikeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DoLikeArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DoLikeArgs) GetFirstArgument() interface{} {
	return p.Req
}

type DoLikeResult struct {
	Success *action.DoLikeResp
}

var DoLikeResult_Success_DEFAULT *action.DoLikeResp

func (p *DoLikeResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(action.DoLikeResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DoLikeResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DoLikeResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DoLikeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *DoLikeResult) Unmarshal(in []byte) error {
	msg := new(action.DoLikeResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DoLikeResult) GetSuccess() *action.DoLikeResp {
	if !p.IsSetSuccess() {
		return DoLikeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DoLikeResult) SetSuccess(x interface{}) {
	p.Success = x.(*action.DoLikeResp)
}

func (p *DoLikeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DoLikeResult) GetResult() interface{} {
	return p.Success
}

func cancelLikeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(action.CancelLikeReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(action.Like).CancelLike(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *CancelLikeArgs:
		success, err := handler.(action.Like).CancelLike(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CancelLikeResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newCancelLikeArgs() interface{} {
	return &CancelLikeArgs{}
}

func newCancelLikeResult() interface{} {
	return &CancelLikeResult{}
}

type CancelLikeArgs struct {
	Req *action.CancelLikeReq
}

func (p *CancelLikeArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(action.CancelLikeReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CancelLikeArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CancelLikeArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CancelLikeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CancelLikeArgs) Unmarshal(in []byte) error {
	msg := new(action.CancelLikeReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CancelLikeArgs_Req_DEFAULT *action.CancelLikeReq

func (p *CancelLikeArgs) GetReq() *action.CancelLikeReq {
	if !p.IsSetReq() {
		return CancelLikeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CancelLikeArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CancelLikeArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CancelLikeResult struct {
	Success *action.CancelLikeResp
}

var CancelLikeResult_Success_DEFAULT *action.CancelLikeResp

func (p *CancelLikeResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(action.CancelLikeResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CancelLikeResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CancelLikeResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CancelLikeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CancelLikeResult) Unmarshal(in []byte) error {
	msg := new(action.CancelLikeResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CancelLikeResult) GetSuccess() *action.CancelLikeResp {
	if !p.IsSetSuccess() {
		return CancelLikeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CancelLikeResult) SetSuccess(x interface{}) {
	p.Success = x.(*action.CancelLikeResp)
}

func (p *CancelLikeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CancelLikeResult) GetResult() interface{} {
	return p.Success
}

func getLikedCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(action.GetLikedCountReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(action.Like).GetLikedCount(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetLikedCountArgs:
		success, err := handler.(action.Like).GetLikedCount(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetLikedCountResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetLikedCountArgs() interface{} {
	return &GetLikedCountArgs{}
}

func newGetLikedCountResult() interface{} {
	return &GetLikedCountResult{}
}

type GetLikedCountArgs struct {
	Req *action.GetLikedCountReq
}

func (p *GetLikedCountArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(action.GetLikedCountReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetLikedCountArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetLikedCountArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetLikedCountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetLikedCountArgs) Unmarshal(in []byte) error {
	msg := new(action.GetLikedCountReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetLikedCountArgs_Req_DEFAULT *action.GetLikedCountReq

func (p *GetLikedCountArgs) GetReq() *action.GetLikedCountReq {
	if !p.IsSetReq() {
		return GetLikedCountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetLikedCountArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetLikedCountArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetLikedCountResult struct {
	Success *action.GetLikedCountResp
}

var GetLikedCountResult_Success_DEFAULT *action.GetLikedCountResp

func (p *GetLikedCountResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(action.GetLikedCountResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetLikedCountResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetLikedCountResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetLikedCountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetLikedCountResult) Unmarshal(in []byte) error {
	msg := new(action.GetLikedCountResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetLikedCountResult) GetSuccess() *action.GetLikedCountResp {
	if !p.IsSetSuccess() {
		return GetLikedCountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetLikedCountResult) SetSuccess(x interface{}) {
	p.Success = x.(*action.GetLikedCountResp)
}

func (p *GetLikedCountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetLikedCountResult) GetResult() interface{} {
	return p.Success
}

func getLikedUsersHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(action.GetLikedUsersReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(action.Like).GetLikedUsers(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetLikedUsersArgs:
		success, err := handler.(action.Like).GetLikedUsers(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetLikedUsersResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetLikedUsersArgs() interface{} {
	return &GetLikedUsersArgs{}
}

func newGetLikedUsersResult() interface{} {
	return &GetLikedUsersResult{}
}

type GetLikedUsersArgs struct {
	Req *action.GetLikedUsersReq
}

func (p *GetLikedUsersArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(action.GetLikedUsersReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetLikedUsersArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetLikedUsersArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetLikedUsersArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetLikedUsersArgs) Unmarshal(in []byte) error {
	msg := new(action.GetLikedUsersReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetLikedUsersArgs_Req_DEFAULT *action.GetLikedUsersReq

func (p *GetLikedUsersArgs) GetReq() *action.GetLikedUsersReq {
	if !p.IsSetReq() {
		return GetLikedUsersArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetLikedUsersArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetLikedUsersArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetLikedUsersResult struct {
	Success *action.GetLikedUsersResp
}

var GetLikedUsersResult_Success_DEFAULT *action.GetLikedUsersResp

func (p *GetLikedUsersResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(action.GetLikedUsersResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetLikedUsersResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetLikedUsersResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetLikedUsersResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetLikedUsersResult) Unmarshal(in []byte) error {
	msg := new(action.GetLikedUsersResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetLikedUsersResult) GetSuccess() *action.GetLikedUsersResp {
	if !p.IsSetSuccess() {
		return GetLikedUsersResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetLikedUsersResult) SetSuccess(x interface{}) {
	p.Success = x.(*action.GetLikedUsersResp)
}

func (p *GetLikedUsersResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetLikedUsersResult) GetResult() interface{} {
	return p.Success
}

func getUserLikedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(action.GetUserLikedReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(action.Like).GetUserLiked(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetUserLikedArgs:
		success, err := handler.(action.Like).GetUserLiked(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetUserLikedResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetUserLikedArgs() interface{} {
	return &GetUserLikedArgs{}
}

func newGetUserLikedResult() interface{} {
	return &GetUserLikedResult{}
}

type GetUserLikedArgs struct {
	Req *action.GetUserLikedReq
}

func (p *GetUserLikedArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(action.GetUserLikedReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetUserLikedArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetUserLikedArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetUserLikedArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetUserLikedArgs) Unmarshal(in []byte) error {
	msg := new(action.GetUserLikedReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetUserLikedArgs_Req_DEFAULT *action.GetUserLikedReq

func (p *GetUserLikedArgs) GetReq() *action.GetUserLikedReq {
	if !p.IsSetReq() {
		return GetUserLikedArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserLikedArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetUserLikedArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetUserLikedResult struct {
	Success *action.GetUserLikedResp
}

var GetUserLikedResult_Success_DEFAULT *action.GetUserLikedResp

func (p *GetUserLikedResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(action.GetUserLikedResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetUserLikedResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetUserLikedResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetUserLikedResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetUserLikedResult) Unmarshal(in []byte) error {
	msg := new(action.GetUserLikedResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserLikedResult) GetSuccess() *action.GetUserLikedResp {
	if !p.IsSetSuccess() {
		return GetUserLikedResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserLikedResult) SetSuccess(x interface{}) {
	p.Success = x.(*action.GetUserLikedResp)
}

func (p *GetUserLikedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserLikedResult) GetResult() interface{} {
	return p.Success
}

func getLikedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(action.GetLikedReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(action.Like).GetLiked(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetLikedArgs:
		success, err := handler.(action.Like).GetLiked(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetLikedResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetLikedArgs() interface{} {
	return &GetLikedArgs{}
}

func newGetLikedResult() interface{} {
	return &GetLikedResult{}
}

type GetLikedArgs struct {
	Req *action.GetLikedReq
}

func (p *GetLikedArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(action.GetLikedReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetLikedArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetLikedArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetLikedArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetLikedArgs) Unmarshal(in []byte) error {
	msg := new(action.GetLikedReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetLikedArgs_Req_DEFAULT *action.GetLikedReq

func (p *GetLikedArgs) GetReq() *action.GetLikedReq {
	if !p.IsSetReq() {
		return GetLikedArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetLikedArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetLikedArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetLikedResult struct {
	Success *action.GetLikedResp
}

var GetLikedResult_Success_DEFAULT *action.GetLikedResp

func (p *GetLikedResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(action.GetLikedResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetLikedResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetLikedResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetLikedResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetLikedResult) Unmarshal(in []byte) error {
	msg := new(action.GetLikedResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetLikedResult) GetSuccess() *action.GetLikedResp {
	if !p.IsSetSuccess() {
		return GetLikedResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetLikedResult) SetSuccess(x interface{}) {
	p.Success = x.(*action.GetLikedResp)
}

func (p *GetLikedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetLikedResult) GetResult() interface{} {
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

func (p *kClient) DoLike(ctx context.Context, Req *action.DoLikeReq) (r *action.DoLikeResp, err error) {
	var _args DoLikeArgs
	_args.Req = Req
	var _result DoLikeResult
	if err = p.c.Call(ctx, "DoLike", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CancelLike(ctx context.Context, Req *action.CancelLikeReq) (r *action.CancelLikeResp, err error) {
	var _args CancelLikeArgs
	_args.Req = Req
	var _result CancelLikeResult
	if err = p.c.Call(ctx, "CancelLike", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetLikedCount(ctx context.Context, Req *action.GetLikedCountReq) (r *action.GetLikedCountResp, err error) {
	var _args GetLikedCountArgs
	_args.Req = Req
	var _result GetLikedCountResult
	if err = p.c.Call(ctx, "GetLikedCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetLikedUsers(ctx context.Context, Req *action.GetLikedUsersReq) (r *action.GetLikedUsersResp, err error) {
	var _args GetLikedUsersArgs
	_args.Req = Req
	var _result GetLikedUsersResult
	if err = p.c.Call(ctx, "GetLikedUsers", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserLiked(ctx context.Context, Req *action.GetUserLikedReq) (r *action.GetUserLikedResp, err error) {
	var _args GetUserLikedArgs
	_args.Req = Req
	var _result GetUserLikedResult
	if err = p.c.Call(ctx, "GetUserLiked", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetLiked(ctx context.Context, Req *action.GetLikedReq) (r *action.GetLikedResp, err error) {
	var _args GetLikedArgs
	_args.Req = Req
	var _result GetLikedResult
	if err = p.c.Call(ctx, "GetLiked", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
