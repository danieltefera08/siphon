package logic

import (
	"context"

	"siphon/app/djigateway/djigateway"
	"siphon/app/djigateway/internal/svc"
	"siphon/common/djisdk"

	"github.com/zeromicro/go-zero/core/logx"
)

type CameraRecordingStopLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCameraRecordingStopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CameraRecordingStopLogic {
	return &CameraRecordingStopLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CameraRecordingStopLogic) CameraRecordingStop(in *djigateway.CameraRecordingStopReq) (*djigateway.CommonRes, error) {
	data := &djisdk.CameraRecordingStopData{
		PayloadIndex: in.PayloadIndex,
	}
	tid, err := l.svcCtx.DjiClient.CameraRecordingStop(l.ctx, in.DeviceSn, data)
	if err != nil {
		l.Errorf("[camera] camera recording stop failed: %v", err)
		return &djigateway.CommonRes{Code: -1, Message: err.Error(), Tid: tid}, nil
	}
	return &djigateway.CommonRes{Code: 0, Message: "success", Tid: tid}, nil
}
