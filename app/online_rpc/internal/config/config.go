package config

import (
	"github.com/wymli/bcsns/common/config"
)

type Config struct {
	config.RpcServerConfig
	Redis config.RedisConfig `yaml:"redis,omitempty"`

	Biz struct {
		RedisKey struct {
			Online struct {
				Pattern string `yaml:"pattern,omitempty"`
				Format  string `yaml:"format,omitempty"`
				Exp     int64  `yaml:"exp,omitempty"`
			} `yaml:"online,omitempty"`
		} `yaml:"redis_key,omitempty"`
	} `yaml:"biz,omitempty"`
}

// Key      struct {
// 	Online struct {

// 	} `yaml:"online,omitempty"`
// 	Deduplicate struct {
// 		Pattern string `yaml:"pattern,omitempty"`
// 		Format  string `yaml:"format,omitempty"`
// 		Exp     int64  `yaml:"exp,omitempty"`
// 	} `yaml:"deduplicate,omitempty"`
// } `yaml:"key,omitempty"`
