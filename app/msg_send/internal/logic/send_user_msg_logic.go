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

type SendUserMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendUserMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendUserMsgLogic {
	return &SendUserMsgLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendUserMsgLogic) SendUserMsg(in *pb.SendUserMsgReq) (*pb.SendUserMsgResp, error) {
	data := mqpb.UserChatMessage{
		ContentType: in.ContentType,
		SendMsgId:   in.SendMsgId,
		ServerMsgId: l.svcCtx.SnowflakeNode.Generate().Int64(),
		SendUid:     in.SendUserId,
		RecvUid:     in.RecvUserId,
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

	return &pb.SendUserMsgResp{
		ServerId: data.ServerMsgId,
	}, nil
}
