package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
	etcdv3 "github.com/coreos/etcd/clientv3"
)

type EtcdClient struct {
	client *etcdv3.Client
}

var (
	etcdClient *EtcdClient
)

func initEtcd(addr, key, localIP string) error {
	cli, err := etcdv3.New(etcdv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:2389", "localhost:2381"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logs.Error("connect etcd failed, err:", err)
		return err
	}

	etcdClient = &EtcdClient{
		client: cli,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if strings.HasSuffix(key, "/") == false {
		key = key + "/"
	}
	etcdKey := fmt.Sprintf("%s%s", key, localIP)
	resp, err := cli.Get(ctx, etcdKey)

	for k, v := range resp.Kvs {
		fmt.Println(k, v)
	}
	return nil
}
