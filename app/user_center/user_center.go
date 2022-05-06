package main

import (
	"flag"
	"fmt"

	"github.com/wymli/bcsns/app/user_center/internal/config"
	"github.com/wymli/bcsns/app/user_center/internal/server"
	"github.com/wymli/bcsns/app/user_center/internal/svc"
	"github.com/wymli/bcsns/app/user_center/pb"
	conf "github.com/wymli/bcsns/common/config"
	"github.com/wymli/bcsns/common/server_framework/rpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/user_center.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoadConfig(*configFile, &c)

	var opts []grpc.ServerOption
	opts = append(opts, rpc.DefaultServerIntercepterOpts()...)

	grpcServer := grpc.NewServer(opts...)

	rpcServer := rpc.NewRpcServer(grpcServer, c.RpcServerConfig)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)

	rpcServer.Init()

	svcCtx := svc.NewServiceContext(c)
	svr := server.NewUserCenterServer(svcCtx)

	pb.RegisterUserCenterServer(grpcServer, svr)

	defer rpcServer.Stop()
	rpcServer.MustStart()
}
