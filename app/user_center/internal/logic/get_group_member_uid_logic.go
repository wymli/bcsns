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

type GetGroupMemberUidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGroupMemberUidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupMemberUidLogic {
	return &GetGroupMemberUidLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGroupMemberUidLogic) GetGroupMemberUid(in *pb.GetGroupMemberUidReq) (*pb.GetGroupMemberUidResp, error) {
	decode := model.GroupMemberUidDecode{}

	if err := model.Query(l.svcCtx.DgraphClient, model.QGroupMemberUid, &decode, in.GroupId); err != nil {
		return nil, errx.Wrapf(err, "failed to get group member uid")
	}

	if len(decode.All) == 0 {
		return nil, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to get group member uids, query result is empty")
	}

	group := decode.All[0]

	uidList := []uint64{}
	for _, m := range group.Members {
		uidList = append(uidList, utils.UintUid(m.Uid))
	}

	return &pb.GetGroupMemberUidResp{
		UidList: uidList,
	}, nil
}
