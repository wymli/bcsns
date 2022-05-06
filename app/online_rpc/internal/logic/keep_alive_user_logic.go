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

type KeepAliveUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKeepAliveUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KeepAliveUserLogic {
	return &KeepAliveUserLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KeepAliveUserLogic) KeepAliveUser(in *pb.KeepAliveUserReq) (*pb.KeepAliveUserResp, error) {
	k := fmt.Sprintf(l.svcCtx.Config.Biz.RedisKey.Online.Format, in.UserId)

	if err := l.svcCtx.RedisClient.Expire(l.ctx, k,
		time.Duration(l.svcCtx.Config.Biz.RedisKey.Online.Exp)*time.Second).Err(); err != nil {
		return nil, errx.Wrapf(errx.ERROR_REDIS, "failed to keepalive user, err:%v", err)
	}

	return &pb.KeepAliveUserResp{}, nil
}
