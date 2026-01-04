package logic

import (
	"context"
	"siphon/socketapp/socketgtw/socketgtw"

	"siphon/socketapp/socketpush/internal/svc"
	"siphon/socketapp/socketpush/socketpush"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
)

type KickMetaSessionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewKickMetaSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KickMetaSessionLogic {
	return &KickMetaSessionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 指定元数据剔除 session
func (l *KickMetaSessionLogic) KickMetaSession(in *socketpush.KickMetaSessionReq) (*socketpush.KickMetaSessionRes, error) {
	baseCtx := context.WithoutCancel(l.ctx)
	for _, cli := range l.svcCtx.SocketContainer.GetClients() {
		threading.GoSafe(func() {
			cli.KickMetaSession(baseCtx, &socketgtw.KickMetaSessionReq{
				ReqId: in.ReqId,
				Key:   in.Key,
				Value: in.Value,
			})
		})
	}
	return &socketpush.KickMetaSessionRes{
		ReqId: in.ReqId,
	}, nil
}
