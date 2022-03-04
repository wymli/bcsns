package push

import (
	"context"

	"github.com/wymli/bcsns/app/push_service/internal/svc"
	"github.com/wymli/bcsns/app/push_service/internal/logic/push"
	"github.com/wymli/bcsns/pkg/mq"
)

func ConsumeRoomMessageHandler(ctx *svc.ServiceContext) mq.Handler {
	return func(value []byte) error {
		l := push.NewConsumeRoomMessageLogic(context.Background(), ctx)
		return l.ConsumeRoomMessage(value)
	}
}