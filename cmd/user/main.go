package main

import (
	"flag"
	"fmt"
	"github.com/zenpk/dorm-system/internal/dal"
	pb "github.com/zenpk/dorm-system/internal/service/user"
	"github.com/zenpk/dorm-system/pkg/config"
	"github.com/zenpk/dorm-system/pkg/zap"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	mode := flag.String("mode", "development", "define program mode")
	flag.Parse()
	// Viper
	if err := config.InitGlobalConfig(*mode); err != nil {
		log.Fatalf("failed to initialize Viper: %v", err)
	}
	// specified config
	userServer := new(pb.Server)
	var err error
	userServer.Config, err = config.InitConfig("user")
	if err != nil {
		log.Fatalf("failed to initialize specified config: %v", err)
	}
	// zap
	if err := zap.InitLogger("user"); err != nil {
		log.Fatalf("failed to initialize zap: %v", err)
	}
	defer zap.Logger.Sync()
	// GORM
	if err := dal.InitDB(); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", userServer.Config.GetInt("server.port")))
	if err != nil {
		log.Fatalf("failed to initialize TCP listener: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterUserServer(server, userServer)
	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
