package logic

import (
	"context"

	"github.com/wymli/bcsns/app/msg_gateway/rpc/pb"
	"github.com/wymli/bcsns/app/msg_gateway/svc"
	pbtcp "github.com/wymli/bcsns/dependency/pb/tcp"

	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
)

type PushRoomMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPushRoomMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushRoomMsgLogic {
	return &PushRoomMsgLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PushRoomMsgLogic) PushRoomMsg(in *pb.PushRoomMsgReq) (*pb.PushRoomMsgResp, error) {
	errMap := map[uint64]error{}

	for _, recvUid := range in.RecvUidList {
		conn, ok := l.svcCtx.UserConnPool.GetConn(recvUid)
		if !ok {
			errMap[recvUid] = errx.ERROR_GATEWAY_USER_NOT_FOUND
			continue
		}

		data := pbtcp.ChatMsg{
			ContentType:  in.ContentType,
			SendMsgId:    in.SendMsgId,
			ServerMsgId:  in.ServerMsgId,
			SendUid:      in.SendUid,
			SendNickname: in.SendNickname,
			SendAvatar:   in.SendAvatar,
			RecvUid:      recvUid,
			RoomId:       &in.RoomId,
			Data:         in.Data,
			IsTest:       in.IsTest,
		}

		if err := conn.SendPBMsg(&data); err != nil {
			// sendpbmsg 内部有足够的错误处理
			errMap[recvUid] = err
			continue
		}
	}

	if len(errMap) != 0 {
		l.Error().Msgf("failed to batch push user msg in gateway, err:%v", errMap)
		return nil, errx.ERROR_PART
	}
	// todo: 修改rsp,返回errMap, 让调用方重新投递

	return &pb.PushRoomMsgResp{}, nil
}
