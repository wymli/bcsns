package logic

import (
	"context"

	"github.com/wymli/bcsns/app/bc_proxy/internal/svc"
	"github.com/wymli/bcsns/app/bc_proxy/pb"
	"google.golang.org/protobuf/proto"

	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
	pbbc "github.com/wymli/bcsns/dependency/pb/blockchain"
)

type OnChainUserMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOnChainUserMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OnChainUserMsgLogic {
	return &OnChainUserMsgLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OnChainUserMsgLogic) OnChainUserMsg(in *pb.OnChainUserMsgReq) (*pb.OnChainUserMsgResp, error) {
	auth, err := CopyBuildTransactionOpts(l.svcCtx.EtherClient, l.svcCtx.BCAuth)
	if err != nil {
		return nil, err
	}

	data := pbbc.Data{
		ContentType: in.ContentType,
		SendMsgId:   in.SendMsgId,
		Data:        in.Data,
	}

	pbuf, err := proto.Marshal(&data)
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_MARSHALL, "failed to marshal blockchain data pb, err:%v", err)
	}

	_, err = l.svcCtx.BcsnsClient.PersistUserMessage(auth, in.SendUserId, in.RecvUserId, in.ServerMsgId, pbuf)
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_BC_TRANSACTION, "failed to persist user message, err:%v", err)
	}

	return &pb.OnChainUserMsgResp{}, nil
}
