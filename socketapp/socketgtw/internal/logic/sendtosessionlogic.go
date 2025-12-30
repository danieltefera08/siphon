package logic

import (
	"context"

	"siphon/socketapp/socketgtw/internal/svc"
	"siphon/socketapp/socketgtw/socketgtw"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendToSessionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendToSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendToSessionLogic {
	return &SendToSessionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 向指定 session 发送消息
func (l *SendToSessionLogic) SendToSession(in *socketgtw.SendToSessionReq) (*socketgtw.SendToSessionRes, error) {
	session := l.svcCtx.SocketServer.GetSession(in.SId)
	if session != nil {
		err := session.Emit(in.Event, in.Payload)
		if err != nil {
			return nil, err
		}
	}
	return &socketgtw.SendToSessionRes{}, nil
}
