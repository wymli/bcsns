package config

import (
	"github.com/wymli/bcsns/common/config"
)

type Config struct {
	config.RpcServerConfig
	Cassandra config.CassandraConfig `yaml:"cassandra,omitempty"`
	Biz       struct {
		Table struct {
			Message       string `yaml:"message,omitempty"`
			Moment        string `yaml:"moment,omitempty"`
			FailedMessage string `yaml:"failed_message,omitempty"`
		} `yaml:"table,omitempty"`
	} `yaml:"biz,omitempty"`
}
