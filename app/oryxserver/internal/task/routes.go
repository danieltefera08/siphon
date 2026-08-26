package task

import (
	"siphon/app/oryxserver/internal/relay"
	"siphon/app/oryxserver/internal/svc"
	"siphon/common/asynqx"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
)

func Register(svcCtx *svc.ServiceContext) *asynq.ServeMux {
	mux := asynqx.NewMux()

	mux.Handle(relay.RelayReconcileTask, NewReconcileHandler(svcCtx))
	logx.Infof("relay task registered: %s", relay.RelayReconcileTask)

	mux.Handle(relay.RelayStopTask, NewStopHandler(svcCtx))
	logx.Infof("relay task registered: %s", relay.RelayStopTask)

	return mux
}
