package logic

import (
	"context"

	"github.com/wymli/bcsns/app/msg_gateway/svc"
	"github.com/wymli/bcsns/app/msg_gateway/rpc/pb"

	"github.com/wymli/bcsns/common/logx"
)

type PushMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPushMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushMsgLogic {
	return &PushMsgLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PushMsgLogic) PushMsg(in *pb.PushMsgReq) (*pb.PushMsgResp, error) {
	
	return &pb.PushMsgResp{}, nil
}
