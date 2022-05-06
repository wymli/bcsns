package gateway

import (
	"context"

	"github.com/wymli/bcsns/app/msg_gateway/svc"
	pbonline "github.com/wymli/bcsns/app/online_rpc/pb"
	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
	"github.com/wymli/bcsns/common/utils"
	pb "github.com/wymli/bcsns/dependency/pb/tcp"
	"github.com/wymli/bcsns/common/server_framework/tcp"
)

type HeartbeatLogic struct {
	logx.Logger
	connCtx *tcp.ConnCtx
	svcCtx  *svc.ServiceContext
}

func NewHeartbeatLogic(connCtx *tcp.ConnCtx, svcCtx *svc.ServiceContext) HeartbeatLogic {
	return HeartbeatLogic{
		Logger:  connCtx.Logger,
		connCtx: connCtx,
		svcCtx:  svcCtx,
	}
}

// Heartbeat 用于客户端和服务端的健康检查,双向检查
func (l *HeartbeatLogic) Heartbeat(req *pb.HeartbeatReq) (resp *pb.CommonResp, err error) {
	userId, ok := utils.ExtractUserIdFromCtx(l.connCtx.Ctx)
	if !ok {
		return nil, errx.Wrap(errx.ERROR_USER_UNAUTHEN, "failed to extract userId from connCtx.ctx")
	}

	// keep alive user online status
	_, err = l.svcCtx.OnlineRpc.KeepAliveUser(context.Background(), &pbonline.KeepAliveUserReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errx.Wrapf(err, "failed to keep alive user")
	}

	return &pb.CommonResp{
		Code: 0,
		Msg:  "OK",
	}, nil
}
