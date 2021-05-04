package global

import (
	"goodssrv/config"

	"xorm.io/xorm"
)

var (
	NacosConfig  *config.NacosConfig  = &config.NacosConfig{}
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	Engine       *xorm.Engine
)
