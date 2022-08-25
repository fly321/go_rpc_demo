package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 1. 用rpc.dial 和rpc 微服务建立链接
	//connection, err := rpc.Dial("tcp", "127.0.0.1:8080")
	connection, err := net.Dial("tcp", "127.0.0.1:8080")

	if err != nil {
		panic(err)
	}

	defer connection.Close()

	// 建立基于json编码的rpc服务
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(connection))

	var reply string
	// 服务名.方法名
	err2 := client.Call("helloServer.SayHello", "我是客户端", &reply)
	if err2 != nil {
	}
	fmt.Printf("%#v", reply)

}
