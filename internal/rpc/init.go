package rpc

import (
	"github.com/zenpk/dorm-system/pkg/gmp"
	"github.com/zenpk/dorm-system/pkg/viperpkg"
	"google.golang.org/grpc"
)

// InitClients initialize RPC clients and return all connections
func InitClients(mode string) ([]*grpc.ClientConn, error) {
	connList := make([]*grpc.ClientConn, 0)
	path, err := gmp.GetModPath()
	if err != nil {
		return nil, err
	}
	path += "configs"

	// dorm
	dormConfig, err := viperpkg.InitConfig("dorm-" + mode)
	if err != nil {
		return nil, err
	}
	dormConn, err := Client.Dorm.initClient(dormConfig)
	if err != nil {
		return nil, err
	}
	connList = append(connList, dormConn)

	// order
	orderConfig, err := viperpkg.InitConfig("order-" + mode)
	if err != nil {
		return nil, err
	}
	orderConn, err := Client.Order.initClient(orderConfig)
	if err != nil {
		return nil, err
	}
	connList = append(connList, orderConn)

	// team
	teamConfig, err := viperpkg.InitConfig("team-" + mode)
	if err != nil {
		return nil, err
	}
	teamConn, err := Client.Team.initClient(teamConfig)
	if err != nil {
		return nil, err
	}
	connList = append(connList, teamConn)

	// token
	tokenConfig, err := viperpkg.InitConfig("token-" + mode)
	if err != nil {
		return nil, err
	}
	tokenConn, err := Client.Token.initClient(tokenConfig)
	if err != nil {
		return nil, err
	}
	connList = append(connList, tokenConn)

	// user
	userConfig, err := viperpkg.InitConfig("user-" + mode)
	if err != nil {
		return nil, err
	}
	userConn, err := Client.User.initClient(userConfig)
	if err != nil {
		return nil, err
	}
	connList = append(connList, userConn)

	return connList, nil
}
