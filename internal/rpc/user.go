package rpc

import (
	"fmt"
	"github.com/spf13/viper"
	pb "github.com/zenpk/dorm-system/internal/service/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type User struct {
	client pb.UserClient
	config *viper.Viper
}

func (u *User) init(config *viper.Viper) (*grpc.ClientConn, error) {
	u.config = config
	addr := fmt.Sprintf("%s:%d", config.GetString("server.target"), config.GetInt("server.port"))
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	u.client = pb.NewUserClient(conn)
	return conn, nil
}

func (u *User) Register(req *pb.RegisterLoginRequest) (*pb.UserReply, error) {
	ctx, cancel := createCtx(u.config.GetInt("timeout"))
	defer cancel()
	resp, err := u.client.Register(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *User) Login(req *pb.RegisterLoginRequest) (*pb.UserReply, error) {
	ctx, cancel := createCtx(u.config.GetInt("timeout"))
	defer cancel()
	resp, err := u.client.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
