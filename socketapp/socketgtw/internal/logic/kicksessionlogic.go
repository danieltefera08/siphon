package logic

import (
	"context"

	"siphon/socketapp/socketgtw/internal/svc"
	"siphon/socketapp/socketgtw/socketgtw"

	"github.com/zeromicro/go-zero/core/logx"
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
func (l *KickSessionLogic) KickSession(in *socketgtw.KickSessionReq) (*socketgtw.KickSessionRes, error) {
	sess := l.svcCtx.SocketServer.GetSession(in.SId)
	if sess != nil {
		err := sess.Close()
		return nil, err
	}
	return &socketgtw.KickSessionRes{}, nil
}
