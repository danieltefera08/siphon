package logic

import (
	"context"
	"time"
	"siphon/common/asynqx"
	"siphon/zerorpc/internal/svc"
	"siphon/zerorpc/internal/taskpayload"
	"siphon/zerorpc/zerorpc"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/jsonx"
	"go.opentelemetry.io/otel/propagation"
	"siphon/common/trace"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendDelayTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendDelayTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendDelayTaskLogic {
	return &SendDelayTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发送延迟任务
func (l *SendDelayTaskLogic) SendDelayTask(in *zerorpc.SendDelayTaskReq) (*zerorpc.SendDelayTaskRes, error) {
	spanCtx, span := svc.StartAsynqProducerSpan(l.ctx, asynqx.DeferDelayTask)
	defer span.End()
	carrier := &propagation.HeaderCarrier{}
	trace.Inject(spanCtx, carrier)
	msg := &taskpayload.HttpPayload{
		MsgId:   in.MsgId,
		Carrier: carrier,
		Msg:     in.Body,
	}
	payload, err := jsonx.Marshal(msg)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(asynqx.DeferDelayTask, []byte(payload)), asynq.TaskID(in.GetMsgId()), asynq.ProcessIn(time.Duration(in.ProcessIn)*time.Minute), asynq.Retention(7*24*time.Hour))
	if err != nil {
		return nil, err
	}
	return &zerorpc.SendDelayTaskRes{}, nil
}
