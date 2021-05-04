package initialize

import (
	"fmt"
	"goodssrv/global"

	_ "github.com/go-sql-driver/mysql"

	"go.uber.org/zap"
	"xorm.io/xorm"
)

var err error

func InitDB() {
	mysqlConf := fmt.Sprintf("%s:%s@tcp(%s:%d)/goods?charset=utf8mb4&parseTime=True&loc=Local",
		global.ServerConfig.MysqlInfo.User,
		global.ServerConfig.MysqlInfo.Password,
		global.ServerConfig.MysqlInfo.Host,
		global.ServerConfig.MysqlInfo.Port)
	global.Engine, err = xorm.NewEngine("mysql", mysqlConf)
	zap.S().Info(mysqlConf)
	if err != nil {
		panic(err)
	}
}
