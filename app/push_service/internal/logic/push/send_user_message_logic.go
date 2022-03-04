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
	return nil
	// data := mqpb.UserChatMessage{
	// 	ContentType: req.ContentType,
	// 	SendMsgId:   req.SendMsgId,
	// 	SendTime:    req.SendTime,
	// 	ServerTime:  time.Now().Unix(),
	// 	ServerMsgId: "", // 是不是应该在push服务中插入mysql后再生成
	// 	SendUid:     userId,
	// 	RecvUid:     req.RecvUserId,
	// 	Data:        req.Data,
	// 	IsTest:      req.IsTest,
	// }

	// pb, err := proto.Marshal(&data)
	// if err != nil {
	// 	return nil, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to marshal pb of data: %#v, err:%v", data, err)
	// }

	// err = l.svcCtx.KafkaClient.WriteMessages(context.Background(), kafka.Message{
	// 	Topic: l.svcCtx.Config.Kafka.Topic.ChatUser,
	// 	Key:   []byte(req.SendMsgId),
	// 	Value: pb,
	// })
	// if err != nil {
	// 	return nil, errx.Wrapf(errx.ERROR_MQ, "failed to write data to mq, err:%v", err)
	// }

	// return &types.SendUserMsgResp{}, nil
}
