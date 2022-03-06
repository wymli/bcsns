package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/wymli/bcsns/app/online_rpc/internal/svc"
	"github.com/wymli/bcsns/app/online_rpc/pb"

	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
)

type OnlineUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOnlineUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OnlineUserLogic {
	return &OnlineUserLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OnlineUserLogic) OnlineUser(in *pb.OnlineUserReq) (*pb.OnlineUserResp, error) {
	k := fmt.Sprintf(l.svcCtx.Config.MyRedis.Key.Online.Format, in.UserId)

	err := l.svcCtx.RedisClient.SetEX(context.Background(), k, in.GatewayAddr, time.Duration(l.svcCtx.Config.MyRedis.Key.Online.Exp)*time.Second).Err()
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_REDIS, "failed to online user, err:%v", err)
	}

	return &pb.OnlineUserResp{
		Exp: l.svcCtx.Config.MyRedis.Key.Online.Exp,
	}, nil
}
