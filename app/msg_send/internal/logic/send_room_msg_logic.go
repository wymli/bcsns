package logic

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
	"github.com/wymli/bcsns/app/msg_send/internal/svc"
	"github.com/wymli/bcsns/app/msg_send/pb"
	mqpb "github.com/wymli/bcsns/dependency/pb/mq"
	"google.golang.org/protobuf/proto"

	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
)

type SendRoomMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendRoomMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendRoomMsgLogic {
	return &SendRoomMsgLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendRoomMsgLogic) SendRoomMsg(in *pb.SendRoomMsgReq) (*pb.SendRoomMsgResp, error) {
	data := mqpb.RoomChatMessage{
		ContentType: in.ContentType,
		SendMsgId:   in.SendMsgId,
		ServerMsgId: l.svcCtx.SnowflakeNode.Generate().Int64(), // 是不是应该在push服务中插入mysql后再生成
		SendUid:     in.SendUserId,
		RoomId:      in.RoomId,
		Data:        in.Data,
		IsTest:      in.IsTest,
	}

	pbData, err := proto.Marshal(&data)
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to marshal pb of data: %#v, err:%v", data, err)
	}

	err = l.svcCtx.KafkaClient.WriteMessages(l.ctx, kafka.Message{
		Topic: l.svcCtx.Config.Biz.Topic.ChatRoom,
		Key:   []byte(fmt.Sprintf("%d", in.SendUserId)),
		Value: pbData,
	})
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_MQ, "failed to write data to mq, err:%v", err)
	}

	return &pb.SendRoomMsgResp{
		ServerId: data.ServerMsgId,
	}, nil
}
