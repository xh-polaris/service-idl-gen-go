// Code generated by Kitex v0.11.3. DO NOT EDIT.

package user

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	core_api "github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/core_api"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	SignUp(ctx context.Context, Req *core_api.SignUpReq, callOptions ...callopt.Option) (r *core_api.SignUpResp, err error)
	SignIn(ctx context.Context, Req *core_api.SignUpReq, callOptions ...callopt.Option) (r *core_api.SignUpResp, err error)
	GetUserInfo(ctx context.Context, Req *core_api.GetUserInfoReq, callOptions ...callopt.Option) (r *core_api.GetUserInfoResp, err error)
	SetUserInfo(ctx context.Context, Req *core_api.SetUserInfoReq, callOptions ...callopt.Option) (r *core_api.Response, err error)
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

func (p *kUserClient) SignUp(ctx context.Context, Req *core_api.SignUpReq, callOptions ...callopt.Option) (r *core_api.SignUpResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SignUp(ctx, Req)
}

func (p *kUserClient) SignIn(ctx context.Context, Req *core_api.SignUpReq, callOptions ...callopt.Option) (r *core_api.SignUpResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SignIn(ctx, Req)
}

func (p *kUserClient) GetUserInfo(ctx context.Context, Req *core_api.GetUserInfoReq, callOptions ...callopt.Option) (r *core_api.GetUserInfoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserInfo(ctx, Req)
}

func (p *kUserClient) SetUserInfo(ctx context.Context, Req *core_api.SetUserInfoReq, callOptions ...callopt.Option) (r *core_api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SetUserInfo(ctx, Req)
}
