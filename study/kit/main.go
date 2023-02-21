package main

import (
	"errors"
	"flag"
	"fmt"
	"golang/study/kit/client"
	"golang/study/kit/consul"
	"golang/study/kit/server"
	"log"

	"github.com/afex/hystrix-go/hystrix"
)

func main() {

	//client.DirectClient()
	//client.ConsulClient()
	StartServer()
}

func StartServer() {
	var errChan = make(chan error)
	//第一步：从命令行启动服务
	var name = flag.String("name", "UserService", "服务名")
	var port = flag.Int("port", 8080, "端口号")
	flag.Parse()
	//放协程避免程序结束后关闭
	go func() {
		errChan <- server.Server(*port)
	}()
	//第二步：将服务发布到consul上
	go func() {
		consul.ServiceDeregister(fmt.Sprintf("%s%d", *name, *port))
		consul.ServiceRegister(*name, *port)
	}()
	//用来同步
	log.Println(<-errChan)
}

func StartClient() {
	var config = hystrix.CommandConfig{
		Timeout:                3000, //超时时间
		MaxConcurrentRequests:  100,  //最大连接数
		RequestVolumeThreshold: 10,   //请求熔断阈值
		ErrorPercentThreshold:  50,   //请求熔断阈值占比
		SleepWindow:            60,   //每分钟判断一次是否继续熔断继续
	}
	hystrix.ConfigureCommand("ConsulClient", config) //关联
	var err = hystrix.Go("ConsulClient", func() error {
		//正式请求
		client.ConsulClient()
		return errors.New("")
	}, func(e error) error {
		//正式请求失败后服务降级
		return errors.New("")
	})
	if err != nil {
		fmt.Println(err)
	}
}
