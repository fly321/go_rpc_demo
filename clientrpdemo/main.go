package main

import (
	"clientrpdemo/proto/greeter"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 建立rpc链接
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	// 关闭链接
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	// 注册客户端
	client := greeter.NewGreeterClient(conn)
	res, err := client.SayHello(context.Background(), &greeter.HelloRequest{Name: "张三"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", res)
}
