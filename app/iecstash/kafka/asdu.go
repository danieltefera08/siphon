package kafka

import (
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/net/context"
	"siphon/app/iecstash/internal/svc"
	"siphon/facade/iecstream/iecstream"
)

type Asdu struct {
	svcCtx *svc.ServiceContext
}

func NewAsdu(svcCtx *svc.ServiceContext) *Asdu {
	return &Asdu{
		svcCtx: svcCtx,
	}
}

func (l Asdu) Consume(ctx context.Context, key, value string) error {
	logx.Infof("asdu, key: %+v msg:%+v", key, value)
	msgBody := iecstream.MsgBody{}
	jsonx.Unmarshal([]byte(value), &msgBody)
	l.svcCtx.IecStreamRpcCli.PushChunkAsdu(ctx, &iecstream.PushChunkAsduReq{
		Info: &msgBody,
	})
	return nil
}
