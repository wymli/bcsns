package push

import (
	"context"

	"github.com/wymli/bcsns/app/push_service/internal/svc"

	"github.com/wymli/bcsns/common/logx"
)

type ConsumeUserMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConsumeUserMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) ConsumeUserMessageLogic {
	return ConsumeUserMessageLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConsumeUserMessageLogic) ConsumeUserMessage(value []byte) error {
	l.svcCtx.RedisClient.Get("online:%d",)
	return nil
}
