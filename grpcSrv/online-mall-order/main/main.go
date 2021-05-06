package main

import (
	"fmt"
	"net"
	"online-mall-order/global"
	"online-mall-order/initialize"
	"online-mall-order/model"
	"online-mall-order/proto"
	"online-mall-order/service"
	"online-mall-order/utils"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
)

var conf = "root:123456@tcp(127.0.0.1:3339)/goods?charset=utf8mb4&parseTime=True&loc=Local"

func main() {

	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()
	initialize.InitConn()
	model.CreateTables()
	// 创建UUID
	u1 := uuid.Must(uuid.NewV4()).String()
	fmt.Printf("UUIDv4: %s\n", u1)

	g := grpc.NewServer()
	proto.RegisterOrderServer(g, &service.OrderService{})
	port, _ := utils.GetFreePort()
	listener, err := net.Listen("tcp", "127.0.0.1:"+strconv.Itoa(port))
	if err != nil {
		panic(err)
	}
	go func() {
		g.Serve(listener)
	}()
	utils.Register("127.0.0.1", port, global.ServerConfig.ServiceName, []string{"mall", "wei"}, u1)
	utils.OnExit(u1)
}
