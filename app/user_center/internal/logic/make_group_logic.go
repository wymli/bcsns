package logic

import (
	"context"

	"github.com/wymli/bcsns/app/user_center/internal/model"
	"github.com/wymli/bcsns/app/user_center/internal/svc"
	"github.com/wymli/bcsns/app/user_center/pb"

	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
	"github.com/wymli/bcsns/common/utils"
)

type MakeGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMakeGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MakeGroupLogic {
	return &MakeGroupLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MakeGroupLogic) MakeGroup(in *pb.MakeGroupReq) (*pb.MakeGroupResp, error) {
	members := []model.User{}
	for _, uid := range in.InvitedUids {
		members = append(members, model.User{Uid: utils.StringUid(uid)})
	}

	group := model.Group{
		GroupAvatar: in.Avatar,
		GroupName:   in.Name,
		Members:     members,
	}

	groupId, err := group.Store(l.svcCtx.DgraphClient)
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_DB, "failed to make new group, err:%v", err)
	}

	return &pb.MakeGroupResp{
		GroupId: groupId,
	}, nil
}
