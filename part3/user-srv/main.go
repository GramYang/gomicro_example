package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"gomicro_example/part3/basic"
	"gomicro_example/part3/basic/config"
	"gomicro_example/part3/user-srv/handler"
	"gomicro_example/part3/user-srv/model"
	s "gomicro_example/part3/user-srv/proto/user"
	"time"
)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()
	//使用consul注册
	micReg := consul.NewRegistry(func(ops *registry.Options) {
		consulCfg := config.GetConsulConfig()
		ops.Timeout = time.Second * 5
		ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.GetHost(), consulCfg.GetPort())}
	})
	// 新建服务
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.user"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			model.Init()
			handler.Init()
		}),
	)

	// 注册handler
	_ = s.RegisterUserHandler(service.Server(), new(handler.Service))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
