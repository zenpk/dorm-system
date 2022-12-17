package rpc

import (
	"github.com/spf13/viper"
	pb "github.com/zenpk/dorm-system/internal/service/user"
	"github.com/zenpk/dorm-system/internal/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type User struct {
	config *viper.Viper
	client pb.UserClient
}

func (u *User) initClient(config *viper.Viper) (*grpc.ClientConn, error) {
	u.config = config
	conn, err := grpc.Dial(u.config.GetString("etcd.target"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	u.client = pb.NewUserClient(conn)
	return conn, nil
}

func (u *User) Register(req *pb.RegisterLoginRequest) (*pb.UserReply, error) {
	ctx, cancel := util.ContextWithTimeout(u.config.GetInt("timeout"))
	defer cancel()
	resp, err := u.client.Register(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *User) Login(req *pb.RegisterLoginRequest) (*pb.UserReply, error) {
	ctx, cancel := util.ContextWithTimeout(u.config.GetInt("timeout"))
	defer cancel()
	resp, err := u.client.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *User) Get(req *pb.GetRequest) (*pb.GetReply, error) {
	ctx, cancel := util.ContextWithTimeout(u.config.GetInt("timeout"))
	defer cancel()
	resp, err := u.client.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *User) Edit(req *pb.EditRequest) (*pb.EditReply, error) {
	ctx, cancel := util.ContextWithTimeout(u.config.GetInt("timeout"))
	defer cancel()
	resp, err := u.client.Edit(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
