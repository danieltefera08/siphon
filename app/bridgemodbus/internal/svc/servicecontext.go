package svc

import (
	"siphon/app/bridgemodbus/internal/config"
	"siphon/common/modbusx"
)

type ServiceContext struct {
	Config           config.Config
	ModbusClientPool *modbusx.ModbusClientPool
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		ModbusClientPool: modbusx.NewModbusClientPool(&c.ModbusClientConf, 3),
	}
}
