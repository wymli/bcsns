package logic

import (
	"context"

	"github.com/wymli/bcsns/app/user_center/internal/model"
	"github.com/wymli/bcsns/app/user_center/internal/svc"
	"github.com/wymli/bcsns/app/user_center/pb"

	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
	"github.com/wymli/bcsns/common/tracer"
	"github.com/wymli/bcsns/common/utils"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	if l.svcCtx.Config.MockAll {
		l.Info().Msg("hello")
		_, span := tracer.Tracer().Start(l.ctx, "biz:get_user_info")
		defer span.End()

		return &pb.GetUserInfoResp{
			UserInfo: &pb.UserInfo{
				UserId:   9999,
				Nickname: "测试",
			},
		}, nil
	}

	if in.UserId == 0 {
		if f, ok := utils.ExtractUserIdFromCtx(l.ctx); ok {
			in.UserId = f
		} else {
			return nil, errx.Wrapf(errx.ERROR_USER_UNAUTHEN, "can't get userid from ctx")
		}
	}

	decode := model.UserInfoDecode{}

	if err := model.Query(l.svcCtx.DgraphClient, model.QUserInfo, &decode, in.UserId); err != nil {
		return nil, errx.Wrapf(err, "failed to get user info")
	}

	if len(decode.All) == 0 {
		return nil, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to get user info, query result is empty")
	}

	m := decode.All[0]

	return &pb.GetUserInfoResp{
		UserInfo: &pb.UserInfo{
			UserId:    utils.UintUid(m.Uid),
			Nickname:  m.Nickname,
			Sex:       int32(m.Sex),
			Age:       int32(m.Age),
			Avater:    m.Avatar,
			Address:   m.Address,
			PublicKey: m.PublicKey,
		},
	}, nil
}
