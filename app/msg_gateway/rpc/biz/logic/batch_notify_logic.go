package logic

import (
	"context"

	"github.com/wymli/bcsns/app/msg_gateway/rpc/pb"
	"github.com/wymli/bcsns/app/msg_gateway/svc"
	pbtcp "github.com/wymli/bcsns/dependency/pb/tcp"

	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
)

type BatchNotifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBatchNotifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchNotifyLogic {
	return &BatchNotifyLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchNotifyLogic) BatchNotify(in *pb.BatchNotifyReq) (*pb.BatchNotifyResp, error) {
	for _, notify := range in.NotifyList {
		conn, ok := l.svcCtx.UserConnPool.GetConn(notify.Uid)
		if !ok {
			return nil, errx.ERROR_GATEWAY_USER_NOT_FOUND
		}

		data := pbtcp.Notification{
			Type:         notify.Type,
			NotifyIdList: notify.NotifyIdList,
		}

		if err := conn.SendPBMsg(&data); err != nil {
			// sendpbmsg 内部有足够的错误处理
			return nil, errx.Wrapf(err, "failed to notificate msg to user")
		}
	}

	return &pb.BatchNotifyResp{}, nil
}
