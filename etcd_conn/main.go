package main

import (
	"context"
	"fmt"
	//"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3"
)

var (
	cli *clientv3.Client
)

func main() {
	Cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})

	if err != nil {
		fmt.Println("conn etcd failed,err :%s ", err)
	}
	fmt.Println("conn etcd successful")
	cli = Cli
	ctx,cancle := context.WithCancel(context.Background())
	resp,_:=cli.Put(ctx,"wmy","111")
    cancle()
	fmt.Println(resp.Header)
	cli.Get(ctx,"wmy")
	//Get()

}


func Get() {
	ctx,cancle := context.WithCancel(context.Background())
	cancle()
	key,_:=cli.Get(ctx,"wmy")
	for _,v:=range key.Kvs {
		fmt.Print(v.Key,v.Value)
	}
}