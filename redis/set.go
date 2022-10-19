package redis

import "github.com/go-redis/redis/v8"

func SetValue(i IRedis, value interface{}) *redis.StatusCmd {
	return redisClient.Set(ctx, i.GetKey(), value, i.GetDuration())
}
