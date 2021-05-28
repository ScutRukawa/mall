package main

import (
	"context"
	"fmt"
	"net"
	"online-mall-inventory/global"
	"online-mall-inventory/initialize"
	"online-mall-inventory/model"
	"online-mall-inventory/proto"
	"online-mall-inventory/service"
	"online-mall-inventory/utils"
	"strconv"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
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
	//创建UUID
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
	//
	mqAddr := fmt.Sprintf("%s:%d", global.ServerConfig.RocketMQInfo.Host, global.ServerConfig.RocketMQInfo.Port)
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName("order_reback_when_failed"),
		consumer.WithNameServer([]string{mqAddr}),
	)
	if err != nil {
		panic(err)
	}

	err = c.Subscribe("reback_inv", consumer.MessageSelector{},
		func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for _, msg := range msgs {
				zap.S().Info("get message from mq:", msg)
				if service.Reback_inv(msg) != nil {
					return consumer.ConsumeRetryLater, nil
				}
			}
			zap.S().Info("消费确认")
			return consumer.ConsumeSuccess, nil
		})
	if err != nil {
		fmt.Println(err.Error())
	}
	if err = c.Start(); err != nil {
		fmt.Println(err.Error())
	}
	utils.OnExit(u1)
	c.Shutdown()
	// cfg := api.DefaultConfig()
	// cfg.Address = "127.0.0.1:8500"
	// client, _ := api.NewClient(cfg)
	// client.Agent().ServiceDeregister("881394ea-46c8-4ad6-b677-a05dc5dc660a")

}
