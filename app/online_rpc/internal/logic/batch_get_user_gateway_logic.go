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

type BatchGetUserGatewayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBatchGetUserGatewayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchGetUserGatewayLogic {
	return &BatchGetUserGatewayLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchGetUserGatewayLogic) BatchGetUserGateway(in *pb.BatchGetUserGatewayReq) (*pb.BatchGetUserGatewayResp, error) {
	res, err := l.svcCtx.RedisClient.Pipelined(context.Background(), func(p redis.Pipeliner) error {
		for _, userId := range in.UserId {
			k := fmt.Sprintf(l.svcCtx.Config.MyRedis.Key.Online.Format, userId)
			p.Get(context.Background(), k)
		}
		return nil
	})

	switch err {
	case nil:
	case redis.Nil:
	default:
		return nil, errx.Wrapf(errx.ERROR_REDIS, "failed to batch get user gateway, err:%v", err)
	}

	addrs := make([]string, len(in.UserId))
	for i, cmd := range res {
		if cmd.Err() == redis.Nil {
			addrs[i] = ""
		} else {
			addrs[i] = cmd.(*redis.StringCmd).Val()
		}
	}

	return &pb.BatchGetUserGatewayResp{
		GatewayAddr: addrs,
	}, nil
}
