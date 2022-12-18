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
	"google.golang.org/grpc/resolver"
	"log"
)

var (
	mode = flag.String("mode", "dev", "define program mode")
)

func main() {
	flag.Parse()
	// Viper
	if err := viperpkg.InitGlobalConfig("global-" + *mode); err != nil {
		log.Fatalf("failed to initialize Viper: %v", err)
	}
	// zap
	if err := zap.InitLogger(*mode); err != nil {
		log.Fatalf("failed to initialize zap: %v", err)
	}
	defer func() {
		if err := zap.Logger.Sync(); err != nil {
			log.Fatalf("failed to close zap: %v", err)
		}
	}()
	// GORM
	if err := dal.InitDB(); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	// Redis
	if err := cache.InitRedis(); err != nil {
		log.Fatalf("failed to initialize Redis: %v", err)
	}
	defer func() {
		if err := cache.Redis.Close(); err != nil {
			log.Fatalf("failed to close Redis connection: %v", err)
		}
	}()
	if err := cache.Warming(); err != nil {
		log.Fatalf("failed to warming Redis: %v", err)
	}
	// etcd
	erb, err := rpc.InitEtcdResolverBuilder()
	if err != nil {
		log.Fatalf("failed to initialize etcd client: %v", err)
	}
	resolver.Register(erb)
	defer func() {
		if err := erb.Close(); err != nil {
			log.Fatalf("failed to close etcd resolver builder: %v", err)
		}
	}()
	// RPC connections
	connList, err := rpc.InitClients(*mode)
	if err != nil {
		log.Fatalf("failed to initialize RPC clients: %v", err)
	}
	defer func() {
		for _, conn := range connList {
			if err := conn.Close(); err != nil {
				log.Fatalf("failed to close RPC connections: %v", err)
			}
		}
	}()
	// Kafka
	producers, err := mq.InitMQ(*mode)
	if err != nil {
		log.Fatalf("failed to init Kafka: %v", err)
	}
	defer func() {
		if err := mq.ClusterAdmin.Close(); err != nil {
			log.Fatalf("failed to close Kafka connection: %v", err)
		}
		for _, p := range producers {
			if err := p.Close(); err != nil {
				log.Fatalf("failed to close Kafka producer: %v: %v", p, err)
			}
		}
	}()
	// Gin
	if err := controller.InitGin(); err != nil {
		log.Fatalf("failed to initialize Gin: %v", err)
	}
}
