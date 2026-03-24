// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package pass

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"siphon/aiapp/aigtw/internal/logic/pass"
	"siphon/aiapp/aigtw/internal/svc"
)

// 列出可用模型
func ListModelsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := pass.NewListModelsLogic(r.Context(), svcCtx)
		resp, err := l.ListModels()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
