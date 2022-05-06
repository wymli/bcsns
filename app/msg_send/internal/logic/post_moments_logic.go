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

type PostMomentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostMomentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostMomentsLogic {
	return &PostMomentsLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostMomentsLogic) PostMoments(in *pb.PostMomentsReq) (*pb.PostMomentsResp, error) {
	data := mqpb.Moments{
		ContentType:     in.ContentType,
		SendMomentsId:   in.SendMomentsId,
		ServerMomentsId: l.svcCtx.SnowflakeNode.Generate().Int64(), // 是不是应该在push服务中插入mysql后再生成
		UserId:          in.UserId,
		Data:            in.Data,
		IsTest:          in.IsTest,
	}

	pbData, err := proto.Marshal(&data)
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to marshal pb of data: %#v, err:%v", data, err)
	}

	err = l.svcCtx.KafkaClient.WriteMessages(l.ctx, kafka.Message{
		Topic: l.svcCtx.Config.Biz.Topic.ChatUser,
		Key:   []byte(fmt.Sprintf("%d", in.UserId)),
		Value: pbData,
	})
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_MQ, "failed to write data to mq, err:%v", err)
	}

	return &pb.PostMomentsResp{
		ServerId: data.ServerMomentsId,
	}, nil
}
