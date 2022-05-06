package config

import (
	"github.com/wymli/bcsns/common/config"
)

type Config struct {
	config.RpcServerConfig `yaml:",inline"`
	Dgraph                 config.DgraphConfig    `yaml:"dgraph,omitempty"`
	AuthRpc                config.RpcClientConfig `yaml:"auth_rpc,omitempty"`
}
