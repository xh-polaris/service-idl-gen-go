// Code generated by Kitex v0.10.3. DO NOT EDIT.

package collection

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	core_api "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/core_api"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetCatPreviews(ctx context.Context, Req *core_api.GetCatPreviewsReq, callOptions ...callopt.Option) (r *core_api.GetCatPreviewsResp, err error)
	GetCatDetail(ctx context.Context, Req *core_api.GetCatDetailReq, callOptions ...callopt.Option) (r *core_api.GetCatDetailResp, err error)
	NewCat(ctx context.Context, Req *core_api.NewCatReq, callOptions ...callopt.Option) (r *core_api.NewCatResp, err error)
	DeleteCat(ctx context.Context, Req *core_api.DeleteCatReq, callOptions ...callopt.Option) (r *core_api.DeleteCatResp, err error)
	SearchCat(ctx context.Context, Req *core_api.SearchCatReq, callOptions ...callopt.Option) (r *core_api.SearchCatResp, err error)
	CreateImage(ctx context.Context, Req *core_api.CreateImageReq, callOptions ...callopt.Option) (r *core_api.CreateImageResp, err error)
	DeleteImage(ctx context.Context, Req *core_api.DeleteImageReq, callOptions ...callopt.Option) (r *core_api.DeleteImageResp, err error)
	GetImageByCat(ctx context.Context, Req *core_api.GetImageByCatReq, callOptions ...callopt.Option) (r *core_api.GetImageByCatResp, err error)
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
	return &kCollectionClient{
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

type kCollectionClient struct {
	*kClient
}

func (p *kCollectionClient) GetCatPreviews(ctx context.Context, Req *core_api.GetCatPreviewsReq, callOptions ...callopt.Option) (r *core_api.GetCatPreviewsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetCatPreviews(ctx, Req)
}

func (p *kCollectionClient) GetCatDetail(ctx context.Context, Req *core_api.GetCatDetailReq, callOptions ...callopt.Option) (r *core_api.GetCatDetailResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetCatDetail(ctx, Req)
}

func (p *kCollectionClient) NewCat(ctx context.Context, Req *core_api.NewCatReq, callOptions ...callopt.Option) (r *core_api.NewCatResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.NewCat(ctx, Req)
}

func (p *kCollectionClient) DeleteCat(ctx context.Context, Req *core_api.DeleteCatReq, callOptions ...callopt.Option) (r *core_api.DeleteCatResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteCat(ctx, Req)
}

func (p *kCollectionClient) SearchCat(ctx context.Context, Req *core_api.SearchCatReq, callOptions ...callopt.Option) (r *core_api.SearchCatResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SearchCat(ctx, Req)
}

func (p *kCollectionClient) CreateImage(ctx context.Context, Req *core_api.CreateImageReq, callOptions ...callopt.Option) (r *core_api.CreateImageResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateImage(ctx, Req)
}

func (p *kCollectionClient) DeleteImage(ctx context.Context, Req *core_api.DeleteImageReq, callOptions ...callopt.Option) (r *core_api.DeleteImageResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteImage(ctx, Req)
}

func (p *kCollectionClient) GetImageByCat(ctx context.Context, Req *core_api.GetImageByCatReq, callOptions ...callopt.Option) (r *core_api.GetImageByCatResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetImageByCat(ctx, Req)
}
