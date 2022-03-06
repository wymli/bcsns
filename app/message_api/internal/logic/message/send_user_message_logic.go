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

type SendUserMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendUserMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) SendUserMessageLogic {
	return SendUserMessageLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendUserMessageLogic) SendUserMessage(req types.SendUserMsgReq) (resp *types.SendUserMsgResp, err error) {
	// get uid from ctx
	userId, err := utils.ExtractUserIdFromCtx(l.ctx)
	if err != nil {
		return nil, errx.Wrap(errx.ERROR_SERVER_COMMON, "failed to get userId")
	}

	_, err = l.svcCtx.MessageRpc.SendUserMsg(l.ctx, &pb.SendUserMsgReq{
		ContentType: req.ContentType,
		SendMsgId:   req.SendMsgId,
		SendUserId:  userId,
		RecvUserId:  req.RecvUserId,
		Data:        req.Data,
		IsTest:      req.IsTest,
	})
	if err != nil {
		return nil, err
	}

	return &types.SendUserMsgResp{}, nil
}
