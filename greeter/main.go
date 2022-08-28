package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)
import "greeter/proto/greeter"

type Hello struct{}

func (receiver Hello) SayHello(ctx context.Context, request *greeter.HelloRequest) (*greeter.HelloReply, error) {
	fmt.Println(request)
	return &greeter.HelloReply{Message: "Hello " + request.Name}, nil
}

func main() {
	// 初始化grpc对象
	grpcServer := grpc.NewServer()
	// 注册服务
	greeter.RegisterGreeterServer(grpcServer, &Hello{})
	// 监听端口
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	// 启动服务
	grpcServer.Serve(listen)
}
