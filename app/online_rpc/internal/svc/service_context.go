package svc

import (
	"github.com/go-redis/redis/v8"
	"github.com/wymli/bcsns/app/online_rpc/internal/config"

	mylogx "github.com/wymli/bcsns/common/logx"
	zerologx "github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	zerologx.Disable()
	mylogx.Init(c.MyLog)
	mylogx.Infof("config: %#v", c)

	return &ServiceContext{
		Config: c,
		RedisClient: redis.NewClient(&redis.Options{
			Addr: c.MyRedis.Host,
			// Password: c.MyRedis.Password,
		}),
	}
}
