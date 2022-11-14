package rpc

import (
	"github.com/zenpk/dorm-system/pkg/gmp"
	"github.com/zenpk/dorm-system/pkg/viperpkg"
	"google.golang.org/grpc"
)

type ClientSet struct {
	Dorm  Dorm
	Token Token
	User  User
}

var Client ClientSet

// InitClient initialize RPC clients and return all connections
func InitClient() ([]*grpc.ClientConn, error) {
	connList := make([]*grpc.ClientConn, 0)
	path, err := gmp.GetModPath()
	if err != nil {
		return nil, err
	}
	path += "configs"

	// dorm
	dormConfig, err := viperpkg.InitConfig("dorm")
	if err != nil {
		return nil, err
	}
	dormConn, err := Client.Dorm.init(dormConfig)
	if err != nil {
		return nil, err
	}
	connList = append(connList, dormConn)

	// user
	userConfig, err := viperpkg.InitConfig("user")
	if err != nil {
		return nil, err
	}
	userConn, err := Client.User.init(userConfig)
	if err != nil {
		return nil, err
	}
	connList = append(connList, userConn)

	// token
	tokenConfig, err := viperpkg.InitConfig("token")
	if err != nil {
		return nil, err
	}
	tokenConn, err := Client.Token.init(tokenConfig)
	if err != nil {
		return nil, err
	}
	connList = append(connList, tokenConn)

	return connList, nil
}
