package main

import (
	"fmt"
	"userweb/global"
	"userweb/initialize"
	myvalidator "userweb/validator"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
)

func main() {
	//1 初始化router
	router := initialize.Routers()
	//2 初始化logger
	initialize.InitLogger()
	//3 初始化配置文件
	initialize.InitConfig()
	//4 初始化翻译
	err := initialize.InitTrans("zh")
	if err != nil {
		panic(err)
	}
	//5 初始化srv用户连接
	initialize.InitServConn()

	//global.ServerConfig.Port, _ = utils.GetFreePort()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myvalidator.ValidatorMobile)
	} // todo translation

	//zap.S() zap.L()  全局logger suger
	zap.S().Debugf("启动服务器,端口：%d", global.ServerConfig.Port)
	err = router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port))
	if err != nil {
		zap.S().Panic("gin run failed", err)
	}

}
