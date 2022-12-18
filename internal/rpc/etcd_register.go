package rpc

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/pkg/zap"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

// initETCDClient initialize ETCD for service registry and discovery
func initETCDClient() (*clientv3.Client, error) {
	endpoint := fmt.Sprintf("%s:%d", viper.GetString("etcd.host"), viper.GetInt("etcd.port"))
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{endpoint},
		DialTimeout: time.Duration(viper.GetInt("etcd.dial_timeout")) * time.Second,
	})
	if err != nil {
		return nil, err
	}
	return etcdClient, nil
}

type ETCDRegister struct {
	etcdClient *clientv3.Client
	leaseID    clientv3.LeaseID
	ctx        context.Context
	cancel     context.CancelFunc
}

func InitETCDRegister() (*ETCDRegister, error) {
	client, err := initETCDClient()
	if err != nil {
		return nil, err
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	er := &ETCDRegister{
		etcdClient: client,
		ctx:        ctx,
		cancel:     cancelFunc,
	}
	return er, nil
}

func (e *ETCDRegister) RegisterServer(serviceName, addr string, expire int64) error {
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
	go e.watcher(serviceName, keepAliveChan)
	return nil
}

func (e *ETCDRegister) createLease(ttl int64) error {
	res, err := e.etcdClient.Grant(e.ctx, ttl)
	if err != nil {
		return err
	}
	e.leaseID = res.ID
	return nil
}

func (e *ETCDRegister) bindLease(key string, value string) error {
	_, err := e.etcdClient.Put(e.ctx, key, value, clientv3.WithLease(e.leaseID))
	return err
}

func (e *ETCDRegister) keepAlive() (<-chan *clientv3.LeaseKeepAliveResponse, error) {
	resChan, err := e.etcdClient.KeepAlive(e.ctx, e.leaseID)
	if err != nil {
		return nil, err
	}
	return resChan, nil
}

func (e *ETCDRegister) watcher(key string, resChan <-chan *clientv3.LeaseKeepAliveResponse) {
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

func (e *ETCDRegister) Close() error {
	e.cancel()
	if _, err := e.etcdClient.Revoke(e.ctx, e.leaseID); err != nil {
		return err
	}
	return e.etcdClient.Close()
}
