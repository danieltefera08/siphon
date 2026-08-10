package logic

import (
	"context"

	"siphon/app/djicloud/djicloud"
	"siphon/app/djicloud/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SupplementLightOpenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSupplementLightOpenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SupplementLightOpenLogic {
	return &SupplementLightOpenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SupplementLightOpen 开启机巢补光灯。
func (l *SupplementLightOpenLogic) SupplementLightOpen(in *djicloud.SupplementLightOpenReq) (*djicloud.CommonRes, error) {
	tid, err := l.svcCtx.DjiClient.SupplementLightOpen(l.ctx, in.DeviceSn)
	if err != nil {
		return commandRes(tid, err)
	}
	return okRes(tid), nil
}
