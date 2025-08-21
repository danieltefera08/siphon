package svc

import (
	"siphon/app/lalhook/internal/config"
	"siphon/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config          config.Config
	HlsTsFilesModel model.HlsTsFilesModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		HlsTsFilesModel: model.NewHlsTsFilesModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
