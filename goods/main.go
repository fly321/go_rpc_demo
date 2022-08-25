package main

import (
	"fmt"
	"net"
	"net/rpc"
)

func main() {
	// 注册服务
	err := rpc.RegisterName("goods", new(Goods))
	if err != nil {
		fmt.Println(err)
	}

	listen, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println(err)
	}
	// 关闭监听
	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(listen)

	for {
		fmt.Println("准备建立链接")
		// 4.监听客户端的链接
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
		}

		// 绑定服务
		rpc.ServeConn(accept)
	}
}

type Goods struct {
}

func (receiver Goods) AddGoods(req AddGoodsReq, res *AddGoodsRes) error {
	fmt.Println(req)
	*res = AddGoodsRes{
		Success: true,
		Message: "增加数据成功",
	}
	return nil
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
