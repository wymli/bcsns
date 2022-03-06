package svc

import (
	"github.com/wymli/bcsns/app/message_api/internal/config"
	"github.com/wymli/bcsns/app/message_rpc/message"
	mylogx "github.com/wymli/bcsns/common/logx"
	zerologx "github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	MessageRpc message.Message
}

func NewServiceContext(c config.Config) *ServiceContext {
	zerologx.Disable()
	mylogx.Init(c.Logx)

	return &ServiceContext{
		Config:     c,
		MessageRpc: message.NewMessage(zrpc.MustNewClient(c.MessageRpcConf)),
	}
}
