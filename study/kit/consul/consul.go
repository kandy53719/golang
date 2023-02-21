package consul

import (
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
)

var client *api.Client

func init() {
	var config = api.DefaultConfig()
	config.Address = "127.0.0.1:8500" //consul服务器地址
	//第一步：创建客户端
	tmp, err := api.NewClient(config)
	if err != nil {
		log.Fatal("conful客户端创建失败")
	}
	client = tmp
}

// 注册服务
func ServiceRegister(name string, port int) error {
	return client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s%d", name, port), //consul服务唯一标识
		Name:    name,                            //测试服务
		Address: "127.0.0.1",                     //服务ip
		Port:    port,                            //服务端口号
		Tags:    []string{"primary"},             //分类标签
		Check: &api.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://127.0.0.1:%d/user", port), //api服务地址
			Interval: "10s",                                         //每10秒检查一次
		},
	})
}

// 注销服务
func ServiceDeregister(id string) error {
	return client.Agent().ServiceDeregister(id)
}
