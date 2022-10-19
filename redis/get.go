package redis

import "github.com/go-redis/redis/v8"

func GetValue(i IRedis) *redis.StringCmd {
	return redisClient.Get(ctx, i.GetKey())
}
