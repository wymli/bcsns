package svc

import (
	"github.com/bwmarrin/snowflake"
	"github.com/go-redis/redis/v8"
	"github.com/segmentio/kafka-go"
	"github.com/wymli/bcsns/app/msg_send/internal/config"

	"github.com/wymli/bcsns/common/logx"
)

type ServiceContext struct {
	Config        config.Config
	KafkaClient   *kafka.Writer
	RedisClient   *redis.Client
	SnowflakeNode *snowflake.Node
}

type KafkaTopic struct{
	
}

func NewServiceContext(c config.Config) *ServiceContext {
	logx.Init(c.Log)

	node, err := snowflake.NewNode(c.Snowflake.NodeNumber)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
		KafkaClient: &kafka.Writer{
			Addr:     kafka.TCP(c.Kafka.Broker.Endpoints...),
			Balancer: &kafka.LeastBytes{},
		},
		// SeqRpc: seq.NewSeq(zrpc.MustNewClient(c.SeqRpcConfig)),
		RedisClient: redis.NewClient(&redis.Options{
			Addr: c.Redis.Host,
		}),
		SnowflakeNode: node,
	}
}
