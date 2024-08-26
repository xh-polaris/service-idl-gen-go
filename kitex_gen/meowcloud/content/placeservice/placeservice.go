// Code generated by Kitex v0.10.3. DO NOT EDIT.

package placeservice

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
	"AddPlace": kitex.NewMethodInfo(
		addPlaceHandler,
		newAddPlaceArgs,
		newAddPlaceResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"UpdatePlace": kitex.NewMethodInfo(
		updatePlaceHandler,
		newUpdatePlaceArgs,
		newUpdatePlaceResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetPlace": kitex.NewMethodInfo(
		getPlaceHandler,
		newGetPlaceArgs,
		newGetPlaceResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"DeletePlace": kitex.NewMethodInfo(
		deletePlaceHandler,
		newDeletePlaceArgs,
		newDeletePlaceResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	placeServiceServiceInfo                = NewServiceInfo()
	placeServiceServiceInfoForClient       = NewServiceInfoForClient()
	placeServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return placeServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return placeServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return placeServiceServiceInfoForClient
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
	serviceName := "PlaceService"
	handlerType := (*content.PlaceService)(nil)
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

func addPlaceHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(content.AddPlaceReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(content.PlaceService).AddPlace(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *AddPlaceArgs:
		success, err := handler.(content.PlaceService).AddPlace(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*AddPlaceResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newAddPlaceArgs() interface{} {
	return &AddPlaceArgs{}
}

func newAddPlaceResult() interface{} {
	return &AddPlaceResult{}
}

type AddPlaceArgs struct {
	Req *content.AddPlaceReq
}

func (p *AddPlaceArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(content.AddPlaceReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *AddPlaceArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *AddPlaceArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *AddPlaceArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *AddPlaceArgs) Unmarshal(in []byte) error {
	msg := new(content.AddPlaceReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var AddPlaceArgs_Req_DEFAULT *content.AddPlaceReq

func (p *AddPlaceArgs) GetReq() *content.AddPlaceReq {
	if !p.IsSetReq() {
		return AddPlaceArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AddPlaceArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *AddPlaceArgs) GetFirstArgument() interface{} {
	return p.Req
}

type AddPlaceResult struct {
	Success *content.AddPlaceResp
}

var AddPlaceResult_Success_DEFAULT *content.AddPlaceResp

func (p *AddPlaceResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(content.AddPlaceResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *AddPlaceResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *AddPlaceResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *AddPlaceResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *AddPlaceResult) Unmarshal(in []byte) error {
	msg := new(content.AddPlaceResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AddPlaceResult) GetSuccess() *content.AddPlaceResp {
	if !p.IsSetSuccess() {
		return AddPlaceResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AddPlaceResult) SetSuccess(x interface{}) {
	p.Success = x.(*content.AddPlaceResp)
}

func (p *AddPlaceResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AddPlaceResult) GetResult() interface{} {
	return p.Success
}

func updatePlaceHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(content.UpdatePlaceReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(content.PlaceService).UpdatePlace(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *UpdatePlaceArgs:
		success, err := handler.(content.PlaceService).UpdatePlace(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdatePlaceResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newUpdatePlaceArgs() interface{} {
	return &UpdatePlaceArgs{}
}

func newUpdatePlaceResult() interface{} {
	return &UpdatePlaceResult{}
}

type UpdatePlaceArgs struct {
	Req *content.UpdatePlaceReq
}

func (p *UpdatePlaceArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(content.UpdatePlaceReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdatePlaceArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdatePlaceArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdatePlaceArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UpdatePlaceArgs) Unmarshal(in []byte) error {
	msg := new(content.UpdatePlaceReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdatePlaceArgs_Req_DEFAULT *content.UpdatePlaceReq

func (p *UpdatePlaceArgs) GetReq() *content.UpdatePlaceReq {
	if !p.IsSetReq() {
		return UpdatePlaceArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdatePlaceArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdatePlaceArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdatePlaceResult struct {
	Success *content.UpdatePlaceResp
}

var UpdatePlaceResult_Success_DEFAULT *content.UpdatePlaceResp

func (p *UpdatePlaceResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(content.UpdatePlaceResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdatePlaceResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdatePlaceResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdatePlaceResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UpdatePlaceResult) Unmarshal(in []byte) error {
	msg := new(content.UpdatePlaceResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdatePlaceResult) GetSuccess() *content.UpdatePlaceResp {
	if !p.IsSetSuccess() {
		return UpdatePlaceResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdatePlaceResult) SetSuccess(x interface{}) {
	p.Success = x.(*content.UpdatePlaceResp)
}

func (p *UpdatePlaceResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdatePlaceResult) GetResult() interface{} {
	return p.Success
}

func getPlaceHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(content.GetPlaceReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(content.PlaceService).GetPlace(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetPlaceArgs:
		success, err := handler.(content.PlaceService).GetPlace(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetPlaceResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetPlaceArgs() interface{} {
	return &GetPlaceArgs{}
}

func newGetPlaceResult() interface{} {
	return &GetPlaceResult{}
}

type GetPlaceArgs struct {
	Req *content.GetPlaceReq
}

func (p *GetPlaceArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(content.GetPlaceReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetPlaceArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetPlaceArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetPlaceArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetPlaceArgs) Unmarshal(in []byte) error {
	msg := new(content.GetPlaceReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetPlaceArgs_Req_DEFAULT *content.GetPlaceReq

func (p *GetPlaceArgs) GetReq() *content.GetPlaceReq {
	if !p.IsSetReq() {
		return GetPlaceArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetPlaceArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetPlaceArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetPlaceResult struct {
	Success *content.GetPlaceResp
}

var GetPlaceResult_Success_DEFAULT *content.GetPlaceResp

func (p *GetPlaceResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(content.GetPlaceResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetPlaceResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetPlaceResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetPlaceResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetPlaceResult) Unmarshal(in []byte) error {
	msg := new(content.GetPlaceResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPlaceResult) GetSuccess() *content.GetPlaceResp {
	if !p.IsSetSuccess() {
		return GetPlaceResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetPlaceResult) SetSuccess(x interface{}) {
	p.Success = x.(*content.GetPlaceResp)
}

func (p *GetPlaceResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetPlaceResult) GetResult() interface{} {
	return p.Success
}

func deletePlaceHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(content.DeletePlaceReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(content.PlaceService).DeletePlace(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *DeletePlaceArgs:
		success, err := handler.(content.PlaceService).DeletePlace(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeletePlaceResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newDeletePlaceArgs() interface{} {
	return &DeletePlaceArgs{}
}

func newDeletePlaceResult() interface{} {
	return &DeletePlaceResult{}
}

type DeletePlaceArgs struct {
	Req *content.DeletePlaceReq
}

func (p *DeletePlaceArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(content.DeletePlaceReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeletePlaceArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeletePlaceArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeletePlaceArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *DeletePlaceArgs) Unmarshal(in []byte) error {
	msg := new(content.DeletePlaceReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeletePlaceArgs_Req_DEFAULT *content.DeletePlaceReq

func (p *DeletePlaceArgs) GetReq() *content.DeletePlaceReq {
	if !p.IsSetReq() {
		return DeletePlaceArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeletePlaceArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DeletePlaceArgs) GetFirstArgument() interface{} {
	return p.Req
}

type DeletePlaceResult struct {
	Success *content.DeletePlaceResp
}

var DeletePlaceResult_Success_DEFAULT *content.DeletePlaceResp

func (p *DeletePlaceResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(content.DeletePlaceResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeletePlaceResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeletePlaceResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeletePlaceResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *DeletePlaceResult) Unmarshal(in []byte) error {
	msg := new(content.DeletePlaceResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeletePlaceResult) GetSuccess() *content.DeletePlaceResp {
	if !p.IsSetSuccess() {
		return DeletePlaceResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeletePlaceResult) SetSuccess(x interface{}) {
	p.Success = x.(*content.DeletePlaceResp)
}

func (p *DeletePlaceResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeletePlaceResult) GetResult() interface{} {
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

func (p *kClient) AddPlace(ctx context.Context, Req *content.AddPlaceReq) (r *content.AddPlaceResp, err error) {
	var _args AddPlaceArgs
	_args.Req = Req
	var _result AddPlaceResult
	if err = p.c.Call(ctx, "AddPlace", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdatePlace(ctx context.Context, Req *content.UpdatePlaceReq) (r *content.UpdatePlaceResp, err error) {
	var _args UpdatePlaceArgs
	_args.Req = Req
	var _result UpdatePlaceResult
	if err = p.c.Call(ctx, "UpdatePlace", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetPlace(ctx context.Context, Req *content.GetPlaceReq) (r *content.GetPlaceResp, err error) {
	var _args GetPlaceArgs
	_args.Req = Req
	var _result GetPlaceResult
	if err = p.c.Call(ctx, "GetPlace", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeletePlace(ctx context.Context, Req *content.DeletePlaceReq) (r *content.DeletePlaceResp, err error) {
	var _args DeletePlaceArgs
	_args.Req = Req
	var _result DeletePlaceResult
	if err = p.c.Call(ctx, "DeletePlace", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
