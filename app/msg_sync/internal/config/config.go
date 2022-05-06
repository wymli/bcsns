package config

import (
	"github.com/wymli/bcsns/common/config"
	connpool "github.com/wymli/bcsns/common/pool/conn"
)

type Config struct {
	config.NormalServerConfig
	Kafka      config.KafkaConfig     `yaml:"kafka,omitempty"`
	Cassandra  config.CassandraConfig `yaml:"cassandra,omitempty"`
	GatewayRpc connpool.Config        `yaml:"gateway_rpc,omitempty"`
	OnlineRpc  config.RpcClientConfig `yaml:"online_rpc,omitempty"`
	UserRpc    config.RpcClientConfig `yaml:"user_rpc,omitempty"`
	BCProxyRpc config.RpcClientConfig `yaml:"bc_proxy_rpc,omitempty"`

	Biz struct {
		Table struct {
			Message       string `yaml:"message,omitempty"`
			Moment        string `yaml:"moment,omitempty"`
			FailedMessage string `yaml:"failed_message,omitempty"`
		} `yaml:"table,omitempty"`

		Topic struct {
			ChatUser     string `yaml:"chat_user,omitempty"`
			ChatRoom     string `yaml:"chat_room,omitempty"`
			Moments      string `yaml:"moments,omitempty"`
			Notification string `yaml:"notification,omitempty"`
		} `yaml:"topic,omitempty"`
	} `yaml:"biz,omitempty"`
}
