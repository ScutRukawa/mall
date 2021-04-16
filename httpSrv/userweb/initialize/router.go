package initialize

import (
	"userweb/middlewares"
	router2 "userweb/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.Cors())
	ApiGroup := router.Group("/u/v1")
	router2.InitUserRouter(ApiGroup) //可以二次分组
	return router
}
