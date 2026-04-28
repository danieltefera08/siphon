package logic

import (
	"context"

	"siphon/app/djigateway/djigateway"
	"siphon/app/djigateway/internal/svc"
	"siphon/common/djisdk"

	"github.com/zeromicro/go-zero/core/logx"
)

type PsdkUIResourceUploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPsdkUIResourceUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PsdkUIResourceUploadLogic {
	return &PsdkUIResourceUploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PsdkUIResourceUpload PSDK UI 资源上传。
func (l *PsdkUIResourceUploadLogic) PsdkUIResourceUpload(in *djigateway.PsdkUIResourceUploadReq) (*djigateway.CommonRes, error) {
	data := &djisdk.PsdkUIResourceUploadData{
		Name:        in.GetName(),
		URL:         in.GetUrl(),
		Fingerprint: in.GetFingerprint(),
	}
	tid, err := l.svcCtx.DjiClient.PsdkUIResourceUpload(l.ctx, in.GetDeviceSn(), data)
	if err != nil {
		l.Errorf("[psdk] ui resource upload failed device_sn=%s tid=%s: %v", in.GetDeviceSn(), tid, err)
		return errRes(tid, err), nil
	}
	return okRes(tid), nil
}
