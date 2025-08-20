package webhook

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"siphon/app/lalhook/internal/logic/webhook"
	"siphon/app/lalhook/internal/svc"
	"siphon/app/lalhook/internal/types"
)

// 别人推流到当前节点
func OnPubStartHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OnPubStartRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := webhook.NewOnPubStartLogic(r.Context(), svcCtx)
		resp, err := l.OnPubStart(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
