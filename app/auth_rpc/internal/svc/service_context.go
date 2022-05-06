package svc

import (
	"github.com/wymli/bcsns/app/auth_rpc/internal/config"
	mylogx "github.com/wymli/bcsns/common/logx"
	zerologx "github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config config.Config
	// RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	zerologx.Disable()
	mylogx.Init(c.Log)

	return &ServiceContext{
		Config: c,
		// RedisClient: redis.NewRedis(c.Redis.Host, c.Redis.Type, c.Redis.Pass),
	}
}
