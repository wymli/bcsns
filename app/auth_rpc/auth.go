package main

import (
	"flag"
	"fmt"

	"github.com/wymli/bcsns/app/auth_rpc/internal/config"
	"github.com/wymli/bcsns/app/auth_rpc/internal/server"
	"github.com/wymli/bcsns/app/auth_rpc/internal/svc"
	"github.com/wymli/bcsns/app/auth_rpc/pb"
	conf "github.com/wymli/bcsns/common/config"
	"github.com/wymli/bcsns/common/server_framework/rpc"

	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/auth.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoadConfig(*configFile, &c)

	grpcServer := grpc.NewServer(rpc.DefaultServerIntercepterOpts()...)

	rpcServer := rpc.NewRpcServer(grpcServer, c.RpcServerConfig)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)

	rpcServer.Init()

	svcCtx := svc.NewServiceContext(c)
	svr := server.NewAuthServer(svcCtx)
	pb.RegisterAuthServer(grpcServer, svr)

	defer rpcServer.Stop()
	rpcServer.MustStart()
}
