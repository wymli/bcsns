package gateway

import (
	"context"
	"fmt"
	"time"

	"github.com/wymli/bcsns/app/msg_gateway/svc"
	pb "github.com/wymli/bcsns/dependency/pb/tcp"
	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
	"github.com/wymli/bcsns/pkg/tcp"
)

type HeartbeatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHeartbeatLogic(ctx context.Context, svcCtx *svc.ServiceContext) HeartbeatLogic {
	return HeartbeatLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Heartbeat 用于客户端和服务端的健康检查,双向检查
func (l *HeartbeatLogic) Heartbeat(connCtx *tcp.ConnCtx, req *pb.HeartbeatReq) (resp *pb.CommonResp, err error) {
	userId, ok := connCtx.Ctx.Value("userId").(uint64)
	if !ok {
		return nil, errx.Wrap(errx.ERROR_USER_UNAUTHEN, "failed to extract userId from connCtx.ctx")
	}

	// refresh redis
	k := fmt.Sprintf(l.svcCtx.Config.Online.KeyFormat, userId)
	err = l.svcCtx.RedisClient.Expire(k, l.svcCtx.Config.Online.RefreshInterval*int(time.Minute))
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_REDIS, "failed to refresh expiration of key:%v, err:%v", k, err)
	}

	return &pb.CommonResp{
		Code: 0,
		Msg:  "OK",
	}, nil
}
