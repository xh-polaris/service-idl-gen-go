// Code generated by Kitex v0.10.3. DO NOT EDIT.

package photoservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	content "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowcloud/content"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	UploadPhoto(ctx context.Context, Req *content.UploadPhotoReq, callOptions ...callopt.Option) (r *content.UploadPhotoResp, err error)
	UpdatePhoto(ctx context.Context, Req *content.UpdatePhotoReq, callOptions ...callopt.Option) (r *content.UpdatePhotoResp, err error)
	GetPhoto(ctx context.Context, Req *content.GetPhotoReq, callOptions ...callopt.Option) (r *content.GetPhotoResp, err error)
	ListPhotos(ctx context.Context, Req *content.ListPhotosReq, callOptions ...callopt.Option) (r *content.ListPhotosResp, err error)
	DeletePhoto(ctx context.Context, Req *content.DeletePhotoReq, callOptions ...callopt.Option) (r *content.DeletePhotoResp, err error)
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
	return &kPhotoServiceClient{
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

type kPhotoServiceClient struct {
	*kClient
}

func (p *kPhotoServiceClient) UploadPhoto(ctx context.Context, Req *content.UploadPhotoReq, callOptions ...callopt.Option) (r *content.UploadPhotoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UploadPhoto(ctx, Req)
}

func (p *kPhotoServiceClient) UpdatePhoto(ctx context.Context, Req *content.UpdatePhotoReq, callOptions ...callopt.Option) (r *content.UpdatePhotoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdatePhoto(ctx, Req)
}

func (p *kPhotoServiceClient) GetPhoto(ctx context.Context, Req *content.GetPhotoReq, callOptions ...callopt.Option) (r *content.GetPhotoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetPhoto(ctx, Req)
}

func (p *kPhotoServiceClient) ListPhotos(ctx context.Context, Req *content.ListPhotosReq, callOptions ...callopt.Option) (r *content.ListPhotosResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListPhotos(ctx, Req)
}

func (p *kPhotoServiceClient) DeletePhoto(ctx context.Context, Req *content.DeletePhotoReq, callOptions ...callopt.Option) (r *content.DeletePhotoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeletePhoto(ctx, Req)
}
