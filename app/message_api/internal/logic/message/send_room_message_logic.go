package message

import (
	"context"

	"github.com/wymli/bcsns/app/message_api/internal/svc"
	"github.com/wymli/bcsns/app/message_api/internal/types"
	"github.com/wymli/bcsns/app/message_rpc/pb"

	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
	"github.com/wymli/bcsns/common/utils"
)

type SendRoomMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendRoomMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) SendRoomMessageLogic {
	return SendRoomMessageLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendRoomMessageLogic) SendRoomMessage(req types.SendRoomMsgReq) (resp *types.SendRoomMsgResp, err error) {
	// get uid from ctx
	userId, err := utils.ExtractUserIdFromCtx(l.ctx)
	if err != nil {
		return nil, errx.Wrap(errx.ERROR_SERVER_COMMON, "failed to get userId")
	}

	_, err = l.svcCtx.MessageRpc.SendRoomMsg(l.ctx, &pb.SendRoomMsgReq{
		ContentType: req.ContentType,
		SendMsgId:   req.SendMsgId,
		SendUserId:  userId,
		RoomId:      req.RoomId,
		Data:        req.Data,
		IsTest:      req.IsTest,
	})
	if err != nil {
		return nil, err
	}

	return &types.SendRoomMsgResp{}, nil
}
