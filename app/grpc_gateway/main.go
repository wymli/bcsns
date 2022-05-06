package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pbuser "github.com/wymli/bcsns/app/user_center/pb"
	"github.com/wymli/bcsns/common/config"
	"github.com/wymli/bcsns/common/grpc/resolver"
	"github.com/wymli/bcsns/common/logx"
	"github.com/wymli/bcsns/common/metadata"
	"github.com/wymli/bcsns/common/server_framework/rpc"
	"github.com/wymli/bcsns/common/tracer"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

var configFile = flag.String("f", "etc/gateway.yaml", "the config file")

func main() {
	flag.Parse()

	c := Config{}
	config.MustLoadConfig(*configFile, &c)

	logx.Init(c.Log)

	resolver.InitRedisFromZero(c.ServiceDiscovery)

	// init tracer
	cancel := tracer.Init(c.Trace)
	defer cancel()

	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(CustomHeaderMatcher),
	)

	if err := pbuser.RegisterUserCenterGWFromEndpoint(context.Background(), mux, c.UserCenterRpc.Endpoint, rpc.DefaultDialOpts()); err != nil {
		panic(err)
	}

	// pbmsgsend.RegisterMsgSendGWFromEndpoint(context.Background(), mux, c.MsgSendRpc.Endpoint, rpcClientDialOpts)
	// pbmsgpull.RegisterMsgPullGWFromEndpoint(context.Background(), mux, c.MsgPullRpc.Endpoint, rpcClientDialOpts)
	// pbbc.RegisterBcProxyGWFromEndpoint(context.Background(), mux, c.BCPullRpc.Endpoint, rpcClientDialOpts)

	fmt.Println("api_gateway serving on " + c.ListenOn)

	if err := http.ListenAndServe(c.ListenOn, otelhttp.NewHandler(mux, c.Name)); err != nil {
		panic(err)
	}
}

func CustomHeaderMatcher(key string) (string, bool) {
	switch key {
	case metadata.XUSERID:
		return key, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}
