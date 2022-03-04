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

type PostMomentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostMomentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) PostMomentsLogic {
	return PostMomentsLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostMomentsLogic) PostMoments(req types.PostMomentsReq) (resp *types.PostMomentsResp, err error) {
	// get uid from ctx
	userId, err := utils.ExtractUserIdFromCtx(l.ctx)
	if err != nil {
		return nil, errx.Wrap(errx.ERROR_SERVER_COMMON, "failed to get userId")
	}

	data := mqpb.Moments{
		ContentType: req.ContentType,
		MomentsId:   req.MomentsId,
		ServerTime:  time.Now().Unix(),
		ServerMsgId: "", // 是不是应该在push服务中插入mysql后再生成
		UserId:      userId,
		Data:        req.Data,
		IsTest:      req.IsTest,
	}

	pb, err := proto.Marshal(&data)
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to marshal pb of data: %#v, err:%v", data, err)
	}

	err = l.svcCtx.KafkaClient.WriteMessages(context.Background(), kafka.Message{
		Topic: l.svcCtx.Config.Kafka.Topic.ChatUser,
		Key:   []byte(req.MomentsId),
		Value: pb,
	})
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_MQ, "failed to write data to mq, err:%v", err)
	}

	return &types.PostMomentsResp{}, nil
}
