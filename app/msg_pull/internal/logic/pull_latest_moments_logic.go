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

type PullLatestMomentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPullLatestMomentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PullLatestMomentsLogic {
	return &PullLatestMomentsLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PullLatestMomentsLogic) PullLatestMoments(in *pb.PullLatestMomentsReq) (*pb.PullLatestMomentsResp, error) {
	uid := metadata.ExtractUserIdFromGRPC(l.ctx)
	qMomentsList, err := model.QuerySyncMomentsLatest(l.svcCtx.CassandraClient, l.svcCtx.Config.Biz.Table.Moment, uid, in.LocalMomentId)
	if err != nil {
		return nil, errx.Wrap(err, "failed to query latest moments")
	}

	rMomentsList := make([]*pb.Moments, len(qMomentsList))
	for i, msg := range qMomentsList {
		rMomentsList[i] = &pb.Moments{
			ContentType: msg.ContentType,
			ServerMsgId: msg.ServerMsgId,
			SendUserId:  msg.SendUid,
			Data:        msg.Data,
		}
	}

	return &pb.PullLatestMomentsResp{
		MomentsList: rMomentsList,
	}, nil
}
