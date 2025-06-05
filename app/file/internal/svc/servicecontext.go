package svc

import (
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"siphon/app/file/internal/config"
	"siphon/model"
)

type ServiceContext struct {
	Config   config.Config
	Validate *validator.Validate
	OssModel model.OssModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		Validate: validator.New(),
		OssModel: model.NewOssModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
