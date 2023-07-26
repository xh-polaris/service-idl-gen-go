// Code generated by Kitex v0.6.2. DO NOT EDIT.

package momentrpc

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	moment "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/moment"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	ListMoment(ctx context.Context, Req *moment.ListMomentReq, callOptions ...callopt.Option) (r *moment.ListMomentResp, err error)
	CountMoment(ctx context.Context, Req *moment.CountMomentReq, callOptions ...callopt.Option) (r *moment.CountMomentResp, err error)
	RetrieveMoment(ctx context.Context, Req *moment.RetrieveMomentReq, callOptions ...callopt.Option) (r *moment.RetrieveMomentResp, err error)
	CreateMoment(ctx context.Context, Req *moment.CreateMomentReq, callOptions ...callopt.Option) (r *moment.CreateMomentResp, err error)
	UpdateMoment(ctx context.Context, Req *moment.UpdateMomentReq, callOptions ...callopt.Option) (r *moment.UpdateMomentResp, err error)
	DeleteMoment(ctx context.Context, Req *moment.DeleteMomentReq, callOptions ...callopt.Option) (r *moment.DeleteMomentResp, err error)
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
	return &kMomentRpcClient{
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

type kMomentRpcClient struct {
	*kClient
}

func (p *kMomentRpcClient) ListMoment(ctx context.Context, Req *moment.ListMomentReq, callOptions ...callopt.Option) (r *moment.ListMomentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListMoment(ctx, Req)
}

func (p *kMomentRpcClient) CountMoment(ctx context.Context, Req *moment.CountMomentReq, callOptions ...callopt.Option) (r *moment.CountMomentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CountMoment(ctx, Req)
}

func (p *kMomentRpcClient) RetrieveMoment(ctx context.Context, Req *moment.RetrieveMomentReq, callOptions ...callopt.Option) (r *moment.RetrieveMomentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RetrieveMoment(ctx, Req)
}

func (p *kMomentRpcClient) CreateMoment(ctx context.Context, Req *moment.CreateMomentReq, callOptions ...callopt.Option) (r *moment.CreateMomentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateMoment(ctx, Req)
}

func (p *kMomentRpcClient) UpdateMoment(ctx context.Context, Req *moment.UpdateMomentReq, callOptions ...callopt.Option) (r *moment.UpdateMomentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateMoment(ctx, Req)
}

func (p *kMomentRpcClient) DeleteMoment(ctx context.Context, Req *moment.DeleteMomentReq, callOptions ...callopt.Option) (r *moment.DeleteMomentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteMoment(ctx, Req)
}
