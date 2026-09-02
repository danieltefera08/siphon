package live

import (
	"context"

	"siphon/app/live/live"
	"siphon/app/livegtw/internal/svc"
	"siphon/app/livegtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type KickParticipantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKickParticipantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KickParticipantLogic {
	return &KickParticipantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KickParticipantLogic) KickParticipant(req *types.KickParticipantRequest) error {
	_, err := l.svcCtx.LiveRpcCli.KickParticipant(l.ctx, &live.KickParticipantReq{
		MeetingNo: req.MeetingNo,
		Identity:  req.Identity,
	})
	return err
}
