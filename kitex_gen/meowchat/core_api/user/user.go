// Code generated by Kitex v0.9.1. DO NOT EDIT.

package user

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
	"GetUserInfo": kitex.NewMethodInfo(
		getUserInfoHandler,
		newGetUserInfoArgs,
		newGetUserInfoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"UpdateUserInfo": kitex.NewMethodInfo(
		updateUserInfoHandler,
		newUpdateUserInfoArgs,
		newUpdateUserInfoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"SearchUser": kitex.NewMethodInfo(
		searchUserHandler,
		newSearchUserArgs,
		newSearchUserResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	userServiceInfo                = NewServiceInfo()
	userServiceInfoForClient       = NewServiceInfoForClient()
	userServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return userServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return userServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return userServiceInfoForClient
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
	serviceName := "user"
	handlerType := (*core_api.User)(nil)
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
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func getUserInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.GetUserInfoReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.User).GetUserInfo(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetUserInfoArgs:
		success, err := handler.(core_api.User).GetUserInfo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetUserInfoResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetUserInfoArgs() interface{} {
	return &GetUserInfoArgs{}
}

func newGetUserInfoResult() interface{} {
	return &GetUserInfoResult{}
}

type GetUserInfoArgs struct {
	Req *core_api.GetUserInfoReq
}

func (p *GetUserInfoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.GetUserInfoReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetUserInfoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetUserInfoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetUserInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetUserInfoArgs) Unmarshal(in []byte) error {
	msg := new(core_api.GetUserInfoReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetUserInfoArgs_Req_DEFAULT *core_api.GetUserInfoReq

func (p *GetUserInfoArgs) GetReq() *core_api.GetUserInfoReq {
	if !p.IsSetReq() {
		return GetUserInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetUserInfoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetUserInfoResult struct {
	Success *core_api.GetUserInfoResp
}

var GetUserInfoResult_Success_DEFAULT *core_api.GetUserInfoResp

func (p *GetUserInfoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.GetUserInfoResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetUserInfoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetUserInfoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetUserInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetUserInfoResult) Unmarshal(in []byte) error {
	msg := new(core_api.GetUserInfoResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserInfoResult) GetSuccess() *core_api.GetUserInfoResp {
	if !p.IsSetSuccess() {
		return GetUserInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.GetUserInfoResp)
}

func (p *GetUserInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserInfoResult) GetResult() interface{} {
	return p.Success
}

func updateUserInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.UpdateUserInfoReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.User).UpdateUserInfo(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *UpdateUserInfoArgs:
		success, err := handler.(core_api.User).UpdateUserInfo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateUserInfoResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newUpdateUserInfoArgs() interface{} {
	return &UpdateUserInfoArgs{}
}

func newUpdateUserInfoResult() interface{} {
	return &UpdateUserInfoResult{}
}

type UpdateUserInfoArgs struct {
	Req *core_api.UpdateUserInfoReq
}

func (p *UpdateUserInfoArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.UpdateUserInfoReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateUserInfoArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateUserInfoArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateUserInfoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateUserInfoArgs) Unmarshal(in []byte) error {
	msg := new(core_api.UpdateUserInfoReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateUserInfoArgs_Req_DEFAULT *core_api.UpdateUserInfoReq

func (p *UpdateUserInfoArgs) GetReq() *core_api.UpdateUserInfoReq {
	if !p.IsSetReq() {
		return UpdateUserInfoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateUserInfoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdateUserInfoArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdateUserInfoResult struct {
	Success *core_api.UpdateUserInfoResp
}

var UpdateUserInfoResult_Success_DEFAULT *core_api.UpdateUserInfoResp

func (p *UpdateUserInfoResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.UpdateUserInfoResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateUserInfoResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateUserInfoResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateUserInfoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateUserInfoResult) Unmarshal(in []byte) error {
	msg := new(core_api.UpdateUserInfoResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateUserInfoResult) GetSuccess() *core_api.UpdateUserInfoResp {
	if !p.IsSetSuccess() {
		return UpdateUserInfoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateUserInfoResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.UpdateUserInfoResp)
}

func (p *UpdateUserInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateUserInfoResult) GetResult() interface{} {
	return p.Success
}

func searchUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.SearchUserReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.User).SearchUser(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *SearchUserArgs:
		success, err := handler.(core_api.User).SearchUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SearchUserResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newSearchUserArgs() interface{} {
	return &SearchUserArgs{}
}

func newSearchUserResult() interface{} {
	return &SearchUserResult{}
}

type SearchUserArgs struct {
	Req *core_api.SearchUserReq
}

func (p *SearchUserArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.SearchUserReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *SearchUserArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *SearchUserArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *SearchUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *SearchUserArgs) Unmarshal(in []byte) error {
	msg := new(core_api.SearchUserReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SearchUserArgs_Req_DEFAULT *core_api.SearchUserReq

func (p *SearchUserArgs) GetReq() *core_api.SearchUserReq {
	if !p.IsSetReq() {
		return SearchUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SearchUserArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SearchUserArgs) GetFirstArgument() interface{} {
	return p.Req
}

type SearchUserResult struct {
	Success *core_api.SearchUserResp
}

var SearchUserResult_Success_DEFAULT *core_api.SearchUserResp

func (p *SearchUserResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.SearchUserResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *SearchUserResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *SearchUserResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *SearchUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *SearchUserResult) Unmarshal(in []byte) error {
	msg := new(core_api.SearchUserResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SearchUserResult) GetSuccess() *core_api.SearchUserResp {
	if !p.IsSetSuccess() {
		return SearchUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SearchUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.SearchUserResp)
}

func (p *SearchUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SearchUserResult) GetResult() interface{} {
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

func (p *kClient) GetUserInfo(ctx context.Context, Req *core_api.GetUserInfoReq) (r *core_api.GetUserInfoResp, err error) {
	var _args GetUserInfoArgs
	_args.Req = Req
	var _result GetUserInfoResult
	if err = p.c.Call(ctx, "GetUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateUserInfo(ctx context.Context, Req *core_api.UpdateUserInfoReq) (r *core_api.UpdateUserInfoResp, err error) {
	var _args UpdateUserInfoArgs
	_args.Req = Req
	var _result UpdateUserInfoResult
	if err = p.c.Call(ctx, "UpdateUserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SearchUser(ctx context.Context, Req *core_api.SearchUserReq) (r *core_api.SearchUserResp, err error) {
	var _args SearchUserArgs
	_args.Req = Req
	var _result SearchUserResult
	if err = p.c.Call(ctx, "SearchUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
