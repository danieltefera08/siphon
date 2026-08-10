package logic

import (
	"context"

	"siphon/app/djicloud/djicloud"
	"siphon/app/djicloud/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnHomeCancelAutoReturnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReturnHomeCancelAutoReturnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnHomeCancelAutoReturnLogic {
	return &ReturnHomeCancelAutoReturnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReturnHomeCancelAutoReturnLogic) ReturnHomeCancelAutoReturn(in *djicloud.ReturnHomeCancelAutoReturnReq) (*djicloud.CommonRes, error) {
	tid, err := l.svcCtx.DjiClient.ReturnHomeCancelAutoReturn(l.ctx, in.DeviceSn)
	if err != nil {
		return commandRes(tid, err)
	}
	return okRes(tid), nil
}
