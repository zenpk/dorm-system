package main

import (
	"flag"
	"fmt"
	"github.com/zenpk/dorm-system/internal/cache"
	"github.com/zenpk/dorm-system/internal/dal"
	"github.com/zenpk/dorm-system/internal/rpc"
	pb "github.com/zenpk/dorm-system/internal/service/dorm"
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
	// specified config
	server := new(pb.Server)
	var err error
	server.Config, err = viperpkg.InitConfig("dorm-" + *mode)
	if err != nil {
		log.Fatalf("failed to initialize specified config: %v", err)
	}
	// zap
	if err := zap.InitLogger("dorm"); err != nil {
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
	// ETCD
	etcdRegister, err := rpc.InitETCDRegister()
	if err != nil {
		log.Fatalf("failed to initialize ETCD: %v", err)
	}
	etcdTarget := server.Config.GetString("etcd.Target")
	etcdAddr := fmt.Sprintf("%s:%d", server.Config.GetString("server.target"), server.Config.GetInt("server.port"))
	ttl := server.Config.GetInt64("etcd.ttl")
	if err := etcdRegister.RegisterServer(etcdTarget, etcdAddr, ttl); err != nil {
		log.Fatalf("failed to register ETCD: %v", err)
	}
	// RPC
	listenAddr := fmt.Sprintf("%s:%d", server.Config.GetString("server.host"), server.Config.GetInt("server.port"))
	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("failed to initialize TCP listener: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterDormServer(grpcServer, server)
	zap.Logger.Infof("server listening at %v", listenAddr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
