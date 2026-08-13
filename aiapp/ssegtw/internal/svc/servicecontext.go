// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"time"

	"siphon/aiapp/ssegtw/internal/config"
	"siphon/common/antsx"
	"siphon/common/grpcx"
	"siphon/zerorpc/zerorpc"

	"github.com/zeromicro/go-zero/zrpc"
)

// SSEEvent SSE 事件结构
type SSEEvent struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}

type ServiceContext struct {
	Config     config.Config
	ZeroRpcCli zerorpc.ZerorpcClient
	Emitter    *antsx.EventEmitter[SSEEvent]
	PendingReg *antsx.ReplyPool[string]
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		ZeroRpcCli: zerorpc.NewZerorpcClient(zrpc.MustNewClient(c.ZeroRpcConf,
			zrpc.WithUnaryClientInterceptor(grpcx.UnaryMetadataInterceptor)).Conn()),
		Emitter:    antsx.NewEventEmitter[SSEEvent](),
		PendingReg: antsx.NewReplyPool[string](antsx.WithDefaultTTL(60 * time.Second)),
	}
}
