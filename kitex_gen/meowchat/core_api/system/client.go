// Code generated by Kitex v0.6.2. DO NOT EDIT.

package system

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	core_api "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/core_api"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetAdmins(ctx context.Context, Req *core_api.GetAdminsReq, callOptions ...callopt.Option) (r *core_api.GetAdminsResp, err error)
	NewAdmin(ctx context.Context, Req *core_api.NewAdminReq, callOptions ...callopt.Option) (r *core_api.NewAdminResp, err error)
	DeleteAdmin(ctx context.Context, Req *core_api.DeleteAdminReq, callOptions ...callopt.Option) (r *core_api.DeleteAdminResp, err error)
	GetNews(ctx context.Context, Req *core_api.GetNewsReq, callOptions ...callopt.Option) (r *core_api.GetNewsResp, err error)
	NewNews(ctx context.Context, Req *core_api.NewNewsReq, callOptions ...callopt.Option) (r *core_api.NewNewsResp, err error)
	DeleteNews(ctx context.Context, Req *core_api.DeleteNewsReq, callOptions ...callopt.Option) (r *core_api.DeleteNewsResp, err error)
	GetNotices(ctx context.Context, Req *core_api.GetNoticesReq, callOptions ...callopt.Option) (r *core_api.GetNoticesResp, err error)
	NewNotice(ctx context.Context, Req *core_api.NewNoticeReq, callOptions ...callopt.Option) (r *core_api.NewNoticeResp, err error)
	DeleteNotice(ctx context.Context, Req *core_api.DeleteNoticeReq, callOptions ...callopt.Option) (r *core_api.DeleteNoticeResp, err error)
	ListCommunity(ctx context.Context, Req *core_api.ListCommunityReq, callOptions ...callopt.Option) (r *core_api.ListCommunityResp, err error)
	NewCommunity(ctx context.Context, Req *core_api.NewCommunityReq, callOptions ...callopt.Option) (r *core_api.NewCommunityResp, err error)
	DeleteCommunity(ctx context.Context, Req *core_api.DeleteCommunityReq, callOptions ...callopt.Option) (r *core_api.DeleteCommunityResp, err error)
	GetUserRoles(ctx context.Context, Req *core_api.GetUserRolesReq, callOptions ...callopt.Option) (r *core_api.GetUserRolesResp, err error)
	UpdateCommunityAdmin(ctx context.Context, Req *core_api.UpdateCommunityAdminReq, callOptions ...callopt.Option) (r *core_api.UpdateCommunityAdminResp, err error)
	UpdateSuperAdmin(ctx context.Context, Req *core_api.UpdateSuperAdminReq, callOptions ...callopt.Option) (r *core_api.UpdateSuperAdminResp, err error)
	GetUserByRole(ctx context.Context, Req *core_api.GetUserByRoleReq, callOptions ...callopt.Option) (r *core_api.GetUserByRoleResp, err error)
	UpdateRole(ctx context.Context, Req *core_api.UpdateRoleReq, callOptions ...callopt.Option) (r *core_api.UpdateRoleResp, err error)
	ListApply(ctx context.Context, Req *core_api.ListApplyReq, callOptions ...callopt.Option) (r *core_api.ListApplyResp, err error)
	HandleApply(ctx context.Context, Req *core_api.HandleApplyReq, callOptions ...callopt.Option) (r *core_api.HandleApplyResp, err error)
	CreateApply(ctx context.Context, Req *core_api.CreateApplyReq, callOptions ...callopt.Option) (r *core_api.CreateApplyResp, err error)
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
	return &kSystemClient{
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

type kSystemClient struct {
	*kClient
}

func (p *kSystemClient) GetAdmins(ctx context.Context, Req *core_api.GetAdminsReq, callOptions ...callopt.Option) (r *core_api.GetAdminsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetAdmins(ctx, Req)
}

func (p *kSystemClient) NewAdmin(ctx context.Context, Req *core_api.NewAdminReq, callOptions ...callopt.Option) (r *core_api.NewAdminResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.NewAdmin(ctx, Req)
}

func (p *kSystemClient) DeleteAdmin(ctx context.Context, Req *core_api.DeleteAdminReq, callOptions ...callopt.Option) (r *core_api.DeleteAdminResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteAdmin(ctx, Req)
}

func (p *kSystemClient) GetNews(ctx context.Context, Req *core_api.GetNewsReq, callOptions ...callopt.Option) (r *core_api.GetNewsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetNews(ctx, Req)
}

func (p *kSystemClient) NewNews(ctx context.Context, Req *core_api.NewNewsReq, callOptions ...callopt.Option) (r *core_api.NewNewsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.NewNews(ctx, Req)
}

func (p *kSystemClient) DeleteNews(ctx context.Context, Req *core_api.DeleteNewsReq, callOptions ...callopt.Option) (r *core_api.DeleteNewsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteNews(ctx, Req)
}

func (p *kSystemClient) GetNotices(ctx context.Context, Req *core_api.GetNoticesReq, callOptions ...callopt.Option) (r *core_api.GetNoticesResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetNotices(ctx, Req)
}

func (p *kSystemClient) NewNotice(ctx context.Context, Req *core_api.NewNoticeReq, callOptions ...callopt.Option) (r *core_api.NewNoticeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.NewNotice(ctx, Req)
}

func (p *kSystemClient) DeleteNotice(ctx context.Context, Req *core_api.DeleteNoticeReq, callOptions ...callopt.Option) (r *core_api.DeleteNoticeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteNotice(ctx, Req)
}

func (p *kSystemClient) ListCommunity(ctx context.Context, Req *core_api.ListCommunityReq, callOptions ...callopt.Option) (r *core_api.ListCommunityResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListCommunity(ctx, Req)
}

func (p *kSystemClient) NewCommunity(ctx context.Context, Req *core_api.NewCommunityReq, callOptions ...callopt.Option) (r *core_api.NewCommunityResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.NewCommunity(ctx, Req)
}

func (p *kSystemClient) DeleteCommunity(ctx context.Context, Req *core_api.DeleteCommunityReq, callOptions ...callopt.Option) (r *core_api.DeleteCommunityResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteCommunity(ctx, Req)
}

func (p *kSystemClient) GetUserRoles(ctx context.Context, Req *core_api.GetUserRolesReq, callOptions ...callopt.Option) (r *core_api.GetUserRolesResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserRoles(ctx, Req)
}

func (p *kSystemClient) UpdateCommunityAdmin(ctx context.Context, Req *core_api.UpdateCommunityAdminReq, callOptions ...callopt.Option) (r *core_api.UpdateCommunityAdminResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateCommunityAdmin(ctx, Req)
}

func (p *kSystemClient) UpdateSuperAdmin(ctx context.Context, Req *core_api.UpdateSuperAdminReq, callOptions ...callopt.Option) (r *core_api.UpdateSuperAdminResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateSuperAdmin(ctx, Req)
}

func (p *kSystemClient) GetUserByRole(ctx context.Context, Req *core_api.GetUserByRoleReq, callOptions ...callopt.Option) (r *core_api.GetUserByRoleResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserByRole(ctx, Req)
}

func (p *kSystemClient) UpdateRole(ctx context.Context, Req *core_api.UpdateRoleReq, callOptions ...callopt.Option) (r *core_api.UpdateRoleResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateRole(ctx, Req)
}

func (p *kSystemClient) ListApply(ctx context.Context, Req *core_api.ListApplyReq, callOptions ...callopt.Option) (r *core_api.ListApplyResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListApply(ctx, Req)
}

func (p *kSystemClient) HandleApply(ctx context.Context, Req *core_api.HandleApplyReq, callOptions ...callopt.Option) (r *core_api.HandleApplyResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.HandleApply(ctx, Req)
}

func (p *kSystemClient) CreateApply(ctx context.Context, Req *core_api.CreateApplyReq, callOptions ...callopt.Option) (r *core_api.CreateApplyResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateApply(ctx, Req)
}
