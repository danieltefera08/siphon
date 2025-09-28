package svc

import (
	"siphon/app/bridgemqtt/internal/config"
	"siphon/common/mqttx"
)

type ServiceContext struct {
	Config     config.Config
	MqttClient *mqttx.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		MqttClient: mqttx.MustNewClient(c.MqttConfig),
	}
}
