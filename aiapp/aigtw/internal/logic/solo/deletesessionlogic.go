package solo

import (
	"context"
	"siphon/aiapp/aisolo/aisolo"

	"siphon/aiapp/aigtw/internal/svc"
	"siphon/aiapp/aigtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSessionLogic {
	return &DeleteSessionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSessionLogic) DeleteSession(req *types.SoloDeleteSessionReq) (resp *types.SoloDeleteSessionResp, err error) {
	protoReq := &aisolo.SessionRequest{
		SessionId: req.SessionId,
	}

	_, err = l.svcCtx.EinoCli.DeleteSession(l.ctx, protoReq)
	if err != nil {
		l.Logger.Errorf("delete session failed: %v", err)
		return nil, err
	}

	return &types.SoloDeleteSessionResp{}, nil
}
