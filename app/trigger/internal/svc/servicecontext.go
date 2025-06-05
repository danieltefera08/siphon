package svc

import (
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/rest/httpc"
	"siphon/app/trigger/internal/config"
	"siphon/common/asynqx"
)

type ServiceContext struct {
	Config      config.Config
	AsynqClient *asynq.Client
	AsynqServer *asynq.Server
	Scheduler   *asynq.Scheduler
	Httpc       httpc.Service
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		AsynqClient: asynqx.NewAsynqClient(c.Redis.Host, c.Redis.Pass),
		AsynqServer: asynqx.NewAsynqServer(c.Redis.Host, c.Redis.Pass),
		Scheduler:   asynqx.NewScheduler(c.Redis.Host, c.Redis.Pass),
		Httpc:       httpc.NewService("httpc"),
	}
}
