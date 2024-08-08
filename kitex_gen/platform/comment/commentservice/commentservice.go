// Code generated by Kitex v0.10.3. DO NOT EDIT.

package commentservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	comment "github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/comment"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"createComment": kitex.NewMethodInfo(
		createCommentHandler,
		newCreateCommentArgs,
		newCreateCommentResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"updateComment": kitex.NewMethodInfo(
		updateCommentHandler,
		newUpdateCommentArgs,
		newUpdateCommentResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"deleteComment": kitex.NewMethodInfo(
		deleteCommentHandler,
		newDeleteCommentArgs,
		newDeleteCommentResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"listCommentByParent": kitex.NewMethodInfo(
		listCommentByParentHandler,
		newListCommentByParentArgs,
		newListCommentByParentResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"countCommentByParent": kitex.NewMethodInfo(
		countCommentByParentHandler,
		newCountCommentByParentArgs,
		newCountCommentByParentResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"retrieveCommentById": kitex.NewMethodInfo(
		retrieveCommentByIdHandler,
		newRetrieveCommentByIdArgs,
		newRetrieveCommentByIdResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"listCommentByAuthorIdAndType": kitex.NewMethodInfo(
		listCommentByAuthorIdAndTypeHandler,
		newListCommentByAuthorIdAndTypeArgs,
		newListCommentByAuthorIdAndTypeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"listCommentByReplyToAndType": kitex.NewMethodInfo(
		listCommentByReplyToAndTypeHandler,
		newListCommentByReplyToAndTypeArgs,
		newListCommentByReplyToAndTypeResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	commentServiceServiceInfo                = NewServiceInfo()
	commentServiceServiceInfoForClient       = NewServiceInfoForClient()
	commentServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return commentServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return commentServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return commentServiceServiceInfoForClient
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
	serviceName := "CommentService"
	handlerType := (*comment.CommentService)(nil)
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
		"PackageName": "platform.comment",
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

func createCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(comment.CreateCommentReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(comment.CommentService).CreateComment(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *CreateCommentArgs:
		success, err := handler.(comment.CommentService).CreateComment(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateCommentResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newCreateCommentArgs() interface{} {
	return &CreateCommentArgs{}
}

func newCreateCommentResult() interface{} {
	return &CreateCommentResult{}
}

type CreateCommentArgs struct {
	Req *comment.CreateCommentReq
}

func (p *CreateCommentArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(comment.CreateCommentReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CreateCommentArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CreateCommentArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CreateCommentArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CreateCommentArgs) Unmarshal(in []byte) error {
	msg := new(comment.CreateCommentReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateCommentArgs_Req_DEFAULT *comment.CreateCommentReq

func (p *CreateCommentArgs) GetReq() *comment.CreateCommentReq {
	if !p.IsSetReq() {
		return CreateCommentArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateCommentArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CreateCommentArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CreateCommentResult struct {
	Success *comment.CreateCommentResp
}

var CreateCommentResult_Success_DEFAULT *comment.CreateCommentResp

func (p *CreateCommentResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(comment.CreateCommentResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CreateCommentResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CreateCommentResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CreateCommentResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CreateCommentResult) Unmarshal(in []byte) error {
	msg := new(comment.CreateCommentResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateCommentResult) GetSuccess() *comment.CreateCommentResp {
	if !p.IsSetSuccess() {
		return CreateCommentResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateCommentResult) SetSuccess(x interface{}) {
	p.Success = x.(*comment.CreateCommentResp)
}

func (p *CreateCommentResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CreateCommentResult) GetResult() interface{} {
	return p.Success
}

func updateCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(comment.UpdateCommentReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(comment.CommentService).UpdateComment(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *UpdateCommentArgs:
		success, err := handler.(comment.CommentService).UpdateComment(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateCommentResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newUpdateCommentArgs() interface{} {
	return &UpdateCommentArgs{}
}

func newUpdateCommentResult() interface{} {
	return &UpdateCommentResult{}
}

type UpdateCommentArgs struct {
	Req *comment.UpdateCommentReq
}

func (p *UpdateCommentArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(comment.UpdateCommentReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateCommentArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateCommentArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateCommentArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateCommentArgs) Unmarshal(in []byte) error {
	msg := new(comment.UpdateCommentReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateCommentArgs_Req_DEFAULT *comment.UpdateCommentReq

func (p *UpdateCommentArgs) GetReq() *comment.UpdateCommentReq {
	if !p.IsSetReq() {
		return UpdateCommentArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateCommentArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdateCommentArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdateCommentResult struct {
	Success *comment.UpdateCommentResp
}

var UpdateCommentResult_Success_DEFAULT *comment.UpdateCommentResp

func (p *UpdateCommentResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(comment.UpdateCommentResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateCommentResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateCommentResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateCommentResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateCommentResult) Unmarshal(in []byte) error {
	msg := new(comment.UpdateCommentResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateCommentResult) GetSuccess() *comment.UpdateCommentResp {
	if !p.IsSetSuccess() {
		return UpdateCommentResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateCommentResult) SetSuccess(x interface{}) {
	p.Success = x.(*comment.UpdateCommentResp)
}

func (p *UpdateCommentResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateCommentResult) GetResult() interface{} {
	return p.Success
}

func deleteCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(comment.DeleteCommentByIdReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(comment.CommentService).DeleteComment(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *DeleteCommentArgs:
		success, err := handler.(comment.CommentService).DeleteComment(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteCommentResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newDeleteCommentArgs() interface{} {
	return &DeleteCommentArgs{}
}

func newDeleteCommentResult() interface{} {
	return &DeleteCommentResult{}
}

type DeleteCommentArgs struct {
	Req *comment.DeleteCommentByIdReq
}

func (p *DeleteCommentArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(comment.DeleteCommentByIdReq)
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
	msg := new(comment.DeleteCommentByIdReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteCommentArgs_Req_DEFAULT *comment.DeleteCommentByIdReq

func (p *DeleteCommentArgs) GetReq() *comment.DeleteCommentByIdReq {
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
	Success *comment.DeleteCommentByIdResp
}

var DeleteCommentResult_Success_DEFAULT *comment.DeleteCommentByIdResp

func (p *DeleteCommentResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(comment.DeleteCommentByIdResp)
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
	msg := new(comment.DeleteCommentByIdResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteCommentResult) GetSuccess() *comment.DeleteCommentByIdResp {
	if !p.IsSetSuccess() {
		return DeleteCommentResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteCommentResult) SetSuccess(x interface{}) {
	p.Success = x.(*comment.DeleteCommentByIdResp)
}

func (p *DeleteCommentResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteCommentResult) GetResult() interface{} {
	return p.Success
}

func listCommentByParentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(comment.ListCommentByParentReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(comment.CommentService).ListCommentByParent(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ListCommentByParentArgs:
		success, err := handler.(comment.CommentService).ListCommentByParent(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ListCommentByParentResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newListCommentByParentArgs() interface{} {
	return &ListCommentByParentArgs{}
}

func newListCommentByParentResult() interface{} {
	return &ListCommentByParentResult{}
}

type ListCommentByParentArgs struct {
	Req *comment.ListCommentByParentReq
}

func (p *ListCommentByParentArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(comment.ListCommentByParentReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ListCommentByParentArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ListCommentByParentArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ListCommentByParentArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ListCommentByParentArgs) Unmarshal(in []byte) error {
	msg := new(comment.ListCommentByParentReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ListCommentByParentArgs_Req_DEFAULT *comment.ListCommentByParentReq

func (p *ListCommentByParentArgs) GetReq() *comment.ListCommentByParentReq {
	if !p.IsSetReq() {
		return ListCommentByParentArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ListCommentByParentArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ListCommentByParentArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ListCommentByParentResult struct {
	Success *comment.ListCommentByParentResp
}

var ListCommentByParentResult_Success_DEFAULT *comment.ListCommentByParentResp

func (p *ListCommentByParentResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(comment.ListCommentByParentResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ListCommentByParentResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ListCommentByParentResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ListCommentByParentResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ListCommentByParentResult) Unmarshal(in []byte) error {
	msg := new(comment.ListCommentByParentResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListCommentByParentResult) GetSuccess() *comment.ListCommentByParentResp {
	if !p.IsSetSuccess() {
		return ListCommentByParentResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ListCommentByParentResult) SetSuccess(x interface{}) {
	p.Success = x.(*comment.ListCommentByParentResp)
}

func (p *ListCommentByParentResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ListCommentByParentResult) GetResult() interface{} {
	return p.Success
}

func countCommentByParentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(comment.CountCommentByParentReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(comment.CommentService).CountCommentByParent(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *CountCommentByParentArgs:
		success, err := handler.(comment.CommentService).CountCommentByParent(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CountCommentByParentResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newCountCommentByParentArgs() interface{} {
	return &CountCommentByParentArgs{}
}

func newCountCommentByParentResult() interface{} {
	return &CountCommentByParentResult{}
}

type CountCommentByParentArgs struct {
	Req *comment.CountCommentByParentReq
}

func (p *CountCommentByParentArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(comment.CountCommentByParentReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CountCommentByParentArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CountCommentByParentArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CountCommentByParentArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CountCommentByParentArgs) Unmarshal(in []byte) error {
	msg := new(comment.CountCommentByParentReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CountCommentByParentArgs_Req_DEFAULT *comment.CountCommentByParentReq

func (p *CountCommentByParentArgs) GetReq() *comment.CountCommentByParentReq {
	if !p.IsSetReq() {
		return CountCommentByParentArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CountCommentByParentArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CountCommentByParentArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CountCommentByParentResult struct {
	Success *comment.CountCommentByParentResp
}

var CountCommentByParentResult_Success_DEFAULT *comment.CountCommentByParentResp

func (p *CountCommentByParentResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(comment.CountCommentByParentResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CountCommentByParentResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CountCommentByParentResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CountCommentByParentResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CountCommentByParentResult) Unmarshal(in []byte) error {
	msg := new(comment.CountCommentByParentResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CountCommentByParentResult) GetSuccess() *comment.CountCommentByParentResp {
	if !p.IsSetSuccess() {
		return CountCommentByParentResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CountCommentByParentResult) SetSuccess(x interface{}) {
	p.Success = x.(*comment.CountCommentByParentResp)
}

func (p *CountCommentByParentResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CountCommentByParentResult) GetResult() interface{} {
	return p.Success
}

func retrieveCommentByIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(comment.RetrieveCommentByIdReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(comment.CommentService).RetrieveCommentById(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *RetrieveCommentByIdArgs:
		success, err := handler.(comment.CommentService).RetrieveCommentById(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RetrieveCommentByIdResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newRetrieveCommentByIdArgs() interface{} {
	return &RetrieveCommentByIdArgs{}
}

func newRetrieveCommentByIdResult() interface{} {
	return &RetrieveCommentByIdResult{}
}

type RetrieveCommentByIdArgs struct {
	Req *comment.RetrieveCommentByIdReq
}

func (p *RetrieveCommentByIdArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(comment.RetrieveCommentByIdReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *RetrieveCommentByIdArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *RetrieveCommentByIdArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *RetrieveCommentByIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *RetrieveCommentByIdArgs) Unmarshal(in []byte) error {
	msg := new(comment.RetrieveCommentByIdReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RetrieveCommentByIdArgs_Req_DEFAULT *comment.RetrieveCommentByIdReq

func (p *RetrieveCommentByIdArgs) GetReq() *comment.RetrieveCommentByIdReq {
	if !p.IsSetReq() {
		return RetrieveCommentByIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RetrieveCommentByIdArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *RetrieveCommentByIdArgs) GetFirstArgument() interface{} {
	return p.Req
}

type RetrieveCommentByIdResult struct {
	Success *comment.RetrieveCommentByIdResp
}

var RetrieveCommentByIdResult_Success_DEFAULT *comment.RetrieveCommentByIdResp

func (p *RetrieveCommentByIdResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(comment.RetrieveCommentByIdResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *RetrieveCommentByIdResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *RetrieveCommentByIdResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *RetrieveCommentByIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *RetrieveCommentByIdResult) Unmarshal(in []byte) error {
	msg := new(comment.RetrieveCommentByIdResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RetrieveCommentByIdResult) GetSuccess() *comment.RetrieveCommentByIdResp {
	if !p.IsSetSuccess() {
		return RetrieveCommentByIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RetrieveCommentByIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*comment.RetrieveCommentByIdResp)
}

func (p *RetrieveCommentByIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RetrieveCommentByIdResult) GetResult() interface{} {
	return p.Success
}

func listCommentByAuthorIdAndTypeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(comment.ListCommentByAuthorIdAndTypeReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(comment.CommentService).ListCommentByAuthorIdAndType(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ListCommentByAuthorIdAndTypeArgs:
		success, err := handler.(comment.CommentService).ListCommentByAuthorIdAndType(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ListCommentByAuthorIdAndTypeResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newListCommentByAuthorIdAndTypeArgs() interface{} {
	return &ListCommentByAuthorIdAndTypeArgs{}
}

func newListCommentByAuthorIdAndTypeResult() interface{} {
	return &ListCommentByAuthorIdAndTypeResult{}
}

type ListCommentByAuthorIdAndTypeArgs struct {
	Req *comment.ListCommentByAuthorIdAndTypeReq
}

func (p *ListCommentByAuthorIdAndTypeArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(comment.ListCommentByAuthorIdAndTypeReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ListCommentByAuthorIdAndTypeArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ListCommentByAuthorIdAndTypeArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ListCommentByAuthorIdAndTypeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ListCommentByAuthorIdAndTypeArgs) Unmarshal(in []byte) error {
	msg := new(comment.ListCommentByAuthorIdAndTypeReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ListCommentByAuthorIdAndTypeArgs_Req_DEFAULT *comment.ListCommentByAuthorIdAndTypeReq

func (p *ListCommentByAuthorIdAndTypeArgs) GetReq() *comment.ListCommentByAuthorIdAndTypeReq {
	if !p.IsSetReq() {
		return ListCommentByAuthorIdAndTypeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ListCommentByAuthorIdAndTypeArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ListCommentByAuthorIdAndTypeArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ListCommentByAuthorIdAndTypeResult struct {
	Success *comment.ListCommentByAuthorIdAndTypeResp
}

var ListCommentByAuthorIdAndTypeResult_Success_DEFAULT *comment.ListCommentByAuthorIdAndTypeResp

func (p *ListCommentByAuthorIdAndTypeResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(comment.ListCommentByAuthorIdAndTypeResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ListCommentByAuthorIdAndTypeResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ListCommentByAuthorIdAndTypeResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ListCommentByAuthorIdAndTypeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ListCommentByAuthorIdAndTypeResult) Unmarshal(in []byte) error {
	msg := new(comment.ListCommentByAuthorIdAndTypeResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListCommentByAuthorIdAndTypeResult) GetSuccess() *comment.ListCommentByAuthorIdAndTypeResp {
	if !p.IsSetSuccess() {
		return ListCommentByAuthorIdAndTypeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ListCommentByAuthorIdAndTypeResult) SetSuccess(x interface{}) {
	p.Success = x.(*comment.ListCommentByAuthorIdAndTypeResp)
}

func (p *ListCommentByAuthorIdAndTypeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ListCommentByAuthorIdAndTypeResult) GetResult() interface{} {
	return p.Success
}

func listCommentByReplyToAndTypeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(comment.ListCommentByReplyToAndTypeReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(comment.CommentService).ListCommentByReplyToAndType(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ListCommentByReplyToAndTypeArgs:
		success, err := handler.(comment.CommentService).ListCommentByReplyToAndType(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ListCommentByReplyToAndTypeResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newListCommentByReplyToAndTypeArgs() interface{} {
	return &ListCommentByReplyToAndTypeArgs{}
}

func newListCommentByReplyToAndTypeResult() interface{} {
	return &ListCommentByReplyToAndTypeResult{}
}

type ListCommentByReplyToAndTypeArgs struct {
	Req *comment.ListCommentByReplyToAndTypeReq
}

func (p *ListCommentByReplyToAndTypeArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(comment.ListCommentByReplyToAndTypeReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ListCommentByReplyToAndTypeArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ListCommentByReplyToAndTypeArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ListCommentByReplyToAndTypeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ListCommentByReplyToAndTypeArgs) Unmarshal(in []byte) error {
	msg := new(comment.ListCommentByReplyToAndTypeReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ListCommentByReplyToAndTypeArgs_Req_DEFAULT *comment.ListCommentByReplyToAndTypeReq

func (p *ListCommentByReplyToAndTypeArgs) GetReq() *comment.ListCommentByReplyToAndTypeReq {
	if !p.IsSetReq() {
		return ListCommentByReplyToAndTypeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ListCommentByReplyToAndTypeArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ListCommentByReplyToAndTypeArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ListCommentByReplyToAndTypeResult struct {
	Success *comment.ListCommentByReplyToAndTypeResp
}

var ListCommentByReplyToAndTypeResult_Success_DEFAULT *comment.ListCommentByReplyToAndTypeResp

func (p *ListCommentByReplyToAndTypeResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(comment.ListCommentByReplyToAndTypeResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ListCommentByReplyToAndTypeResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ListCommentByReplyToAndTypeResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ListCommentByReplyToAndTypeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ListCommentByReplyToAndTypeResult) Unmarshal(in []byte) error {
	msg := new(comment.ListCommentByReplyToAndTypeResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListCommentByReplyToAndTypeResult) GetSuccess() *comment.ListCommentByReplyToAndTypeResp {
	if !p.IsSetSuccess() {
		return ListCommentByReplyToAndTypeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ListCommentByReplyToAndTypeResult) SetSuccess(x interface{}) {
	p.Success = x.(*comment.ListCommentByReplyToAndTypeResp)
}

func (p *ListCommentByReplyToAndTypeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ListCommentByReplyToAndTypeResult) GetResult() interface{} {
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

func (p *kClient) CreateComment(ctx context.Context, Req *comment.CreateCommentReq) (r *comment.CreateCommentResp, err error) {
	var _args CreateCommentArgs
	_args.Req = Req
	var _result CreateCommentResult
	if err = p.c.Call(ctx, "createComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateComment(ctx context.Context, Req *comment.UpdateCommentReq) (r *comment.UpdateCommentResp, err error) {
	var _args UpdateCommentArgs
	_args.Req = Req
	var _result UpdateCommentResult
	if err = p.c.Call(ctx, "updateComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteComment(ctx context.Context, Req *comment.DeleteCommentByIdReq) (r *comment.DeleteCommentByIdResp, err error) {
	var _args DeleteCommentArgs
	_args.Req = Req
	var _result DeleteCommentResult
	if err = p.c.Call(ctx, "deleteComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListCommentByParent(ctx context.Context, Req *comment.ListCommentByParentReq) (r *comment.ListCommentByParentResp, err error) {
	var _args ListCommentByParentArgs
	_args.Req = Req
	var _result ListCommentByParentResult
	if err = p.c.Call(ctx, "listCommentByParent", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CountCommentByParent(ctx context.Context, Req *comment.CountCommentByParentReq) (r *comment.CountCommentByParentResp, err error) {
	var _args CountCommentByParentArgs
	_args.Req = Req
	var _result CountCommentByParentResult
	if err = p.c.Call(ctx, "countCommentByParent", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) RetrieveCommentById(ctx context.Context, Req *comment.RetrieveCommentByIdReq) (r *comment.RetrieveCommentByIdResp, err error) {
	var _args RetrieveCommentByIdArgs
	_args.Req = Req
	var _result RetrieveCommentByIdResult
	if err = p.c.Call(ctx, "retrieveCommentById", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListCommentByAuthorIdAndType(ctx context.Context, Req *comment.ListCommentByAuthorIdAndTypeReq) (r *comment.ListCommentByAuthorIdAndTypeResp, err error) {
	var _args ListCommentByAuthorIdAndTypeArgs
	_args.Req = Req
	var _result ListCommentByAuthorIdAndTypeResult
	if err = p.c.Call(ctx, "listCommentByAuthorIdAndType", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListCommentByReplyToAndType(ctx context.Context, Req *comment.ListCommentByReplyToAndTypeReq) (r *comment.ListCommentByReplyToAndTypeResp, err error) {
	var _args ListCommentByReplyToAndTypeArgs
	_args.Req = Req
	var _result ListCommentByReplyToAndTypeResult
	if err = p.c.Call(ctx, "listCommentByReplyToAndType", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
