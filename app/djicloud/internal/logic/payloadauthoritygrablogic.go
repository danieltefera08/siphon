package logic

import (
	"context"

	"siphon/app/djicloud/djicloud"
	"siphon/app/djicloud/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PayloadAuthorityGrabLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPayloadAuthorityGrabLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayloadAuthorityGrabLogic {
	return &PayloadAuthorityGrabLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PayloadAuthorityGrabLogic) PayloadAuthorityGrab(in *djicloud.PayloadAuthorityGrabReq) (*djicloud.CommonRes, error) {
	tid, err := l.svcCtx.DjiClient.PayloadAuthorityGrab(l.ctx, in.DeviceSn)
	if err != nil {
		return commandRes(tid, err)
	}
	return okRes(tid), nil
}
