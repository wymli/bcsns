package logic

import (
	"context"

	"github.com/wymli/bcsns/app/bc_proxy/internal/svc"
	"github.com/wymli/bcsns/app/bc_proxy/pb"

	"github.com/wymli/bcsns/common/logx"
)

type PullUserMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPullUserMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PullUserMsgLogic {
	return &PullUserMsgLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PullUserMsgLogic) PullUserMsg(in *pb.PullUserMsgReq) (*pb.PullUserMsgResp, error) {
	// todo: add your logic here and delete this line
	l.Debug().Msg("in PullUserMsg")

	return &pb.PullUserMsgResp{}, nil
}
