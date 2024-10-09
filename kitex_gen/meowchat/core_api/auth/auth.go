// Code generated by Kitex v0.11.3. DO NOT EDIT.

package auth

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
	"SignIn": kitex.NewMethodInfo(
		signInHandler,
		newSignInArgs,
		newSignInResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"SendVerifyCode": kitex.NewMethodInfo(
		sendVerifyCodeHandler,
		newSendVerifyCodeArgs,
		newSendVerifyCodeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"SetPassword": kitex.NewMethodInfo(
		setPasswordHandler,
		newSetPasswordArgs,
		newSetPasswordResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	authServiceInfo                = NewServiceInfo()
	authServiceInfoForClient       = NewServiceInfoForClient()
	authServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return authServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return authServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return authServiceInfoForClient
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
	serviceName := "auth"
	handlerType := (*core_api.Auth)(nil)
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
		KiteXGenVersion: "v0.11.3",
		Extra:           extra,
	}
	return svcInfo
}

func signInHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.SignInReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Auth).SignIn(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *SignInArgs:
		success, err := handler.(core_api.Auth).SignIn(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SignInResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newSignInArgs() interface{} {
	return &SignInArgs{}
}

func newSignInResult() interface{} {
	return &SignInResult{}
}

type SignInArgs struct {
	Req *core_api.SignInReq
}

func (p *SignInArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.SignInReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *SignInArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *SignInArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *SignInArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *SignInArgs) Unmarshal(in []byte) error {
	msg := new(core_api.SignInReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SignInArgs_Req_DEFAULT *core_api.SignInReq

func (p *SignInArgs) GetReq() *core_api.SignInReq {
	if !p.IsSetReq() {
		return SignInArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SignInArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SignInArgs) GetFirstArgument() interface{} {
	return p.Req
}

type SignInResult struct {
	Success *core_api.SignInResp
}

var SignInResult_Success_DEFAULT *core_api.SignInResp

func (p *SignInResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.SignInResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *SignInResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *SignInResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *SignInResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *SignInResult) Unmarshal(in []byte) error {
	msg := new(core_api.SignInResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SignInResult) GetSuccess() *core_api.SignInResp {
	if !p.IsSetSuccess() {
		return SignInResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SignInResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.SignInResp)
}

func (p *SignInResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SignInResult) GetResult() interface{} {
	return p.Success
}

func sendVerifyCodeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.SendVerifyCodeReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Auth).SendVerifyCode(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *SendVerifyCodeArgs:
		success, err := handler.(core_api.Auth).SendVerifyCode(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SendVerifyCodeResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newSendVerifyCodeArgs() interface{} {
	return &SendVerifyCodeArgs{}
}

func newSendVerifyCodeResult() interface{} {
	return &SendVerifyCodeResult{}
}

type SendVerifyCodeArgs struct {
	Req *core_api.SendVerifyCodeReq
}

func (p *SendVerifyCodeArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.SendVerifyCodeReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *SendVerifyCodeArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *SendVerifyCodeArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *SendVerifyCodeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *SendVerifyCodeArgs) Unmarshal(in []byte) error {
	msg := new(core_api.SendVerifyCodeReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SendVerifyCodeArgs_Req_DEFAULT *core_api.SendVerifyCodeReq

func (p *SendVerifyCodeArgs) GetReq() *core_api.SendVerifyCodeReq {
	if !p.IsSetReq() {
		return SendVerifyCodeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SendVerifyCodeArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SendVerifyCodeArgs) GetFirstArgument() interface{} {
	return p.Req
}

type SendVerifyCodeResult struct {
	Success *core_api.SendVerifyCodeResp
}

var SendVerifyCodeResult_Success_DEFAULT *core_api.SendVerifyCodeResp

func (p *SendVerifyCodeResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.SendVerifyCodeResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *SendVerifyCodeResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *SendVerifyCodeResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *SendVerifyCodeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *SendVerifyCodeResult) Unmarshal(in []byte) error {
	msg := new(core_api.SendVerifyCodeResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SendVerifyCodeResult) GetSuccess() *core_api.SendVerifyCodeResp {
	if !p.IsSetSuccess() {
		return SendVerifyCodeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SendVerifyCodeResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.SendVerifyCodeResp)
}

func (p *SendVerifyCodeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SendVerifyCodeResult) GetResult() interface{} {
	return p.Success
}

func setPasswordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.SetPasswordReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Auth).SetPassword(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *SetPasswordArgs:
		success, err := handler.(core_api.Auth).SetPassword(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SetPasswordResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newSetPasswordArgs() interface{} {
	return &SetPasswordArgs{}
}

func newSetPasswordResult() interface{} {
	return &SetPasswordResult{}
}

type SetPasswordArgs struct {
	Req *core_api.SetPasswordReq
}

func (p *SetPasswordArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.SetPasswordReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *SetPasswordArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *SetPasswordArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *SetPasswordArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *SetPasswordArgs) Unmarshal(in []byte) error {
	msg := new(core_api.SetPasswordReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SetPasswordArgs_Req_DEFAULT *core_api.SetPasswordReq

func (p *SetPasswordArgs) GetReq() *core_api.SetPasswordReq {
	if !p.IsSetReq() {
		return SetPasswordArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SetPasswordArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SetPasswordArgs) GetFirstArgument() interface{} {
	return p.Req
}

type SetPasswordResult struct {
	Success *core_api.SetPasswordResp
}

var SetPasswordResult_Success_DEFAULT *core_api.SetPasswordResp

func (p *SetPasswordResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.SetPasswordResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *SetPasswordResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *SetPasswordResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *SetPasswordResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *SetPasswordResult) Unmarshal(in []byte) error {
	msg := new(core_api.SetPasswordResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SetPasswordResult) GetSuccess() *core_api.SetPasswordResp {
	if !p.IsSetSuccess() {
		return SetPasswordResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SetPasswordResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.SetPasswordResp)
}

func (p *SetPasswordResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SetPasswordResult) GetResult() interface{} {
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

func (p *kClient) SignIn(ctx context.Context, Req *core_api.SignInReq) (r *core_api.SignInResp, err error) {
	var _args SignInArgs
	_args.Req = Req
	var _result SignInResult
	if err = p.c.Call(ctx, "SignIn", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SendVerifyCode(ctx context.Context, Req *core_api.SendVerifyCodeReq) (r *core_api.SendVerifyCodeResp, err error) {
	var _args SendVerifyCodeArgs
	_args.Req = Req
	var _result SendVerifyCodeResult
	if err = p.c.Call(ctx, "SendVerifyCode", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SetPassword(ctx context.Context, Req *core_api.SetPasswordReq) (r *core_api.SetPasswordResp, err error) {
	var _args SetPasswordArgs
	_args.Req = Req
	var _result SetPasswordResult
	if err = p.c.Call(ctx, "SetPassword", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
