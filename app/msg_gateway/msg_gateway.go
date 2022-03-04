package main

import (
	"flag"
	"fmt"

	"github.com/wymli/bcsns/app/msg_gateway/config"
	"github.com/wymli/bcsns/app/msg_gateway/rpc/biz/server"
	"github.com/wymli/bcsns/app/msg_gateway/svc"
	"github.com/wymli/bcsns/app/msg_gateway/rpc/pb"
	"github.com/wymli/bcsns/common/interceptor/rpcfilter"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/msg_gateway.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewGatewayServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterGatewayServer(grpcServer, srv)

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

// func startTcpServer() {
// 	s := tcp.NewServer()
// 	s.AddRoutes()

// 	frame := tcpprotocol.TcpFrame{}
// 	s.Framer = frame.Framer()
// 	s.Decoder = frame.Decoder()
// }
