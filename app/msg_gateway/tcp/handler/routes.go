package handler

import (
	"github.com/wymli/bcsns/app/msg_gateway/svc"
	"github.com/wymli/bcsns/app/msg_gateway/tcp/handler/gateway"
	"github.com/wymli/bcsns/common/server_framework/tcp"
)

func RegisterHandlers(server *tcp.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes([]tcp.Route{
		{
			Path:    1,
			Handler: gateway.HeatbeatHandler(serverCtx),
		},
		{
			Path:    2,
			Handler: gateway.OnlineUserHandler(serverCtx),
		},
		{
			Path:    3,
			Handler: gateway.OfflineUserHandler(serverCtx),
		},
	})
}
