package rpc

import (
	"github.com/zenpk/dorm-system/pkg/gmp"
	"github.com/zenpk/dorm-system/pkg/viperpkg"
	"google.golang.org/grpc"
)

type ClientSet struct {
	User  User
	Order Order
	Dorm  Dorm
}

var Client ClientSet

func InitClient() ([]*grpc.ClientConn, error) {
	connList := make([]*grpc.ClientConn, 0)
	path, err := gmp.GetModPath()
	if err != nil {
		return nil, err
	}
	path += "configs"

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

	// order
	orderConfig, err := viperpkg.InitConfig("order")
	if err != nil {
		return nil, err
	}
	orderConn, err := Client.Order.init(orderConfig)
	if err != nil {
		return nil, err
	}
	connList = append(connList, orderConn)

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

	return connList, nil
}
