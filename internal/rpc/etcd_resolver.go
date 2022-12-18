package rpc

import (
	"context"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/pkg/zap"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
	"sync"
)

type ResolverBuilder struct {
	etcdClient *clientv3.Client
}

func InitETCDResolverBuilder() (*ResolverBuilder, error) {
	etcdClient, err := initETCDClient()
	if err != nil {
		return nil, err
	}
	rb := new(ResolverBuilder)
	rb.etcdClient = etcdClient
	return rb, nil
}

func (r ResolverBuilder) Close() error {
	return r.etcdClient.Close()
}

func (r ResolverBuilder) Scheme() string {
	return viper.GetString("etcd.scheme")
}

func (r ResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	prefix := target.URL.Scheme + "://" + target.URL.Host + target.URL.Path
	res, err := r.etcdClient.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	er := &etcdResolver{
		etcdClient: r.etcdClient,
		cc:         cc,
		ctx:        ctx,
		cancel:     cancelFunc,
		scheme:     target.URL.Scheme,
	}
	for _, kv := range res.Kvs {
		er.store(kv.Key, kv.Value)
	}
	if err := er.updateState(); err != nil {
		return nil, err
	}
	go er.watcher()
	return er, nil
}

type etcdResolver struct {
	etcdClient *clientv3.Client
	cc         resolver.ClientConn
	ctx        context.Context
	cancel     context.CancelFunc
	scheme     string
	ipPool     sync.Map
}

func (e *etcdResolver) ResolveNow(resolver.ResolveNowOptions) {
}

func (e *etcdResolver) Close() {
	e.cancel()
}

func (e *etcdResolver) watcher() {
	watchChan := e.etcdClient.Watch(e.ctx, e.scheme)
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
		name, ok := k.(string)
		if !ok {
			return false
		}
		addr, ok := v.(string)
		if !ok {
			return false
		}
		addrList.Addresses = append(addrList.Addresses, resolver.Address{Addr: addr, ServerName: name})
		return true
	})
	return e.cc.UpdateState(addrList)
}
