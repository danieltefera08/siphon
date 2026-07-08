package logic

import (
	"context"

	"siphon/app/ispagent/internal/svc"
	"siphon/app/ispagent/ispagent"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExecuteCommandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExecuteCommandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExecuteCommandLogic {
	return &ExecuteCommandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExecuteCommandLogic) ExecuteCommand(in *ispagent.CommandReq) (*ispagent.CommandRes, error) {
	msg, err := l.svcCtx.IspClient.Execute(l.ctx, in.GetType(), in.GetCommand(), in.GetCode(), protoItems(in.GetItems()))
	if err != nil {
		return nil, err
	}
	return commandResponse(msg), nil
}
