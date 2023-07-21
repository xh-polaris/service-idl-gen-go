// Code generated by Kitex v0.6.1. DO NOT EDIT.

package user

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	core_api "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/core_api"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetUserInfo(ctx context.Context, Req *core_api.GetUserInfoReq, callOptions ...callopt.Option) (r *core_api.GetUserInfoResp, err error)
	UpdateUserInfo(ctx context.Context, Req *core_api.UpdateUserInfoReq, callOptions ...callopt.Option) (r *core_api.UpdateUserInfoResp, err error)
	SearchUser(ctx context.Context, Req *core_api.SearchUserReq, callOptions ...callopt.Option) (r *core_api.SearchUserResp, err error)
	SearchUserForAdmin(ctx context.Context, Req *core_api.SearchUserForAdminReq, callOptions ...callopt.Option) (r *core_api.SearchUserForAdminResp, err error)
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
	return &kUserClient{
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

type kUserClient struct {
	*kClient
}

func (p *kUserClient) GetUserInfo(ctx context.Context, Req *core_api.GetUserInfoReq, callOptions ...callopt.Option) (r *core_api.GetUserInfoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserInfo(ctx, Req)
}

func (p *kUserClient) UpdateUserInfo(ctx context.Context, Req *core_api.UpdateUserInfoReq, callOptions ...callopt.Option) (r *core_api.UpdateUserInfoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateUserInfo(ctx, Req)
}

func (p *kUserClient) SearchUser(ctx context.Context, Req *core_api.SearchUserReq, callOptions ...callopt.Option) (r *core_api.SearchUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SearchUser(ctx, Req)
}

func (p *kUserClient) SearchUserForAdmin(ctx context.Context, Req *core_api.SearchUserForAdminReq, callOptions ...callopt.Option) (r *core_api.SearchUserForAdminResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SearchUserForAdmin(ctx, Req)
}
