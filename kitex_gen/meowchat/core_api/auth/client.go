// Code generated by Kitex v0.6.1. DO NOT EDIT.

package auth

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	core_api "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/core_api"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	SignIn(ctx context.Context, Req *core_api.SignInReq, callOptions ...callopt.Option) (r *core_api.SignInResp, err error)
	SendVerifyCode(ctx context.Context, Req *core_api.SendVerifyCodeReq, callOptions ...callopt.Option) (r *core_api.SendVerifyCodeResp, err error)
	SetPassword(ctx context.Context, Req *core_api.SetPasswordReq, callOptions ...callopt.Option) (r *core_api.SetPasswordResp, err error)
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
	return &kAuthClient{
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

type kAuthClient struct {
	*kClient
}

func (p *kAuthClient) SignIn(ctx context.Context, Req *core_api.SignInReq, callOptions ...callopt.Option) (r *core_api.SignInResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SignIn(ctx, Req)
}

func (p *kAuthClient) SendVerifyCode(ctx context.Context, Req *core_api.SendVerifyCodeReq, callOptions ...callopt.Option) (r *core_api.SendVerifyCodeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SendVerifyCode(ctx, Req)
}

func (p *kAuthClient) SetPassword(ctx context.Context, Req *core_api.SetPasswordReq, callOptions ...callopt.Option) (r *core_api.SetPasswordResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SetPassword(ctx, Req)
}
