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

type GetMyFollowsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMyFollowsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyFollowsLogic {
	return &GetMyFollowsLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMyFollowsLogic) GetMyFollows(in *pb.GetMyFollowsReq) (*pb.GetMyFollowsResp, error) {
	uid := metadata.ExtractUserIdFromGRPC(l.ctx)
	if uid == 0 {
		return nil, errx.Wrapf(errx.ERROR_USER_UNAUTHEN, "can't get userid from ctx")
	}

	decode := model.FollowsDecode{}

	if err := model.Query(l.svcCtx.DgraphClient, model.QFollows, &decode, uid); err != nil {
		return nil, errx.Wrapf(err, "failed to get follows info")
	}

	if len(decode.All) == 0 {
		return nil, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to get follows info, query result is empty")
	}

	userInfos := []*pb.UserInfo{}
	for _, m := range decode.All {
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

	return &pb.GetMyFollowsResp{
		FollowsList: userInfos,
	}, nil
}
