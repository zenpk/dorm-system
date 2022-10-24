package order

import (
	"flag"
	"fmt"
	"github.com/zenpk/dorm-system/internal/dal"
	pb "github.com/zenpk/dorm-system/internal/service/order"
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
	orderServer := new(pb.Server)
	var err error
	orderServer.Config, err = viperpkg.InitConfig("order")
	if err != nil {
		log.Fatalf("failed to initialize specified config: %v", err)
	}
	// zap
	if err := zap.InitLogger("order"); err != nil {
		log.Fatalf("failed to initialize zap: %v", err)
	}
	defer zap.Logger.Sync()
	// GORM
	if err := dal.InitDB(); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	addr := fmt.Sprintf("%s:%d", orderServer.Config.GetString("server.host"), orderServer.Config.GetInt("server.port"))
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to initialize TCP listener: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterOrderServer(server, orderServer)
	zap.Logger.Infof("server listening at %v", addr)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
