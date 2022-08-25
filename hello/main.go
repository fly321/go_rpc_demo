package main

import (
	"fmt"
	"net"
	"net/rpc"
)

func main() {
	// 注册一个rpc服务
	err1 := rpc.RegisterName("helloServer", new(HelloServer))
	if err1 != nil {
		panic(err1)
	}

	// 监听端口
	listen, err2 := net.Listen("tcp", "127.0.0.1:8080")
	if err2 != nil {
		panic(err2)
	}

	// 应用退出关闭端口
	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {

		}
	}(listen)

	for {
		fmt.Println("开始建立链接")
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		// 绑定服务
		rpc.ServeConn(conn)
	}

}

// 定于i一个远程调用的方法
type HelloServer struct {
}

//SayHello
/**
req 表示客户端传过来的数据
res 表示给客户端返回数据
req res 的类型不能是 :channel , func 均不能序列化
*/
func (that HelloServer) SayHello(req string, res *string) error {
	fmt.Println(req)
	*res = "你好:" + req
	return nil
}
