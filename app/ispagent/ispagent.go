package main

import (
	"flag"
	"fmt"
	"siphon/common/grpcx"

	"siphon/app/ispagent/internal/config"
	"siphon/app/ispagent/internal/server"
	"siphon/app/ispagent/internal/svc"
	"siphon/app/ispagent/ispagent"

	_ "siphon/common/carbonx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/ispagent.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		ispagent.RegisterIspAgentServer(grpcServer, server.NewIspAgentServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	s.AddUnaryInterceptors(grpcx.LoggerInterceptor)
	logx.AddGlobalFields(logx.Field("app", c.Name))
	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()
	serviceGroup.Add(s)
	if ctx.Scheduler != nil {
		serviceGroup.Add(ctx.Scheduler)
	}

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	serviceGroup.Start()
}
