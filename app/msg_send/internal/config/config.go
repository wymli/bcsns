package config

import (
	"github.com/wymli/bcsns/common/config"
)

type Config struct {
	config.RpcServerConfig `yaml:",inline"`
	Kafka                  config.KafkaConfig     `yaml:"kafka,omitempty"`
	Redis                  config.RedisConfig     `yaml:"redis,omitempty"`
	Snowflake              config.SnowflakeConfig `yaml:"snowflake,omitempty"`

	Biz struct {
		RedisKey struct {
			Online struct {
				Pattern string `yaml:"pattern,omitempty"`
				Format  string `yaml:"format,omitempty"`
				Exp     int64  `yaml:"exp,omitempty"`
			} `yaml:"online,omitempty"`
		} `yaml:"redis_key,omitempty"`

		Topic struct {
			ChatUser     string `yaml:"chat_user,omitempty"`
			ChatRoom     string `yaml:"chat_room,omitempty"`
			Moments      string `yaml:"moments,omitempty"`
			Notification string `yaml:"notification,omitempty"`
		} `yaml:"topic,omitempty"`
	} `yaml:"biz,omitempty"`
}
