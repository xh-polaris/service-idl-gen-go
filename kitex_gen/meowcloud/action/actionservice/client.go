// Code generated by Kitex v0.12.0. DO NOT EDIT.

package actionservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	action "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowcloud/action"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	DoLike(ctx context.Context, Req *action.DoLikeReq, callOptions ...callopt.Option) (r *action.DoLikeResp, err error)
	CancelLike(ctx context.Context, Req *action.CancelLikeReq, callOptions ...callopt.Option) (r *action.CancelLikeResp, err error)
	GetLikedCount(ctx context.Context, Req *action.GetLikedCountReq, callOptions ...callopt.Option) (r *action.GetLikedCountResp, err error)
	GetLikedUsers(ctx context.Context, Req *action.GetLikedUsersReq, callOptions ...callopt.Option) (r *action.GetLikedUsersResp, err error)
	GetUserLiked(ctx context.Context, Req *action.GetUserLikedReq, callOptions ...callopt.Option) (r *action.GetUserLikedResp, err error)
	GetLiked(ctx context.Context, Req *action.GetLikedReq, callOptions ...callopt.Option) (r *action.GetLikedResp, err error)
	DoShare(ctx context.Context, Req *action.DoShareReq, callOptions ...callopt.Option) (r *action.DoShareResp, err error)
	GetSharedCount(ctx context.Context, Req *action.GetSharedCountReq, callOptions ...callopt.Option) (r *action.GetSharedCountResp, err error)
	GetSharedUsers(ctx context.Context, Req *action.GetSharedUsersReq, callOptions ...callopt.Option) (r *action.GetSharedUsersResp, err error)
	GetUserShared(ctx context.Context, Req *action.GetUserSharedReq, callOptions ...callopt.Option) (r *action.GetUserSharedResp, err error)
	GetShared(ctx context.Context, Req *action.GetSharedReq, callOptions ...callopt.Option) (r *action.GetSharedResp, err error)
	DoFollow(ctx context.Context, Req *action.DoFollowReq, callOptions ...callopt.Option) (r *action.DoFollowResp, err error)
	CancelFollow(ctx context.Context, Req *action.CancelFollowReq, callOptions ...callopt.Option) (r *action.CancelFollowResp, err error)
	GetFollowedCount(ctx context.Context, Req *action.GetFollowedCountReq, callOptions ...callopt.Option) (r *action.GetFollowedCountResp, err error)
	GetFollowedUsers(ctx context.Context, Req *action.GetFollowedUsersReq, callOptions ...callopt.Option) (r *action.GetFollowedUsersResp, err error)
	GetUserFollowed(ctx context.Context, Req *action.GetUserFollowedReq, callOptions ...callopt.Option) (r *action.GetUserFollowedResp, err error)
	GetFollowed(ctx context.Context, Req *action.GetFollowedReq, callOptions ...callopt.Option) (r *action.GetFollowedResp, err error)
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
	return &kActionServiceClient{
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

type kActionServiceClient struct {
	*kClient
}

func (p *kActionServiceClient) DoLike(ctx context.Context, Req *action.DoLikeReq, callOptions ...callopt.Option) (r *action.DoLikeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DoLike(ctx, Req)
}

func (p *kActionServiceClient) CancelLike(ctx context.Context, Req *action.CancelLikeReq, callOptions ...callopt.Option) (r *action.CancelLikeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CancelLike(ctx, Req)
}

func (p *kActionServiceClient) GetLikedCount(ctx context.Context, Req *action.GetLikedCountReq, callOptions ...callopt.Option) (r *action.GetLikedCountResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetLikedCount(ctx, Req)
}

func (p *kActionServiceClient) GetLikedUsers(ctx context.Context, Req *action.GetLikedUsersReq, callOptions ...callopt.Option) (r *action.GetLikedUsersResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetLikedUsers(ctx, Req)
}

func (p *kActionServiceClient) GetUserLiked(ctx context.Context, Req *action.GetUserLikedReq, callOptions ...callopt.Option) (r *action.GetUserLikedResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserLiked(ctx, Req)
}

func (p *kActionServiceClient) GetLiked(ctx context.Context, Req *action.GetLikedReq, callOptions ...callopt.Option) (r *action.GetLikedResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetLiked(ctx, Req)
}

func (p *kActionServiceClient) DoShare(ctx context.Context, Req *action.DoShareReq, callOptions ...callopt.Option) (r *action.DoShareResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DoShare(ctx, Req)
}

func (p *kActionServiceClient) GetSharedCount(ctx context.Context, Req *action.GetSharedCountReq, callOptions ...callopt.Option) (r *action.GetSharedCountResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetSharedCount(ctx, Req)
}

func (p *kActionServiceClient) GetSharedUsers(ctx context.Context, Req *action.GetSharedUsersReq, callOptions ...callopt.Option) (r *action.GetSharedUsersResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetSharedUsers(ctx, Req)
}

func (p *kActionServiceClient) GetUserShared(ctx context.Context, Req *action.GetUserSharedReq, callOptions ...callopt.Option) (r *action.GetUserSharedResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserShared(ctx, Req)
}

func (p *kActionServiceClient) GetShared(ctx context.Context, Req *action.GetSharedReq, callOptions ...callopt.Option) (r *action.GetSharedResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetShared(ctx, Req)
}

func (p *kActionServiceClient) DoFollow(ctx context.Context, Req *action.DoFollowReq, callOptions ...callopt.Option) (r *action.DoFollowResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DoFollow(ctx, Req)
}

func (p *kActionServiceClient) CancelFollow(ctx context.Context, Req *action.CancelFollowReq, callOptions ...callopt.Option) (r *action.CancelFollowResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CancelFollow(ctx, Req)
}

func (p *kActionServiceClient) GetFollowedCount(ctx context.Context, Req *action.GetFollowedCountReq, callOptions ...callopt.Option) (r *action.GetFollowedCountResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollowedCount(ctx, Req)
}

func (p *kActionServiceClient) GetFollowedUsers(ctx context.Context, Req *action.GetFollowedUsersReq, callOptions ...callopt.Option) (r *action.GetFollowedUsersResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollowedUsers(ctx, Req)
}

func (p *kActionServiceClient) GetUserFollowed(ctx context.Context, Req *action.GetUserFollowedReq, callOptions ...callopt.Option) (r *action.GetUserFollowedResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserFollowed(ctx, Req)
}

func (p *kActionServiceClient) GetFollowed(ctx context.Context, Req *action.GetFollowedReq, callOptions ...callopt.Option) (r *action.GetFollowedResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFollowed(ctx, Req)
}
