package handler

import (
	"github.com/wymli/bcsns/app/msg_sync/internal/handler/sync"
	"github.com/wymli/bcsns/app/msg_sync/internal/svc"
	"github.com/wymli/bcsns/common/server_framework/mq"
)

func RegisterHandlers(svr *mq.Server, serverCtx *svc.ServiceContext) {
	svr.AddRoutes(
		[]mq.Route{
			{
				Topic:   serverCtx.Config.Biz.Topic.ChatUser,
				Handler: sync.ConsumeUserMessageHandler(serverCtx),
			},
			{
				Topic:   serverCtx.Config.Biz.Topic.ChatRoom,
				Handler: sync.ConsumeRoomMessageHandler(serverCtx),
			},
			{
				Topic:   serverCtx.Config.Biz.Topic.Moments,
				Handler: sync.ConsumeMomentsHandler(serverCtx),
			},
			{
				Topic:   serverCtx.Config.Biz.Topic.Notification,
				Handler: sync.ConsumeNotificationHandler(serverCtx),
			},
		},
	)
}
