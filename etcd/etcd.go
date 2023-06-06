package etcd

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
)

func EtcdPull(endpoints []string, serviceName string) *grpc.ClientConn {
	// 获取etcd服务
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: endpoints,
		DialTimeout: time.Second << 2,
	})
	//etcdClient, err := clientv3.NewFromURL(etcdUrl1)
	if err != nil {
		panic(err)
	}

	etcdResolver, err := resolver.NewBuilder(etcdClient)
	if err != nil {
		panic(err)
	}

	// 与etcd建立连接
	conn, err := grpc.Dial("etcd:///" + serviceName, grpc.WithResolvers(etcdResolver), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))

	if err != nil {
		panic(err)
	}

	// watch
	go func() {
		c := etcdClient.Watch(context.Background(), serviceName)
		for resp := range c {
			for _, evt := range resp.Events {
				fmt.Printf("Type: %s Key:%s Value:%s\n", evt.Type, evt.Kv.Key, evt.Kv.Value)
			}
		}
	}()

	return conn
}