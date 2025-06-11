package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"siphon/app/xfusionmock/cron"
	"siphon/app/xfusionmock/internal/config"
	"siphon/app/xfusionmock/internal/server"
	"siphon/app/xfusionmock/internal/svc"
	"siphon/app/xfusionmock/kafka"
	"siphon/app/xfusionmock/xfusionmock"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	_ "siphon/common/carbonx"
)

var configFile = flag.String("f", "etc/xfusionmock.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		xfusionmock.RegisterXFusionMockRpcServer(grpcServer, server.NewXFusionMockRpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()
	logx.AddGlobalFields(logx.Field("app", c.Name))
	serviceGroup.Add(s)

	// kafka
	serviceGroup.Add(kq.MustNewQueue(c.KafkaTestConfig, kafka.NewTest(ctx)))

	// cron
	serviceGroup.Add(cron.NewCronService(ctx))
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	serviceGroup.Start()
}
