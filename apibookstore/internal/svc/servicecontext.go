package svc

import (
	"github.com/zeromicro/go-zero/zrpc"

	"hellozero/apibookstore/internal/config"
	"hellozero/rpc/add/adder"
	"hellozero/rpc/check/checker"
)

type ServiceContext struct {
	Config config.Config
	Add    adder.Adder
	Check  checker.Checker
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Add:    adder.NewAdder(zrpc.MustNewClient(c.Add)),
		Check:  checker.NewChecker(zrpc.MustNewClient(c.Check)),
	}
}
