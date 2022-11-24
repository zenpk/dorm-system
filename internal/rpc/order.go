package rpc

//import (
//	"fmt"
//	"github.com/spf13/viper"
//	pb "github.com/zenpk/dorm-system/internal/service/order"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/credentials/insecure"
//)
//
//type Order struct {
//	client pb.OrderClient
//	config *viper.Viper
//}
//
//func (o *Order) init(config *viper.Viper) (*grpc.ClientConn, error) {
//	o.config = config
//	addr := fmt.Sprintf("%s:%d", config.GetString("server.target"), config.GetInt("server.port"))
//	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
//	if err != nil {
//		return nil, err
//	}
//	o.client = pb.NewOrderClient(conn)
//	return conn, nil
//}
//
//func (o *Order) Submit(req *pb.OrderRequest) (*pb.OrderReply, error) {
//	ctx, cancel := ContextWithTimeout(o.config.GetInt("timeout"))
//	defer cancel()
//	resp, err := o.client.Submit(ctx, req)
//	if err != nil {
//		return nil, err
//	}
//	return resp, nil
//}
