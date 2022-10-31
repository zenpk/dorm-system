package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/internal/dal"
)

var Redis *redis.Client

func InitRedis() error {
	Redis = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.host") + ":" + viper.GetString("redis.port"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})
	ctx := context.Background()
	if _, err := Redis.Ping(ctx).Result(); err != nil {
		return err
	}
	if err := Redis.FlushAll(ctx).Err(); err != nil { // flush all caches, comment this line in production
		return err
	}
	return nil
}

// Warming preload warm data into Redis
func Warming() error {
	if err := warmAvailable(); err != nil {
		return err
	}
	return nil
}

// warmAvailable preload available bed num into Redis
func warmAvailable() error {
	ctx := context.Background()
	all := int64(0)
	availIds, err := dal.Table.Building.FindAllAvailableIds(ctx)
	if err != nil {
		return err
	}
	for _, id := range availIds {
		num, err := dal.Table.Dorm.SumAvailableByBuildingId(ctx, id)
		if err != nil {
			return err
		}
		err = Redis.HSet(ctx, "available", id, num).Err()
		if err != nil {
			return err
		}
		all += num
	}
	return Redis.HSet(ctx, "available", "all", all).Err()
}
