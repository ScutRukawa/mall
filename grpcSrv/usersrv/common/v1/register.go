package v1

import (
	"github.com/hashicorp/consul/api"
)

func Register(addr string, port int, name string, tags []string, id string) error {
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"

	client, err := api.NewClient(cfg)

	if err != nil {
		panic(err)
	}
	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	// check := &api.AgentServiceCheck{
	// 	GRPC:                           "http://192.168.0.112:8088/grpc_health_v1",
	// 	Timeout:                        "5s",
	// 	Interval:                       "5s",
	// 	DeregisterCriticalServiceAfter: "10s",
	// }
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = addr
	// registration.Check = check
	//健康检查参数设置

	if err = client.Agent().ServiceRegister(registration); err != nil {
		panic(err)
	}
	return nil

}

