package initialize

import (
	"net/http"
	"online-mall-order-web/middlewares"
	router2 "online-mall-order-web/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})
	router.Use(middlewares.Cors())
	ApiGroup := router.Group("/u/v1")
	router2.InitOrderRouter(ApiGroup)
	router2.InitShopCartRouter(ApiGroup)
	return router
}
