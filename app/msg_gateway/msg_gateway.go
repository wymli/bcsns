package main

import (
	"flag"
	"fmt"

	"github.com/wymli/bcsns/app/msg_gateway/config"
	"github.com/wymli/bcsns/app/msg_gateway/rpc/biz/server"
	"github.com/wymli/bcsns/app/msg_gateway/rpc/pb"
	"github.com/wymli/bcsns/app/msg_gateway/svc"
	"github.com/wymli/bcsns/app/msg_gateway/tcp/handler"
	"github.com/wymli/bcsns/common/interceptor/rpcfilter"
	"github.com/wymli/bcsns/pkg/codec"
	"github.com/wymli/bcsns/pkg/server_framework/tcp"

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

	startRpcServer(&c, ctx)
}

func startRpcServer(c *config.Config, svcCtx *svc.ServiceContext) {
	srv := server.NewGatewayServer(svcCtx)

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

func startTcpServer(c *config.Config, svcCtx *svc.ServiceContext) {
	s := tcp.NewServer()
	handler.RegisterHandlers(s, svcCtx)
	s.RegisterFramer(codec.Framer())
	s.RegisterDecoder(codec.Decoder())

	defer s.Stop()
	fmt.Printf("Starting tcp server at %s...\n", c.ListenOn)
	s.Start()
}
