package main

import (
	"flag"
	"fmt"
	"github.com/zenpk/dorm-system/internal/dal"
	pb "github.com/zenpk/dorm-system/internal/service/token"
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
	if err := viperpkg.InitGlobalConfig(*mode); err != nil {
		log.Fatalf("failed to initialize Viper, error: %v", err)
	}
	// specified config
	server := new(pb.Server)
	var err error
	server.Config, err = viperpkg.InitConfig("token")
	if err != nil {
		log.Fatalf("failed to initialize specified config, error: %v", err)
	}
	// zap
	if err := zap.InitLogger("user"); err != nil {
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
	// RPC
	addr := fmt.Sprintf("%s:%d", server.Config.GetString("server.host"), server.Config.GetInt("server.port"))
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to initialize TCP listener, error: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTokenServer(grpcServer, server)
	zap.Logger.Infof("server listening at %v", addr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve, error: %v", err)
	}
}
