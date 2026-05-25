package logic

import (
	"context"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"

	"siphon/aiapp/aisolo/aisolo"
	"siphon/aiapp/aisolo/internal/svc"
	"siphon/common/tool"
	"siphon/third_party/extproto"
)

type GetSessionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSessionLogic {
	return &GetSessionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSessionLogic) GetSession(in *aisolo.GetSessionReq) (*aisolo.GetSessionResp, error) {
	userID := strings.TrimSpace(in.GetUserId())
	sessionID := strings.TrimSpace(in.GetSessionId())
	if userID == "" || sessionID == "" {
		return nil, tool.NewErrorByPbCode(extproto.Code__1_01_PARAM_MISSING, "user_id and session_id are required")
	}
	sess, err := l.svcCtx.Sessions.GetSession(l.ctx, userID, sessionID)
	if err != nil {
		return nil, err
	}
	return &aisolo.GetSessionResp{Session: toProtoSession(sess)}, nil
}
