package logic

import (
	"context"

	"siphon/app/gis/gis"
	"siphon/app/gis/internal/svc"
	"siphon/common/tool"
	"siphon/third_party/extproto"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFenceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFenceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFenceLogic {
	return &GetFenceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetFence 按 ID 获取围栏详情（含多边形顶点和空间索引 cells）。
func (l *GetFenceLogic) GetFence(in *gis.GetFenceReq) (*gis.GetFenceRes, error) {
	if in.FenceId == "" {
		return nil, tool.NewErrorByPbCode(extproto.Code__1_01_PARAM_MISSING, "fenceId")
	}

	info, err := l.svcCtx.FenceStore.GetFence(l.ctx, in.FenceId)
	if err != nil {
		return nil, err
	}

	return &gis.GetFenceRes{
		Fence: fenceInfoToDetail(info),
	}, nil
}
