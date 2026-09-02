package live

import (
	"context"

	"siphon/app/live/live"
	"siphon/app/livegtw/internal/svc"
	"siphon/app/livegtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EndMeetingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEndMeetingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EndMeetingLogic {
	return &EndMeetingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EndMeetingLogic) EndMeeting(req *types.EndMeetingRequest) error {
	_, err := l.svcCtx.LiveRpcCli.EndMeeting(l.ctx, &live.EndMeetingReq{
		MeetingNo: req.MeetingNo,
	})
	return err
}
