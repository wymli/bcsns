package main

import (
	"flag"
	"fmt"

	"github.com/wymli/bcsns/app/msg_pull/internal/config"
	"github.com/wymli/bcsns/app/msg_pull/internal/server"
	"github.com/wymli/bcsns/app/msg_pull/internal/svc"
	"github.com/wymli/bcsns/app/msg_pull/pb"
	conf "github.com/wymli/bcsns/common/config"
	"github.com/wymli/bcsns/common/grpc/interceptor"
	"github.com/wymli/bcsns/common/server_framework/rpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/msg_send.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoadConfig(*configFile, &c)

	svcCtx := svc.NewServiceContext(c)
	svr := server.NewMsgPullServer(svcCtx)

	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(interceptor.RpcLogFilter))
	opts = append(opts, grpc.UnaryInterceptor(interceptor.RpcErrConvertFilter))

	grpcServer := grpc.NewServer(opts...)

	pb.RegisterMsgPullServer(grpcServer, svr)

	rpcServer := rpc.NewRpcServer(grpcServer, c.RpcServerConfig)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)

	defer rpcServer.Stop()
	rpcServer.MustStart()
}
