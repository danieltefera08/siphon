package logic

import (
	"context"

	"siphon/app/oryxserver/internal/svc"
	"siphon/app/oryxserver/oryxserver"
	"siphon/common/tool"
	"siphon/third_party/extproto"

	"github.com/zeromicro/go-zero/core/logx"
)

type DvrApplyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDvrApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DvrApplyLogic {
	return &DvrApplyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 应用 DVR 云录制配置
func (l *DvrApplyLogic) DvrApply(in *oryxserver.DvrApplyReq) (*oryxserver.DvrApplyRes, error) {
	err := l.svcCtx.OryxClient.DvrApply(l.ctx, in.All)
	if err != nil {
		l.Logger.Errorf("调用 Oryx API 失败: %v", err)
		return nil, tool.NewErrorByPbCodeWrap(extproto.Code__1_06_THIRD_PARTY, err, "调用 Oryx API 失败")
	}
	return &oryxserver.DvrApplyRes{}, nil
}
