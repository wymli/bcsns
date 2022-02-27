package rpcfilter

import (
	"context"

	"github.com/wymli/bcsns/common/errx"
	"google.golang.org/grpc"
)

func RpcErrConvertFilter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (rsp interface{}, err error) {
	rsp, err = handler(ctx, req)

	return rsp, errx.ToGrpcError(err)
}
