// Code generated by Kitex v0.11.3. DO NOT EDIT.

package dataservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	data "github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/data"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Insert": kitex.NewMethodInfo(
		insertHandler,
		newInsertArgs,
		newInsertResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	dataServiceServiceInfo                = NewServiceInfo()
	dataServiceServiceInfoForClient       = NewServiceInfoForClient()
	dataServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return dataServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return dataServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return dataServiceServiceInfoForClient
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
	serviceName := "DataService"
	handlerType := (*data.DataService)(nil)
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
		"PackageName": "platform.data",
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

func insertHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(data.InsertReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(data.DataService).Insert(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *InsertArgs:
		success, err := handler.(data.DataService).Insert(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*InsertResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newInsertArgs() interface{} {
	return &InsertArgs{}
}

func newInsertResult() interface{} {
	return &InsertResult{}
}

type InsertArgs struct {
	Req *data.InsertReq
}

func (p *InsertArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(data.InsertReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *InsertArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *InsertArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *InsertArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *InsertArgs) Unmarshal(in []byte) error {
	msg := new(data.InsertReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var InsertArgs_Req_DEFAULT *data.InsertReq

func (p *InsertArgs) GetReq() *data.InsertReq {
	if !p.IsSetReq() {
		return InsertArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *InsertArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *InsertArgs) GetFirstArgument() interface{} {
	return p.Req
}

type InsertResult struct {
	Success *data.InsertResp
}

var InsertResult_Success_DEFAULT *data.InsertResp

func (p *InsertResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(data.InsertResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *InsertResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *InsertResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *InsertResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *InsertResult) Unmarshal(in []byte) error {
	msg := new(data.InsertResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *InsertResult) GetSuccess() *data.InsertResp {
	if !p.IsSetSuccess() {
		return InsertResult_Success_DEFAULT
	}
	return p.Success
}

func (p *InsertResult) SetSuccess(x interface{}) {
	p.Success = x.(*data.InsertResp)
}

func (p *InsertResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *InsertResult) GetResult() interface{} {
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

func (p *kClient) Insert(ctx context.Context, Req *data.InsertReq) (r *data.InsertResp, err error) {
	var _args InsertArgs
	_args.Req = Req
	var _result InsertResult
	if err = p.c.Call(ctx, "Insert", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
