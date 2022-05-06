package sync

import (
	"context"

	pbbc "github.com/wymli/bcsns/app/bc_proxy/pb"
	"github.com/wymli/bcsns/app/msg_sync/internal/model"
	"github.com/wymli/bcsns/app/msg_sync/internal/svc"
	pbuser "github.com/wymli/bcsns/app/user_center/pb"
	pbmq "github.com/wymli/bcsns/dependency/pb/mq"

	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
)

type ConsumeMomentsLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConsumeMomentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) ConsumeMomentsLogic {
	return ConsumeMomentsLogic{
		log:    logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConsumeMomentsLogic) ConsumeMoments(req *pbmq.Moments) error {
	// 朋友圈采取拉取的策略,只推送红点

	// 1. query friends
	fansRes, err := l.svcCtx.UserRpc.GetMyFansUid(l.ctx, &pbuser.GetMyFansUidReq{})
	if err != nil {
		return errx.Wrap(err, "failed to get fans")
	}

	// 2. store in friends recv box
	moments := model.SyncMoments{}
	for _, uid := range fansRes.FansUidList {
		moments = append(moments, model.SyncMoment{
			Uid:         uid,
			SendUid:     req.UserId,
			ServerMsgId: req.ServerMomentsId,
			ContentType: req.ContentType,
			Data:        req.Data,
		})
	}
	if err := moments.BatchStore(l.svcCtx.CassandraClient, l.svcCtx.Config.Biz.Table.Moment); err != nil {
		return errx.Wrapf(err, "failed to store moments in db")
	}

	// 3. async on chain
	if _, err := l.svcCtx.BCProxyRpc.OnChainMoments(l.ctx, &pbbc.OnChainMomentsReq{
		ContentType:     req.ContentType,
		SendMomentsId:   req.SendMomentsId,
		ServerMomentsId: req.ServerMomentsId,
		UserId:          req.UserId,
		Data:            req.Data,
		IsTest:          req.IsTest,
	}); err != nil {
		return errx.Wrapf(err, "failed to persist moments in blockchain")
	}

	// 4. push red point
	// 略，让客户端轮询拉取？

	return nil
}
