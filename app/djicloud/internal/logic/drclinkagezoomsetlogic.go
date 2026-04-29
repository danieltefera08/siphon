package logic

import (
	"context"

	"siphon/app/djicloud/djicloud"
	"siphon/app/djicloud/internal/svc"
	"siphon/common/djisdk"

	"github.com/zeromicro/go-zero/core/logx"
)

type DrcLinkageZoomSetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDrcLinkageZoomSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DrcLinkageZoomSetLogic {
	return &DrcLinkageZoomSetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DrcLinkageZoomSetLogic) DrcLinkageZoomSet(in *djicloud.DrcLinkageZoomSetReq) (*djicloud.CommonRes, error) {
	data := &djisdk.DrcLinkageZoomSetData{PayloadIndex: in.GetPayloadIndex(), State: in.GetState()}
	tid, err := l.svcCtx.DjiClient.DrcLinkageZoomSet(l.ctx, in.GetDeviceSn(), data)
	if err != nil {
		return errRes(tid, err), nil
	}
	return okRes(tid), nil
}
