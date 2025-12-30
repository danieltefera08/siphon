package logic

import (
	"context"
	"siphon/socketapp/socketgtw/internal/svc"
	"siphon/socketapp/socketgtw/socketgtw"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinRoomLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJoinRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinRoomLogic {
	return &JoinRoomLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 加入房间
func (l *JoinRoomLogic) JoinRoom(in *socketgtw.JoinRoomReq) (*socketgtw.JoinRoomRes, error) {
	session := l.svcCtx.SocketServer.GetSession(in.SId)
	if session != nil {
		err := session.JoinRoom(in.Room)
		if err != nil {
			return nil, err
		}
	}
	return &socketgtw.JoinRoomRes{}, nil
}
