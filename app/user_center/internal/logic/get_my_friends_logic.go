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

type GetMyFriendsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMyFriendsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyFriendsLogic {
	return &GetMyFriendsLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMyFriendsLogic) GetMyFriends(in *pb.GetMyFriendsReq) (*pb.GetMyFriendsResp, error) {
	uid := metadata.ExtractUserIdFromGRPC(l.ctx)
	if uid == 0 {
		return nil, errx.Wrapf(errx.ERROR_USER_UNAUTHEN, "can't get userid from ctx")
	}

	decode := model.FriendsDecode{}

	if err := model.Query(l.svcCtx.DgraphClient, model.QFriends, &decode, uid); err != nil {
		return nil, errx.Wrapf(err, "failed to get friends info")
	}

	if len(decode.All) == 0 {
		return nil, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to get friends info, query result is empty")
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

	return &pb.GetMyFriendsResp{
		FriendsList: userInfos,
	}, nil
}
