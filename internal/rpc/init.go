package rpc

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/pkg/gmp"
	"github.com/zenpk/dorm-system/pkg/viperpkg"
	"go.etcd.io/etcd/client/v3"
	resolverv3 "go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"time"
)

var EtcdClient *clientv3.Client
var EtcdResolver resolver.Builder

// InitETCD initialize ETCD for service registry and discovery
func InitETCD() error {
	endpoint := fmt.Sprintf("%s:%d", viper.GetString("etcd.host"), viper.GetInt("etcd.port"))
	var err error
	EtcdClient, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{endpoint},
		DialTimeout: time.Duration(viper.GetInt("etcd.dial_timeout")) * time.Second,
	})
	if err != nil {
		return err
	}
	EtcdResolver, err = resolverv3.NewBuilder(EtcdClient)
	if err != nil {
		return err
	}
	return nil
}

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
