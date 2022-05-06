package resolver

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/wymli/bcsns/common/config"
	"github.com/wymli/bcsns/common/logx"
	"google.golang.org/grpc/resolver"
)

func init() {
	resolver.Register(&redisResolverBuilder{})
}

var redisClient *redis.Client

func InitRedis(rc *redis.Client) {
	redisClient = rc
}

func InitRedisFromZero(c config.ServiceDiscovery) {
	if c.Discovery != "redis" {
		panic("discovery is not redis")
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr: c.Endpoints[0],
	})
}

type redisResolverBuilder struct{}

func (b *redisResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	if len(target.URL.Path) == 0 {
		return nil, fmt.Errorf("endpoint not set")
	}

	url := b.removeScheme(target.URL.Path)

	redisResolver := &redisResolver{redisKey: url, cc: cc}
	redisResolver.ResolveNow(resolver.ResolveNowOptions{})

	return redisResolver, nil
}

func (b *redisResolverBuilder) Scheme() string {
	return "redis"
}

func (b *redisResolverBuilder) removeScheme(url string) string {
	return strings.TrimLeft(url, b.Scheme()+":///")
}

type redisResolver struct {
	redisKey string
	cc       resolver.ClientConn
}

func (rr *redisResolver) ResolveNow(_ resolver.ResolveNowOptions) {
	if redisClient == nil {
		logx.GlobalLogger.Error().Msg("redis not set in service discovery")
		return
	}

	res, err := redisClient.SMembers(context.Background(), rr.redisKey).Result()
	if err != nil {
		logx.GlobalLogger.Err(err).Msg("failed to access redis in 'ResolveNow' in service discovery")
		return
	}

	addrs := make([]resolver.Address, len(res))
	for i, url := range res {
		addrs[i] = resolver.Address{Addr: url}
	}

	if err := rr.cc.UpdateState(resolver.State{
		Addresses: addrs,
	}); err != nil {
		return
	}
}

func (b *redisResolver) Close() {}
