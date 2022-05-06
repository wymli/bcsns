package svc

import (
	"github.com/gocql/gocql"
	"github.com/segmentio/kafka-go"
	pbbc "github.com/wymli/bcsns/app/bc_proxy/pb"
	pbgw "github.com/wymli/bcsns/app/msg_gateway/rpc/pb"
	"github.com/wymli/bcsns/app/msg_sync/internal/config"
	pbonline "github.com/wymli/bcsns/app/online_rpc/pb"
	pbuser "github.com/wymli/bcsns/app/user_center/pb"
	"github.com/wymli/bcsns/common/logx"
	connpool "github.com/wymli/bcsns/common/pool/conn"
	"github.com/wymli/bcsns/common/server_framework/rpc"
	"google.golang.org/grpc"
)

type ServiceContext struct {
	Config          config.Config
	KafkaClient     *kafka.Writer
	GatewayRpcPool  connpool.ConnPool
	CassandraClient *gocql.Session
	OnlineRpc       pbonline.OnlineClient
	UserRpc         pbuser.UserCenterClient
	BCProxyRpc      pbbc.BcProxyClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	logx.Init(c.Log)

	gwConnPool := connpool.MustNewConnPool(func(endpoint string) (interface{}, error) {
		client := pbgw.NewGatewayClient(rpc.Must(grpc.Dial(endpoint)))
		return client, nil
	}, c.GatewayRpc)

	clusterConf := gocql.NewCluster(c.Cassandra.Endpoints...)
	clusterConf.Keyspace = c.Cassandra.KeySpace
	clusterConf.Consistency = gocql.Consistency(c.Cassandra.Consistency)

	sess, err := clusterConf.CreateSession()
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
		KafkaClient: &kafka.Writer{
			Addr:     kafka.TCP(c.Kafka.Broker.Endpoints...),
			Balancer: &kafka.LeastBytes{},
		},
		GatewayRpcPool:  gwConnPool,
		CassandraClient: sess,
		OnlineRpc:       pbonline.NewOnlineClient(rpc.Must(grpc.Dial(c.OnlineRpc.Endpoint))),
		UserRpc:         pbuser.NewUserCenterClient(rpc.Must(grpc.Dial(c.UserRpc.Endpoint))),
		BCProxyRpc:      pbbc.NewBcProxyClient(rpc.Must(grpc.Dial(c.BCProxyRpc.Endpoint))),
	}
}
