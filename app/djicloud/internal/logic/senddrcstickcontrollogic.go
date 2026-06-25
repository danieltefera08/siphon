package logic

import (
	"context"

	"siphon/app/djicloud/djicloud"
	"siphon/app/djicloud/internal/svc"
	"siphon/common/djisdk"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendDrcStickControlLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendDrcStickControlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendDrcStickControlLogic {
	return &SendDrcStickControlLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendDrcStickControlLogic) SendDrcStickControl(in *djicloud.DrcStickControlReq) (*djicloud.DrcStickControlRes, error) {
	deviceSn := in.GetDeviceSn()
	seq, err := l.svcCtx.DjiClient.DrcNextSeq(deviceSn)
	if err != nil {
		return nil, err
	}
	data := buildDrcStickControlData(in)
	if _, err := l.svcCtx.DjiClient.SendDrcStickControl(l.ctx, deviceSn, int(seq), data); err != nil {
		l.Errorf("[drc] send stick control failed device_sn=%s seq=%d: %v", deviceSn, seq, err)
		return nil, err
	}
	return &djicloud.DrcStickControlRes{Seq: int32(seq)}, nil
}

func buildDrcStickControlData(in *djicloud.DrcStickControlReq) *djisdk.DrcStickControlData {
	return &djisdk.DrcStickControlData{
		Roll:        in.GetRoll(),
		Pitch:       in.GetPitch(),
		Throttle:    in.GetThrottle(),
		Yaw:         in.GetYaw(),
		GimbalPitch: in.GetGimbalPitch(),
	}
}
