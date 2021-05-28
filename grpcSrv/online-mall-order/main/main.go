package main

import (
	"context"
	"fmt"
	"net"
	"online-mall-order/global"
	"online-mall-order/initialize"
	"online-mall-order/model"
	"online-mall-order/proto"
	"online-mall-order/service"
	"online-mall-order/utils"
	"strconv"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
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
	orderService := service.NewOrderService()
	proto.RegisterOrderServer(g, orderService)
	port, _ := utils.GetFreePort()
	listener, err := net.Listen("tcp", "127.0.0.1:"+strconv.Itoa(port))
	if err != nil {
		panic(err)
	}
	go func() {
		g.Serve(listener)
	}()

	mqAddr := fmt.Sprintf("%s:%d", global.ServerConfig.RocketMQInfo.Host, global.ServerConfig.RocketMQInfo.Port)
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName("order_delay_group"),
		consumer.WithNameServer([]string{mqAddr}),
	)
	if err != nil {
		panic(err)
	}

	err = c.Subscribe("order_delay_topic", consumer.MessageSelector{},
		func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for _, msg := range msgs {
				zap.S().Info("get message from mq:", msg)
				if service.OrderDelayProcess(msg) != nil {
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
	utils.Register("127.0.0.1", port, global.ServerConfig.ServiceName, []string{"mall", "wei"}, u1)
	utils.OnExit(u1)
}
