package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/wymli/bcsns/app/online_rpc/internal/svc"
	"github.com/wymli/bcsns/app/online_rpc/pb"

	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
)

type BatchOnlineUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBatchOnlineUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchOnlineUserLogic {
	return &BatchOnlineUserLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchOnlineUserLogic) BatchOnlineUser(in *pb.BatchOnlineUserReq) (*pb.BatchOnlineUserResp, error) {
	_, err := l.svcCtx.RedisClient.Pipelined(l.ctx, func(p redis.Pipeliner) error {
		for _, userId := range in.UserIdList {
			k := fmt.Sprintf(l.svcCtx.Config.Biz.RedisKey.Online.Format, userId)
			p.SetEX(l.ctx, k, in.GatewayAddr, time.Duration(l.svcCtx.Config.Biz.RedisKey.Online.Exp)*time.Second)
		}
		return nil
	})
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_REDIS, "failed to batch online user, err:%v", err)
	}

	return &pb.BatchOnlineUserResp{
		Exp: l.svcCtx.Config.Biz.RedisKey.Online.Exp,
	}, nil
}
