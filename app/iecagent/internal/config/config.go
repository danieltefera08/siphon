package config

import (
	iecserver "siphon/common/iec104/server"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	IecServer iecserver.ServerConfig
}
