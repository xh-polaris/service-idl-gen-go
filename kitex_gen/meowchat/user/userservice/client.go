// Code generated by Kitex v0.8.0. DO NOT EDIT.

package userservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	user "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/user"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetUser(ctx context.Context, Req *user.GetUserReq, callOptions ...callopt.Option) (r *user.GetUserResp, err error)
	GetUserDetail(ctx context.Context, Req *user.GetUserDetailReq, callOptions ...callopt.Option) (r *user.GetUserDetailResp, err error)
	UpdateUser(ctx context.Context, Req *user.UpdateUserReq, callOptions ...callopt.Option) (r *user.UpdateUserResp, err error)
	SearchUser(ctx context.Context, Req *user.SearchUserReq, callOptions ...callopt.Option) (r *user.SearchUserResp, err error)
	CheckIn(ctx context.Context, Req *user.CheckInReq, callOptions ...callopt.Option) (r *user.CheckInResp, err error)
	DoLike(ctx context.Context, Req *user.DoLikeReq, callOptions ...callopt.Option) (r *user.DoLikeResp, err error)
	GetUserLike(ctx context.Context, Req *user.GetUserLikedReq, callOptions ...callopt.Option) (r *user.GetUserLikedResp, err error)
	GetTargetLikes(ctx context.Context, Req *user.GetTargetLikesReq, callOptions ...callopt.Option) (r *user.GetTargetLikesResp, err error)
	GetUserLikes(ctx context.Context, Req *user.GetUserLikesReq, callOptions ...callopt.Option) (r *user.GetUserLikesResp, err error)
	GetLikedUsers(ctx context.Context, Req *user.GetLikedUsersReq, callOptions ...callopt.Option) (r *user.GetLikedUsersResp, err error)
	GetUserMission(ctx context.Context, Req *user.GetUserMissionReq, callOptions ...callopt.Option) (r *user.GetUserMissionResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) GetUser(ctx context.Context, Req *user.GetUserReq, callOptions ...callopt.Option) (r *user.GetUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUser(ctx, Req)
}

func (p *kUserServiceClient) GetUserDetail(ctx context.Context, Req *user.GetUserDetailReq, callOptions ...callopt.Option) (r *user.GetUserDetailResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserDetail(ctx, Req)
}

func (p *kUserServiceClient) UpdateUser(ctx context.Context, Req *user.UpdateUserReq, callOptions ...callopt.Option) (r *user.UpdateUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateUser(ctx, Req)
}

func (p *kUserServiceClient) SearchUser(ctx context.Context, Req *user.SearchUserReq, callOptions ...callopt.Option) (r *user.SearchUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SearchUser(ctx, Req)
}

func (p *kUserServiceClient) CheckIn(ctx context.Context, Req *user.CheckInReq, callOptions ...callopt.Option) (r *user.CheckInResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CheckIn(ctx, Req)
}

func (p *kUserServiceClient) DoLike(ctx context.Context, Req *user.DoLikeReq, callOptions ...callopt.Option) (r *user.DoLikeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DoLike(ctx, Req)
}

func (p *kUserServiceClient) GetUserLike(ctx context.Context, Req *user.GetUserLikedReq, callOptions ...callopt.Option) (r *user.GetUserLikedResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserLike(ctx, Req)
}

func (p *kUserServiceClient) GetTargetLikes(ctx context.Context, Req *user.GetTargetLikesReq, callOptions ...callopt.Option) (r *user.GetTargetLikesResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetTargetLikes(ctx, Req)
}

func (p *kUserServiceClient) GetUserLikes(ctx context.Context, Req *user.GetUserLikesReq, callOptions ...callopt.Option) (r *user.GetUserLikesResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserLikes(ctx, Req)
}

func (p *kUserServiceClient) GetLikedUsers(ctx context.Context, Req *user.GetLikedUsersReq, callOptions ...callopt.Option) (r *user.GetLikedUsersResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetLikedUsers(ctx, Req)
}

func (p *kUserServiceClient) GetUserMission(ctx context.Context, Req *user.GetUserMissionReq, callOptions ...callopt.Option) (r *user.GetUserMissionResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserMission(ctx, Req)
}
