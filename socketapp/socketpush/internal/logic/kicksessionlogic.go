package logic

import (
	"context"
	"siphon/socketapp/socketgtw/socketgtw"

	"siphon/socketapp/socketpush/internal/svc"
	"siphon/socketapp/socketpush/socketpush"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
)

type KickSessionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewKickSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KickSessionLogic {
	return &KickSessionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 剔除 session
func (l *KickSessionLogic) KickSession(in *socketpush.KickSessionReq) (*socketpush.KickSessionRes, error) {
	baseCtx := context.WithoutCancel(l.ctx)
	for _, cli := range l.svcCtx.SocketContainer.GetClients() {
		threading.GoSafe(func() {
			cli.KickSession(baseCtx, &socketgtw.KickSessionReq{
				ReqId:    in.ReqId,
				SocketId: in.SocketId,
			})
		})
	}
	return &socketpush.KickSessionRes{
		ReqId: in.ReqId,
	}, nil
}
