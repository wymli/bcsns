package svc

import (
	"github.com/gocql/gocql"
	"github.com/wymli/bcsns/app/msg_pull/internal/config"

	"github.com/wymli/bcsns/common/logx"
)

type ServiceContext struct {
	Config          config.Config
	CassandraClient *gocql.Session
}

func NewServiceContext(c config.Config) *ServiceContext {
	logx.Init(c.Log)

	clusterConf := gocql.NewCluster(c.Cassandra.Endpoints...)
	clusterConf.Keyspace = c.Cassandra.KeySpace
	clusterConf.Consistency = gocql.Consistency(c.Cassandra.Consistency)

	sess, err := clusterConf.CreateSession()
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:          c,
		CassandraClient: sess,
	}
}
