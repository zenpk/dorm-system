package test

func initTestEnv() {
	//// Viper
	//if err := viperpkg.InitGlobalConfig("test"); err != nil {
	//	log.Fatalf("failed to initialize Viper: %v", err)
	//}
	//// zap
	//if err := zap.InitLogger("test"); err != nil {
	//	log.Fatalf("failed to initialize zap: %v", err)
	//}
	//defer zap.Logger.Sync()
	//// RPC connections
	//connList, err := rpc.InitClients("dev", nil)
	//if err != nil {
	//	log.Fatalf("failed to initialize RPC clients: %v", err)
	//}
	//for _, conn := range connList {
	//	defer conn.Close()
	//}
}
