package test

import (
	"context"
	"online-mall-order/proto"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func TestCartItemList(t *testing.T) {
	Convey("获取购物车列表", t, func() {
		conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		client := proto.NewOrderClient(conn)
		rsp, err := client.CartItemList(context.Background(), &proto.UserInfo{Id: 171})
		if err != nil {
			zap.S().Error("CartItemList error: ", err)
		}
		So(len(rsp.Data), ShouldEqual, 2)
	})

}
func TestCreateCartItem(t *testing.T) {
	Convey("加入购物车", t, func() {
		conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		client := proto.NewOrderClient(conn)
		CartItemRequest := proto.CartItemRequest{
			UserId:  172,
			GoodsId: 3,
			Nums:    1,
			Checked: true,
		}
		rsp, err := client.CreateCartItem(context.Background(), &CartItemRequest)
		if err != nil {
			zap.S().Error("CreateCartItem error: ", err)
		}
		So(rsp.Code, ShouldEqual, proto.RetCode_SUCCESS)
	})
}
func TestDeleteCartItem(t *testing.T) {
	Convey("加入购物车", t, func() {
		conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		client := proto.NewOrderClient(conn)
		CartItemRequest := proto.CartItemRequest{
			UserId: 172,
		}
		rsp, err := client.DeleteCartItem(context.Background(), &CartItemRequest)
		if err != nil {
			zap.S().Error("DeleteCartItem error: ", err)
		}
		So(rsp.Code, ShouldEqual, proto.RetCode_SUCCESS)
	})
}
func TestUpdateCartItem(t *testing.T) {
	Convey("加入购物车", t, func() {
		conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		client := proto.NewOrderClient(conn)
		CartItemRequest := proto.CartItemRequest{
			UserId:  171,
			Nums:    100,
			Checked: true,
			GoodsId: 5,
		}
		rsp, err := client.UpdateCartItem(context.Background(), &CartItemRequest)
		if err != nil {
			zap.S().Error("UpdateCartItem error: ", err)
		}
		So(rsp.Code, ShouldEqual, proto.RetCode_SUCCESS)
	})
}
