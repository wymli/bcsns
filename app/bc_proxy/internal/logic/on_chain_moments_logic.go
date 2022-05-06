package logic

import (
	"context"

	"github.com/wymli/bcsns/app/bc_proxy/internal/svc"
	"github.com/wymli/bcsns/app/bc_proxy/pb"
	pbbc "github.com/wymli/bcsns/dependency/pb/blockchain"
	"google.golang.org/protobuf/proto"

	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
)

type OnChainMomentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOnChainMomentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OnChainMomentsLogic {
	return &OnChainMomentsLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OnChainMomentsLogic) OnChainMoments(in *pb.OnChainMomentsReq) (*pb.OnChainMomentsResp, error) {
	auth, err := CopyBuildTransactionOpts(l.svcCtx.EtherClient, l.svcCtx.BCAuth)
	if err != nil {
		return nil, err
	}

	data := pbbc.Data{
		ContentType: in.ContentType,
		SendMsgId:   in.SendMomentsId,
		Data:        in.Data,
	}

	pbuf, err := proto.Marshal(&data)
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_MARSHALL, "failed to marshal blockchain data pb, err:%v", err)
	}

	_, err = l.svcCtx.BcsnsClient.PersistMoments(auth, in.UserId, in.ServerMomentsId, pbuf)
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_BC_TRANSACTION, "failed to persist moments, err:%v", err)
	}

	return &pb.OnChainMomentsResp{}, nil
}
