package model

import "time"

type InventoryHistoryStatus int

const (
	Reback   InventoryHistoryStatus = 1
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
	Id             int                    `xorm:"pk autoincr BIGINT"`
	OrderId        string                 `xorm:"comment('订单id') varchar(50)"`
	OrderInvDetail []byte                 `xorm:"comment('订单详情') BLOB"`
	Status         InventoryHistoryStatus `xorm:"comment('is reback') TINYINT"`
	CreateTime     time.Time              `xorm:"not null created   comment('创建时间') DATETIME"`
	UpdateTime     time.Time              `xorm:"not null updated  comment('修改时间') DATETIME"`
}

type OrderMsgBody struct {
	OrderId string `json:"order_id"`
	UserId  int32  `json:"user_id"`
	Address string `json:"address"`
	Mobile  string `json:"mobile"`
	Post    string `json:"post"`
	Name    string `json:"name"`
}
type OrderInvDetial struct {
	GoodsNumsMap map[int32]interface{} `json:"goods_nums_map"`
}

func NewOrderInvDetial() *OrderInvDetial {
	goodsNumsMap := make(map[int32]interface{}, 2)
	return &OrderInvDetial{
		GoodsNumsMap: goodsNumsMap,
	}
}
