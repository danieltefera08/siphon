package logic

import (
	"context"

	"siphon/app/djicloud/djicloud"
	"siphon/app/djicloud/internal/svc"
	"siphon/common/djisdk"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlightTaskRecoveryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlightTaskRecoveryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlightTaskRecoveryLogic {
	return &FlightTaskRecoveryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlightTaskRecoveryLogic) FlightTaskRecovery(in *djicloud.FlightTaskRecoveryReq) (*djicloud.CommonRes, error) {
	tid, err := l.svcCtx.DjiClient.FlightTaskRecovery(l.ctx, in.GetDeviceSn(), &djisdk.FlightTaskRecoveryData{
		FlightID:  in.GetFlightId(),
		WaylineID: int(in.GetWaylineId()),
	})
	if err != nil {
		return errRes(tid, err), nil
	}
	return okRes(tid), nil
}
