package logic

import (
	"context"

	"hellozero/apibookstore/internal/svc"
	"hellozero/apibookstore/internal/types"
	"hellozero/rpc/check/check"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckLogic) Check(req *types.CheckReq) (resp *types.CheckResp, err error) {
	checkResp, err := l.svcCtx.Check.Check(l.ctx, &check.CheckReq{Book: req.Book})
	if err != nil {
		return
	}

	return &types.CheckResp{
		Found: checkResp.Found,
		Price: checkResp.Price,
	}, nil
}
