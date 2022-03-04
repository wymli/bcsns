package svc

import (
	"github.com/segmentio/kafka-go"
	"github.com/wymli/bcsns/app/message_api/internal/config"
	mylogx "github.com/wymli/bcsns/common/logx"
	zerologx "github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config      config.Config
	KafkaClient *kafka.Writer
}

func NewServiceContext(c config.Config) *ServiceContext {
	zerologx.Disable()
	mylogx.Init(c.Logx)

	return &ServiceContext{
		Config: c,
		KafkaClient: &kafka.Writer{
			Addr:     kafka.TCP(c.Kafka.Broker.Endpoints...),
			Balancer: &kafka.LeastBytes{},
		},
	}
}
