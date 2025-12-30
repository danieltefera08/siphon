package handler

import (
	"net/http"
	"siphon/common/socketiox"
	"siphon/socketapp/socketgtw/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/socket.io",
				Handler: socketiox.SocketioHandler(serverCtx.SocketServer),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
	)
}
