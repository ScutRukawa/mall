package model

import "time"

type InventoryHistoryStatus int

const (
	InStock  InventoryHistoryStatus = 1
	OutStock InventoryHistoryStatus = 2
)

type Inventory struct {
	Id         int32     `xorm:"pk autoincr comment('商品库存ID') BIGINT"`
	GoodsId    int32     `xorm:"not null UNIQUE comment('商品ID') BIGINT"`
	Num        int32     `xorm:"not null default 0 comment('库存数量') BIGINT"`
	Version    int32     `xorm:"not null version default 1 comment('商品版本') INT "`
	CreateTime time.Time `xorm:"not null created comment('创建时间') DATETIME"`
	UpdateTime time.Time `xorm:"not null updated comment('修改时间') DATETIME"`
}

type InventoryHistory struct {
	Id         int                    `xorm:"pk autoincr BIGINT"`
	UserId     int                    `xorm:"comment('商品id') UNIQUE BIGINT"`
	GoodsId    int                    `xorm:"comment('商品id') BIGINT"`
	Nums       int                    `xorm:"comment('操作数量') INT"`
	OrderId    int                    `xorm:"comment('订单id') index BIGINT"`
	Status     InventoryHistoryStatus `xorm:"comment('是否出库') TINYINT"`
	CreateTime time.Time              `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') DATETIME"`
	UpdateTime time.Time              `xorm:"not null default CURRENT_TIMESTAMP comment('修改时间') DATETIME"`
}
