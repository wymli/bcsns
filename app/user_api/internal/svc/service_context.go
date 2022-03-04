package svc

import (
	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"github.com/wymli/bcsns/app/auth_rpc/auth"
	"github.com/wymli/bcsns/app/user_api/internal/config"
	"github.com/wymli/bcsns/common/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type ServiceContext struct {
	Config       config.Config
	DgraphClient *dgo.Dgraph
	AuthClient   auth.Auth
}

func NewServiceContext(c config.Config) *ServiceContext {
	logx.Init(c.Logx)

	return &ServiceContext{
		Config:       c,
		DgraphClient: newDgraphClient(c.Dgraph.Endpoint),
		AuthClient:   auth.NewAuth(zrpc.MustNewClient(c.AuthRpcConf)),
	}
}

func newDgraphClient(endpoint string) *dgo.Dgraph {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	d, err := grpc.Dial(endpoint, grpc.WithInsecure())
	logx.FatalIfErrf(err, "failed to dial dgraph endpoint:%s", endpoint)

	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
}
