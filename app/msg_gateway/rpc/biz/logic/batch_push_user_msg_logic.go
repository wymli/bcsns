package logic

import (
	"context"

	"github.com/wymli/bcsns/app/msg_gateway/rpc/pb"
	"github.com/wymli/bcsns/app/msg_gateway/svc"
	pbtcp "github.com/wymli/bcsns/dependency/pb/tcp"

	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
)

type BatchPushUserMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBatchPushUserMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchPushUserMsgLogic {
	return &BatchPushUserMsgLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchPushUserMsgLogic) BatchPushUserMsg(in *pb.BatchPushUserMsgReq) (*pb.BatchPushUserMsgResp, error) {
	errMap := map[uint64]error{}

	for _, msg := range in.PushUserMsgList {
		conn, ok := l.svcCtx.UserConnPool.GetConn(msg.RecvUid)
		if !ok {
			errMap[msg.RecvUid] = errx.ERROR_GATEWAY_USER_NOT_FOUND
			continue
		}

		data := pbtcp.ChatMsg{
			ContentType:  msg.ContentType,
			SendMsgId:    msg.SendMsgId,
			ServerMsgId:  msg.ServerMsgId,
			SendUid:      msg.SendUid,
			SendNickname: msg.SendNickname,
			SendAvatar:   msg.SendAvatar,
			RecvUid:      msg.RecvUid,
			RoomId:       nil,
			Data:         msg.Data,
			IsTest:       msg.IsTest,
		}

		if err := conn.SendPBMsg(&data); err != nil {
			// sendpbmsg 内部有足够的错误处理
			errMap[msg.RecvUid] = err
			continue
		}
	}

	if len(errMap) != 0 {
		l.Error().Msgf("failed to batch push user msg in gateway, err:%v", errMap)
		return nil, errx.ERROR_PART
	}
	// todo: 修改rsp,返回errMap, 让调用方重新投递

	return &pb.BatchPushUserMsgResp{}, nil
}
