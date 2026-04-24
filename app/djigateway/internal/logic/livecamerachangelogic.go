package logic

import (
	"context"

	"siphon/app/djigateway/djigateway"
	"siphon/app/djigateway/internal/svc"
	"siphon/common/djisdk"

	"github.com/zeromicro/go-zero/core/logx"
)

type LiveCameraChangeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLiveCameraChangeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LiveCameraChangeLogic {
	return &LiveCameraChangeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LiveCameraChangeLogic) LiveCameraChange(in *djigateway.LiveCameraChangeReq) (*djigateway.CommonRes, error) {
	data := &djisdk.LiveCameraChangeData{
		VideoID:     in.VideoId,
		CameraIndex: in.CameraPosition,
	}
	tid, err := l.svcCtx.DjiClient.LiveCameraChange(l.ctx, in.DeviceSn, data)
	if err != nil {
		l.Errorf("[live] live camera change failed: %v", err)
		return errRes(tid, err), nil
	}
	return okRes(tid), nil
}
