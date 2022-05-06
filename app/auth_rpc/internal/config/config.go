package config

import (
	"github.com/wymli/bcsns/common/config"
)

type Config struct {
	config.RpcServerConfig `yaml:",inline"`
	JwtAuth                config.JwtConfig `yaml:"jwt_auth,omitempty"`
}
