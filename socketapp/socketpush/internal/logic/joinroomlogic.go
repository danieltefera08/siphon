package logic

import (
	"context"
	"time"
	"siphon/common/tool"
	"siphon/socketapp/socketgtw/socketgtw"

	"siphon/socketapp/socketpush/internal/svc"
	"siphon/socketapp/socketpush/socketpush"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
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
func (l *JoinRoomLogic) JoinRoom(in *socketpush.JoinRoomReq) (*socketpush.JoinRoomRes, error) {
	baseCtx := context.WithoutCancel(l.ctx)
	for _, cli := range l.svcCtx.SocketContainer.GetClients() {
		threading.GoSafe(func() {
			reqId, _ := tool.SimpleUUID()
			socktCTx, cancel := context.WithTimeout(baseCtx, 10*time.Second)
			defer cancel()
			cli.JoinRoom(socktCTx, &socketgtw.JoinRoomReq{
				ReqId: reqId,
				SId:   in.SId,
				Room:  in.Room,
			})
		})
	}
	return &socketpush.JoinRoomRes{}, nil
}
