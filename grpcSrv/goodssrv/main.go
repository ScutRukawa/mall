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

	// goodsList := make([]*model.Goods, 10)
	// var goodsList []model.Goods
	// query := "name LIKE ? AND is_hot=? AND is_new=? AND ?<=shop_price<=? AND brands_id=?"
	// db.GetDB().Where(query, "%"+"goods1"+"%", false, false, 0, 100, 10).Find(&goodsList)

	// fmt.Println("v", goodsList)

}
