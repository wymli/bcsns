package rpc

import (
	"github.com/wymli/bcsns/common/tracer"
	"google.golang.org/grpc"
)

func Must(cc *grpc.ClientConn, err error) *grpc.ClientConn {
	if err != nil {
		panic(err)
	}

	return cc
}

func Dial(target string, opts ...grpc.DialOption) *grpc.ClientConn {
	opts = append(opts, DefaultDialOpts()...)
	return Must(grpc.Dial(target, opts...))
}

func DefaultDialOpts() []grpc.DialOption {
	return []grpc.DialOption{grpc.WithInsecure(), grpc.WithUnaryInterceptor(tracer.RpcClientTraceInterceptor)}
}
