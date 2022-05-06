package logic

import (
	"context"

	"github.com/wymli/bcsns/app/user_center/internal/model"
	"github.com/wymli/bcsns/app/user_center/internal/svc"
	"github.com/wymli/bcsns/app/user_center/pb"

	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
	"github.com/wymli/bcsns/common/metadata"
	"github.com/wymli/bcsns/common/utils"
)

type GetMyGroupsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMyGroupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyGroupsLogic {
	return &GetMyGroupsLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMyGroupsLogic) GetMyGroups(in *pb.GetMyGroupsReq) (*pb.GetMyGroupsResp, error) {
	uid := metadata.ExtractUserIdFromGRPC(l.ctx)
	if uid == 0 {
		return nil, errx.Wrapf(errx.ERROR_USER_UNAUTHEN, "can't get userid from ctx")
	}

	decode := model.MyGroupsDecode{}

	if err := model.Query(l.svcCtx.DgraphClient, model.QMyGroups, &decode, uid); err != nil {
		return nil, errx.Wrapf(err, "failed to get my groups info")
	}

	if len(decode.All) == 0 {
		return nil, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to get my groups info, query result is empty")
	}

	groupsInfo := []*pb.GroupInfo{}
	for _, m := range decode.All {
		groupsInfo = append(groupsInfo, &pb.GroupInfo{
			GroupId: utils.UintUid(m.Uid),
			Avatar:  m.GroupAvatar,
			Name:    m.GroupName,
			Members: nil,
		})
	}

	return &pb.GetMyGroupsResp{
		GroupsList: groupsInfo,
	}, nil
}
