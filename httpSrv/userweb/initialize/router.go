package initialize

import (
	"net/http"
	"userweb/middlewares"
	router2 "userweb/router"

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
	router2.InitUserRouter(ApiGroup)
	// router2.InitBaseRouter(ApiGroup)
	return router
}
