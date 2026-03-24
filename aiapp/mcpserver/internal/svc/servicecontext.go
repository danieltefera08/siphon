package svc

import (
	"siphon/aiapp/mcpserver/internal/config"
	"siphon/app/bridgemodbus/bridgemodbus"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	BridgeModbusCli bridgemodbus.BridgeModbusClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		BridgeModbusCli: bridgemodbus.NewBridgeModbusClient(
			zrpc.MustNewClient(c.BridgeModbusRpcConf).Conn()),
	}
}
