package svc

import (
	"github.com/wymli/bcsns/app/auth_api/internal/config"
	"github.com/wymli/bcsns/app/auth/rpc/auth"
	mylogx "github.com/wymli/bcsns/common/logx"
	zerologx "github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	AuthRpc auth.Auth
}

func NewServiceContext(c config.Config) *ServiceContext {
	zerologx.Disable()
	mylogx.Init(c.Logx)

	return &ServiceContext{
		Config:  c,
		AuthRpc: auth.NewAuth(zrpc.MustNewClient(c.AuthRpcConf)),
	}
}
