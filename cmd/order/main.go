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
	mode = flag.String("mode", "dev", "define program mode")
)

func main() {
	flag.Parse()
	// Viper
	if err := viperpkg.InitGlobalConfig("global-" + *mode); err != nil {
		log.Fatalf("failed to initialize Viper: %v", err)
	}
	orderConfig, err := viperpkg.InitConfig("order-" + *mode)
	if err != nil {
		log.Fatalf("failed to initialize specified config: %v", err)
	}
	// zap
	if err := zap.InitLogger("order"); err != nil {
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
	// RPC
	server := &pb.Server{Config: orderConfig}
	listenAddr := fmt.Sprintf("%s:%d", server.Config.GetString("server.host"), server.Config.GetInt("server.port"))
	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("failed to initialize TCP listener: %v", err)
	}
	go func() {
		grpcServer := grpc.NewServer()
		pb.RegisterOrderServer(grpcServer, server)
		zap.Logger.Infof("server listening at %v", listenAddr)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	// Kafka
	zap.Logger.Infof("order consumer is subscribed")
	if err := mq.Consumer.Order.Init(orderConfig, server); err != nil { // defer Close() is inside Init()
		log.Fatalf("failed to initialize consumer: %v", err)
	}
}
