package gw

import (
	"github.com/wymli/bcsns/pkg/tcp"
)

type GatewayServer struct {
	Server *tcp.Server
	// srvcCtx *svc.ServiceContext
}
