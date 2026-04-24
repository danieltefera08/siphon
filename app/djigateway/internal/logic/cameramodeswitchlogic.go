package logic

import (
	"context"

	"siphon/app/djigateway/djigateway"
	"siphon/app/djigateway/internal/svc"
	"siphon/common/djisdk"

	"github.com/zeromicro/go-zero/core/logx"
)

type CameraModeSwitchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCameraModeSwitchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CameraModeSwitchLogic {
	return &CameraModeSwitchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CameraModeSwitch 切换相机拍摄模式。
func (l *CameraModeSwitchLogic) CameraModeSwitch(in *djigateway.CameraModeSwitchReq) (*djigateway.CommonRes, error) {
	data := &djisdk.CameraModeSwitchData{
		PayloadIndex: in.PayloadIndex,
		CameraMode:   int(in.CameraMode),
	}
	tid, err := l.svcCtx.DjiClient.CameraModeSwitch(l.ctx, in.DeviceSn, data)
	if err != nil {
		l.Errorf("[camera] camera mode switch failed: %v", err)
		return &djigateway.CommonRes{Code: -1, Message: err.Error(), Tid: tid}, nil
	}
	return &djigateway.CommonRes{Code: 0, Message: "success", Tid: tid}, nil
}
