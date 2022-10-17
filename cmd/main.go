package main

import (
	"flag"
	"github.com/zenpk/dorm-system/internal/controller"
	"github.com/zenpk/dorm-system/internal/dal"
	"github.com/zenpk/dorm-system/internal/middleware"
	"github.com/zenpk/dorm-system/internal/util"
)

func main() {
	// 从命令行中读取运行模式，默认为 development
	mode := flag.String("mode", "development", "define program mode")
	flag.Parse()
	// 加载 Viper
	if err := util.InitConfig(*mode, "./configs"); err != nil {
		panic(err)
	}
	// 加载 zap
	if err := middleware.InitLogger(); err != nil {
		panic(err)
	}
	defer middleware.Logger.Sync()
	// 加载 Gorm，连接 MySQL
	if err := dal.InitDB(); err != nil {
		panic(err)
	}
	// 连接 Redis，暂不使用
	//if err := cache.InitRedis(); err != nil {
	//	panic(err)
	//}
	// 加载 Gin 并开始监听
	if err := controller.InitGin(); err != nil {
		panic(err)
	}
}
