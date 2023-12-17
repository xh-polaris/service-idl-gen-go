// Code generated by Kitex v0.8.0. DO NOT EDIT.

package contentservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	content "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/content"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	SearchCat(ctx context.Context, Req *content.SearchCatReq, callOptions ...callopt.Option) (r *content.SearchCatResp, err error)
	ListCat(ctx context.Context, Req *content.ListCatReq, callOptions ...callopt.Option) (r *content.ListCatResp, err error)
	RetrieveCat(ctx context.Context, Req *content.RetrieveCatReq, callOptions ...callopt.Option) (r *content.RetrieveCatResp, err error)
	CreateCat(ctx context.Context, Req *content.CreateCatReq, callOptions ...callopt.Option) (r *content.CreateCatResp, err error)
	UpdateCat(ctx context.Context, Req *content.UpdateCatReq, callOptions ...callopt.Option) (r *content.UpdateCatResp, err error)
	DeleteCat(ctx context.Context, Req *content.DeleteCatReq, callOptions ...callopt.Option) (r *content.DeleteCatResp, err error)
	CreateImage(ctx context.Context, Req *content.CreateImageReq, callOptions ...callopt.Option) (r *content.CreateImageResp, err error)
	DeleteImage(ctx context.Context, Req *content.DeleteImageReq, callOptions ...callopt.Option) (r *content.DeleteImageResp, err error)
	ListImage(ctx context.Context, Req *content.ListImageReq, callOptions ...callopt.Option) (r *content.ListImageResp, err error)
	ListMoment(ctx context.Context, Req *content.ListMomentReq, callOptions ...callopt.Option) (r *content.ListMomentResp, err error)
	CountMoment(ctx context.Context, Req *content.CountMomentReq, callOptions ...callopt.Option) (r *content.CountMomentResp, err error)
	RetrieveMoment(ctx context.Context, Req *content.RetrieveMomentReq, callOptions ...callopt.Option) (r *content.RetrieveMomentResp, err error)
	CreateMoment(ctx context.Context, Req *content.CreateMomentReq, callOptions ...callopt.Option) (r *content.CreateMomentResp, err error)
	UpdateMoment(ctx context.Context, Req *content.UpdateMomentReq, callOptions ...callopt.Option) (r *content.UpdateMomentResp, err error)
	DeleteMoment(ctx context.Context, Req *content.DeleteMomentReq, callOptions ...callopt.Option) (r *content.DeleteMomentResp, err error)
	CreatePost(ctx context.Context, Req *content.CreatePostReq, callOptions ...callopt.Option) (r *content.CreatePostResp, err error)
	RetrievePost(ctx context.Context, Req *content.RetrievePostReq, callOptions ...callopt.Option) (r *content.RetrievePostResp, err error)
	UpdatePost(ctx context.Context, Req *content.UpdatePostReq, callOptions ...callopt.Option) (r *content.UpdatePostResp, err error)
	DeletePost(ctx context.Context, Req *content.DeletePostReq, callOptions ...callopt.Option) (r *content.DeletePostResp, err error)
	ListPost(ctx context.Context, Req *content.ListPostReq, callOptions ...callopt.Option) (r *content.ListPostResp, err error)
	CountPost(ctx context.Context, Req *content.CountPostReq, callOptions ...callopt.Option) (r *content.CountPostResp, err error)
	SetOfficial(ctx context.Context, Req *content.SetOfficialReq, callOptions ...callopt.Option) (r *content.SetOfficialResp, err error)
	ListPlan(ctx context.Context, Req *content.ListPlanReq, callOptions ...callopt.Option) (r *content.ListPlanResp, err error)
	CountPlan(ctx context.Context, Req *content.CountPlanReq, callOptions ...callopt.Option) (r *content.CountPlanResp, err error)
	RetrievePlan(ctx context.Context, Req *content.RetrievePlanReq, callOptions ...callopt.Option) (r *content.RetrievePlanResp, err error)
	CreatePlan(ctx context.Context, Req *content.CreatePlanReq, callOptions ...callopt.Option) (r *content.CreatePlanResp, err error)
	UpdatePlan(ctx context.Context, Req *content.UpdatePlanReq, callOptions ...callopt.Option) (r *content.UpdatePlanResp, err error)
	DeletePlan(ctx context.Context, Req *content.DeletePlanReq, callOptions ...callopt.Option) (r *content.DeletePlanResp, err error)
	DonateFish(ctx context.Context, Req *content.DonateFishReq, callOptions ...callopt.Option) (r *content.DonateFishResp, err error)
	AddUserFish(ctx context.Context, Req *content.AddUserFishReq, callOptions ...callopt.Option) (r *content.AddUserFishResp, err error)
	ListDonateByUser(ctx context.Context, Req *content.ListDonateByUserReq, callOptions ...callopt.Option) (r *content.ListDonateByUserResp, err error)
	ListFishByPlan(ctx context.Context, Req *content.ListFishByPlanReq, callOptions ...callopt.Option) (r *content.ListFishByPlanResp, err error)
	RetrieveUserFish(ctx context.Context, Req *content.RetrieveUserFishReq, callOptions ...callopt.Option) (r *content.RetrieveUserFishResp, err error)
	CountDonateByUser(ctx context.Context, Req *content.CountDonateByUserReq, callOptions ...callopt.Option) (r *content.CountDonateByUserResp, err error)
	CountDonateByPlan(ctx context.Context, Req *content.CountDonateByPlanReq, callOptions ...callopt.Option) (r *content.CountDonateByPlanResp, err error)
	GetContentMission(ctx context.Context, Req *content.GetContentMissionReq, callOptions ...callopt.Option) (r *content.GetContentMissionResp, err error)
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

func (p *kContentServiceClient) SearchCat(ctx context.Context, Req *content.SearchCatReq, callOptions ...callopt.Option) (r *content.SearchCatResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SearchCat(ctx, Req)
}

func (p *kContentServiceClient) ListCat(ctx context.Context, Req *content.ListCatReq, callOptions ...callopt.Option) (r *content.ListCatResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListCat(ctx, Req)
}

func (p *kContentServiceClient) RetrieveCat(ctx context.Context, Req *content.RetrieveCatReq, callOptions ...callopt.Option) (r *content.RetrieveCatResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RetrieveCat(ctx, Req)
}

func (p *kContentServiceClient) CreateCat(ctx context.Context, Req *content.CreateCatReq, callOptions ...callopt.Option) (r *content.CreateCatResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateCat(ctx, Req)
}

func (p *kContentServiceClient) UpdateCat(ctx context.Context, Req *content.UpdateCatReq, callOptions ...callopt.Option) (r *content.UpdateCatResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateCat(ctx, Req)
}

func (p *kContentServiceClient) DeleteCat(ctx context.Context, Req *content.DeleteCatReq, callOptions ...callopt.Option) (r *content.DeleteCatResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteCat(ctx, Req)
}

func (p *kContentServiceClient) CreateImage(ctx context.Context, Req *content.CreateImageReq, callOptions ...callopt.Option) (r *content.CreateImageResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateImage(ctx, Req)
}

func (p *kContentServiceClient) DeleteImage(ctx context.Context, Req *content.DeleteImageReq, callOptions ...callopt.Option) (r *content.DeleteImageResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteImage(ctx, Req)
}

func (p *kContentServiceClient) ListImage(ctx context.Context, Req *content.ListImageReq, callOptions ...callopt.Option) (r *content.ListImageResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListImage(ctx, Req)
}

func (p *kContentServiceClient) ListMoment(ctx context.Context, Req *content.ListMomentReq, callOptions ...callopt.Option) (r *content.ListMomentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListMoment(ctx, Req)
}

func (p *kContentServiceClient) CountMoment(ctx context.Context, Req *content.CountMomentReq, callOptions ...callopt.Option) (r *content.CountMomentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CountMoment(ctx, Req)
}

func (p *kContentServiceClient) RetrieveMoment(ctx context.Context, Req *content.RetrieveMomentReq, callOptions ...callopt.Option) (r *content.RetrieveMomentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RetrieveMoment(ctx, Req)
}

func (p *kContentServiceClient) CreateMoment(ctx context.Context, Req *content.CreateMomentReq, callOptions ...callopt.Option) (r *content.CreateMomentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateMoment(ctx, Req)
}

func (p *kContentServiceClient) UpdateMoment(ctx context.Context, Req *content.UpdateMomentReq, callOptions ...callopt.Option) (r *content.UpdateMomentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateMoment(ctx, Req)
}

func (p *kContentServiceClient) DeleteMoment(ctx context.Context, Req *content.DeleteMomentReq, callOptions ...callopt.Option) (r *content.DeleteMomentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteMoment(ctx, Req)
}

func (p *kContentServiceClient) CreatePost(ctx context.Context, Req *content.CreatePostReq, callOptions ...callopt.Option) (r *content.CreatePostResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreatePost(ctx, Req)
}

func (p *kContentServiceClient) RetrievePost(ctx context.Context, Req *content.RetrievePostReq, callOptions ...callopt.Option) (r *content.RetrievePostResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RetrievePost(ctx, Req)
}

func (p *kContentServiceClient) UpdatePost(ctx context.Context, Req *content.UpdatePostReq, callOptions ...callopt.Option) (r *content.UpdatePostResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdatePost(ctx, Req)
}

func (p *kContentServiceClient) DeletePost(ctx context.Context, Req *content.DeletePostReq, callOptions ...callopt.Option) (r *content.DeletePostResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeletePost(ctx, Req)
}

func (p *kContentServiceClient) ListPost(ctx context.Context, Req *content.ListPostReq, callOptions ...callopt.Option) (r *content.ListPostResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListPost(ctx, Req)
}

func (p *kContentServiceClient) CountPost(ctx context.Context, Req *content.CountPostReq, callOptions ...callopt.Option) (r *content.CountPostResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CountPost(ctx, Req)
}

func (p *kContentServiceClient) SetOfficial(ctx context.Context, Req *content.SetOfficialReq, callOptions ...callopt.Option) (r *content.SetOfficialResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SetOfficial(ctx, Req)
}

func (p *kContentServiceClient) ListPlan(ctx context.Context, Req *content.ListPlanReq, callOptions ...callopt.Option) (r *content.ListPlanResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListPlan(ctx, Req)
}

func (p *kContentServiceClient) CountPlan(ctx context.Context, Req *content.CountPlanReq, callOptions ...callopt.Option) (r *content.CountPlanResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CountPlan(ctx, Req)
}

func (p *kContentServiceClient) RetrievePlan(ctx context.Context, Req *content.RetrievePlanReq, callOptions ...callopt.Option) (r *content.RetrievePlanResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RetrievePlan(ctx, Req)
}

func (p *kContentServiceClient) CreatePlan(ctx context.Context, Req *content.CreatePlanReq, callOptions ...callopt.Option) (r *content.CreatePlanResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreatePlan(ctx, Req)
}

func (p *kContentServiceClient) UpdatePlan(ctx context.Context, Req *content.UpdatePlanReq, callOptions ...callopt.Option) (r *content.UpdatePlanResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdatePlan(ctx, Req)
}

func (p *kContentServiceClient) DeletePlan(ctx context.Context, Req *content.DeletePlanReq, callOptions ...callopt.Option) (r *content.DeletePlanResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeletePlan(ctx, Req)
}

func (p *kContentServiceClient) DonateFish(ctx context.Context, Req *content.DonateFishReq, callOptions ...callopt.Option) (r *content.DonateFishResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DonateFish(ctx, Req)
}

func (p *kContentServiceClient) AddUserFish(ctx context.Context, Req *content.AddUserFishReq, callOptions ...callopt.Option) (r *content.AddUserFishResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddUserFish(ctx, Req)
}

func (p *kContentServiceClient) ListDonateByUser(ctx context.Context, Req *content.ListDonateByUserReq, callOptions ...callopt.Option) (r *content.ListDonateByUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListDonateByUser(ctx, Req)
}

func (p *kContentServiceClient) ListFishByPlan(ctx context.Context, Req *content.ListFishByPlanReq, callOptions ...callopt.Option) (r *content.ListFishByPlanResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListFishByPlan(ctx, Req)
}

func (p *kContentServiceClient) RetrieveUserFish(ctx context.Context, Req *content.RetrieveUserFishReq, callOptions ...callopt.Option) (r *content.RetrieveUserFishResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RetrieveUserFish(ctx, Req)
}

func (p *kContentServiceClient) CountDonateByUser(ctx context.Context, Req *content.CountDonateByUserReq, callOptions ...callopt.Option) (r *content.CountDonateByUserResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CountDonateByUser(ctx, Req)
}

func (p *kContentServiceClient) CountDonateByPlan(ctx context.Context, Req *content.CountDonateByPlanReq, callOptions ...callopt.Option) (r *content.CountDonateByPlanResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CountDonateByPlan(ctx, Req)
}

func (p *kContentServiceClient) GetContentMission(ctx context.Context, Req *content.GetContentMissionReq, callOptions ...callopt.Option) (r *content.GetContentMissionResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetContentMission(ctx, Req)
}
