package logic

import (
	"context"

	"siphon/app/lalproxy/internal/svc"
	"siphon/app/lalproxy/lalproxy"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartRelayPullLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStartRelayPullLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartRelayPullLogic {
	return &StartRelayPullLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 控制服务器从远端拉流至本地
func (l *StartRelayPullLogic) StartRelayPull(in *lalproxy.StartRelayPullReq) (*lalproxy.StartRelayPullRes, error) {
	// todo: add your logic here and delete this line

	return &lalproxy.StartRelayPullRes{}, nil
}
