package logic

import (
	"context"

	"github.com/wymli/bcsns/app/seq_service/internal/svc"
	"github.com/wymli/bcsns/app/seq_service/pb"

	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
)

type GetSeqIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSeqIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSeqIdLogic {
	return &GetSeqIdLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSeqIdLogic) GetSeqId(in *pb.GetSeqIdReq) (*pb.GetSeqIdResp, error) {
	seqId, err := l.svcCtx.UserSeqMgr.GetNextId(in.UserId)
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_SEQID, "failed to generate next seq id, err:%v", err)
	}

	return &pb.GetSeqIdResp{
		SeqId: seqId,
	}, nil
}
