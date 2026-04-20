package solo

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"siphon/aiapp/aigtw/internal/logic/solo"
	"siphon/aiapp/aigtw/internal/svc"
)

func ListSkillsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := solo.NewListSkillsLogic(r.Context(), svcCtx)
		resp, err := l.ListSkills()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
