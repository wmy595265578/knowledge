package main

import (
	"fmt"

	"github.com/coreos/etcd/clientv3"

	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:            []string{"127.0.0.1:2379"},
		AutoSyncInterval:     0,
		DialTimeout:          5 * time.Second,
		DialKeepAliveTime:    0,
		DialKeepAliveTimeout: 0,
		MaxCallSendMsgSize:   0,
		MaxCallRecvMsgSize:   0,
		TLS:                  nil,
		Username:             "",
		Password:             "",
		RejectOldCluster:     false,
		DialOptions:          nil,
		LogConfig:            nil,
		Context:              nil,
		PermitWithoutStream:  false,
	})

	if err != nil {
		fmt.Println("conn etcd failed,err :%s ", err)
	}
	fmt.Println("conn etcd successful")
	defer cli.Close()
}
