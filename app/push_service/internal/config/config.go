package config

import (
	"github.com/wymli/bcsns/common/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Config struct {
	Redis redis.RedisKeyConf

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
	GroupId string
}
