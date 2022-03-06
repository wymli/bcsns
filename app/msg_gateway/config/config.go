package config

import (
	"github.com/wymli/bcsns/common/logx"
	"github.com/wymli/bcsns/pkg/server_framework/tcp"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Logx            logx.Config
	AuthRpcConfig   zrpc.RpcClientConf
	OnlineRpcConfig zrpc.RpcClientConf
	Deploy          struct {
		Use string `json=",options=k8s|docker|docker-compose"`
	}
	TcpConfig tcp.Config
}
