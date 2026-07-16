package config

import (
	"siphon/common/isp"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	IspConf isp.ServerConfig
}
