package utils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/hashicorp/consul/api"
)

func OnExit(serviceId string) {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT)
	<-quit
	//注销
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"
	client, _ := api.NewClient(cfg)
	client.Agent().ServiceDeregister(serviceId)
	fmt.Println("SIGINT service deregister...")
}
