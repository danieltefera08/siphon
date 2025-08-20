package webhook

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"siphon/app/lalhook/internal/logic/webhook"
	"siphon/app/lalhook/internal/svc"
	"siphon/app/lalhook/internal/types"
)

// 回源拉流停止
func OnRelayPullStopHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OnRelayPullStopRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := webhook.NewOnRelayPullStopLogic(r.Context(), svcCtx)
		resp, err := l.OnRelayPullStop(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
