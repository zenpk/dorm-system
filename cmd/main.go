package main

import (
	"flag"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/internal/controller"
	"github.com/zenpk/dorm-system/internal/dal"
	"github.com/zenpk/dorm-system/pkg/config"
	"github.com/zenpk/dorm-system/pkg/gmp"
	"github.com/zenpk/dorm-system/pkg/zap"
	"log"
)

func main() {
	// read mode from commandline, default as development
	mode := flag.String("mode", "development", "define program mode")
	flag.Parse()
	path, err := gmp.GetModPath()
	if err != nil {
		log.Fatalln(err)
	}
	// Viper
	if err := config.InitConfig(*mode, path+"configs"); err != nil {
		log.Fatalln(err)
	}
	// zap
	if err := zap.InitLogger(path + viper.GetString("zap.log_path")); err != nil {
		log.Fatalln(err)
	}
	defer zap.Logger.Sync()
	// GORM
	if err := dal.InitDB(); err != nil {
		log.Fatalln(err)
	}
	// Redis
	//if err := cache.InitRedis(); err != nil {
	//	log.Fatalln(err)
	//}
	// Gin
	if err := controller.InitGin(); err != nil {
		log.Fatalln(err)
	}
}
