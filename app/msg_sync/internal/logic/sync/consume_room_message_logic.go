package sync

import (
	"context"

	pbbc "github.com/wymli/bcsns/app/bc_proxy/pb"
	pbgw "github.com/wymli/bcsns/app/msg_gateway/rpc/pb"
	"github.com/wymli/bcsns/app/msg_sync/internal/model"
	"github.com/wymli/bcsns/app/msg_sync/internal/svc"
	pbonline "github.com/wymli/bcsns/app/online_rpc/pb"
	pbuser "github.com/wymli/bcsns/app/user_center/pb"
	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
	pbmq "github.com/wymli/bcsns/dependency/pb/mq"
)

type ConsumeRoomMessageLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConsumeRoomMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) ConsumeRoomMessageLogic {
	return ConsumeRoomMessageLogic{
		log:    logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConsumeRoomMessageLogic) ConsumeRoomMessage(req *pbmq.RoomChatMessage) error {
	// 1. get userids from roomid
	gInfo, err := l.svcCtx.UserRpc.GetGroupMemberUid(l.ctx, &pbuser.GetGroupMemberUidReq{
		GroupId: req.RoomId,
	})
	if err != nil {
		return errx.Wrapf(err, "failed to get member of group:%v", req.RoomId)
	}

	roomUids := []uint64{}
	for _, uid := range gInfo.UidList {
		roomUids = append(roomUids, uid)
	}

	// 2. store it in sync db
	sms := model.SyncMessages{}
	for _, uid := range roomUids {
		sms = append(sms, model.SyncMessage{
			Uid:         uid,
			SendUid:     req.SendUid,
			RoomId:      req.RoomId,
			ServerMsgId: req.ServerMsgId,
			SendMsgId:   req.SendMsgId,
			ContentType: req.ContentType,
			Data:        req.Data,
		})
	}

	if err := sms.BatchStore(l.svcCtx.CassandraClient, l.svcCtx.Config.Biz.Table.Message); err != nil {
		return errx.Wrapf(errx.ERROR_DB, "failed to store sync message, err:%v", err)
	}

	// 3. async on chain
	go func() {
		if _, err := l.svcCtx.BCProxyRpc.OnChainRoomMsg(l.ctx, &pbbc.OnChainRoomMsgReq{
			ContentType: req.ContentType,
			SendMsgId:   req.SendMsgId,
			ServerMsgId: req.ServerMsgId,
			SendUserId:  req.SendUid,
			RoomId:      req.RoomId,
			Data:        req.Data,
			IsTest:      req.IsTest,
		}); err != nil {
			l.log.Error().Msgf("faield to on chain group msg, err:%v", err)
		}
	}()

	// 4. push it to gateway
	olrsp, err := l.svcCtx.OnlineRpc.BatchGetUserGateway(l.ctx, &pbonline.BatchGetUserGatewayReq{
		UserIdList: roomUids,
	})
	if err != nil {
		return errx.Wrapf(err, "failed to batch get user gateway")
	}

	addr2uidMap := map[string][]uint64{}
	for i := range olrsp.GatewayAddrList {
		addr2uidMap[olrsp.GatewayAddrList[i]] = append(addr2uidMap[olrsp.GatewayAddrList[i]], roomUids[i])
	}

	for addr, uidList := range addr2uidMap {
		client, err := l.svcCtx.GatewayRpcPool.Get(addr)
		if err != nil {
			return errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to get rpc client from pool, err:%v", err)
		}

		_, err = client.(pbgw.GatewayClient).PushRoomMsg(l.ctx, &pbgw.PushRoomMsgReq{
			ContentType:  req.ContentType,
			SendMsgId:    req.SendMsgId,
			ServerMsgId:  req.ServerMsgId,
			SendUid:      req.SendUid,
			SendNickname: "",
			SendAvatar:   "",
			RoomId:       req.RoomId,
			RecvUidList:  uidList,
			Data:         req.Data,
			IsTest:       req.IsTest,
		})
		if err != nil {
			return errx.Wrap(err, "failed to push room msg to gateway")
		}
	}

	return nil
}
