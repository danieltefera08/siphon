package live

import (
	"context"

	"siphon/app/live/live"
	"siphon/app/livegtw/internal/svc"
	"siphon/app/livegtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListSipProvidersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListSipProvidersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListSipProvidersLogic {
	return &ListSipProvidersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListSipProvidersLogic) ListSipProviders(req *types.ListSipProvidersRequest) (resp *types.ListSipProvidersReply, err error) {
	r, err := l.svcCtx.LiveRpcCli.ListSipProviders(l.ctx, &live.ListSipProvidersReq{})
	if err != nil {
		return nil, err
	}
	providers := make([]types.SipProviderInfo, 0, len(r.GetProviders()))
	for _, p := range r.GetProviders() {
		providers = append(providers, toSipProviderInfo(p))
	}
	return &types.ListSipProvidersReply{
		Providers: providers,
	}, nil
}
