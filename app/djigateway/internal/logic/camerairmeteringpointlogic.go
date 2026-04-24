package logic

import (
	"context"

	"siphon/app/djigateway/djigateway"
	"siphon/app/djigateway/internal/svc"
	"siphon/common/djisdk"

	"github.com/zeromicro/go-zero/core/logx"
)

type CameraIrMeteringPointLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCameraIrMeteringPointLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CameraIrMeteringPointLogic {
	return &CameraIrMeteringPointLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CameraIrMeteringPointLogic) CameraIrMeteringPoint(in *djigateway.CameraIrMeteringPointReq) (*djigateway.CommonRes, error) {
	data := &djisdk.CameraIrMeteringPointData{
		PayloadIndex: in.PayloadIndex,
		X:            in.X,
		Y:            in.Y,
	}
	tid, err := l.svcCtx.DjiClient.CameraIrMeteringPoint(l.ctx, in.DeviceSn, data)
	if err != nil {
		l.Errorf("[camera] camera ir metering point failed: %v", err)
		return errRes(tid, err), nil
	}
	return okRes(tid), nil
}
