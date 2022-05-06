package logic

import (
	"context"

	"github.com/wymli/bcsns/app/msg_pull/internal/model"
	"github.com/wymli/bcsns/app/msg_pull/internal/svc"
	"github.com/wymli/bcsns/app/msg_pull/pb"

	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
	"github.com/wymli/bcsns/common/metadata"
)

type PullLatestMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPullLatestMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PullLatestMsgLogic {
	return &PullLatestMsgLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PullLatestMsgLogic) PullLatestMsg(in *pb.PullLatestMsgReq) (*pb.PullLatestMsgResp, error) {
	uid := metadata.ExtractUserIdFromGRPC(l.ctx)
	qMsgList, err := model.QuerySyncMsgLatest(l.svcCtx.CassandraClient, l.svcCtx.Config.Biz.Table.Message, uid, in.LocalMsgId)
	if err != nil {
		return nil, errx.Wrap(err, "failed to query latest msg")
	}

	// todo: 查询失败消息表

	rMsgList := make([]*pb.Message, len(qMsgList))
	for i, msg := range qMsgList {
		rMsgList[i] = &pb.Message{
			ContentType: msg.ContentType,
			ServerMsgId: msg.ServerMsgId,
			SendUserId:  msg.SendUid,
			RoomId:      msg.RoomId,
			Data:        msg.Data,
		}
	}

	return &pb.PullLatestMsgResp{
		MsgList: rMsgList,
	}, nil
}
