package rpc

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/pkg/zap"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
	"sync"
	"time"
)

// initETCDClient initialize ETCD for service registry and discovery
func initETCDClient() (*clientv3.Client, error) {
	endpoint := fmt.Sprintf("%s:%d", viper.GetString("etcd.host"), viper.GetInt("etcd.port"))
	var err error
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{endpoint},
		DialTimeout: time.Duration(viper.GetInt("etcd.dial_timeout")) * time.Second,
	})
	if err != nil {
		return nil, err
	}
	return etcdClient, nil
}

var ETCDResolver etcdResolverBuilder

type etcdResolverBuilder struct {
	etcdClient *clientv3.Client
}

func InitETCDResolver() error {
	var err error
	ETCDResolver.etcdClient, err = initETCDClient()
	resolver.Register(ETCDResolver)
	return err
}

func (e etcdResolverBuilder) Close() error {
	return e.etcdClient.Close()
}

func (e etcdResolverBuilder) Scheme() string {
	return viper.GetString("etcd.scheme")
}

func (e etcdResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	res, err := e.etcdClient.Get(context.Background(), target.URL.Scheme)
	if err != nil {
		return nil, err
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	es := &etcdResolver{
		cc:         cc,
		etcdClient: e.etcdClient,
		ctx:        ctx,
		cancel:     cancelFunc,
		scheme:     target.URL.Scheme,
	}
	for _, kv := range res.Kvs {
		es.store(kv.Key, kv.Value)
	}
	if err := es.updateState(); err != nil {
		return nil, err
	}
	go es.watcher()
	return es, nil
}

type etcdResolver struct {
	ctx        context.Context
	cancel     context.CancelFunc
	cc         resolver.ClientConn
	etcdClient *clientv3.Client
	scheme     string
	ipPool     sync.Map
}

func (e *etcdResolver) ResolveNow(resolver.ResolveNowOptions) {
	zap.Logger.Info("ETCD resolver is resolving")
}

func (e *etcdResolver) Close() {
	e.cancel()
}

func (e *etcdResolver) watcher() {
	watchChan := e.etcdClient.Watch(context.Background(), e.scheme)
	for {
		select {
		case val := <-watchChan:
			for _, event := range val.Events {
				switch event.Type {
				case 0:
					e.store(event.Kv.Key, event.Kv.Value)
					if err := e.updateState(); err != nil {
						zap.Logger.Error(err)
						return
					}
				case 1:
					e.del(event.Kv.Key)
					if err := e.updateState(); err != nil {
						zap.Logger.Error(err)
						return
					}
				}
			}
		case <-e.ctx.Done():
			return
		}
	}
}

func (e *etcdResolver) store(k, v []byte) {
	e.ipPool.Store(string(k), string(v))
}

func (e *etcdResolver) del(key []byte) {
	e.ipPool.Delete(string(key))
}

func (e *etcdResolver) updateState() error {
	var addrList resolver.State
	e.ipPool.Range(func(k, v interface{}) bool {
		tA, ok := v.(string)
		if !ok {
			return false
		}
		addrList.Addresses = append(addrList.Addresses, resolver.Address{Addr: tA})
		return true
	})
	return e.cc.UpdateState(addrList)
}
