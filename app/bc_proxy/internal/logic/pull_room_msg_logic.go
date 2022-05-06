package logic

import (
	"context"

	"github.com/wymli/bcsns/app/bc_proxy/internal/svc"
	"github.com/wymli/bcsns/app/bc_proxy/pb"

	"github.com/wymli/bcsns/common/logx"
)

type PullRoomMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPullRoomMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PullRoomMsgLogic {
	return &PullRoomMsgLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PullRoomMsgLogic) PullRoomMsg(in *pb.PullRoomMsgReq) (*pb.PullRoomMsgResp, error) {
	// todo: add your logic here and delete this line
	l.Debug().Msg("in PullRoomMsg")

	return &pb.PullRoomMsgResp{}, nil
}
