package webhook

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"siphon/app/lalhook/internal/logic/webhook"
	"siphon/app/lalhook/internal/svc"
	"siphon/app/lalhook/internal/types"
)

// 回源拉流成功
func OnRelayPullStartHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OnRelayPullStartRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := webhook.NewOnRelayPullStartLogic(r.Context(), svcCtx)
		resp, err := l.OnRelayPullStart(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
