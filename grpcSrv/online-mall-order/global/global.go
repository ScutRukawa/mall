package global

import (
	"online-mall-order/config"
	"online-mall-order/proto"

	"xorm.io/xorm"
)

var (
	NacosConfig    *config.NacosConfig  = &config.NacosConfig{}
	ServerConfig   *config.ServerConfig = &config.ServerConfig{}
	Engine         *xorm.Engine
	GoodsSrvClient proto.GoodsClient
	InvSrvClient   proto.InventoryClient
)
