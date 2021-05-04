package global

import (
	"online-mall-inventory/config"

	"xorm.io/xorm"
)

var (
	NacosConfig  *config.NacosConfig  = &config.NacosConfig{}
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	Engine       *xorm.Engine
)
