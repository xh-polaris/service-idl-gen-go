// Code generated by Kitex v0.7.3. DO NOT EDIT.

package plan

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	core_api "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/core_api"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetPlanPreviews(ctx context.Context, Req *core_api.GetPlanPreviewsReq, callOptions ...callopt.Option) (r *core_api.GetPlanPreviewsResp, err error)
	GetPlanDetail(ctx context.Context, Req *core_api.GetPlanDetailReq, callOptions ...callopt.Option) (r *core_api.GetPlanDetailResp, err error)
	NewPlan(ctx context.Context, Req *core_api.NewPlanReq, callOptions ...callopt.Option) (r *core_api.NewPlanResp, err error)
	DeletePlan(ctx context.Context, Req *core_api.DeletePlanReq, callOptions ...callopt.Option) (r *core_api.DeletePlanResp, err error)
	DonateFish(ctx context.Context, Req *core_api.DonateFishReq, callOptions ...callopt.Option) (r *core_api.DonateFishResp, err error)
	GetUserFish(ctx context.Context, Req *core_api.GetUserFishReq, callOptions ...callopt.Option) (r *core_api.GetUserFishResp, err error)
	ListFishByPlan(ctx context.Context, Req *core_api.ListFishByPlanReq, callOptions ...callopt.Option) (r *core_api.ListFishByPlanResp, err error)
	ListDonateByUser(ctx context.Context, Req *core_api.ListDonateByUserReq, callOptions ...callopt.Option) (r *core_api.ListDonateByUserResp, err error)
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
	return &kPlanClient{
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

type kPlanClient struct {
	*kClient
}

func (p *kPlanClient) GetPlanPreviews(ctx context.Context, Req *core_api.GetPlanPreviewsReq, callOptions ...callopt.Option) (r *core_api.GetPlanPreviewsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPlanPreviews(ctx, Req)
}

func (p *kPlanClient) GetPlanDetail(ctx context.Context, Req *core_api.GetPlanDetailReq, callOptions ...callopt.Option) (r *core_api.GetPlanDetailResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPlanDetail(ctx, Req)
}

func (p *kPlanClient) NewPlan(ctx context.Context, Req *core_api.NewPlanReq, callOptions ...callopt.Option) (r *core_api.NewPlanResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.NewPlan(ctx, Req)
}

func (p *kPlanClient) DeletePlan(ctx context.Context, Req *core_api.DeletePlanReq, callOptions ...callopt.Option) (r *core_api.DeletePlanResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeletePlan(ctx, Req)
}

func (p *kPlanClient) DonateFish(ctx context.Context, Req *core_api.DonateFishReq, callOptions ...callopt.Option) (r *core_api.DonateFishResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DonateFish(ctx, Req)
}

func (p *kPlanClient) GetUserFish(ctx context.Context, Req *core_api.GetUserFishReq, callOptions ...callopt.Option) (r *core_api.GetUserFishResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserFish(ctx, Req)
}

func (p *kPlanClient) ListFishByPlan(ctx context.Context, Req *core_api.ListFishByPlanReq, callOptions ...callopt.Option) (r *core_api.ListFishByPlanResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListFishByPlan(ctx, Req)
}

func (p *kPlanClient) ListDonateByUser(ctx context.Context, Req *core_api.ListDonateByUserReq, callOptions ...callopt.Option) (r *core_api.ListDonateByUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListDonateByUser(ctx, Req)
}
