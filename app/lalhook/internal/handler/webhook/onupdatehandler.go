package webhook

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"siphon/app/lalhook/internal/logic/webhook"
	"siphon/app/lalhook/internal/svc"
	"siphon/app/lalhook/internal/types"
)

// 定时汇报所有group、session的信息
func OnUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OnUpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := webhook.NewOnUpdateLogic(r.Context(), svcCtx)
		resp, err := l.OnUpdate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
