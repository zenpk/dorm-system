package rpc

import (
	"fmt"
	"github.com/spf13/viper"
	pb "github.com/zenpk/dorm-system/internal/service/team"
	"github.com/zenpk/dorm-system/internal/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Team struct {
	config *viper.Viper
	client pb.TeamClient
}

func (t *Team) initClient(config *viper.Viper) (*grpc.ClientConn, error) {
	t.config = config
	addr := fmt.Sprintf("%s:%d", config.GetString("server.target"), config.GetInt("server.port"))
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	t.client = pb.NewTeamClient(conn)
	return conn, nil
}

func (t *Team) Create(req *pb.CreateRequest) (*pb.CreateReply, error) {
	ctx, cancel := util.ContextWithTimeout(t.config.GetInt("timeout"))
	defer cancel()
	resp, err := t.client.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *Team) Get(req *pb.GetRequest) (*pb.GetReply, error) {
	ctx, cancel := util.ContextWithTimeout(t.config.GetInt("timeout"))
	defer cancel()
	resp, err := t.client.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *Team) Join(req *pb.JoinRequest) (*pb.JoinReply, error) {
	ctx, cancel := util.ContextWithTimeout(t.config.GetInt("timeout"))
	defer cancel()
	resp, err := t.client.Join(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *Team) Leave(req *pb.LeaveRequest) (*pb.LeaveReply, error) {
	ctx, cancel := util.ContextWithTimeout(t.config.GetInt("timeout"))
	defer cancel()
	resp, err := t.client.Leave(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *Team) Transfer(req *pb.TransferRequest) (*pb.TransferReply, error) {
	ctx, cancel := util.ContextWithTimeout(t.config.GetInt("timeout"))
	defer cancel()
	resp, err := t.client.Transfer(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
