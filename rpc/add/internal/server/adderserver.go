// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: add.proto

package server

import (
	"context"

	"hellozero/rpc/add/add"
	"hellozero/rpc/add/internal/logic"
	"hellozero/rpc/add/internal/svc"
)

type AdderServer struct {
	svcCtx *svc.ServiceContext
	add.UnimplementedAdderServer
}

func NewAdderServer(svcCtx *svc.ServiceContext) *AdderServer {
	return &AdderServer{
		svcCtx: svcCtx,
	}
}

func (s *AdderServer) Add(ctx context.Context, in *add.AddReq) (*add.AddResp, error) {
	l := logic.NewAddLogic(ctx, s.svcCtx)
	return l.Add(in)
}