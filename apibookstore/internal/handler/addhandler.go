package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"hellozero/apibookstore/internal/logic"
	"hellozero/apibookstore/internal/svc"
	"hellozero/apibookstore/internal/types"
)

func AddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc { // info Handler直接对接http handlerFunc 无业务逻辑
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAddLogic(r.Context(), svcCtx)
		resp, err := l.Add(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
