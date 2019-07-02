package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/web"
	"github.com/prometheus/common/log"
	"gomicro_example/part1/user-web/basic"
	"gomicro_example/part1/user-web/basic/config"
	"gomicro_example/part1/user-web/handler"
	"time"
)

func main() {
	basic.Init()
	micReg := consul.NewRegistry(func(ops *registry.Options) {
		consulCfg := config.GetConsulConfig()
		ops.Timeout = time.Second * 5
		ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.GetHost(), consulCfg.GetPort())}
	})
	service := web.NewService(
		web.Name("mu.micro.book.web.user"),
		web.Version("latest"),
		web.Registry(micReg),
		web.Address(":8088"),
	)
	if err := service.Init(
		web.Action(func(c *cli.Context) {
			handler.Init()
		}),
	); err != nil {
		log.Fatal(err)
	}
	service.HandleFunc("/user/login", handler.Login)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
