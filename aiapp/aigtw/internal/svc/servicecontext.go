package svc

import (
	"siphon/aiapp/aigtw/internal/config"
	interceptor "siphon/common/Interceptor/rpcclient"
	"siphon/zerorpc/zerorpc"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	ZeroRpcCli zerorpc.ZerorpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		ZeroRpcCli: zerorpc.NewZerorpcClient(zrpc.MustNewClient(c.ZeroRpcConf,
			zrpc.WithUnaryClientInterceptor(interceptor.UnaryMetadataInterceptor)).Conn()),
	}
}
