package gateway

import (
	"context"

	"github.com/wymli/bcsns/app/msg_gateway/svc"
	"github.com/wymli/bcsns/app/online_rpc/online"
	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
	"github.com/wymli/bcsns/common/utils"
	pb "github.com/wymli/bcsns/dependency/pb/tcp"
	"github.com/wymli/bcsns/pkg/server_framework/tcp"
)

type OfflineUserLogic struct {
	logx.Logger
	connCtx *tcp.ConnCtx
	svcCtx  *svc.ServiceContext
}

func NewOfflineUserLogic(connCtx *tcp.ConnCtx, svcCtx *svc.ServiceContext) OfflineUserLogic {
	return OfflineUserLogic{
		Logger:  connCtx.Logger,
		connCtx: connCtx,
		svcCtx:  svcCtx,
	}
}

func (l *OfflineUserLogic) OfflineUser(req *pb.OfflineUserReq) (resp *pb.CommonResp, err error) {
	userId, ok := utils.ExtractUserIdFromCtx(l.connCtx.Ctx)
	if !ok {
		return nil, errx.Wrap(errx.ERROR_USER_UNAUTHEN, "failed to extract userId from connCtx.ctx")
	}

	err = l.svcCtx.UserConnPool.CloseConn(userId)
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to close connection of userId:%v, err:%v", userId, err)
	}

	_, err = l.svcCtx.OnlineRpc.OfflineUser(context.Background(), &online.OfflineUserReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errx.Wrapf(err, "failed to offline user")
	}

	return &pb.CommonResp{
		Code: 0,
		Msg:  "OK",
	}, nil
}
