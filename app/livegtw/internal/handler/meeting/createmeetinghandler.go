// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package meeting

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"siphon/app/livegtw/internal/logic/meeting"
	"siphon/app/livegtw/internal/svc"
	"siphon/app/livegtw/internal/types"

	xhttp "github.com/zeromicro/x/http"
)

// 创建会议
func CreateMeetingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateMeetingReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		l := meeting.NewCreateMeetingLogic(r.Context(), svcCtx)
		resp, err := l.CreateMeeting(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
