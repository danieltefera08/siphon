package logic

import (
	"context"
	"strconv"

	"siphon/app/djicloud/djicloud"
	"siphon/app/djicloud/internal/svc"
	"siphon/common/djisdk"

	"github.com/zeromicro/go-zero/core/logx"
)

type DrcIntervalPhotoSetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDrcIntervalPhotoSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DrcIntervalPhotoSetLogic {
	return &DrcIntervalPhotoSetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DrcIntervalPhotoSetLogic) DrcIntervalPhotoSet(in *djicloud.DrcIntervalPhotoSetReq) (*djicloud.CommonRes, error) {
	interval, _ := strconv.ParseFloat(in.GetInterval(), 64)
	data := &djisdk.DrcIntervalPhotoSetData{PayloadIndex: in.GetPayloadIndex(), Interval: interval}
	tid, err := l.svcCtx.DjiClient.DrcIntervalPhotoSet(l.ctx, in.GetDeviceSn(), data)
	if err != nil {
		return errRes(tid, err), nil
	}
	return okRes(tid), nil
}
