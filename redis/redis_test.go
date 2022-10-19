package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
)

func TestRedis(t *testing.T) {
	err := redisClient.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := redisClient.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := redisClient.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
