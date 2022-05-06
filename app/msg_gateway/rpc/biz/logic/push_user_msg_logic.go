package logic

import (
	"context"

	"github.com/wymli/bcsns/app/msg_gateway/rpc/pb"
	"github.com/wymli/bcsns/app/msg_gateway/svc"
	pbtcp "github.com/wymli/bcsns/dependency/pb/tcp"

	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
)

type PushUserMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPushUserMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushUserMsgLogic {
	return &PushUserMsgLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PushUserMsgLogic) PushUserMsg(in *pb.PushUserMsgReq) (*pb.PushUserMsgResp, error) {
	conn, ok := l.svcCtx.UserConnPool.GetConn(in.RecvUid)
	if !ok {
		return nil, errx.Wrapf(errx.ERROR_GATEWAY_USER_NOT_FOUND, "user not found in this gateway")
	}

	data := pbtcp.ChatMsg{
		ContentType:  in.ContentType,
		SendMsgId:    in.SendMsgId,
		ServerMsgId:  in.ServerMsgId,
		SendUid:      in.SendUid,
		SendNickname: in.SendNickname,
		SendAvatar:   in.SendAvatar,
		RecvUid:      in.RecvUid,
		RoomId:       nil,
		Data:         in.Data,
		IsTest:       in.IsTest,
	}

	if err := conn.SendPBMsg(&data); err != nil {
		// sendpbmsg 内部有足够的错误处理
		return nil, errx.Wrapf(err, "failed to push user msg")
	}

	return &pb.PushUserMsgResp{}, nil
}
