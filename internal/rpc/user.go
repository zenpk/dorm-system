package rpc

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	pb "github.com/zenpk/dorm-system/internal/service/user"
	"github.com/zenpk/dorm-system/internal/util"
	"github.com/zenpk/dorm-system/pkg/zap"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type User struct {
	config  *viper.Viper
	client  pb.UserClient
	leaseId clientv3.LeaseID
}

func (u *User) initClient(config *viper.Viper) (*grpc.ClientConn, error) {
	u.config = config
	conn, err := grpc.Dial(u.config.GetString("etcd.target"), grpc.WithResolvers(EtcdResolver), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	u.client = pb.NewUserClient(conn)
	return conn, nil
}

func (u *User) ServiceRegistry(config *viper.Viper) error {
	u.config = config
	// create a new lease
	resp, err := EtcdClient.Grant(context.TODO(), u.config.GetInt64("etcd.ttl"))
	if err != nil {
		return err
	}
	u.leaseId = resp.ID
	// register
	target := u.config.GetString("etcd.target")
	addr := fmt.Sprintf("%s:%d", u.config.GetString("server.target"), u.config.GetInt("server.port"))
	if _, err = EtcdClient.Put(context.TODO(), target, addr, clientv3.WithLease(u.leaseId)); err != nil {
		return err
	}
	chanKeepAlive, err := EtcdClient.KeepAlive(context.TODO(), u.leaseId)
	if err != nil {
		return err
	}
	go func() {
		for {
			select {
			case _, ok := <-chanKeepAlive:
				if !ok {
					if err := u.ServiceRevoke(); err != nil {
						zap.Logger.Errorf("failed to revoke service from ETCD, error: %v", err)
					}
					return
				}
			}
		}
	}()
	return nil
}

func (u *User) ServiceRevoke() error {
	_, err := EtcdClient.Revoke(context.TODO(), u.leaseId)
	return err
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
