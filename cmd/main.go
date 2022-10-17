package main

import (
	"flag"
	"github.com/zenpk/dorm-system/internal/controller"
	"github.com/zenpk/dorm-system/internal/dal"
	"github.com/zenpk/dorm-system/pkg/viper"
	"github.com/zenpk/dorm-system/pkg/zap"
)

func main() {
	// read mode from commandline, default as development
	mode := flag.String("mode", "development", "define program mode")
	flag.Parse()
	// Viper
	if err := viper.InitConfig(*mode, "./configs"); err != nil {
		panic(err)
	}
	// zap
	if err := zap.InitLogger(); err != nil {
		panic(err)
	}
	defer zap.Logger.Sync()
	// GORM
	if err := dal.InitDB(); err != nil {
		panic(err)
	}
	// Redis
	//if err := cache.InitRedis(); err != nil {
	//	panic(err)
	//}
	// Gin
	if err := controller.InitGin(); err != nil {
		panic(err)
	}
}
