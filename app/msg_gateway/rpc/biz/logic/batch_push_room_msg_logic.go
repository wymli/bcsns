package logic

import (
	"context"

	"github.com/wymli/bcsns/app/msg_gateway/rpc/pb"
	"github.com/wymli/bcsns/app/msg_gateway/svc"

	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
)

type BatchPushRoomMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBatchPushRoomMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchPushRoomMsgLogic {
	return &BatchPushRoomMsgLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchPushRoomMsgLogic) BatchPushRoomMsg(in *pb.BatchPushRoomMsgReq) (*pb.BatchPushRoomMsgResp, error) {
	return nil, errx.Wrapf(errx.ERROR_SERVER_UNIMPLEMENTED, "sorry, use push_room_msg instead")
}
