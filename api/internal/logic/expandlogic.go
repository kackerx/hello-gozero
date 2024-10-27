package logic

import (
	"context"

	"hellozero/api/internal/svc"
	"hellozero/api/internal/types"
	"hellozero/rpc/transform/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExpandLogic) Expand(req *types.ExpandReq) (resp *types.ExpandResp, err error) {
	rpcResp, err := l.svcCtx.Transformer.Expand(l.ctx, &transform.ExpandReq{Shorten: req.Shorten})
	if err != nil {
		return
	}

	return &types.ExpandResp{Url: rpcResp.Url}, nil
}
