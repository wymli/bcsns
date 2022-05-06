package rpc

import (
	"net"
	"os"
	"strings"

	"github.com/go-redis/redis/v8"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/wymli/bcsns/common/config"
	"github.com/wymli/bcsns/common/discovery"
	"github.com/wymli/bcsns/common/grpc/interceptor"
	"github.com/wymli/bcsns/common/grpc/resolver"
	"github.com/wymli/bcsns/common/metric"
	"github.com/wymli/bcsns/common/signal"
	"github.com/wymli/bcsns/common/tracer"
	"google.golang.org/grpc"
)

type RpcServer struct {
	GrpcServer *grpc.Server
	Config     config.RpcServerConfig
	StopTrace  func()
	Unregister func()
}

func NewRpcServer(svr *grpc.Server, config config.RpcServerConfig) *RpcServer {
	return &RpcServer{
		GrpcServer: svr,
		Config:     config,
		StopTrace:  nil,
		Unregister: nil,
	}
}

func (s *RpcServer) Init() {
	// init redis
	switch s.Config.ServiceDiscovery.Discovery {
	case "redis":
		rc := redis.NewClient(&redis.Options{
			Addr: s.Config.ServiceDiscovery.Endpoints[0],
		})

		resolver.InitRedis(rc)
		discovery.InitRedis(rc)

		// 获取 hostip
		ip := os.Getenv("HOST_IP")
		idx := strings.Index(s.Config.ListenOn, ":")
		port := s.Config.ListenOn[idx:]

		s.Unregister = discovery.Register(s.Config.Name, ip+port)

	case "direct":
		// do nothing
	default:
		panic("not implemented service discovery")
	}

	// start trace
	s.StopTrace = tracer.Init(s.Config.Trace)
	// start metric
	metric.RegisterGrpcServer(s.GrpcServer)
	metric.StartMetricServer(s.Config.Metric)
}

func (s *RpcServer) MustStart() {
	done := signal.SignalCatch(s.Stop)

	lis, err := net.Listen("tcp", s.Config.ListenOn)
	if err != nil {
		panic(err)
	}

	if err := s.GrpcServer.Serve(lis); err != nil {
		panic(err)
	}

	<-done
}

func (s *RpcServer) Stop() {
	s.GrpcServer.GracefulStop()
	s.StopTrace()
	s.Unregister()
}

func DefaultServerIntercepterOpts() []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			interceptor.RpcErrConvertFilter,
			interceptor.RpcLogFilter,
			tracer.RpcServerTraceInterceptor,
			metric.RpcServerMetricInterceptor,
		)),
	}
}
