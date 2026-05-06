package logic

import (
	"context"

	"siphon/app/djicloud/djicloud"
	"siphon/app/djicloud/internal/svc"
	"siphon/common/djisdk"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResumeFlightTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResumeFlightTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResumeFlightTaskLogic {
	return &ResumeFlightTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ResumeFlightTask 恢复已暂停的飞行任务。
func (l *ResumeFlightTaskLogic) ResumeFlightTask(in *djicloud.ResumeFlightTaskReq) (*djicloud.CommonRes, error) {
	tid, err := l.svcCtx.DjiClient.ResumeFlightTask(l.ctx, in.GetDeviceSn(), &djisdk.FlightTaskResumeData{
		FlightID:  in.GetFlightId(),
		WaylineID: int(in.GetWaylineId()),
	})
	if err != nil {
		l.Errorf("[flight-task] resume flight task failed: %v", err)
		return errRes(tid, err), nil
	}
	return okRes(tid), nil
}
