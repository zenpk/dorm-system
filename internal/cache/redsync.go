package cache

import (
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

type RedSync struct {
	rs *redsync.Redsync
}

// Init the distributed lock from client
func (r *RedSync) Init(client *redis.Client) {
	// Create a pool with go-redis (or redigo) which is the pool redisync will
	// use while communicating with Redis. This can also be any pool that
	// implements the `redis.Pool` interface.
	pool := goredis.NewPool(client)

	// Create an instance of redisync to be used to obtain a mutual exclusion
	// lock.
	r.rs = redsync.New(pool)
}

// NewMutex get a distributed lock from name, for this project, mutexName will be the buildingId
// possible errors: Lock():
func (r *RedSync) NewMutex(mutexName string) *redsync.Mutex {
	return r.rs.NewMutex(mutexName)
}
