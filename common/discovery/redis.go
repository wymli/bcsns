package discovery

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func InitRedis(rc *redis.Client) {
	redisClient = rc
}

func checkInit() {
	if redisClient == nil {
		panic("redis not set in discovery")
	}
}

func Register(serviceName, endpoint string) (cancel func()) {
	checkInit()

	if err := redisClient.SAdd(context.Background(), serviceName, endpoint).Err(); err != nil {
		panic("failed to regiser service discovery: " + err.Error())
	}

	return func() {
		Unregister(serviceName, endpoint)
	}
}

func Unregister(serviceName, endpoint string) {
	checkInit()

	if err := redisClient.SRem(context.Background(), serviceName, endpoint).Err(); err != nil {
		panic("failed to unregiser service discovery: " + err.Error())
	}
}
