// Code generated by Kitex v0.11.3. DO NOT EDIT.

package charge

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	charge "github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/charge"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateBaseInterface(ctx context.Context, Req *charge.CreateBaseInterfaceReq, callOptions ...callopt.Option) (r *charge.CreateBaseInterfaceResp, err error)
	UpdateBaseInterface(ctx context.Context, Req *charge.UpdateBaseInterfaceReq, callOptions ...callopt.Option) (r *charge.UpdateBaseInterfaceResp, err error)
	DeleteBaseInterface(ctx context.Context, Req *charge.DeleteBaseInterfaceReq, callOptions ...callopt.Option) (r *charge.DeleteBaseInterfaceResp, err error)
	GetBaseInterface(ctx context.Context, Req *charge.GetBaseInterfaceReq, callOptions ...callopt.Option) (r *charge.GetBaseInterfaceResp, err error)
	CreateFullInterface(ctx context.Context, Req *charge.CreateFullInterfaceReq, callOptions ...callopt.Option) (r *charge.CreateFullInterfaceResp, err error)
	UpdateFullInterface(ctx context.Context, Req *charge.UpdateFullInterfaceReq, callOptions ...callopt.Option) (r *charge.UpdateFullInterfaceResp, err error)
	DeleteFullInterface(ctx context.Context, Req *charge.DeleteFullInterfaceReq, callOptions ...callopt.Option) (r *charge.DeleteFullInterfaceResp, err error)
	GetFullInterface(ctx context.Context, Req *charge.GetFullInterfaceReq, callOptions ...callopt.Option) (r *charge.GetFullInterfaceResp, err error)
	GetOneFullInterface(ctx context.Context, Req *charge.GetOneFullInterfaceReq, callOptions ...callopt.Option) (r *charge.GetOneFullInterfaceResp, err error)
	CreateMargin(ctx context.Context, Req *charge.CreateMarginReq, callOptions ...callopt.Option) (r *charge.CreateMarginResp, err error)
	UpdateMargin(ctx context.Context, Req *charge.UpdateMarginReq, callOptions ...callopt.Option) (r *charge.UpdateMarginResp, err error)
	GetMargin(ctx context.Context, Req *charge.GetMarginReq, callOptions ...callopt.Option) (r *charge.GetMarginResp, err error)
	DeleteMargin(ctx context.Context, Req *charge.DeleteMarginReq, callOptions ...callopt.Option) (r *charge.DeleteMarginResp, err error)
	GetFullAndBaseInterfaceForCheck(ctx context.Context, Req *charge.GetFullAndBaseInterfaceForCheckReq, callOptions ...callopt.Option) (r *charge.GetFullAndBaseInterfaceForCheckResp, err error)
	CreateGradient(ctx context.Context, Req *charge.CreateGradientReq, callOptions ...callopt.Option) (r *charge.CreateGradientResp, err error)
	UpdateGradient(ctx context.Context, Req *charge.UpdateGradientReq, callOptions ...callopt.Option) (r *charge.UpdateGradientResp, err error)
	GetGradient(ctx context.Context, Req *charge.GetGradientReq, callOptions ...callopt.Option) (r *charge.GetGradientResp, err error)
	GetAmount(ctx context.Context, Req *charge.GetAmountReq, callOptions ...callopt.Option) (r *charge.GetAmountResp, err error)
	CreateLog(ctx context.Context, Req *charge.CreateLogReq, callOptions ...callopt.Option) (r *charge.CreateLogResp, err error)
	GetLog(ctx context.Context, Req *charge.GetLogReq, callOptions ...callopt.Option) (r *charge.GetLogResp, err error)
	GetAccountByTxId(ctx context.Context, Req *charge.GetAccountByTxIdReq, callOptions ...callopt.Option) (r *charge.GetAccountByTxIdResp, err error)
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
	return &kChargeClient{
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

type kChargeClient struct {
	*kClient
}

func (p *kChargeClient) CreateBaseInterface(ctx context.Context, Req *charge.CreateBaseInterfaceReq, callOptions ...callopt.Option) (r *charge.CreateBaseInterfaceResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateBaseInterface(ctx, Req)
}

func (p *kChargeClient) UpdateBaseInterface(ctx context.Context, Req *charge.UpdateBaseInterfaceReq, callOptions ...callopt.Option) (r *charge.UpdateBaseInterfaceResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateBaseInterface(ctx, Req)
}

func (p *kChargeClient) DeleteBaseInterface(ctx context.Context, Req *charge.DeleteBaseInterfaceReq, callOptions ...callopt.Option) (r *charge.DeleteBaseInterfaceResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteBaseInterface(ctx, Req)
}

func (p *kChargeClient) GetBaseInterface(ctx context.Context, Req *charge.GetBaseInterfaceReq, callOptions ...callopt.Option) (r *charge.GetBaseInterfaceResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetBaseInterface(ctx, Req)
}

func (p *kChargeClient) CreateFullInterface(ctx context.Context, Req *charge.CreateFullInterfaceReq, callOptions ...callopt.Option) (r *charge.CreateFullInterfaceResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateFullInterface(ctx, Req)
}

func (p *kChargeClient) UpdateFullInterface(ctx context.Context, Req *charge.UpdateFullInterfaceReq, callOptions ...callopt.Option) (r *charge.UpdateFullInterfaceResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateFullInterface(ctx, Req)
}

func (p *kChargeClient) DeleteFullInterface(ctx context.Context, Req *charge.DeleteFullInterfaceReq, callOptions ...callopt.Option) (r *charge.DeleteFullInterfaceResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteFullInterface(ctx, Req)
}

func (p *kChargeClient) GetFullInterface(ctx context.Context, Req *charge.GetFullInterfaceReq, callOptions ...callopt.Option) (r *charge.GetFullInterfaceResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFullInterface(ctx, Req)
}

func (p *kChargeClient) GetOneFullInterface(ctx context.Context, Req *charge.GetOneFullInterfaceReq, callOptions ...callopt.Option) (r *charge.GetOneFullInterfaceResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetOneFullInterface(ctx, Req)
}

func (p *kChargeClient) CreateMargin(ctx context.Context, Req *charge.CreateMarginReq, callOptions ...callopt.Option) (r *charge.CreateMarginResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateMargin(ctx, Req)
}

func (p *kChargeClient) UpdateMargin(ctx context.Context, Req *charge.UpdateMarginReq, callOptions ...callopt.Option) (r *charge.UpdateMarginResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateMargin(ctx, Req)
}

func (p *kChargeClient) GetMargin(ctx context.Context, Req *charge.GetMarginReq, callOptions ...callopt.Option) (r *charge.GetMarginResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetMargin(ctx, Req)
}

func (p *kChargeClient) DeleteMargin(ctx context.Context, Req *charge.DeleteMarginReq, callOptions ...callopt.Option) (r *charge.DeleteMarginResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteMargin(ctx, Req)
}

func (p *kChargeClient) GetFullAndBaseInterfaceForCheck(ctx context.Context, Req *charge.GetFullAndBaseInterfaceForCheckReq, callOptions ...callopt.Option) (r *charge.GetFullAndBaseInterfaceForCheckResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetFullAndBaseInterfaceForCheck(ctx, Req)
}

func (p *kChargeClient) CreateGradient(ctx context.Context, Req *charge.CreateGradientReq, callOptions ...callopt.Option) (r *charge.CreateGradientResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateGradient(ctx, Req)
}

func (p *kChargeClient) UpdateGradient(ctx context.Context, Req *charge.UpdateGradientReq, callOptions ...callopt.Option) (r *charge.UpdateGradientResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateGradient(ctx, Req)
}

func (p *kChargeClient) GetGradient(ctx context.Context, Req *charge.GetGradientReq, callOptions ...callopt.Option) (r *charge.GetGradientResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetGradient(ctx, Req)
}

func (p *kChargeClient) GetAmount(ctx context.Context, Req *charge.GetAmountReq, callOptions ...callopt.Option) (r *charge.GetAmountResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetAmount(ctx, Req)
}

func (p *kChargeClient) CreateLog(ctx context.Context, Req *charge.CreateLogReq, callOptions ...callopt.Option) (r *charge.CreateLogResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateLog(ctx, Req)
}

func (p *kChargeClient) GetLog(ctx context.Context, Req *charge.GetLogReq, callOptions ...callopt.Option) (r *charge.GetLogResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetLog(ctx, Req)
}

func (p *kChargeClient) GetAccountByTxId(ctx context.Context, Req *charge.GetAccountByTxIdReq, callOptions ...callopt.Option) (r *charge.GetAccountByTxIdResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetAccountByTxId(ctx, Req)
}
