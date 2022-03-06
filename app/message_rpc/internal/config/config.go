package config

import (
	"github.com/wymli/bcsns/common/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Logx                logx.Config
	Kafka               KafkaConfig
	SeqRpcConfig        zrpc.RpcClientConf
	SnowflakeNodeNumber int64
}

type KafkaConfig struct {
	Broker struct {
		Endpoints []string
	}
	Topic struct {
		ChatUser string
		ChatRoom string
		Moments  string
	}
}
