package logic

import (
	"context"

	"siphon/app/bridgemodbus/bridgemodbus"
	"siphon/app/bridgemodbus/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *bridgemodbus.Req) (*bridgemodbus.Res, error) {
	// todo: add your logic here and delete this line

	return &bridgemodbus.Res{}, nil
}
