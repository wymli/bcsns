package config

import (
	"github.com/wymli/bcsns/common/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Logx          logx.Config
	AuthRpcConfig zrpc.RpcClientConf
	Deploy        struct {
		Use string `json=",options=k8s|docker|docker-compose"`
	}
	Online struct{
		RefreshInterval int 
		KeyFormat string
	}
}
