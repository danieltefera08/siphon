package logic

import (
	"context"

	"siphon/app/djicloud/djicloud"
	"siphon/app/djicloud/internal/svc"
	"siphon/common/djisdk"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlyToPointStopLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlyToPointStopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlyToPointStopLogic {
	return &FlyToPointStopLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlyToPointStopLogic) FlyToPointStop(in *djicloud.FlyToPointStopReq) (*djicloud.CommonRes, error) {
	tid, err := l.svcCtx.DjiClient.FlyToPointStop(l.ctx, in.GetDeviceSn(), &djisdk.FlyToPointStopData{
		FlyToID: in.GetFlyToId(),
	})
	if err != nil {
		return errRes(tid, err), nil
	}
	return okRes(tid), nil
}
