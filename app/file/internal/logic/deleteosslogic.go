package logic

import (
	"context"

	"siphon/app/file/file"
	"siphon/app/file/internal/svc"
	"siphon/common/ossx"
	"siphon/model/gormmodel"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteOssLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOssLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOssLogic {
	return &DeleteOssLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteOssLogic) DeleteOss(in *file.DeleteOssReq) (*file.DeleteOssRes, error) {
	var oss gormmodel.Oss
	if err := l.svcCtx.DB.WithContext(l.ctx).First(&oss, in.Id).Error; err != nil {
		return nil, err
	}
	if err := l.svcCtx.DB.WithContext(l.ctx).Delete(&oss).Error; err != nil {
		return nil, err
	}
	ossx.CacheInvalidate(oss.TenantId)
	return &file.DeleteOssRes{Id: oss.Id}, nil
}
