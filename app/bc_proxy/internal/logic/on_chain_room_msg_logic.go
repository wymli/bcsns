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

type OnChainRoomMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOnChainRoomMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OnChainRoomMsgLogic {
	return &OnChainRoomMsgLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OnChainRoomMsgLogic) OnChainRoomMsg(in *pb.OnChainRoomMsgReq) (*pb.OnChainRoomMsgResp, error) {
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

	_, err = l.svcCtx.BcsnsClient.PersistRoomMessage(auth, in.RoomId, in.SendUserId, in.ServerMsgId, pbuf)
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_BC_TRANSACTION, "failed to persist room message, err:%v", err)
	}

	return &pb.OnChainRoomMsgResp{}, nil
}
