package file

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"siphon/gtw/internal/logic/file"
	"siphon/gtw/internal/svc"
	"siphon/gtw/internal/types"
)

// 获取文件信息
func StatFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StatFileRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := file.NewStatFileLogic(r.Context(), svcCtx)
		resp, err := l.StatFile(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
