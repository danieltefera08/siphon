package logic

import (
	"context"

	"siphon/app/live/internal/svc"
	"siphon/app/live/live"

	"github.com/zeromicro/go-zero/core/logx"
	"siphon/common/tool"
	"siphon/third_party/extproto"
)

type ListMeetingsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListMeetingsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMeetingsLogic {
	return &ListMeetingsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页查询会议列表
func (l *ListMeetingsLogic) ListMeetings(in *live.ListMeetingsReq) (*live.ListMeetingsRes, error) {
	page := in.Page
	if page <= 0 {
		page = 1
	}
	pageSize := in.PageSize
	if pageSize <= 0 {
		pageSize = 20
	}
	if pageSize > 100 {
		pageSize = 100
	}
	meetings, total, err := l.svcCtx.MeetingRepo.ListMeetings(l.ctx, in.Status, page, pageSize)
	if err != nil {
		return nil, tool.NewErrorByPbCodeWrap(extproto.Code__1_02_DB, err, "查询会议列表失败")
	}
	resp := &live.ListMeetingsRes{Total: total}
	for i := range meetings {
		resp.Meetings = append(resp.Meetings, toMeetingInfo(&meetings[i]))
	}
	return resp, nil
}
