package main

import (
	"github.com/wymli/bcsns/common/config"
)

type Config struct {
	config.RpcServerConfig `yaml:",inline"`
	UserCenterRpc          config.RpcClientConfig `yaml:"user_center_rpc,omitempty"`
	MsgSendRpc             config.RpcClientConfig `yaml:"msg_send_rpc,omitempty"`
	MsgPullRpc             config.RpcClientConfig `yaml:"msg_pull_rpc,omitempty"`
	BCPullRpc              config.RpcClientConfig `yaml:"bc_pull_rpc,omitempty"`
}
