package main

import (
	"flag"
	"fmt"

	"github.com/wymli/bcsns/app/seq_service/internal/config"
	"github.com/wymli/bcsns/app/seq_service/internal/server"
	"github.com/wymli/bcsns/app/seq_service/internal/svc"
	"github.com/wymli/bcsns/app/seq_service/pb"

	conf "github.com/wymli/bcsns/common/config"
	"github.com/wymli/bcsns/common/grpc/interceptor"
	"github.com/wymli/bcsns/common/server_framework/rpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/seq.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoadConfig(*configFile, &c)

	svcCtx := svc.NewServiceContext(c)
	svr := server.NewSeqServer(svcCtx)

	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(interceptor.RpcLogFilter))
	opts = append(opts, grpc.UnaryInterceptor(interceptor.RpcErrConvertFilter))

	grpcServer := grpc.NewServer(opts...)

	pb.RegisterSeqServer(grpcServer, svr)

	rpcServer := rpc.NewRpcServer(grpcServer, c.RpcServerConf)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)

	defer rpcServer.Stop()
	rpcServer.MustStart()
}
