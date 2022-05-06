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

type BatchKeepAliveUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBatchKeepAliveUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchKeepAliveUserLogic {
	return &BatchKeepAliveUserLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//  is not online
func (l *BatchKeepAliveUserLogic) BatchKeepAliveUser(in *pb.BatchKeepAliveUserReq) (*pb.BatchKeepAliveUserResp, error) {
	_, err := l.svcCtx.RedisClient.Pipelined(l.ctx, func(p redis.Pipeliner) error {
		for _, userId := range in.UserIdList {
			k := fmt.Sprintf(l.svcCtx.Config.Biz.RedisKey.Online.Format, userId)
			p.Expire(l.ctx, k, time.Duration(l.svcCtx.Config.Biz.RedisKey.Online.Exp)*time.Second)
		}
		return nil
	})
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_REDIS, "failed to batch keepalive user, err:%v", err)
	}

	return &pb.BatchKeepAliveUserResp{}, nil
}
