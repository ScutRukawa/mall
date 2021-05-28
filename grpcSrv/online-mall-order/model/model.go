package model

import "time"

type PayType int32

type TradeStatus string

const (
	TradeStatusWaitBuyerPay TradeStatus = "WAIT_BUYER_PAY" //（交易创建，等待买家付款）
	TradeStatusClosed       TradeStatus = "TRADE_CLOSED"   //（未付款交易超时关闭，或支付完成后全额退款）
	TradeStatusSuccess      TradeStatus = "TRADE_SUCCESS"  //（交易支付成功）
	TradeStatusFinished     TradeStatus = "TRADE_FINISHED" //（交易结束，不可退款）
)
const (
	ALIPAY PayType = 1
)

type ShoppingCart struct {
	Id      int32 `xorm:"pk autoincr comment('自增ID') BIGINT"`
	UserId  int32 `xorm:"comment('用户id') BIGINT"`
	GoodsId int32 `xorm:"comment('商品id') BIGINT"`
	Nums    int32 `xorm:"comment('购买数量') INT"`
	Checked bool  `xorm:"default false comment('商品id')"`
}

type OrderInfo struct {
	Id         int32       `xorm:"pk autoincr comment('自增ID') BIGINT"`
	UserId     int32       `xorm:"comment('用户id') BIGINT"`
	OrderId    string      `xorm:"unique index comment('订单号') varchar(50)"`
	PayType    int32       `xorm:"comment('支付方式') TINYINT"`
	Status     TradeStatus `xorm:"comment('订单状态') varchar(20)"`
	TradeNo    string      `xorm:"index comment('交易号') varchar(100)"`
	OrderMount float32     `xorm:"default 0 comment('订单金额') DOUBLE"`
	PayTime    time.Time   `xorm:"comment('支付时间') datetime"`
	Address    string      `xorm:"comment('用户地址') varchar(100)"`
	Mobile     string      `xorm:"comment('用户号码') varchar(11)"`
	Post       string      `xorm:"comment('用户备注') varchar(100)"`
	Total      float32     `xorm:"comment('金额') FLOAT"`
	SignerName string      `xorm:"comment('签收人') varchar(20)"`
	SignerNum  string      `xorm:"comment('签收人电话') varchar(11)"`
}

type OrderGoods struct {
	Id         int32   `xorm:"pk autoincr comment('自增ID') BIGINT"`
	GoodsId    int32   `xorm:"comment('商品ID') BIGINT"`
	OrderId    string  `xorm:"comment('订单号') varchar(50)"`
	GoodsName  string  `xorm:"comment('商品名称') varchar(30)"`
	GoodsPrice float32 `xorm:"comment('商品名称') float"`
	Nums       int32   `xorm:"comment('商品数量') INT"`
}

