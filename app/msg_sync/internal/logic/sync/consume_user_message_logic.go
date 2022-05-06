package sync

import (
	"context"

	pbbc "github.com/wymli/bcsns/app/bc_proxy/pb"
	pbgw "github.com/wymli/bcsns/app/msg_gateway/rpc/pb"
	"github.com/wymli/bcsns/app/msg_sync/internal/model"
	"github.com/wymli/bcsns/app/msg_sync/internal/svc"
	pbonline "github.com/wymli/bcsns/app/online_rpc/pb"
	pbmq "github.com/wymli/bcsns/dependency/pb/mq"

	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
)

type ConsumeUserMessageLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConsumeUserMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) ConsumeUserMessageLogic {
	return ConsumeUserMessageLogic{
		log:    logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConsumeUserMessageLogic) ConsumeUserMessage(req *pbmq.UserChatMessage) error {
	// 1. store it in sync db
	sm := model.SyncMessage{
		Uid:         req.RecvUid,
		SendUid:     req.SendUid,
		RoomId:      0,
		ServerMsgId: req.ServerMsgId,
		SendMsgId:   req.SendMsgId,
		ContentType: req.ContentType,
		Data:        req.Data,
	}

	if err := sm.Store(l.svcCtx.CassandraClient, l.svcCtx.Config.Biz.Table.Message); err != nil {
		return errx.Wrapf(errx.ERROR_DB, "failed to store sync message, err:%v", err)
	}

	// 2. async on chain
	go func() {
		if _, err := l.svcCtx.BCProxyRpc.OnChainUserMsg(l.ctx, &pbbc.OnChainUserMsgReq{
			ContentType: req.ContentType,
			SendMsgId:   req.SendMsgId,
			ServerMsgId: req.ServerMsgId,
			SendUserId:  req.SendUid,
			RecvUserId:  req.RecvUid,
			Data:        req.Data,
			IsTest:      req.IsTest,
		}); err != nil {
			l.log.Error().Msgf("faield to on chain user msg, err:%v", err)
		}
	}()

	// 3. push it to gateway
	olrsp, err := l.svcCtx.OnlineRpc.GetUserGateway(l.ctx, &pbonline.GetUserGatewayReq{
		UserId: req.RecvUid,
	})
	if err != nil {
		return errx.Wrapf(err, "failed to get user gateway")
	}

	client, err := l.svcCtx.GatewayRpcPool.Get(olrsp.GatewayAddr)
	if err != nil {
		return errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to get rpc client from pool, err:%v", err)
	}

	_, err = client.(pbgw.GatewayClient).PushUserMsg(l.ctx, &pbgw.PushUserMsgReq{
		ContentType:  req.ContentType,
		SendMsgId:    req.SendMsgId,
		ServerMsgId:  req.ServerMsgId,
		SendUid:      req.SendUid,
		SendNickname: "",
		SendAvatar:   "",
		RecvUid:      req.RecvUid,
		Data:         req.Data,
		IsTest:       req.IsTest,
	})
	if err != nil {
		return errx.Wrap(err, "failed to push user msg to gateway")
	}

	if errx.Is(err, errx.ERROR_GATEWAY_USER_NOT_FOUND) {
		// 用户不在该gateway上, 在线服务状态出错
	}

	// todo: 重试, 又或者没有重试的必要? 只要出错就认为用户不在线即可

	return nil
}
