package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
	"gomicro_example/part3/basic"
	"gomicro_example/part3/basic/config"
	"gomicro_example/part3/order-web/handler"
	"net/http"
	"time"
)

func main() {
	// 初始化配置
	basic.Init()

	// 使用consul注册
	micReg := consul.NewRegistry(registryOptions)

	// 创建新服务
	service := web.NewService(
		web.Name("mu.micro.book.web.orders"),
		web.Version("latest"),
		web.Registry(micReg),
		web.Address(":8091"),
	)

	// 初始化服务
	if err := service.Init(
		web.Action(
			func(c *cli.Context) {
				// 初始化handler
				handler.Init()
			}),
	); err != nil {
		log.Fatal(err)
	}

	// 新建订单接口
	authHandler := http.HandlerFunc(handler.New)
	service.Handle("/orders/new", handler.AuthWrapper(authHandler))
	service.HandleFunc("/", handler.Hello)

	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	consulCfg := config.GetConsulConfig()
	ops.Timeout = time.Second * 5
	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.GetHost(), consulCfg.GetPort())}
}
