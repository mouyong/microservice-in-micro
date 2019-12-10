package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/util/log"
	"net/http"
	"microservice-in-micro/basic"
	"microservice-in-micro/basic/config"

	"github.com/micro/go-micro/web"
	"microservice-in-micro/user-web/handler"
)

func main() {
	basic.Init()

	micReg := etcd.NewRegistry(registryOptions)

	// create new web service
	service := web.NewService(
		web.Name("mu.micro.book.web.user"),
		web.Version("latest"),
		web.Registry(micReg),
		web.Address(":8088"),
	)

	// initialise service
	if err := service.Init(web.Action(func(c *cli.Context) {
		handler.Init()
	})); err != nil {
		log.Fatal(err)
	}

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/user/login", handler.Login)
	service.HandleFunc("/user/logout", handler.Logout)
	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
