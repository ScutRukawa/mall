package test

import (
	"context"
	"goodssrv/proto"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func TestCreateOrder(t *testing.T) {
	Convey("获取购物车列表", t, func() {
		conn, err := grpc.Dial("127.0.0.1:37003", grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		client := proto.NewGoodsClient(conn)
		rsp, err := client.BatchGetGoods(context.Background(),
			&proto.BatchGoodsIdInfo{
				Id: []int32{4, 5, 6},
			})
		if err != nil {
			zap.S().Error("CartItemList error: ", err)
		}
		So(rsp.Total, ShouldEqual, 3)
	})

}
