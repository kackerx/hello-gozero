package svc

import (
	"github.com/zeromicro/go-zero/zrpc"

	"hellozero/api/internal/config"
	"hellozero/rpc/transform/transformer"
)

type ServiceContext struct {
	Config      config.Config
	Transformer transformer.Transformer // client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		Transformer: transformer.NewTransformer(zrpc.MustNewClient(c.Transform)),
	}
}
