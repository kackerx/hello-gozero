package logic

import (
	"context"

	"hellozero/api/internal/svc"
	"hellozero/api/internal/types"
	"hellozero/rpc/transform/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortenLogic) Shorten(req *types.ShortenReq) (resp *types.ShortenResp, err error) {
	shorten, err := l.svcCtx.Transformer.Shorten(l.ctx, &transform.ShortenReq{Url: req.Url})
	if err != nil {
		return
	}

	return &types.ShortenResp{shorten.Shorten}, nil
}
