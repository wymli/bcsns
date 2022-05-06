package sync

import (
	"context"

	"github.com/wymli/bcsns/app/msg_sync/internal/logic/sync"
	"github.com/wymli/bcsns/app/msg_sync/internal/svc"
	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/server_framework/mq"
	mqpb "github.com/wymli/bcsns/dependency/pb/mq"
	"google.golang.org/protobuf/proto"
)

func ConsumeUserMessageHandler(ctx *svc.ServiceContext) mq.Handler {
	return func(value []byte) error {
		data := mqpb.UserChatMessage{}
		if err := proto.Unmarshal(value, &data); err != nil {
			return errx.Wrapf(errx.ERROR_MARSHALL, "failed to consume user msg, err:%v", err)
		}

		l := sync.NewConsumeUserMessageLogic(context.Background(), ctx)
		return l.ConsumeUserMessage(&data)
	}
}
