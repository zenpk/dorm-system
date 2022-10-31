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
	var ctx = context.Background()
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
	redisCli := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.host") + ":" + viper.GetString("redis.port"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})
	if err := warmAvailable(redisCli); err != nil {
		return err
	}
	return nil
}

// warmAvailable preload available bed num into Redis
func warmAvailable(redisCli *redis.Client) error {
	var ctx context.Context
	buildingTable := new(dal.Building)
	dormTable := new(dal.Dorm)
	all := int64(0)
	availIds, err := buildingTable.FindAllAvailableIds(ctx)
	if err != nil {
		return err
	}
	for _, id := range availIds {
		num, err := dormTable.SumAvailableByBuildingId(ctx, id)
		if err != nil {
			return err
		}
		Redis.HSet(ctx, "available", id, num)
		all += num
	}
	Redis.HSet(ctx, "available", "all", all)
	return redisCli.Close()
}
