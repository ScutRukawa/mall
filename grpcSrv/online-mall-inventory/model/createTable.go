package model

import "online-mall-inventory/global"

func CreateTables() {
	inventory := Inventory{}
	inventoryHistory := InventoryHistory{}
	err := global.Engine.Sync2(&inventory, &inventoryHistory)
	if err != nil {
		panic(err)
	}
}
