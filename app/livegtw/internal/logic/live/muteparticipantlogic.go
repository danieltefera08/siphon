package live

import (
	"context"

	"siphon/app/live/live"
	"siphon/app/livegtw/internal/svc"
	"siphon/app/livegtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MuteParticipantLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMuteParticipantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MuteParticipantLogic {
	return &MuteParticipantLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MuteParticipantLogic) MuteParticipant(req *types.MuteParticipantRequest) error {
	_, err := l.svcCtx.LiveRpcCli.MuteParticipant(l.ctx, &live.MuteParticipantReq{
		MeetingNo: req.MeetingNo,
		Identity:  req.Identity,
		Muted:     req.Muted,
		Kind:      req.Kind,
	})
	return err
}
