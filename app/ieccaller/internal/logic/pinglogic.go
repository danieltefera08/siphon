package logic

import (
	"context"

	"siphon/app/ieccaller/ieccaller"
	"siphon/app/ieccaller/internal/svc"

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

func (l *PingLogic) Ping(in *ieccaller.Req) (*ieccaller.Res, error) {
	return &ieccaller.Res{Pong: "pong"}, nil
}
