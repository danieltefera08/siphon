// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package live

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"siphon/app/livegtw/internal/logic/live"
	"siphon/app/livegtw/internal/svc"
	"siphon/app/livegtw/internal/types"
)

// 停止会议录制
func StopMeetingRecordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StopMeetingRecordRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := live.NewStopMeetingRecordLogic(r.Context(), svcCtx)
		resp, err := l.StopMeetingRecord(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
