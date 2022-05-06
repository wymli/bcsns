package main

import (
	"flag"
	"fmt"

	"github.com/wymli/bcsns/app/bc_proxy/internal/config"
	"github.com/wymli/bcsns/app/bc_proxy/internal/server"
	"github.com/wymli/bcsns/app/bc_proxy/internal/svc"
	"github.com/wymli/bcsns/app/bc_proxy/pb"
	conf "github.com/wymli/bcsns/common/config"
	_ "github.com/wymli/bcsns/common/grpc/resolver"
	"github.com/wymli/bcsns/common/server_framework/rpc"

	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/bc_proxy.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoadConfig(*configFile, &c)

	svcCtx := svc.NewServiceContext(c)
	svr := server.NewBcProxyServer(svcCtx)

	var opts []grpc.ServerOption
	opts = append(opts, rpc.DefaultServerIntercepterOpts()...)

	grpcServer := grpc.NewServer(opts...)

	pb.RegisterBcProxyServer(grpcServer, svr)

	rpcServer := rpc.NewRpcServer(grpcServer, c.RpcServerConfig)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)

	defer rpcServer.Stop()
	rpcServer.MustStart()
}
