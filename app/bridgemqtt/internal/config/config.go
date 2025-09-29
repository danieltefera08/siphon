package config

import (
	"siphon/common/mqttx"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	MqttConfig      mqttx.MqttConfig
	StreamEventConf zrpc.RpcClientConf
}
