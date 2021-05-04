package model

import "online-mall-order/global"

func CreateTables() {
	orderGoods := OrderGoods{}
	orderInfo := OrderInfo{}
	shoppingCart := ShoppingCart{}
	err := global.Engine.Sync2(&orderGoods, &orderInfo, &shoppingCart)
	if err != nil {
		panic(err)
	}
}
