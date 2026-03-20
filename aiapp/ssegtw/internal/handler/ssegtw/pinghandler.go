// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package ssegtw

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"siphon/aiapp/ssegtw/internal/logic/ssegtw"
	"siphon/aiapp/ssegtw/internal/svc"
)

// ping
func PingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := ssegtw.NewPingLogic(r.Context(), svcCtx)
		resp, err := l.Ping()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
