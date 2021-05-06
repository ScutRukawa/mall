package main

import (
	"fmt"
	"goodssrv/api/goods"
	v1 "goodssrv/common/v1"
	"goodssrv/global"
	"goodssrv/grpc_health_v1"
	"goodssrv/initialize"
	"goodssrv/model"
	"goodssrv/proto"
	"goodssrv/utils"
	"net"
	"strconv"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
)

func main() {
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()
	model.CreateTables()
	// 创建UUID
	u1 := uuid.Must(uuid.NewV4()).String()
	fmt.Printf("UUIDv4: %s\n", u1)

	g := grpc.NewServer()
	proto.RegisterGoodsServer(g, &goods.Goods{})
	grpc_health_v1.RegisterHealthServer(g, &v1.HealthImpl{})
	port, _ := utils.GetFreePort()
	listener, err := net.Listen("tcp", "127.0.0.1:"+strconv.Itoa(port))
	fmt.Println("Using port:", port)
	if err != nil {
		panic(err)
	}
	go func() {
		g.Serve(listener)
	}()
	v1.Register("127.0.0.1", port, global.ServerConfig.ServiceName, []string{"mall", "wei"}, u1)
	v1.OnExit(u1)

	// cfg := api.DefaultConfig()
	// cfg.Address = "127.0.0.1:8500"
	// client, _ := api.NewClient(cfg)
	// client.Agent().ServiceDeregister("5bc377bf-32e8-498c-a290-c3166ef3a7aa")

}
