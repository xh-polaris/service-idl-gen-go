// Code generated by Kitex v0.12.1. DO NOT EDIT.

package systemrpc

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	system "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/system"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	RetrieveNotice(ctx context.Context, Req *system.RetrieveNoticeReq, callOptions ...callopt.Option) (r *system.RetrieveNoticeResp, err error)
	ListNotice(ctx context.Context, Req *system.ListNoticeReq, callOptions ...callopt.Option) (r *system.ListNoticeResp, err error)
	CreateNotice(ctx context.Context, Req *system.CreateNoticeReq, callOptions ...callopt.Option) (r *system.CreateNoticeResp, err error)
	UpdateNotice(ctx context.Context, Req *system.UpdateNoticeReq, callOptions ...callopt.Option) (r *system.UpdateNoticeResp, err error)
	DeleteNotice(ctx context.Context, Req *system.DeleteNoticeReq, callOptions ...callopt.Option) (r *system.DeleteNoticeResp, err error)
	RetrieveNews(ctx context.Context, Req *system.RetrieveNewsReq, callOptions ...callopt.Option) (r *system.RetrieveNewsResp, err error)
	ListNews(ctx context.Context, Req *system.ListNewsReq, callOptions ...callopt.Option) (r *system.ListNewsResp, err error)
	CreateNews(ctx context.Context, Req *system.CreateNewsReq, callOptions ...callopt.Option) (r *system.CreateNewsResp, err error)
	UpdateNews(ctx context.Context, Req *system.UpdateNewsReq, callOptions ...callopt.Option) (r *system.UpdateNewsResp, err error)
	DeleteNews(ctx context.Context, Req *system.DeleteNewsReq, callOptions ...callopt.Option) (r *system.DeleteNewsResp, err error)
	RetrieveAdmin(ctx context.Context, Req *system.RetrieveAdminReq, callOptions ...callopt.Option) (r *system.RetrieveAdminResp, err error)
	ListAdmin(ctx context.Context, Req *system.ListAdminReq, callOptions ...callopt.Option) (r *system.ListAdminResp, err error)
	CreateAdmin(ctx context.Context, Req *system.CreateAdminReq, callOptions ...callopt.Option) (r *system.CreateAdminResp, err error)
	UpdateAdmin(ctx context.Context, Req *system.UpdateAdminReq, callOptions ...callopt.Option) (r *system.UpdateAdminResp, err error)
	DeleteAdmin(ctx context.Context, Req *system.DeleteAdminReq, callOptions ...callopt.Option) (r *system.DeleteAdminResp, err error)
	RetrieveUserRole(ctx context.Context, Req *system.RetrieveUserRoleReq, callOptions ...callopt.Option) (r *system.RetrieveUserRoleResp, err error)
	ListUserIdByRole(ctx context.Context, Req *system.ListUserIdByRoleReq, callOptions ...callopt.Option) (r *system.ListUserIdByRoleResp, err error)
	UpdateUserRole(ctx context.Context, Req *system.UpdateUserRoleReq, callOptions ...callopt.Option) (r *system.UpdateUserRoleResp, err error)
	ContainsRole(ctx context.Context, Req *system.ContainsRoleReq, callOptions ...callopt.Option) (r *system.ContainsRoleResp, err error)
	CreateApply(ctx context.Context, Req *system.CreateApplyReq, callOptions ...callopt.Option) (r *system.CreateApplyResp, err error)
	HandleApply(ctx context.Context, Req *system.HandleApplyReq, callOptions ...callopt.Option) (r *system.HandleApplyResp, err error)
	ListApply(ctx context.Context, Req *system.ListApplyReq, callOptions ...callopt.Option) (r *system.ListApplyResp, err error)
	RetrieveCommunity(ctx context.Context, Req *system.RetrieveCommunityReq, callOptions ...callopt.Option) (r *system.RetrieveCommunityResp, err error)
	ListCommunity(ctx context.Context, Req *system.ListCommunityReq, callOptions ...callopt.Option) (r *system.ListCommunityResp, err error)
	CreateCommunity(ctx context.Context, Req *system.CreateCommunityReq, callOptions ...callopt.Option) (r *system.CreateCommunityResp, err error)
	UpdateCommunity(ctx context.Context, Req *system.UpdateCommunityReq, callOptions ...callopt.Option) (r *system.UpdateCommunityResp, err error)
	DeleteCommunity(ctx context.Context, Req *system.DeleteCommunityReq, callOptions ...callopt.Option) (r *system.DeleteCommunityResp, err error)
	ListNotification(ctx context.Context, Req *system.ListNotificationReq, callOptions ...callopt.Option) (r *system.ListNotificationResp, err error)
	CleanNotification(ctx context.Context, Req *system.CleanNotificationReq, callOptions ...callopt.Option) (r *system.CleanNotificationResp, err error)
	CountNotification(ctx context.Context, Req *system.CountNotificationReq, callOptions ...callopt.Option) (r *system.CountNotificationResp, err error)
	ReadNotification(ctx context.Context, Req *system.ReadNotificationReq, callOptions ...callopt.Option) (r *system.ReadNotificationResp, err error)
	AddNotification(ctx context.Context, Req *system.AddNotificationReq, callOptions ...callopt.Option) (r *system.AddNotificationResp, err error)
	ReadRangeNotification(ctx context.Context, Req *system.ReadRangeNotificationReq, callOptions ...callopt.Option) (r *system.ReadRangeNotificationResp, err error)
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
	return &kSystemRpcClient{
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

type kSystemRpcClient struct {
	*kClient
}

func (p *kSystemRpcClient) RetrieveNotice(ctx context.Context, Req *system.RetrieveNoticeReq, callOptions ...callopt.Option) (r *system.RetrieveNoticeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RetrieveNotice(ctx, Req)
}

func (p *kSystemRpcClient) ListNotice(ctx context.Context, Req *system.ListNoticeReq, callOptions ...callopt.Option) (r *system.ListNoticeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListNotice(ctx, Req)
}

func (p *kSystemRpcClient) CreateNotice(ctx context.Context, Req *system.CreateNoticeReq, callOptions ...callopt.Option) (r *system.CreateNoticeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateNotice(ctx, Req)
}

func (p *kSystemRpcClient) UpdateNotice(ctx context.Context, Req *system.UpdateNoticeReq, callOptions ...callopt.Option) (r *system.UpdateNoticeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateNotice(ctx, Req)
}

func (p *kSystemRpcClient) DeleteNotice(ctx context.Context, Req *system.DeleteNoticeReq, callOptions ...callopt.Option) (r *system.DeleteNoticeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteNotice(ctx, Req)
}

func (p *kSystemRpcClient) RetrieveNews(ctx context.Context, Req *system.RetrieveNewsReq, callOptions ...callopt.Option) (r *system.RetrieveNewsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RetrieveNews(ctx, Req)
}

func (p *kSystemRpcClient) ListNews(ctx context.Context, Req *system.ListNewsReq, callOptions ...callopt.Option) (r *system.ListNewsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListNews(ctx, Req)
}

func (p *kSystemRpcClient) CreateNews(ctx context.Context, Req *system.CreateNewsReq, callOptions ...callopt.Option) (r *system.CreateNewsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateNews(ctx, Req)
}

func (p *kSystemRpcClient) UpdateNews(ctx context.Context, Req *system.UpdateNewsReq, callOptions ...callopt.Option) (r *system.UpdateNewsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateNews(ctx, Req)
}

func (p *kSystemRpcClient) DeleteNews(ctx context.Context, Req *system.DeleteNewsReq, callOptions ...callopt.Option) (r *system.DeleteNewsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteNews(ctx, Req)
}

func (p *kSystemRpcClient) RetrieveAdmin(ctx context.Context, Req *system.RetrieveAdminReq, callOptions ...callopt.Option) (r *system.RetrieveAdminResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RetrieveAdmin(ctx, Req)
}

func (p *kSystemRpcClient) ListAdmin(ctx context.Context, Req *system.ListAdminReq, callOptions ...callopt.Option) (r *system.ListAdminResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListAdmin(ctx, Req)
}

func (p *kSystemRpcClient) CreateAdmin(ctx context.Context, Req *system.CreateAdminReq, callOptions ...callopt.Option) (r *system.CreateAdminResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateAdmin(ctx, Req)
}

func (p *kSystemRpcClient) UpdateAdmin(ctx context.Context, Req *system.UpdateAdminReq, callOptions ...callopt.Option) (r *system.UpdateAdminResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateAdmin(ctx, Req)
}

func (p *kSystemRpcClient) DeleteAdmin(ctx context.Context, Req *system.DeleteAdminReq, callOptions ...callopt.Option) (r *system.DeleteAdminResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteAdmin(ctx, Req)
}

func (p *kSystemRpcClient) RetrieveUserRole(ctx context.Context, Req *system.RetrieveUserRoleReq, callOptions ...callopt.Option) (r *system.RetrieveUserRoleResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RetrieveUserRole(ctx, Req)
}

func (p *kSystemRpcClient) ListUserIdByRole(ctx context.Context, Req *system.ListUserIdByRoleReq, callOptions ...callopt.Option) (r *system.ListUserIdByRoleResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListUserIdByRole(ctx, Req)
}

func (p *kSystemRpcClient) UpdateUserRole(ctx context.Context, Req *system.UpdateUserRoleReq, callOptions ...callopt.Option) (r *system.UpdateUserRoleResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateUserRole(ctx, Req)
}

func (p *kSystemRpcClient) ContainsRole(ctx context.Context, Req *system.ContainsRoleReq, callOptions ...callopt.Option) (r *system.ContainsRoleResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ContainsRole(ctx, Req)
}

func (p *kSystemRpcClient) CreateApply(ctx context.Context, Req *system.CreateApplyReq, callOptions ...callopt.Option) (r *system.CreateApplyResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateApply(ctx, Req)
}

func (p *kSystemRpcClient) HandleApply(ctx context.Context, Req *system.HandleApplyReq, callOptions ...callopt.Option) (r *system.HandleApplyResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.HandleApply(ctx, Req)
}

func (p *kSystemRpcClient) ListApply(ctx context.Context, Req *system.ListApplyReq, callOptions ...callopt.Option) (r *system.ListApplyResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListApply(ctx, Req)
}

func (p *kSystemRpcClient) RetrieveCommunity(ctx context.Context, Req *system.RetrieveCommunityReq, callOptions ...callopt.Option) (r *system.RetrieveCommunityResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RetrieveCommunity(ctx, Req)
}

func (p *kSystemRpcClient) ListCommunity(ctx context.Context, Req *system.ListCommunityReq, callOptions ...callopt.Option) (r *system.ListCommunityResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListCommunity(ctx, Req)
}

func (p *kSystemRpcClient) CreateCommunity(ctx context.Context, Req *system.CreateCommunityReq, callOptions ...callopt.Option) (r *system.CreateCommunityResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateCommunity(ctx, Req)
}

func (p *kSystemRpcClient) UpdateCommunity(ctx context.Context, Req *system.UpdateCommunityReq, callOptions ...callopt.Option) (r *system.UpdateCommunityResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateCommunity(ctx, Req)
}

func (p *kSystemRpcClient) DeleteCommunity(ctx context.Context, Req *system.DeleteCommunityReq, callOptions ...callopt.Option) (r *system.DeleteCommunityResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteCommunity(ctx, Req)
}

func (p *kSystemRpcClient) ListNotification(ctx context.Context, Req *system.ListNotificationReq, callOptions ...callopt.Option) (r *system.ListNotificationResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListNotification(ctx, Req)
}

func (p *kSystemRpcClient) CleanNotification(ctx context.Context, Req *system.CleanNotificationReq, callOptions ...callopt.Option) (r *system.CleanNotificationResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CleanNotification(ctx, Req)
}

func (p *kSystemRpcClient) CountNotification(ctx context.Context, Req *system.CountNotificationReq, callOptions ...callopt.Option) (r *system.CountNotificationResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CountNotification(ctx, Req)
}

func (p *kSystemRpcClient) ReadNotification(ctx context.Context, Req *system.ReadNotificationReq, callOptions ...callopt.Option) (r *system.ReadNotificationResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ReadNotification(ctx, Req)
}

func (p *kSystemRpcClient) AddNotification(ctx context.Context, Req *system.AddNotificationReq, callOptions ...callopt.Option) (r *system.AddNotificationResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddNotification(ctx, Req)
}

func (p *kSystemRpcClient) ReadRangeNotification(ctx context.Context, Req *system.ReadRangeNotificationReq, callOptions ...callopt.Option) (r *system.ReadRangeNotificationResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ReadRangeNotification(ctx, Req)
}
