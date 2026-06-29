package logic

import (
	"context"

	"siphon/app/djicloud/djicloud"
	"siphon/app/djicloud/internal/svc"
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

func (l *CameraRecordingStopLogic) CameraRecordingStop(in *djicloud.CameraRecordingStopReq) (*djicloud.CommonRes, error) {
	data := &djisdk.CameraRecordingStopData{
		PayloadIndex: in.PayloadIndex,
	}
	tid, err := l.svcCtx.DjiClient.CameraRecordingStop(l.ctx, in.DeviceSn, data)
	if err != nil {
		return errRes(tid, err), nil
	}
	return okRes(tid), nil
}
