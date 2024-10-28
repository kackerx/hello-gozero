package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"hellozero/rpc/add/internal/config"
	"hellozero/rpc/model"
)

type ServiceContext struct {
	Config config.Config
	Model  model.BookModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewBookModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
