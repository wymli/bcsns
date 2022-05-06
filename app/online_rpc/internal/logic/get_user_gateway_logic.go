package logic

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/wymli/bcsns/app/online_rpc/internal/svc"
	"github.com/wymli/bcsns/app/online_rpc/pb"

	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
)

type GetUserGatewayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserGatewayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserGatewayLogic {
	return &GetUserGatewayLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserGatewayLogic) GetUserGateway(in *pb.GetUserGatewayReq) (*pb.GetUserGatewayResp, error) {
	k := fmt.Sprintf(l.svcCtx.Config.Biz.RedisKey.Online.Format, in.UserId)

	res := l.svcCtx.RedisClient.Get(l.ctx, k)
	switch res.Err() {
	case redis.Nil:
		return nil, errx.Wrapf(errx.ERROR_USER_OFFLINE, "user:%d is not online", in.UserId)
	case nil:
	default:
		return nil, errx.Wrapf(errx.ERROR_REDIS, "failed to get user gateway, err:%v", res.Err())
	}

	return &pb.GetUserGatewayResp{
		GatewayAddr: res.String(),
	}, nil
}
