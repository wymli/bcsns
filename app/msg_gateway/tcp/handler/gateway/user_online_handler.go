package gateway

import (
	"github.com/wymli/bcsns/app/msg_gateway/svc"
	"github.com/wymli/bcsns/app/msg_gateway/tcp/logic/gateway"
	pb "github.com/wymli/bcsns/dependency/pb/tcp"
	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/result"
	"github.com/wymli/bcsns/pkg/tcp"
	"google.golang.org/protobuf/proto"
)

func UserOnlineHandler(ctx *svc.ServiceContext) tcp.Handler {
	return func(connCtx *tcp.ConnCtx, body []byte) {
		var req *pb.UserOnlineReq
		if err := proto.Unmarshal(body, req); err != nil {
			err = errx.Wrapf(errx.ERROR_BAD_REQUEST, "failed to unmarshal pb, invalid argument, err:%v", err)
			result.TcpResult(connCtx.Conn, req, nil, err)
		}

		l := gateway.NewUserOnlineLogic(connCtx.Ctx, ctx)
		resp, err := l.UserOnline(connCtx, req)

		result.TcpResult(connCtx.Conn, req, resp, err)
	}
}