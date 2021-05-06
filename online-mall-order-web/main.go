package main

import (
	"fmt"
	"online-mall-order-web/global"
	"online-mall-order-web/initialize"
	"online-mall-order-web/utils"
	myvalidator "online-mall-order-web/validator"

	"github.com/gofrs/uuid"
	_ "github.com/nacos-group/nacos-sdk-go/common/constant"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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
	//5 初始化srv连接
	initialize.InitSrvConn()

	// 创建UUID
	u1 := uuid.Must(uuid.NewV4()).String()
	fmt.Printf("UUIDv4: %s\n", u1)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myvalidator.ValidatorMobile)
	} // todo translation

	zap.S().Debugf("启动服务器,端口：%d", global.ServerConfig.Port)
	go func() {
		err := router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port))
		if err != nil {
			zap.S().Panic("gin run failed", err)
		}
	}()

	utils.Register("127.0.0.1", global.ServerConfig.Port, global.ServerConfig.ServiceName, []string{"mall", "wei"}, u1)
	utils.OnExit(u1)

}
