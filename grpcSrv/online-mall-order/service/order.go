package service

import (
	"context"
	"encoding/json"
	"fmt"
	"online-mall-order/global"
	"online-mall-order/model"
	"online-mall-order/proto"
	"strconv"
	"sync"
	"time"

	"math/rand"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type localTransStatus int

const (
	localTrans_success localTransStatus = 1
	localTrans_fail    localTransStatus = 2
)

type LocalTransMsg struct {
	OrderId   string
	Total     float32
	ErrorCode codes.Code
	ErrorInfo string
	Status    localTransStatus
}

type OrderService struct {
	LocalTransMsgMap *sync.Map
}
type OrderMsgBody struct {
	OrderId string `json:"order_id"`
	UserId  int32  `json:"user_id"`
	Address string `json:"address"`
	Mobile  string `json:"mobile"`
	Post    string `json:"post"`
	Name    string `json:"name"`
}

func NewOrderService() *OrderService {
	os := OrderService{}
	os.LocalTransMsgMap = new(sync.Map)
	return &os
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
		rsp.Code = proto.OrderRetCode_ERROR
		return &rsp, err
	}
	zap.S().Info("insert insertRsp:", insertRsp)
	rsp.Code = proto.OrderRetCode_SUCCESS
	return &rsp, nil
}

//删除已支付过的商品
func (o *OrderService) DeleteCartItem(ctx context.Context, request *proto.CartItemRequest) (*proto.CommonResponse, error) {
	rsp := proto.CommonResponse{Code: proto.OrderRetCode_SUCCESS}
	shoppingCartItem := model.ShoppingCart{}
	_, err := global.Engine.Where("user_Id=? and checked=?", request.UserId, true).Delete(&shoppingCartItem)
	if err != nil {
		rsp.Code = proto.OrderRetCode_ERROR
	}
	zap.S().Info("DeleteCartItem:", rsp)
	return &rsp, err
}

func (o *OrderService) UpdateCartItem(ctx context.Context, request *proto.CartItemRequest) (*proto.CommonResponse, error) {
	zap.S().Info("更新商品：", request)
	rsp := proto.CommonResponse{}
	updateField := map[string]interface{}{"user_id": request.UserId, "checked": request.Checked}
	affected, err := global.Engine.Table(new(model.ShoppingCart)).Where("user_id=? and goods_id=?", request.UserId, request.GoodsId).Update(updateField)
	// shoppingCartItem := model.ShoppingCart{}
	// shoppingCartItem.Checked = request.Checked
	// shoppingCartItem.Nums = request.Nums
	// updateRsp, err := global.Engine.Where("user_id=? and goods_id=?", request.UserId, request.GoodsId).Update(&shoppingCartItem)
	if err != nil || affected == 0 {
		zap.S().Error(err)
		rsp.Code = proto.OrderRetCode_ERROR
		return &rsp, status.Errorf(codes.NotFound, "商品不存在")
	}
	zap.S().Info("更新商品状态成功:", affected)
	rsp.Code = proto.OrderRetCode_SUCCESS
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
	commenRsp.Code = proto.OrderRetCode_SUCCESS
	// orderInfo := model.OrderInfo{Status: model.TradeStatus(request.Status)}
	orderInfo := map[string]interface{}{"status": request.Status}
	_, err := global.Engine.Table(new(model.OrderInfo)).Where("order_id=?", request.OrderId).Update(&orderInfo)
	if err != nil {
		zap.S().Error("UpdateOrderStatus failed : ", err)
		commenRsp.Code = proto.OrderRetCode_ERROR
		return &commenRsp, nil
	}
	return &commenRsp, nil
}
func (o *OrderService) CheckLocalTransaction(msg *primitive.MessageExt) primitive.LocalTransactionState {
	msg_body := OrderMsgBody{}
	err := json.Unmarshal(msg.Body, &msg_body)
	if err != nil {
		zap.S().Error("parse msg body error: ", err)
		//todo
	}
	//查询本地订单是否已经入库
	orderInfo := model.OrderInfo{}
	_, err = global.Engine.Where("order_id=?", msg_body.OrderId).Get(&orderInfo)
	if orderInfo.Id == 0 {
		zap.S().Info("未查询到本地事务执行成功")
		return primitive.RollbackMessageState
	}
	return primitive.CommitMessageState
}

//订单
func (o *OrderService) CreateOrder(ctx context.Context, request *proto.OrderRequest) (*proto.OrderInfoResponse, error) {
	zap.S().Info("recieve a request:", request)
	//half消息 一个服务只需要一个producer todo
	mqAddr := fmt.Sprintf("%s:%d", global.ServerConfig.RocketMQInfo.Host, global.ServerConfig.RocketMQInfo.Port)
	zap.S().Info("rocketmq address:", mqAddr)

	p, err := rocketmq.NewTransactionProducer(
		o,
		producer.WithNameServer([]string{mqAddr}),
		producer.WithGroupName("order_reback_when_failed"),
		producer.WithRetry(1),
	)
	if err != nil {
		zap.S().Info("创建事务生产者失败：", err)
	}
	err = p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s\n", err.Error())
	}
	order_id := generate_orderId(request.UserId)
	msg_body := OrderMsgBody{
		OrderId: order_id,
		UserId:  request.UserId,
		Address: request.Address,
		Mobile:  request.Mobile,
		Post:    request.Post,
		Name:    request.Name,
	}
	ch := make(chan LocalTransMsg, 1)
	o.LocalTransMsgMap.Store(order_id, ch)

	body, err := json.Marshal(msg_body)
	if err != nil {
		zap.S().Error("msg body marshal error:", err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	msg := primitive.Message{
		Topic: "reback_inv",
		Body:  body,
	}
	res, err := p.SendMessageInTransaction(context.Background(), &msg)
	zap.S().Info("SendMessageInTransaction", res, err)
	if err != nil || res == nil || res.Status != primitive.SendOK {
		zap.S().Error("SendMessageInTransaction failed:", res.Status)
		return nil, status.Errorf(codes.Internal, "新建订单失败")
	}
	zap.S().Info("before ch")
	localTransMsg := <-ch
	zap.S().Info("after ch", localTransMsg)
	p.Shutdown()
	if localTransMsg.Status != localTrans_success {
		return nil, status.Errorf(localTransMsg.ErrorCode, localTransMsg.ErrorInfo)
	}
	rsp := proto.OrderInfoResponse{
		OrderId: localTransMsg.OrderId,
		Total:   localTransMsg.Total,
	}
	return &rsp, nil
}

func (o *OrderService) ExecuteLocalTransaction(msg *primitive.Message) primitive.LocalTransactionState {
	msg_body := OrderMsgBody{}
	err := json.Unmarshal(msg.Body, &msg_body)
	if err != nil {
		zap.S().Error("parse msg body error: ", err)
		//todo
	}
	zap.S().Info("OrderMsgBody:", msg_body)
	//1,价格-访问商品服务
	//2,库存扣减-访问库存服务
	//3,创建订单基本信息表-订单的商品信息表
	items := make([]model.ShoppingCart, 0)
	goods_nums := make(map[int32]int32, 0)
	goods_id := make([]int32, 0)
	goods_sell_info := proto.SellInfo{
		OrderId:   msg_body.OrderId,
		GoodsInfo: make([]*proto.GoodsInvInfo, 0),
	}

	//获取购物车记录
	// err := global.Engine.Find(&it	orderId := generate_orderId(request.UserId)ems, &model.ShoppingCart{UserId: request.UserId, Checked: true}) //todo
	err = global.Engine.Where("user_id=? and checked=true", msg_body.UserId).Find(&items) //todo
	zap.S().Info("获取购物车记录", items)
	if err != nil {
		localTransMsg := LocalTransMsg{
			Status:    localTrans_fail,
			ErrorCode: codes.ResourceExhausted,
			ErrorInfo: "获取购物车记录失败",
		}
		return o.transReturnHandler(localTransMsg, false)
	}

	for _, item := range items {
		goods_nums[item.GoodsId] = item.Nums
		goods_id = append(goods_id, item.GoodsId)
		zap.S().Info("item.GoodsId:", item.GoodsId)
	}

	if len(goods_id) == 0 {
		localTransMsg := LocalTransMsg{
			OrderId:   msg_body.OrderId,
			Status:    localTrans_fail,
			ErrorCode: codes.NotFound,
			ErrorInfo: "获取购物车商品失败",
		}
		return o.transReturnHandler(localTransMsg, false)
	}

	//查询商品信息
	req := proto.BatchGoodsIdInfo{Id: goods_id}
	order_goods_slice := make([]*model.OrderGoods, 0)
	goodsListRsp, err := global.GoodsSrvClient.BatchGetGoods(context.Background(), &req)
	if err != nil {
		zap.S().Info("goodsListRsp error", err)
		localTransMsg := LocalTransMsg{
			OrderId:   msg_body.OrderId,
			Status:    localTrans_fail,
			ErrorCode: codes.NotFound,
			ErrorInfo: "查询商品列表失败",
		}
		return o.transReturnHandler(localTransMsg, false)
	}
	zap.S().Info("goodsListRsp:", goodsListRsp.Data)
	var order_amount float32 = 0

	for _, goodsInfo := range goodsListRsp.Data {
		order_goods := model.OrderGoods{}
		order_amount += goodsInfo.ShopPrice * float32(goods_nums[goodsInfo.GoodsId])
		zap.S().Info("goodsInfo:", goodsInfo)
		order_goods.GoodsId = goodsInfo.GoodsId
		order_goods.GoodsName = goodsInfo.Name
		order_goods.GoodsPrice = goodsInfo.ShopPrice
		order_goods_slice = append(order_goods_slice, &order_goods)

		goodsInvInfo := proto.GoodsInvInfo{GoodsId: goodsInfo.GoodsId, Num: goods_nums[goodsInfo.GoodsId]}
		goods_sell_info.GoodsInfo = append(goods_sell_info.GoodsInfo, &goodsInvInfo)
		zap.S().Info("goodsInvInfo:", goodsInvInfo)
	}

	//扣减库存 失败的情况分析
	zap.S().Info("InvSrvClient.Sell:", goods_sell_info.GoodsInfo)
	_, err = global.InvSrvClient.Sell(context.Background(), &goods_sell_info)
	if err != nil {
		zap.S().Error("InvSrvClient.Sell error:", err)
		localTransMsg := LocalTransMsg{
			OrderId:   msg_body.OrderId,
			Status:    localTrans_fail,
			ErrorCode: codes.NotFound,
			ErrorInfo: "扣减库存失败",
		}
		e, _ := status.FromError(err)
		if e.Code() == codes.Unknown || e.Code() == codes.DeadlineExceeded {
			return o.transReturnHandler(localTransMsg, true) //无法预知是否扣减库存成功，处理重复归还以及库存不足的情况，弱库存不足那么，扣减历史表里肯定没有记录，在consumer处直接确认次消息即可，本次订单创建失败
		} else {
			return o.transReturnHandler(localTransMsg, false)
		}
	}

	session := global.Engine.NewSession()
	defer session.Close()
	err = session.Begin()
	//创建订单
	order := model.OrderInfo{}
	order.Address = msg_body.Address
	order.UserId = msg_body.UserId
	order.SignerNum = msg_body.Mobile
	order.OrderId = msg_body.OrderId
	order.OrderMount = order_amount
	order.Status = model.TradeStatusWaitBuyerPay
	order.Post = msg_body.Post
	order.Total = order_amount
	order.SignerName = msg_body.Name

	affected, err := session.Insert(&order)
	if err != nil || affected == 0 {
		zap.S().Error("插入订单失败：", err)
		session.Rollback()
		localTransMsg := LocalTransMsg{
			OrderId:   msg_body.OrderId,
			Status:    localTrans_fail,
			ErrorCode: codes.Internal,
			ErrorInfo: "插入订单失败",
		}
		return o.transReturnHandler(localTransMsg, true)
	}

	//批量插入订单商品表
	for _, order_goods := range order_goods_slice {
		order_goods.OrderId = order.OrderId
	}
	zap.S().Info("order_goods_slice", order_goods_slice)
	affected, err = session.Insert(&order_goods_slice)

	if err != nil || affected == 0 {
		zap.S().Error("批量插入订单商品表失败：", err)
		session.Rollback()
		localTransMsg := LocalTransMsg{
			OrderId:   msg_body.OrderId,
			Status:    localTrans_fail,
			ErrorCode: codes.NotFound,
			ErrorInfo: "创建订单商品表失败",
		}
		return o.transReturnHandler(localTransMsg, true)
	}

	//删除购物车已支付商品记录
	shoppingCartItem := model.ShoppingCart{}
	_, err = session.Where("user_Id=? and checked=?", msg_body.UserId, true).Delete(&shoppingCartItem)
	if err != nil {
		session.Rollback()
		zap.S().Info("goodsListRsp error", err)
		localTransMsg := LocalTransMsg{
			OrderId:   msg_body.OrderId,
			Status:    localTrans_fail,
			ErrorCode: codes.Internal,
			ErrorInfo: "删除购物车记录失败",
		}
		return o.transReturnHandler(localTransMsg, true)
	}

	mqAddr := fmt.Sprintf("%s:%d", global.ServerConfig.RocketMQInfo.Host, global.ServerConfig.RocketMQInfo.Port)
	p, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{mqAddr}),
		producer.WithGroupName("order_delay_group"),
	)
	p.Start()
	orderTimeOutMsg := OrderMsgBody{
		OrderId: msg_body.OrderId,
	}
	orderDelayMsg, _ := json.Marshal(orderTimeOutMsg)
	priMsg := &primitive.Message{
		Topic: "order_delay_topic",
		Body:  orderDelayMsg,
	}
	priMsg.WithDelayTimeLevel(2)
	res, err := p.SendSync(context.Background(), priMsg)
	if err != nil || res.Status != primitive.SendOK {
		session.Rollback()
		zap.S().Error("发送订单延时消息失败：", err)
		localTransMsg := LocalTransMsg{
			OrderId:   msg_body.OrderId,
			Status:    localTrans_fail,
			ErrorCode: codes.Internal,
			ErrorInfo: "发送延时消息失败",
		}
		o.transReturnHandler(localTransMsg, true)
	}
	session.Commit()

	localTransMsg := LocalTransMsg{
		OrderId: msg_body.OrderId,
		Total:   order_amount,
		Status:  localTrans_success,
	}
	zap.S().Info("finish locak transaction:", localTransMsg)
	return o.transReturnHandler(localTransMsg, false)
}
func OrderDelayProcess(msg *primitive.MessageExt) error {
	orderMsgBody := OrderMsgBody{}
	err := json.Unmarshal(msg.Body, &orderMsgBody)
	if err != nil {
		zap.S().Error("messgae body parse error:", err) //todo
		return err
	}
	//查询订单状态
	orderInfo := model.OrderInfo{}
	ok, err := global.Engine.Where("order_id=?", orderMsgBody.OrderId).Get(&orderInfo)
	zap.S().Info("查询订单状态失败：", orderInfo, ok)
	if err != nil || !ok {
		zap.S().Error("查询订单状态失败：", err)
		return nil
	}
	if orderInfo.Status != model.TradeStatusWaitBuyerPay {
		return nil
	}
	//订单已超时 未支付
	mqAddr := fmt.Sprintf("%s:%d", global.ServerConfig.RocketMQInfo.Host, global.ServerConfig.RocketMQInfo.Port)
	p, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{mqAddr}),
		producer.WithGroupName("order_reback_when_failed"),
	)
	p.Start()
	orderTimeOutMsg := OrderMsgBody{
		OrderId: orderMsgBody.OrderId,
	}
	orderDelayMsg, _ := json.Marshal(orderTimeOutMsg)
	priMsg := &primitive.Message{
		Topic: "reback_inv",
		Body:  orderDelayMsg,
	}
	priMsg.WithDelayTimeLevel(2)
	res, err := p.SendSync(context.Background(), priMsg)
	if err != nil || res.Status != primitive.SendOK {
		zap.S().Error("发送订单延时消息失败：", err)
		return err
	}
	return nil
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
	orderInfoDto.Status = string(orderInfo.Status)
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

func (o *OrderService) transReturnHandler(msg LocalTransMsg, isCommmit bool) primitive.LocalTransactionState {
	ch, ok := o.LocalTransMsgMap.Load(msg.OrderId)
	if !ok {
		zap.S().Error("order id 不存在")
		//todo
	}
	c, ok := ch.(chan LocalTransMsg)
	zap.S().Info("写入数据到channel:")
	c <- msg
	zap.S().Info("写入数据到channel完成:")
	if isCommmit {
		return primitive.CommitMessageState
	}
	return primitive.RollbackMessageState
}
