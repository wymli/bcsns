package logic

import (
	"context"
	"fmt"

	"github.com/wymli/bcsns/app/online_rpc/internal/svc"
	"github.com/wymli/bcsns/app/online_rpc/pb"

	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
)

type OfflineUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOfflineUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OfflineUserLogic {
	return &OfflineUserLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OfflineUserLogic) OfflineUser(in *pb.OfflineUserReq) (*pb.OfflineUserResp, error) {
	k := fmt.Sprintf(l.svcCtx.Config.MyRedis.Key.Online.Format, in.UserId)

	if err := l.svcCtx.RedisClient.Del(context.Background(), k).Err(); err != nil {
		return nil, errx.Wrapf(errx.ERROR_REDIS, "failed to offline user, err:%v", err)
	}

	return &pb.OfflineUserResp{}, nil
}
