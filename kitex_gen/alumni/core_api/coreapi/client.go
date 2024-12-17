// Code generated by Kitex v0.12.0. DO NOT EDIT.

package coreapi

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	core_api "github.com/xh-polaris/service-idl-gen-go/kitex_gen/alumni/core_api"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	SignUp(ctx context.Context, Req *core_api.SignUpReq, callOptions ...callopt.Option) (r *core_api.SignUpResp, err error)
	SignIn(ctx context.Context, Req *core_api.SignInReq, callOptions ...callopt.Option) (r *core_api.SignInResp, err error)
	UpdateUserInfo(ctx context.Context, Req *core_api.UpdateUserInfoReq, callOptions ...callopt.Option) (r *core_api.Response, err error)
	UpdateEducation(ctx context.Context, Req *core_api.UpdateEducationReq, callOptions ...callopt.Option) (r *core_api.Response, err error)
	UpdateEmployment(ctx context.Context, Req *core_api.UpdateEmploymentReq, callOptions ...callopt.Option) (r *core_api.Response, err error)
	GetUserInfo(ctx context.Context, Req *core_api.GetUserInfoReq, callOptions ...callopt.Option) (r *core_api.GetUserInfoResp, err error)
	CreateActivity(ctx context.Context, Req *core_api.CreateActivityReq, callOptions ...callopt.Option) (r *core_api.Response, err error)
	UpdateActivity(ctx context.Context, Req *core_api.UpdateActivityReq, callOptions ...callopt.Option) (r *core_api.Response, err error)
	GetActivities(ctx context.Context, Req *core_api.GetActivitiesReq, callOptions ...callopt.Option) (r *core_api.GetActivitiesResp, err error)
	GetActivity(ctx context.Context, Req *core_api.GetActivityReq, callOptions ...callopt.Option) (r *core_api.GetActivityResp, err error)
	RegisterActivity(ctx context.Context, Req *core_api.RegisterActivityReq, callOptions ...callopt.Option) (r *core_api.Response, err error)
	CheckIn(ctx context.Context, Req *core_api.CheckInReq, callOptions ...callopt.Option) (r *core_api.Response, err error)
	GetRegisters(ctx context.Context, Req *core_api.GetRegistersReq, callOptions ...callopt.Option) (r *core_api.GetRegisterResp, err error)
	ApplySignedUrl(ctx context.Context, Req *core_api.ApplySignedUrlReq, callOptions ...callopt.Option) (r *core_api.ApplySignedUrlResp, err error)
	SendVerifyCode(ctx context.Context, Req *core_api.SendVerifyCodeReq, callOptions ...callopt.Option) (r *core_api.Response, err error)
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
	return &kCoreApiClient{
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

type kCoreApiClient struct {
	*kClient
}

func (p *kCoreApiClient) SignUp(ctx context.Context, Req *core_api.SignUpReq, callOptions ...callopt.Option) (r *core_api.SignUpResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SignUp(ctx, Req)
}

func (p *kCoreApiClient) SignIn(ctx context.Context, Req *core_api.SignInReq, callOptions ...callopt.Option) (r *core_api.SignInResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SignIn(ctx, Req)
}

func (p *kCoreApiClient) UpdateUserInfo(ctx context.Context, Req *core_api.UpdateUserInfoReq, callOptions ...callopt.Option) (r *core_api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateUserInfo(ctx, Req)
}

func (p *kCoreApiClient) UpdateEducation(ctx context.Context, Req *core_api.UpdateEducationReq, callOptions ...callopt.Option) (r *core_api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateEducation(ctx, Req)
}

func (p *kCoreApiClient) UpdateEmployment(ctx context.Context, Req *core_api.UpdateEmploymentReq, callOptions ...callopt.Option) (r *core_api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateEmployment(ctx, Req)
}

func (p *kCoreApiClient) GetUserInfo(ctx context.Context, Req *core_api.GetUserInfoReq, callOptions ...callopt.Option) (r *core_api.GetUserInfoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserInfo(ctx, Req)
}

func (p *kCoreApiClient) CreateActivity(ctx context.Context, Req *core_api.CreateActivityReq, callOptions ...callopt.Option) (r *core_api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateActivity(ctx, Req)
}

func (p *kCoreApiClient) UpdateActivity(ctx context.Context, Req *core_api.UpdateActivityReq, callOptions ...callopt.Option) (r *core_api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateActivity(ctx, Req)
}

func (p *kCoreApiClient) GetActivities(ctx context.Context, Req *core_api.GetActivitiesReq, callOptions ...callopt.Option) (r *core_api.GetActivitiesResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetActivities(ctx, Req)
}

func (p *kCoreApiClient) GetActivity(ctx context.Context, Req *core_api.GetActivityReq, callOptions ...callopt.Option) (r *core_api.GetActivityResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetActivity(ctx, Req)
}

func (p *kCoreApiClient) RegisterActivity(ctx context.Context, Req *core_api.RegisterActivityReq, callOptions ...callopt.Option) (r *core_api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RegisterActivity(ctx, Req)
}

func (p *kCoreApiClient) CheckIn(ctx context.Context, Req *core_api.CheckInReq, callOptions ...callopt.Option) (r *core_api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CheckIn(ctx, Req)
}

func (p *kCoreApiClient) GetRegisters(ctx context.Context, Req *core_api.GetRegistersReq, callOptions ...callopt.Option) (r *core_api.GetRegisterResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetRegisters(ctx, Req)
}

func (p *kCoreApiClient) ApplySignedUrl(ctx context.Context, Req *core_api.ApplySignedUrlReq, callOptions ...callopt.Option) (r *core_api.ApplySignedUrlResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ApplySignedUrl(ctx, Req)
}

func (p *kCoreApiClient) SendVerifyCode(ctx context.Context, Req *core_api.SendVerifyCodeReq, callOptions ...callopt.Option) (r *core_api.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SendVerifyCode(ctx, Req)
}