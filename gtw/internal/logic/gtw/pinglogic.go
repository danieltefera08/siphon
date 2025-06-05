package gtw

import (
	"context"
	"siphon/zerorpc/zerorpc"

	"siphon/gtw/internal/svc"
	"siphon/gtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping() (resp *types.PingReply, err error) {
	res, err := l.svcCtx.ZeroRpcCli.Ping(l.ctx, &zerorpc.Req{Ping: "gtw"})
	if err != nil {
		return nil, err
	}
	return &types.PingReply{Msg: res.Pong}, nil
}
