package test

import (
	"context"
	"fmt"
	"online-mall-inventory/proto"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"google.golang.org/grpc"
)

func TestSell(t *testing.T) {
	Convey("减库存测试，库存不够", t, func() {
		conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		client := proto.NewInventoryClient(conn)
		goodsInfo := []*proto.GoodsInvInfo{
			{GoodsId: 1, Num: 1},
			{GoodsId: 2, Num: 1},
			{GoodsId: 3, Num: 100},
		}
		sellInfo := proto.SellInfo{GoodsInfo: goodsInfo}
		rsp, err := client.Sell(context.Background(), &sellInfo)
		if err != nil {
			fmt.Println(err)
		}
		So(rsp.Code, ShouldEqual, proto.RetCode_inventory_insufficient)
	})
	Convey("减库存测试，库存足够", t, func() {
		conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		client := proto.NewInventoryClient(conn)
		goodsInfo := []*proto.GoodsInvInfo{
			{GoodsId: 1, Num: 1},
			{GoodsId: 2, Num: 1},
			{GoodsId: 3, Num: 1},
		}
		rsp0, err := client.GetInv(context.Background(), goodsInfo[0])
		sellInfo := proto.SellInfo{GoodsInfo: goodsInfo}
		rsp, err := client.Sell(context.Background(), &sellInfo)
		if err != nil {
			fmt.Println(err)
		}
		rsp1, err := client.GetInv(context.Background(), goodsInfo[0])
		So(rsp.Code, ShouldEqual, proto.RetCode_SUCCESS)
		So(rsp0.Num, ShouldEqual, rsp1.Num+goodsInfo[0].Num)
	})

}

func TestReBack(t *testing.T) {
	Convey("库存归还", t, func() {
		conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		client := proto.NewInventoryClient(conn)
		goodsInfo := []*proto.GoodsInvInfo{
			{GoodsId: 1, Num: 1},
			{GoodsId: 2, Num: 1},
			{GoodsId: 3, Num: 1},
		}
		sellInfo := proto.SellInfo{GoodsInfo: goodsInfo}
		rsp, err := client.Reback(context.Background(), &sellInfo)
		if err != nil {
			fmt.Println(err)
		}
		So(rsp.Code, ShouldEqual, proto.RetCode_SUCCESS)
	})
}

func TestGetInv(t *testing.T) {
	Convey("获取库存,设置库存", t, func() {
		conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		client := proto.NewInventoryClient(conn)
		goodsInfo := proto.GoodsInvInfo{GoodsId: 1, Num: 99}
		rsp, err := client.SetInv(context.Background(), &goodsInfo)
		if err != nil {
			fmt.Println("xxxxxxxxxxxxxx", err)
		}
		So(rsp.Code, ShouldEqual, proto.RetCode_SUCCESS)
		rsp2, err := client.GetInv(context.Background(), &goodsInfo)
		if err != nil {
			fmt.Println("获取库存失败", err)
		}
		So(rsp2.Num, ShouldEqual, 99)
	})

}

// func TestInsertInv(t *testing.T) {
// 	Convey("获取库存,设置库存", t, func() {
// 		conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
// 		if err != nil {
// 			panic(err)
// 		}
// 		defer conn.Close()
// 		client := proto.NewInventoryClient(conn)
// 		for i := 100; i < 300; i++ {
// 			goodsInfo := proto.GoodsInvInfo{GoodsId: int32(i), Num: 99}
// 			zap.S().Info(i)
// 			_, err := client.SetInv(context.Background(), &goodsInfo)
// 			if err != nil {
// 				fmt.Println("xxxxxxxxxxxxxx", err)
// 			}
// 		}

// 	})

// }
