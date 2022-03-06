package push

import (
	"context"

	"github.com/wymli/bcsns/app/push_service/internal/svc"

	"github.com/wymli/bcsns/common/logx"
)

type ConsumeMomentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConsumeMomentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) ConsumeMomentsLogic {
	return ConsumeMomentsLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConsumeMomentsLogic) ConsumeMoments(value []byte) error {
	return nil
}
