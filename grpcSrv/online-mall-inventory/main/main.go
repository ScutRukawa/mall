package main

import (
	"fmt"
	"net"
	"strconv"

	"online-mall-inventory/global"
	"online-mall-inventory/initialize"
	"online-mall-inventory/model"
	"online-mall-inventory/proto"
	"online-mall-inventory/service"
	"online-mall-inventory/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
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

	port, _ := utils.GetFreePort()
	zap.S().Info("inventory service port:", port)
	g := grpc.NewServer()
	proto.RegisterInventoryServer(g, &service.InventoryService{})
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
