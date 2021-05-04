package service

import (
	"context"
	"fmt"
	"online-mall-order/global"
	"online-mall-order/model"
	"online-mall-order/proto"
	"strconv"
	"time"

	"math/rand"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

type OrderService struct {
}

//购物车
func (o *OrderService) CartItemList(ctx context.Context, request *proto.UserInfo) (*proto.CartItemListResponse, error) {
	zap.S().Info("recieve a request:", request)
	cartListInfo := make([]model.ShoppingCart, 0)
	shopCartInfoResponse := make([]*proto.ShopCartInfoResponse, 0)
	cartItemListResponse := proto.CartItemListResponse{}
	err := global.Engine.Where("user_id=?", request.Id).Find(&cartListInfo)
	zap.S().Info("cartListInfo:", cartListInfo)
	if err != nil {
		zap.S().Error("get CartItemList error:", err)
		return nil, err
	}
	for _, v := range cartListInfo {
		cartInfo := proto.ShopCartInfoResponse{}
		cartInfo.Id = v.Id
		cartInfo.GoodsId = v.GoodsId
		cartInfo.Nums = v.Nums
		cartInfo.Checked = v.Checked
		cartInfo.UserId = v.UserId
		shopCartInfoResponse = append(shopCartInfoResponse, &cartInfo)
	}
	cartItemListResponse.Total = int32(len(shopCartInfoResponse))
	cartItemListResponse.Data = shopCartInfoResponse
	zap.S().Info(cartItemListResponse.Data)
	return &cartItemListResponse, nil
}

func (o *OrderService) CreateCartItem(ctx context.Context, request *proto.CartItemRequest) (*proto.CommonResponse, error) {
	rsp := proto.CommonResponse{}
	shoppingCartItem := model.ShoppingCart{}
	shoppingCartItem.Checked = request.Checked
	shoppingCartItem.GoodsId = request.GoodsId
	shoppingCartItem.UserId = request.UserId
	shoppingCartItem.Nums = request.Nums
	insertRsp, err := global.Engine.Insert(&shoppingCartItem)
	if err != nil {
		zap.S().Error(err)
		rsp.Code = proto.RetCode_ERROR
		return &rsp, err
	}
	zap.S().Info("insert insertRsp:", insertRsp)
	rsp.Code = proto.RetCode_SUCCESS
	return &rsp, nil
}

//删除已支付过的商品
func (o *OrderService) DeleteCartItem(ctx context.Context, request *proto.CartItemRequest) (*proto.CommonResponse, error) {
	rsp := proto.CommonResponse{Code: proto.RetCode_SUCCESS}
	shoppingCartItem := model.ShoppingCart{}
	_, err := global.Engine.Where("user_Id=? and checked=?", request.UserId, true).Delete(&shoppingCartItem)
	if err != nil {
		rsp.Code = proto.RetCode_ERROR
	}
	zap.S().Info("DeleteCartItem:", rsp)
	return &rsp, err
}

func (o *OrderService) UpdateCartItem(ctx context.Context, request *proto.CartItemRequest) (*proto.CommonResponse, error) {
	rsp := proto.CommonResponse{}
	shoppingCartItem := model.ShoppingCart{}
	shoppingCartItem.Checked = request.Checked
	shoppingCartItem.Nums = request.Nums
	shoppingCartItem.GoodsId = request.GoodsId
	updateRsp, err := global.Engine.Update(&shoppingCartItem,
		model.ShoppingCart{UserId: request.UserId, GoodsId: request.GoodsId})
	if err != nil {
		zap.S().Error(err)
		rsp.Code = proto.RetCode_ERROR
		return &rsp, err
	}
	zap.S().Info("insert updateRsp:", updateRsp)
	rsp.Code = proto.RetCode_SUCCESS
	return &rsp, nil
}

//订单
func (o *OrderService) CreateOrder(ctx context.Context, request *proto.OrderRequest) (*proto.OrderInfoResponse, error) {
	//1,价格-访问商品服务
	//2,库存扣减-访问库存服务
	//3,创建订单基本信息表-订单的商品信息表
	items := make([]model.ShoppingCart, 0)
	goods_nums := make(map[int32]int32, 0)
	goods_id := make([]int32, 0)
	goods_sell_info := proto.SellInfo{
		GoodsInfo: make([]*proto.GoodsInvInfo, 0),
	}

	//获取购物车记录
	// err := global.Engine.Find(&items, &model.ShoppingCart{UserId: request.UserId, Checked: true}) //todo
	err := global.Engine.Where("user_id=? and checked=true", request.UserId).Find(&items) //todo
	zap.S().Info("获取购物车记录", items)
	if err != nil {
		zap.S().Error("获取购物车记录失败", err)
		return nil, err
	}
	for _, item := range items {
		goods_nums[item.GoodsId] = item.Nums
		goods_id = append(goods_id, item.GoodsId)
		zap.S().Info("item.GoodsId:", item.GoodsId)
	}
	if len(goods_id) == 0 {
		md := metadata.Pairs("code", strconv.Itoa(int(codes.NotFound)), "message", "没有选中的商品")
		ctx = metadata.NewOutgoingContext(context.Background(), md)
		return nil, nil
	}

	//查询商品信息
	req := proto.BatchGoodsIdInfo{Id: goods_id}
	order_goods_slice := make([]model.OrderGoods, 0)
	goodsListRsp, err := global.GoodsSrvClient.BatchGetGoods(ctx, &req)
	if err != nil {
		md := metadata.Pairs("code", strconv.Itoa(int(codes.NotFound)), "message", "获取商品信息失败")
		ctx = metadata.NewOutgoingContext(context.Background(), md)
		zap.S().Info("goodsListRsp error", err)
		return nil, err
	}
	zap.S().Info("goodsListRsp:", goodsListRsp.Data)
	var order_amount float32 = 0
	order_goods := model.OrderGoods{}
	for _, goodsInfo := range goodsListRsp.Data {
		order_amount += goodsInfo.ShopPrice * float32(goods_nums[goods_id[goodsInfo.GoodsId]])
		order_goods.GoodsId = goodsInfo.Id
		order_goods.GoodsName = goodsInfo.Name
		order_goods.GoodsPrice = goodsInfo.ShopPrice
		order_goods_slice = append(order_goods_slice, order_goods)

		goodsInvInfo := proto.GoodsInvInfo{GoodsId: goodsInfo.Id, Num: goods_nums[goods_id[goodsInfo.GoodsId]]}
		goods_sell_info.GoodsInfo = append(goods_sell_info.GoodsInfo, &goodsInvInfo)
		zap.S().Info("goodsInvInfo:", goodsInvInfo)
	}
	zap.S().Info("goods_nums:", goods_nums)
	//扣减库存
	zap.S().Info("InvSrvClient.Sell:", goods_sell_info.GoodsInfo)
	_, err = global.InvSrvClient.Sell(ctx, &goods_sell_info)
	if err != nil {
		zap.S().Error("InvSrvClient.Sell error:", err)
		md := metadata.Pairs("code", strconv.Itoa(int(codes.ResourceExhausted)), "message", "扣减库存失败")
		ctx = metadata.NewOutgoingContext(context.Background(), md)
		return nil, err
	}
	session := global.Engine.NewSession()
	session.Begin()
	//创建订单
	order := model.OrderInfo{}
	order.Address = request.Address
	order.UserId = request.UserId
	order.SignerNum = request.Mobile
	order.OrderId = generate_orderId(request.UserId)
	order.OrderMount = order_amount
	order.Status = model.WAIT_BUYER_PAY
	order.Post = request.Post
	order.Total = order_amount
	order.SignerName = request.Name

	zap.S().Info("order.OrderId", order.OrderId)

	//批量插入订单商品表
	for _, order_goods := range order_goods_slice {
		order_goods.OrderId = order.OrderId
	}
	_, err = global.Engine.Insert(&order_goods_slice)
	if err != nil {
		zap.S().Error("批量插入订单商品表失败：", err)
		md := metadata.Pairs("code", strconv.Itoa(int(codes.Unknown)), "message", "插入订单商品失败")
		ctx = metadata.NewOutgoingContext(context.Background(), md)
		session.Rollback()
		return nil, err
	}

	//删除购物车已支付商品记录
	shoppingCartItem := model.ShoppingCart{}
	_, err = global.Engine.Where("user_Id=? and checked=?", request.UserId, true).Delete(&shoppingCartItem)
	if err != nil {
		session.Rollback()
		return nil, err
	}
	session.Commit()
	rsp := proto.OrderInfoResponse{
		Id:      order.Id,
		OrderId: order.OrderId,
		Total:   order.Total,
	}
	return &rsp, nil
}
func (o *OrderService) OrderList(ctx context.Context, request *proto.OrderFilterRequest) (*proto.OrderListResponse, error) {
	orderList := make([]*model.OrderInfo, 0)
	orderListRsp := proto.OrderListResponse{}
	err := global.Engine.Where("user_id=?", request.UserId).Limit(int(request.Pages)*int(request.PagePerNum), int(request.PagePerNum)).Find(&orderList)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}
	for _, v := range orderList {
		orderInfoRsp := proto.OrderInfoResponse{}
		convertOrderInfo2Dto(&orderInfoRsp, v)
		orderListRsp.Data = append(orderListRsp.Data, &orderInfoRsp)
	}
	orderListRsp.Total = int32(len(orderListRsp.Data))
	return &orderListRsp, nil
}

func (o *OrderService) OrderDetail(ctx context.Context, request *proto.OrderRequest) (*proto.OrderInfoDetailResponse, error) {
	orderInfoDetailRsp := proto.OrderInfoDetailResponse{}
	orderInfo := model.OrderInfo{}
	orderGoodsItem := make([]*model.OrderGoods, 0)
	_, err := global.Engine.Where("order_id=?", request.OrderId).Get(orderInfo)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}
	err = global.Engine.Where("order_id=?", request.OrderId).Find(&orderGoodsItem)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}
	for _, v := range orderGoodsItem {
		orderItemRsp := proto.OrderItemResponse{}
		convertOrderItem2Dto(&orderItemRsp, v)
		orderInfoDetailRsp.Data = append(orderInfoDetailRsp.Data, &orderItemRsp)
	}
	orderInfoRsp := proto.OrderInfoResponse{}
	convertOrderInfo2Dto(&orderInfoRsp, &orderInfo)
	orderInfoDetailRsp.OrderInfo = &orderInfoRsp
	return &orderInfoDetailRsp, nil
}

func (o *OrderService) UpdateOrderStatus(ctx context.Context, request *proto.OrderStatus) (*proto.CommonResponse, error) {
	commenRsp := proto.CommonResponse{}
	commenRsp.Code = proto.RetCode_SUCCESS
	orderInfo := model.OrderInfo{Status: model.OrderStatus(request.Status)}
	_, err := global.Engine.Where("order_id=?", request.OrderId).Update(&orderInfo)
	if err != nil {
		zap.S().Error("UpdateOrderStatus failed : ", err)
		commenRsp.Code = proto.RetCode_ERROR
		return &commenRsp, nil
	}
	return &commenRsp, nil
}

func convertOrderInfo2Dto(orderInfoDto *proto.OrderInfoResponse, orderInfo *model.OrderInfo) {
	orderInfoDto.Address = orderInfo.Address
	orderInfoDto.Id = orderInfo.Id
	orderInfoDto.Mobile = orderInfo.Mobile
	orderInfoDto.OrderId = orderInfo.OrderId
	orderInfoDto.PayType = orderInfo.PayType
	orderInfoDto.Name = orderInfo.SignerName
	orderInfoDto.Post = orderInfo.Post
	orderInfoDto.UserId = orderInfo.UserId
	orderInfoDto.Status = int32(orderInfo.Status)
	orderInfoDto.Total = float32(orderInfo.Total)
}

func convertOrderItem2Dto(orderItemRsp *proto.OrderItemResponse, orderGoods *model.OrderGoods) {
	orderItemRsp.GoodsId = int32(orderGoods.GoodsId)
	orderItemRsp.GoodsName = orderGoods.GoodsName
	orderItemRsp.GoodsPrice = orderGoods.GoodsPrice
	orderItemRsp.Nums = orderGoods.Nums
	orderItemRsp.OrderId = orderGoods.OrderId
	orderItemRsp.Id = orderGoods.Id
}

func generate_orderId(userId int32) string {
	//current time + userId+random
	random := rand.NewSource(time.Now().UnixNano())
	orderId := fmt.Sprintf("%s%s%s",
		strconv.FormatInt(time.Now().UnixNano(), 10),
		strconv.FormatInt(int64(userId), 10),
		strconv.FormatInt(random.Int63(), 10),
	)
	return orderId
}
