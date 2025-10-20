package main

import (
	"flag"
	"fmt"
	"siphon/app/bridgemqtt/bridgemqtt"
	"siphon/app/bridgemqtt/internal/config"
	"siphon/app/bridgemqtt/internal/server"
	"siphon/app/bridgemqtt/internal/svc"
	interceptor "siphon/common/Interceptor/rpcserver"
	"siphon/common/nacosx"

	_ "siphon/common/carbonx"
	_ "siphon/common/nacosx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/bridgemqtt.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		bridgemqtt.RegisterBridgeMqttServer(grpcServer, server.NewBridgeMqttServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	nacosx.SetUpLogger(nacosx.LoggerConfig{
		AppendToStdout: true,
		Level:          "error",
		LogDir:         "/tmp/nacos/log",
	})
	s.AddUnaryInterceptors(interceptor.LoggerInterceptor)
	logx.AddGlobalFields(logx.Field("app", c.Name))
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
