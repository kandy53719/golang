package main

import (
	"context"
	"fmt"
	"golang/study/grpc/server/services"
	"log"

	"google.golang.org/grpc"
)

func main() {
	//建立连接
	conn, err := grpc.Dial(":8080", grpc.WithInsecure()) //不适用证书的连接
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	//创建客户端
	client := services.NewProductServiceClient(conn)
	//远程调用服务
	res, err2 := client.GetProductStock(context.Background(), &services.ProductRequest{ProductId: 12})
	if err2 != nil {
		log.Fatal(err2)
	} else {
		fmt.Println(res.ProductStock)
	}
}
