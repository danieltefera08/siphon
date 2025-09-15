package logic

import (
	"context"

	"siphon/app/bridgedump/bridgedump"
	"siphon/app/bridgedump/internal/svc"

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

func (l *PingLogic) Ping(in *bridgedump.Req) (*bridgedump.Res, error) {
	return &bridgedump.Res{Pong: "pong"}, nil
}
