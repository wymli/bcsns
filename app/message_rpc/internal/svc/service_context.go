package svc

import (
	"github.com/bwmarrin/snowflake"
	"github.com/segmentio/kafka-go"
	"github.com/wymli/bcsns/app/message_rpc/internal/config"
	"github.com/wymli/bcsns/app/seq_service/seq"

	mylogx "github.com/wymli/bcsns/common/logx"
	zerologx "github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	KafkaClient *kafka.Writer
	SeqRpc      seq.Seq
	SnowflakeNode        *snowflake.Node
}

func NewServiceContext(c config.Config) *ServiceContext {
	zerologx.Disable()
	mylogx.Init(c.Logx)

	node, err := snowflake.NewNode(c.SnowflakeNodeNumber)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
		KafkaClient: &kafka.Writer{
			Addr:     kafka.TCP(c.Kafka.Broker.Endpoints...),
			Balancer: &kafka.LeastBytes{},
		},
		SeqRpc: seq.NewSeq(zrpc.MustNewClient(c.SeqRpcConfig)),
		SnowflakeNode:   node,
	}
}
