package logic

import (
	"context"

	"siphon/app/djicloud/djicloud"
	"siphon/app/djicloud/internal/svc"
	"siphon/common/djisdk"

	"github.com/zeromicro/go-zero/core/logx"
)

type DrcCameraShutterSetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDrcCameraShutterSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DrcCameraShutterSetLogic {
	return &DrcCameraShutterSetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DrcCameraShutterSetLogic) DrcCameraShutterSet(in *djicloud.DrcCameraShutterSetReq) (*djicloud.CommonRes, error) {
	data := &djisdk.DrcCameraShutterSetData{PayloadIndex: in.GetPayloadIndex(), CameraType: in.GetCameraType(), ShutterValue: int(in.GetShutterValue())}
	tid, err := l.svcCtx.DjiClient.DrcCameraShutterSet(l.ctx, in.GetDeviceSn(), data)
	if err != nil {
		return errRes(tid, err), nil
	}
	return okRes(tid), nil
}
