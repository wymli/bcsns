package logic

import (
	"context"
	"strconv"

	"github.com/wymli/bcsns/app/online_rpc/internal/svc"
	"github.com/wymli/bcsns/app/online_rpc/pb"

	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
)

type GetAllOnlineUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllOnlineUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllOnlineUserLogic {
	return &GetAllOnlineUserLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllOnlineUserLogic) GetAllOnlineUser(in *pb.GetAllOnlineUserReq) (*pb.GetAllOnlineUserResp, error) {
	pattern := l.svcCtx.Config.Biz.RedisKey.Online.Pattern
	res := l.svcCtx.RedisClient.Keys(l.ctx, pattern)
	if res.Err() != nil {
		return nil, errx.Wrapf(errx.ERROR_REDIS, "failed to get all online users, err:%v", res.Err())
	}

	logx.Infof("res.val: %#v", res.Val())

	prefix := pattern[:len(pattern)-1]
	userIds := make([]uint64, len(res.Val()))
	for i, userId := range res.Val() {
		uid, err := strconv.ParseUint(userId[len(prefix):], 10, 64)
		if err != nil {
			l.Err(err).Msg("")
		}
		userIds[i] = uid
	}

	return &pb.GetAllOnlineUserResp{
		UserIdList: userIds,
	}, nil
}
