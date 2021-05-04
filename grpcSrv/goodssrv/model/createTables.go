package model

import "goodssrv/global"

func CreateTables() {
	goods := Goods{}
	err := global.Engine.Sync2(&goods)
	if err != nil {
		panic(err)
	}
}
