package seqmgr

import (
	"fmt"
	"strconv"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type RedisStore struct {
	key         string
	redisClient *redis.Redis
}

func NewRedisStorer(key string, redisClient *redis.Redis) (StoreKVer, error) {
	return &RedisStore{
		key:         key,
		redisClient: redisClient,
	}, nil
}

func (rs *RedisStore) Store(v uint64) error {
	return rs.redisClient.Set(rs.key, fmt.Sprintf("%d", v))
}

func (rs *RedisStore) Read() (uint64, error) {
	val, err := rs.redisClient.Get(rs.key)
	if err != nil {
		return 0, err
	}

	return strconv.ParseUint(val, 10, 64)
}
