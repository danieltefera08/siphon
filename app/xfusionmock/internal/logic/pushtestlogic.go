package logic

import (
	"context"
	"fmt"
	"time"

	"siphon/app/xfusionmock/internal/svc"
	"siphon/app/xfusionmock/xfusionmock"

	"github.com/zeromicro/go-zero/core/logx"
)

type PushTestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPushTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushTestLogic {
	return &PushTestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PushTestLogic) PushTest(in *xfusionmock.ReqPushTest) (*xfusionmock.ResPushTest, error) {
	l.Info("PushTest")
	l.svcCtx.KafkaTestPusher.Push(context.Background(), fmt.Sprintf("%s 定时消息 @ %v", l.svcCtx.Config.Name, time.Now()))
	return &xfusionmock.ResPushTest{}, nil
}
