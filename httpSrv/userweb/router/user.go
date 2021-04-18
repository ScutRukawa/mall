package router

import (
	"userweb/api"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("/user") //二次分组.Use(middlewares.JWTAuth())
	{
		// userRouter.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		userRouter.GET("list", api.GetUserList)
		userRouter.POST("pwd_login", api.PasswordLogin)

	}
}
