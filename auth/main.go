package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/util/log"
	"microservice-in-micro/auth/handler"
	"microservice-in-micro/auth/model"
	"microservice-in-micro/basic"
	"microservice-in-micro/basic/config"

	auth "microservice-in-micro/auth/proto/auth"
)

func main() {
	basic.Init()

	micReg := etcd.NewRegistry(registryOptions)

	// New Service
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.auth"),
		micro.Version("latest"),
		micro.Registry(micReg),
	)

	// Initialise service
	service.Init(micro.Action(func(c *cli.Context) {
		model.Init()

		handler.Init()
	}))

	// Register Handler
	auth.RegisterServiceHandler(service.Server(), new(handler.Service))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}