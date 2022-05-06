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

type GetGroupInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGroupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupInfoLogic {
	return &GetGroupInfoLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGroupInfoLogic) GetGroupInfo(in *pb.GetGroupInfoReq) (*pb.GetGroupInfoResp, error) {
	decode := model.GroupInfoDecode{}

	if err := model.Query(l.svcCtx.DgraphClient, model.QGroupInfo, &decode, in.GroupId); err != nil {
		return nil, errx.Wrapf(err, "failed to get group info")
	}

	if len(decode.All) == 0 {
		return nil, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to get group info, query result is empty")
	}

	group := decode.All[0]

	userInfos := []*pb.UserInfo{}
	for _, m := range group.Members {
		userInfos = append(userInfos, &pb.UserInfo{
			UserId:    utils.UintUid(m.Uid),
			Nickname:  m.Nickname,
			Sex:       int32(m.Sex),
			Age:       int32(m.Age),
			Avater:    m.Avatar,
			Address:   m.Address,
			PublicKey: m.PublicKey,
		})
	}

	return &pb.GetGroupInfoResp{
		GroupInfo: &pb.GroupInfo{
			GroupId: utils.UintUid(group.Uid),
			Avatar:  group.GroupAvatar,
			Name:    group.GroupName,
			Members: userInfos,
		},
	}, nil
}
