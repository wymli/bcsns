package rpcfilter

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
)

func RpcLogFilter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (rsp interface{}, err error) {
	rsp, err = handler(ctx, req)
	// todo: 一般是使用结构化日志,以便查询,不过es也支持全文搜索,所以这里就算了,毕竟这个框架也没实现结构化日志功能
	if err != nil {
		logx.Errorf("req:%v serverInfo:%v rsp:%v err:%v", req, *info, rsp, err)
	} else {
		logx.Infof("req:%v serverInfo:%v rsp:%v err:nil", req, *info, rsp)
	}

	return rsp, err
}
