package logic

import (
	"context"

	"siphon/app/ieccaller/ieccaller"
	"siphon/app/ieccaller/internal/svc"
	"siphon/common/iec104/client"
	"siphon/common/tool"
	"siphon/third_party/extproto"

	"github.com/wendy512/go-iecp5/asdu"
	"github.com/zeromicro/go-zero/core/logx"
)

type SendSetpointFloatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendSetpointFloatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSetpointFloatLogic {
	return &SendSetpointFloatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendSetpointFloatLogic) SendSetpointFloat(in *ieccaller.SendSetpointFloatReq) (*ieccaller.SendSetpointFloatRes, error) {
	cli, err := l.svcCtx.ClientManager.GetClientOrNil(in.Host, int(in.Port))
	if err != nil {
		return nil, tool.NewErrorByPbCodeWrap(extproto.Code__1_06_RPC, err, "获取IEC客户端失败")
	}
	if cli == nil && l.svcCtx.IsBroadcast() {
		var res ieccaller.SendSetpointFloatRes
		err = l.svcCtx.PushPbBroadcastWithAck(l.ctx, ieccaller.IecCaller_SendSetpointFloat_FullMethodName, in, &res)
		if err != nil {
			return nil, wrapCommandAckError(err, "集群推送ACK失败")
		}
		return &res, nil
	} else if cli != nil {
		ack, err := cli.SendSetpointFloatCmd(l.ctx, uint16(in.Coa), asdu.InfoObjAddr(in.Ioa), float32(in.Value), in.WithTime, client.WithAck())
		if err != nil {
			return nil, wrapCommandAckError(err, "IEC发送浮点设点命令失败")
		}
		value, err := ackFloat32Value(ack.Value)
		if err != nil {
			return nil, wrapCommandAckError(err, "IEC浮点设点ACK解析失败")
		}
		return &ieccaller.SendSetpointFloatRes{Value: float64(value)}, nil
	}
	return nil, tool.NewErrorByPbCode(extproto.Code__1_06_RPC, "IEC客户端不存在: %s:%d", in.Host, in.Port)
}
