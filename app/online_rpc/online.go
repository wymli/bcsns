package main

import (
	"flag"
	"fmt"

	"github.com/wymli/bcsns/app/online_rpc/internal/config"
	"github.com/wymli/bcsns/app/online_rpc/internal/server"
	"github.com/wymli/bcsns/app/online_rpc/internal/svc"
	"github.com/wymli/bcsns/app/online_rpc/pb"
	conf "github.com/wymli/bcsns/common/config"
	"github.com/wymli/bcsns/common/server_framework/rpc"

	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/online.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoadConfig(*configFile, &c)

	svcCtx := svc.NewServiceContext(c)
	svr := server.NewOnlineServer(svcCtx)

	var opts []grpc.ServerOption
	opts = append(opts, rpc.DefaultServerIntercepterOpts()...)

	grpcServer := grpc.NewServer(opts...)

	pb.RegisterOnlineServer(grpcServer, svr)

	rpcServer := rpc.NewRpcServer(grpcServer, c.RpcServerConfig)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)

	defer rpcServer.Stop()
	rpcServer.MustStart()
}
