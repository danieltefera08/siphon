package main

import (
	"flag"
	"fmt"
	interceptor "siphon/common/Interceptor/rpcserver"

	"siphon/app/ispserver/internal/config"
	"siphon/app/ispserver/internal/server"
	"siphon/app/ispserver/internal/svc"
	"siphon/app/ispserver/ispserver"

	_ "siphon/common/carbonx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/ispserver.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		ispserver.RegisterIspServerServer(grpcServer, server.NewIspServerServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	s.AddUnaryInterceptors(interceptor.LoggerInterceptor)
	logx.AddGlobalFields(logx.Field("app", c.Name))

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()
	serviceGroup.Add(s)
	serviceGroup.Add(ctx.IspServer)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	fmt.Printf("ISP TCP server listening on %s...\n", c.IspConf.ListenAddr)
	serviceGroup.Start()
}
