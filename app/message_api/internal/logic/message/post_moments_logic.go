package message

import (
	"context"

	"github.com/wymli/bcsns/app/message_api/internal/svc"
	"github.com/wymli/bcsns/app/message_api/internal/types"
	"github.com/wymli/bcsns/app/message_rpc/pb"

	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
	"github.com/wymli/bcsns/common/utils"
)

type PostMomentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostMomentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) PostMomentsLogic {
	return PostMomentsLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostMomentsLogic) PostMoments(req types.PostMomentsReq) (resp *types.PostMomentsResp, err error) {
	// get uid from ctx
	userId, err := utils.ExtractUserIdFromCtx(l.ctx)
	if err != nil {
		return nil, errx.Wrap(errx.ERROR_SERVER_COMMON, "failed to get userId")
	}

	_, err = l.svcCtx.MessageRpc.PostMoments(l.ctx, &pb.PostMomentsReq{
		ContentType:   req.ContentType,
		SendMomentsId: req.SendMomentsId,
		UserId:        userId,
		Data:          req.Data,
		IsTest:        req.IsTest,
	})
	if err != nil {
		return nil, err
	}

	return &types.PostMomentsResp{}, nil
}
