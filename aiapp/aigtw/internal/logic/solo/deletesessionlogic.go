package solo

import (
	"context"
	"errors"

	"siphon/aiapp/aigtw/internal/svc"
	"siphon/aiapp/aigtw/internal/types"
	"siphon/aiapp/aisolo/aisolo"
	"siphon/common/ctxdata"

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

func (l *DeleteSessionLogic) DeleteSession(req *types.SoloDeleteSessionRequest) (*types.SoloDeleteSessionResponse, error) {
	userID := ctxdata.GetUserId(l.ctx)
	if userID == "" {
		return nil, errors.New("missing user id in context")
	}
	resp, err := l.svcCtx.AiSoloCli.DeleteSession(l.ctx, &aisolo.DeleteSessionReq{
		SessionId: req.SessionId,
		UserId:    userID,
	})
	if err != nil {
		return nil, err
	}
	return &types.SoloDeleteSessionResponse{Success: resp.GetSuccess()}, nil
}
