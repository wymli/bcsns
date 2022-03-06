package svc

import (
	"github.com/wymli/bcsns/app/auth_rpc/auth"
	"github.com/wymli/bcsns/app/msg_gateway/config"
	gw "github.com/wymli/bcsns/app/msg_gateway/tcp"
	"github.com/wymli/bcsns/app/online_rpc/online"

	mylogx "github.com/wymli/bcsns/common/logx"
	zerologx "github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	AuthRpc      auth.Auth
	UserConnPool *gw.UserConnPool
	OnlineRpc    online.Online
}

func NewServiceContext(c config.Config) *ServiceContext {
	zerologx.Disable()
	mylogx.Init(c.Logx)

	return &ServiceContext{
		Config:       c,
		AuthRpc:      auth.NewAuth(zrpc.MustNewClient(c.AuthRpcConfig)),
		UserConnPool: gw.NewUserConnPool(),
		OnlineRpc: online.NewOnline(zrpc.MustNewClient(c.OnlineRpcConfig)),
	}
}
