package gateway

import (
	"github.com/wymli/bcsns/app/msg_gateway/svc"
	"github.com/wymli/bcsns/app/msg_gateway/tcp/logic/gateway"
	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/result"
	pb "github.com/wymli/bcsns/dependency/pb/tcp"
	"github.com/wymli/bcsns/common/server_framework/tcp"
	"google.golang.org/protobuf/proto"
)

func OfflineUserHandler(ctx *svc.ServiceContext) tcp.Handler {
	return func(connCtx *tcp.ConnCtx, body []byte) {
		var req *pb.OfflineUserReq
		if err := proto.Unmarshal(body, req); err != nil {
			err = errx.Wrapf(errx.ERROR_BAD_REQUEST, "failed to unmarshal pb, invalid argument, err:%v", err)
			result.TcpResult(connCtx, req, nil, err)
		}

		l := gateway.NewOfflineUserLogic(connCtx, ctx)
		resp, err := l.OfflineUser(req)

		result.TcpResult(connCtx, req, resp, err)
	}
}
