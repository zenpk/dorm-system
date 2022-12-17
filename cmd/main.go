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

var (
	mode = flag.String("mode", "development", "define program mode")
)

func main() {
	flag.Parse()
	// Viper
	if err := viperpkg.InitGlobalConfig("global-" + *mode); err != nil {
		log.Fatalf("failed to initialize Viper, error: %v", err)
	}
	// zap
	if err := zap.InitLogger(*mode); err != nil {
		log.Fatalf("failed to initialize zap, error: %v", err)
	}
	defer func() {
		if err := zap.Logger.Sync(); err != nil {
			log.Fatalf("failed to close zap, error: %v", err)
		}
	}()
	// GORM
	if err := dal.InitDB(); err != nil {
		log.Fatalf("failed to initialize database, error: %v", err)
	}
	// Redis
	if err := cache.InitRedis(); err != nil {
		log.Fatalf("failed to initialize Redis, error: %v", err)
	}
	defer func() {
		if err := cache.Redis.Close(); err != nil {
			log.Fatalf("failed to close Redis connection, error: %v", err)
		}
	}()
	if err := cache.Warming(); err != nil {
		log.Fatalf("failed to warming Redis, error: %v", err)
	}
	// ETCD
	if err := rpc.InitETCDResolver(); err != nil {
		log.Fatalf("failed to initialize ETCD client, error: %v", err)
	}
	// RPC connections
	connList, err := rpc.InitClients(*mode)
	if err != nil {
		log.Fatalf("failed to initialize RPC clients, error: %v", err)
	}
	defer func() {
		for _, conn := range connList {
			if err := conn.Close(); err != nil {
				log.Fatalf("failed to close RPC connections, error: %v", err)
			}
		}
	}()
	// Kafka
	producers, err := mq.InitMQ(*mode)
	if err != nil {
		log.Fatalf("failed to init Kafka, error: %v", err)
	}
	defer func() {
		if err := mq.ClusterAdmin.Close(); err != nil {
			log.Fatalf("failed to close Kafka connection, error: %v", err)
		}
		for _, p := range producers {
			if err := p.Close(); err != nil {
				log.Fatalf("failed to close Kafka producer: %v, error: %v", p, err)
			}
		}
	}()
	// Gin
	if err := controller.InitGin(); err != nil {
		log.Fatalf("failed to initialize Gin, error: %v", err)
	}
}
