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

type GetMyFansUidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMyFansUidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyFansUidLogic {
	return &GetMyFansUidLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMyFansUidLogic) GetMyFansUid(in *pb.GetMyFansUidReq) (*pb.GetMyFansUidResp, error) {
	uid := metadata.ExtractUserIdFromGRPC(l.ctx)
	if uid == 0 {
		return nil, errx.Wrapf(errx.ERROR_USER_UNAUTHEN, "can't get userid from ctx")
	}

	decode := model.FansDecode{}

	if err := model.Query(l.svcCtx.DgraphClient, model.QFollows, &decode, uid); err != nil {
		return nil, errx.Wrapf(err, "failed to get fans uid info")
	}

	if len(decode.All) == 0 {
		return nil, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to get fans uid info, query result is empty")
	}

	uidList := []uint64{}
	for _, m := range decode.All {
		uidList = append(uidList, utils.UintUid(m.Uid))
	}

	return &pb.GetMyFansUidResp{
		FansUidList: uidList,
	}, nil
}
