package push

import (
	"context"

	"github.com/wymli/bcsns/app/push_service/internal/logic/push"
	"github.com/wymli/bcsns/app/push_service/internal/svc"
	"github.com/wymli/bcsns/pkg/server_framework/mq"
)

func ConsumeMomentsHandler(ctx *svc.ServiceContext) mq.Handler {
	return func(value []byte) error {
		l := push.NewConsumeMomentsLogic(context.Background(), ctx)
		return l.ConsumeMoments(value)
	}
}
