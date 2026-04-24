package logic

import (
	"context"

	"siphon/app/djigateway/djigateway"
	"siphon/app/djigateway/internal/svc"
	"siphon/common/djisdk"

	"github.com/zeromicro/go-zero/core/logx"
)

type CameraPointFocusActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCameraPointFocusActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CameraPointFocusActionLogic {
	return &CameraPointFocusActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CameraPointFocusActionLogic) CameraPointFocusAction(in *djigateway.CameraPointFocusActionReq) (*djigateway.CommonRes, error) {
	data := &djisdk.CameraPointFocusActionData{
		PayloadIndex: in.PayloadIndex,
		CameraType:   int(in.CameraType),
		X:            in.X,
		Y:            in.Y,
	}
	tid, err := l.svcCtx.DjiClient.CameraPointFocusAction(l.ctx, in.DeviceSn, data)
	if err != nil {
		l.Errorf("[camera] camera point focus action failed: %v", err)
		return errRes(tid, err), nil
	}
	return okRes(tid), nil
}
