package logic

import (
	"context"

	"siphon/app/djigateway/djigateway"
	"siphon/app/djigateway/internal/svc"
	"siphon/common/djisdk"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoteLogFileUploadCancelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoteLogFileUploadCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoteLogFileUploadCancelLogic {
	return &RemoteLogFileUploadCancelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// RemoteLogFileUploadCancel 取消远程日志文件上传任务。
func (l *RemoteLogFileUploadCancelLogic) RemoteLogFileUploadCancel(in *djigateway.RemoteLogFileUploadReq) (*djigateway.CommonRes, error) {
	files := make([]djisdk.RemoteLogFile, 0, len(in.Files))
	for _, f := range in.Files {
		files = append(files, djisdk.RemoteLogFile{
			DeviceSN: f.DeviceSn,
			Module:   f.Module,
			Key:      f.Key,
			Name:     f.Name,
			URL:      f.Url,
			Size:     f.Size,
		})
	}
	data := &djisdk.RemoteLogFileUploadCancelData{Files: files}
	tid, err := l.svcCtx.DjiClient.RemoteLogFileUploadCancel(l.ctx, in.DeviceSn, data)
	if err != nil {
		l.Errorf("[remote-log] file upload cancel failed: %v", err)
		return errRes(tid, err), nil
	}
	return okRes(tid), nil
}
