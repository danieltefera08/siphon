package logic

import (
	"context"

	"siphon/app/djigateway/djigateway"
	"siphon/app/djigateway/internal/hooks"
	"siphon/app/djigateway/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsDeviceOnlineLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsDeviceOnlineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsDeviceOnlineLogic {
	return &IsDeviceOnlineLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// IsDeviceOnline 查询机巢在线状态。
func (l *IsDeviceOnlineLogic) IsDeviceOnline(in *djigateway.DeviceSnReq) (*djigateway.DeviceOnlineRes, error) {
	online := hooks.IsOnline(l.svcCtx.OnlineCache, in.DeviceSn)
	return &djigateway.DeviceOnlineRes{IsOnline: online}, nil
}
