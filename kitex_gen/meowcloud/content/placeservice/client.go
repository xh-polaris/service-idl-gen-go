// Code generated by Kitex v0.10.3. DO NOT EDIT.

package placeservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	content "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowcloud/content"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AddPlace(ctx context.Context, Req *content.AddPlaceReq, callOptions ...callopt.Option) (r *content.AddPlaceResp, err error)
	UpdatePlace(ctx context.Context, Req *content.UpdatePlaceReq, callOptions ...callopt.Option) (r *content.UpdatePlaceResp, err error)
	GetPlace(ctx context.Context, Req *content.GetPlaceReq, callOptions ...callopt.Option) (r *content.GetPlaceResp, err error)
	DeletePlace(ctx context.Context, Req *content.DeletePlaceReq, callOptions ...callopt.Option) (r *content.DeletePlaceResp, err error)
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
	return &kPlaceServiceClient{
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

type kPlaceServiceClient struct {
	*kClient
}

func (p *kPlaceServiceClient) AddPlace(ctx context.Context, Req *content.AddPlaceReq, callOptions ...callopt.Option) (r *content.AddPlaceResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddPlace(ctx, Req)
}

func (p *kPlaceServiceClient) UpdatePlace(ctx context.Context, Req *content.UpdatePlaceReq, callOptions ...callopt.Option) (r *content.UpdatePlaceResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdatePlace(ctx, Req)
}

func (p *kPlaceServiceClient) GetPlace(ctx context.Context, Req *content.GetPlaceReq, callOptions ...callopt.Option) (r *content.GetPlaceResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPlace(ctx, Req)
}

func (p *kPlaceServiceClient) DeletePlace(ctx context.Context, Req *content.DeletePlaceReq, callOptions ...callopt.Option) (r *content.DeletePlaceResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeletePlace(ctx, Req)
}
