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

type JoinGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJoinGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinGroupLogic {
	return &JoinGroupLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JoinGroupLogic) JoinGroup(in *pb.JoinGroupReq) (*pb.JoinGroupResp, error) {
	uid := metadata.ExtractUserIdFromGRPC(l.ctx)
	if uid == 0 {
		return nil, errx.Wrapf(errx.ERROR_USER_UNAUTHEN, "can't get userid from ctx")
	}

	if err := model.JoinGroup(l.svcCtx.DgraphClient, in.GroupId, uid); err != nil {
		return nil, errx.Wrapf(err, "failed to join group")
	}

	return &pb.JoinGroupResp{}, nil
}
