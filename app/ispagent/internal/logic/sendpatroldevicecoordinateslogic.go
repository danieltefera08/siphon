package logic

import (
	"context"

	"siphon/app/ispagent/internal/svc"
	"siphon/app/ispagent/ispagent"
	"siphon/common/isp"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendPatrolDeviceCoordinatesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendPatrolDeviceCoordinatesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendPatrolDeviceCoordinatesLogic {
	return &SendPatrolDeviceCoordinatesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendPatrolDeviceCoordinatesLogic) SendPatrolDeviceCoordinates(in *ispagent.SendPatrolDeviceCoordinatesReq) (*ispagent.CommandRes, error) {
	msg, err := l.svcCtx.IspClient.Execute(l.ctx, isp.TypePatrolDeviceCoordinates, isp.CommandReport, in.GetCode(), patrolDeviceCoordinatesToItems(in.GetItems()))
	if err != nil {
		return nil, err
	}
	return commandResponse(msg), nil
}
