package logic

import (
	"context"

	"hellozero/apibookstore/internal/svc"
	"hellozero/apibookstore/internal/types"
	"hellozero/rpc/add/adder"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.AddReq) (resp *types.AddResp, err error) {
	addResp, err := l.svcCtx.Add.Add(l.ctx, &adder.AddReq{
		Book:  req.Book,
		Price: req.Price,
	})
	if err != nil {
		return
	}

	return &types.AddResp{
		Ok: addResp.Ok,
	}, nil
}
