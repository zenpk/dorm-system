package rpc

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	pb "github.com/zenpk/dorm-system/internal/service/dorm"
	"github.com/zenpk/dorm-system/internal/util"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

type Dorm struct {
	config *viper.Viper
	client pb.DormClient
}

func (d *Dorm) initClient(config *viper.Viper) (*grpc.ClientConn, error) {
	d.config = config

	conn, err := grpc.Dial(config.GetString("etcd.target"), grpc.WithResolvers(EtcdResolver))
	if err != nil {
		return nil, err
	}
	d.client = pb.NewDormClient(conn)
	return conn, nil
}

func (d *Dorm) ServerRegistry() error {
	// create a new lease
	resp, err := EtcdClient.Grant(context.TODO(), d.config.GetInt64("etcd.ttl"))
	if err != nil {
		return err
	}
	// register
	target := viper.GetString("etcd.target")
	addr := fmt.Sprintf("%s:%d", d.config.GetString("server.target"), d.config.GetInt("server.port"))
	if _, err = EtcdClient.Put(context.TODO(), target, addr, clientv3.WithLease(resp.ID)); err != nil {
		return err
	}
	chanKeepAlive, err := EtcdClient.KeepAlive(context.TODO(), resp.ID)
	if err != nil {
		return err
	}
	go func() {
		for {
			select {
			case _, ok := <-chanKeepAlive:
				if !ok {
					d.ServerRevoke()
				}
			}
		}
	}()
	return nil
}

func (d *Dorm) ServerRevoke() error {
	return nil
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
