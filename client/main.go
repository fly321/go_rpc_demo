package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 1. 用rpc.dial 和rpc 微服务建立链接
	//connection, err := rpc.Dial("tcp", "127.0.0.1:8080")
	connection, err := rpc.Dial("tcp", "127.0.0.1:8081")

	if err != nil {
		panic(err)
	}

	defer func(connection *rpc.Client) {
		err := connection.Close()
		if err != nil {
			panic(err)
		}
	}(connection)
	//var reply string
	var reply AddGoodsRes
	// 服务名.方法名
	//err1 := connection.Call("helloServer.SayHello", "我是客户端", &reply)
	err1 := connection.Call("goods.AddGoods", AddGoodsReq{
		Id:      1,
		Title:   "标题",
		Price:   1.1,
		Content: "我是内容",
	}, &reply)
	if err != nil {
		panic(err1)
	}
	fmt.Printf("%#v", reply)

}

type AddGoodsReq struct {
	Id      int
	Title   string
	Price   float32
	Content string
}

type AddGoodsRes struct {
	Success bool
	Message string
}
