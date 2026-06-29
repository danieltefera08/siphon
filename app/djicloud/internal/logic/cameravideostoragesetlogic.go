package logic

import (
	"context"

	"siphon/app/djicloud/djicloud"
	"siphon/app/djicloud/internal/svc"
	"siphon/common/djisdk"

	"github.com/zeromicro/go-zero/core/logx"
)

type CameraVideoStorageSetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCameraVideoStorageSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CameraVideoStorageSetLogic {
	return &CameraVideoStorageSetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CameraVideoStorageSetLogic) CameraVideoStorageSet(in *djicloud.CameraVideoStorageSetReq) (*djicloud.CommonRes, error) {
	data := &djisdk.CameraStorageSetData{
		PayloadIndex: in.PayloadIndex,
		StorageType:  int(in.VideoStorageType),
	}
	tid, err := l.svcCtx.DjiClient.CameraVideoStorageSet(l.ctx, in.DeviceSn, data)
	if err != nil {
		return errRes(tid, err), nil
	}
	return okRes(tid), nil
}
