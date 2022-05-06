package logic

import (
	"context"

	"github.com/wymli/bcsns/app/user_center/internal/svc"
	"github.com/wymli/bcsns/app/user_center/pb"

	"github.com/wymli/bcsns/common/logx"
)

type LeaveGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLeaveGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LeaveGroupLogic {
	return &LeaveGroupLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LeaveGroupLogic) LeaveGroup(in *pb.LeaveGroupReq) (*pb.LeaveGroupResp, error) {
	// todo: implemented
	return &pb.LeaveGroupResp{}, nil
}
