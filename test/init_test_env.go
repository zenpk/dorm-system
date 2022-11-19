package test

import (
	"github.com/zenpk/dorm-system/internal/rpc"
	"github.com/zenpk/dorm-system/pkg/viperpkg"
	"github.com/zenpk/dorm-system/pkg/zap"
	"log"
)

func initTestEnv() {
	// Viper
	if err := viperpkg.InitGlobalConfig("testing"); err != nil {
		log.Fatalf("failed to initialize Viper: %v", err)
	}
	// zap
	if err := zap.InitLogger("testing"); err != nil {
		log.Fatalf("failed to initialize zap: %v", err)
	}
	defer zap.Logger.Sync()
	// RPC connections
	connList, err := rpc.InitClients()
	if err != nil {
		log.Fatalf("failed to initialize RPC clients: %v", err)
	}
	for _, conn := range connList {
		defer conn.Close()
	}
}
