package svc

import (
	"siphon/aiapp/aichat/internal/config"
	"siphon/aiapp/aichat/internal/provider"
)

type ServiceContext struct {
	Config   config.Config
	Registry *provider.Registry
}

func NewServiceContext(c config.Config) *ServiceContext {
	registry, err := provider.NewRegistry(c.Providers, c.Models)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:   c,
		Registry: registry,
	}
}
