package handler

import (
	"github.com/wymli/bcsns/app/push_service/internal/handler/push"
	"github.com/wymli/bcsns/app/push_service/internal/svc"
	"github.com/wymli/bcsns/pkg/server_framework/mq"
)

func RegisterHandlers(svr *mq.Server, serverCtx *svc.ServiceContext) {
	svr.AddRoutes(
		[]mq.Route{
			{
				Topic:   "chat_user",
				Handler: push.ConsumeUserMessageHandler(serverCtx),
			},
			{
				Topic:   "chat_room",
				Handler: push.ConsumeRoomMessageHandler(serverCtx),
			},
			{
				Topic:   "moments",
				Handler: push.ConsumeMomentsHandler(serverCtx),
			},
		},
	)
}
