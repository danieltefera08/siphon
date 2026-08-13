package main

import (
	"flag"
	"fmt"

	"siphon/aiapp/aichat/aichat"
	"siphon/aiapp/aichat/internal/config"
	"siphon/aiapp/aichat/internal/server"
	"siphon/aiapp/aichat/internal/svc"
	"siphon/common/grpcx"
	"siphon/common/tool"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/aichat.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	tool.PrintGoVersion()

	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		aichat.RegisterAiChatServer(grpcServer, server.NewAiChatServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	s.AddUnaryInterceptors(grpcx.LoggerInterceptor)
	s.AddStreamInterceptors(grpcx.StreamLoggerInterceptor)
	defer s.Stop()

	logx.AddGlobalFields(logx.Field("app", c.Name))

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
