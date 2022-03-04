package gateway

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/wymli/bcsns/app/auth_rpc/auth"
	"github.com/wymli/bcsns/app/msg_gateway/svc"
	pb "github.com/wymli/bcsns/dependency/pb/tcp"
	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
	"github.com/wymli/bcsns/pkg/tcp"
)

type UserOnlineLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserOnlineLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserOnlineLogic {
	return UserOnlineLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserOnlineLogic) UserOnline(connCtx *tcp.ConnCtx, req *pb.UserOnlineReq) (resp *pb.CommonResp, err error) {
	rsp, err := l.svcCtx.AuthClient.ValidateToken(context.Background(), &auth.ValidateTokenReq{
		UserId: req.UserId,
		Token:  req.Token,
	})
	if err != nil {
		return nil, errx.Wrapf(err, "failed to validate token, err:%v", err)
	}
	if !rsp.Ok {
		return nil, errx.ERROR_TOKEN_INVALID
	}

	l.svcCtx.UserConnPool.AddConn(req.UserId, connCtx.Conn)

	connCtx.Ctx = context.WithValue(connCtx.Ctx, "userId", req.UserId)

	// todo: 统一一下配置获取,如果部署在k8s内,如果是暴露自己,那么肯定是通过环境变量的方式,统一一下配置获取,看要不要上viper
	// todo: k8s内,pod要以某种方式暴露给集群外,只有nodeport,ingress gateway几种方式,ingress似乎偏向于http,对于tcp可能就直接nodeport即可(存疑)
	// todo: 如果暴露nodeport,那么需要获知本机ip和nodeport
	// todo: docker-compose 在水平拓展服务中,一般是将容器的某个端口映射到主机上未分配的端口,外部访问要先知道该随机端口,所以一般并不建立端口映射,而是使用负载均衡器/网关(比如nginx),nginx将会使用docker dns直接访问水平拓展的container,比如nginx listen:80 proxy_pass tcp://app:8080
	// todo: k8s 中,与之类似,如果强制进行端口映射,那么就限制了一台机器最多一个服务,比较受限. 可以建立nodeport service,也是使用dns转发请求到对应名字
	addr := ""
	switch l.svcCtx.Config.Deploy.Use {
	case "docker":
		addr = "localhost:" + strings.Split(l.svcCtx.Config.ListenOn, ":")[1]
	case "docker-compose":
		// 返回负载均衡的地址
		return nil, errx.Wrap(errx.ERROR_SERVER_UNIMPLEMENTED, "deploy environment: docker-compose is not implemented")
	case "k8s":
		// 返回nodeport/ingress service地址,这个地址会在环境变量里
		addr = os.Getenv("SERVICE_MSG_GATEWAY_HOST") + ":" + os.Getenv("SERVICE_MSG_GATEWAY_PORT")
	}

	// set redis
	k := fmt.Sprintf(l.svcCtx.Config.Online.KeyFormat, req.UserId)
	err = l.svcCtx.RedisClient.Setex(k, addr,
		l.svcCtx.Config.Online.RefreshInterval*int(time.Minute))
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_REDIS, "failed to setex user online of key:%v, err:%v", k, err)
	}

	return &pb.CommonResp{
		Code: 0,
		Msg:  "OK",
	}, nil
}
