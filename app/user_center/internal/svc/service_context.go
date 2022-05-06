package svc

import (
	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	pbauth "github.com/wymli/bcsns/app/auth_rpc/pb"
	"github.com/wymli/bcsns/app/user_center/internal/config"
	"google.golang.org/grpc"

	"github.com/wymli/bcsns/common/logx"
	"github.com/wymli/bcsns/common/server_framework/rpc"
)

type ServiceContext struct {
	Config       config.Config
	DgraphClient *dgo.Dgraph
	AuthClient   pbauth.AuthClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	logx.Init(c.Log)

	if c.MockAll {
		return &ServiceContext{
			Config: c,
		}
	}

	return &ServiceContext{
		Config:       c,
		DgraphClient: newDgraphClient(c.Dgraph.Endpoint),
		AuthClient:   pbauth.NewAuthClient(rpc.Dial(c.AuthRpc.Endpoint)),
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
