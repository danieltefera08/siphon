package logic

import (
	"context"

	"siphon/app/djicloud/djicloud"
	"siphon/app/djicloud/internal/svc"
	"siphon/common/djisdk"

	"github.com/zeromicro/go-zero/core/logx"
)

type DrcNightLightsStateSetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDrcNightLightsStateSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DrcNightLightsStateSetLogic {
	return &DrcNightLightsStateSetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DrcNightLightsStateSetLogic) DrcNightLightsStateSet(in *djicloud.DrcNightLightsStateSetReq) (*djicloud.CommonRes, error) {
	data := &djisdk.DrcNightLightsStateSetData{NightLightsState: int(in.GetNightLightsState())}
	tid, err := l.svcCtx.DjiClient.DrcNightLightsStateSet(l.ctx, in.GetDeviceSn(), data)
	if err != nil {
		return errRes(tid, err), nil
	}
	return okRes(tid), nil
}
