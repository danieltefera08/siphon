package logic

import (
	"context"

	"siphon/app/djigateway/djigateway"
	"siphon/app/djigateway/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlightTaskExecuteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlightTaskExecuteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlightTaskExecuteLogic {
	return &FlightTaskExecuteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlightTaskExecuteLogic) FlightTaskExecute(in *djigateway.FlightTaskExecuteReq) (*djigateway.CommonRes, error) {
	tid, err := l.svcCtx.DjiClient.FlightTaskExecute(l.ctx, in.DeviceSn, in.FlightId)
	if err != nil {
		l.Errorf("[flight-task] flighttask_execute failed: %v", err)
		return errRes(tid, err), nil
	}
	return okRes(tid), nil
}
