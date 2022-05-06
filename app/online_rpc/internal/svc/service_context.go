package svc

import (
	"github.com/go-redis/redis/v8"
	"github.com/wymli/bcsns/app/online_rpc/internal/config"

	"github.com/wymli/bcsns/common/logx"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	logx.Init(c.Log)

	return &ServiceContext{
		Config: c,
		RedisClient: redis.NewClient(&redis.Options{
			Addr: c.Redis.Host,
			// Password: c.MyRedis.Password,
		}),
	}
}
