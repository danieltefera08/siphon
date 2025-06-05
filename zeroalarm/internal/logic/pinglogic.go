package logic

import (
	"context"

	"siphon/zeroalarm/internal/svc"
	"siphon/zeroalarm/zeroalarm"

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

func (l *PingLogic) Ping(in *zeroalarm.Req) (*zeroalarm.Res, error) {
	return &zeroalarm.Res{Pong: "alarm"}, nil
}
