package logic

import (
	"context"

	"siphon/app/djicloud/djicloud"
	"siphon/app/djicloud/internal/svc"
	"siphon/common/djisdk"

	"github.com/zeromicro/go-zero/core/logx"
)

type DrcCameraMechanicalShutterSetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDrcCameraMechanicalShutterSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DrcCameraMechanicalShutterSetLogic {
	return &DrcCameraMechanicalShutterSetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DrcCameraMechanicalShutterSetLogic) DrcCameraMechanicalShutterSet(in *djicloud.DrcCameraMechanicalShutterSetReq) (*djicloud.CommonRes, error) {
	data := &djisdk.DrcCameraMechanicalShutterSetData{PayloadIndex: in.GetPayloadIndex(), CameraType: in.GetCameraType(), MechanicalShutterState: int(in.GetDewarpingState())}
	tid, err := l.svcCtx.DjiClient.DrcCameraMechanicalShutterSet(l.ctx, in.GetDeviceSn(), data)
	if err != nil {
		return errRes(tid, err), nil
	}
	return okRes(tid), nil
}
