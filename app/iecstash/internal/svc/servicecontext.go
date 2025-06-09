package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"siphon/app/iecstash/internal/config"
	interceptor "siphon/common/Interceptor/rpcclient"
	"siphon/facade/iecstream/iecstream"
)

type ServiceContext struct {
	Config          config.Config
	IecStreamRpcCli iecstream.IecStreamRpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		IecStreamRpcCli: iecstream.NewIecStreamRpcClient(zrpc.MustNewClient(c.IecStreamRpcConf,
			zrpc.WithUnaryClientInterceptor(interceptor.UnaryMetadataInterceptor)).Conn()),
	}
}
