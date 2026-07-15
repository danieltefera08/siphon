package handler

import (
	"context"

	"siphon/common/gnetx"
	"siphon/common/isp"
)

// HandleHeartbeat 处理 251-2 心跳指令，对标 Java SipEndpoint.T2512。
func HandleHeartbeat(ctx context.Context, conn gnetx.Conn, req *isp.Message) (*isp.Message, error) {
	return isp.NewResponse(req, isp.SessionSourceServer, isp.StatusSuccess, isp.CommandGenericResponseWithoutItems, nil), nil
}
