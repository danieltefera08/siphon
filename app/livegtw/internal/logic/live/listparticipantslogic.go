package live

import (
	"context"

	"siphon/app/live/live"
	"siphon/app/livegtw/internal/svc"
	"siphon/app/livegtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListParticipantsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListParticipantsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListParticipantsLogic {
	return &ListParticipantsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListParticipantsLogic) ListParticipants(req *types.ListParticipantsRequest) (resp *types.ListParticipantsReply, err error) {
	r, err := l.svcCtx.LiveRpcCli.ListParticipants(l.ctx, &live.ListParticipantsReq{
		MeetingNo: req.MeetingNo,
	})
	if err != nil {
		return nil, err
	}
	participants := make([]types.ParticipantInfo, 0, len(r.GetParticipants()))
	for _, p := range r.GetParticipants() {
		participants = append(participants, toParticipantInfo(p))
	}
	return &types.ListParticipantsReply{
		Participants: participants,
	}, nil
}
