package svc

import (
	pbauth "github.com/wymli/bcsns/app/auth_rpc/pb"
	"github.com/wymli/bcsns/app/msg_gateway/config"
	gw "github.com/wymli/bcsns/app/msg_gateway/tcp"
	pbonline "github.com/wymli/bcsns/app/online_rpc/pb"
	"google.golang.org/grpc"

	"github.com/wymli/bcsns/common/logx"
	"github.com/wymli/bcsns/common/server_framework/rpc"
)

type ServiceContext struct {
	Config       config.Config
	AuthRpc      pbauth.AuthClient
	UserConnPool *gw.UserConnPool
	OnlineRpc    pbonline.OnlineClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	logx.Init(c.Log)

	return &ServiceContext{
		Config:       c,
		AuthRpc:      pbauth.NewAuthClient(rpc.Must(grpc.Dial(c.AuthRpc.Endpoint))),
		UserConnPool: gw.NewUserConnPool(),
		OnlineRpc:    pbonline.NewOnlineClient(rpc.Must(grpc.Dial(c.OnlineRpc.Endpoint))),
	}
}
