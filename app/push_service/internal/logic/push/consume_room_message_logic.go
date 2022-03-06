package push

import (
	"context"

	"github.com/wymli/bcsns/app/push_service/internal/svc"

	"github.com/wymli/bcsns/common/logx"
)

type ConsumeRoomMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConsumeRoomMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) ConsumeRoomMessageLogic {
	return ConsumeRoomMessageLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConsumeRoomMessageLogic) ConsumeRoomMessage(value []byte) error {
	return nil
	// // get uid from ctx
	// userId, err := utils.ExtractUserIdFromCtx(l.ctx)
	// if err != nil {
	// 	return nil, errx.Wrap(errx.ERROR_SERVER_COMMON, "failed to get userId")
	// }

	// data := mqpb.RoomChatMessage{
	// 	ContentType: req.ContentType,
	// 	SendMsgId:   req.SendMsgId,
	// 	SendTime:    req.SendTime,
	// 	ServerTime:  time.Now().Unix(),
	// 	ServerMsgId: "", // 是不是应该在push服务中插入mysql后再生成
	// 	SendUid:     userId,
	// 	RoomId:      req.RoomId,
	// 	Data:        req.Data,
	// 	IsTest:      req.IsTest,
	// }

	// pb, err := proto.Marshal(&data)
	// if err != nil {
	// 	return nil, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to marshal pb of data: %#v, err:%v", data, err)
	// }

	// err = l.svcCtx.KafkaClient.WriteMessages(context.Background(), kafka.Message{
	// 	Topic: l.svcCtx.Config.Kafka.Topic.ChatRoom,
	// 	Key:   []byte(req.SendMsgId),
	// 	Value: pb,
	// })
	// if err != nil {
	// 	return nil, errx.Wrapf(errx.ERROR_MQ, "failed to write data to mq, err:%v", err)
	// }

	// return &types.SendRoomMsgResp{}, nil
}
