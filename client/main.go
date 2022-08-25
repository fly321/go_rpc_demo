package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 1. 用rpc.dial 和rpc 微服务建立链接
	connection, err := rpc.Dial("tcp", "127.0.0.1:8080")

	if err != nil {
		panic(err)
	}

	defer func(connection *rpc.Client) {
		err := connection.Close()
		if err != nil {
			panic(err)
		}
	}(connection)
	var reply string
	// 服务名.方法名
	err1 := connection.Call("helloServer.SayHello", "我是客户端", &reply)
	if err != nil {
		panic(err1)
	}
	fmt.Printf("%#v", reply)

}
