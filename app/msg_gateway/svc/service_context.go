package svc

import (
	"github.com/wymli/bcsns/app/auth_rpc/auth"
	"github.com/wymli/bcsns/app/msg_gateway/config"
	gw "github.com/wymli/bcsns/app/msg_gateway/tcp"

	mylogx "github.com/wymli/bcsns/common/logx"
	zerologx "github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	AuthClient   auth.Auth
	UserConnPool *gw.UserConnPool
	RedisClient  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	zerologx.Disable()
	mylogx.Init(c.Logx)

	return &ServiceContext{
		Config:       c,
		AuthClient:   auth.NewAuth(zrpc.MustNewClient(c.AuthRpcConfig)),
		UserConnPool: gw.NewUserConnPool(),
		RedisClient:  redis.NewRedis(c.Redis.Host, c.Redis.Type, c.Redis.Pass),
	}
}
