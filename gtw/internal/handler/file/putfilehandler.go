package file

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"siphon/gtw/internal/logic/file"
	"siphon/gtw/internal/svc"
	"siphon/gtw/internal/types"
)

// 上传文件
func PutFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PutFileRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := file.NewPutFileLogic(r.Context(), svcCtx, r)
		resp, err := l.PutFile(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
