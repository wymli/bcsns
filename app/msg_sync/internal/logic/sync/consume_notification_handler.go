package sync

import (
	"context"

	"github.com/wymli/bcsns/app/msg_gateway/rpc/pb"
	"github.com/wymli/bcsns/app/msg_sync/internal/svc"
	pbonline "github.com/wymli/bcsns/app/online_rpc/pb"
	pbuser "github.com/wymli/bcsns/app/user_center/pb"
	mqpb "github.com/wymli/bcsns/dependency/pb/mq"
	"github.com/wymli/bcsns/dependency/pb/notify"

	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
	"github.com/wymli/bcsns/common/metadata"
)

type ConsumeNotificationLogic struct {
	log    logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConsumeNotificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) ConsumeNotificationLogic {
	return ConsumeNotificationLogic{
		log:    logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConsumeNotificationLogic) ConsumeNotification(req *mqpb.Notification) error {
	switch req.Type {
	case notify.NotifyType_MSG_ON_CHAIN, notify.NotifyType_MOMENTS_ON_CHAIN:
		uidList := []uint64{}

		switch z := req.RecvId.RecvId.(type) {
		case *notify.NotifyRecvId_RoomId:
			res, err := l.svcCtx.UserRpc.GetGroupMemberUid(l.ctx, &pbuser.GetGroupMemberUidReq{
				GroupId: z.RoomId,
			})
			if err != nil {
				return errx.Wrapf(err, "failed to get group info of group:%v", z.RoomId)
			}
			for _, uid := range res.UidList {
				uidList = append(uidList, uid)
			}
		case *notify.NotifyRecvId_UserId:
			uidList = append(uidList, req.SendUid, z.UserId)
		case nil:
			uidList = append(uidList, req.SendUid)
		}

		return l.NotifyOnChian(req.Type, uidList, req.NotifyId.GetMsgId())

	case notify.NotifyType_MSG_READ:
		return l.NotifyMsgRead(req.Type, req.SendUid, req.NotifyId.GetMsgId())

	case notify.NotifyType_FOLLOWS_ONLINE:
		uid := req.NotifyId.GetUserId()
		ctx := metadata.CtxWithUserId(l.ctx, uid)

		rsp, err := l.svcCtx.UserRpc.GetMyFollows(ctx, &pbuser.GetMyFollowsReq{})
		if err != nil {
			return errx.Wrapf(err, "failed to get my follows")
		}

		followsList := []uint64{}
		for _, user := range rsp.FollowsList {
			followsList = append(followsList, user.UserId)
		}

		return l.NotifyOnline(req.Type, followsList, req.NotifyId.GetUserId())

	case notify.NotifyType_FRIENDS_ONLINE:
		uid := req.NotifyId.GetUserId()
		ctx := metadata.CtxWithUserId(l.ctx, uid)

		rsp, err := l.svcCtx.UserRpc.GetMyFriends(ctx, &pbuser.GetMyFriendsReq{})
		if err != nil {
			return errx.Wrapf(err, "failed to get my follows")
		}

		friendsList := []uint64{}
		for _, user := range rsp.FriendsList {
			friendsList = append(friendsList, user.UserId)
		}

		return l.NotifyOnline(req.Type, friendsList, req.NotifyId.GetUserId())

	default:
		return errx.ERROR_SERVER_UNIMPLEMENTED
	}
}

func (l *ConsumeNotificationLogic) NotifyOnChian(typ notify.NotifyType, uidList []uint64, msgId int64) error {
	olrsp, err := l.svcCtx.OnlineRpc.BatchGetUserGateway(l.ctx, &pbonline.BatchGetUserGatewayReq{
		UserIdList: uidList,
	})
	if err != nil {
		return errx.Wrap(err, "failed to batch get user online gateway")
	}

	addr2uidMap := map[string][]uint64{}
	for i := range olrsp.GatewayAddrList {
		addr2uidMap[olrsp.GatewayAddrList[i]] = append(addr2uidMap[olrsp.GatewayAddrList[i]], uidList[i])
	}

	for addr, uidList := range addr2uidMap {
		client, err := l.svcCtx.GatewayRpcPool.Get(addr)
		if err != nil {
			return errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to get rpc client from pool, err:%v", err)
		}

		notifyList := []*pb.BatchNotifyReqNotifyItem{}
		for _, uid := range uidList {
			notifyList = append(notifyList, &pb.BatchNotifyReqNotifyItem{
				Uid:  uid,
				Type: typ,
				NotifyIdList: []*notify.NotifyId{
					{Id: &notify.NotifyId_MsgId{MsgId: msgId}},
				},
			})
		}

		if _, err = client.(pb.GatewayClient).BatchNotify(l.ctx, &pb.BatchNotifyReq{
			NotifyList: notifyList,
		}); err != nil {
			return errx.Wrap(err, "failed to batch notify ack_on_chain to gateway")
		}
	}

	return nil
}

func (l *ConsumeNotificationLogic) NotifyOnline(typ notify.NotifyType, recvUidList []uint64, onlineUid uint64) error {
	return nil
}

func (l *ConsumeNotificationLogic) NotifyMsgRead(typ notify.NotifyType, uid uint64, msgId int64) error {
	return nil
}
