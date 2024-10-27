package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"hellozero/rpc/transform/internal/config"
	"hellozero/rpc/transform/model"
)

type ServiceContext struct {
	Config config.Config
	Model  model.ShorturlModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewShorturlModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
