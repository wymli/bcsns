package config

import (
	"github.com/wymli/bcsns/common/config"
	"github.com/wymli/bcsns/common/server_framework/tcp"
)

type Config struct {
	config.RpcServerConfig
	AuthRpc   config.RpcClientConfig `yaml:"auth_rpc,omitempty"`
	OnlineRpc config.RpcClientConfig `yaml:"online_rpc,omitempty"`
	Deploy    config.DeployConfig    `yaml:"deploy,omitempty"`
	TcpConfig tcp.Config             `yaml:"tcp_config,omitempty"`
}
