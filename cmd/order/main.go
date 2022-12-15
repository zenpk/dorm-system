package main

import (
	"flag"
	"fmt"
	"github.com/zenpk/dorm-system/internal/cache"
	"github.com/zenpk/dorm-system/internal/dal"
	"github.com/zenpk/dorm-system/internal/mq"
	pb "github.com/zenpk/dorm-system/internal/service/order"
	"github.com/zenpk/dorm-system/pkg/viperpkg"
	"github.com/zenpk/dorm-system/pkg/zap"
	"google.golang.org/grpc"
	"log"
	"net"
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
	orderConfig, err := viperpkg.InitConfig("order-" + *mode)
	if err != nil {
		log.Fatalf("failed to initialize specified config, error: %v", err)
	}
	// zap
	if err := zap.InitLogger("order"); err != nil {
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
		log.Fatalf("failed to initialize Redis, error, error: %v", err)
	}
	defer func() {
		if err := cache.Redis.Close(); err != nil {
			log.Fatalf("failed to close Redis connection, error, error: %v", err)
		}
	}()
	// RPC
	server := &pb.Server{Config: orderConfig}
	addr := fmt.Sprintf("%s:%d", server.Config.GetString("server.host"), server.Config.GetInt("server.port"))
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to initialize TCP listener, error: %v", err)
	}
	go func() {
		grpcServer := grpc.NewServer()
		pb.RegisterOrderServer(grpcServer, server)
		zap.Logger.Infof("server listening at %v", addr)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve, error: %v", err)
		}
	}()
	// Kafka
	zap.Logger.Infof("order consumer is subscribed")
	if err := mq.Consumer.Order.Init(orderConfig, server); err != nil { // defer Close() is inside Init()
		log.Fatalf("failed to initialize consumer, error: %v", err)
	}
}
