package pay

import (
	"context"
	"net/http"
	"online-mall-order-web/global"
	"online-mall-order-web/proto"

	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
	"go.uber.org/zap"
)

func Notify(ctx *gin.Context) {
	//支付宝回调通知
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
	noti, _ := client.GetTradeNotification(ctx.Request)
	if noti == nil {
		//zap.S().Info("交易状态为:", noti.TradeStatus)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "不合法的通知",
		})
		//返回给支付宝,此处可以不返回任何信息
	}

	_, err = global.OrderSrvClient.UpdateOrderStatus(context.Background(), &proto.OrderStatus{
		OrderId: noti.OutTradeNo,
		Status:  string(noti.TradeStatus),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	return
}
