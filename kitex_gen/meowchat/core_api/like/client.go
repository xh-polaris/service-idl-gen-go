// Code generated by Kitex v0.7.3. DO NOT EDIT.

package like

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	core_api "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/core_api"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	DoLike(ctx context.Context, Req *core_api.DoLikeReq, callOptions ...callopt.Option) (r *core_api.DoLikeResp, err error)
	GetUserLiked(ctx context.Context, Req *core_api.GetUserLikedReq, callOptions ...callopt.Option) (r *core_api.GetUserLikedResp, err error)
	GetLikedCount(ctx context.Context, Req *core_api.GetLikedCountReq, callOptions ...callopt.Option) (r *core_api.GetLikedCountResp, err error)
	GetLikedUsers(ctx context.Context, Req *core_api.GetLikedUsersReq, callOptions ...callopt.Option) (r *core_api.GetLikedUsersResp, err error)
	GetUserLikes(ctx context.Context, Req *core_api.GetUserLikesReq, callOptions ...callopt.Option) (r *core_api.GetUserLikesResp, err error)
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
	return &kLikeClient{
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

type kLikeClient struct {
	*kClient
}

func (p *kLikeClient) DoLike(ctx context.Context, Req *core_api.DoLikeReq, callOptions ...callopt.Option) (r *core_api.DoLikeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DoLike(ctx, Req)
}

func (p *kLikeClient) GetUserLiked(ctx context.Context, Req *core_api.GetUserLikedReq, callOptions ...callopt.Option) (r *core_api.GetUserLikedResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserLiked(ctx, Req)
}

func (p *kLikeClient) GetLikedCount(ctx context.Context, Req *core_api.GetLikedCountReq, callOptions ...callopt.Option) (r *core_api.GetLikedCountResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetLikedCount(ctx, Req)
}

func (p *kLikeClient) GetLikedUsers(ctx context.Context, Req *core_api.GetLikedUsersReq, callOptions ...callopt.Option) (r *core_api.GetLikedUsersResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetLikedUsers(ctx, Req)
}

func (p *kLikeClient) GetUserLikes(ctx context.Context, Req *core_api.GetUserLikesReq, callOptions ...callopt.Option) (r *core_api.GetUserLikesResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserLikes(ctx, Req)
}
