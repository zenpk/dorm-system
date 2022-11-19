package main

import (
	"flag"
	"github.com/zenpk/dorm-system/internal/dal"
	"github.com/zenpk/dorm-system/internal/mq"
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
	if err := viperpkg.InitGlobalConfig(*mode); err != nil {
		log.Fatalf("failed to initialize Viper: %v", err)
	}
	orderConfig, err := viperpkg.InitConfig("order")
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
	zap.Logger.Infof("order consumer is subscribed")
	if err := mq.Consumer.Order.Init(orderConfig); err != nil {
		log.Fatalf("failed to initialize consumer: %v", err)
	}
}
