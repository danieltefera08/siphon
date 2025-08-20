package webhook

import (
	"context"

	"siphon/app/lalhook/internal/svc"
	"siphon/app/lalhook/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OnRelayPullStartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 回源拉流成功
func NewOnRelayPullStartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OnRelayPullStartLogic {
	return &OnRelayPullStartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OnRelayPullStartLogic) OnRelayPullStart(req *types.OnRelayPullStartRequest) (resp *types.EmptyReply, err error) {
	// todo: add your logic here and delete this line

	return
}
