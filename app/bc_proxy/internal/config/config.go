package config

import (
	"github.com/wymli/bcsns/common/config"
	"github.com/wymli/bcsns/common/logx"
)

type Config struct {
	config.RpcServerConfig
	BlockChain config.BlockChainConfig `yaml:"block_chain,omitempty"`
	Log        logx.Config             `yaml:"log,omitempty"`
	Kafka      config.KafkaConfig      `yaml:"kafka,omitempty"`

	Biz struct {
		Topic struct {
			ChatUser     string `yaml:"chat_user,omitempty"`
			ChatRoom     string `yaml:"chat_room,omitempty"`
			Moments      string `yaml:"moments,omitempty"`
			Notification string `yaml:"notification,omitempty"`
		} `yaml:"topic,omitempty"`
	} `yaml:"biz,omitempty"`
}
