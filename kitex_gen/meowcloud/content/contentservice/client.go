// Code generated by Kitex v0.10.3. DO NOT EDIT.

package contentservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	content "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowcloud/content"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateCatAlbum(ctx context.Context, Req *content.CreateCatAlbumReq, callOptions ...callopt.Option) (r *content.CreateCatAlbumResp, err error)
	RetrieveCatAlbum(ctx context.Context, Req *content.RetrieveCatAlbumReq, callOptions ...callopt.Option) (r *content.RetrieveCatAlbumResp, err error)
	UpdateCatAlbum(ctx context.Context, Req *content.UpdateCatAlbumReq, callOptions ...callopt.Option) (r *content.UpdateCatAlbumResp, err error)
	DeleteCatAlbum(ctx context.Context, Req *content.DeleteCatAlbumReq, callOptions ...callopt.Option) (r *content.DeleteCatAlbumResp, err error)
	ListCatAlbum(ctx context.Context, Req *content.ListCatAlbumReq, callOptions ...callopt.Option) (r *content.ListCatAlbumResp, err error)
	CreateLocationAlbum(ctx context.Context, Req *content.CreateLocationAlbumReq, callOptions ...callopt.Option) (r *content.CreateLocationAlbumResp, err error)
	RetrieveLocationAlbum(ctx context.Context, Req *content.RetrieveLocationAlbumReq, callOptions ...callopt.Option) (r *content.RetrieveLocationAlbumResp, err error)
	UpdateLocationAlbum(ctx context.Context, Req *content.UpdateLocationAlbumReq, callOptions ...callopt.Option) (r *content.UpdateLocationAlbumResp, err error)
	DeleteLocationAlbum(ctx context.Context, Req *content.DeleteLocationAlbumReq, callOptions ...callopt.Option) (r *content.DeleteLocationAlbumResp, err error)
	ListLocationAlbum(ctx context.Context, Req *content.ListLocationAlbumReq, callOptions ...callopt.Option) (r *content.ListLocationAlbumResp, err error)
	CreatePhoto(ctx context.Context, Req *content.CreatePhotoReq, callOptions ...callopt.Option) (r *content.CreatePhotoResp, err error)
	RetrievePhoto(ctx context.Context, Req *content.RetrievePhotoReq, callOptions ...callopt.Option) (r *content.RetrievePhotoResp, err error)
	UpdatePhoto(ctx context.Context, Req *content.UpdatePhotoReq, callOptions ...callopt.Option) (r *content.UpdatePhotoResp, err error)
	DeletePhoto(ctx context.Context, Req *content.DeletePhotoReq, callOptions ...callopt.Option) (r *content.DeletePhotoResp, err error)
	ListPhoto(ctx context.Context, Req *content.ListPhotoReq, callOptions ...callopt.Option) (r *content.ListPhotoResp, err error)
	ListFeaturedPhoto(ctx context.Context, Req *content.ListFeaturedPhotoReq, callOptions ...callopt.Option) (r *content.ListFeaturedPhotoResp, err error)
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
	return &kContentServiceClient{
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

type kContentServiceClient struct {
	*kClient
}

func (p *kContentServiceClient) CreateCatAlbum(ctx context.Context, Req *content.CreateCatAlbumReq, callOptions ...callopt.Option) (r *content.CreateCatAlbumResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateCatAlbum(ctx, Req)
}

func (p *kContentServiceClient) RetrieveCatAlbum(ctx context.Context, Req *content.RetrieveCatAlbumReq, callOptions ...callopt.Option) (r *content.RetrieveCatAlbumResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RetrieveCatAlbum(ctx, Req)
}

func (p *kContentServiceClient) UpdateCatAlbum(ctx context.Context, Req *content.UpdateCatAlbumReq, callOptions ...callopt.Option) (r *content.UpdateCatAlbumResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateCatAlbum(ctx, Req)
}

func (p *kContentServiceClient) DeleteCatAlbum(ctx context.Context, Req *content.DeleteCatAlbumReq, callOptions ...callopt.Option) (r *content.DeleteCatAlbumResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteCatAlbum(ctx, Req)
}

func (p *kContentServiceClient) ListCatAlbum(ctx context.Context, Req *content.ListCatAlbumReq, callOptions ...callopt.Option) (r *content.ListCatAlbumResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListCatAlbum(ctx, Req)
}

func (p *kContentServiceClient) CreateLocationAlbum(ctx context.Context, Req *content.CreateLocationAlbumReq, callOptions ...callopt.Option) (r *content.CreateLocationAlbumResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateLocationAlbum(ctx, Req)
}

func (p *kContentServiceClient) RetrieveLocationAlbum(ctx context.Context, Req *content.RetrieveLocationAlbumReq, callOptions ...callopt.Option) (r *content.RetrieveLocationAlbumResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RetrieveLocationAlbum(ctx, Req)
}

func (p *kContentServiceClient) UpdateLocationAlbum(ctx context.Context, Req *content.UpdateLocationAlbumReq, callOptions ...callopt.Option) (r *content.UpdateLocationAlbumResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateLocationAlbum(ctx, Req)
}

func (p *kContentServiceClient) DeleteLocationAlbum(ctx context.Context, Req *content.DeleteLocationAlbumReq, callOptions ...callopt.Option) (r *content.DeleteLocationAlbumResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteLocationAlbum(ctx, Req)
}

func (p *kContentServiceClient) ListLocationAlbum(ctx context.Context, Req *content.ListLocationAlbumReq, callOptions ...callopt.Option) (r *content.ListLocationAlbumResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListLocationAlbum(ctx, Req)
}

func (p *kContentServiceClient) CreatePhoto(ctx context.Context, Req *content.CreatePhotoReq, callOptions ...callopt.Option) (r *content.CreatePhotoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreatePhoto(ctx, Req)
}

func (p *kContentServiceClient) RetrievePhoto(ctx context.Context, Req *content.RetrievePhotoReq, callOptions ...callopt.Option) (r *content.RetrievePhotoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RetrievePhoto(ctx, Req)
}

func (p *kContentServiceClient) UpdatePhoto(ctx context.Context, Req *content.UpdatePhotoReq, callOptions ...callopt.Option) (r *content.UpdatePhotoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdatePhoto(ctx, Req)
}

func (p *kContentServiceClient) DeletePhoto(ctx context.Context, Req *content.DeletePhotoReq, callOptions ...callopt.Option) (r *content.DeletePhotoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeletePhoto(ctx, Req)
}

func (p *kContentServiceClient) ListPhoto(ctx context.Context, Req *content.ListPhotoReq, callOptions ...callopt.Option) (r *content.ListPhotoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListPhoto(ctx, Req)
}

func (p *kContentServiceClient) ListFeaturedPhoto(ctx context.Context, Req *content.ListFeaturedPhotoReq, callOptions ...callopt.Option) (r *content.ListFeaturedPhotoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListFeaturedPhoto(ctx, Req)
}