package main

import (
	"fmt"
	"golang/study/rpc/api"
	"net/rpc"
)

func main() {
	// 1.建立连接
	c, err := rpc.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println("客户端远程连接建立失败：", err)
	}
	defer c.Close()

	// 2.调用远程服务
	//var res string //用来接收返回值
	//err2 := c.Call("hello.Hello", "这是客户端消息", &res) //单值返回
	var req = api.UserRequest{Name: "张三", Message: "多值测试"}
	var res = api.UserResponse{}
	err2 := c.Call("hello.HelloUser", req, &res)
	if err2 != nil {
		fmt.Println("远程服务调用失败：", err2)
	}

	fmt.Println(res)
}
