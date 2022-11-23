package rpc

import (
	"fmt"
	"github.com/spf13/viper"
	pb "github.com/zenpk/dorm-system/internal/service/dorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Dorm struct {
	config *viper.Viper
	client pb.DormClient
}

func (d *Dorm) init(config *viper.Viper) (*grpc.ClientConn, error) {
	d.config = config
	addr := fmt.Sprintf("%s:%d", config.GetString("server.target"), config.GetInt("server.port"))
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	d.client = pb.NewDormClient(conn)
	return conn, nil
}

func (d *Dorm) GetRemainCnt(req *pb.EmptyRequest) (*pb.MapReply, error) {
	ctx, cancel := createCtx(d.config.GetInt("timeout"))
	defer cancel()
	resp, err := d.client.GetRemainCnt(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
