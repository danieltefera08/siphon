package logic

import (
	"context"

	"siphon/app/djicloud/djicloud"
	"siphon/app/djicloud/internal/svc"
	"siphon/common/djisdk"

	"github.com/zeromicro/go-zero/core/logx"
)

type DrcCameraIsoSetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDrcCameraIsoSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DrcCameraIsoSetLogic {
	return &DrcCameraIsoSetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DrcCameraIsoSetLogic) DrcCameraIsoSet(in *djicloud.DrcCameraIsoSetReq) (*djicloud.DrcCameraIsoSetRes, error) {
	deviceSn := in.GetDeviceSn()
	seq, err := l.svcCtx.DjiClient.DrcNextSeq(deviceSn)
	if err != nil {
		return nil, drcError(err, "DRC 命令调用失败")
	}
	data := &djisdk.DrcCameraIsoSetData{PayloadIndex: in.GetPayloadIndex(), CameraType: in.GetCameraType(), ISOValue: int(in.GetIsoValue())}
	if _, err := l.svcCtx.DjiClient.DrcCameraIsoSet(l.ctx, deviceSn, seq, data); err != nil {
		return nil, drcError(err, "DRC 命令调用失败")
	}
	return &djicloud.DrcCameraIsoSetRes{Seq: int32(seq)}, nil
}
