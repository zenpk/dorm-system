package main

import (
	"flag"
	"github.com/zenpk/dorm-system/internal/cache"
	"github.com/zenpk/dorm-system/internal/controller"
	"github.com/zenpk/dorm-system/internal/dal"
	"github.com/zenpk/dorm-system/internal/mq"
	"github.com/zenpk/dorm-system/internal/rpc"
	"github.com/zenpk/dorm-system/pkg/viperpkg"
	"github.com/zenpk/dorm-system/pkg/zap"
	"log"
)

func main() {
	// read mode from commandline, default as development
	mode := flag.String("mode", "development", "define program mode")
	flag.Parse()
	// Viper
	if err := viperpkg.InitGlobalConfig(*mode); err != nil {
		log.Fatalf("failed to initialize Viper: %v", err)
	}
	// zap
	if err := zap.InitLogger(*mode); err != nil {
		log.Fatalf("failed to initialize zap: %v", err)
	}
	defer zap.Logger.Sync()
	// GORM
	if err := dal.InitDB(); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	// Redis
	if err := cache.InitRedis(); err != nil {
		log.Fatalf("failed to initialize Redis: %v", err)
	}
	defer cache.Redis.Close()
	if err := cache.Warming(); err != nil {
		log.Fatalf("failed to warming Redis: %v", err)
	}
	// RPC connections
	connList, err := rpc.InitClient()
	if err != nil {
		log.Fatalf("failed to initialize RPC clients: %v", err)
	}
	for _, conn := range connList {
		defer conn.Close()
	}
	// Kafka
	if err := mq.InitProducer(); err != nil {
		log.Fatalf("failed to init Kafka producer: %v", err)
	}
	// Gin
	if err := controller.InitGin(); err != nil {
		log.Fatalf("failed to initialize Gin: %v", err)
	}
}
