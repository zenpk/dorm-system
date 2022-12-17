package rpc

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	pb "github.com/zenpk/dorm-system/internal/service/dorm"
	"github.com/zenpk/dorm-system/internal/util"
	"github.com/zenpk/dorm-system/pkg/zap"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Dorm struct {
	config  *viper.Viper
	client  pb.DormClient
	leaseId clientv3.LeaseID
}

func (d *Dorm) initClient(config *viper.Viper) (*grpc.ClientConn, error) {
	d.config = config
	conn, err := grpc.Dial(d.config.GetString("etcd.target"), grpc.WithResolvers(EtcdResolver), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("here")
		return nil, err
	}
	d.client = pb.NewDormClient(conn)
	return conn, nil
}

func (d *Dorm) ServiceRegistry(config *viper.Viper) error {
	d.config = config
	// create a new lease
	resp, err := EtcdClient.Grant(context.TODO(), d.config.GetInt64("etcd.ttl"))
	if err != nil {
		return err
	}
	d.leaseId = resp.ID
	// register
	target := d.config.GetString("etcd.target")
	addr := fmt.Sprintf("%s:%d", d.config.GetString("server.target"), d.config.GetInt("server.port"))
	if _, err = EtcdClient.Put(context.TODO(), target, addr, clientv3.WithLease(d.leaseId)); err != nil {
		return err
	}
	chanKeepAlive, err := EtcdClient.KeepAlive(context.TODO(), d.leaseId)
	if err != nil {
		return err
	}
	go func() {
		for {
			select {
			case _, ok := <-chanKeepAlive:
				if !ok {
					if err := d.ServiceRevoke(); err != nil {
						zap.Logger.Errorf("failed to revoke service from ETCD, error: %v", err)
					}
					return
				}
			}
		}
	}()
	return nil
}

func (d *Dorm) ServiceRevoke() error {
	_, err := EtcdClient.Revoke(context.TODO(), d.leaseId)
	return err
}

func (d *Dorm) GetRemainCnt(req *pb.EmptyRequest) (*pb.MapReply, error) {
	ctx, cancel := util.ContextWithTimeout(d.config.GetInt("timeout"))
	defer cancel()
	resp, err := d.client.GetRemainCnt(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *Dorm) GetAll(req *pb.EmptyRequest) (*pb.GetAllReply, error) {
	ctx, cancel := util.ContextWithTimeout(d.config.GetInt("timeout"))
	defer cancel()
	resp, err := d.client.GetAll(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
