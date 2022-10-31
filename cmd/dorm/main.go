package main

import (
	"flag"
	"fmt"
	"github.com/zenpk/dorm-system/internal/cache"
	"github.com/zenpk/dorm-system/internal/dal"
	pb "github.com/zenpk/dorm-system/internal/service/dorm"
	"github.com/zenpk/dorm-system/pkg/viperpkg"
	"github.com/zenpk/dorm-system/pkg/zap"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	mode := flag.String("mode", "development", "define program mode")
	flag.Parse()
	// Viper
	if err := viperpkg.InitGlobalConfig(*mode); err != nil {
		log.Fatalf("failed to initialize Viper: %v", err)
	}
	// specified config
	dormServer := new(pb.Server)
	var err error
	dormServer.Config, err = viperpkg.InitConfig("dorm")
	if err != nil {
		log.Fatalf("failed to initialize specified config: %v", err)
	}
	// zap
	if err := zap.InitLogger("dorm"); err != nil {
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
	addr := fmt.Sprintf("%s:%d", dormServer.Config.GetString("server.host"), dormServer.Config.GetInt("server.port"))
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to initialize TCP listener: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterDormServer(server, dormServer)
	zap.Logger.Infof("server listening at %v", addr)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
