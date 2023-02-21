package client

import (
	"context"
	"fmt"
	"golang/study/kit/client/transport"
	"golang/study/kit/server/endpoint"
	"io"
	"os"
	"time"

	"net/url"

	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	transporthttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/hashicorp/consul/api"
)

// 直连模式
func DirectClient() {
	//创建客户端
	url, _ := url.Parse("http://127.0.0.1:8080")
	client := transporthttp.NewClient("GET", url, transport.EncodeRequest, transport.DecodeResponse)
	//生成服务
	userService := client.Endpoint()
	//上下文日志追踪
	ctx := context.Background()
	//开始调用
	response, err := userService(ctx, endpoint.UserRequest{Id: 123})
	if err != nil {
		fmt.Println(err)
	}
	//处理返回
	var user = response.(endpoint.UserResponse)
	fmt.Println(user.Result)
}

// 通过consul获取服务
func ConsulClient() {
	var config = api.DefaultConfig()
	config.Address = "127.0.0.1:8500" //consul服务器地址
	//第一步：创建客户端
	api_client, _ := api.NewClient(config)
	client := consul.NewClient(api_client)
	//第二部创建一个consul的实例
	var logger log.Logger = log.NewLogfmtLogger(os.Stdout)
	var Tag = []string{"primary"}
	instancer := consul.NewInstancer(client, logger, "test", Tag, true)
	factory := func(service_url string) (kitendpoint.Endpoint, io.Closer, error) { //factory定义了如何获得服务端的endpoint,这里的service_url是从consul中读取到的service的address我这里是192.168.3.14:8000
		tart, _ := url.Parse("http://" + service_url)
		return transporthttp.NewClient("GET", tart, transport.EncodeRequest, transport.DecodeResponse).Endpoint(), nil, nil //我再GetUserInfo_Request里面定义了访问哪一个api把url拼接成了http://192.168.3.14:8000/v1/user/{uid}的形式
	}
	endpointer := sd.NewEndpointer(instancer, factory, logger)
	endpoints, _ := endpointer.Endpoints()
	fmt.Println("服务有", len(endpoints), "条")

	var myBalance = lb.NewRoundRobin(endpointer)                //go-kit自带的轮询
	myBalance = lb.NewRandom(endpointer, time.Now().UnixNano()) //go-kit自带的随机
	for {
		//getUserInfo := endpoints[0] //写死获取第一个
		getUserInfo, err := myBalance.Endpoint()
		if err != nil {
			fmt.Println("轮询获取服务失败")
		}
		ctx := context.Background() //第三步：创建一个context上下文对象
		//第四步：执行
		res, err := getUserInfo(ctx, endpoint.UserRequest{Id: 123})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		//第五步：断言，得到响应值
		userinfo := res.(endpoint.UserResponse)
		fmt.Println(userinfo.Result)
	}

}
