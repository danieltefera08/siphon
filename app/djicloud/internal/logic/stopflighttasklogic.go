package logic

import (
	"context"

	"siphon/app/djicloud/djicloud"
	"siphon/app/djicloud/internal/svc"
	"siphon/common/djisdk"

	"github.com/zeromicro/go-zero/core/logx"
)

type StopFlightTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStopFlightTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StopFlightTaskLogic {
	return &StopFlightTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// StopFlightTask 强制停止当前航线任务。
func (l *StopFlightTaskLogic) StopFlightTask(in *djicloud.StopFlightTaskReq) (*djicloud.CommonRes, error) {
	tid, err := l.svcCtx.DjiClient.StopFlightTask(l.ctx, in.GetDeviceSn(), &djisdk.FlightTaskStopData{
		FlightID:  in.GetFlightId(),
		WaylineID: int(in.GetWaylineId()),
	})
	if err != nil {
		return commandRes(tid, err)
	}
	return okRes(tid), nil
}
