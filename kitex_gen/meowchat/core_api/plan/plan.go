// Code generated by Kitex v0.7.1. DO NOT EDIT.

package plan

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	core_api "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/core_api"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return planServiceInfo
}

var planServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "plan"
	handlerType := (*core_api.Plan)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetPlanPreviews": kitex.NewMethodInfo(getPlanPreviewsHandler, newGetPlanPreviewsArgs, newGetPlanPreviewsResult, false),
		"GetPlanDetail":   kitex.NewMethodInfo(getPlanDetailHandler, newGetPlanDetailArgs, newGetPlanDetailResult, false),
		"NewPlan":         kitex.NewMethodInfo(newPlanHandler, newNewPlanArgs, newNewPlanResult, false),
		"DeletePlan":      kitex.NewMethodInfo(deletePlanHandler, newDeletePlanArgs, newDeletePlanResult, false),
		"SearchPlan":      kitex.NewMethodInfo(searchPlanHandler, newSearchPlanArgs, newSearchPlanResult, false),
		"DonateFish":      kitex.NewMethodInfo(donateFishHandler, newDonateFishArgs, newDonateFishResult, false),
		"GetUserFish":     kitex.NewMethodInfo(getUserFishHandler, newGetUserFishArgs, newGetUserFishResult, false),
		"ListFishByPlan":  kitex.NewMethodInfo(listFishByPlanHandler, newListFishByPlanArgs, newListFishByPlanResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "meowchat.core_api",
		"ServiceFilePath": "",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.7.1",
		Extra:           extra,
	}
	return svcInfo
}

func getPlanPreviewsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.GetPlanPreviewsReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Plan).GetPlanPreviews(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetPlanPreviewsArgs:
		success, err := handler.(core_api.Plan).GetPlanPreviews(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetPlanPreviewsResult)
		realResult.Success = success
	}
	return nil
}
func newGetPlanPreviewsArgs() interface{} {
	return &GetPlanPreviewsArgs{}
}

func newGetPlanPreviewsResult() interface{} {
	return &GetPlanPreviewsResult{}
}

type GetPlanPreviewsArgs struct {
	Req *core_api.GetPlanPreviewsReq
}

func (p *GetPlanPreviewsArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.GetPlanPreviewsReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetPlanPreviewsArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetPlanPreviewsArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetPlanPreviewsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetPlanPreviewsArgs) Unmarshal(in []byte) error {
	msg := new(core_api.GetPlanPreviewsReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetPlanPreviewsArgs_Req_DEFAULT *core_api.GetPlanPreviewsReq

func (p *GetPlanPreviewsArgs) GetReq() *core_api.GetPlanPreviewsReq {
	if !p.IsSetReq() {
		return GetPlanPreviewsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetPlanPreviewsArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetPlanPreviewsArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetPlanPreviewsResult struct {
	Success *core_api.GetPlanPreviewsResp
}

var GetPlanPreviewsResult_Success_DEFAULT *core_api.GetPlanPreviewsResp

func (p *GetPlanPreviewsResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.GetPlanPreviewsResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetPlanPreviewsResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetPlanPreviewsResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetPlanPreviewsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetPlanPreviewsResult) Unmarshal(in []byte) error {
	msg := new(core_api.GetPlanPreviewsResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPlanPreviewsResult) GetSuccess() *core_api.GetPlanPreviewsResp {
	if !p.IsSetSuccess() {
		return GetPlanPreviewsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetPlanPreviewsResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.GetPlanPreviewsResp)
}

func (p *GetPlanPreviewsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetPlanPreviewsResult) GetResult() interface{} {
	return p.Success
}

func getPlanDetailHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.GetPlanDetailReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Plan).GetPlanDetail(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetPlanDetailArgs:
		success, err := handler.(core_api.Plan).GetPlanDetail(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetPlanDetailResult)
		realResult.Success = success
	}
	return nil
}
func newGetPlanDetailArgs() interface{} {
	return &GetPlanDetailArgs{}
}

func newGetPlanDetailResult() interface{} {
	return &GetPlanDetailResult{}
}

type GetPlanDetailArgs struct {
	Req *core_api.GetPlanDetailReq
}

func (p *GetPlanDetailArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.GetPlanDetailReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetPlanDetailArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetPlanDetailArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetPlanDetailArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetPlanDetailArgs) Unmarshal(in []byte) error {
	msg := new(core_api.GetPlanDetailReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetPlanDetailArgs_Req_DEFAULT *core_api.GetPlanDetailReq

func (p *GetPlanDetailArgs) GetReq() *core_api.GetPlanDetailReq {
	if !p.IsSetReq() {
		return GetPlanDetailArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetPlanDetailArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetPlanDetailArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetPlanDetailResult struct {
	Success *core_api.GetPlanDetailResp
}

var GetPlanDetailResult_Success_DEFAULT *core_api.GetPlanDetailResp

func (p *GetPlanDetailResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.GetPlanDetailResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetPlanDetailResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetPlanDetailResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetPlanDetailResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetPlanDetailResult) Unmarshal(in []byte) error {
	msg := new(core_api.GetPlanDetailResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPlanDetailResult) GetSuccess() *core_api.GetPlanDetailResp {
	if !p.IsSetSuccess() {
		return GetPlanDetailResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetPlanDetailResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.GetPlanDetailResp)
}

func (p *GetPlanDetailResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetPlanDetailResult) GetResult() interface{} {
	return p.Success
}

func newPlanHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.NewPlanReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Plan).NewPlan(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *NewPlanArgs:
		success, err := handler.(core_api.Plan).NewPlan(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*NewPlanResult)
		realResult.Success = success
	}
	return nil
}
func newNewPlanArgs() interface{} {
	return &NewPlanArgs{}
}

func newNewPlanResult() interface{} {
	return &NewPlanResult{}
}

type NewPlanArgs struct {
	Req *core_api.NewPlanReq
}

func (p *NewPlanArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.NewPlanReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *NewPlanArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *NewPlanArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *NewPlanArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *NewPlanArgs) Unmarshal(in []byte) error {
	msg := new(core_api.NewPlanReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var NewPlanArgs_Req_DEFAULT *core_api.NewPlanReq

func (p *NewPlanArgs) GetReq() *core_api.NewPlanReq {
	if !p.IsSetReq() {
		return NewPlanArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *NewPlanArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *NewPlanArgs) GetFirstArgument() interface{} {
	return p.Req
}

type NewPlanResult struct {
	Success *core_api.NewPlanResp
}

var NewPlanResult_Success_DEFAULT *core_api.NewPlanResp

func (p *NewPlanResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.NewPlanResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *NewPlanResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *NewPlanResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *NewPlanResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *NewPlanResult) Unmarshal(in []byte) error {
	msg := new(core_api.NewPlanResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *NewPlanResult) GetSuccess() *core_api.NewPlanResp {
	if !p.IsSetSuccess() {
		return NewPlanResult_Success_DEFAULT
	}
	return p.Success
}

func (p *NewPlanResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.NewPlanResp)
}

func (p *NewPlanResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *NewPlanResult) GetResult() interface{} {
	return p.Success
}

func deletePlanHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.DeletePlanReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Plan).DeletePlan(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *DeletePlanArgs:
		success, err := handler.(core_api.Plan).DeletePlan(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeletePlanResult)
		realResult.Success = success
	}
	return nil
}
func newDeletePlanArgs() interface{} {
	return &DeletePlanArgs{}
}

func newDeletePlanResult() interface{} {
	return &DeletePlanResult{}
}

type DeletePlanArgs struct {
	Req *core_api.DeletePlanReq
}

func (p *DeletePlanArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.DeletePlanReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeletePlanArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeletePlanArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeletePlanArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *DeletePlanArgs) Unmarshal(in []byte) error {
	msg := new(core_api.DeletePlanReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeletePlanArgs_Req_DEFAULT *core_api.DeletePlanReq

func (p *DeletePlanArgs) GetReq() *core_api.DeletePlanReq {
	if !p.IsSetReq() {
		return DeletePlanArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeletePlanArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DeletePlanArgs) GetFirstArgument() interface{} {
	return p.Req
}

type DeletePlanResult struct {
	Success *core_api.DeletePlanResp
}

var DeletePlanResult_Success_DEFAULT *core_api.DeletePlanResp

func (p *DeletePlanResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.DeletePlanResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeletePlanResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeletePlanResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeletePlanResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *DeletePlanResult) Unmarshal(in []byte) error {
	msg := new(core_api.DeletePlanResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeletePlanResult) GetSuccess() *core_api.DeletePlanResp {
	if !p.IsSetSuccess() {
		return DeletePlanResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeletePlanResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.DeletePlanResp)
}

func (p *DeletePlanResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeletePlanResult) GetResult() interface{} {
	return p.Success
}

func searchPlanHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.SearchPlanReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Plan).SearchPlan(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *SearchPlanArgs:
		success, err := handler.(core_api.Plan).SearchPlan(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SearchPlanResult)
		realResult.Success = success
	}
	return nil
}
func newSearchPlanArgs() interface{} {
	return &SearchPlanArgs{}
}

func newSearchPlanResult() interface{} {
	return &SearchPlanResult{}
}

type SearchPlanArgs struct {
	Req *core_api.SearchPlanReq
}

func (p *SearchPlanArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.SearchPlanReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *SearchPlanArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *SearchPlanArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *SearchPlanArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *SearchPlanArgs) Unmarshal(in []byte) error {
	msg := new(core_api.SearchPlanReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SearchPlanArgs_Req_DEFAULT *core_api.SearchPlanReq

func (p *SearchPlanArgs) GetReq() *core_api.SearchPlanReq {
	if !p.IsSetReq() {
		return SearchPlanArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SearchPlanArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SearchPlanArgs) GetFirstArgument() interface{} {
	return p.Req
}

type SearchPlanResult struct {
	Success *core_api.SearchPlanResp
}

var SearchPlanResult_Success_DEFAULT *core_api.SearchPlanResp

func (p *SearchPlanResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.SearchPlanResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *SearchPlanResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *SearchPlanResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *SearchPlanResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *SearchPlanResult) Unmarshal(in []byte) error {
	msg := new(core_api.SearchPlanResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SearchPlanResult) GetSuccess() *core_api.SearchPlanResp {
	if !p.IsSetSuccess() {
		return SearchPlanResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SearchPlanResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.SearchPlanResp)
}

func (p *SearchPlanResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SearchPlanResult) GetResult() interface{} {
	return p.Success
}

func donateFishHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.DonateFishReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Plan).DonateFish(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *DonateFishArgs:
		success, err := handler.(core_api.Plan).DonateFish(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DonateFishResult)
		realResult.Success = success
	}
	return nil
}
func newDonateFishArgs() interface{} {
	return &DonateFishArgs{}
}

func newDonateFishResult() interface{} {
	return &DonateFishResult{}
}

type DonateFishArgs struct {
	Req *core_api.DonateFishReq
}

func (p *DonateFishArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.DonateFishReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DonateFishArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DonateFishArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DonateFishArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *DonateFishArgs) Unmarshal(in []byte) error {
	msg := new(core_api.DonateFishReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DonateFishArgs_Req_DEFAULT *core_api.DonateFishReq

func (p *DonateFishArgs) GetReq() *core_api.DonateFishReq {
	if !p.IsSetReq() {
		return DonateFishArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DonateFishArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DonateFishArgs) GetFirstArgument() interface{} {
	return p.Req
}

type DonateFishResult struct {
	Success *core_api.DonateFishResp
}

var DonateFishResult_Success_DEFAULT *core_api.DonateFishResp

func (p *DonateFishResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.DonateFishResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DonateFishResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DonateFishResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DonateFishResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *DonateFishResult) Unmarshal(in []byte) error {
	msg := new(core_api.DonateFishResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DonateFishResult) GetSuccess() *core_api.DonateFishResp {
	if !p.IsSetSuccess() {
		return DonateFishResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DonateFishResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.DonateFishResp)
}

func (p *DonateFishResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DonateFishResult) GetResult() interface{} {
	return p.Success
}

func getUserFishHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.GetUserFishReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Plan).GetUserFish(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetUserFishArgs:
		success, err := handler.(core_api.Plan).GetUserFish(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetUserFishResult)
		realResult.Success = success
	}
	return nil
}
func newGetUserFishArgs() interface{} {
	return &GetUserFishArgs{}
}

func newGetUserFishResult() interface{} {
	return &GetUserFishResult{}
}

type GetUserFishArgs struct {
	Req *core_api.GetUserFishReq
}

func (p *GetUserFishArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.GetUserFishReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetUserFishArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetUserFishArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetUserFishArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetUserFishArgs) Unmarshal(in []byte) error {
	msg := new(core_api.GetUserFishReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetUserFishArgs_Req_DEFAULT *core_api.GetUserFishReq

func (p *GetUserFishArgs) GetReq() *core_api.GetUserFishReq {
	if !p.IsSetReq() {
		return GetUserFishArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserFishArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetUserFishArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetUserFishResult struct {
	Success *core_api.GetUserFishResp
}

var GetUserFishResult_Success_DEFAULT *core_api.GetUserFishResp

func (p *GetUserFishResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.GetUserFishResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetUserFishResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetUserFishResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetUserFishResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetUserFishResult) Unmarshal(in []byte) error {
	msg := new(core_api.GetUserFishResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserFishResult) GetSuccess() *core_api.GetUserFishResp {
	if !p.IsSetSuccess() {
		return GetUserFishResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserFishResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.GetUserFishResp)
}

func (p *GetUserFishResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserFishResult) GetResult() interface{} {
	return p.Success
}

func listFishByPlanHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core_api.ListFishByPlanReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core_api.Plan).ListFishByPlan(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *ListFishByPlanArgs:
		success, err := handler.(core_api.Plan).ListFishByPlan(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ListFishByPlanResult)
		realResult.Success = success
	}
	return nil
}
func newListFishByPlanArgs() interface{} {
	return &ListFishByPlanArgs{}
}

func newListFishByPlanResult() interface{} {
	return &ListFishByPlanResult{}
}

type ListFishByPlanArgs struct {
	Req *core_api.ListFishByPlanReq
}

func (p *ListFishByPlanArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(core_api.ListFishByPlanReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ListFishByPlanArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ListFishByPlanArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ListFishByPlanArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ListFishByPlanArgs) Unmarshal(in []byte) error {
	msg := new(core_api.ListFishByPlanReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ListFishByPlanArgs_Req_DEFAULT *core_api.ListFishByPlanReq

func (p *ListFishByPlanArgs) GetReq() *core_api.ListFishByPlanReq {
	if !p.IsSetReq() {
		return ListFishByPlanArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ListFishByPlanArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ListFishByPlanArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ListFishByPlanResult struct {
	Success *core_api.ListFishByPlanResp
}

var ListFishByPlanResult_Success_DEFAULT *core_api.ListFishByPlanResp

func (p *ListFishByPlanResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(core_api.ListFishByPlanResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ListFishByPlanResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ListFishByPlanResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ListFishByPlanResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ListFishByPlanResult) Unmarshal(in []byte) error {
	msg := new(core_api.ListFishByPlanResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListFishByPlanResult) GetSuccess() *core_api.ListFishByPlanResp {
	if !p.IsSetSuccess() {
		return ListFishByPlanResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ListFishByPlanResult) SetSuccess(x interface{}) {
	p.Success = x.(*core_api.ListFishByPlanResp)
}

func (p *ListFishByPlanResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ListFishByPlanResult) GetResult() interface{} {
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

func (p *kClient) GetPlanPreviews(ctx context.Context, Req *core_api.GetPlanPreviewsReq) (r *core_api.GetPlanPreviewsResp, err error) {
	var _args GetPlanPreviewsArgs
	_args.Req = Req
	var _result GetPlanPreviewsResult
	if err = p.c.Call(ctx, "GetPlanPreviews", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetPlanDetail(ctx context.Context, Req *core_api.GetPlanDetailReq) (r *core_api.GetPlanDetailResp, err error) {
	var _args GetPlanDetailArgs
	_args.Req = Req
	var _result GetPlanDetailResult
	if err = p.c.Call(ctx, "GetPlanDetail", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) NewPlan(ctx context.Context, Req *core_api.NewPlanReq) (r *core_api.NewPlanResp, err error) {
	var _args NewPlanArgs
	_args.Req = Req
	var _result NewPlanResult
	if err = p.c.Call(ctx, "NewPlan", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeletePlan(ctx context.Context, Req *core_api.DeletePlanReq) (r *core_api.DeletePlanResp, err error) {
	var _args DeletePlanArgs
	_args.Req = Req
	var _result DeletePlanResult
	if err = p.c.Call(ctx, "DeletePlan", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SearchPlan(ctx context.Context, Req *core_api.SearchPlanReq) (r *core_api.SearchPlanResp, err error) {
	var _args SearchPlanArgs
	_args.Req = Req
	var _result SearchPlanResult
	if err = p.c.Call(ctx, "SearchPlan", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DonateFish(ctx context.Context, Req *core_api.DonateFishReq) (r *core_api.DonateFishResp, err error) {
	var _args DonateFishArgs
	_args.Req = Req
	var _result DonateFishResult
	if err = p.c.Call(ctx, "DonateFish", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserFish(ctx context.Context, Req *core_api.GetUserFishReq) (r *core_api.GetUserFishResp, err error) {
	var _args GetUserFishArgs
	_args.Req = Req
	var _result GetUserFishResult
	if err = p.c.Call(ctx, "GetUserFish", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListFishByPlan(ctx context.Context, Req *core_api.ListFishByPlanReq) (r *core_api.ListFishByPlanResp, err error) {
	var _args ListFishByPlanArgs
	_args.Req = Req
	var _result ListFishByPlanResult
	if err = p.c.Call(ctx, "ListFishByPlan", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
