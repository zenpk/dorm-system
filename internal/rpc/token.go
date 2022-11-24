package rpc

import (
	"fmt"
	"github.com/spf13/viper"
	pb "github.com/zenpk/dorm-system/internal/service/token"
	"github.com/zenpk/dorm-system/internal/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Token struct {
	config *viper.Viper
	client pb.TokenClient
}

func (t *Token) init(config *viper.Viper) (*grpc.ClientConn, error) {
	t.config = config
	addr := fmt.Sprintf("%s:%d", config.GetString("server.target"), config.GetInt("server.port"))
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	t.client = pb.NewTokenClient(conn)
	return conn, nil
}

func (t *Token) GenAllToken(req *pb.GenAllTokenRequest) (*pb.TokenReply, error) {
	ctx, cancel := util.ContextWithTimeout(t.config.GetInt("timeout"))
	defer cancel()
	resp, err := t.client.GenAllToken(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *Token) GenAccessToken(req *pb.GenAccessTokenRequest) (*pb.TokenReply, error) {
	ctx, cancel := util.ContextWithTimeout(t.config.GetInt("timeout"))
	defer cancel()
	resp, err := t.client.GenAccessToken(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
