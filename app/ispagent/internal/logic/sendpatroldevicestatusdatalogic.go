package logic

import (
	"context"

	"siphon/app/ispagent/internal/svc"
	"siphon/app/ispagent/ispagent"
	"siphon/common/isp"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendPatrolDeviceStatusDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendPatrolDeviceStatusDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendPatrolDeviceStatusDataLogic {
	return &SendPatrolDeviceStatusDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendPatrolDeviceStatusDataLogic) SendPatrolDeviceStatusData(in *ispagent.SendPatrolDeviceStatusDataReq) (*ispagent.CommandRes, error) {
	msg, err := l.svcCtx.IspClient.Execute(l.ctx, isp.TypePatrolDeviceStatusData, isp.CommandReport, in.GetCode(), patrolDeviceStatusDataToItems(in.GetItems()))
	if err != nil {
		return nil, err
	}
	return commandResponse(msg), nil
}
