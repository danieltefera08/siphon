package file

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"siphon/gtw/internal/logic/file"
	"siphon/gtw/internal/svc"
	"siphon/gtw/internal/types"
)

// 生成文件url
func SignUrlHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SignUrlRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := file.NewSignUrlLogic(r.Context(), svcCtx)
		resp, err := l.SignUrl(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
