package main

import (
	"flag"
	"fmt"

	"github.com/wymli/bcsns/app/message_rpc/internal/config"
	"github.com/wymli/bcsns/app/message_rpc/internal/server"
	"github.com/wymli/bcsns/app/message_rpc/internal/svc"
	"github.com/wymli/bcsns/app/message_rpc/pb"
	"github.com/wymli/bcsns/common/interceptor/rpcfilter"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/message.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewMessageServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterMessageServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	// rpc filter
	s.AddUnaryInterceptors(rpcfilter.RpcLogFilter)
	s.AddUnaryInterceptors(rpcfilter.RpcErrConvertFilter)

	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
