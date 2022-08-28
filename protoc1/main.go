package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"protoc1/proto/userServices"
)

func main() {
	u := &userServices.Userinfo{
		Username: "张三",
		Age:      18,
		Hobby:    []string{"篮球", "足球"},
	}
	fmt.Println(u.GetUsername())

	// protobuf 的序列化
	marshal, _ := proto.Marshal(u)
	fmt.Printf("%v", marshal)
	// protobuf 的反序列化
	user := userServices.Userinfo{}
	err := proto.Unmarshal(marshal, &user)
	if err != nil {
		return
	}
	fmt.Println(user)

}
