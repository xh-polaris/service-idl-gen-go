// Code generated by Kitex v0.7.2. DO NOT EDIT.

package commentservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	comment "github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/comment"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateComment(ctx context.Context, Req *comment.CreateCommentReq, callOptions ...callopt.Option) (r *comment.CreateCommentResp, err error)
	UpdateComment(ctx context.Context, Req *comment.UpdateCommentReq, callOptions ...callopt.Option) (r *comment.UpdateCommentResp, err error)
	DeleteComment(ctx context.Context, Req *comment.DeleteCommentByIdReq, callOptions ...callopt.Option) (r *comment.DeleteCommentByIdResp, err error)
	ListCommentByParentOrFirstLevelId(ctx context.Context, Req *comment.ListCommentByParentOrFirstLevelIdReq, callOptions ...callopt.Option) (r *comment.ListCommentByParentOrFirstLevelIdResp, err error)
	CountCommentByParent(ctx context.Context, Req *comment.CountCommentByParentReq, callOptions ...callopt.Option) (r *comment.CountCommentByParentResp, err error)
	RetrieveCommentById(ctx context.Context, Req *comment.RetrieveCommentByIdReq, callOptions ...callopt.Option) (r *comment.RetrieveCommentByIdResp, err error)
	ListCommentByAuthorIdAndType(ctx context.Context, Req *comment.ListCommentByAuthorIdAndTypeReq, callOptions ...callopt.Option) (r *comment.ListCommentByAuthorIdAndTypeResp, err error)
	ListCommentByReplyToAndType(ctx context.Context, Req *comment.ListCommentByReplyToAndTypeReq, callOptions ...callopt.Option) (r *comment.ListCommentByReplyToAndTypeResp, err error)
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
	return &kCommentServiceClient{
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

type kCommentServiceClient struct {
	*kClient
}

func (p *kCommentServiceClient) CreateComment(ctx context.Context, Req *comment.CreateCommentReq, callOptions ...callopt.Option) (r *comment.CreateCommentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateComment(ctx, Req)
}

func (p *kCommentServiceClient) UpdateComment(ctx context.Context, Req *comment.UpdateCommentReq, callOptions ...callopt.Option) (r *comment.UpdateCommentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateComment(ctx, Req)
}

func (p *kCommentServiceClient) DeleteComment(ctx context.Context, Req *comment.DeleteCommentByIdReq, callOptions ...callopt.Option) (r *comment.DeleteCommentByIdResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteComment(ctx, Req)
}

func (p *kCommentServiceClient) ListCommentByParentOrFirstLevelId(ctx context.Context, Req *comment.ListCommentByParentOrFirstLevelIdReq, callOptions ...callopt.Option) (r *comment.ListCommentByParentOrFirstLevelIdResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListCommentByParentOrFirstLevelId(ctx, Req)
}

func (p *kCommentServiceClient) CountCommentByParent(ctx context.Context, Req *comment.CountCommentByParentReq, callOptions ...callopt.Option) (r *comment.CountCommentByParentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CountCommentByParent(ctx, Req)
}

func (p *kCommentServiceClient) RetrieveCommentById(ctx context.Context, Req *comment.RetrieveCommentByIdReq, callOptions ...callopt.Option) (r *comment.RetrieveCommentByIdResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RetrieveCommentById(ctx, Req)
}

func (p *kCommentServiceClient) ListCommentByAuthorIdAndType(ctx context.Context, Req *comment.ListCommentByAuthorIdAndTypeReq, callOptions ...callopt.Option) (r *comment.ListCommentByAuthorIdAndTypeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListCommentByAuthorIdAndType(ctx, Req)
}

func (p *kCommentServiceClient) ListCommentByReplyToAndType(ctx context.Context, Req *comment.ListCommentByReplyToAndTypeReq, callOptions ...callopt.Option) (r *comment.ListCommentByReplyToAndTypeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListCommentByReplyToAndType(ctx, Req)
}
