package config

import (
	"github.com/wymli/bcsns/common/logx"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	Logx  logx.Config
	Kafka KafkaConfig
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
