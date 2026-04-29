package logic

import (
	"context"

	"siphon/app/djicloud/djicloud"
	"siphon/app/djicloud/internal/svc"
	"siphon/common/djisdk"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoteLogFileListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoteLogFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoteLogFileListLogic {
	return &RemoteLogFileListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// RemoteLogFileList 查询可上传的远程日志文件列表。
func (l *RemoteLogFileListLogic) RemoteLogFileList(in *djicloud.RemoteLogFileListReq) (*djicloud.CommonRes, error) {
	data := &djisdk.RemoteLogFileListData{
		DeviceSN: in.TargetDeviceSn,
		Module:   in.Module,
	}
	tid, err := l.svcCtx.DjiClient.RemoteLogFileList(l.ctx, in.DeviceSn, data)
	if err != nil {
		l.Errorf("[remote-log] file list failed: %v", err)
		return errRes(tid, err), nil
	}
	return okRes(tid), nil
}
