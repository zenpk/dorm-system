package rpc

import (
	"fmt"
	"github.com/spf13/viper"
	pb "github.com/zenpk/dorm-system/internal/service/order"
	"github.com/zenpk/dorm-system/internal/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Order struct {
	client pb.OrderClient
	config *viper.Viper
}

func (o *Order) initClient(config *viper.Viper) (*grpc.ClientConn, error) {
	o.config = config
	addr := fmt.Sprintf("%s:%d", config.GetString("server.target"), config.GetInt("server.port"))
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	o.client = pb.NewOrderClient(conn)
	return conn, nil
}

func (o *Order) Get(req *pb.GetRequest) (*pb.GetReply, error) {
	ctx, cancel := util.ContextWithTimeout(o.config.GetInt("timeout.rpc"))
	defer cancel()
	resp, err := o.client.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *Order) Delete(req *pb.DeleteRequest) (*pb.DeleteReply, error) {
	ctx, cancel := util.ContextWithTimeout(o.config.GetInt("timeout.rpc"))
	defer cancel()
	resp, err := o.client.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
