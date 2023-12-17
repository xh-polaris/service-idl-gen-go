// Code generated by Kitex v0.8.0. DO NOT EDIT.

package incentive

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	core_api "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/core_api"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return incentiveServiceInfo
}

var incentiveServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "incentive"
	handlerType := (*core_api.Incentive)(nil)
	methods := map[string]kitex.MethodInfo{
		"CheckIn":    kitex.NewMethodInfo(checkInHandler, newCheckInArgs, newCheckInResult, false),
		"GetMission": kitex.NewMethodInfo(getMissionHandler, newGetMissionArgs, newGetMissionResult, false),
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

func checkInHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.CheckInReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Incentive).CheckIn(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CheckInArgs:
		success, err := handler.(core_api.Incentive).CheckIn(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CheckInResult)
		realResult.Success = success
	}
	return nil
}
func newCheckInArgs() interface{} {
	return &CheckInArgs{}
}

func newCheckInResult() interface{} {
	return &CheckInResult{}
}

type CheckInArgs struct {
	Req *core_api.CheckInReq
}

func (p *CheckInArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.CheckInReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CheckInArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CheckInArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CheckInArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CheckInArgs) Unmarshal(in []byte) error {
	msg := new(core_api.CheckInReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CheckInArgs_Req_DEFAULT *core_api.CheckInReq

func (p *CheckInArgs) GetReq() *core_api.CheckInReq {
	if !p.IsSetReq() {
		return CheckInArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CheckInArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CheckInArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CheckInResult struct {
	Success *core_api.CheckInResp
}

var CheckInResult_Success_DEFAULT *core_api.CheckInResp

func (p *CheckInResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.CheckInResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CheckInResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CheckInResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CheckInResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CheckInResult) Unmarshal(in []byte) error {
	msg := new(core_api.CheckInResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckInResult) GetSuccess() *core_api.CheckInResp {
	if !p.IsSetSuccess() {
		return CheckInResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CheckInResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.CheckInResp)
}

func (p *CheckInResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CheckInResult) GetResult() interface{} {
	return p.Success
}

func getMissionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.GetMissionReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Incentive).GetMission(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetMissionArgs:
		success, err := handler.(core_api.Incentive).GetMission(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetMissionResult)
		realResult.Success = success
	}
	return nil
}
func newGetMissionArgs() interface{} {
	return &GetMissionArgs{}
}

func newGetMissionResult() interface{} {
	return &GetMissionResult{}
}

type GetMissionArgs struct {
	Req *core_api.GetMissionReq
}

func (p *GetMissionArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.GetMissionReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetMissionArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetMissionArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetMissionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetMissionArgs) Unmarshal(in []byte) error {
	msg := new(core_api.GetMissionReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetMissionArgs_Req_DEFAULT *core_api.GetMissionReq

func (p *GetMissionArgs) GetReq() *core_api.GetMissionReq {
	if !p.IsSetReq() {
		return GetMissionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetMissionArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetMissionArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetMissionResult struct {
	Success *core_api.GetMissionResp
}

var GetMissionResult_Success_DEFAULT *core_api.GetMissionResp

func (p *GetMissionResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.GetMissionResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetMissionResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetMissionResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetMissionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetMissionResult) Unmarshal(in []byte) error {
	msg := new(core_api.GetMissionResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetMissionResult) GetSuccess() *core_api.GetMissionResp {
	if !p.IsSetSuccess() {
		return GetMissionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetMissionResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.GetMissionResp)
}

func (p *GetMissionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetMissionResult) GetResult() interface{} {
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

func (p *kClient) CheckIn(ctx context.Context, Req *core_api.CheckInReq) (r *core_api.CheckInResp, err error) {
	var _args CheckInArgs
	_args.Req = Req
	var _result CheckInResult
	if err = p.c.Call(ctx, "CheckIn", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetMission(ctx context.Context, Req *core_api.GetMissionReq) (r *core_api.GetMissionResp, err error) {
	var _args GetMissionArgs
	_args.Req = Req
	var _result GetMissionResult
	if err = p.c.Call(ctx, "GetMission", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
