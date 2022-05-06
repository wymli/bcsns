package metric

import (
	"net/http"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	// panicIf(prometheus.Register(serverReqDuration))
	// panicIf(prometheus.Register(serverRspCode))
	// panicIf(prometheus.Register(clientReqDuration))
	// panicIf(prometheus.Register(clientRspCode))
}

type Config struct {
	Path     string
	ListenOn string
}

func StartMetricServer(c Config) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				panic(err)
			}
		}()

		if c.Path == "" {
			c.Path = "/metrics"
		}
		http.Handle(c.Path, promhttp.Handler())
		http.ListenAndServe(c.ListenOn, nil)
	}()
}

var RpcServerMetricInterceptor = grpc_prometheus.UnaryServerInterceptor

var RpcClientMetricInterceptor = grpc_prometheus.UnaryClientInterceptor

var RegisterGrpcServer = grpc_prometheus.Register

// const (
// 	RPC_SERVER_NS  = "rpc_server"
// 	RPC_CLIENT_NS  = "rpc_client"
// 	HTTP_SERVER_NS = "http_server"
// )

// var serverReqDuration = prometheus.NewGaugeVec(prometheus.GaugeOpts{
// 	Namespace:   RPC_SERVER_NS,
// 	Subsystem:   "requests",
// 	Name:        "server_req_dur_ms",
// 	Help:        "rpc server requests duration(ms).",
// 	ConstLabels: map[string]string{},
// }, []string{"method"})

// // var reqDurHist = prometheus.NewHistogramVec(prometheus.HistogramOpts{
// // 	Namespace:   RPC_SERVER_NS,
// // 	Subsystem:   "requests",
// // 	Name:        "req_dur_ms",
// // 	Help:        "rpc server requests duration(ms).",
// // 	ConstLabels: map[string]string{},
// // 	Buckets:     []float64{},
// // }, []string{"method"})

// var serverRspCode = prometheus.NewGaugeVec(prometheus.GaugeOpts{
// 	Namespace:   RPC_SERVER_NS,
// 	Subsystem:   "response",
// 	Name:        "server_rsp_code_cnt",
// 	Help:        "rpc server response code count.",
// 	ConstLabels: map[string]string{},
// }, []string{"method", "code"})

// func RpcServerMetricInterceptor(ctx context.Context, req interface{},
// 	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
// ) (interface{}, error) {
// 	startTime := time.Now()
// 	resp, err := handler(ctx, req)

// 	serverRspCode.WithLabelValues(info.FullMethod, strconv.Itoa(int(status.Code(err)))).Inc()
// 	serverReqDuration.WithLabelValues(info.FullMethod).Set(float64(time.Since(startTime) / time.Millisecond))

// 	return resp, err
// }

// var clientReqDuration = prometheus.NewGaugeVec(prometheus.GaugeOpts{
// 	Namespace:   RPC_CLIENT_NS,
// 	Subsystem:   "requests",
// 	Name:        "client_req_dur_ms",
// 	Help:        "rpc client requests duration(ms).",
// 	ConstLabels: map[string]string{},
// }, []string{"method"})

// var clientRspCode = prometheus.NewGaugeVec(prometheus.GaugeOpts{
// 	Namespace:   RPC_CLIENT_NS,
// 	Subsystem:   "response",
// 	Name:        "client_rsp_code_cnt",
// 	Help:        "rpc client response code count.",
// 	ConstLabels: map[string]string{},
// }, []string{"method", "code"})

// func RpcClientMetricInterceptor(
// 	ctx context.Context,
// 	method string,
// 	req, reply interface{},
// 	cc *grpc.ClientConn,
// 	invoker grpc.UnaryInvoker,
// 	callOpts ...grpc.CallOption,
// ) error {
// 	startTime := time.Now()
// 	err := invoker(ctx, method, req, reply, cc, callOpts)

// 	clientRspCode.WithLabelValues(method, strconv.Itoa(int(status.Code(err)))).Inc()
// 	clientReqDuration.WithLabelValues(method).Set(float64(time.Since(startTime) / time.Millisecond))

// 	return err
// }

// func panicIf(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

// var httpServerReqDuration = prometheus.NewGaugeVec(prometheus.GaugeOpts{
// 	Namespace:   HTTP_SERVER_NS,
// 	Subsystem:   "requests",
// 	Name:        "server_req_dur_ms",
// 	Help:        "http server requests duration(ms).",
// 	ConstLabels: map[string]string{},
// }, []string{"method"})

// var httpServerRspCode = prometheus.NewGaugeVec(prometheus.GaugeOpts{
// 	Namespace:   HTTP_SERVER_NS,
// 	Subsystem:   "response",
// 	Name:        "server_rsp_code_cnt",
// 	Help:        "http server response code count.",
// 	ConstLabels: map[string]string{},
// }, []string{"method", "code"})

// func HttpMetricHandler(handler http.Handler) http.Handler {
// 	return
// }
