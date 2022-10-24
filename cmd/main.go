package main

import (
	"flag"
	"github.com/zenpk/dorm-system/internal/controller"
	"github.com/zenpk/dorm-system/internal/rpc"
	"github.com/zenpk/dorm-system/pkg/config"
	"github.com/zenpk/dorm-system/pkg/zap"
	"log"
)

func main() {
	// read mode from commandline, default as development
	mode := flag.String("mode", "development", "define program mode")
	flag.Parse()
	// Viper
	if err := config.InitGlobalConfig(*mode); err != nil {
		log.Fatalf("failed to initialize Viper: %v", err)
	}
	// zap
	if err := zap.InitLogger(*mode); err != nil {
		log.Fatalf("failed to initialize zap: %v", err)
	}
	defer zap.Logger.Sync()
	// GORM
	//if err := dal.InitDB(); err != nil {
	//	log.Fatalf("failed to initialize database: %v", err)
	//}
	// Redis
	//if err := cache.InitRedis(); err != nil {
	//	log.Fatalf("failed to initialize Redis: %v", err)
	//}
	// RPC connections
	connList, err := rpc.InitClient()
	if err != nil {
		log.Fatalf("failed to initialize RPC clients: %v", err)
	}
	for _, conn := range connList {
		defer conn.Close()
	}
	// Gin
	if err := controller.InitGin(); err != nil {
		log.Fatalf("failed to initialize Gin: %v", err)
	}
}
