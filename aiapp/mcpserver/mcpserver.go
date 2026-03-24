package main

import (
	"flag"
	"fmt"

	"siphon/aiapp/mcpserver/internal/config"
	"siphon/aiapp/mcpserver/internal/svc"
	"siphon/aiapp/mcpserver/internal/tools"
	"siphon/common/mcpx"
	"siphon/common/tool"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/mcpserver.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	tool.PrintGoVersion()
	logx.DisableStat()

	// 创建带鉴权的 MCP 服务器（与 go-zero mcp.NewMcpServer 对齐）
	server := mcpx.NewMcpServer(c.McpServerConf)
	defer server.Stop()

	// 注册所有工具
	svcCtx := svc.NewServiceContext(c)
	tools.RegisterAll(server.Server(), svcCtx)

	fmt.Printf("Starting MCP server at %s:%d ...\n", c.Host, c.Port)
	server.Start()
}
