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

type GetSessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSessionLogic {
	return &GetSessionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSessionLogic) GetSession(req *types.SoloGetSessionRequest) (*types.SoloGetSessionResponse, error) {
	userID := ctxdata.GetUserId(l.ctx)
	if userID == "" {
		return nil, errors.New("missing user id in context")
	}
	resp, err := l.svcCtx.AiSoloCli.GetSession(l.ctx, &aisolo.GetSessionReq{
		SessionId: req.SessionId,
		UserId:    userID,
	})
	if err != nil {
		return nil, err
	}
	return &types.SoloGetSessionResponse{Session: sessionToType(resp.GetSession())}, nil
}
