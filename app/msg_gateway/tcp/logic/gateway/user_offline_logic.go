package gateway

import (
	"context"

	"github.com/wymli/bcsns/app/msg_gateway/svc"
	pb "github.com/wymli/bcsns/dependency/pb/tcp"
	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
	"github.com/wymli/bcsns/pkg/tcp"
)

type UserOfflineLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserOfflineLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserOfflineLogic {
	return UserOfflineLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserOfflineLogic) UserOffline(connCtx *tcp.ConnCtx, req *pb.UserOfflineReq) (resp *pb.CommonResp, err error) {
	userId, ok := connCtx.Ctx.Value("userId").(uint64)
	if !ok {
		return nil, errx.Wrap(errx.ERROR_USER_UNAUTHEN, "failed to extract userId from connCtx.ctx")
	}

	err = l.svcCtx.UserConnPool.CloseConn(userId)
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to close connection of userId:%v, err:%v", userId, err)
	}

	return &pb.CommonResp{
		Code: 0,
		Msg:  "OK",
	}, nil
}
