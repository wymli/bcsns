package config

import (
	"github.com/wymli/bcsns/app/seq_service/internal/seqmgr"
	"github.com/wymli/bcsns/common/config"
)

type Config struct {
	config.RpcServerConfig
	Seq seqmgr.SeqGenConfig
}
