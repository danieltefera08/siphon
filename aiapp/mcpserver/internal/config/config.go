package config

import (
	"siphon/common/mcpx"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	mcpx.McpServerConf
	BridgeModbusRpcConf zrpc.RpcClientConf
}
