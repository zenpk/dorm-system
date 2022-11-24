package test

import (
	"github.com/gin-gonic/gin"
)

// SetupHandler is for setting up a temporary test environment
func SetupHandler(c *gin.Context) {
	//ctx := context.Background()
	//cache.Redis.Set(ctx, "count", 10, 10*time.Minute)
	//cache.All.RedSync.Init(cache.Redis)
	//c.JSON(200, "ok")
}

// Handler is for testing
func Handler(c *gin.Context) {
	//ctx := context.Background()
	//m := cache.All.RedSync.NewMutex("a")
	//if err := m.Lock(); err != nil {
	//	fmt.Println(err)
	//}
	//now, _ := cache.Redis.Get(ctx, "count").Int()
	//if now > 0 {
	//	now, _ = cache.Redis.Get(ctx, "count").Int()
	//	cache.Redis.Set(ctx, "count", now-1, 10*time.Minute)
	//}
	//if ok, err := m.Unlock(); err != nil {
	//	fmt.Println(ok, err)
	//}
	//c.JSON(200, cache.Redis.Get(ctx, "count").String())
}
