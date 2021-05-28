package initialize

import (
	"fmt"
	"goodssrv/global"

	"github.com/go-redis/redis/v8"
)

func InitRedis() {
	// builder := strings.Builder{}
	addr := fmt.Sprintf("%s:%d", global.ServerConfig.Redisinfo.Host, global.ServerConfig.Redisinfo.Port)
	cfg := redis.Options{Addr: addr}
	global.RedisCli = redis.NewClient(&cfg)
}
