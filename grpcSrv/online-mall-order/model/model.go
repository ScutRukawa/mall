package model

import "time"

type OrderStatus int32
type PayType int32

const (
	TRADE_SUCCESS  OrderStatus = 1
	TRADE_CLOSED   OrderStatus = 2
	WAIT_BUYER_PAY OrderStatus = 3
	TRADE_FINISHED OrderStatus = 4
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
	OrderId    string      `xorm:"unique index comment('订单号') varchar(40)"`
	PayType    int32       `xorm:"comment('支付方式') TINYINT"`
	Status     OrderStatus `xorm:"comment('订单状态') TINYINT"`
	TradeNo    string      `xorm:"unique index comment('交易号') varchar(100)"`
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
	OrderId    string  `xorm:"comment('订单号') varchar(40)"`
	GoodsName  string  `xorm:"comment('商品名称') varchar(30)"`
	GoodsPrice float32 `xorm:"comment('商品名称') float"`
	Nums       int32   `xorm:"comment('商品数量') INT"`
}
