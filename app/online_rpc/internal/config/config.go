package config

import (
	"github.com/wymli/bcsns/common/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	MyLog   logx.Config
	MyRedis RedisConf
}

type RedisConf struct {
	Host     string
	Type     string
	Password string
	Key      struct {
		Online struct {
			Pattern string
			Format string
			Exp    int64
		}
	}
}
