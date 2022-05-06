package logic

import (
	"context"

	"github.com/wymli/bcsns/app/user_center/internal/model"
	"github.com/wymli/bcsns/app/user_center/internal/svc"
	"github.com/wymli/bcsns/app/user_center/pb"

	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
	"github.com/wymli/bcsns/common/metadata"
)

type FollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowLogic) Follow(in *pb.FollowReq) (*pb.FollowResp, error) {
	uid := metadata.ExtractUserIdFromGRPC(l.ctx)
	if uid == 0 {
		return nil, errx.Wrapf(errx.ERROR_USER_UNAUTHEN, "can't get userid from ctx")
	}

	if err := model.Follow(l.svcCtx.DgraphClient, uid, in.To); err != nil {
		return nil, errx.Wrapf(err, "failed to follow user")
	}

	return &pb.FollowResp{}, nil
}
