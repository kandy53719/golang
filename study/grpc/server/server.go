package main

import (
	"fmt"
	"golang/study/grpc/server/services"
	"net"

	"google.golang.org/grpc"
)

func main() {
	//建立服务端
	var server = grpc.NewServer()
	//注册服务
	services.RegisterProductServiceServer(server, new(services.ProductService))
	//server.RegisterService(services.ProductService, )
	//启动监听
	listener, _ := net.Listen("tcp", ":8080")
	server.Serve(listener)

	go func() {
		fmt.Println("")
	}()
}
