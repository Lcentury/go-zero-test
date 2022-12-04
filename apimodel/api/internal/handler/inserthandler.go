package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"test/apimodel/api/internal/logic"
	"test/apimodel/api/internal/svc"
	"test/apimodel/api/internal/types"
)

func InsertHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InsertReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewInsertLogic(r.Context(), svcCtx)
		resp, err := l.Insert(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
