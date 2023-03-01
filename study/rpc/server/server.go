package main

import (
	"fmt"
	"golang/study/rpc/api"
	"net"
	"net/rpc"
)

func main() {
	//1.建立监听端口
	l, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println("服务端口绑定失败：", err)
		return
	} else {
		defer l.Close()
	}

	// 2.注册rpc服务
	//err2 := rpc.RegisterName("hello", new(api.Hello)) //单值请求及返回
	err2 := rpc.RegisterName("hello", new(api.Hello))
	if err2 != nil {
		fmt.Println("rpc服务注册失败", err2)
		return
	}

	for {
		// 3.启动监听
		c, err3 := l.Accept()
		if err3 != nil {
			fmt.Println("请求连接失败", err3)
		}

		// 4.将rpc服务绑定到监听连接
		go rpc.ServeConn(c)
	}
}
