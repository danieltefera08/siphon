package logic

import (
	"context"

	"siphon/app/djicloud/djicloud"
	"siphon/app/djicloud/internal/svc"
	"siphon/common/tool"
	"siphon/third_party/extproto"

	"github.com/zeromicro/go-zero/core/logx"
)

type CustomDataTransmissionToPsdkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCustomDataTransmissionToPsdkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CustomDataTransmissionToPsdkLogic {
	return &CustomDataTransmissionToPsdkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CustomDataTransmissionToPsdkLogic) CustomDataTransmissionToPsdk(in *djicloud.CustomDataTransmissionToPsdkReq) (*djicloud.CommonRes, error) {
	l.Infof("[psdk-transmit] sn=%s value_len=%d", in.DeviceSn, len(in.Value))

	if len(in.Value) >= 256 {
		return nil, tool.NewErrorByPbCode(extproto.Code__1_01_PARAM_INVALID, "value 长度必须小于 256")
	}

	tid, err := l.svcCtx.DjiClient.CustomDataTransmissionToPsdk(l.ctx, in.DeviceSn, in.Value)
	if err != nil {
		return commandRes(tid, err)
	}

	return okRes(tid), nil
}
