package initialize

import (
	"fmt"
	"online-mall-order-web/global"
	"online-mall-order-web/proto"

	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func InitSrvConn() {
	consulInfo := global.ServerConfig.ConsulInfo
	goodsConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.GoodsSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Info(fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.GoodsSrvInfo.Name))
		zap.S().Fatal("[InitSrvConn] 连接 【商品服务失败】")
	}

	global.GoodsSrvClient = proto.NewGoodsClient(goodsConn)

	orderConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.OrderSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 【订单服务失败】")
	}

	global.OrderSrvClient = proto.NewOrderClient(orderConn)

	invConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.InventorySrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 【库存服务失败】")
	}

	global.InventorySrvClient = proto.NewInventoryClient(invConn)

}

// func InitServConn2() {
// 	cfg := consulApi.DefaultConfig()
// 	consulInfo := global.ServerConfig.ConsulInfo
// 	cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)
// 	zap.S().Debug("cfg.Address", cfg.Address)
// 	client, err := consulApi.NewClient(cfg)

// 	if err != nil {
// 		panic(err)
// 	}
// 	filterStr := fmt.Sprintf("Service==\"%s\"", global.ServerConfig.UserSrvInfo.SrvName)
// 	data, err := client.Agent().ServicesWithFilter(filterStr)
// 	if err != nil {
// 		panic(err)
// 	}
// 	var host string
// 	var port int
// 	for _, value := range data {
// 		host = value.Address
// 		port = value.Port
// 		break
// 	}
// 	if host == "" {
// 		zap.S().Fatal("连接用户服务失败")
// 	}

// 	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithInsecure())
// 	if err != nil {
// 		zap.S().Errorw("连接用户服务失败", "msg", err)
// 	}
// 	//如何维护全局变量
// 	// 1 后续服务下线了 2 该端口 3ip  负载均衡
// 	//一个连接会有多个groutine共用，连接池 或者 负载均衡  todo
// 	userSrvClient := proto.NewUserClient(userConn)
// 	global.UseSrvClient = userSrvClient

// }
