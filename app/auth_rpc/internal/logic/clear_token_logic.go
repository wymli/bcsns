package logic

import (
	"context"

	"github.com/wymli/bcsns/app/auth_rpc/internal/svc"
	"github.com/wymli/bcsns/app/auth_rpc/pb"

	"github.com/wymli/bcsns/common/logx"
)

type ClearTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClearTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClearTokenLogic {
	return &ClearTokenLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//  清除token，只针对用户服务开放访问
func (l *ClearTokenLogic) ClearToken(in *pb.ClearTokenReq) (*pb.ClearTokenResp, error) {
	// todo: add your logic here and delete this line
	l.Debug().Msg("in ClearToken")

	return &pb.ClearTokenResp{}, nil
}
