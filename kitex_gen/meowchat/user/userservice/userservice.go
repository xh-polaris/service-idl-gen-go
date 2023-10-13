// Code generated by Kitex v0.7.2. DO NOT EDIT.

package userservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	user "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/user"
	proto "google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetUser":        kitex.NewMethodInfo(getUserHandler, newGetUserArgs, newGetUserResult, false),
		"GetUserDetail":  kitex.NewMethodInfo(getUserDetailHandler, newGetUserDetailArgs, newGetUserDetailResult, false),
		"UpdateUser":     kitex.NewMethodInfo(updateUserHandler, newUpdateUserArgs, newUpdateUserResult, false),
		"SearchUser":     kitex.NewMethodInfo(searchUserHandler, newSearchUserArgs, newSearchUserResult, false),
		"DoLike":         kitex.NewMethodInfo(doLikeHandler, newDoLikeArgs, newDoLikeResult, false),
		"GetUserLike":    kitex.NewMethodInfo(getUserLikeHandler, newGetUserLikeArgs, newGetUserLikeResult, false),
		"GetTargetLikes": kitex.NewMethodInfo(getTargetLikesHandler, newGetTargetLikesArgs, newGetTargetLikesResult, false),
		"GetUserLikes":   kitex.NewMethodInfo(getUserLikesHandler, newGetUserLikesArgs, newGetUserLikesResult, false),
		"GetLikedUsers":  kitex.NewMethodInfo(getLikedUsersHandler, newGetLikedUsersArgs, newGetLikedUsersResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "meowchat.user",
		"ServiceFilePath": ``,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.7.2",
		Extra:           extra,
	}
	return svcInfo
}

func getUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.GetUserReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).GetUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetUserArgs:
		success, err := handler.(user.UserService).GetUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetUserResult)
		realResult.Success = success
	}
	return nil
}
func newGetUserArgs() interface{} {
	return &GetUserArgs{}
}

func newGetUserResult() interface{} {
	return &GetUserResult{}
}

type GetUserArgs struct {
	Req *user.GetUserReq
}

func (p *GetUserArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.GetUserReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetUserArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetUserArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetUserArgs) Unmarshal(in []byte) error {
	msg := new(user.GetUserReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetUserArgs_Req_DEFAULT *user.GetUserReq

func (p *GetUserArgs) GetReq() *user.GetUserReq {
	if !p.IsSetReq() {
		return GetUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetUserArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetUserResult struct {
	Success *user.GetUserResp
}

var GetUserResult_Success_DEFAULT *user.GetUserResp

func (p *GetUserResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.GetUserResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetUserResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetUserResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetUserResult) Unmarshal(in []byte) error {
	msg := new(user.GetUserResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserResult) GetSuccess() *user.GetUserResp {
	if !p.IsSetSuccess() {
		return GetUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.GetUserResp)
}

func (p *GetUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserResult) GetResult() interface{} {
	return p.Success
}

func getUserDetailHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.GetUserDetailReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).GetUserDetail(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetUserDetailArgs:
		success, err := handler.(user.UserService).GetUserDetail(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetUserDetailResult)
		realResult.Success = success
	}
	return nil
}
func newGetUserDetailArgs() interface{} {
	return &GetUserDetailArgs{}
}

func newGetUserDetailResult() interface{} {
	return &GetUserDetailResult{}
}

type GetUserDetailArgs struct {
	Req *user.GetUserDetailReq
}

func (p *GetUserDetailArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.GetUserDetailReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetUserDetailArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetUserDetailArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetUserDetailArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetUserDetailArgs) Unmarshal(in []byte) error {
	msg := new(user.GetUserDetailReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetUserDetailArgs_Req_DEFAULT *user.GetUserDetailReq

func (p *GetUserDetailArgs) GetReq() *user.GetUserDetailReq {
	if !p.IsSetReq() {
		return GetUserDetailArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserDetailArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetUserDetailArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetUserDetailResult struct {
	Success *user.GetUserDetailResp
}

var GetUserDetailResult_Success_DEFAULT *user.GetUserDetailResp

func (p *GetUserDetailResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.GetUserDetailResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetUserDetailResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetUserDetailResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetUserDetailResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetUserDetailResult) Unmarshal(in []byte) error {
	msg := new(user.GetUserDetailResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserDetailResult) GetSuccess() *user.GetUserDetailResp {
	if !p.IsSetSuccess() {
		return GetUserDetailResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserDetailResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.GetUserDetailResp)
}

func (p *GetUserDetailResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserDetailResult) GetResult() interface{} {
	return p.Success
}

func updateUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.UpdateUserReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).UpdateUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UpdateUserArgs:
		success, err := handler.(user.UserService).UpdateUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateUserResult)
		realResult.Success = success
	}
	return nil
}
func newUpdateUserArgs() interface{} {
	return &UpdateUserArgs{}
}

func newUpdateUserResult() interface{} {
	return &UpdateUserResult{}
}

type UpdateUserArgs struct {
	Req *user.UpdateUserReq
}

func (p *UpdateUserArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.UpdateUserReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateUserArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateUserArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateUserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateUserArgs) Unmarshal(in []byte) error {
	msg := new(user.UpdateUserReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateUserArgs_Req_DEFAULT *user.UpdateUserReq

func (p *UpdateUserArgs) GetReq() *user.UpdateUserReq {
	if !p.IsSetReq() {
		return UpdateUserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateUserArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdateUserArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdateUserResult struct {
	Success *user.UpdateUserResp
}

var UpdateUserResult_Success_DEFAULT *user.UpdateUserResp

func (p *UpdateUserResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.UpdateUserResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateUserResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateUserResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateUserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateUserResult) Unmarshal(in []byte) error {
	msg := new(user.UpdateUserResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateUserResult) GetSuccess() *user.UpdateUserResp {
	if !p.IsSetSuccess() {
		return UpdateUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.UpdateUserResp)
}

func (p *UpdateUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateUserResult) GetResult() interface{} {
	return p.Success
}

func searchUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.SearchUserReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).SearchUser(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *SearchUserArgs:
		success, err := handler.(user.UserService).SearchUser(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SearchUserResult)
		realResult.Success = success
	}
	return nil
}
func newSearchUserArgs() interface{} {
	return &SearchUserArgs{}
}

func newSearchUserResult() interface{} {
	return &SearchUserResult{}
}

type SearchUserArgs struct {
	Req *user.SearchUserReq
}

func (p *SearchUserArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.SearchUserReq)
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
	msg := new(user.SearchUserReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SearchUserArgs_Req_DEFAULT *user.SearchUserReq

func (p *SearchUserArgs) GetReq() *user.SearchUserReq {
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
	Success *user.SearchUserResp
}

var SearchUserResult_Success_DEFAULT *user.SearchUserResp

func (p *SearchUserResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.SearchUserResp)
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
	msg := new(user.SearchUserResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SearchUserResult) GetSuccess() *user.SearchUserResp {
	if !p.IsSetSuccess() {
		return SearchUserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SearchUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.SearchUserResp)
}

func (p *SearchUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SearchUserResult) GetResult() interface{} {
	return p.Success
}

func doLikeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.DoLikeReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).DoLike(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *DoLikeArgs:
		success, err := handler.(user.UserService).DoLike(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DoLikeResult)
		realResult.Success = success
	}
	return nil
}
func newDoLikeArgs() interface{} {
	return &DoLikeArgs{}
}

func newDoLikeResult() interface{} {
	return &DoLikeResult{}
}

type DoLikeArgs struct {
	Req *user.DoLikeReq
}

func (p *DoLikeArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.DoLikeReq)
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
	msg := new(user.DoLikeReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DoLikeArgs_Req_DEFAULT *user.DoLikeReq

func (p *DoLikeArgs) GetReq() *user.DoLikeReq {
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
	Success *user.DoLikeResp
}

var DoLikeResult_Success_DEFAULT *user.DoLikeResp

func (p *DoLikeResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.DoLikeResp)
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
	msg := new(user.DoLikeResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DoLikeResult) GetSuccess() *user.DoLikeResp {
	if !p.IsSetSuccess() {
		return DoLikeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DoLikeResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.DoLikeResp)
}

func (p *DoLikeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DoLikeResult) GetResult() interface{} {
	return p.Success
}

func getUserLikeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.GetUserLikedReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).GetUserLike(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetUserLikeArgs:
		success, err := handler.(user.UserService).GetUserLike(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetUserLikeResult)
		realResult.Success = success
	}
	return nil
}
func newGetUserLikeArgs() interface{} {
	return &GetUserLikeArgs{}
}

func newGetUserLikeResult() interface{} {
	return &GetUserLikeResult{}
}

type GetUserLikeArgs struct {
	Req *user.GetUserLikedReq
}

func (p *GetUserLikeArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.GetUserLikedReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetUserLikeArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetUserLikeArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetUserLikeArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetUserLikeArgs) Unmarshal(in []byte) error {
	msg := new(user.GetUserLikedReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetUserLikeArgs_Req_DEFAULT *user.GetUserLikedReq

func (p *GetUserLikeArgs) GetReq() *user.GetUserLikedReq {
	if !p.IsSetReq() {
		return GetUserLikeArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserLikeArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetUserLikeArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetUserLikeResult struct {
	Success *user.GetUserLikedResp
}

var GetUserLikeResult_Success_DEFAULT *user.GetUserLikedResp

func (p *GetUserLikeResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.GetUserLikedResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetUserLikeResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetUserLikeResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetUserLikeResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetUserLikeResult) Unmarshal(in []byte) error {
	msg := new(user.GetUserLikedResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserLikeResult) GetSuccess() *user.GetUserLikedResp {
	if !p.IsSetSuccess() {
		return GetUserLikeResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserLikeResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.GetUserLikedResp)
}

func (p *GetUserLikeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserLikeResult) GetResult() interface{} {
	return p.Success
}

func getTargetLikesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.GetTargetLikesReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).GetTargetLikes(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetTargetLikesArgs:
		success, err := handler.(user.UserService).GetTargetLikes(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetTargetLikesResult)
		realResult.Success = success
	}
	return nil
}
func newGetTargetLikesArgs() interface{} {
	return &GetTargetLikesArgs{}
}

func newGetTargetLikesResult() interface{} {
	return &GetTargetLikesResult{}
}

type GetTargetLikesArgs struct {
	Req *user.GetTargetLikesReq
}

func (p *GetTargetLikesArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.GetTargetLikesReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetTargetLikesArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetTargetLikesArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetTargetLikesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetTargetLikesArgs) Unmarshal(in []byte) error {
	msg := new(user.GetTargetLikesReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetTargetLikesArgs_Req_DEFAULT *user.GetTargetLikesReq

func (p *GetTargetLikesArgs) GetReq() *user.GetTargetLikesReq {
	if !p.IsSetReq() {
		return GetTargetLikesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetTargetLikesArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetTargetLikesArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetTargetLikesResult struct {
	Success *user.GetTargetLikesResp
}

var GetTargetLikesResult_Success_DEFAULT *user.GetTargetLikesResp

func (p *GetTargetLikesResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.GetTargetLikesResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetTargetLikesResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetTargetLikesResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetTargetLikesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetTargetLikesResult) Unmarshal(in []byte) error {
	msg := new(user.GetTargetLikesResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetTargetLikesResult) GetSuccess() *user.GetTargetLikesResp {
	if !p.IsSetSuccess() {
		return GetTargetLikesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetTargetLikesResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.GetTargetLikesResp)
}

func (p *GetTargetLikesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetTargetLikesResult) GetResult() interface{} {
	return p.Success
}

func getUserLikesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.GetUserLikesReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).GetUserLikes(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetUserLikesArgs:
		success, err := handler.(user.UserService).GetUserLikes(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetUserLikesResult)
		realResult.Success = success
	}
	return nil
}
func newGetUserLikesArgs() interface{} {
	return &GetUserLikesArgs{}
}

func newGetUserLikesResult() interface{} {
	return &GetUserLikesResult{}
}

type GetUserLikesArgs struct {
	Req *user.GetUserLikesReq
}

func (p *GetUserLikesArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.GetUserLikesReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetUserLikesArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetUserLikesArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetUserLikesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetUserLikesArgs) Unmarshal(in []byte) error {
	msg := new(user.GetUserLikesReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetUserLikesArgs_Req_DEFAULT *user.GetUserLikesReq

func (p *GetUserLikesArgs) GetReq() *user.GetUserLikesReq {
	if !p.IsSetReq() {
		return GetUserLikesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserLikesArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetUserLikesArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetUserLikesResult struct {
	Success *user.GetUserLikesResp
}

var GetUserLikesResult_Success_DEFAULT *user.GetUserLikesResp

func (p *GetUserLikesResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.GetUserLikesResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetUserLikesResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetUserLikesResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetUserLikesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetUserLikesResult) Unmarshal(in []byte) error {
	msg := new(user.GetUserLikesResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserLikesResult) GetSuccess() *user.GetUserLikesResp {
	if !p.IsSetSuccess() {
		return GetUserLikesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserLikesResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.GetUserLikesResp)
}

func (p *GetUserLikesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetUserLikesResult) GetResult() interface{} {
	return p.Success
}

func getLikedUsersHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.GetLikedUsersReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserService).GetLikedUsers(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetLikedUsersArgs:
		success, err := handler.(user.UserService).GetLikedUsers(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetLikedUsersResult)
		realResult.Success = success
	}
	return nil
}
func newGetLikedUsersArgs() interface{} {
	return &GetLikedUsersArgs{}
}

func newGetLikedUsersResult() interface{} {
	return &GetLikedUsersResult{}
}

type GetLikedUsersArgs struct {
	Req *user.GetLikedUsersReq
}

func (p *GetLikedUsersArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.GetLikedUsersReq)
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
	msg := new(user.GetLikedUsersReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetLikedUsersArgs_Req_DEFAULT *user.GetLikedUsersReq

func (p *GetLikedUsersArgs) GetReq() *user.GetLikedUsersReq {
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
	Success *user.GetLikedUsersResp
}

var GetLikedUsersResult_Success_DEFAULT *user.GetLikedUsersResp

func (p *GetLikedUsersResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.GetLikedUsersResp)
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
	msg := new(user.GetLikedUsersResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetLikedUsersResult) GetSuccess() *user.GetLikedUsersResp {
	if !p.IsSetSuccess() {
		return GetLikedUsersResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetLikedUsersResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.GetLikedUsersResp)
}

func (p *GetLikedUsersResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetLikedUsersResult) GetResult() interface{} {
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

func (p *kClient) GetUser(ctx context.Context, Req *user.GetUserReq) (r *user.GetUserResp, err error) {
	var _args GetUserArgs
	_args.Req = Req
	var _result GetUserResult
	if err = p.c.Call(ctx, "GetUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserDetail(ctx context.Context, Req *user.GetUserDetailReq) (r *user.GetUserDetailResp, err error) {
	var _args GetUserDetailArgs
	_args.Req = Req
	var _result GetUserDetailResult
	if err = p.c.Call(ctx, "GetUserDetail", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateUser(ctx context.Context, Req *user.UpdateUserReq) (r *user.UpdateUserResp, err error) {
	var _args UpdateUserArgs
	_args.Req = Req
	var _result UpdateUserResult
	if err = p.c.Call(ctx, "UpdateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SearchUser(ctx context.Context, Req *user.SearchUserReq) (r *user.SearchUserResp, err error) {
	var _args SearchUserArgs
	_args.Req = Req
	var _result SearchUserResult
	if err = p.c.Call(ctx, "SearchUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DoLike(ctx context.Context, Req *user.DoLikeReq) (r *user.DoLikeResp, err error) {
	var _args DoLikeArgs
	_args.Req = Req
	var _result DoLikeResult
	if err = p.c.Call(ctx, "DoLike", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserLike(ctx context.Context, Req *user.GetUserLikedReq) (r *user.GetUserLikedResp, err error) {
	var _args GetUserLikeArgs
	_args.Req = Req
	var _result GetUserLikeResult
	if err = p.c.Call(ctx, "GetUserLike", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetTargetLikes(ctx context.Context, Req *user.GetTargetLikesReq) (r *user.GetTargetLikesResp, err error) {
	var _args GetTargetLikesArgs
	_args.Req = Req
	var _result GetTargetLikesResult
	if err = p.c.Call(ctx, "GetTargetLikes", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserLikes(ctx context.Context, Req *user.GetUserLikesReq) (r *user.GetUserLikesResp, err error) {
	var _args GetUserLikesArgs
	_args.Req = Req
	var _result GetUserLikesResult
	if err = p.c.Call(ctx, "GetUserLikes", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetLikedUsers(ctx context.Context, Req *user.GetLikedUsersReq) (r *user.GetLikedUsersResp, err error) {
	var _args GetLikedUsersArgs
	_args.Req = Req
	var _result GetLikedUsersResult
	if err = p.c.Call(ctx, "GetLikedUsers", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
