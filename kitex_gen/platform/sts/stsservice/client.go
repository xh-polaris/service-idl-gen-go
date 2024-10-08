// Code generated by Kitex v0.11.3. DO NOT EDIT.

package stsservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	sts "github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/sts"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GenCosSts(ctx context.Context, Req *sts.GenCosStsReq, callOptions ...callopt.Option) (r *sts.GenCosStsResp, err error)
	GenSignedUrl(ctx context.Context, Req *sts.GenSignedUrlReq, callOptions ...callopt.Option) (r *sts.GenSignedUrlResp, err error)
	DeleteObject(ctx context.Context, Req *sts.DeleteObjectReq, callOptions ...callopt.Option) (r *sts.DeleteObjectResp, err error)
	TextCheck(ctx context.Context, Req *sts.TextCheckReq, callOptions ...callopt.Option) (r *sts.TextCheckResp, err error)
	PhotoCheck(ctx context.Context, Req *sts.PhotoCheckReq, callOptions ...callopt.Option) (r *sts.PhotoCheckResp, err error)
	SignIn(ctx context.Context, Req *sts.SignInReq, callOptions ...callopt.Option) (r *sts.SignInResp, err error)
	SetPassword(ctx context.Context, Req *sts.SetPasswordReq, callOptions ...callopt.Option) (r *sts.SetPasswordResp, err error)
	SendVerifyCode(ctx context.Context, Req *sts.SendVerifyCodeReq, callOptions ...callopt.Option) (r *sts.SendVerifyCodeResp, err error)
	AddUserAuth(ctx context.Context, Req *sts.AddUserAuthReq, callOptions ...callopt.Option) (r *sts.AddUserAuthResp, err error)
	SendMessage(ctx context.Context, Req *sts.SendMessageReq, callOptions ...callopt.Option) (r *sts.SendMessageResp, err error)
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
	return &kStsServiceClient{
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

type kStsServiceClient struct {
	*kClient
}

func (p *kStsServiceClient) GenCosSts(ctx context.Context, Req *sts.GenCosStsReq, callOptions ...callopt.Option) (r *sts.GenCosStsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GenCosSts(ctx, Req)
}

func (p *kStsServiceClient) GenSignedUrl(ctx context.Context, Req *sts.GenSignedUrlReq, callOptions ...callopt.Option) (r *sts.GenSignedUrlResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GenSignedUrl(ctx, Req)
}

func (p *kStsServiceClient) DeleteObject(ctx context.Context, Req *sts.DeleteObjectReq, callOptions ...callopt.Option) (r *sts.DeleteObjectResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteObject(ctx, Req)
}

func (p *kStsServiceClient) TextCheck(ctx context.Context, Req *sts.TextCheckReq, callOptions ...callopt.Option) (r *sts.TextCheckResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.TextCheck(ctx, Req)
}

func (p *kStsServiceClient) PhotoCheck(ctx context.Context, Req *sts.PhotoCheckReq, callOptions ...callopt.Option) (r *sts.PhotoCheckResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PhotoCheck(ctx, Req)
}

func (p *kStsServiceClient) SignIn(ctx context.Context, Req *sts.SignInReq, callOptions ...callopt.Option) (r *sts.SignInResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SignIn(ctx, Req)
}

func (p *kStsServiceClient) SetPassword(ctx context.Context, Req *sts.SetPasswordReq, callOptions ...callopt.Option) (r *sts.SetPasswordResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SetPassword(ctx, Req)
}

func (p *kStsServiceClient) SendVerifyCode(ctx context.Context, Req *sts.SendVerifyCodeReq, callOptions ...callopt.Option) (r *sts.SendVerifyCodeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SendVerifyCode(ctx, Req)
}

func (p *kStsServiceClient) AddUserAuth(ctx context.Context, Req *sts.AddUserAuthReq, callOptions ...callopt.Option) (r *sts.AddUserAuthResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddUserAuth(ctx, Req)
}

func (p *kStsServiceClient) SendMessage(ctx context.Context, Req *sts.SendMessageReq, callOptions ...callopt.Option) (r *sts.SendMessageResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SendMessage(ctx, Req)
}
