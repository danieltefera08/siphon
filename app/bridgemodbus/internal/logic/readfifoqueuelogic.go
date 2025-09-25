package logic

import (
	"context"

	"siphon/app/bridgemodbus/bridgemodbus"
	"siphon/app/bridgemodbus/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReadFIFOQueueLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReadFIFOQueueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReadFIFOQueueLogic {
	return &ReadFIFOQueueLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 读取 FIFO 队列 (Function Code 0x18)
func (l *ReadFIFOQueueLogic) ReadFIFOQueue(in *bridgemodbus.ReadFIFOQueueReq) (*bridgemodbus.ReadFIFOQueueRes, error) {
	// todo: add your logic here and delete this line

	return &bridgemodbus.ReadFIFOQueueRes{}, nil
}
