// Code generated by Kitex v0.12.0. DO NOT EDIT.

package show

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	show "github.com/xh-polaris/service-idl-gen-go/kitex_gen/essay/show"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	SignUp(ctx context.Context, Req *show.SignUpReq, callOptions ...callopt.Option) (r *show.SignUpResp, err error)
	SignIn(ctx context.Context, Req *show.SignInReq, callOptions ...callopt.Option) (r *show.SignInResp, err error)
	GetUserInfo(ctx context.Context, Req *show.GetUserInfoReq, callOptions ...callopt.Option) (r *show.GetUserInfoResp, err error)
	UpdatePassword(ctx context.Context, Req *show.UpdatePasswordReq, callOptions ...callopt.Option) (r *show.UpdatePasswordReq, err error)
	UpdateUserInfo(ctx context.Context, Req *show.UpdateUserInfoReq, callOptions ...callopt.Option) (r *show.Response, err error)
	EssayEvaluate(ctx context.Context, Req *show.EssayEvaluateReq, callOptions ...callopt.Option) (r *show.EssayEvaluateResp, err error)
	GetEvaluateLogs(ctx context.Context, Req *show.GetEssayEvaluateLogsReq, callOptions ...callopt.Option) (r *show.GetEssayEvaluateLogsResp, err error)
	OCR(ctx context.Context, Req *show.OCRReq, callOptions ...callopt.Option) (r *show.OCRResp, err error)
	ApplySignedUrl(ctx context.Context, Req *show.ApplySignedUrlReq, callOptions ...callopt.Option) (r *show.ApplySignedUrlResp, err error)
	SendVerifyCode(ctx context.Context, Req *show.SendVerifyCodeReq, callOptions ...callopt.Option) (r *show.Response, err error)
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
	return &kShowClient{
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

type kShowClient struct {
	*kClient
}

func (p *kShowClient) SignUp(ctx context.Context, Req *show.SignUpReq, callOptions ...callopt.Option) (r *show.SignUpResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SignUp(ctx, Req)
}

func (p *kShowClient) SignIn(ctx context.Context, Req *show.SignInReq, callOptions ...callopt.Option) (r *show.SignInResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SignIn(ctx, Req)
}

func (p *kShowClient) GetUserInfo(ctx context.Context, Req *show.GetUserInfoReq, callOptions ...callopt.Option) (r *show.GetUserInfoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserInfo(ctx, Req)
}

func (p *kShowClient) UpdatePassword(ctx context.Context, Req *show.UpdatePasswordReq, callOptions ...callopt.Option) (r *show.UpdatePasswordReq, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdatePassword(ctx, Req)
}

func (p *kShowClient) UpdateUserInfo(ctx context.Context, Req *show.UpdateUserInfoReq, callOptions ...callopt.Option) (r *show.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateUserInfo(ctx, Req)
}

func (p *kShowClient) EssayEvaluate(ctx context.Context, Req *show.EssayEvaluateReq, callOptions ...callopt.Option) (r *show.EssayEvaluateResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EssayEvaluate(ctx, Req)
}

func (p *kShowClient) GetEvaluateLogs(ctx context.Context, Req *show.GetEssayEvaluateLogsReq, callOptions ...callopt.Option) (r *show.GetEssayEvaluateLogsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetEvaluateLogs(ctx, Req)
}

func (p *kShowClient) OCR(ctx context.Context, Req *show.OCRReq, callOptions ...callopt.Option) (r *show.OCRResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.OCR(ctx, Req)
}

func (p *kShowClient) ApplySignedUrl(ctx context.Context, Req *show.ApplySignedUrlReq, callOptions ...callopt.Option) (r *show.ApplySignedUrlResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ApplySignedUrl(ctx, Req)
}

func (p *kShowClient) SendVerifyCode(ctx context.Context, Req *show.SendVerifyCodeReq, callOptions ...callopt.Option) (r *show.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SendVerifyCode(ctx, Req)
}