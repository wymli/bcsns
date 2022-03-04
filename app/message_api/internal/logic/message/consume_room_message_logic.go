package message

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/wymli/bcsns/app/message_api/internal/svc"
	"github.com/wymli/bcsns/app/message_api/internal/types"
	mqpb "github.com/wymli/bcsns/dependency/pb/mq"
	"google.golang.org/protobuf/proto"

	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
	"github.com/wymli/bcsns/common/utils"
)

type SendRoomMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendRoomMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) SendRoomMessageLogic {
	return SendRoomMessageLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendRoomMessageLogic) SendRoomMessage(req types.SendRoomMsgReq) (resp *types.SendRoomMsgResp, err error) {
	// get uid from ctx
	userId, err := utils.ExtractUserIdFromCtx(l.ctx)
	if err != nil {
		return nil, errx.Wrap(errx.ERROR_SERVER_COMMON, "failed to get userId")
	}

	data := mqpb.RoomChatMessage{
		ContentType: req.ContentType,
		SendMsgId:   req.SendMsgId,
		SendTime:    req.SendTime,
		ServerTime:  time.Now().Unix(),
		ServerMsgId: "", // 是不是应该在push服务中插入mysql后再生成
		SendUid:     userId,
		RoomId:      req.RoomId,
		Data:        req.Data,
		IsTest:      req.IsTest,
	}

	pb, err := proto.Marshal(&data)
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to marshal pb of data: %#v, err:%v", data, err)
	}

	err = l.svcCtx.KafkaClient.WriteMessages(context.Background(), kafka.Message{
		Topic: l.svcCtx.Config.Kafka.Topic.ChatRoom,
		Key:   []byte(req.SendMsgId),
		Value: pb,
	})
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_MQ, "failed to write data to mq, err:%v", err)
	}

	return &types.SendRoomMsgResp{}, nil
}
