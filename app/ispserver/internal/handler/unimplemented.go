package handler

import (
	"context"

	"siphon/common/gnetx"
	"siphon/common/isp"
)

// HandleUnimplemented 对所有未实现指令返回 251-3 code=500，对标 Java SipHandlerInterceptor.notSupported。
func HandleUnimplemented(ctx context.Context, conn gnetx.Conn, req *isp.Message) (*isp.Message, error) {
	return isp.NewResponse(req, isp.SessionSourceServer, isp.ResponseCode(isp.ErrUnimplemented), isp.CommandGenericResponseWithoutItems, nil), nil
}
