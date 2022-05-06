package logic

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/wymli/bcsns/app/bc_proxy/internal/svc"
	"github.com/wymli/bcsns/common/logx"
	"github.com/wymli/bcsns/dependency/pb/mq"
	"github.com/wymli/bcsns/dependency/pb/notify"
	bcsns "github.com/wymli/bcsns/smart_contract"
	"google.golang.org/protobuf/proto"
)

type EventNotifyLogic struct {
	logx.Logger
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	notifyCh chan interface{}
}

func NewEventNotifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EventNotifyLogic {
	return &EventNotifyLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (ntf *EventNotifyLogic) Notify() {
	for e := range ntf.notifyCh {
		var mqdata *mq.Notification

		switch z := e.(type) {
		case *bcsns.BcsnsMomentsPersistedEvent:
			mqdata = &mq.Notification{
				Type:     notify.NotifyType_MOMENTS_ON_CHAIN,
				SendUid:  z.SendUid,
				RecvId:   nil,
				NotifyId: &notify.NotifyId{Id: &notify.NotifyId_MsgId{MsgId: z.ServerMsgId}},
			}
		case *bcsns.BcsnsUserMsgPersistedEvent:
			mqdata = &mq.Notification{
				Type:     notify.NotifyType_MSG_ON_CHAIN,
				SendUid:  z.SendUid,
				RecvId:   &notify.NotifyRecvId{RecvId: &notify.NotifyRecvId_UserId{UserId: z.RecvUid}},
				NotifyId: &notify.NotifyId{Id: &notify.NotifyId_MsgId{MsgId: z.ServerMsgId}},
			}
		case *bcsns.BcsnsRoomMsgPersistedEvent:
			mqdata = &mq.Notification{
				Type:     notify.NotifyType_MSG_ON_CHAIN,
				SendUid:  z.SendUid,
				RecvId:   &notify.NotifyRecvId{RecvId: &notify.NotifyRecvId_RoomId{RoomId: z.RoomUid}},
				NotifyId: &notify.NotifyId{Id: &notify.NotifyId_MsgId{MsgId: z.ServerMsgId}},
			}
		}

		pbuf, err := proto.Marshal(mqdata)
		if err != nil {
			ntf.Error().Msgf("failed to marshal ack-on-chain mq protobuf, err:%v", err)
			continue
		}

		if err := ntf.svcCtx.KafkaProducer.WriteMessages(ntf.ctx, kafka.Message{
			Topic: ntf.svcCtx.Config.Biz.Topic.Notification,
			Value: pbuf,
		}); err != nil {
			ntf.Error().Msgf("failed to produce ack-on-chain mq data to kafka, err:%v", err)
		}
	}
}
