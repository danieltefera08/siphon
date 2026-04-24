package logic

import (
	"context"

	"siphon/app/djigateway/djigateway"
	"siphon/app/djigateway/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DrcModeExitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDrcModeExitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DrcModeExitLogic {
	return &DrcModeExitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DrcModeExitLogic) DrcModeExit(in *djigateway.DeviceSnReq) (*djigateway.CommonRes, error) {
	tid, err := l.svcCtx.DjiClient.DrcModeExit(l.ctx, in.DeviceSn)
	if err != nil {
		l.Errorf("[drc] drc mode exit failed: %v", err)
		return errRes(tid, err), nil
	}
	return okRes(tid), nil
}
