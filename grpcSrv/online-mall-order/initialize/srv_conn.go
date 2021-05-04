package initialize

import (
	"fmt"
	"log"
	"online-mall-order/global"
	"online-mall-order/proto"

	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func InitConn() {
	consulInfo := global.ServerConfig.ConsulInfo
	address := fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)
	srvName := global.ServerConfig.GoodsSrv.Name
	zap.S().Info(fmt.Sprintf("consul://%s/%s?wait=14s", address, srvName))
	goodsConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s/%s?wait=14s", address, srvName),
		//wait 解析等待时间 limit 解析出多少个服务
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	zap.S().Info(goodsConn)
	if err != nil {
		log.Fatal(err)
	}
	goodsSrvClient := proto.NewGoodsClient(goodsConn)
	global.GoodsSrvClient = goodsSrvClient
	zap.S().Info("GoodsSrvClient :", goodsSrvClient)

	//inventory connection
	srvName = global.ServerConfig.InventorySrv.Name
	zap.S().Info(fmt.Sprintf("consul://%s/%s?wait=14s", address, srvName))
	invConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s/%s?wait=14s", address, srvName),
		//wait 解析等待时间 limit 解析出多少个服务
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	zap.S().Info(invConn)
	if err != nil {
		log.Fatal(err)
	}
	InvSrvClient := proto.NewInventoryClient(invConn)
	global.InvSrvClient = InvSrvClient
}

// func initGoodsConn() {

// }

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
