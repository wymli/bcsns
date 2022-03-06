package logic

import (
	"context"
	"fmt"

	"github.com/wymli/bcsns/app/online_rpc/internal/svc"
	"github.com/wymli/bcsns/app/online_rpc/pb"

	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
)

type BatchOfflineUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBatchOfflineUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchOfflineUserLogic {
	return &BatchOfflineUserLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchOfflineUserLogic) BatchOfflineUser(in *pb.BatchOfflineUserReq) (*pb.BatchOfflineUserResp, error) {
	ks := make([]string, len(in.UserId))
	for _, userId := range in.UserId {
		ks = append(ks, fmt.Sprintf(l.svcCtx.Config.MyRedis.Key.Online.Format, userId))
	}

	if err := l.svcCtx.RedisClient.Del(context.Background(), ks...).Err(); err != nil {
		return nil, errx.Wrapf(errx.ERROR_REDIS, "failed to batch offline user, err:%v", err)
	}

	return &pb.BatchOfflineUserResp{}, nil
}
