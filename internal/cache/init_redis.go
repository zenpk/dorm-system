package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var Redis *redis.Client
var CTX = context.Background()

func InitRedis() error {
	Redis = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.host") + ":" + viper.GetString("redis.port"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})
	if _, err := Redis.Ping(CTX).Result(); err != nil {
		return err
	}
	if err := Redis.FlushAll(CTX).Err(); err != nil { // flush all caches, comment this line in production
		return err
	}
	return nil
}
