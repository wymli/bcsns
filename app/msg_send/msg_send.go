package main

import (
	"flag"
	"fmt"

	"github.com/wymli/bcsns/app/msg_send/internal/config"
	"github.com/wymli/bcsns/app/msg_send/internal/server"
	"github.com/wymli/bcsns/app/msg_send/internal/svc"
	"github.com/wymli/bcsns/app/msg_send/pb"
	conf "github.com/wymli/bcsns/common/config"
	"github.com/wymli/bcsns/common/server_framework/rpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/msg_send.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoadConfig(*configFile, &c)

	svcCtx := svc.NewServiceContext(c)
	svr := server.NewMsgSendServer(svcCtx)

	var opts []grpc.ServerOption
	opts = append(opts, rpc.DefaultServerIntercepterOpts()...)

	grpcServer := grpc.NewServer(opts...)

	pb.RegisterMsgSendServer(grpcServer, svr)

	rpcServer := rpc.NewRpcServer(grpcServer, c.RpcServerConfig)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)

	defer rpcServer.Stop()
	rpcServer.MustStart()
}
