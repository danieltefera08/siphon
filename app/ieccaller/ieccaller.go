package main

import (
	"flag"
	"fmt"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/zeromicro/go-zero/core/logx"
	"siphon/app/ieccaller/cron"
	"siphon/app/ieccaller/ieccaller"
	"siphon/app/ieccaller/internal/config"
	"siphon/app/ieccaller/internal/iec"
	"siphon/app/ieccaller/internal/server"
	"siphon/app/ieccaller/internal/svc"
	interceptor "siphon/common/Interceptor/rpcserver"
	"siphon/common/nacos"
	"siphon/iec104/iec104client"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	_ "siphon/common/carbonx"
)

var configFile = flag.String("f", "etc/ieccaller.yaml", "the config file")

// GOARCH=amd64 GOOS=linux GOOS=linux go build -o app
// GOARCH=arm GOOS=linux go build -o app
// go build -o app
func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		ieccaller.RegisterIecCallerServer(grpcServer, server.NewIecCallerServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	// register service to nacos
	if c.NacosConfig.IsRegister {
		sc := []constant.ServerConfig{
			*constant.NewServerConfig(c.NacosConfig.Host, c.NacosConfig.Port),
		}
		cc := &constant.ClientConfig{
			NamespaceId:         c.NacosConfig.NamespaceId,
			Username:            c.NacosConfig.Username,
			Password:            c.NacosConfig.PassWord,
			TimeoutMs:           5000,
			NotLoadCacheAtStart: true,
			LogDir:              "/tmp/nacos/log",
			CacheDir:            "/tmp/nacos/cache",
			//RotateTime:          "1h",
			//MaxAge:              3,
			LogLevel: "debug",
		}
		m := map[string]string{
			"gRPC_port":                 strutil.After(c.RpcServerConf.ListenOn, ":"),
			"preserved.register.source": "go-zero",
		}
		opts := nacos.NewNacosConfig(c.NacosConfig.ServiceName, c.ListenOn, sc, cc, nacos.WithMetadata(m))
		_ = nacos.RegisterService(opts)
	}
	s.AddUnaryInterceptors(interceptor.LoggerInterceptor)
	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()
	logx.AddGlobalFields(logx.Field("app", c.Name))
	serviceGroup.Add(s)

	for _, cf := range c.IecServerConfig {
		serviceGroup.Add(iec104client.MustNewIecServerClient(cf, iec.NewClientCall(ctx, cf.Host, cf.Port, cf.MetaData), ctx.ClientManager))
	}

	// cron
	serviceGroup.Add(cron.NewCronService(ctx))
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	serviceGroup.Start()
}
