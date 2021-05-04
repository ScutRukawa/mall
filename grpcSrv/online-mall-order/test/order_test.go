package test

import (
	"context"
	"online-mall-order/proto"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func TestCreateOrder(t *testing.T) {
	Convey("获取购物车列表", t, func() {
		conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		client := proto.NewOrderClient(conn)
		rsp, err := client.CreateOrder(context.Background(),
			&proto.OrderRequest{
				UserId:  173,
				Address: "guangzhou",
				Mobile:  "17306674706",
				Name:    "wei",
				Post:    "sounds good",
			})
		if err != nil {
			zap.S().Error("CartItemList error: ", err)
		}
		So(rsp, ShouldNotBeNil)
		So(rsp.OrderId, ShouldNotEqual, "")
	})

}
