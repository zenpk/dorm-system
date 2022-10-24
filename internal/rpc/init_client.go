package rpc

import (
	"github.com/zenpk/dorm-system/pkg/config"
	"github.com/zenpk/dorm-system/pkg/gmp"
	"google.golang.org/grpc"
)

type ClientSet struct {
	User User
}

var Client *ClientSet

func InitClient() ([]*grpc.ClientConn, error) {
	connList := make([]*grpc.ClientConn, 0)
	path, err := gmp.GetModPath()
	if err != nil {
		return nil, err
	}
	path += "configs"

	// user
	userConfig, err := config.InitConfig("user")
	if err != nil {
		return nil, err
	}
	userConn, err := Client.User.initUser(userConfig)
	if err != nil {
		return nil, err
	}
	connList = append(connList, userConn)

	return connList, nil
}
