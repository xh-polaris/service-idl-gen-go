// Code generated by Kitex v0.12.1. DO NOT EDIT.

package data

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	core_api "github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/core_api"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"ReportEvent": kitex.NewMethodInfo(
		reportEventHandler,
		newReportEventArgs,
		newReportEventResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	dataServiceInfo                = NewServiceInfo()
	dataServiceInfoForClient       = NewServiceInfoForClient()
	dataServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return dataServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return dataServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return dataServiceInfoForClient
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
	serviceName := "data"
	handlerType := (*core_api.Data)(nil)
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
		"PackageName": "platform.core_api",
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

func reportEventHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.ReportEventRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Data).ReportEvent(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ReportEventArgs:
		success, err := handler.(core_api.Data).ReportEvent(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ReportEventResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newReportEventArgs() interface{} {
	return &ReportEventArgs{}
}

func newReportEventResult() interface{} {
	return &ReportEventResult{}
}

type ReportEventArgs struct {
	Req *core_api.ReportEventRequest
}

func (p *ReportEventArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.ReportEventRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ReportEventArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ReportEventArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ReportEventArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ReportEventArgs) Unmarshal(in []byte) error {
	msg := new(core_api.ReportEventRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ReportEventArgs_Req_DEFAULT *core_api.ReportEventRequest

func (p *ReportEventArgs) GetReq() *core_api.ReportEventRequest {
	if !p.IsSetReq() {
		return ReportEventArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ReportEventArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ReportEventArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ReportEventResult struct {
	Success *core_api.ReportEventResponse
}

var ReportEventResult_Success_DEFAULT *core_api.ReportEventResponse

func (p *ReportEventResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.ReportEventResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ReportEventResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ReportEventResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ReportEventResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ReportEventResult) Unmarshal(in []byte) error {
	msg := new(core_api.ReportEventResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ReportEventResult) GetSuccess() *core_api.ReportEventResponse {
	if !p.IsSetSuccess() {
		return ReportEventResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ReportEventResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.ReportEventResponse)
}

func (p *ReportEventResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ReportEventResult) GetResult() interface{} {
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

func (p *kClient) ReportEvent(ctx context.Context, Req *core_api.ReportEventRequest) (r *core_api.ReportEventResponse, err error) {
	var _args ReportEventArgs
	_args.Req = Req
	var _result ReportEventResult
	if err = p.c.Call(ctx, "ReportEvent", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
