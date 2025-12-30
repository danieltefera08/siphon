package svc

import (
	"siphon/common/socketio"
	"siphon/socketapp/socketpush/internal/config"
)

type ServiceContext struct {
	Config          config.Config
	SocketContainer *socketiox.SocketContainer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		SocketContainer: socketiox.MustNewPubContainer(c.SocketGtwConf),
	}
}
