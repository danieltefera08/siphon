package svc

import (
	"siphon/aiapp/aichat/internal/config"
	"siphon/aiapp/aichat/internal/provider"
	"siphon/common/mcpx"

	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config    config.Config
	Registry  *provider.Registry
	McpClient *mcpx.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	logx.Must(logx.SetUp(c.Log))
	registry, err := provider.NewRegistry(c.Providers, c.Models)
	if err != nil {
		panic(err)
	}

	var mc *mcpx.Client
	mcpCfg := c.Mcpx
	if len(mcpCfg.Servers) > 0 {
		mc = mcpx.NewClient(mcpCfg)
	}

	return &ServiceContext{
		Config:    c,
		Registry:  registry,
		McpClient: mc,
	}
}
