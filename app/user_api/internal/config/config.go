package config

import (
	"github.com/wymli/bcsns/common/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Dgraph      DgraphConfig
	Logx        logx.Config
	AuthRpcConf zrpc.RpcClientConf
}

type DgraphConfig struct {
	Endpoint string
}
