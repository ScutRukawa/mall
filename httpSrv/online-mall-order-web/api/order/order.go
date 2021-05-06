package order

import (
	"context"
	"net/http"
	"online-mall-order-web/api"
	"online-mall-order-web/forms"
	"online-mall-order-web/global"
	"online-mall-order-web/models"
	"online-mall-order-web/proto"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
	"go.uber.org/zap"
)

func List(ctx *gin.Context) {
	//订单的列表
	userId, _ := ctx.Get("userId")
	request := proto.OrderFilterRequest{}
	request.UserId = int32(userId.(uint))

	pages := ctx.DefaultQuery("p", "0")
	pagesInt, _ := strconv.Atoi(pages)
	request.Pages = int32(pagesInt)

	perNums := ctx.DefaultQuery("pnum", "0")
	perNumsInt, _ := strconv.Atoi(perNums)
	request.PagePerNum = int32(perNumsInt)

	request.Pages = int32(pagesInt)
	request.PagePerNum = int32(perNumsInt)

	rsp, err := global.OrderSrvClient.OrderList(context.Background(), &request)
	if err != nil {
		zap.S().Errorw("获取订单列表失败")
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	/*
		{
			"total":100,
			"data":[
				{
					"
				}
			]
		}
	*/
	reMap := gin.H{
		"total": rsp.Total,
	}
	orderList := make([]interface{}, 0)

	for _, item := range rsp.Data {
		tmpMap := map[string]interface{}{}

		tmpMap["id"] = item.Id
		tmpMap["status"] = item.Status
		tmpMap["pay_type"] = item.PayType
		tmpMap["user"] = item.UserId
		tmpMap["post"] = item.Post
		tmpMap["total"] = item.Total
		tmpMap["address"] = item.Address
		tmpMap["name"] = item.Name
		tmpMap["mobile"] = item.Mobile
		tmpMap["order_id"] = item.OrderId
		tmpMap["id"] = item.Id

		orderList = append(orderList, tmpMap)
	}
	reMap["data"] = orderList
	ctx.JSON(http.StatusOK, reMap)
}

func Detail(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, _ := ctx.Get("userId")
	i, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": "url格式出错",
		})
		return
	}

	request := proto.OrderRequest{
		UserId: int32(i),
	}
	claims, _ := ctx.Get("claims")
	model := claims.(*models.CustomClaims)
	if model.AuthorityId == 1 {
		request.UserId = int32(userId.(uint))
	}

	rsp, err := global.OrderSrvClient.OrderDetail(context.Background(), &request)
	if err != nil {
		zap.S().Errorw("获取订单详情失败")
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	reMap := gin.H{}
	reMap["id"] = rsp.OrderInfo.Id
	reMap["status"] = rsp.OrderInfo.Status
	reMap["user"] = rsp.OrderInfo.UserId
	reMap["post"] = rsp.OrderInfo.Post
	reMap["total"] = rsp.OrderInfo.Total
	reMap["address"] = rsp.OrderInfo.Address
	reMap["name"] = rsp.OrderInfo.Name
	reMap["mobile"] = rsp.OrderInfo.Mobile
	reMap["pay_type"] = rsp.OrderInfo.PayType
	reMap["order_id"] = rsp.OrderInfo.OrderId

	goodsList := make([]interface{}, 0)
	for _, item := range rsp.Data {
		tmpMap := gin.H{
			"id":    item.GoodsId,
			"name":  item.GoodsName,
			"price": item.GoodsPrice,
			"nums":  item.Nums,
		}

		goodsList = append(goodsList, tmpMap)
	}
	//todo 查询订单详情 同样的方式生成支付url
	reMap["goods"] = goodsList
	ctx.JSON(http.StatusOK, reMap)
}

func New(ctx *gin.Context) {
	orderForm := forms.CreateOrderForm{}
	if err := ctx.ShouldBindJSON(&orderForm); err != nil {
		api.HandleValidatorError(ctx, err)
	}
	userId, _ := ctx.Get("userId")
	rsp, err := global.OrderSrvClient.CreateOrder(context.WithValue(context.Background(), "ginContext", ctx), &proto.OrderRequest{
		UserId:  int32(userId.(uint)),
		Name:    orderForm.Name,
		Mobile:  orderForm.Mobile,
		Address: orderForm.Address,
		Post:    orderForm.Post,
	})
	zap.S().Info("新建订单：", rsp)
	if err != nil {
		zap.S().Error("新建订单失败", err)
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	//生成支付宝支付url
	client, err := alipay.New(global.ServerConfig.AliPayInfo.AppID, global.ServerConfig.AliPayInfo.PrivateKey, false)
	if err != nil {
		zap.S().Errorw("生成支付实例失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
	}
	err = client.LoadAliPayPublicKey(global.ServerConfig.AliPayInfo.AliPublicKey)
	if err != nil {
		zap.S().Errorw("加载支付宝公钥失败")
	}
	p := alipay.TradePagePay{}
	p.NotifyURL = global.ServerConfig.AliPayInfo.NotifyURL
	p.ReturnURL = global.ServerConfig.AliPayInfo.ReturnURL
	p.Subject = "天天特价-订单" + rsp.OrderId
	p.OutTradeNo = rsp.OrderId
	p.TotalAmount = strconv.FormatFloat(float64(rsp.Total), 'f', 2, 64)
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, err := client.TradePagePay(p)
	if err != nil {
		zap.S().Errorw("生成支付宝url失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"order_id":    rsp.OrderId,
		"total":       rsp.Total,
		"aplipay_url": url.String(),
	})
}
