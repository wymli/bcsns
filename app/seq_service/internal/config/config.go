package config

import (
	"github.com/wymli/bcsns/app/seq_service/internal/seqmgr"
	"github.com/wymli/bcsns/common/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Logx logx.Config
	Seq  seqmgr.SeqGenConfig
}
