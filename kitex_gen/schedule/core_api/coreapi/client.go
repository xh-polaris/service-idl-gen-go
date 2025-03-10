// Code generated by Kitex v0.12.1. DO NOT EDIT.

package coreapi

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	core_api "github.com/xh-polaris/service-idl-gen-go/kitex_gen/schedule/core_api"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	SignUp(ctx context.Context, Req *core_api.SignUpReq, callOptions ...callopt.Option) (r *core_api.SignUpResp, err error)
	SignIn(ctx context.Context, Req *core_api.SignInReq, callOptions ...callopt.Option) (r *core_api.SignInResp, err error)
	SendVerifyCode(ctx context.Context, Req *core_api.SendVerifyCodeReq, callOptions ...callopt.Option) (r *core_api.SendVerifyCodeResp, err error)
	CreateSchedule(ctx context.Context, Req *core_api.CreateScheduleReq, callOptions ...callopt.Option) (r *core_api.CreateScheduleResp, err error)
	CreateScheduleFromOrigin(ctx context.Context, Req *core_api.CreateScheduleFromOriginReq, callOptions ...callopt.Option) (r *core_api.CreateScheduleFromOriginResp, err error)
	CreateSchedules(ctx context.Context, Req *core_api.CreateSchedulesReq, callOptions ...callopt.Option) (r *core_api.CreateSchedulesResp, err error)
	UpdateSchedule(ctx context.Context, Req *core_api.UpdateScheduleReq, callOptions ...callopt.Option) (r *core_api.UpdateScheduleResp, err error)
	GetSchedules(ctx context.Context, Req *core_api.GetSchedulesReq, callOptions ...callopt.Option) (r *core_api.GetSchedulesResp, err error)
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

func (p *kCoreApiClient) SendVerifyCode(ctx context.Context, Req *core_api.SendVerifyCodeReq, callOptions ...callopt.Option) (r *core_api.SendVerifyCodeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SendVerifyCode(ctx, Req)
}

func (p *kCoreApiClient) CreateSchedule(ctx context.Context, Req *core_api.CreateScheduleReq, callOptions ...callopt.Option) (r *core_api.CreateScheduleResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateSchedule(ctx, Req)
}

func (p *kCoreApiClient) CreateScheduleFromOrigin(ctx context.Context, Req *core_api.CreateScheduleFromOriginReq, callOptions ...callopt.Option) (r *core_api.CreateScheduleFromOriginResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateScheduleFromOrigin(ctx, Req)
}

func (p *kCoreApiClient) CreateSchedules(ctx context.Context, Req *core_api.CreateSchedulesReq, callOptions ...callopt.Option) (r *core_api.CreateSchedulesResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateSchedules(ctx, Req)
}

func (p *kCoreApiClient) UpdateSchedule(ctx context.Context, Req *core_api.UpdateScheduleReq, callOptions ...callopt.Option) (r *core_api.UpdateScheduleResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateSchedule(ctx, Req)
}

func (p *kCoreApiClient) GetSchedules(ctx context.Context, Req *core_api.GetSchedulesReq, callOptions ...callopt.Option) (r *core_api.GetSchedulesResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetSchedules(ctx, Req)
}
