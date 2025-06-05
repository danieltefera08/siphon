package gtw

import (
	"context"
	"encoding/json"

	"siphon/gtw/internal/svc"
	"siphon/gtw/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ForwardLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewForwardLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ForwardLogic {
	return &ForwardLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ForwardLogic) Forward(req *types.ForwardRequest) (resp *types.ForwardReply, err error) {
	content, _ := json.Marshal(req)
	logx.Infof("forward req:%s", content)
	return &types.ForwardReply{}, nil
}
