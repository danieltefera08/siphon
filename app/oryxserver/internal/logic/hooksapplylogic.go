package logic

import (
	"context"

	"siphon/app/oryxserver/internal/svc"
	"siphon/app/oryxserver/oryxserver"
	"siphon/common/oryxx"
	"siphon/common/tool"
	"siphon/third_party/extproto"

	"github.com/zeromicro/go-zero/core/logx"
)

type HooksApplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHooksApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HooksApplyLogic {
	return &HooksApplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 应用回调配置
func (l *HooksApplyLogic) HooksApply(in *oryxserver.HooksApplyReq) (*oryxserver.HooksApplyRes, error) {
	err := l.svcCtx.OryxClient.HooksApply(l.ctx, oryxx.HooksApplyReq{
		Target: in.Target,
		Opaque: in.Opaque,
		All:    in.All,
		Host:   in.Host,
	})
	if err != nil {
		l.Logger.Errorf("调用 Oryx API 失败: %v", err)
		return nil, tool.NewErrorByPbCodeWrap(extproto.Code__1_06_THIRD_PARTY, err, "调用 Oryx API 失败")
	}
	return &oryxserver.HooksApplyRes{}, nil
}
