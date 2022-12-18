package rpc

import (
	"context"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/pkg/zap"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

// initEtcdClient initialize etcd for service registry and discovery
func initEtcdClient() (*clientv3.Client, error) {
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   viper.GetStringSlice("etcd.endpoints"),
		DialTimeout: time.Duration(viper.GetInt("etcd.dial_timeout")) * time.Second,
	})
	if err != nil {
		return nil, err
	}
	return etcdClient, nil
}

type EtcdRegister struct {
	etcdClient *clientv3.Client
	leaseID    clientv3.LeaseID
	ctx        context.Context
	cancel     context.CancelFunc
}

func InitEtcdRegister() (*EtcdRegister, error) {
	client, err := initEtcdClient()
	if err != nil {
		return nil, err
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	er := &EtcdRegister{
		etcdClient: client,
		ctx:        ctx,
		cancel:     cancelFunc,
	}
	return er, nil
}

func (e *EtcdRegister) RegisterServer(serviceName, addr string, expire int64) error {
	if err := e.createLease(expire); err != nil {
		return err
	}
	if err := e.bindLease(serviceName, addr); err != nil {
		return err
	}
	keepAliveChan, err := e.keepAlive()
	if err != nil {
		return err
	}
	go e.watcher(keepAliveChan)
	return nil
}

func (e *EtcdRegister) createLease(ttl int64) error {
	res, err := e.etcdClient.Grant(e.ctx, ttl)
	if err != nil {
		return err
	}
	e.leaseID = res.ID
	return nil
}

func (e *EtcdRegister) bindLease(key string, value string) error {
	_, err := e.etcdClient.Put(e.ctx, key, value, clientv3.WithLease(e.leaseID))
	return err
}

func (e *EtcdRegister) keepAlive() (<-chan *clientv3.LeaseKeepAliveResponse, error) {
	resChan, err := e.etcdClient.KeepAlive(e.ctx, e.leaseID)
	if err != nil {
		return nil, err
	}
	return resChan, nil
}

func (e *EtcdRegister) watcher(resChan <-chan *clientv3.LeaseKeepAliveResponse) {
	for {
		select {
		case _, ok := <-resChan:
			// keeping alive
			if !ok {
				zap.Logger.Error("can't keep service alive, revoking")
				if err := e.Close(); err != nil {
					zap.Logger.Error(err)
				}
			}
		case <-e.ctx.Done():
			return
		}
	}
}

func (e *EtcdRegister) Close() error {
	e.cancel()
	if _, err := e.etcdClient.Revoke(e.ctx, e.leaseID); err != nil {
		return err
	}
	return e.etcdClient.Close()
}
