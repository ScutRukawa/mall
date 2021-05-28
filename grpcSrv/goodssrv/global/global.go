package global

import (
	"goodssrv/config"

	"github.com/go-redis/redis/v8"
	"xorm.io/xorm"
)

var (
	NacosConfig  *config.NacosConfig  = &config.NacosConfig{}
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	Engine       *xorm.Engine
	RedisCli     *redis.Client
)
