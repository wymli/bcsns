package config

import (
	"github.com/wymli/bcsns/common/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
	}
	NoAuthUrls []string

	AuthRpcConf zrpc.RpcClientConf
	Logx        logx.Config
}
