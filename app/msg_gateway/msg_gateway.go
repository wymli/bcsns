package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/wymli/bcsns/app/msg_gateway/config"
	"github.com/wymli/bcsns/app/msg_gateway/rpc/biz/server"
	"github.com/wymli/bcsns/app/msg_gateway/rpc/pb"
	"github.com/wymli/bcsns/app/msg_gateway/svc"
	"github.com/wymli/bcsns/app/msg_gateway/tcp/handler"
	"github.com/wymli/bcsns/common/codec"
	"github.com/wymli/bcsns/common/server_framework/rpc"
	"github.com/wymli/bcsns/common/server_framework/tcp"

	conf "github.com/wymli/bcsns/common/config"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/msg_gateway.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoadConfig(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		startRpcServer(&c, svcCtx)
		wg.Done()
	}()
	go func() {
		startTcpServer(&c, svcCtx)
		wg.Done()
	}()
	wg.Wait()
}

func startRpcServer(c *config.Config, svcCtx *svc.ServiceContext) {
	svr := server.NewGatewayServer(svcCtx)

	var opts []grpc.ServerOption
	opts = append(opts, rpc.DefaultServerIntercepterOpts()...)

	grpcServer := grpc.NewServer(opts...)

	pb.RegisterGatewayServer(grpcServer, svr)

	rpcServer := rpc.NewRpcServer(grpcServer, c.RpcServerConfig)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)

	defer rpcServer.Stop()
	rpcServer.MustStart()
}

func startTcpServer(c *config.Config, svcCtx *svc.ServiceContext) {
	tcpServer := tcp.NewServer()
	handler.RegisterHandlers(tcpServer, svcCtx)
	tcpServer.RegisterFramer(codec.Framer())
	tcpServer.RegisterDecoder(codec.Decoder())

	defer tcpServer.Stop()
	fmt.Printf("Starting tcp server at %s...\n", c.ListenOn)
	tcpServer.Start()
}
