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

func ConsumeRoomMessageHandler(ctx *svc.ServiceContext) mq.Handler {
	return func(value []byte) error {
		data := mqpb.RoomChatMessage{}
		if err := proto.Unmarshal(value, &data); err != nil {
			return errx.Wrapf(errx.ERROR_MARSHALL, "failed to consume room msg, err:%v", err)
		}

		l := sync.NewConsumeRoomMessageLogic(context.Background(), ctx)
		return l.ConsumeRoomMessage(&data)
	}
}
